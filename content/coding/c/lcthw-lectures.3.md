+++
tags = ["c"]
title = "C Lecture - 3"
description = "Exercise 41~ 48"
+++

Author: Zed A. Shaw

All content comes from Zed's [Lecture Repository](https://github.com/zedshaw/learn-c-the-hard-way-lectures.git) and [Libraries Repository](https://github.com/zedshaw/liblcthw).  All credit goes to Zed.



### Exercise 41 Project devpkg
.\ex41\devpkg

.\ex41\devpkg\commands.c

```c
#include <apr_uri.h>
#include <apr_fnmatch.h>
#include <unistd.h>

#include "commands.h"
#include "dbg.h"
#include "bstrlib.h"
#include "db.h"
#include "shell.h"

int Command_depends(apr_pool_t * p, const char *path)
{
    FILE *in = NULL;
    bstring line = NULL;

    in = fopen(path, "r");
    check(in != NULL, "Failed to open downloaded depends: %s", path);

    for (line = bgets((bNgetc) fgetc, in, '\n');
            line != NULL; 
            line = bgets((bNgetc) fgetc, in, '\n')) 
    {
        btrimws(line);
        log_info("Processing depends: %s", bdata(line));
        int rc = Command_install(p, bdata(line), NULL, NULL, NULL);
        check(rc == 0, "Failed to install: %s", bdata(line));
        bdestroy(line);
    }

    fclose(in);
    return 0;

error:
    if (line) bdestroy(line);
    if (in) fclose(in);
    return -1;
}

int Command_fetch(apr_pool_t * p, const char *url, int fetch_only)
{
    apr_uri_t info = {.port = 0 };
    int rc = 0;
    const char *depends_file = NULL;
    apr_status_t rv = apr_uri_parse(p, url, &info);

    check(rv == APR_SUCCESS, "Failed to parse URL: %s", url);

    if (apr_fnmatch(GIT_PAT, info.path, 0) == APR_SUCCESS) {
        rc = Shell_exec(GIT_SH, "URL", url, NULL);
        check(rc == 0, "git failed.");
    } else if (apr_fnmatch(DEPEND_PAT, info.path, 0) == APR_SUCCESS) {
        check(!fetch_only, "No point in fetching a DEPENDS file.");

        if (info.scheme) {
            depends_file = DEPENDS_PATH;
            rc = Shell_exec(CURL_SH, "URL", url, "TARGET", depends_file,
                    NULL);
            check(rc == 0, "Curl failed.");
        } else {
            depends_file = info.path;
        }

        // recursively process the devpkg list
        log_info("Building according to DEPENDS: %s", url);
        rv = Command_depends(p, depends_file);
        check(rv == 0, "Failed to process the DEPENDS: %s", url);

        // this indicates that nothing needs to be done
        return 0;

    } else if (apr_fnmatch(TAR_GZ_PAT, info.path, 0) == APR_SUCCESS) {
        if (info.scheme) {
            rc = Shell_exec(CURL_SH,
                    "URL", url, "TARGET", TAR_GZ_SRC, NULL);
            check(rc == 0, "Failed to curl source: %s", url);
        }

        rv = apr_dir_make_recursive(BUILD_DIR,
                APR_UREAD | APR_UWRITE |
                APR_UEXECUTE, p);
        check(rv == APR_SUCCESS, "Failed to make directory %s",
                BUILD_DIR);

        rc = Shell_exec(TAR_SH, "FILE", TAR_GZ_SRC, NULL);
        check(rc == 0, "Failed to untar %s", TAR_GZ_SRC);
    } else if (apr_fnmatch(TAR_BZ2_PAT, info.path, 0) == APR_SUCCESS) {
        if (info.scheme) {
            rc = Shell_exec(CURL_SH, "URL", url, "TARGET", TAR_BZ2_SRC,
                    NULL);
            check(rc == 0, "Curl failed.");
        }

        apr_status_t rc = apr_dir_make_recursive(BUILD_DIR,
                APR_UREAD | APR_UWRITE
                | APR_UEXECUTE, p);

        check(rc == 0, "Failed to make directory %s", BUILD_DIR);
        rc = Shell_exec(TAR_SH, "FILE", TAR_BZ2_SRC, NULL);
        check(rc == 0, "Failed to untar %s", TAR_BZ2_SRC);
    } else {
        sentinel("Don't now how to handle %s", url);
    }

    // indicates that an install needs to actually run
    return 1;
error:
    return -1;
}

int Command_build(apr_pool_t * p, const char *url,
        const char *configure_opts, const char *make_opts,
        const char *install_opts)
{
    int rc = 0;

    check(access(BUILD_DIR, X_OK | R_OK | W_OK) == 0,
            "Build directory doesn't exist: %s", BUILD_DIR);

    // actually do an install
    if (access(CONFIG_SCRIPT, X_OK) == 0) {
        log_info("Has a configure script, running it.");
        rc = Shell_exec(CONFIGURE_SH, "OPTS", configure_opts, NULL);
        check(rc == 0, "Failed to configure.");
    }

    rc = Shell_exec(MAKE_SH, "OPTS", make_opts, NULL);
    check(rc == 0, "Failed to build.");

    rc = Shell_exec(INSTALL_SH,
            "TARGET", install_opts ? install_opts : "install",
            NULL);
    check(rc == 0, "Failed to install.");

    rc = Shell_exec(CLEANUP_SH, NULL);
    check(rc == 0, "Failed to cleanup after build.");

    rc = DB_update(url);
    check(rc == 0, "Failed to add this package to the database.");

    return 0;

error:
    return -1;
}

int Command_install(apr_pool_t * p, const char *url,
        const char *configure_opts, const char *make_opts,
        const char *install_opts)
{
    int rc = 0;
    check(Shell_exec(CLEANUP_SH, NULL) == 0,
            "Failed to cleanup before building.");

    rc = DB_find(url);
    check(rc != -1, "Error checking the install database.");

    if (rc == 1) {
        log_info("Package %s already installed.", url);
        return 0;
    }

    rc = Command_fetch(p, url, 0);

    if (rc == 1) {
        rc = Command_build(p, url, configure_opts, make_opts,
                install_opts);
        check(rc == 0, "Failed to build: %s", url);
    } else if (rc == 0) {
        // no install needed
        log_info("Depends successfully installed: %s", url);
    } else {
        // had an error
        sentinel("Install failed: %s", url);
    }

    Shell_exec(CLEANUP_SH, NULL);
    return 0;

error:
    Shell_exec(CLEANUP_SH, NULL);
    return -1;
}

```

.\ex41\devpkg\commands.h

```c
#ifndef _commands_h
#define _commands_h

#include <apr_pools.h>

#define DEPENDS_PATH "/tmp/DEPENDS"
#define TAR_GZ_SRC "/tmp/pkg-src.tar.gz"
#define TAR_BZ2_SRC "/tmp/pkg-src.tar.bz2"
#define BUILD_DIR "/tmp/pkg-build"
#define GIT_PAT "*.git"
#define DEPEND_PAT "*DEPENDS"
#define TAR_GZ_PAT "*.tar.gz"
#define TAR_BZ2_PAT "*.tar.bz2"
#define CONFIG_SCRIPT "/tmp/pkg-build/configure"

enum CommandType {
    COMMAND_NONE, COMMAND_INSTALL, COMMAND_LIST, COMMAND_FETCH,
    COMMAND_INIT, COMMAND_BUILD
};

int Command_fetch(apr_pool_t * p, const char *url, int fetch_only);

int Command_install(apr_pool_t * p, const char *url,
        const char *configure_opts, const char *make_opts,
        const char *install_opts);

int Command_depends(apr_pool_t * p, const char *path);

int Command_build(apr_pool_t * p, const char *url,
        const char *configure_opts, const char *make_opts,
        const char *install_opts);

#endif

```

.\ex41\devpkg\db.c

```c
#include <unistd.h>
#include <apr_errno.h>
#include <apr_file_io.h>

#include "db.h"
#include "bstrlib.h"
#include "dbg.h"

static FILE *DB_open(const char *path, const char *mode)
{
    return fopen(path, mode);
}

static void DB_close(FILE * db)
{
    fclose(db);
}

static bstring DB_load()
{
    FILE *db = NULL;
    bstring data = NULL;

    db = DB_open(DB_FILE, "r");
    check(db, "Failed to open database: %s", DB_FILE);

    data = bread((bNread) fread, db);
    check(data, "Failed to read from db file: %s", DB_FILE);

    DB_close(db);
    return data;

error:
    if (db)
        DB_close(db);
    if (data)
        bdestroy(data);
    return NULL;
}

int DB_update(const char *url)
{
    if (DB_find(url)) {
        log_info("Already recorded as installed: %s", url);
    }

    FILE *db = DB_open(DB_FILE, "a+");
    check(db, "Failed to open DB file: %s", DB_FILE);

    bstring line = bfromcstr(url);
    bconchar(line, '\n');
    int rc = fwrite(line->data, blength(line), 1, db);
    check(rc == 1, "Failed to append to the db.");

    return 0;
error:
    if (db)
        DB_close(db);
    return -1;
}

int DB_find(const char *url)
{
    bstring data = NULL;
    bstring line = bfromcstr(url);
    int res = -1;

    data = DB_load();
    check(data, "Failed to load: %s", DB_FILE);

    if (binstr(data, 0, line) == BSTR_ERR) {
        res = 0;
    } else {
        res = 1;
    }

error:			// fallthrough
    if (data)
        bdestroy(data);
    if (line)
        bdestroy(line);

    return res;
}

int DB_init()
{
    apr_pool_t *p = NULL;
    apr_pool_initialize();
    apr_pool_create(&p, NULL);

    if (access(DB_DIR, W_OK | X_OK) == -1) {
        apr_status_t rc = apr_dir_make_recursive(DB_DIR,
                APR_UREAD | APR_UWRITE
                | APR_UEXECUTE |
                APR_GREAD | APR_GWRITE
                | APR_GEXECUTE, p);
        check(rc == APR_SUCCESS, "Failed to make database dir: %s",
                DB_DIR);
    }

    if (access(DB_FILE, W_OK) == -1) {
        FILE *db = DB_open(DB_FILE, "w");
        check(db, "Cannot open database: %s", DB_FILE);
        DB_close(db);
    }

    apr_pool_destroy(p);
    return 0;

error:
    apr_pool_destroy(p);
    return -1;
}

int DB_list()
{
    bstring data = DB_load();
    check(data, "Failed to read load: %s", DB_FILE);

    printf("%s", bdata(data));
    bdestroy(data);
    return 0;

error:
    return -1;
}

```

.\ex41\devpkg\db.h

```c
#ifndef _db_h
#define _db_h

#define DB_FILE "/usr/local/.devpkg/db"
#define DB_DIR "/usr/local/.devpkg"

int DB_init();
int DB_list();
int DB_update(const char *url);
int DB_find(const char *url);

#endif

```

.\ex41\devpkg\dbg.h

```c
#ifndef __dbg_h__
#define __dbg_h__

#include <stdio.h>
#include <errno.h>
#include <string.h>

#ifdef NDEBUG
#define debug(M, ...)
#else
#define debug(M, ...) fprintf(stderr, "DEBUG %s:%d: " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)
#endif

#define clean_errno() (errno == 0 ? "None" : strerror(errno))

#define log_err(M, ...) fprintf(stderr,\
        "[ERROR] (%s:%d: errno: %s) " M "\n", __FILE__, __LINE__,\
        clean_errno(), ##__VA_ARGS__)

#define log_warn(M, ...) fprintf(stderr,\
        "[WARN] (%s:%d: errno: %s) " M "\n",\
        __FILE__, __LINE__, clean_errno(), ##__VA_ARGS__)

#define log_info(M, ...) fprintf(stderr, "[INFO] (%s:%d) " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)

#define check(A, M, ...) if(!(A)) {\
    log_err(M, ##__VA_ARGS__); errno=0; goto error; }

#define sentinel(M, ...)  { log_err(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#define check_mem(A) check((A), "Out of memory.")

#define check_debug(A, M, ...) if(!(A)) { debug(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#endif

```

.\ex41\devpkg\devpkg.c

```c
#include <stdio.h>
#include <apr_general.h>
#include <apr_getopt.h>
#include <apr_strings.h>
#include <apr_lib.h>

#include "dbg.h"
#include "db.h"
#include "commands.h"

int main(int argc, const char *argv[])
{
    apr_pool_t *p = NULL;
    apr_pool_initialize();
    apr_pool_create(&p, NULL);

    apr_getopt_t *opt;
    apr_status_t rv;

    char ch = '\0';
    const char *optarg = NULL;
    const char *config_opts = NULL;
    const char *install_opts = NULL;
    const char *make_opts = NULL;
    const char *url = NULL;
    enum CommandType request = COMMAND_NONE;

    rv = apr_getopt_init(&opt, p, argc, argv);

    while (apr_getopt(opt, "I:Lc:m:i:d:SF:B:", &ch, &optarg) ==
            APR_SUCCESS) {
        switch (ch) {
            case 'I':
                request = COMMAND_INSTALL;
                url = optarg;
                break;

            case 'L':
                request = COMMAND_LIST;
                break;

            case 'c':
                config_opts = optarg;
                break;

            case 'm':
                make_opts = optarg;
                break;

            case 'i':
                install_opts = optarg;
                break;

            case 'S':
                request = COMMAND_INIT;
                break;

            case 'F':
                request = COMMAND_FETCH;
                url = optarg;
                break;

            case 'B':
                request = COMMAND_BUILD;
                url = optarg;
                break;
        }
    }

    switch (request) {
        case COMMAND_INSTALL:
            check(url, "You must at least give a URL.");
            Command_install(p, url, config_opts, make_opts, install_opts);
            break;

        case COMMAND_LIST:
            DB_list();
            break;

        case COMMAND_FETCH:
            check(url != NULL, "You must give a URL.");
            Command_fetch(p, url, 1);
            log_info("Downloaded to %s and in /tmp/", BUILD_DIR);
            break;

        case COMMAND_BUILD:
            check(url, "You must at least give a URL.");
            Command_build(p, url, config_opts, make_opts, install_opts);
            break;

        case COMMAND_INIT:
            rv = DB_init();
            check(rv == 0, "Failed to make the database.");
            break;

        default:
            sentinel("Invalid command given.");
    }

    return 0;

error:
    return 1;
}

```

.\ex41\devpkg\shell.c

```c
#include "shell.h"
#include "dbg.h"
#include <stdarg.h>

int Shell_exec(Shell template, ...)
{
    apr_pool_t *p = NULL;
    int rc = -1;
    apr_status_t rv = APR_SUCCESS;
    va_list argp;
    const char *key = NULL;
    const char *arg = NULL;
    int i = 0;

    rv = apr_pool_create(&p, NULL);
    check(rv == APR_SUCCESS, "Failed to create pool.");

    va_start(argp, template);

    for (key = va_arg(argp, const char *);
            key != NULL; key = va_arg(argp, const char *)) {
        arg = va_arg(argp, const char *);

        for (i = 0; template.args[i] != NULL; i++) {
            if (strcmp(template.args[i], key) == 0) {
                template.args[i] = arg;
                break;		// found it
            }
        }
    }

    rc = Shell_run(p, &template);
    apr_pool_destroy(p);
    va_end(argp);
    return rc;

error:
    if (p) {
        apr_pool_destroy(p);
    }
    return rc;
}

int Shell_run(apr_pool_t * p, Shell * cmd)
{
    apr_procattr_t *attr;
    apr_status_t rv;
    apr_proc_t newproc;

    rv = apr_procattr_create(&attr, p);
    check(rv == APR_SUCCESS, "Failed to create proc attr.");

    rv = apr_procattr_io_set(attr, APR_NO_PIPE, APR_NO_PIPE,
            APR_NO_PIPE);
    check(rv == APR_SUCCESS, "Failed to set IO of command.");

    rv = apr_procattr_dir_set(attr, cmd->dir);
    check(rv == APR_SUCCESS, "Failed to set root to %s", cmd->dir);

    rv = apr_procattr_cmdtype_set(attr, APR_PROGRAM_PATH);
    check(rv == APR_SUCCESS, "Failed to set cmd type.");

    rv = apr_proc_create(&newproc, cmd->exe, cmd->args, NULL, attr, p);
    check(rv == APR_SUCCESS, "Failed to run command.");

    rv = apr_proc_wait(&newproc, &cmd->exit_code, &cmd->exit_why,
            APR_WAIT);
    check(rv == APR_CHILD_DONE, "Failed to wait.");

    check(cmd->exit_code == 0, "%s exited badly.", cmd->exe);
    check(cmd->exit_why == APR_PROC_EXIT, "%s was killed or crashed",
            cmd->exe);

    return 0;

error:
    return -1;
}

Shell CLEANUP_SH = {
    .exe = "rm",
    .dir = "/tmp",
    .args = {"rm", "-rf", "/tmp/pkg-build", "/tmp/pkg-src.tar.gz",
        "/tmp/pkg-src.tar.bz2", "/tmp/DEPENDS", NULL}
};

Shell GIT_SH = {
    .dir = "/tmp",
    .exe = "git",
    .args = {"git", "clone", "URL", "pkg-build", NULL}
};

Shell TAR_SH = {
    .dir = "/tmp/pkg-build",
    .exe = "tar",
    .args = {"tar", "-xzf", "FILE", "--strip-components", "1", NULL}
};

Shell CURL_SH = {
    .dir = "/tmp",
    .exe = "curl",
    .args = {"curl", "-L", "-o", "TARGET", "URL", NULL}
};

Shell CONFIGURE_SH = {
    .exe = "./configure",
    .dir = "/tmp/pkg-build",
    .args = {"configure", "OPTS", NULL}
    ,
};

Shell MAKE_SH = {
    .exe = "make",
    .dir = "/tmp/pkg-build",
    .args = {"make", "OPTS", NULL}
};

Shell INSTALL_SH = {
    .exe = "sudo",
    .dir = "/tmp/pkg-build",
    .args = {"sudo", "make", "TARGET", NULL}
};

```

.\ex41\devpkg\shell.h

```c
#ifndef _shell_h
#define _shell_h

#define MAX_COMMAND_ARGS 100

#include <apr_thread_proc.h>

typedef struct Shell {
    const char *dir;
    const char *exe;

    apr_procattr_t *attr;
    apr_proc_t proc;
    apr_exit_why_e exit_why;
    int exit_code;

    const char *args[MAX_COMMAND_ARGS];
} Shell;

int Shell_run(apr_pool_t * p, Shell * cmd);
int Shell_exec(Shell cmd, ...);

extern Shell CLEANUP_SH;
extern Shell GIT_SH;
extern Shell TAR_SH;
extern Shell CURL_SH;
extern Shell CONFIGURE_SH;
extern Shell MAKE_SH;
extern Shell INSTALL_SH;

#endif

```

.\ex41\ex41.1.sh

```bash
set -e

## go somewhere safe
cd /tmp

## get the source to base APR 1.5.2
curl -L -O http://archive.apache.org/dist/apr/apr-1.5.2.tar.gz

## extract it and go into the source
tar -xzvf apr-1.5.2.tar.gz
cd apr-1.5.2

## you need this on OSX Yosemite
touch libtoolT

## configure, make, make install
./configure
make
sudo make install

## reset and cleanup
cd /tmp
rm -rf apr-1.5.2 apr-1.5.2.tar.gz

## do the same with apr-util
curl -L -O http://archive.apache.org/dist/apr/apr-util-1.5.4.tar.gz

## extract
tar -xzvf apr-util-1.5.4.tar.gz
cd apr-util-1.5.4

## you need this on OSX Yosemite
touch libtoolT

## configure, make, make install
./configure --with-apr=/usr/local/apr
## you need that extra parameter to configure because
## apr-util can't really find it because...who knows.

make
sudo make install

#cleanup
cd /tmp
rm -rf apr-util-1.5.4* apr-1.5.2*

```

.\ex41\ex41.2.sh

```bash
mkdir devpkg
cd devpkg
touch README Makefile

```

.\ex41\devpkg

.\ex41\devpkg\commands.c

```c
#include <apr_uri.h>
#include <apr_fnmatch.h>
#include <unistd.h>

#include "commands.h"
#include "dbg.h"
#include "bstrlib.h"
#include "db.h"
#include "shell.h"

int Command_depends(apr_pool_t * p, const char *path)
{
    FILE *in = NULL;
    bstring line = NULL;

    in = fopen(path, "r");
    check(in != NULL, "Failed to open downloaded depends: %s", path);

    for (line = bgets((bNgetc) fgetc, in, '\n');
            line != NULL; 
            line = bgets((bNgetc) fgetc, in, '\n')) 
    {
        btrimws(line);
        log_info("Processing depends: %s", bdata(line));
        int rc = Command_install(p, bdata(line), NULL, NULL, NULL);
        check(rc == 0, "Failed to install: %s", bdata(line));
        bdestroy(line);
    }

    fclose(in);
    return 0;

error:
    if (line) bdestroy(line);
    if (in) fclose(in);
    return -1;
}

int Command_fetch(apr_pool_t * p, const char *url, int fetch_only)
{
    apr_uri_t info = {.port = 0 };
    int rc = 0;
    const char *depends_file = NULL;
    apr_status_t rv = apr_uri_parse(p, url, &info);

    check(rv == APR_SUCCESS, "Failed to parse URL: %s", url);

    if (apr_fnmatch(GIT_PAT, info.path, 0) == APR_SUCCESS) {
        rc = Shell_exec(GIT_SH, "URL", url, NULL);
        check(rc == 0, "git failed.");
    } else if (apr_fnmatch(DEPEND_PAT, info.path, 0) == APR_SUCCESS) {
        check(!fetch_only, "No point in fetching a DEPENDS file.");

        if (info.scheme) {
            depends_file = DEPENDS_PATH;
            rc = Shell_exec(CURL_SH, "URL", url, "TARGET", depends_file,
                    NULL);
            check(rc == 0, "Curl failed.");
        } else {
            depends_file = info.path;
        }

        // recursively process the devpkg list
        log_info("Building according to DEPENDS: %s", url);
        rv = Command_depends(p, depends_file);
        check(rv == 0, "Failed to process the DEPENDS: %s", url);

        // this indicates that nothing needs to be done
        return 0;

    } else if (apr_fnmatch(TAR_GZ_PAT, info.path, 0) == APR_SUCCESS) {
        if (info.scheme) {
            rc = Shell_exec(CURL_SH,
                    "URL", url, "TARGET", TAR_GZ_SRC, NULL);
            check(rc == 0, "Failed to curl source: %s", url);
        }

        rv = apr_dir_make_recursive(BUILD_DIR,
                APR_UREAD | APR_UWRITE |
                APR_UEXECUTE, p);
        check(rv == APR_SUCCESS, "Failed to make directory %s",
                BUILD_DIR);

        rc = Shell_exec(TAR_SH, "FILE", TAR_GZ_SRC, NULL);
        check(rc == 0, "Failed to untar %s", TAR_GZ_SRC);
    } else if (apr_fnmatch(TAR_BZ2_PAT, info.path, 0) == APR_SUCCESS) {
        if (info.scheme) {
            rc = Shell_exec(CURL_SH, "URL", url, "TARGET", TAR_BZ2_SRC,
                    NULL);
            check(rc == 0, "Curl failed.");
        }

        apr_status_t rc = apr_dir_make_recursive(BUILD_DIR,
                APR_UREAD | APR_UWRITE
                | APR_UEXECUTE, p);

        check(rc == 0, "Failed to make directory %s", BUILD_DIR);
        rc = Shell_exec(TAR_SH, "FILE", TAR_BZ2_SRC, NULL);
        check(rc == 0, "Failed to untar %s", TAR_BZ2_SRC);
    } else {
        sentinel("Don't now how to handle %s", url);
    }

    // indicates that an install needs to actually run
    return 1;
error:
    return -1;
}

int Command_build(apr_pool_t * p, const char *url,
        const char *configure_opts, const char *make_opts,
        const char *install_opts)
{
    int rc = 0;

    check(access(BUILD_DIR, X_OK | R_OK | W_OK) == 0,
            "Build directory doesn't exist: %s", BUILD_DIR);

    // actually do an install
    if (access(CONFIG_SCRIPT, X_OK) == 0) {
        log_info("Has a configure script, running it.");
        rc = Shell_exec(CONFIGURE_SH, "OPTS", configure_opts, NULL);
        check(rc == 0, "Failed to configure.");
    }

    rc = Shell_exec(MAKE_SH, "OPTS", make_opts, NULL);
    check(rc == 0, "Failed to build.");

    rc = Shell_exec(INSTALL_SH,
            "TARGET", install_opts ? install_opts : "install",
            NULL);
    check(rc == 0, "Failed to install.");

    rc = Shell_exec(CLEANUP_SH, NULL);
    check(rc == 0, "Failed to cleanup after build.");

    rc = DB_update(url);
    check(rc == 0, "Failed to add this package to the database.");

    return 0;

error:
    return -1;
}

int Command_install(apr_pool_t * p, const char *url,
        const char *configure_opts, const char *make_opts,
        const char *install_opts)
{
    int rc = 0;
    check(Shell_exec(CLEANUP_SH, NULL) == 0,
            "Failed to cleanup before building.");

    rc = DB_find(url);
    check(rc != -1, "Error checking the install database.");

    if (rc == 1) {
        log_info("Package %s already installed.", url);
        return 0;
    }

    rc = Command_fetch(p, url, 0);

    if (rc == 1) {
        rc = Command_build(p, url, configure_opts, make_opts,
                install_opts);
        check(rc == 0, "Failed to build: %s", url);
    } else if (rc == 0) {
        // no install needed
        log_info("Depends successfully installed: %s", url);
    } else {
        // had an error
        sentinel("Install failed: %s", url);
    }

    Shell_exec(CLEANUP_SH, NULL);
    return 0;

error:
    Shell_exec(CLEANUP_SH, NULL);
    return -1;
}

```

.\ex41\devpkg\commands.h

```c
#ifndef _commands_h
#define _commands_h

#include <apr_pools.h>

#define DEPENDS_PATH "/tmp/DEPENDS"
#define TAR_GZ_SRC "/tmp/pkg-src.tar.gz"
#define TAR_BZ2_SRC "/tmp/pkg-src.tar.bz2"
#define BUILD_DIR "/tmp/pkg-build"
#define GIT_PAT "*.git"
#define DEPEND_PAT "*DEPENDS"
#define TAR_GZ_PAT "*.tar.gz"
#define TAR_BZ2_PAT "*.tar.bz2"
#define CONFIG_SCRIPT "/tmp/pkg-build/configure"

enum CommandType {
    COMMAND_NONE, COMMAND_INSTALL, COMMAND_LIST, COMMAND_FETCH,
    COMMAND_INIT, COMMAND_BUILD
};

int Command_fetch(apr_pool_t * p, const char *url, int fetch_only);

int Command_install(apr_pool_t * p, const char *url,
        const char *configure_opts, const char *make_opts,
        const char *install_opts);

int Command_depends(apr_pool_t * p, const char *path);

int Command_build(apr_pool_t * p, const char *url,
        const char *configure_opts, const char *make_opts,
        const char *install_opts);

#endif

```

.\ex41\devpkg\db.c

```c
#include <unistd.h>
#include <apr_errno.h>
#include <apr_file_io.h>

#include "db.h"
#include "bstrlib.h"
#include "dbg.h"

static FILE *DB_open(const char *path, const char *mode)
{
    return fopen(path, mode);
}

static void DB_close(FILE * db)
{
    fclose(db);
}

static bstring DB_load()
{
    FILE *db = NULL;
    bstring data = NULL;

    db = DB_open(DB_FILE, "r");
    check(db, "Failed to open database: %s", DB_FILE);

    data = bread((bNread) fread, db);
    check(data, "Failed to read from db file: %s", DB_FILE);

    DB_close(db);
    return data;

error:
    if (db)
        DB_close(db);
    if (data)
        bdestroy(data);
    return NULL;
}

int DB_update(const char *url)
{
    if (DB_find(url)) {
        log_info("Already recorded as installed: %s", url);
    }

    FILE *db = DB_open(DB_FILE, "a+");
    check(db, "Failed to open DB file: %s", DB_FILE);

    bstring line = bfromcstr(url);
    bconchar(line, '\n');
    int rc = fwrite(line->data, blength(line), 1, db);
    check(rc == 1, "Failed to append to the db.");

    return 0;
error:
    if (db)
        DB_close(db);
    return -1;
}

int DB_find(const char *url)
{
    bstring data = NULL;
    bstring line = bfromcstr(url);
    int res = -1;

    data = DB_load();
    check(data, "Failed to load: %s", DB_FILE);

    if (binstr(data, 0, line) == BSTR_ERR) {
        res = 0;
    } else {
        res = 1;
    }

error:			// fallthrough
    if (data)
        bdestroy(data);
    if (line)
        bdestroy(line);

    return res;
}

int DB_init()
{
    apr_pool_t *p = NULL;
    apr_pool_initialize();
    apr_pool_create(&p, NULL);

    if (access(DB_DIR, W_OK | X_OK) == -1) {
        apr_status_t rc = apr_dir_make_recursive(DB_DIR,
                APR_UREAD | APR_UWRITE
                | APR_UEXECUTE |
                APR_GREAD | APR_GWRITE
                | APR_GEXECUTE, p);
        check(rc == APR_SUCCESS, "Failed to make database dir: %s",
                DB_DIR);
    }

    if (access(DB_FILE, W_OK) == -1) {
        FILE *db = DB_open(DB_FILE, "w");
        check(db, "Cannot open database: %s", DB_FILE);
        DB_close(db);
    }

    apr_pool_destroy(p);
    return 0;

error:
    apr_pool_destroy(p);
    return -1;
}

int DB_list()
{
    bstring data = DB_load();
    check(data, "Failed to read load: %s", DB_FILE);

    printf("%s", bdata(data));
    bdestroy(data);
    return 0;

error:
    return -1;
}

```

.\ex41\devpkg\db.h

```c
#ifndef _db_h
#define _db_h

#define DB_FILE "/usr/local/.devpkg/db"
#define DB_DIR "/usr/local/.devpkg"

int DB_init();
int DB_list();
int DB_update(const char *url);
int DB_find(const char *url);

#endif

```

.\ex41\devpkg\dbg.h

```c
#ifndef __dbg_h__
#define __dbg_h__

#include <stdio.h>
#include <errno.h>
#include <string.h>

#ifdef NDEBUG
#define debug(M, ...)
#else
#define debug(M, ...) fprintf(stderr, "DEBUG %s:%d: " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)
#endif

#define clean_errno() (errno == 0 ? "None" : strerror(errno))

#define log_err(M, ...) fprintf(stderr,\
        "[ERROR] (%s:%d: errno: %s) " M "\n", __FILE__, __LINE__,\
        clean_errno(), ##__VA_ARGS__)

#define log_warn(M, ...) fprintf(stderr,\
        "[WARN] (%s:%d: errno: %s) " M "\n",\
        __FILE__, __LINE__, clean_errno(), ##__VA_ARGS__)

#define log_info(M, ...) fprintf(stderr, "[INFO] (%s:%d) " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)

#define check(A, M, ...) if(!(A)) {\
    log_err(M, ##__VA_ARGS__); errno=0; goto error; }

#define sentinel(M, ...)  { log_err(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#define check_mem(A) check((A), "Out of memory.")

#define check_debug(A, M, ...) if(!(A)) { debug(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#endif

```

.\ex41\devpkg\devpkg.c

```c
#include <stdio.h>
#include <apr_general.h>
#include <apr_getopt.h>
#include <apr_strings.h>
#include <apr_lib.h>

#include "dbg.h"
#include "db.h"
#include "commands.h"

int main(int argc, const char *argv[])
{
    apr_pool_t *p = NULL;
    apr_pool_initialize();
    apr_pool_create(&p, NULL);

    apr_getopt_t *opt;
    apr_status_t rv;

    char ch = '\0';
    const char *optarg = NULL;
    const char *config_opts = NULL;
    const char *install_opts = NULL;
    const char *make_opts = NULL;
    const char *url = NULL;
    enum CommandType request = COMMAND_NONE;

    rv = apr_getopt_init(&opt, p, argc, argv);

    while (apr_getopt(opt, "I:Lc:m:i:d:SF:B:", &ch, &optarg) ==
            APR_SUCCESS) {
        switch (ch) {
            case 'I':
                request = COMMAND_INSTALL;
                url = optarg;
                break;

            case 'L':
                request = COMMAND_LIST;
                break;

            case 'c':
                config_opts = optarg;
                break;

            case 'm':
                make_opts = optarg;
                break;

            case 'i':
                install_opts = optarg;
                break;

            case 'S':
                request = COMMAND_INIT;
                break;

            case 'F':
                request = COMMAND_FETCH;
                url = optarg;
                break;

            case 'B':
                request = COMMAND_BUILD;
                url = optarg;
                break;
        }
    }

    switch (request) {
        case COMMAND_INSTALL:
            check(url, "You must at least give a URL.");
            Command_install(p, url, config_opts, make_opts, install_opts);
            break;

        case COMMAND_LIST:
            DB_list();
            break;

        case COMMAND_FETCH:
            check(url != NULL, "You must give a URL.");
            Command_fetch(p, url, 1);
            log_info("Downloaded to %s and in /tmp/", BUILD_DIR);
            break;

        case COMMAND_BUILD:
            check(url, "You must at least give a URL.");
            Command_build(p, url, config_opts, make_opts, install_opts);
            break;

        case COMMAND_INIT:
            rv = DB_init();
            check(rv == 0, "Failed to make the database.");
            break;

        default:
            sentinel("Invalid command given.");
    }

    return 0;

error:
    return 1;
}

```

.\ex41\devpkg\shell.c

```c
#include "shell.h"
#include "dbg.h"
#include <stdarg.h>

int Shell_exec(Shell template, ...)
{
    apr_pool_t *p = NULL;
    int rc = -1;
    apr_status_t rv = APR_SUCCESS;
    va_list argp;
    const char *key = NULL;
    const char *arg = NULL;
    int i = 0;

    rv = apr_pool_create(&p, NULL);
    check(rv == APR_SUCCESS, "Failed to create pool.");

    va_start(argp, template);

    for (key = va_arg(argp, const char *);
            key != NULL; key = va_arg(argp, const char *)) {
        arg = va_arg(argp, const char *);

        for (i = 0; template.args[i] != NULL; i++) {
            if (strcmp(template.args[i], key) == 0) {
                template.args[i] = arg;
                break;		// found it
            }
        }
    }

    rc = Shell_run(p, &template);
    apr_pool_destroy(p);
    va_end(argp);
    return rc;

error:
    if (p) {
        apr_pool_destroy(p);
    }
    return rc;
}

int Shell_run(apr_pool_t * p, Shell * cmd)
{
    apr_procattr_t *attr;
    apr_status_t rv;
    apr_proc_t newproc;

    rv = apr_procattr_create(&attr, p);
    check(rv == APR_SUCCESS, "Failed to create proc attr.");

    rv = apr_procattr_io_set(attr, APR_NO_PIPE, APR_NO_PIPE,
            APR_NO_PIPE);
    check(rv == APR_SUCCESS, "Failed to set IO of command.");

    rv = apr_procattr_dir_set(attr, cmd->dir);
    check(rv == APR_SUCCESS, "Failed to set root to %s", cmd->dir);

    rv = apr_procattr_cmdtype_set(attr, APR_PROGRAM_PATH);
    check(rv == APR_SUCCESS, "Failed to set cmd type.");

    rv = apr_proc_create(&newproc, cmd->exe, cmd->args, NULL, attr, p);
    check(rv == APR_SUCCESS, "Failed to run command.");

    rv = apr_proc_wait(&newproc, &cmd->exit_code, &cmd->exit_why,
            APR_WAIT);
    check(rv == APR_CHILD_DONE, "Failed to wait.");

    check(cmd->exit_code == 0, "%s exited badly.", cmd->exe);
    check(cmd->exit_why == APR_PROC_EXIT, "%s was killed or crashed",
            cmd->exe);

    return 0;

error:
    return -1;
}

Shell CLEANUP_SH = {
    .exe = "rm",
    .dir = "/tmp",
    .args = {"rm", "-rf", "/tmp/pkg-build", "/tmp/pkg-src.tar.gz",
        "/tmp/pkg-src.tar.bz2", "/tmp/DEPENDS", NULL}
};

Shell GIT_SH = {
    .dir = "/tmp",
    .exe = "git",
    .args = {"git", "clone", "URL", "pkg-build", NULL}
};

Shell TAR_SH = {
    .dir = "/tmp/pkg-build",
    .exe = "tar",
    .args = {"tar", "-xzf", "FILE", "--strip-components", "1", NULL}
};

Shell CURL_SH = {
    .dir = "/tmp",
    .exe = "curl",
    .args = {"curl", "-L", "-o", "TARGET", "URL", NULL}
};

Shell CONFIGURE_SH = {
    .exe = "./configure",
    .dir = "/tmp/pkg-build",
    .args = {"configure", "OPTS", NULL}
    ,
};

Shell MAKE_SH = {
    .exe = "make",
    .dir = "/tmp/pkg-build",
    .args = {"make", "OPTS", NULL}
};

Shell INSTALL_SH = {
    .exe = "sudo",
    .dir = "/tmp/pkg-build",
    .args = {"sudo", "make", "TARGET", NULL}
};

```

.\ex41\devpkg\shell.h

```c
#ifndef _shell_h
#define _shell_h

#define MAX_COMMAND_ARGS 100

#include <apr_thread_proc.h>

typedef struct Shell {
    const char *dir;
    const char *exe;

    apr_procattr_t *attr;
    apr_proc_t proc;
    apr_exit_why_e exit_why;
    int exit_code;

    const char *args[MAX_COMMAND_ARGS];
} Shell;

int Shell_run(apr_pool_t * p, Shell * cmd);
int Shell_exec(Shell cmd, ...);

extern Shell CLEANUP_SH;
extern Shell GIT_SH;
extern Shell TAR_SH;
extern Shell CURL_SH;
extern Shell CONFIGURE_SH;
extern Shell MAKE_SH;
extern Shell INSTALL_SH;

#endif

```

.\ex41\ex41.1.sh

```bash
set -e

## go somewhere safe
cd /tmp

## get the source to base APR 1.5.2
curl -L -O http://archive.apache.org/dist/apr/apr-1.5.2.tar.gz

## extract it and go into the source
tar -xzvf apr-1.5.2.tar.gz
cd apr-1.5.2

## you need this on OSX Yosemite
touch libtoolT

## configure, make, make install
./configure
make
sudo make install

## reset and cleanup
cd /tmp
rm -rf apr-1.5.2 apr-1.5.2.tar.gz

## do the same with apr-util
curl -L -O http://archive.apache.org/dist/apr/apr-util-1.5.4.tar.gz

## extract
tar -xzvf apr-util-1.5.4.tar.gz
cd apr-util-1.5.4

## you need this on OSX Yosemite
touch libtoolT

## configure, make, make install
./configure --with-apr=/usr/local/apr
## you need that extra parameter to configure because
## apr-util can't really find it because...who knows.

make
sudo make install

#cleanup
cd /tmp
rm -rf apr-util-1.5.4* apr-1.5.2*

```

.\ex41\ex41.2.sh

```bash
mkdir devpkg
cd devpkg
touch README Makefile

```

The Plan

Create a handy little tool called devpkg.

This will be a *lot* of work, so this video is more complete.

Demonstration

I'll demonstrate how devpkg works so you get a better idea.

Read the book's description as well for more details.

The Apache Portable Runtime

Review of the APR and installing it.

The Analysis

Walk through the code, where everything is, and what to watch out for.

Getting My Code

If you get stuck you can check out the learn-c-the-hard-way-lectures project:

And look in ex41/devpkg for the code.

Extra Credit

* Compare your code to my code available online.  Starting with 100%,
  remove 1% for each line you got wrong.
* Take the notes.txt file that you previously created and implement your improvements to the the code and functionality
  of ``devpkg``.
* Write an alternative version of ``devpkg`` using your other
  favorite language or the one you think can do this the best.  Compare
  the two, then improve your *C* version of ``devpkg`` based on what
  you've learned.

### Exercise 42 Stacks and Queues

The Plan

Create a Stack and Queue data structure from just the unit tests.

PAUSE!

WARNING!  Stop the video now and try to solve this yourself!

I'll show you how I did it after you try it (or you can cheat).

Code Review

.\ex42\queue_tests.c

```c
#include "minunit.h"
#include <lcthw/queue.h>
#include <assert.h>

static Queue *queue = NULL;
char *tests[] = { "test1 data", "test2 data", "test3 data" };

#define NUM_TESTS 3

char *test_create()
{
    queue = Queue_create();
    mu_assert(queue != NULL, "Failed to create queue.");

    return NULL;
}

char *test_destroy()
{
    mu_assert(queue != NULL, "Failed to make queue #2");
    Queue_destroy(queue);

    return NULL;
}

char *test_send_recv()
{
    int i = 0;
    for (i = 0; i < NUM_TESTS; i++) {
        Queue_send(queue, tests[i]);
        mu_assert(Queue_peek(queue) == tests[0], "Wrong next value.");
    }

    mu_assert(Queue_count(queue) == NUM_TESTS, "Wrong count on send.");

    QUEUE_FOREACH(queue, cur) {
        debug("VAL: %s", (char *)cur->value);
    }

    for (i = 0; i < NUM_TESTS; i++) {
        char *val = Queue_recv(queue);
        mu_assert(val == tests[i], "Wrong value on recv.");
    }

    mu_assert(Queue_count(queue) == 0, "Wrong count after recv.");

    return NULL;
}

char *all_tests()
{
    mu_suite_start();

    mu_run_test(test_create);
    mu_run_test(test_send_recv);
    mu_run_test(test_destroy);

    return NULL;
}

RUN_TESTS(all_tests);

```

.\ex42\stack_tests.c

```c
#include "minunit.h"
#include <lcthw/stack.h>
#include <assert.h>

static Stack *stack = NULL;
char *tests[] = { "test1 data", "test2 data", "test3 data" };

#define NUM_TESTS 3

char *test_create()
{
    stack = Stack_create();
    mu_assert(stack != NULL, "Failed to create stack.");

    return NULL;
}

char *test_destroy()
{
    mu_assert(stack != NULL, "Failed to make stack #2");
    Stack_destroy(stack);

    return NULL;
}

char *test_push_pop()
{
    int i = 0;
    for (i = 0; i < NUM_TESTS; i++) {
        Stack_push(stack, tests[i]);
        mu_assert(Stack_peek(stack) == tests[i], "Wrong next value.");
    }

    mu_assert(Stack_count(stack) == NUM_TESTS, "Wrong count on push.");

    STACK_FOREACH(stack, cur) {
        debug("VAL: %s", (char *)cur->value);
    }

    for (i = NUM_TESTS - 1; i >= 0; i--) {
        char *val = Stack_pop(stack);
        mu_assert(val == tests[i], "Wrong value on pop.");
    }

    mu_assert(Stack_count(stack) == 0, "Wrong count after pop.");

    return NULL;
}

char *all_tests()
{
    mu_suite_start();

    mu_run_test(test_create);
    mu_run_test(test_push_pop);
    mu_run_test(test_destroy);

    return NULL;
}

RUN_TESTS(all_tests);

```

Extra Credit

* Implement ``Stack`` using ``DArray`` instead of ``List``, but without changing the unit test.  That means you'll have to create your own ``STACK_FOREACH``.

### Exercise 43 A Simple Statistics Engine
.\ex43\stats.h

```c
#ifndef lcthw_stats_h
#define lcthw_stats_h

typedef struct Stats {
    double sum;
    double sumsq;
    unsigned long n;
    double min;
    double max;
} Stats;

Stats *Stats_recreate(double sum, double sumsq, unsigned long n,
        double min, double max);

Stats *Stats_create();

double Stats_mean(Stats * st);

double Stats_stddev(Stats * st);

void Stats_sample(Stats * st, double s);

void Stats_dump(Stats * st);

#endif

```

.\ex43\stats.c

```c
#include <math.h>
#include <lcthw/stats.h>
#include <stdlib.h>
#include <lcthw/dbg.h>

Stats *Stats_recreate(double sum, double sumsq, unsigned long n,
        double min, double max)
{
    Stats *st = malloc(sizeof(Stats));
    check_mem(st);

    st->sum = sum;
    st->sumsq = sumsq;
    st->n = n;
    st->min = min;
    st->max = max;

    return st;

error:
    return NULL;
}

Stats *Stats_create()
{
    return Stats_recreate(0.0, 0.0, 0L, 0.0, 0.0);
}

double Stats_mean(Stats * st)
{
    return st->sum / st->n;
}

double Stats_stddev(Stats * st)
{
    return sqrt((st->sumsq - (st->sum * st->sum / st->n)) /
            (st->n - 1));
}

void Stats_sample(Stats * st, double s)
{
    st->sum += s;
    st->sumsq += s * s;

    if (st->n == 0) {
        st->min = s;
        st->max = s;
    } else {
        if (st->min > s)
            st->min = s;
        if (st->max < s)
            st->max = s;
    }

    st->n += 1;
}

void Stats_dump(Stats * st)
{
    fprintf(stderr,
            "sum: %f, sumsq: %f, n: %ld, "
            "min: %f, max: %f, mean: %f, stddev: %f",
            st->sum, st->sumsq, st->n, st->min, st->max, Stats_mean(st),
            Stats_stddev(st));
}

```

.\ex43\stats_tests.c

```c
#include "minunit.h"
#include <lcthw/stats.h>
#include <math.h>

const int NUM_SAMPLES = 10;
double samples[] = {
    6.1061334, 9.6783204, 1.2747090, 8.2395131, 0.3333483,
    6.9755066, 1.0626275, 7.6587523, 4.9382973, 9.5788115
};

Stats expect = {
    .sumsq = 425.1641,
    .sum = 55.84602,
    .min = 0.333,
    .max = 9.678,
    .n = 10,
};

double expect_mean = 5.584602;
double expect_stddev = 3.547868;

#define EQ(X,Y,N) (round((X) * pow(10, N)) == round((Y) * pow(10, N)))

char *test_operations()
{
    int i = 0;
    Stats *st = Stats_create();
    mu_assert(st != NULL, "Failed to create stats.");

    for (i = 0; i < NUM_SAMPLES; i++) {
        Stats_sample(st, samples[i]);
    }

    Stats_dump(st);

    mu_assert(EQ(st->sumsq, expect.sumsq, 3), "sumsq not valid");
    mu_assert(EQ(st->sum, expect.sum, 3), "sum not valid");
    mu_assert(EQ(st->min, expect.min, 3), "min not valid");
    mu_assert(EQ(st->max, expect.max, 3), "max not valid");
    mu_assert(EQ(st->n, expect.n, 3), "max not valid");
    mu_assert(EQ(expect_mean, Stats_mean(st), 3), "mean not valid");
    mu_assert(EQ(expect_stddev, Stats_stddev(st), 3),
            "stddev not valid");

    return NULL;
}

char *test_recreate()
{
    Stats *st = Stats_recreate(
            expect.sum, expect.sumsq, expect.n, expect.min, expect.max);

    mu_assert(st->sum == expect.sum, "sum not equal");
    mu_assert(st->sumsq == expect.sumsq, "sumsq not equal");
    mu_assert(st->n == expect.n, "n not equal");
    mu_assert(st->min == expect.min, "min not equal");
    mu_assert(st->max == expect.max, "max not equal");
    mu_assert(EQ(expect_mean, Stats_mean(st), 3), "mean not valid");
    mu_assert(EQ(expect_stddev, Stats_stddev(st), 3),
            "stddev not valid");

    return NULL;
}

char *all_tests()
{
    mu_suite_start();

    mu_run_test(test_operations);
    mu_run_test(test_recreate);

    return NULL;
}

RUN_TESTS(all_tests);

```

.\ex43\stats.h

```c
#ifndef lcthw_stats_h
#define lcthw_stats_h

typedef struct Stats {
    double sum;
    double sumsq;
    unsigned long n;
    double min;
    double max;
} Stats;

Stats *Stats_recreate(double sum, double sumsq, unsigned long n,
        double min, double max);

Stats *Stats_create();

double Stats_mean(Stats * st);

double Stats_stddev(Stats * st);

void Stats_sample(Stats * st, double s);

void Stats_dump(Stats * st);

#endif

```

.\ex43\stats.c

```c
#include <math.h>
#include <lcthw/stats.h>
#include <stdlib.h>
#include <lcthw/dbg.h>

Stats *Stats_recreate(double sum, double sumsq, unsigned long n,
        double min, double max)
{
    Stats *st = malloc(sizeof(Stats));
    check_mem(st);

    st->sum = sum;
    st->sumsq = sumsq;
    st->n = n;
    st->min = min;
    st->max = max;

    return st;

error:
    return NULL;
}

Stats *Stats_create()
{
    return Stats_recreate(0.0, 0.0, 0L, 0.0, 0.0);
}

double Stats_mean(Stats * st)
{
    return st->sum / st->n;
}

double Stats_stddev(Stats * st)
{
    return sqrt((st->sumsq - (st->sum * st->sum / st->n)) /
            (st->n - 1));
}

void Stats_sample(Stats * st, double s)
{
    st->sum += s;
    st->sumsq += s * s;

    if (st->n == 0) {
        st->min = s;
        st->max = s;
    } else {
        if (st->min > s)
            st->min = s;
        if (st->max < s)
            st->max = s;
    }

    st->n += 1;
}

void Stats_dump(Stats * st)
{
    fprintf(stderr,
            "sum: %f, sumsq: %f, n: %ld, "
            "min: %f, max: %f, mean: %f, stddev: %f",
            st->sum, st->sumsq, st->n, st->min, st->max, Stats_mean(st),
            Stats_stddev(st));
}

```

.\ex43\stats_tests.c

```c
#include "minunit.h"
#include <lcthw/stats.h>
#include <math.h>

const int NUM_SAMPLES = 10;
double samples[] = {
    6.1061334, 9.6783204, 1.2747090, 8.2395131, 0.3333483,
    6.9755066, 1.0626275, 7.6587523, 4.9382973, 9.5788115
};

Stats expect = {
    .sumsq = 425.1641,
    .sum = 55.84602,
    .min = 0.333,
    .max = 9.678,
    .n = 10,
};

double expect_mean = 5.584602;
double expect_stddev = 3.547868;

#define EQ(X,Y,N) (round((X) * pow(10, N)) == round((Y) * pow(10, N)))

char *test_operations()
{
    int i = 0;
    Stats *st = Stats_create();
    mu_assert(st != NULL, "Failed to create stats.");

    for (i = 0; i < NUM_SAMPLES; i++) {
        Stats_sample(st, samples[i]);
    }

    Stats_dump(st);

    mu_assert(EQ(st->sumsq, expect.sumsq, 3), "sumsq not valid");
    mu_assert(EQ(st->sum, expect.sum, 3), "sum not valid");
    mu_assert(EQ(st->min, expect.min, 3), "min not valid");
    mu_assert(EQ(st->max, expect.max, 3), "max not valid");
    mu_assert(EQ(st->n, expect.n, 3), "max not valid");
    mu_assert(EQ(expect_mean, Stats_mean(st), 3), "mean not valid");
    mu_assert(EQ(expect_stddev, Stats_stddev(st), 3),
            "stddev not valid");

    return NULL;
}

char *test_recreate()
{
    Stats *st = Stats_recreate(
            expect.sum, expect.sumsq, expect.n, expect.min, expect.max);

    mu_assert(st->sum == expect.sum, "sum not equal");
    mu_assert(st->sumsq == expect.sumsq, "sumsq not equal");
    mu_assert(st->n == expect.n, "n not equal");
    mu_assert(st->min == expect.min, "min not equal");
    mu_assert(st->max == expect.max, "max not equal");
    mu_assert(EQ(expect_mean, Stats_mean(st), 3), "mean not valid");
    mu_assert(EQ(expect_stddev, Stats_stddev(st), 3),
            "stddev not valid");

    return NULL;
}

char *all_tests()
{
    mu_suite_start();

    mu_run_test(test_operations);
    mu_run_test(test_recreate);

    return NULL;
}

RUN_TESTS(all_tests);

```

The Plan

* A fun and handy little statistics engine for simple analysis.
* Comparing it to the same in R.

Comparing Test vs. R

I'll use R to show you how this works vs. normal calculations using all data.

Breaking It

Easiest way to break this is to just feed it bad data once then the whole
stream is broken.

Extra Credit

* Convert the ``Stats_stddev`` and ``Stats_mean`` to ``static inline`` functions in the ``stats.h`` file instead of in the ``stats.c`` file.
* Use this code to write a performance test of the ``string_algos_test.c``.
  Make it optional, and have it run the base test as a series of samples, and then report
  the results.
* Write a version of this in another programming language you know.  Confirm that this
  version is correct based on what I have here.

Extra Credit

* Write a little program that can take a file full of numbers and spit these statistics
  out for them.
* Make the program accept a table of data that has headers on one line, then all
  of the other numbers on lines after it are separated by any number of spaces.  Your program
  should then print out these statistics for each column by the header name.

### Exercise 44 Ring Buffer

The Plan

Learn about a handy data structure for I/O processing:

Ring Buffers

The Code

.\ex44\netclient.c

```c
#undef NDEBUG
#include <stdlib.h>
#include <sys/select.h>
#include <stdio.h>
#include <lcthw/ringbuffer.h>
#include <lcthw/dbg.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <sys/uio.h>
#include <arpa/inet.h>
#include <netdb.h>
#include <unistd.h>
#include <fcntl.h>

struct tagbstring NL = bsStatic("\n");
struct tagbstring CRLF = bsStatic("\r\n");

int nonblock(int fd)
{
    int flags = fcntl(fd, F_GETFL, 0);
    check(flags >= 0, "Invalid flags on nonblock.");

    int rc = fcntl(fd, F_SETFL, flags | O_NONBLOCK);
    check(rc == 0, "Can't set nonblocking.");

    return 0;
error:
    return -1;
}

int client_connect(char *host, char *port)
{
    int rc = 0;
    struct addrinfo *addr = NULL;

    rc = getaddrinfo(host, port, NULL, &addr);
    check(rc == 0, "Failed to lookup %s:%s", host, port);

    int sock = socket(AF_INET, SOCK_STREAM, 0);
    check(sock >= 0, "Cannot create a socket.");

    rc = connect(sock, addr->ai_addr, addr->ai_addrlen);
    check(rc == 0, "Connect failed.");

    rc = nonblock(sock);
    check(rc == 0, "Can't set nonblocking.");

    freeaddrinfo(addr);
    return sock;

error:
    freeaddrinfo(addr);
    return -1;
}

int read_some(RingBuffer * buffer, int fd, int is_socket)
{
    int rc = 0;

    if (RingBuffer_available_data(buffer) == 0) {
        buffer->start = buffer->end = 0;
    }

    if (is_socket) {
        rc = recv(fd, RingBuffer_starts_at(buffer),
                RingBuffer_available_space(buffer), 0);
    } else {
        rc = read(fd, RingBuffer_starts_at(buffer),
                RingBuffer_available_space(buffer));
    }

    check(rc >= 0, "Failed to read from fd: %d", fd);

    RingBuffer_commit_write(buffer, rc);

    return rc;

error:
    return -1;
}

int write_some(RingBuffer * buffer, int fd, int is_socket)
{
    int rc = 0;
    bstring data = RingBuffer_get_all(buffer);

    check(data != NULL, "Failed to get from the buffer.");
    check(bfindreplace(data, &NL, &CRLF, 0) == BSTR_OK,
            "Failed to replace NL.");

    if (is_socket) {
        rc = send(fd, bdata(data), blength(data), 0);
    } else {
        rc = write(fd, bdata(data), blength(data));
    }

    check(rc == blength(data), "Failed to write everything to fd: %d.",
            fd);
    bdestroy(data);

    return rc;

error:
    return -1;
}

int main(int argc, char *argv[])
{
    fd_set allreads;
    fd_set readmask;

    int socket = 0;
    int rc = 0;
    RingBuffer *in_rb = RingBuffer_create(1024 * 10);
    RingBuffer *sock_rb = RingBuffer_create(1024 * 10);

    check(argc == 3, "USAGE: netclient host port");

    socket = client_connect(argv[1], argv[2]);
    check(socket >= 0, "connect to %s:%s failed.", argv[1], argv[2]);

    FD_ZERO(&allreads);
    FD_SET(socket, &allreads);
    FD_SET(0, &allreads);

    while (1) {
        readmask = allreads;
        rc = select(socket + 1, &readmask, NULL, NULL, NULL);
        check(rc >= 0, "select failed.");

        if (FD_ISSET(0, &readmask)) {
            rc = read_some(in_rb, 0, 0);
            check_debug(rc != -1, "Failed to read from stdin.");
        }

        if (FD_ISSET(socket, &readmask)) {
            rc = read_some(sock_rb, socket, 0);
            check_debug(rc != -1, "Failed to read from socket.");
        }

        while (!RingBuffer_empty(sock_rb)) {
            rc = write_some(sock_rb, 1, 0);
            check_debug(rc != -1, "Failed to write to stdout.");
        }

        while (!RingBuffer_empty(in_rb)) {
            rc = write_some(in_rb, socket, 1);
            check_debug(rc != -1, "Failed to write to socket.");
        }
    }

    return 0;

error:
    return -1;
}

```

It's basically a DArray with dynamic start and end settings.
You can *also* use a Queue of bstrings to do almost the same thing.

Code Review

.\ex44\netclient.c

```c
#undef NDEBUG
#include <stdlib.h>
#include <sys/select.h>
#include <stdio.h>
#include <lcthw/ringbuffer.h>
#include <lcthw/dbg.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <sys/uio.h>
#include <arpa/inet.h>
#include <netdb.h>
#include <unistd.h>
#include <fcntl.h>

struct tagbstring NL = bsStatic("\n");
struct tagbstring CRLF = bsStatic("\r\n");

int nonblock(int fd)
{
    int flags = fcntl(fd, F_GETFL, 0);
    check(flags >= 0, "Invalid flags on nonblock.");

    int rc = fcntl(fd, F_SETFL, flags | O_NONBLOCK);
    check(rc == 0, "Can't set nonblocking.");

    return 0;
error:
    return -1;
}

int client_connect(char *host, char *port)
{
    int rc = 0;
    struct addrinfo *addr = NULL;

    rc = getaddrinfo(host, port, NULL, &addr);
    check(rc == 0, "Failed to lookup %s:%s", host, port);

    int sock = socket(AF_INET, SOCK_STREAM, 0);
    check(sock >= 0, "Cannot create a socket.");

    rc = connect(sock, addr->ai_addr, addr->ai_addrlen);
    check(rc == 0, "Connect failed.");

    rc = nonblock(sock);
    check(rc == 0, "Can't set nonblocking.");

    freeaddrinfo(addr);
    return sock;

error:
    freeaddrinfo(addr);
    return -1;
}

int read_some(RingBuffer * buffer, int fd, int is_socket)
{
    int rc = 0;

    if (RingBuffer_available_data(buffer) == 0) {
        buffer->start = buffer->end = 0;
    }

    if (is_socket) {
        rc = recv(fd, RingBuffer_starts_at(buffer),
                RingBuffer_available_space(buffer), 0);
    } else {
        rc = read(fd, RingBuffer_starts_at(buffer),
                RingBuffer_available_space(buffer));
    }

    check(rc >= 0, "Failed to read from fd: %d", fd);

    RingBuffer_commit_write(buffer, rc);

    return rc;

error:
    return -1;
}

int write_some(RingBuffer * buffer, int fd, int is_socket)
{
    int rc = 0;
    bstring data = RingBuffer_get_all(buffer);

    check(data != NULL, "Failed to get from the buffer.");
    check(bfindreplace(data, &NL, &CRLF, 0) == BSTR_OK,
            "Failed to replace NL.");

    if (is_socket) {
        rc = send(fd, bdata(data), blength(data), 0);
    } else {
        rc = write(fd, bdata(data), blength(data));
    }

    check(rc == blength(data), "Failed to write everything to fd: %d.",
            fd);
    bdestroy(data);

    return rc;

error:
    return -1;
}

int main(int argc, char *argv[])
{
    fd_set allreads;
    fd_set readmask;

    int socket = 0;
    int rc = 0;
    RingBuffer *in_rb = RingBuffer_create(1024 * 10);
    RingBuffer *sock_rb = RingBuffer_create(1024 * 10);

    check(argc == 3, "USAGE: netclient host port");

    socket = client_connect(argv[1], argv[2]);
    check(socket >= 0, "connect to %s:%s failed.", argv[1], argv[2]);

    FD_ZERO(&allreads);
    FD_SET(socket, &allreads);
    FD_SET(0, &allreads);

    while (1) {
        readmask = allreads;
        rc = select(socket + 1, &readmask, NULL, NULL, NULL);
        check(rc >= 0, "select failed.");

        if (FD_ISSET(0, &readmask)) {
            rc = read_some(in_rb, 0, 0);
            check_debug(rc != -1, "Failed to read from stdin.");
        }

        if (FD_ISSET(socket, &readmask)) {
            rc = read_some(sock_rb, socket, 0);
            check_debug(rc != -1, "Failed to read from socket.");
        }

        while (!RingBuffer_empty(sock_rb)) {
            rc = write_some(sock_rb, 1, 0);
            check_debug(rc != -1, "Failed to write to stdout.");
        }

        while (!RingBuffer_empty(in_rb)) {
            rc = write_some(in_rb, socket, 1);
            check_debug(rc != -1, "Failed to write to socket.");
        }
    }

    return 0;

error:
    return -1;
}

```

The Analysis

* Watch a ring buffer work in the debugger.
* Draw it visually to explore it.
* The purpose is to efficiently add and remove data when the amount added and removed is random.

Pause!

I will next review the unit test I wrote so if you want to attempt
solving it yourself then pause now.

The Unit Test

Here's my version of the unit test.

Breaking It

* The biggest mistake you'll make with a ring buffer is off-by-one errors.
* This is why the RingBuffer\_commit\_ and other macros exist.
* Another common mistake is to use it between threads, but that's a whole other book.

Extra Credit

* Create an alternative implementation of ``RingBuffer`` that uses
  the POSIX trick and a unit test for it.
* Add a performance comparison test to this unit test that compares the
  two versions by fuzzing them with random data and random read/write operations.
  Make sure that you set up this fuzzing so that the same operations are done
  to each version, and you can compare them between runs.

### Exercise 45 A Simple TCP/IP Client

The Plan

* Learn to use the *select* method and a RingBuffer to write a simple command line network client.

How select Works

Code Review

Improving It

These read functions are useful so I can put them in RingBuffer.

Extra Credit

* As I mentioned, there are quite a few functions you may not know, so
  look them up.  In fact, look them all up even if you think you know
  them.
* Go back through and add various defensive programming checks to
  the functions to improve them.
* Use the *getopt* function to allow the user
  the option *not* to translate *\n* to *\r\n*. This
  is only needed on protocols that require it for line endings, like HTTP.
  Sometimes you don't want the translation, so give the user the option.

### Exercise 46 Ternary Search Tree

The Plan

Learn about my favorite data structure ever:

Ternary Search Tree

The Code

.\ex46\tstree.h

```c
#ifndef _lcthw_TSTree_h
#define _lcthw_TSTree_h

#include <stdlib.h>
#include <lcthw/darray.h>

typedef struct TSTree {
    char splitchar;
    struct TSTree *low;
    struct TSTree *equal;
    struct TSTree *high;
    void *value;
} TSTree;

void *TSTree_search(TSTree * root, const char *key, size_t len);

void *TSTree_search_prefix(TSTree * root, const char *key, size_t len);

typedef void (*TSTree_traverse_cb) (void *value, void *data);

TSTree *TSTree_insert(TSTree * node, const char *key, size_t len,
        void *value);

void TSTree_traverse(TSTree * node, TSTree_traverse_cb cb, void *data);

void TSTree_destroy(TSTree * root);

#endif

```

.\ex46\tstree.c

```c
#include <stdlib.h>
#include <stdio.h>
#include <assert.h>
#include <lcthw/dbg.h>
#include <lcthw/tstree.h>

static inline TSTree *TSTree_insert_base(TSTree * root, TSTree * node,
        const char *key, size_t len,
        void *value)
{
    if (node == NULL) {
        node = (TSTree *) calloc(1, sizeof(TSTree));

        if (root == NULL) {
            root = node;
        }

        node->splitchar = *key;
    }

    if (*key < node->splitchar) {
        node->low = TSTree_insert_base(
                root, node->low, key, len, value);
    } else if (*key == node->splitchar) {
        if (len > 1) {
            node->equal = TSTree_insert_base(
                    root, node->equal, key + 1, len - 1, value);
        } else {
            assert(node->value == NULL && "Duplicate insert into tst.");
            node->value = value;
        }
    } else {
        node->high = TSTree_insert_base(
                root, node->high, key, len, value);
    }

    return node;
}

TSTree *TSTree_insert(TSTree * node, const char *key, size_t len,
        void *value)
{
    return TSTree_insert_base(node, node, key, len, value);
}

void *TSTree_search(TSTree * root, const char *key, size_t len)
{
    TSTree *node = root;
    size_t i = 0;

    while (i < len && node) {
        if (key[i] < node->splitchar) {
            node = node->low;
        } else if (key[i] == node->splitchar) {
            i++;
            if (i < len)
                node = node->equal;
        } else {
            node = node->high;
        }
    }

    if (node) {
        return node->value;
    } else {
        return NULL;
    }
}

void *TSTree_search_prefix(TSTree * root, const char *key, size_t len)
{
    if (len == 0)
        return NULL;

    TSTree *node = root;
    TSTree *last = NULL;
    size_t i = 0;

    while (i < len && node) {
        if (key[i] < node->splitchar) {
            node = node->low;
        } else if (key[i] == node->splitchar) {
            i++;
            if (i < len) {
                if (node->value)
                    last = node;
                node = node->equal;
            }
        } else {
            node = node->high;
        }
    }

    node = node ? node : last;

    // traverse until we find the first value in the equal chain
    // this is then the first node with this prefix
    while (node && !node->value) {
        node = node->equal;
    }

    return node ? node->value : NULL;
}

void TSTree_traverse(TSTree * node, TSTree_traverse_cb cb, void *data)
{
    if (!node)
        return;

    if (node->low)
        TSTree_traverse(node->low, cb, data);

    if (node->equal) {
        TSTree_traverse(node->equal, cb, data);
    }

    if (node->high)
        TSTree_traverse(node->high, cb, data);

    if (node->value)
        cb(node->value, data);
}

void TSTree_destroy(TSTree * node)
{
    if (node == NULL)
        return;

    if (node->low)
        TSTree_destroy(node->low);

    if (node->equal) {
        TSTree_destroy(node->equal);
    }

    if (node->high)
        TSTree_destroy(node->high);

    free(node);
}

```

.\ex46\tstree_tests.c

```c
#include "minunit.h"
#include <lcthw/tstree.h>
#include <string.h>
#include <assert.h>
#include <lcthw/bstrlib.h>

TSTree *node = NULL;
char *valueA = "VALUEA";
char *valueB = "VALUEB";
char *value2 = "VALUE2";
char *value4 = "VALUE4";
char *reverse = "VALUER";
int traverse_count = 0;

struct tagbstring test1 = bsStatic("TEST");
struct tagbstring test2 = bsStatic("TEST2");
struct tagbstring test3 = bsStatic("TSET");
struct tagbstring test4 = bsStatic("T");

char *test_insert()
{
    node = TSTree_insert(node, bdata(&test1), blength(&test1), valueA);
    mu_assert(node != NULL, "Failed to insert into tst.");

    node = TSTree_insert(node, bdata(&test2), blength(&test2), value2);
    mu_assert(node != NULL,
            "Failed to insert into tst with second name.");

    node = TSTree_insert(node, bdata(&test3), blength(&test3), reverse);
    mu_assert(node != NULL,
            "Failed to insert into tst with reverse name.");

    node = TSTree_insert(node, bdata(&test4), blength(&test4), value4);
    mu_assert(node != NULL,
            "Failed to insert into tst with second name.");

    return NULL;
}

char *test_search_exact()
{
    // tst returns the last one inserted
    void *res = TSTree_search(node, bdata(&test1), blength(&test1));
    mu_assert(res == valueA,
            "Got the wrong value back, should get A not B.");

    // tst does not find if not exact
    res = TSTree_search(node, "TESTNO", strlen("TESTNO"));
    mu_assert(res == NULL, "Should not find anything.");

    return NULL;
}

char *test_search_prefix()
{
    void *res = TSTree_search_prefix(
            node, bdata(&test1), blength(&test1));
    debug("result: %p, expected: %p", res, valueA);
    mu_assert(res == valueA, "Got wrong valueA by prefix.");

    res = TSTree_search_prefix(node, bdata(&test1), 1);
    debug("result: %p, expected: %p", res, valueA);
    mu_assert(res == value4, "Got wrong value4 for prefix of 1.");

    res = TSTree_search_prefix(node, "TE", strlen("TE"));
    mu_assert(res != NULL, "Should find for short prefix.");

    res = TSTree_search_prefix(node, "TE--", strlen("TE--"));
    mu_assert(res != NULL, "Should find for partial prefix.");

    return NULL;
}

void TSTree_traverse_test_cb(void *value, void *data)
{
    assert(value != NULL && "Should not get NULL value.");
    assert(data == valueA && "Expecting valueA as the data.");
    traverse_count++;
}

char *test_traverse()
{
    traverse_count = 0;
    TSTree_traverse(node, TSTree_traverse_test_cb, valueA);
    debug("traverse count is: %d", traverse_count);
    mu_assert(traverse_count == 4, "Didn't find 4 keys.");

    return NULL;
}

char *test_destroy()
{
    TSTree_destroy(node);

    return NULL;
}

char *all_tests()
{
    mu_suite_start();

    mu_run_test(test_insert);
    mu_run_test(test_search_exact);
    mu_run_test(test_search_prefix);
    mu_run_test(test_traverse);
    mu_run_test(test_destroy);

    return NULL;
}

RUN_TESTS(all_tests);

```

Similar to a Binary Search Tree, but it has 3 branches per node based on
the characters in strings.

Advantages

* Find any string comparing at most N characters.
* Detect *missing* strings as fast, usually faster.
* Find all strings that start with, or contain, any substring as fast.
* Find all similar known strings quickly.

Disadvantages

* Delete is a pain, as in most trees.
* Uses lots of memory to store keys, so bad for sets of large keys.
* Kind of weird for most programmers.

Improving It

* You could allow duplicates by using a *DArray* instead of the
  *value*.
* As I mentioned earlier, deleting is hard, but you could simulate it by setting
  the values to *NULL* so that they are effectively gone.
* There are no ways to collect all of the possible matching values.  I'll have
  you implement that in an extra credit.
* There are other algorithms that are more complex but have slightly
  better properties.  Take a look at suffix array, suffix tree, and
  radix tree structures.

Extra Credit

* Implement a *TSTree_collect* that returns a *DArray* containing
  all of the keys that match the given prefix.
* Implement *TSTree_search_suffix* and a *TSTree_insert_suffix*
  so you can do suffix searches and inserts.
* Use the debugger to see how this structure is used in memory 
  compared to the *BSTree* and *Hashmap*.

### Exercise 47 A Fast URL Router

The Plan

Use the *TSTree* to do something useful:

Route URLs

.\ex47\ex47_urls.txt

```
/test.tst TestHandler
/ IndexHandler
/test/this/out/index.html PageHandler
/index.html PageHandler
/and/then/i/have/things/to/test.html PageHandler

```

Code Review

.\ex47\urlor.c

```c
#include <lcthw/tstree.h>
#include <lcthw/bstrlib.h>

TSTree *add_route_data(TSTree * routes, bstring line)
{
    struct bstrList *data = bsplit(line, ' ');
    check(data->qty == 2, "Line '%s' does not have 2 columns",
            bdata(line));

    routes = TSTree_insert(routes,
            bdata(data->entry[0]),
            blength(data->entry[0]),
            bstrcpy(data->entry[1]));

    bstrListDestroy(data);

    return routes;

error:
    return NULL;
}

TSTree *load_routes(const char *file)
{
    TSTree *routes = NULL;
    bstring line = NULL;
    FILE *routes_map = NULL;

    routes_map = fopen(file, "r");
    check(routes_map != NULL, "Failed to open routes: %s", file);

    while ((line = bgets((bNgetc) fgetc, routes_map, '\n')) != NULL) {
        check(btrimws(line) == BSTR_OK, "Failed to trim line.");
        routes = add_route_data(routes, line);
        check(routes != NULL, "Failed to add route.");
        bdestroy(line);
    }

    fclose(routes_map);
    return routes;

error:
    if (routes_map) fclose(routes_map);
    if (line) bdestroy(line);

    return NULL;
}

bstring match_url(TSTree * routes, bstring url)
{
    bstring route = TSTree_search(routes, bdata(url), blength(url));

    if (route == NULL) {
        printf("No exact match found, trying prefix.\n");
        route = TSTree_search_prefix(routes, bdata(url), blength(url));
    }

    return route;
}

bstring read_line(const char *prompt)
{
    printf("%s", prompt);

    bstring result = bgets((bNgetc) fgetc, stdin, '\n');
    check_debug(result != NULL, "stdin closed.");

    check(btrimws(result) == BSTR_OK, "Failed to trim.");

    return result;

error:
    return NULL;
}

void bdestroy_cb(void *value, void *ignored)
{
    (void)ignored;
    bdestroy((bstring) value);
}

void destroy_routes(TSTree * routes)
{
    TSTree_traverse(routes, bdestroy_cb, NULL);
    TSTree_destroy(routes);
}

int main(int argc, char *argv[])
{
    bstring url = NULL;
    bstring route = NULL;
    TSTree *routes = NULL;

    check(argc == 2, "USAGE: urlor <urlfile>");

    routes = load_routes(argv[1]);
    check(routes != NULL, "Your route file has an error.");

    while (1) {
        url = read_line("URL> ");
        check_debug(url != NULL, "goodbye.");

        route = match_url(routes, url);

        if (route) {
            printf("MATCH: %s == %s\n", bdata(url), bdata(route));
        } else {
            printf("FAIL: %s\n", bdata(url));
        }

        bdestroy(url);
    }

    destroy_routes(routes);
    return 0;

error:
    destroy_routes(routes);
    return 1;
}

```

The Analysis

Watch me play with it and then tell you how it's working.

Improving It

* Collect all possible matches then choose the longest as winner.
* Use TSTree to find prefixes, then regex to choose winner.

Extra Credit

* Instead of just storing the string for the handler, create an actual engine that uses a
  *Handler* struct to store the application.  The structure would store the URL to which it's attached, the name, and anything else you'd need to make an actual routing system.

Extra Credit

* Instead of mapping URLs to arbitrary names, map them to .so files and use the *dlopen*
  system to load handlers on the fly and call callbacks they contain.  Put these callbacks that
  in your *Handler* struct, and then you have yourself a fully dynamic callback
  handler system in C.

### Exercise 48a A Simple Network Server:

Project Description

The Plan

Start your first long running project:

statserve

The Purpose

You'll get the project started and get a minimum first hack going.

The Requirements

1. Create a simple network server that accepts a connection on port 7899 from
   *netclient* or the *nc* command, and echoes back anything you type.
2. You'll need to learn how to bind a port, listen on the socket, and answer it.
   Use your research skills to study how this is done and attempt to implement it
   yourself.

The Requirements

3. The more important part of this project is laying out the project directory
   from the *c-skeleton*, and making sure you can build everything and get it
   working.
4. Don't worry about things like daemons or anything else.  Your server just has
   to run from the command line and keep running.

The Clues

I will now give you some clues:

* USE liblcthw!
* Remember you did a client already, you just need to make a server.
* Do NOT use select! Use fork() for the server.
* Keep it *simple*.  Don't worry about anything other than accepting a connection and closing.
* Stay small, build slowly.

Important References

* Research online for "echo server in C".
* Read man (2) pages for *accept*, *bind*, *listen*, *connect*, *select*, *socket*, and *shutdown*.

Encouragement

This will be *HARD*!  Try it your best, and take it piece by piece.  You can do it, but remember if you give up the next video (48b) will show you the code to my solution and how to solve it.  You can peek there then come back when you're stuck.

### Exercise 48b A Simple Network Server:
.\ex48b\c-skeleton

.\ex48b\c-skeleton\src\dbg.h

```c
#ifndef __dbg_h__
#define __dbg_h__

#include <stdio.h>
#include <errno.h>
#include <string.h>

#ifdef NDEBUG
#define debug(M, ...)
#else
#define debug(M, ...) fprintf(stderr, "DEBUG %s:%d: " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)
#endif

#define clean_errno() (errno == 0 ? "None" : strerror(errno))

#define log_err(M, ...) fprintf(stderr,\
        "[ERROR] (%s:%d: errno: %s) " M "\n", __FILE__, __LINE__,\
        clean_errno(), ##__VA_ARGS__)

#define log_warn(M, ...) fprintf(stderr,\
        "[WARN] (%s:%d: errno: %s) " M "\n",\
        __FILE__, __LINE__, clean_errno(), ##__VA_ARGS__)

#define log_info(M, ...) fprintf(stderr, "[INFO] (%s:%d) " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)

#define check(A, M, ...) if(!(A)) {\
    log_err(M, ##__VA_ARGS__); errno=0; goto error; }

#define sentinel(M, ...)  { log_err(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#define check_mem(A) check((A), "Out of memory.")

#define check_debug(A, M, ...) if(!(A)) { debug(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#endif

```

.\ex48b\c-skeleton\src\libex29.c

```c
#include <stdio.h>
#include <ctype.h>
#include "dbg.h"

int print_a_message(const char *msg)
{
    printf("A STRING: %s\n", msg);

    return 0;
}

int uppercase(const char *msg)
{
    int i = 0;

    // BUG: \0 termination problems
    for(i = 0; msg[i] != '\0'; i++) {
        printf("%c", toupper(msg[i]));
    }

    printf("\n");

    return 0;
}

int lowercase(const char *msg)
{
    int i = 0;

    // BUG: \0 termination problems
    for(i = 0; msg[i] != '\0'; i++) {
        printf("%c", tolower(msg[i]));
    }

    printf("\n");

    return 0;
}

int fail_on_purpose(const char *msg)
{
    return 1;
} 

```

.\ex48b\c-skeleton\tests\libex29_tests.c

```c
#include "minunit.h"
#include <dlfcn.h>

typedef int (*lib_function) (const char *data);
char *lib_file = "build/libYOUR_LIBRARY.so";
void *lib = NULL;

int check_function(const char *func_to_run, const char *data,
        int expected)
{
    lib_function func = dlsym(lib, func_to_run);
    check(func != NULL,
            "Did not find %s function in the library %s: %s", func_to_run,
            lib_file, dlerror());

    int rc = func(data);
    check(rc == expected, "Function %s return %d for data: %s",
            func_to_run, rc, data);

    return 1;
error:
    return 0;
}

char *test_dlopen()
{
    lib = dlopen(lib_file, RTLD_NOW);
    mu_assert(lib != NULL, "Failed to open the library to test.");

    return NULL;
}

char *test_functions()
{
    mu_assert(check_function("print_a_message", "Hello", 0),
            "print_a_message failed.");
    mu_assert(check_function("uppercase", "Hello", 0),
            "uppercase failed.");
    mu_assert(check_function("lowercase", "Hello", 0),
            "lowercase failed.");

    return NULL;
}

char *test_failures()
{
    mu_assert(check_function("fail_on_purpose", "Hello", 1),
            "fail_on_purpose should fail.");

    return NULL;
}

char *test_dlclose()
{
    int rc = dlclose(lib);
    mu_assert(rc == 0, "Failed to close lib.");

    return NULL;
}

char *all_tests()
{
    mu_suite_start();

    mu_run_test(test_dlopen);
    mu_run_test(test_functions);
    mu_run_test(test_failures);
    mu_run_test(test_dlclose);

    return NULL;
}

RUN_TESTS(all_tests);

```

.\ex48b\statserve

.\ex48b\statserve\bin\statserve.c

```c
#include <stdio.h>
#include <lcthw/dbg.h>
#include "statserve.h"
#include "net.h"

int main(int argc, char *argv[])
{
    check(argc == 3, "USAGE: statserve host port");

    const char *host = argv[1];
    const char *port = argv[2];

    check(echo_server(host, port), "Failed to run the echo server.");

    return 0;

error:

    return 1;
}

```

.\ex48b\statserve\src\dbg.h

```c
#ifndef __dbg_h__
#define __dbg_h__

#include <stdio.h>
#include <errno.h>
#include <string.h>

#ifdef NDEBUG
#define debug(M, ...)
#else
#define debug(M, ...) fprintf(stderr, "DEBUG %s:%d: " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)
#endif

#define clean_errno() (errno == 0 ? "None" : strerror(errno))

#define log_err(M, ...) fprintf(stderr,\
        "[ERROR] (%s:%d: errno: %s) " M "\n", __FILE__, __LINE__,\
        clean_errno(), ##__VA_ARGS__)

#define log_warn(M, ...) fprintf(stderr,\
        "[WARN] (%s:%d: errno: %s) " M "\n",\
        __FILE__, __LINE__, clean_errno(), ##__VA_ARGS__)

#define log_info(M, ...) fprintf(stderr, "[INFO] (%s:%d) " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)

#define check(A, M, ...) if(!(A)) {\
    log_err(M, ##__VA_ARGS__); errno=0; goto error; }

#define sentinel(M, ...)  { log_err(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#define check_mem(A) check((A), "Out of memory.")

#define check_debug(A, M, ...) if(!(A)) { debug(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#endif

```

.\ex48b\statserve\src\net.c

```c
#include <stdlib.h>
#include <sys/select.h>
#include <stdio.h>
#include <lcthw/ringbuffer.h>
#include <lcthw/dbg.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <sys/uio.h>
#include <arpa/inet.h>
#include <netdb.h>
#include <unistd.h>
#include <fcntl.h>
#include "net.h"

struct tagbstring NL = bsStatic("\n");
struct tagbstring CRLF = bsStatic("\r\n");

int nonblock(int fd)
{
    int flags = fcntl(fd, F_GETFL, 0);
    check(flags >= 0, "Invalid flags on nonblock.");

    int rc = fcntl(fd, F_SETFL, flags | O_NONBLOCK);
    check(rc == 0, "Can't set nonblocking.");

    return 0;
error:
    return -1;
}

int client_connect(char *host, char *port)
{
    int rc = 0;
    struct addrinfo *addr = NULL;

    rc = getaddrinfo(host, port, NULL, &addr);
    check(rc == 0, "Failed to lookup %s:%s", host, port);

    int sock = socket(AF_INET, SOCK_STREAM, 0);
    check(sock >= 0, "Cannot create a socket.");

    rc = connect(sock, addr->ai_addr, addr->ai_addrlen);
    check(rc == 0, "Connect failed.");

    rc = nonblock(sock);
    check(rc == 0, "Can't set nonblocking.");

    freeaddrinfo(addr);
    return sock;

error:
    freeaddrinfo(addr);
    return -1;
}

int read_some(RingBuffer * buffer, int fd, int is_socket)
{
    int rc = 0;

    if (RingBuffer_available_data(buffer) == 0) {
        buffer->start = buffer->end = 0;
    }

    if (is_socket) {
        rc = recv(fd, RingBuffer_starts_at(buffer),
                RingBuffer_available_space(buffer), 0);
    } else {
        rc = read(fd, RingBuffer_starts_at(buffer),
                RingBuffer_available_space(buffer));
    }

    check(rc >= 0, "Failed to read from fd: %d", fd);

    RingBuffer_commit_write(buffer, rc);

    return rc;

error:
    return -1;
}

int write_some(RingBuffer * buffer, int fd, int is_socket)
{
    int rc = 0;
    bstring data = RingBuffer_get_all(buffer);

    check(data != NULL, "Failed to get from the buffer.");
    check(bfindreplace(data, &NL, &CRLF, 0) == BSTR_OK,
            "Failed to replace NL.");

    if (is_socket) {
        rc = send(fd, bdata(data), blength(data), 0);
    } else {
        rc = write(fd, bdata(data), blength(data));
    }

    check(rc == blength(data), "Failed to write everything to fd: %d.",
            fd);
    bdestroy(data);

    return rc;

error:
    return -1;
}

int attempt_listen(struct addrinfo *info)
{
    int sockfd = -1; // default fail
    int rc = -1;
    int yes = 1;

    check(info != NULL, "Invalid addrinfo.");

    // create a socket with the addrinfo
    sockfd = socket(info->ai_family, info->ai_socktype,
            info->ai_protocol);
    check_debug(sockfd != -1, "Failed to bind to address. Trying more.");

    // set the SO_REUSEADDR option on the socket
    rc = setsockopt(sockfd, SOL_SOCKET, SO_REUSEADDR, &yes, sizeof(int));
    check_debug(rc == 0, "Failed to set SO_REUSADDR.");

    // attempt to bind to it
    rc = bind(sockfd, info->ai_addr, info->ai_addrlen);
    check_debug(rc == 0, "Failed to find socket.");

    // finally listen with a backlog
    rc = listen(sockfd, BACKLOG);
    check_debug(rc == 0, "Failed to listen to socket.");

    return sockfd;

error:
    return -1;
}

int server_listen(const char *host, const char *port)
{
    int rc = 0;
    int sockfd = -1; // default fail value
    struct addrinfo *info = NULL;
    struct addrinfo *next_p = NULL;
    struct addrinfo addr = {
        .ai_family = AF_UNSPEC,
        .ai_socktype = SOCK_STREAM,
        .ai_flags = AI_PASSIVE
    };

    check(host != NULL, "Invalid host.");
    check(port != NULL, "Invalid port.");

    // get the address info for host and port
    rc = getaddrinfo(NULL, port, &addr, &info);
    check(rc == 0, "Failed to get address info for connect.");

    // cycle through the available list to find one
    for(next_p = info; next_p != NULL; next_p = next_p->ai_next)
    {
        // attempt to listen to each one
        sockfd = attempt_listen(next_p);
        if(sockfd != -1) break;
    }

    // either we found one and were able to listen or nothing.
    check(sockfd != -1, "All possible addresses failed.");

error: //fallthrough
    if(info) freeaddrinfo(info);
    // this gets set by the above to either -1 or valid
    return sockfd;
}

```

.\ex48b\statserve\src\net.h

```c
#ifndef _net_h
#define _net_h

#include <lcthw/ringbuffer.h>

#define BACKLOG 10

int nonblock(int fd);
int client_connect(char *host, char *port);
int read_some(RingBuffer * buffer, int fd, int is_socket);
int write_some(RingBuffer * buffer, int fd, int is_socket);
int server_listen(const char *host, const char *port);

#endif

```

.\ex48b\statserve\src\statserve.c

```c
#include <stdio.h>
#include <ctype.h>
#include <lcthw/dbg.h>
#include <unistd.h>
#include <stdlib.h>
#include <signal.h>
#include <sys/wait.h>
#include "net.h"
#include <netdb.h>

const int RB_SIZE = 1024 * 10;

void handle_sigchild(int sig) {
    sig = 0; // ignore it
    while(waitpid(-1, NULL, WNOHANG) > 0) {
    }
}

void client_handler(int client_fd)
{
    int rc = 0;
    // need a ringbuffer for the input
    RingBuffer *sock_rb = RingBuffer_create(RB_SIZE);

    // read_some in a loop
    while(read_some(sock_rb, client_fd, 1) != -1) {
        // write_it back off the ringbuffer
        if(write_some(sock_rb, client_fd, 1) == -1) {
            debug("Client closed.");
            break;
        }
    }

    // close the socket
    rc = close(client_fd);
    check(rc != -1, "Failed to close the socket.");

error: // fallthrough
    if(sock_rb) RingBuffer_destroy(sock_rb);
    exit(0); // just exit the child process
}

int echo_server(const char *host, const char *port)
{
    int rc = 0;
    struct sockaddr_in client_addr;
    socklen_t sin_size = sizeof(client_addr);
    int server_socket = 0;
    int client_fd = 0;

    struct sigaction sa = {
        .sa_handler = handle_sigchild,
        .sa_flags = SA_RESTART | SA_NOCLDSTOP
    };

    check(host != NULL, "Invalid host.");
    check(port != NULL, "Invalid port.");

    // create a sigaction that handles SIGCHLD
    sigemptyset(&sa.sa_mask);
    rc = sigaction(SIGCHLD, &sa, 0);
    check(rc != -1, "Failed to setup signal handler for child processes.");

    // listen on the given port and host
    server_socket = server_listen(host, port);
    check(server_socket >= 0, "bind to %s:%s failed.", host, port);

    while(1) {
        // accept the connection
        client_fd = accept(server_socket, (struct sockaddr *)&client_addr, &sin_size); 
        check(client_fd >= 0, "Failed to accept connection.");

        debug("Client connected.");

        rc = fork();
        if(rc == 0) {
            // child process
            close(server_socket); // don't need this
            // handle the client
            client_handler(client_fd);
        } else {
            // server process
            close(client_fd); // don't need this
        }
    }

error:  // fallthrough
    return -1;
}

```

.\ex48b\statserve\src\statserve.h

```c
#ifndef _statserve_h
#define _statserve_h

int echo_server(const char *host, const char *port);

#endif

```

.\ex48b\statserve\tests\statserve_tests.c

```c
#include "minunit.h"
#include <dlfcn.h>
#include "statserve.h"

char *test_dummy()
{
    return NULL;
}

char *all_tests()
{
    mu_suite_start();

    mu_run_test(test_dummy);

    return NULL;
}

RUN_TESTS(all_tests);

```

.\ex48b\c-skeleton

.\ex48b\c-skeleton\src\dbg.h

```c
#ifndef __dbg_h__
#define __dbg_h__

#include <stdio.h>
#include <errno.h>
#include <string.h>

#ifdef NDEBUG
#define debug(M, ...)
#else
#define debug(M, ...) fprintf(stderr, "DEBUG %s:%d: " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)
#endif

#define clean_errno() (errno == 0 ? "None" : strerror(errno))

#define log_err(M, ...) fprintf(stderr,\
        "[ERROR] (%s:%d: errno: %s) " M "\n", __FILE__, __LINE__,\
        clean_errno(), ##__VA_ARGS__)

#define log_warn(M, ...) fprintf(stderr,\
        "[WARN] (%s:%d: errno: %s) " M "\n",\
        __FILE__, __LINE__, clean_errno(), ##__VA_ARGS__)

#define log_info(M, ...) fprintf(stderr, "[INFO] (%s:%d) " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)

#define check(A, M, ...) if(!(A)) {\
    log_err(M, ##__VA_ARGS__); errno=0; goto error; }

#define sentinel(M, ...)  { log_err(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#define check_mem(A) check((A), "Out of memory.")

#define check_debug(A, M, ...) if(!(A)) { debug(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#endif

```

.\ex48b\c-skeleton\src\libex29.c

```c
#include <stdio.h>
#include <ctype.h>
#include "dbg.h"

int print_a_message(const char *msg)
{
    printf("A STRING: %s\n", msg);

    return 0;
}

int uppercase(const char *msg)
{
    int i = 0;

    // BUG: \0 termination problems
    for(i = 0; msg[i] != '\0'; i++) {
        printf("%c", toupper(msg[i]));
    }

    printf("\n");

    return 0;
}

int lowercase(const char *msg)
{
    int i = 0;

    // BUG: \0 termination problems
    for(i = 0; msg[i] != '\0'; i++) {
        printf("%c", tolower(msg[i]));
    }

    printf("\n");

    return 0;
}

int fail_on_purpose(const char *msg)
{
    return 1;
} 

```

.\ex48b\c-skeleton\tests\libex29_tests.c

```c
#include "minunit.h"
#include <dlfcn.h>

typedef int (*lib_function) (const char *data);
char *lib_file = "build/libYOUR_LIBRARY.so";
void *lib = NULL;

int check_function(const char *func_to_run, const char *data,
        int expected)
{
    lib_function func = dlsym(lib, func_to_run);
    check(func != NULL,
            "Did not find %s function in the library %s: %s", func_to_run,
            lib_file, dlerror());

    int rc = func(data);
    check(rc == expected, "Function %s return %d for data: %s",
            func_to_run, rc, data);

    return 1;
error:
    return 0;
}

char *test_dlopen()
{
    lib = dlopen(lib_file, RTLD_NOW);
    mu_assert(lib != NULL, "Failed to open the library to test.");

    return NULL;
}

char *test_functions()
{
    mu_assert(check_function("print_a_message", "Hello", 0),
            "print_a_message failed.");
    mu_assert(check_function("uppercase", "Hello", 0),
            "uppercase failed.");
    mu_assert(check_function("lowercase", "Hello", 0),
            "lowercase failed.");

    return NULL;
}

char *test_failures()
{
    mu_assert(check_function("fail_on_purpose", "Hello", 1),
            "fail_on_purpose should fail.");

    return NULL;
}

char *test_dlclose()
{
    int rc = dlclose(lib);
    mu_assert(rc == 0, "Failed to close lib.");

    return NULL;
}

char *all_tests()
{
    mu_suite_start();

    mu_run_test(test_dlopen);
    mu_run_test(test_functions);
    mu_run_test(test_failures);
    mu_run_test(test_dlclose);

    return NULL;
}

RUN_TESTS(all_tests);

```

.\ex48b\statserve

.\ex48b\statserve\bin\statserve.c

```c
#include <stdio.h>
#include <lcthw/dbg.h>
#include "statserve.h"
#include "net.h"

int main(int argc, char *argv[])
{
    check(argc == 3, "USAGE: statserve host port");

    const char *host = argv[1];
    const char *port = argv[2];

    check(echo_server(host, port), "Failed to run the echo server.");

    return 0;

error:

    return 1;
}

```

.\ex48b\statserve\src\dbg.h

```c
#ifndef __dbg_h__
#define __dbg_h__

#include <stdio.h>
#include <errno.h>
#include <string.h>

#ifdef NDEBUG
#define debug(M, ...)
#else
#define debug(M, ...) fprintf(stderr, "DEBUG %s:%d: " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)
#endif

#define clean_errno() (errno == 0 ? "None" : strerror(errno))

#define log_err(M, ...) fprintf(stderr,\
        "[ERROR] (%s:%d: errno: %s) " M "\n", __FILE__, __LINE__,\
        clean_errno(), ##__VA_ARGS__)

#define log_warn(M, ...) fprintf(stderr,\
        "[WARN] (%s:%d: errno: %s) " M "\n",\
        __FILE__, __LINE__, clean_errno(), ##__VA_ARGS__)

#define log_info(M, ...) fprintf(stderr, "[INFO] (%s:%d) " M "\n",\
        __FILE__, __LINE__, ##__VA_ARGS__)

#define check(A, M, ...) if(!(A)) {\
    log_err(M, ##__VA_ARGS__); errno=0; goto error; }

#define sentinel(M, ...)  { log_err(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#define check_mem(A) check((A), "Out of memory.")

#define check_debug(A, M, ...) if(!(A)) { debug(M, ##__VA_ARGS__);\
    errno=0; goto error; }

#endif

```

.\ex48b\statserve\src\net.c

```c
#include <stdlib.h>
#include <sys/select.h>
#include <stdio.h>
#include <lcthw/ringbuffer.h>
#include <lcthw/dbg.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <sys/uio.h>
#include <arpa/inet.h>
#include <netdb.h>
#include <unistd.h>
#include <fcntl.h>
#include "net.h"

struct tagbstring NL = bsStatic("\n");
struct tagbstring CRLF = bsStatic("\r\n");

int nonblock(int fd)
{
    int flags = fcntl(fd, F_GETFL, 0);
    check(flags >= 0, "Invalid flags on nonblock.");

    int rc = fcntl(fd, F_SETFL, flags | O_NONBLOCK);
    check(rc == 0, "Can't set nonblocking.");

    return 0;
error:
    return -1;
}

int client_connect(char *host, char *port)
{
    int rc = 0;
    struct addrinfo *addr = NULL;

    rc = getaddrinfo(host, port, NULL, &addr);
    check(rc == 0, "Failed to lookup %s:%s", host, port);

    int sock = socket(AF_INET, SOCK_STREAM, 0);
    check(sock >= 0, "Cannot create a socket.");

    rc = connect(sock, addr->ai_addr, addr->ai_addrlen);
    check(rc == 0, "Connect failed.");

    rc = nonblock(sock);
    check(rc == 0, "Can't set nonblocking.");

    freeaddrinfo(addr);
    return sock;

error:
    freeaddrinfo(addr);
    return -1;
}

int read_some(RingBuffer * buffer, int fd, int is_socket)
{
    int rc = 0;

    if (RingBuffer_available_data(buffer) == 0) {
        buffer->start = buffer->end = 0;
    }

    if (is_socket) {
        rc = recv(fd, RingBuffer_starts_at(buffer),
                RingBuffer_available_space(buffer), 0);
    } else {
        rc = read(fd, RingBuffer_starts_at(buffer),
                RingBuffer_available_space(buffer));
    }

    check(rc >= 0, "Failed to read from fd: %d", fd);

    RingBuffer_commit_write(buffer, rc);

    return rc;

error:
    return -1;
}

int write_some(RingBuffer * buffer, int fd, int is_socket)
{
    int rc = 0;
    bstring data = RingBuffer_get_all(buffer);

    check(data != NULL, "Failed to get from the buffer.");
    check(bfindreplace(data, &NL, &CRLF, 0) == BSTR_OK,
            "Failed to replace NL.");

    if (is_socket) {
        rc = send(fd, bdata(data), blength(data), 0);
    } else {
        rc = write(fd, bdata(data), blength(data));
    }

    check(rc == blength(data), "Failed to write everything to fd: %d.",
            fd);
    bdestroy(data);

    return rc;

error:
    return -1;
}

int attempt_listen(struct addrinfo *info)
{
    int sockfd = -1; // default fail
    int rc = -1;
    int yes = 1;

    check(info != NULL, "Invalid addrinfo.");

    // create a socket with the addrinfo
    sockfd = socket(info->ai_family, info->ai_socktype,
            info->ai_protocol);
    check_debug(sockfd != -1, "Failed to bind to address. Trying more.");

    // set the SO_REUSEADDR option on the socket
    rc = setsockopt(sockfd, SOL_SOCKET, SO_REUSEADDR, &yes, sizeof(int));
    check_debug(rc == 0, "Failed to set SO_REUSADDR.");

    // attempt to bind to it
    rc = bind(sockfd, info->ai_addr, info->ai_addrlen);
    check_debug(rc == 0, "Failed to find socket.");

    // finally listen with a backlog
    rc = listen(sockfd, BACKLOG);
    check_debug(rc == 0, "Failed to listen to socket.");

    return sockfd;

error:
    return -1;
}

int server_listen(const char *host, const char *port)
{
    int rc = 0;
    int sockfd = -1; // default fail value
    struct addrinfo *info = NULL;
    struct addrinfo *next_p = NULL;
    struct addrinfo addr = {
        .ai_family = AF_UNSPEC,
        .ai_socktype = SOCK_STREAM,
        .ai_flags = AI_PASSIVE
    };

    check(host != NULL, "Invalid host.");
    check(port != NULL, "Invalid port.");

    // get the address info for host and port
    rc = getaddrinfo(NULL, port, &addr, &info);
    check(rc == 0, "Failed to get address info for connect.");

    // cycle through the available list to find one
    for(next_p = info; next_p != NULL; next_p = next_p->ai_next)
    {
        // attempt to listen to each one
        sockfd = attempt_listen(next_p);
        if(sockfd != -1) break;
    }

    // either we found one and were able to listen or nothing.
    check(sockfd != -1, "All possible addresses failed.");

error: //fallthrough
    if(info) freeaddrinfo(info);
    // this gets set by the above to either -1 or valid
    return sockfd;
}

```

.\ex48b\statserve\src\net.h

```c
#ifndef _net_h
#define _net_h

#include <lcthw/ringbuffer.h>

#define BACKLOG 10

int nonblock(int fd);
int client_connect(char *host, char *port);
int read_some(RingBuffer * buffer, int fd, int is_socket);
int write_some(RingBuffer * buffer, int fd, int is_socket);
int server_listen(const char *host, const char *port);

#endif

```

.\ex48b\statserve\src\statserve.c

```c
#include <stdio.h>
#include <ctype.h>
#include <lcthw/dbg.h>
#include <unistd.h>
#include <stdlib.h>
#include <signal.h>
#include <sys/wait.h>
#include "net.h"
#include <netdb.h>

const int RB_SIZE = 1024 * 10;

void handle_sigchild(int sig) {
    sig = 0; // ignore it
    while(waitpid(-1, NULL, WNOHANG) > 0) {
    }
}

void client_handler(int client_fd)
{
    int rc = 0;
    // need a ringbuffer for the input
    RingBuffer *sock_rb = RingBuffer_create(RB_SIZE);

    // read_some in a loop
    while(read_some(sock_rb, client_fd, 1) != -1) {
        // write_it back off the ringbuffer
        if(write_some(sock_rb, client_fd, 1) == -1) {
            debug("Client closed.");
            break;
        }
    }

    // close the socket
    rc = close(client_fd);
    check(rc != -1, "Failed to close the socket.");

error: // fallthrough
    if(sock_rb) RingBuffer_destroy(sock_rb);
    exit(0); // just exit the child process
}

int echo_server(const char *host, const char *port)
{
    int rc = 0;
    struct sockaddr_in client_addr;
    socklen_t sin_size = sizeof(client_addr);
    int server_socket = 0;
    int client_fd = 0;

    struct sigaction sa = {
        .sa_handler = handle_sigchild,
        .sa_flags = SA_RESTART | SA_NOCLDSTOP
    };

    check(host != NULL, "Invalid host.");
    check(port != NULL, "Invalid port.");

    // create a sigaction that handles SIGCHLD
    sigemptyset(&sa.sa_mask);
    rc = sigaction(SIGCHLD, &sa, 0);
    check(rc != -1, "Failed to setup signal handler for child processes.");

    // listen on the given port and host
    server_socket = server_listen(host, port);
    check(server_socket >= 0, "bind to %s:%s failed.", host, port);

    while(1) {
        // accept the connection
        client_fd = accept(server_socket, (struct sockaddr *)&client_addr, &sin_size); 
        check(client_fd >= 0, "Failed to accept connection.");

        debug("Client connected.");

        rc = fork();
        if(rc == 0) {
            // child process
            close(server_socket); // don't need this
            // handle the client
            client_handler(client_fd);
        } else {
            // server process
            close(client_fd); // don't need this
        }
    }

error:  // fallthrough
    return -1;
}

```

.\ex48b\statserve\src\statserve.h

```c
#ifndef _statserve_h
#define _statserve_h

int echo_server(const char *host, const char *port);

#endif

```

.\ex48b\statserve\tests\statserve_tests.c

```c
#include "minunit.h"
#include <dlfcn.h>
#include "statserve.h"

char *test_dummy()
{
    return NULL;
}

char *all_tests()
{
    mu_suite_start();

    mu_run_test(test_dummy);

    return NULL;
}

RUN_TESTS(all_tests);

```

Solution

The Plan

Show you how I solved the *statserve* project.

The Purpose

Watch me solve the first project quickly, then review the code.

The Setup

First I need to install liblcthw since I'll be using that.

Then I make the project skeleton and get something, anything going.

The Server

Then I just get it accepting a connection.

The Echo

Then I decided to just make it echo back what I type.

The Final Code

