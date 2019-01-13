+++
tags = ["c"]
title = "C Lecture - 4"
description = "Exercise 48 ~ 51"
+++

Author: Zed A. Shaw

All content comes from Zed's [Lecture Repository](https://github.com/zedshaw/learn-c-the-hard-way-lectures.git) and [Libraries Repository](https://github.com/zedshaw/liblcthw).  All credit goes to Zed.



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

### Exercise 49a A Statistics Server

Project Description

The Plan

Make the *statsserver* do something using a simple protocol.

The Purpose

Learn the first steps in creating a server that answers a protocol.

The Requirements

Create this protocol:

    create Create a new statistic.
    mean   Get the current mean of a statistic.
    sample Add a new sample to a statistics.
    dump   Get all of the elements of a statistic (sum, sumsq, n, min, and max).

The Requirements

1. You'll need to allow people to name these statistics, which means using one of the map style data structures to map names to ``Stats`` structs.
2. You'll need to add the ``CRUD`` standard operations for each name.  CRUD stands for create read update delete.  Currently, the list of commands above has create, mean, and dump for reading; and sample for updating.  You need a delete command now.
3. Make the protocol *strict*! Abort any client that makes any mistakes in protocols.

Strict Protocol

Once again, in case you missed it, be ruthless!

Abort all deviant clients.

Pause!

I'm going to give you clues to solve this, so if you want to try on your own pause now!

The Clues

* Create the data structures first for holding the information for each of these commands.
* Then write a protocol parser to handle it and fill in the data.
* Then pass that data to a function that knows how to do that command.
* You can just store the stats in a Hashmap, BSTree, or TSTree for now.
* KEEP IT SIMPLE!

Important References

* You'll want to refer to the bstring documentation as much as possible to know what functions to use.

Encouragement

* Remember that this is *supposed* to be hard.
* You are *supposed* to struggle with this.
* This could take you a while, but keep up the struggle, do it bit by bit, and test little pieces as you go.
* Automate your tests!

### Exercise 49b A Statistics Server:
.\ex49b\statserve

.\ex49b\statserve\bin\statserve.c

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

.\ex49b\statserve\src\dbg.h

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

.\ex49b\statserve\src\net.c

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

bstring read_line(RingBuffer *input, const char line_ending)
{
    int i = 0;
    bstring result = NULL;

    // not super efficient
    // read a character at a time from the ring buffer
    for(i = 0; i < RingBuffer_available_data(input); i++) {
        // if the buffer has line ending
        if(input->buffer[i] == line_ending) {
            // get that much fromt he ring buffer
            result = RingBuffer_gets(input, i);
            check(result, "Failed to get line from RingBuffer");
            // make sure that we got the right amount
            check(RingBuffer_available_data(input) >= 1, 
                    "Not enough data in the RingBuffer after reading line.");
            // and commit it
            RingBuffer_commit_read(input, 1);
            break;
        }
    }

    // notice this will fail in the cases where we get a set of data
    // on the wire that does not have a line ending yet
    return result;
error:
    return NULL;
}

```

.\ex49b\statserve\src\net.h

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
bstring read_line(RingBuffer *input, const char line_ending);

#endif

```

.\ex49b\statserve\src\statserve.c

```c
#include <stdio.h>
#include <ctype.h>
#include <lcthw/dbg.h>
#include <lcthw/hashmap.h>
#include <lcthw/stats.h>
#include <unistd.h>
#include <stdlib.h>
#include <signal.h>
#include <sys/wait.h>
#include "net.h"
#include <netdb.h>

struct tagbstring LINE_SPLIT = bsStatic(" ");
struct tagbstring CREATE = bsStatic("create");
struct tagbstring STDDEV = bsStatic("stddev");
struct tagbstring MEAN = bsStatic("mean");
struct tagbstring SAMPLE = bsStatic("sample");
struct tagbstring DUMP = bsStatic("dump");
struct tagbstring DELETE = bsStatic("delete");
struct tagbstring OK = bsStatic("OK\n");
struct tagbstring ERR = bsStatic("ERR\n");
struct tagbstring DNE = bsStatic("DNE\n");
struct tagbstring EXISTS = bsStatic("EXISTS\n");
const char LINE_ENDING = '\n';

const int RB_SIZE = 1024 * 10;

Hashmap *DATA = NULL;

struct Command;

typedef int (*handler_cb)(struct Command *cmd, RingBuffer *send_rb);

typedef struct Command {
    bstring command;
    bstring name;
    bstring number;
    handler_cb handler;
} Command;

typedef struct Record {
    bstring name;
    Stats *stat;
} Record;

void handle_sigchild(int sig) {
    sig = 0; // ignore it
    while(waitpid(-1, NULL, WNOHANG) > 0) {
    }
}

void send_reply(RingBuffer *send_rb, bstring reply)
{
    RingBuffer_puts(send_rb, reply);
}

int handle_create(Command *cmd, RingBuffer *send_rb)
{
    int rc = 0;

    // if the name is in the DATA map then return exists
    if(Hashmap_get(DATA, cmd->name)) {
        send_reply(send_rb, &EXISTS);
    } else {
        // allocate a recrod
        debug("create: %s %s", bdata(cmd->name), bdata(cmd->number));

        Record *info = calloc(sizeof(Record), 1);
        check_mem(info);

        // set its stat element
        info->stat = Stats_create();
        check_mem(info->stat);

        // set its name element
        info->name = bstrcpy(cmd->name);
        check_mem(info->name);

        // do a first sample
        Stats_sample(info->stat, atof(bdata(cmd->number)));

        // add it to the hashmap
        rc = Hashmap_set(DATA, info->name, info);
        check(rc == 0, "Failed to add data to map.");

        // send an OK
        send_reply(send_rb, &OK);
    }

    return 0;
error:
    return -1;
}

int handle_sample(Command *cmd, RingBuffer *send_rb)
{
    // get the info from the hashmap
    Record *info = Hashmap_get(DATA, cmd->name);

    if(info == NULL) {
        // if it doesn't exist then DNE
        send_reply(send_rb, &DNE);
    } else {
        // else run sample on it, return the mean
        Stats_sample(info->stat, atof(bdata(cmd->number)));
        bstring reply = bformat("%f\n", Stats_mean(info->stat));
        send_reply(send_rb, reply);
        bdestroy(reply);
    }

    return 0;
}

int handle_delete(Command *cmd, RingBuffer *send_rb)
{
    log_info("delete: %s", bdata(cmd->name));
    Record *info = Hashmap_get(DATA, cmd->name);

    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        Hashmap_delete(DATA, cmd->name);

        free(info->stat);
        bdestroy(info->name);
        free(info);

        send_reply(send_rb, &OK);
    }

    return 0;
}

int handle_mean(Command *cmd, RingBuffer *send_rb)
{
    log_info("mean: %s", bdata(cmd->name));
    Record *info = Hashmap_get(DATA, cmd->name);

    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        bstring reply = bformat("%f\n", Stats_mean(info->stat));
        send_reply(send_rb, reply);
        bdestroy(reply);
    }

    return 0;
}

int handle_stddev(Command *cmd, RingBuffer *send_rb)
{
    log_info("stddev: %s", bdata(cmd->name));
    Record *info = Hashmap_get(DATA, cmd->name);

    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        bstring reply = bformat("%f\n", Stats_stddev(info->stat));
        send_reply(send_rb, reply);
        bdestroy(reply);
    }

    return 0;
}

int handle_dump(Command *cmd, RingBuffer *send_rb)
{
    log_info("dump: %s", bdata(cmd->name));
    Record *info = Hashmap_get(DATA, cmd->name);

    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        bstring reply = bformat("%f %f %f %f %ld %f %f\n",
                Stats_mean(info->stat),
                Stats_stddev(info->stat),
                info->stat->sum,
                info->stat->sumsq,
                info->stat->n,
                info->stat->min,
                info->stat->max);

        send_reply(send_rb, reply);
        bdestroy(reply);
    }

    return 0;
}

int parse_command(struct bstrList *splits, Command *cmd)
{
    // get the command
    cmd->command = splits->entry[0];

    if(biseq(cmd->command, &CREATE)) {
        check(splits->qty == 3, "Failed to parse create: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->number = splits->entry[2];
        cmd->handler = handle_create;
    } else if(biseq(cmd->command, &MEAN)) {
        check(splits->qty == 2, "Failed to parse mean: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_mean;
    } else if(biseq(cmd->command, &SAMPLE)) {
        check(splits->qty == 3, "Failed to parse sample: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->number = splits->entry[2];
        cmd->handler = handle_sample;
    } else if(biseq(cmd->command, &DUMP)) {
        check(splits->qty == 2, "Failed to parse dump: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_dump;
    } else if(biseq(cmd->command, &DELETE)) {
        check(splits->qty == 2, "Failed to parse delete: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_delete;
    } else if(biseq(cmd->command, &STDDEV)) {
        check(splits->qty == 2, "Failed to parse stddev: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_stddev;
    } else {
        sentinel("Failed to parse the command.");
    }

    return 0;
error:
    return -1;
}

int parse_line(bstring data, RingBuffer *send_rb)
{
    int rc = -1;
    Command cmd = {.command = NULL};

    // split data on line boundaries
    struct bstrList *splits = bsplits(data, &LINE_SPLIT);
    check(splits != NULL, "Bad data.");

    // parse it into a command
    rc = parse_command(splits, &cmd);
    check(rc == 0, "Failed to parse command.");

    // call the command handler for that command
    rc = cmd.handler(&cmd, send_rb);

error: // fallthrough
    if(splits) bstrListDestroy(splits);
    return rc;

}

void client_handler(int client_fd)
{
    int rc = 0;
    RingBuffer *recv_rb = RingBuffer_create(RB_SIZE);
    RingBuffer *send_rb = RingBuffer_create(RB_SIZE);

    check_mem(recv_rb);
    check_mem(send_rb);

    // keep reading into the recv buffer and sending on send
    while(read_some(recv_rb, client_fd, 1) != -1) {
        // read a line from the recv_rb
        bstring data = read_line(recv_rb, LINE_ENDING);
        check(data != NULL, "Client closed.");

        // parse it, close on any protocol errors
        rc = parse_line(data, send_rb);
        bdestroy(data); // cleanup here
        check(rc == 0, "Failed to parse user. Closing.");

        // and as long as there's something to send, send it
        if(RingBuffer_available_data(send_rb)) {
            write_some(send_rb, client_fd, 1);
        }
    }

    // close the socket
    rc = close(client_fd);
    check(rc != -1, "Failed to close the socket.");

error: // fallthrough
    if(recv_rb) RingBuffer_destroy(recv_rb);
    if(send_rb) RingBuffer_destroy(send_rb);
    exit(0); // just exit the child process
}

int setup_data_store()
{
    // a more advanced design simply wouldn't use this
    DATA = Hashmap_create(NULL, NULL);
    check_mem(DATA);

    return 0;
error:
    return -1;
}

int echo_server(const char *host, const char *port)
{
    int rc = 0;
    struct sockaddr_in client_addr;
    socklen_t sin_size = sizeof(client_addr);
    int server_socket = 0;
    int client_fd = 0;

    rc = setup_data_store();
    check(rc == 0, "Failed to setup the data store.");

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

.\ex49b\statserve\src\statserve.h

```c
#ifndef _statserve_h
#define _statserve_h

#include <lcthw/bstrlib.h>
#include <lcthw/ringbuffer.h>

struct tagbstring OK;

int setup_data_store();

int parse_line(bstring data, RingBuffer *send_rb);

int echo_server(const char *host, const char *port);

#endif

```

.\ex49b\statserve\tests\statserve_tests.c

```c
#include "minunit.h"
#include <dlfcn.h>
#include "statserve.h"
#include <lcthw/bstrlib.h>
#include <lcthw/ringbuffer.h>
#include <assert.h>

typedef struct LineTest {
    char *line;
    bstring result;
    char *description;
} LineTest;

int attempt_line(LineTest test)
{
    int rc = -1;
    bstring result = NULL;

    bstring line = bfromcstr(test.line);
    RingBuffer *send_rb = RingBuffer_create(1024);

    rc = parse_line(line, send_rb);
    check(rc == 0, "Failed to parse line.");

    result = RingBuffer_get_all(send_rb);
    check(result != NULL, "Ring buffer empty.");
    check(biseq(result, test.result), "Got the wrong output: %s expected %s",
            bdata(result), bdata(test.result));

    bdestroy(line);
    RingBuffer_destroy(send_rb);
    return 1; // using 1 for tests
error:

    log_err("Failed to process test %s: got %s", test.line, bdata(result));
    if(line) bdestroy(line);
    if(send_rb) RingBuffer_destroy(send_rb);
    return 0;
}

int run_test_lines(LineTest *tests, int count)
{
    int i = 0;

    for(i = 0; i < count; i++) {
        check(attempt_line(tests[i]), "Failed to run %s", tests[i].description);
    }

    return 1;
error:
    return 0;
}

char *test_create()
{
    LineTest tests[] = {
        {.line = "create /zed 100", .result = &OK, .description = "create zed failed"},
        {.line = "create /joe 100", .result = &OK, .description = "create joe failed"},

    };

    mu_assert(run_test_lines(tests, 2), "Failed to run create tests.");

    return NULL;
}

char *test_sample()
{
    struct tagbstring sample1 = bsStatic("100.000000\n");

    LineTest tests[] = {
        {.line = "sample /zed 100", .result = &sample1, .description = "sample zed failed."}
    };

    mu_assert(run_test_lines(tests, 1), "Failed to run sample tests.");

    return NULL;
}

char *all_tests()
{
    mu_suite_start();

    int rc = setup_data_store();
    mu_assert(rc == 0, "Failed to setup the data store.");

    mu_run_test(test_create);
    mu_run_test(test_sample);

    return NULL;
}

RUN_TESTS(all_tests);

```

.\ex49b\statserve

.\ex49b\statserve\bin\statserve.c

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

.\ex49b\statserve\src\dbg.h

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

.\ex49b\statserve\src\net.c

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

bstring read_line(RingBuffer *input, const char line_ending)
{
    int i = 0;
    bstring result = NULL;

    // not super efficient
    // read a character at a time from the ring buffer
    for(i = 0; i < RingBuffer_available_data(input); i++) {
        // if the buffer has line ending
        if(input->buffer[i] == line_ending) {
            // get that much fromt he ring buffer
            result = RingBuffer_gets(input, i);
            check(result, "Failed to get line from RingBuffer");
            // make sure that we got the right amount
            check(RingBuffer_available_data(input) >= 1, 
                    "Not enough data in the RingBuffer after reading line.");
            // and commit it
            RingBuffer_commit_read(input, 1);
            break;
        }
    }

    // notice this will fail in the cases where we get a set of data
    // on the wire that does not have a line ending yet
    return result;
error:
    return NULL;
}

```

.\ex49b\statserve\src\net.h

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
bstring read_line(RingBuffer *input, const char line_ending);

#endif

```

.\ex49b\statserve\src\statserve.c

```c
#include <stdio.h>
#include <ctype.h>
#include <lcthw/dbg.h>
#include <lcthw/hashmap.h>
#include <lcthw/stats.h>
#include <unistd.h>
#include <stdlib.h>
#include <signal.h>
#include <sys/wait.h>
#include "net.h"
#include <netdb.h>

struct tagbstring LINE_SPLIT = bsStatic(" ");
struct tagbstring CREATE = bsStatic("create");
struct tagbstring STDDEV = bsStatic("stddev");
struct tagbstring MEAN = bsStatic("mean");
struct tagbstring SAMPLE = bsStatic("sample");
struct tagbstring DUMP = bsStatic("dump");
struct tagbstring DELETE = bsStatic("delete");
struct tagbstring OK = bsStatic("OK\n");
struct tagbstring ERR = bsStatic("ERR\n");
struct tagbstring DNE = bsStatic("DNE\n");
struct tagbstring EXISTS = bsStatic("EXISTS\n");
const char LINE_ENDING = '\n';

const int RB_SIZE = 1024 * 10;

Hashmap *DATA = NULL;

struct Command;

typedef int (*handler_cb)(struct Command *cmd, RingBuffer *send_rb);

typedef struct Command {
    bstring command;
    bstring name;
    bstring number;
    handler_cb handler;
} Command;

typedef struct Record {
    bstring name;
    Stats *stat;
} Record;

void handle_sigchild(int sig) {
    sig = 0; // ignore it
    while(waitpid(-1, NULL, WNOHANG) > 0) {
    }
}

void send_reply(RingBuffer *send_rb, bstring reply)
{
    RingBuffer_puts(send_rb, reply);
}

int handle_create(Command *cmd, RingBuffer *send_rb)
{
    int rc = 0;

    // if the name is in the DATA map then return exists
    if(Hashmap_get(DATA, cmd->name)) {
        send_reply(send_rb, &EXISTS);
    } else {
        // allocate a recrod
        debug("create: %s %s", bdata(cmd->name), bdata(cmd->number));

        Record *info = calloc(sizeof(Record), 1);
        check_mem(info);

        // set its stat element
        info->stat = Stats_create();
        check_mem(info->stat);

        // set its name element
        info->name = bstrcpy(cmd->name);
        check_mem(info->name);

        // do a first sample
        Stats_sample(info->stat, atof(bdata(cmd->number)));

        // add it to the hashmap
        rc = Hashmap_set(DATA, info->name, info);
        check(rc == 0, "Failed to add data to map.");

        // send an OK
        send_reply(send_rb, &OK);
    }

    return 0;
error:
    return -1;
}

int handle_sample(Command *cmd, RingBuffer *send_rb)
{
    // get the info from the hashmap
    Record *info = Hashmap_get(DATA, cmd->name);

    if(info == NULL) {
        // if it doesn't exist then DNE
        send_reply(send_rb, &DNE);
    } else {
        // else run sample on it, return the mean
        Stats_sample(info->stat, atof(bdata(cmd->number)));
        bstring reply = bformat("%f\n", Stats_mean(info->stat));
        send_reply(send_rb, reply);
        bdestroy(reply);
    }

    return 0;
}

int handle_delete(Command *cmd, RingBuffer *send_rb)
{
    log_info("delete: %s", bdata(cmd->name));
    Record *info = Hashmap_get(DATA, cmd->name);

    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        Hashmap_delete(DATA, cmd->name);

        free(info->stat);
        bdestroy(info->name);
        free(info);

        send_reply(send_rb, &OK);
    }

    return 0;
}

int handle_mean(Command *cmd, RingBuffer *send_rb)
{
    log_info("mean: %s", bdata(cmd->name));
    Record *info = Hashmap_get(DATA, cmd->name);

    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        bstring reply = bformat("%f\n", Stats_mean(info->stat));
        send_reply(send_rb, reply);
        bdestroy(reply);
    }

    return 0;
}

int handle_stddev(Command *cmd, RingBuffer *send_rb)
{
    log_info("stddev: %s", bdata(cmd->name));
    Record *info = Hashmap_get(DATA, cmd->name);

    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        bstring reply = bformat("%f\n", Stats_stddev(info->stat));
        send_reply(send_rb, reply);
        bdestroy(reply);
    }

    return 0;
}

int handle_dump(Command *cmd, RingBuffer *send_rb)
{
    log_info("dump: %s", bdata(cmd->name));
    Record *info = Hashmap_get(DATA, cmd->name);

    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        bstring reply = bformat("%f %f %f %f %ld %f %f\n",
                Stats_mean(info->stat),
                Stats_stddev(info->stat),
                info->stat->sum,
                info->stat->sumsq,
                info->stat->n,
                info->stat->min,
                info->stat->max);

        send_reply(send_rb, reply);
        bdestroy(reply);
    }

    return 0;
}

int parse_command(struct bstrList *splits, Command *cmd)
{
    // get the command
    cmd->command = splits->entry[0];

    if(biseq(cmd->command, &CREATE)) {
        check(splits->qty == 3, "Failed to parse create: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->number = splits->entry[2];
        cmd->handler = handle_create;
    } else if(biseq(cmd->command, &MEAN)) {
        check(splits->qty == 2, "Failed to parse mean: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_mean;
    } else if(biseq(cmd->command, &SAMPLE)) {
        check(splits->qty == 3, "Failed to parse sample: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->number = splits->entry[2];
        cmd->handler = handle_sample;
    } else if(biseq(cmd->command, &DUMP)) {
        check(splits->qty == 2, "Failed to parse dump: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_dump;
    } else if(biseq(cmd->command, &DELETE)) {
        check(splits->qty == 2, "Failed to parse delete: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_delete;
    } else if(biseq(cmd->command, &STDDEV)) {
        check(splits->qty == 2, "Failed to parse stddev: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_stddev;
    } else {
        sentinel("Failed to parse the command.");
    }

    return 0;
error:
    return -1;
}

int parse_line(bstring data, RingBuffer *send_rb)
{
    int rc = -1;
    Command cmd = {.command = NULL};

    // split data on line boundaries
    struct bstrList *splits = bsplits(data, &LINE_SPLIT);
    check(splits != NULL, "Bad data.");

    // parse it into a command
    rc = parse_command(splits, &cmd);
    check(rc == 0, "Failed to parse command.");

    // call the command handler for that command
    rc = cmd.handler(&cmd, send_rb);

error: // fallthrough
    if(splits) bstrListDestroy(splits);
    return rc;

}

void client_handler(int client_fd)
{
    int rc = 0;
    RingBuffer *recv_rb = RingBuffer_create(RB_SIZE);
    RingBuffer *send_rb = RingBuffer_create(RB_SIZE);

    check_mem(recv_rb);
    check_mem(send_rb);

    // keep reading into the recv buffer and sending on send
    while(read_some(recv_rb, client_fd, 1) != -1) {
        // read a line from the recv_rb
        bstring data = read_line(recv_rb, LINE_ENDING);
        check(data != NULL, "Client closed.");

        // parse it, close on any protocol errors
        rc = parse_line(data, send_rb);
        bdestroy(data); // cleanup here
        check(rc == 0, "Failed to parse user. Closing.");

        // and as long as there's something to send, send it
        if(RingBuffer_available_data(send_rb)) {
            write_some(send_rb, client_fd, 1);
        }
    }

    // close the socket
    rc = close(client_fd);
    check(rc != -1, "Failed to close the socket.");

error: // fallthrough
    if(recv_rb) RingBuffer_destroy(recv_rb);
    if(send_rb) RingBuffer_destroy(send_rb);
    exit(0); // just exit the child process
}

int setup_data_store()
{
    // a more advanced design simply wouldn't use this
    DATA = Hashmap_create(NULL, NULL);
    check_mem(DATA);

    return 0;
error:
    return -1;
}

int echo_server(const char *host, const char *port)
{
    int rc = 0;
    struct sockaddr_in client_addr;
    socklen_t sin_size = sizeof(client_addr);
    int server_socket = 0;
    int client_fd = 0;

    rc = setup_data_store();
    check(rc == 0, "Failed to setup the data store.");

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

.\ex49b\statserve\src\statserve.h

```c
#ifndef _statserve_h
#define _statserve_h

#include <lcthw/bstrlib.h>
#include <lcthw/ringbuffer.h>

struct tagbstring OK;

int setup_data_store();

int parse_line(bstring data, RingBuffer *send_rb);

int echo_server(const char *host, const char *port);

#endif

```

.\ex49b\statserve\tests\statserve_tests.c

```c
#include "minunit.h"
#include <dlfcn.h>
#include "statserve.h"
#include <lcthw/bstrlib.h>
#include <lcthw/ringbuffer.h>
#include <assert.h>

typedef struct LineTest {
    char *line;
    bstring result;
    char *description;
} LineTest;

int attempt_line(LineTest test)
{
    int rc = -1;
    bstring result = NULL;

    bstring line = bfromcstr(test.line);
    RingBuffer *send_rb = RingBuffer_create(1024);

    rc = parse_line(line, send_rb);
    check(rc == 0, "Failed to parse line.");

    result = RingBuffer_get_all(send_rb);
    check(result != NULL, "Ring buffer empty.");
    check(biseq(result, test.result), "Got the wrong output: %s expected %s",
            bdata(result), bdata(test.result));

    bdestroy(line);
    RingBuffer_destroy(send_rb);
    return 1; // using 1 for tests
error:

    log_err("Failed to process test %s: got %s", test.line, bdata(result));
    if(line) bdestroy(line);
    if(send_rb) RingBuffer_destroy(send_rb);
    return 0;
}

int run_test_lines(LineTest *tests, int count)
{
    int i = 0;

    for(i = 0; i < count; i++) {
        check(attempt_line(tests[i]), "Failed to run %s", tests[i].description);
    }

    return 1;
error:
    return 0;
}

char *test_create()
{
    LineTest tests[] = {
        {.line = "create /zed 100", .result = &OK, .description = "create zed failed"},
        {.line = "create /joe 100", .result = &OK, .description = "create joe failed"},

    };

    mu_assert(run_test_lines(tests, 2), "Failed to run create tests.");

    return NULL;
}

char *test_sample()
{
    struct tagbstring sample1 = bsStatic("100.000000\n");

    LineTest tests[] = {
        {.line = "sample /zed 100", .result = &sample1, .description = "sample zed failed."}
    };

    mu_assert(run_test_lines(tests, 1), "Failed to run sample tests.");

    return NULL;
}

char *all_tests()
{
    mu_suite_start();

    int rc = setup_data_store();
    mu_assert(rc == 0, "Failed to setup the data store.");

    mu_run_test(test_create);
    mu_run_test(test_sample);

    return NULL;
}

RUN_TESTS(all_tests);

```

Solution

The Plan

I'll show you how I implemented the protocol in the smallest code possible.

I won't implement all of the CRUD operations, so you can go look at the 
git repo for this project to see a full implementation.

The Setup

First I setup the data, then the protocol parser, then the handlers.

The Protocol

    create Create a new statistic.
    mean   Get the current mean of a statistic.
    sample Add a new sample to a statistics.
    dump   Get all of the elements of a statistic (sum, sumsq, n, min, and max).Final Code

The Command Structure

    typedef struct Command {
        bstring command;
        bstring name;
        bstring number;
        handler_cb handler;
    } Command;

The Storage Record

    typedef struct Record {
        bstring name;
        Stats *stat;
    } Record;

The Design

* Accept a connection
* Parse the line into the Command
* Run a handler function to process it
* Temporarily store into a Hashmap

Final Thoughts

The last thing I would do is add better tests and round out the protocol with CRUD operations.

### Exercise 50a Routing The Statistics

Project Description

The Plan

You are now given vague instructions and have to "solve" as best you can.

The Purpose

To give you freedom to be creative, and also taste a real project with vague
specifications.

Many times all you get is a single sentence in a bug tracker. Oh well.

The Requirements

Allow people to work with statistics at arbitrary URLs in the server.
You get to define what that means, but think "web application".

Pause!

Try to solve it on your own then continue.

The Clues

Answer these questions:

1. What happens when I have a statistics "under" another, as in /age/northamerica/ is under /age/.
2. Could you do the summary statistics we talked about?  A mean of means and mean of standard deviations that are rolled up the tree?
3. What data structures do you need?  Starting with data is key here too. Data data data.
4. Are your tests good enough?  Before you start you might want to get good tests that use the protocol.

Important References

* Definitely look at the statistics code you built in liblcthw if you do the summary statistics.

Encouragement

This is hard, as I've said all along, however it is all doable. It's simply a matter of breaking the problems down and tackling each little piece.

### Exercise 50b Routing the Statistics

Solution

The Plan

Show you how I solved the problem of routing the names of statistics as URLs.

The Setup

1. First thing I did was make sure my tests were really good.
2. Then I designed the data structures I'd need.
3. Then I did the work to make them functions.
4. The protocol shouldn't need to change.

"URLs"

I'll define paths as simply names separated by /.

Real URLs are way more complex than that.

Data Structure

I just added:

    struct bstrList *path;

To the Command struct to hold paths.

URL Meaning

Kind of weird, but:

    Deepest part of URL is "parent", this is the main stat.
    Children are next segments up, and are mean-of-mean stats.

New Processing

1. Change to a loop over all paths with a "scan path" function.
2. Add optional path parameter to handlers.
3. Parse the path in *parse\_command* to set path in Commands.
4. In sample and create, change root processing vs. child processing.
5. Move *send\_reply* over to *net.c* instead.

Test First Path Parsing

I'll write a small test for just the *scan\_paths* part first.

Then wire that in and use the existing tests to confirm the old code
works.

The Code

.\ex50b\statserve

.\ex50b\statserve\bin\statserve.c

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

.\ex50b\statserve\src\dbg.h

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

.\ex50b\statserve\src\net.c

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

bstring read_line(RingBuffer *input, const char line_ending)
{
    int i = 0;
    bstring result = NULL;

    // not super efficient
    // read a character at a time from the ring buffer
    for(i = 0; i < RingBuffer_available_data(input); i++) {
        // if the buffer has line ending
        if(input->buffer[i] == line_ending) {
            // get that much fromt he ring buffer
            result = RingBuffer_gets(input, i);
            check(result, "Failed to get line from RingBuffer");
            // make sure that we got the right amount
            check(RingBuffer_available_data(input) >= 1, 
                    "Not enough data in the RingBuffer after reading line.");
            // and commit it
            RingBuffer_commit_read(input, 1);
            break;
        }
    }

    // notice this will fail in the cases where we get a set of data
    // on the wire that does not have a line ending yet
    return result;
error:
    return NULL;
}

void send_reply(RingBuffer *send_rb, bstring reply)
{
    RingBuffer_puts(send_rb, reply);
}

```

.\ex50b\statserve\src\net.h

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
bstring read_line(RingBuffer *input, const char line_ending);
void send_reply(RingBuffer *send_rb, bstring reply);

#endif

```

.\ex50b\statserve\src\statserve.c

```c
#include <stdio.h>
#include <ctype.h>
#include <lcthw/dbg.h>
#include <lcthw/hashmap.h>
#include <unistd.h>
#include <stdlib.h>
#include <signal.h>
#include <sys/wait.h>
#include "net.h"
#include <netdb.h>
#include "statserve.h"

struct tagbstring LINE_SPLIT = bsStatic(" ");
struct tagbstring CREATE = bsStatic("create");
struct tagbstring STDDEV = bsStatic("stddev");
struct tagbstring MEAN = bsStatic("mean");
struct tagbstring SAMPLE = bsStatic("sample");
struct tagbstring DUMP = bsStatic("dump");
struct tagbstring DELETE = bsStatic("delete");
struct tagbstring OK = bsStatic("OK\n");
struct tagbstring ERR = bsStatic("ERR\n");
struct tagbstring DNE = bsStatic("DNE\n");
struct tagbstring EXISTS = bsStatic("EXISTS\n");
struct tagbstring SLASH = bsStatic("/");
const char LINE_ENDING = '\n';

const int RB_SIZE = 1024 * 10;

Hashmap *DATA = NULL;

void handle_sigchild(int sig) {
    sig = 0; // ignore it
    while(waitpid(-1, NULL, WNOHANG) > 0) {
    }
}

int handle_create(Command *cmd, RingBuffer *send_rb, bstring path)
{
    int rc = 0;
    int is_root = biseq(path, cmd->name);
    log_info("create: %s %s %s", bdata(cmd->name), bdata(path), bdata(cmd->number));

    Record *info = Hashmap_get(DATA, path);

    if(info != NULL && is_root) {
        // report if root exists, just skip children
        send_reply(send_rb, &EXISTS);
    } else if(info != NULL) {
        debug("Child %s exists, skipping it.", bdata(path));
        return 0;
    } else {
        // new child so make it
        debug("create: %s %s", bdata(path), bdata(cmd->number));

        Record *info = calloc(sizeof(Record), 1);
        check_mem(info);

        // set its stat element
        info->stat = Stats_create();
        check_mem(info->stat);

        // set its name element
        info->name = bstrcpy(path);
        check_mem(info->name);

        // do a first sample
        Stats_sample(info->stat, atof(bdata(cmd->number)));

        // add it to the hashmap
        rc = Hashmap_set(DATA, info->name, info);
        check(rc == 0, "Failed to add data to map.");

        // only send the for the root part
        if(is_root) {
            send_reply(send_rb, &OK);
        }
    }

    return 0;
error:
    return -1;
}

int handle_sample(Command *cmd, RingBuffer *send_rb, bstring path)
{
    // get the info from the hashmap
    Record *info = Hashmap_get(DATA, path);
    int is_root = biseq(path, cmd->name);
    log_info("sample %s %s %s", bdata(cmd->name), bdata(path), bdata(cmd->number));
    bstring child_path = NULL;

    if(info == NULL) {
        // if it doesn't exist then DNE
        send_reply(send_rb, &DNE);
        return 0;
    } else {
        if(is_root) {
            // just sample the root like normal
            Stats_sample(info->stat, atof(bdata(cmd->number)));
        } else {
            // need to do some hackery to get the child path
            // for rolling up mean-of-means on it

            // increase the qty on path up one
            cmd->path->qty++;
            // get the "child path" (previous path?)
            child_path = bjoin(cmd->path, &SLASH);
            // get that info from the DATA
            Record *child_info = Hashmap_get(DATA, child_path);
            bdestroy(child_path);

            // if it exists then sample on it
            if(child_info) {
                // info is /logins, child_info is /logins/zed 
                // we want /logins/zed's mean to be a new sample on /logins
                Stats_sample(info->stat, Stats_mean(child_info->stat));
            }
            // drop the path back to where it was
            cmd->path->qty--;
        }

    }

    // do the reply for the mean last
    bstring reply = bformat("%f\n", Stats_mean(info->stat));
    send_reply(send_rb, reply);
    bdestroy(reply);

    return 0;
}

int handle_delete(Command *cmd, RingBuffer *send_rb, bstring path)
{
    log_info("delete: %s %s", bdata(cmd->name), bdata(path));
    Record *info = Hashmap_get(DATA, path);
    int is_root = biseq(path, cmd->name);

    // BUG: should just decide that this isn't scanned 
    // but run once, for now just only run on root
    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else if(is_root) {
        Hashmap_delete(DATA, path);

        free(info->stat);
        bdestroy(info->name);
        free(info);

        send_reply(send_rb, &OK);
    }

    return 0;
}

int handle_mean(Command *cmd, RingBuffer *send_rb, bstring path)
{
    log_info("mean: %s %s %s", bdata(cmd->name), bdata(path), bdata(path));
    Record *info = Hashmap_get(DATA, path);

    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        bstring reply = bformat("%f\n", Stats_mean(info->stat));
        send_reply(send_rb, reply);
        bdestroy(reply);
    }

    return 0;
}

int handle_stddev(Command *cmd, RingBuffer *send_rb, bstring path)
{
    log_info("stddev: %s %s %s", bdata(cmd->name), bdata(path), bdata(path));
    Record *info = Hashmap_get(DATA, path);

    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        bstring reply = bformat("%f\n", Stats_stddev(info->stat));
        send_reply(send_rb, reply);
        bdestroy(reply);
    }

    return 0;
}

int handle_dump(Command *cmd, RingBuffer *send_rb, bstring path)
{
    log_info("dump: %s, %s, %s", bdata(cmd->name), bdata(path), bdata(path));
    Record *info = Hashmap_get(DATA, path);

    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        bstring reply = bformat("%f %f %f %f %ld %f %f\n",
                Stats_mean(info->stat),
                Stats_stddev(info->stat),
                info->stat->sum,
                info->stat->sumsq,
                info->stat->n,
                info->stat->min,
                info->stat->max);

        send_reply(send_rb, reply);
        bdestroy(reply);
    }

    return 0;
}

int parse_command(struct bstrList *splits, Command *cmd)
{
    // get the command
    cmd->command = splits->entry[0];

    if(biseq(cmd->command, &CREATE)) {
        check(splits->qty == 3, "Failed to parse create: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->number = splits->entry[2];
        cmd->handler = handle_create;
    } else if(biseq(cmd->command, &MEAN)) {
        check(splits->qty == 2, "Failed to parse mean: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_mean;
    } else if(biseq(cmd->command, &SAMPLE)) {
        check(splits->qty == 3, "Failed to parse sample: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->number = splits->entry[2];
        cmd->handler = handle_sample;
    } else if(biseq(cmd->command, &DUMP)) {
        check(splits->qty == 2, "Failed to parse dump: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_dump;
    } else if(biseq(cmd->command, &DELETE)) {
        check(splits->qty == 2, "Failed to parse delete: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_delete;
    } else if(biseq(cmd->command, &STDDEV)) {
        check(splits->qty == 2, "Failed to parse stddev: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_stddev;
    } else {
        sentinel("Failed to parse the command.");
    }

    return 0;
error:
    return -1;
}

int scan_paths(Command *cmd, RingBuffer *send_rb)
{
    check(cmd->path != NULL, "Path was not set in command.");

    int rc = 0;
    // save the original path length
    size_t qty = cmd->path->qty;

    // starting at the longest path, shorten it and call
    // for each one:
    for(; cmd->path->qty > 1; cmd->path->qty--) {
        // remake the path with / again
        bstring path = bjoin(cmd->path, &SLASH);
        // call the handler with the path
        rc = cmd->handler(cmd, send_rb, path);
        // if the handler returns != 0 then abort and return that
        bdestroy(path);

        if(rc != 0) break;
    }

    // restore path length
    cmd->path->qty = qty;
    return rc;
error:
    return -1;
}

struct bstrList *parse_name(bstring name)
{
    return bsplits(name, &SLASH);
}

int parse_line(bstring data, RingBuffer *send_rb)
{
    int rc = -1;
    Command cmd = {.command = NULL};

    // split data on line boundaries
    struct bstrList *splits = bsplits(data, &LINE_SPLIT);
    check(splits != NULL, "Bad data.");

    // parse it into a command
    rc = parse_command(splits, &cmd);
    check(rc == 0, "Failed to parse command.");

    // parse the name into the path we need for scan_paths
    cmd.path = parse_name(cmd.name);
    check(cmd.path != NULL, "Invalid path.");

    // scan the path and call the handlers
    rc = scan_paths(&cmd, send_rb);
    check(rc == 0, "Failure running command against path: %s", bdata(cmd.name));

    bstrListDestroy(cmd.path);
    bstrListDestroy(splits);

    return 0;

error: // fallthrough
    if(cmd.path) bstrListDestroy(cmd.path);
    if(splits) bstrListDestroy(splits);
    return -1;
}

void client_handler(int client_fd)
{
    int rc = 0;
    RingBuffer *recv_rb = RingBuffer_create(RB_SIZE);
    RingBuffer *send_rb = RingBuffer_create(RB_SIZE);

    check_mem(recv_rb);
    check_mem(send_rb);

    // keep reading into the recv buffer and sending on send
    while(read_some(recv_rb, client_fd, 1) != -1) {
        // read a line from the recv_rb
        bstring data = read_line(recv_rb, LINE_ENDING);
        check(data != NULL, "Client closed.");

        // parse it, close on any protocol errors
        rc = parse_line(data, send_rb);
        bdestroy(data); // cleanup here
        check(rc == 0, "Failed to parse user. Closing.");

        // and as long as there's something to send, send it
        if(RingBuffer_available_data(send_rb)) {
            write_some(send_rb, client_fd, 1);
        }
    }

    // close the socket
    rc = close(client_fd);
    check(rc != -1, "Failed to close the socket.");

error: // fallthrough
    if(recv_rb) RingBuffer_destroy(recv_rb);
    if(send_rb) RingBuffer_destroy(send_rb);
    exit(0); // just exit the child process
}

int setup_data_store()
{
    // a more advanced design simply wouldn't use this
    DATA = Hashmap_create(NULL, NULL);
    check_mem(DATA);

    return 0;
error:
    return -1;
}

int echo_server(const char *host, const char *port)
{
    int rc = 0;
    struct sockaddr_in client_addr;
    socklen_t sin_size = sizeof(client_addr);
    int server_socket = 0;
    int client_fd = 0;

    rc = setup_data_store();
    check(rc == 0, "Failed to setup the data store.");

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

.\ex50b\statserve\src\statserve.h

```c
#ifndef _statserve_h
#define _statserve_h

#include <lcthw/bstrlib.h>
#include <lcthw/ringbuffer.h>
#include <lcthw/stats.h>

struct Command;

typedef int (*handler_cb)(struct Command *cmd, RingBuffer *send_rb, bstring path);

typedef struct Command {
    bstring command;
    bstring name;
    struct bstrList *path;
    bstring number;
    handler_cb handler;
} Command;

typedef struct Record {
    bstring name;
    Stats *stat;
} Record;

struct tagbstring OK;

int setup_data_store();

struct bstrList *parse_name(bstring name);

int scan_paths(Command *cmd, RingBuffer *send_rb);

int parse_line(bstring data, RingBuffer *send_rb);

int echo_server(const char *host, const char *port);

#endif

```

.\ex50b\statserve\tests\statserve_tests.c

```c
#include "minunit.h"
#include <dlfcn.h>
#include "statserve.h"
#include <lcthw/bstrlib.h>
#include <lcthw/ringbuffer.h>
#include <assert.h>

typedef struct LineTest {
    char *line;
    bstring result;
    char *description;
} LineTest;

int attempt_line(LineTest test)
{
    int rc = -1;
    bstring result = NULL;

    bstring line = bfromcstr(test.line);
    RingBuffer *send_rb = RingBuffer_create(1024);

    rc = parse_line(line, send_rb);
    check(rc == 0, "Failed to parse line.");

    result = RingBuffer_get_all(send_rb);
    check(result != NULL, "Ring buffer empty.");
    check(biseq(result, test.result), "Got the wrong output: %s expected %s",
            bdata(result), bdata(test.result));

    bdestroy(line);
    RingBuffer_destroy(send_rb);
    return 1; // using 1 for tests
error:

    log_err("Failed to process test %s: got %s", test.line, bdata(result));
    if(line) bdestroy(line);
    if(send_rb) RingBuffer_destroy(send_rb);
    return 0;
}

int run_test_lines(LineTest *tests, int count)
{
    int i = 0;

    for(i = 0; i < count; i++) {
        check(attempt_line(tests[i]), "Failed to run %s", tests[i].description);
    }

    return 1;
error:
    return 0;
}

int fake_command(Command *cmd, RingBuffer *send_rb, bstring path)
{
    check(cmd != NULL, "Bad cmd.");
    check(cmd->path != NULL, "Bad path.");
    check(send_rb != NULL, "Bad send_rb.");
    check(path != NULL, "Bad path given.");

    return 0;
error:
    return -1;
}

char *test_path_parsing()
{
    struct bstrList *result = NULL;
    struct tagbstring slash = bsStatic("/");
    struct tagbstring logins_zed = bsStatic("/logins/zed");
    struct tagbstring command_name = bsStatic("dump");
    RingBuffer *send_rb = RingBuffer_create(1024);
    struct bstrList *path = bsplits(&logins_zed, &slash);
    int rc = 0;

    Command fake = {
        .command = &command_name,
        .name = &logins_zed,
        .number = NULL,
        .handler = fake_command,
        .path = path
    };

    result = parse_name(&logins_zed);
    mu_assert(result != NULL, "Failed to parse /logins/zed");

    rc = scan_paths(&fake, send_rb); 
    mu_assert(rc != -1, "scan_paths failed.");

    return NULL;
}

char *test_create()
{
    LineTest tests[] = {
        {.line = "create /zed 100", .result = &OK, .description = "create zed failed"},
        {.line = "create /joe 100", .result = &OK, .description = "create joe failed"},

    };

    mu_assert(run_test_lines(tests, 2), "Failed to run create tests.");

    return NULL;
}

char *test_sample()
{
    struct tagbstring sample1 = bsStatic("100.000000\n");

    LineTest tests[] = {
        {.line = "sample /zed 100", .result = &sample1, .description = "sample zed failed."}
    };

    mu_assert(run_test_lines(tests, 1), "Failed to run sample tests.");

    return NULL;
}

char *all_tests()
{
    mu_suite_start();

    int rc = setup_data_store();
    mu_assert(rc == 0, "Failed to setup the data store.");

    mu_run_test(test_create);
    mu_run_test(test_sample);
    mu_run_test(test_path_parsing);

    return NULL;
}

RUN_TESTS(all_tests);

```

Final Review

### Exercise 51a Storing the Statistics

Project Description

The Plan

Learn to store the statistics to the hard disk.

There are meany issues with this.

The Purpose

To teach you about various problems related to securely storing files.

The Requirements

For this exercise, you'll add two commands for storing to and loading statistics
from a hard drive:

store
    If there's a URL, store it to a hard drive.

load
    If there are two URLs, load the statistic from the hard drive based on the first URL, and then put it into the second URL that's in memory.

The Requirements

1. If URLs have ``/`` characters in them, then that conflicts with the filesystem's use of slashes.  How will you solve this?
2. If URLs have ``/`` characters in them, then someone can use your server to overwrite files on a hard drive by giving paths to them.  How will you solve this?
3. If you choose to use deeply nested directories, then traversing directories to find files will be very slow.  What will you do here?

The Requirements

4. If you choose to use one directory and hash URLs (oops, I gave a hint), then directories with too many files in them are slow.  How will you solve this?
5. What happens when someone loads a statistic from a hard drive into a URL that already exists?
6. How will someone running ``statserve`` know where the storage should be?

The Clues

There are no clues.  You can do this.

### Exercise 51b Storing the Statistics

Solution

The Plan

Show you how I solved the problem of storing the statistics to disk.

Security Requirements

* Use *realpath* to make sure that paths are in one place.
* Use _BAD_ encryption to mangle the stored names.
* No other security beyond that. Just a demo of the path issue.

XTEA Encryption

* For an extra challenge I decided to "hash" names with XTEA.
* https://en.wikipedia.org/wiki/XTEA for the code.
* Normally I wouldn't do this, but wanted to show you XTEA.
* Because XTEA if cool and fun, although broken.
* DON'T USE XTEA FOR ENCRYPTION.

Improvements

* Let commands set cmd->path = NULL to indicate non-recursive.
* Change *echo_server* to *run_server* finally.
* Allow a 3rd storage path argument on the command line.
* Allow an additional argument to Command.

Weirdness

* Forking means I can't share data between clients without storage.
* Storing doesn't happen automatically, only explicitly.
* Loading acts as a copy command.
* XTEA isn't the best algorithm at all.  Just for fun.

How I Did It

1. Create the LOAD and STORE handlers.
2. Add Command.arg and set those in *parse\_command*.
3. Move path parsing up to allow non-recursive handling with path = NULL.
4. Write a *sanitize\_location* and *encrypt\_armor\_name* function, test them.
5. Write *handle\_store* first to allow testing *handle\_load*.
6. Use *open* (man 2 open) with O_EXLOCK to get exclusive locks on files.
7. Using *close* (man 2 close) should release EXLOCK, but not clear on this.

The Code

.\ex51b\statserve

.\ex51b\statserve\bin\statserve.c

```c
#include <stdio.h>
#include <lcthw/dbg.h>
#include "statserve.h"
#include "net.h"

int main(int argc, char *argv[])
{
    check(argc == 4, "USAGE: statserve host port store_path");

    const char *host = argv[1];
    const char *port = argv[2];
    const char *store_path = argv[3];

    check(run_server(host, port, store_path), "Failed to run the echo server.");

    return 0;

error:

    return 1;
}

```

.\ex51b\statserve\src\dbg.h

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

.\ex51b\statserve\src\net.c

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

bstring read_line(RingBuffer *input, const char line_ending)
{
    int i = 0;
    bstring result = NULL;

    // not super efficient
    // read a character at a time from the ring buffer
    for(i = 0; i < RingBuffer_available_data(input); i++) {
        // if the buffer has line ending
        if(input->buffer[i] == line_ending) {
            // get that much fromt he ring buffer
            result = RingBuffer_gets(input, i);
            check(result, "Failed to get line from RingBuffer");
            // make sure that we got the right amount
            check(RingBuffer_available_data(input) >= 1, 
                    "Not enough data in the RingBuffer after reading line.");
            // and commit it
            RingBuffer_commit_read(input, 1);
            break;
        }
    }

    // notice this will fail in the cases where we get a set of data
    // on the wire that does not have a line ending yet
    return result;
error:
    return NULL;
}

void send_reply(RingBuffer *send_rb, bstring reply)
{
    RingBuffer_puts(send_rb, reply);
}

```

.\ex51b\statserve\src\net.h

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
bstring read_line(RingBuffer *input, const char line_ending);
void send_reply(RingBuffer *send_rb, bstring reply);

#endif

```

.\ex51b\statserve\src\statserve.c

```c
#include <stdio.h>
#include <ctype.h>
#include <lcthw/dbg.h>
#include <lcthw/hashmap.h>
#include <unistd.h>
#include <stdlib.h>
#include <signal.h>
#include <sys/wait.h>
#include "net.h"
#include <netdb.h>
#include <fcntl.h>
#include "statserve.h"

struct tagbstring LINE_SPLIT = bsStatic(" ");
struct tagbstring CREATE = bsStatic("create");
struct tagbstring STDDEV = bsStatic("stddev");
struct tagbstring MEAN = bsStatic("mean");
struct tagbstring SAMPLE = bsStatic("sample");
struct tagbstring DUMP = bsStatic("dump");
struct tagbstring DELETE = bsStatic("delete");
struct tagbstring STORE = bsStatic("store");
struct tagbstring LOAD = bsStatic("load");
struct tagbstring OK = bsStatic("OK\n");
struct tagbstring ERR = bsStatic("ERR\n");
struct tagbstring DNE = bsStatic("DNE\n");
struct tagbstring EXISTS = bsStatic("EXISTS\n");
struct tagbstring SLASH = bsStatic("/");
const char LINE_ENDING = '\n';

const int RB_SIZE = 1024 * 10;

Hashmap *DATA = NULL;
bstring STORE_PATH = NULL;

void handle_sigchild(int sig) {
    sig = 0; // ignore it
    while(waitpid(-1, NULL, WNOHANG) > 0) {
    }
}

// BUG: this is stupid, use md5
void encipher(unsigned int num_rounds, uint32_t v[2], uint32_t const key[4]) {
    unsigned int i;
    uint32_t v0=v[0], v1=v[1], sum=0, delta=0x9E3779B9;
    for (i=0; i < num_rounds; i++) {
        v0 += (((v1 << 4) ^ (v1 >> 5)) + v1) ^ (sum + key[sum & 3]);
        sum += delta;
        v1 += (((v0 << 4) ^ (v0 >> 5)) + v0) ^ (sum + key[(sum>>11) & 3]);
    }
    v[0]=v0; v[1]=v1;
}

/// TOTALLY RANDOM! LOL  BUG: not secure
const uint32_t STORE_KEY[4] = {18748274, 228374, 193034845, 85726348};
struct tagbstring FAUX64 = bsStatic("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890");

// BUG: this all dies
bstring encrypt_armor_name(bstring name)
{
    // copy the name to encrypt
    bstring encname = bstrcpy(name);
    size_t i = 0;
    // point the encrypt pointer at it
    // BUG: this cast is weird, why?
    uint32_t *v = (uint32_t *)bdata(encname);

    // extend the encname so that it can hold everything
    // BUG: use a correct padding algorithm
    while(blength(encname) % (sizeof(uint32_t) * 2) > 0) {
        bconchar(encname, ' ');
    }

    // run encipher on this
    // BUG: get rid of encipher
    for(i = 0; i < (size_t)blength(encname) / (sizeof(uint32_t) * 2); i+=2) {
        encipher(1, v+i, STORE_KEY);
    }

    // do a lame "base 64" kind of thing on it
    // BUG: this is NOT the best way, it's a quick hack to get it working
    // replace with real BASE64 later
    for(i = 0; i < (size_t)blength(encname); i++) {
        int at = encname->data[i] % blength(&FAUX64);
        encname->data[i] = FAUX64.data[at];
    }

    // that's our final hack encrypted name
    return encname;
}

bstring sanitize_location(bstring base, bstring path)
{
    bstring attempt = NULL;
    bstring encpath = NULL;

    // encrypt armore the name
    // BUG: ditch encryption, it was dumb
    encpath = encrypt_armor_name(path);
    check(encpath != NULL, "Failed to encrypt path name: %s", bdata(path));

    // combine it with the base, this means that we've armored the 
    // path so we can just append it
    attempt = bformat("%s/%s", bdata(base), bdata(encpath));
    bdestroy(encpath);
    return attempt;

error:
    if(encpath) bdestroy(encpath);
    if(attempt) bdestroy(attempt);
    return NULL;
}

int handle_create(Command *cmd, RingBuffer *send_rb, bstring path)
{
    int rc = 0;
    int is_root = biseq(path, cmd->name);
    log_info("create: %s %s %s", bdata(cmd->name), bdata(path), bdata(cmd->number));

    Record *info = Hashmap_get(DATA, path);

    if(info != NULL && is_root) {
        // report if root exists, just skip children
        send_reply(send_rb, &EXISTS);
    } else if(info != NULL) {
        debug("Child %s exists, skipping it.", bdata(path));
        return 0;
    } else {
        // new child so make it
        debug("create: %s %s", bdata(path), bdata(cmd->number));

        Record *info = calloc(1, sizeof(Record));
        check_mem(info);

        // set its stat element
        info->stat = Stats_create();
        check_mem(info->stat);

        // set its name element
        info->name = bstrcpy(path);
        check_mem(info->name);

        // do a first sample
        Stats_sample(info->stat, atof(bdata(cmd->number)));

        // add it to the hashmap
        rc = Hashmap_set(DATA, info->name, info);
        check(rc == 0, "Failed to add data to map.");

        // only send the for the root part
        if(is_root) {
            send_reply(send_rb, &OK);
        }
    }

    return 0;
error:
    return -1;
}

int handle_sample(Command *cmd, RingBuffer *send_rb, bstring path)
{
    // get the info from the hashmap
    Record *info = Hashmap_get(DATA, path);
    int is_root = biseq(path, cmd->name);
    log_info("sample %s %s %s", bdata(cmd->name), bdata(path), bdata(cmd->number));
    bstring child_path = NULL;

    if(info == NULL) {
        // if it doesn't exist then DNE
        send_reply(send_rb, &DNE);
        return 0;
    } else {
        if(is_root) {
            // just sample the root like normal
            Stats_sample(info->stat, atof(bdata(cmd->number)));
        } else {
            // need to do some hackery to get the child path
            // for rolling up mean-of-means on it

            // increase the qty on path up one
            cmd->path->qty++;
            // get the "child path" (previous path?)
            child_path = bjoin(cmd->path, &SLASH);
            // get that info from the DATA
            Record *child_info = Hashmap_get(DATA, child_path);
            bdestroy(child_path);

            // if it exists then sample on it
            if(child_info) {
                // info is /logins, child_info is /logins/zed 
                // we want /logins/zed's mean to be a new sample on /logins
                Stats_sample(info->stat, Stats_mean(child_info->stat));
            }
            // drop the path back to where it was
            cmd->path->qty--;
        }

    }

    // do the reply for the mean last
    bstring reply = bformat("%f\n", Stats_mean(info->stat));
    send_reply(send_rb, reply);
    bdestroy(reply);

    return 0;
}

int handle_delete(Command *cmd, RingBuffer *send_rb, bstring path)
{
    log_info("delete: %s", bdata(cmd->name));
    Record *info = Hashmap_get(DATA, cmd->name);
    check(path == NULL && cmd->path == NULL, "Should be a recursive command.");

    // BUG: should just decide that this isn't scanned 
    // but run once, for now just only run on root
    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        Hashmap_delete(DATA, cmd->name);

        free(info->stat);
        bdestroy(info->name);
        free(info);

        send_reply(send_rb, &OK);
    }

    return 0;
error:
    return -1;
}

int handle_mean(Command *cmd, RingBuffer *send_rb, bstring path)
{
    log_info("mean: %s %s %s", bdata(cmd->name), bdata(path), bdata(path));
    Record *info = Hashmap_get(DATA, path);

    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        bstring reply = bformat("%f\n", Stats_mean(info->stat));
        send_reply(send_rb, reply);
        bdestroy(reply);
    }

    return 0;
}

int handle_stddev(Command *cmd, RingBuffer *send_rb, bstring path)
{
    log_info("stddev: %s %s %s", bdata(cmd->name), bdata(path), bdata(path));
    Record *info = Hashmap_get(DATA, path);

    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        bstring reply = bformat("%f\n", Stats_stddev(info->stat));
        send_reply(send_rb, reply);
        bdestroy(reply);
    }

    return 0;
}

int handle_dump(Command *cmd, RingBuffer *send_rb, bstring path)
{
    log_info("dump: %s, %s, %s", bdata(cmd->name), bdata(path), bdata(path));
    Record *info = Hashmap_get(DATA, path);

    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        bstring reply = bformat("%f %f %f %f %ld %f %f\n",
                Stats_mean(info->stat),
                Stats_stddev(info->stat),
                info->stat->sum,
                info->stat->sumsq,
                info->stat->n,
                info->stat->min,
                info->stat->max);

        send_reply(send_rb, reply);
        bdestroy(reply);
    }

    return 0;
}

int handle_store(Command *cmd, RingBuffer *send_rb, bstring path)
{
    Record *info = Hashmap_get(DATA, cmd->name);
    bstring location = NULL;
    bstring from = cmd->name;
    int rc = 0;
    int fd = -1;

    check(cmd != NULL, "Invalid command.");
    debug("store %s", bdata(cmd->name));
    check(path == NULL && cmd->path == NULL, "Store is non-recursive.");

    if(info == NULL) {
        send_reply(send_rb, &DNE);
    } else {
        // it exists so we sanitize the name
        location = sanitize_location(STORE_PATH, from);
        check(location, "Failed to sanitize the location.");

        // open the file we need with EXLOCK
        fd = open(bdata(location), O_WRONLY | O_CREAT | O_EXLOCK, S_IRWXU);
        check(fd >= 0, "Cannot open file for writing: %s", bdata(location));

        // write the Stats part of info to it
        rc = write(fd, info->stat, sizeof(Stats));
        check(rc == sizeof(Stats), "Failed to write to %s", bdata(location));

        // close, which should release the lock
        close(fd);

        // then send OK
        send_reply(send_rb, &OK);
    }

    return 0;
error: 
    if(fd < 0) close(fd);
    return -1;
}

int handle_load(Command *cmd, RingBuffer *send_rb, bstring path)
{
    bstring to = cmd->arg;
    bstring from = cmd->name;
    bstring location = NULL;
    Record *info = Hashmap_get(DATA, to);
    int fd = -1;

    check(path == NULL && cmd->path == NULL, "Load is non-recursive.");

    if(info != NULL) {
        // don't do it if the target to exists
        send_reply(send_rb, &EXISTS);
    } else {
        location = sanitize_location(STORE_PATH, from);
        check(location, "Failed to sanitize location.");

        // make a new record for the to target
        // TODO: make regular CRUD methods for Record
        info = calloc(1, sizeof(Record));
        check_mem(info);

        info->stat = calloc(1, sizeof(Stats));
        check_mem(info->stat);

        // open the file to read from readonly and locked
        fd = open(bdata(location), O_RDONLY | O_EXLOCK);
        check(fd >= 0, "Error opening file: %s", bdata(location));

        // read into the stats record 
        int rc = read(fd, info->stat, sizeof(Stats));
        check(rc == sizeof(Stats), "Failed to read record at %s", bdata(location));

        // close so we release the lock quick
        close(fd);

        // make a copy of to as the name for the info
        info->name = bstrcpy(to);
        check_mem(info->name);

        // put it in the hashmap
        rc = Hashmap_set(DATA, info->name, info);
        check(rc == 0, "Failed to ass to data map: %s", bdata(info->name));

        // and send the reply
        send_reply(send_rb, &OK);
    }

    return 0;
error:
    if(fd < 0) close(fd);
    return -1;
}

int parse_command(struct bstrList *splits, Command *cmd)
{
    check(splits != NULL, "Invalid split line.");

    // get the command
    cmd->command = splits->entry[0];

    if(biseq(cmd->command, &CREATE)) {
        check(splits->qty == 3, "Failed to parse create: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->number = splits->entry[2];
        cmd->handler = handle_create;
        cmd->path = parse_name(cmd->name);
    } else if(biseq(cmd->command, &MEAN)) {
        check(splits->qty == 2, "Failed to parse mean: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_mean;
        cmd->path = parse_name(cmd->name);
    } else if(biseq(cmd->command, &SAMPLE)) {
        check(splits->qty == 3, "Failed to parse sample: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->number = splits->entry[2];
        cmd->handler = handle_sample;
        cmd->path = parse_name(cmd->name);
    } else if(biseq(cmd->command, &DUMP)) {
        check(splits->qty == 2, "Failed to parse dump: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_dump;
        cmd->path = parse_name(cmd->name);
    } else if(biseq(cmd->command, &DELETE)) {
        check(splits->qty == 2, "Failed to parse delete: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_delete;
        cmd->path = NULL;
    } else if(biseq(cmd->command, &STDDEV)) {
        check(splits->qty == 2, "Failed to parse stddev: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_stddev;
        cmd->path = parse_name(cmd->name);
    } else if(biseq(cmd->command, &STORE)) {
        // store URL
        check(splits->qty == 2, "Failed to parse store: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->handler = handle_store;
        cmd->path = NULL;
    } else if(biseq(cmd->command, &LOAD)) {
        // load FROM TO
        check(splits->qty == 3, "Failed to parse load: %d", splits->qty);
        cmd->name = splits->entry[1];
        cmd->arg = splits->entry[2];
        cmd->handler = handle_load;
        cmd->path = NULL;
    } else {
        sentinel("Failed to parse the command.");
    }

    return 0;
error:
    return -1;
}

int scan_paths(Command *cmd, RingBuffer *send_rb)
{
    check(cmd->path != NULL, "Path was not set in command.");

    int rc = 0;
    // save the original path length
    size_t qty = cmd->path->qty;

    // starting at the longest path, shorten it and call
    // for each one:
    for(; cmd->path->qty > 1; cmd->path->qty--) {
        // remake the path with / again
        bstring path = bjoin(cmd->path, &SLASH);
        // call the handler with the path
        rc = cmd->handler(cmd, send_rb, path);
        // if the handler returns != 0 then abort and return that
        bdestroy(path);

        if(rc != 0) break;
    }

    // restore path length
    cmd->path->qty = qty;
    return rc;
error:
    return -1;
}

struct bstrList *parse_name(bstring name)
{
    return bsplits(name, &SLASH);
}

int parse_line(bstring data, RingBuffer *send_rb)
{
    int rc = -1;
    Command cmd = {.command = NULL};

    // split data on line boundaries
    struct bstrList *splits = bsplits(data, &LINE_SPLIT);
    check(splits != NULL, "Bad data.");

    // parse it into a command
    rc = parse_command(splits, &cmd);
    check(rc == 0, "Failed to parse command.");

    // scan the path and call the handlers
    if(cmd.path) { 
        check(cmd.path->qty > 1, "Didn't give a valid URL.");
        rc = scan_paths(&cmd, send_rb);
        check(rc == 0, "Failure running recursive command against path: %s", bdata(cmd.name));
        bstrListDestroy(cmd.path);
    } else {
        rc = cmd.handler(&cmd, send_rb, NULL);
        check(rc == 0, "Failed running command against path: %s", bdata(cmd.name));
    }

    bstrListDestroy(splits);

    return 0;

error: // fallthrough
    if(cmd.path) bstrListDestroy(cmd.path);
    if(splits) bstrListDestroy(splits);
    return -1;
}

void client_handler(int client_fd)
{
    int rc = 0;
    RingBuffer *recv_rb = RingBuffer_create(RB_SIZE);
    RingBuffer *send_rb = RingBuffer_create(RB_SIZE);

    check_mem(recv_rb);
    check_mem(send_rb);

    // keep reading into the recv buffer and sending on send
    while(read_some(recv_rb, client_fd, 1) != -1) {
        // read a line from the recv_rb
        bstring data = read_line(recv_rb, LINE_ENDING);
        check(data != NULL, "Client closed.");

        // parse it, close on any protocol errors
        rc = parse_line(data, send_rb);
        bdestroy(data); // cleanup here
        check(rc == 0, "Failed to parse user. Closing.");

        // and as long as there's something to send, send it
        if(RingBuffer_available_data(send_rb)) {
            write_some(send_rb, client_fd, 1);
        }
    }

    // close the socket
    rc = close(client_fd);
    check(rc != -1, "Failed to close the socket.");

error: // fallthrough
    if(recv_rb) RingBuffer_destroy(recv_rb);
    if(send_rb) RingBuffer_destroy(send_rb);
    exit(0); // just exit the child process
}

int setup_data_store(const char *store_path)
{
    // a more advanced design simply wouldn't use this
    DATA = Hashmap_create(NULL, NULL);
    check_mem(DATA);

    char *path = realpath(store_path, NULL);
    check(path != NULL, "Failed to get the real path for storage: %s", store_path);

    STORE_PATH = bfromcstr(path);
    free(path);

    return 0;
error:
    return -1;
}

int run_server(const char *host, const char *port, const char *store_path)
{
    int rc = 0;
    struct sockaddr_in client_addr;
    socklen_t sin_size = sizeof(client_addr);
    int server_socket = 0;
    int client_fd = 0;

    rc = setup_data_store(store_path);
    check(rc == 0, "Failed to setup the data store.");

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

.\ex51b\statserve\src\statserve.h

```c
#ifndef _statserve_h
#define _statserve_h

#include <lcthw/bstrlib.h>
#include <lcthw/ringbuffer.h>
#include <lcthw/stats.h>

struct Command;

typedef int (*handler_cb)(struct Command *cmd, RingBuffer *send_rb, bstring path);

typedef struct Command {
    bstring command;
    bstring name;
    struct bstrList *path;
    bstring number;
    bstring arg;
    handler_cb handler;
} Command;

typedef struct Record {
    bstring name;
    Stats *stat;
} Record;

struct tagbstring OK;

int setup_data_store(const char *store_path);

struct bstrList *parse_name(bstring name);

int scan_paths(Command *cmd, RingBuffer *send_rb);

int parse_line(bstring data, RingBuffer *send_rb);

int run_server(const char *host, const char *port, const char *store_path);

bstring sanitize_location(bstring base, bstring path);

bstring encrypt_armor_name(bstring name);

#endif

```

.\ex51b\statserve\tests\statserve_tests.c

```c
#include "minunit.h"
#include <dlfcn.h>
#include "statserve.h"
#include <lcthw/bstrlib.h>
#include <lcthw/ringbuffer.h>
#include <assert.h>

typedef struct LineTest {
    char *line;
    bstring result;
    char *description;
} LineTest;

int attempt_line(LineTest test)
{
    int rc = -1;
    bstring result = NULL;

    bstring line = bfromcstr(test.line);
    RingBuffer *send_rb = RingBuffer_create(1024);

    rc = parse_line(line, send_rb);
    check(rc == 0, "Failed to parse line.");
    result = RingBuffer_get_all(send_rb);
    check(result != NULL, "Ring buffer empty.");
    check(biseq(result, test.result), "Got the wrong output: %s expected %s",
            bdata(result), bdata(test.result));

    bdestroy(line);
    RingBuffer_destroy(send_rb);
    return 1; // using 1 for tests
error:

    log_err("Failed to process test %s: got %s", test.line, bdata(result));
    if(line) bdestroy(line);
    if(send_rb) RingBuffer_destroy(send_rb);
    return 0;
}

int run_test_lines(LineTest *tests, int count)
{
    int i = 0;

    for(i = 0; i < count; i++) {
        check(attempt_line(tests[i]), "Failed to run %s", tests[i].description);
    }

    return 1;
error:
    return 0;
}

int fake_command(Command *cmd, RingBuffer *send_rb, bstring path)
{
    check(cmd != NULL, "Bad cmd.");
    check(cmd->path != NULL, "Bad path.");
    check(send_rb != NULL, "Bad send_rb.");
    check(path != NULL, "Bad path given.");

    return 0;
error:
    return -1;
}

char *test_path_parsing()
{
    struct bstrList *result = NULL;
    struct tagbstring slash = bsStatic("/");
    struct tagbstring logins_zed = bsStatic("/logins/zed");
    struct tagbstring command_name = bsStatic("dump");
    RingBuffer *send_rb = RingBuffer_create(1024);
    struct bstrList *path = bsplits(&logins_zed, &slash);
    int rc = 0;

    Command fake = {
        .command = &command_name,
        .name = &logins_zed,
        .number = NULL,
        .handler = fake_command,
        .path = path
    };

    result = parse_name(&logins_zed);
    mu_assert(result != NULL, "Failed to parse /logins/zed");

    rc = scan_paths(&fake, send_rb); 
    mu_assert(rc != -1, "scan_paths failed.");

    return NULL;
}

char *test_create()
{
    LineTest tests[] = {
        {.line = "create /zed 100", .result = &OK, .description = "create zed failed"},
        {.line = "create /joe 100", .result = &OK, .description = "create joe failed"},

    };

    mu_assert(run_test_lines(tests, 2), "Failed to run create tests.");

    return NULL;
}

char *test_sample()
{
    struct tagbstring sample1 = bsStatic("100.000000\n");

    LineTest tests[] = {
        {.line = "sample /zed 100", .result = &sample1, .description = "sample zed failed."}
    };

    mu_assert(run_test_lines(tests, 1), "Failed to run sample tests.");

    return NULL;
}

char *test_store_load()
{
    LineTest tests[] = {
        {.line = "delete /zed", .result = &OK, .description = "delete zed failed"},
        {.line = "create /zed 100", .result = &OK, .description = "create zed failed"},
        {.line = "store /zed", .result = &OK, .description = "store zed failed"},
        {.line = "load /zed /sam", .result = &OK, .description = "load zed failed"},
        {.line = "delete /sam", .result = &OK, .description = "load zed failed"},
    };

    mu_assert(run_test_lines(tests, 3), "Failed to run sample tests.");

    return NULL;
}

char *test_encrypt_armor_name()
{
    struct tagbstring test1 = bsStatic("/logins");
    struct tagbstring expect1 = bsStatic("vtmTmzNI");
    struct tagbstring test2 = bsStatic("../../../../../../../../etc/passwd");
    struct tagbstring expect2 = bsStatic("pVOBpFjHEIhB7cuT3BGUvyZGn3lvyj226mgggggg");

    bstring result = encrypt_armor_name(&test1);
    debug("Got encrypted name %s", bdata(result));
    mu_assert(biseq(result, &expect1), "Failed to encrypt test2.");
    bdestroy(result);

    result = encrypt_armor_name(&test2);
    debug("Got encrypted name %s", bdata(result));
    mu_assert(biseq(result, &expect2), "Failed to encrypt test2.");
    bdestroy(result);

    return NULL;
}

char *test_path_sanitize_armor()
{
    struct tagbstring base = bsStatic("/tmp");
    struct tagbstring test1 = bsStatic("/somepath/here/there");
    bstring encname = encrypt_armor_name(&test1);
    bstring expect = bformat("%s/%s", bdata(&base), bdata(encname));
    struct tagbstring test2 = bsStatic("../../../../../../../../etc/passwd");

    bstring result = sanitize_location(&base, &test1);
    mu_assert(result != NULL, "Failed to sanitize path.");
    mu_assert(biseq(result, expect), "failed to sanitize test1");

    // this should be pulled up into a tester function
    // BUG: just get rid of this and use md5
    encname = encrypt_armor_name(&test2);
    expect = bformat("%s/%s", bdata(&base), bdata(encname));
    result = sanitize_location(&base, &test2);
    mu_assert(result != NULL, "Failed to sanitize path.");
    mu_assert(biseq(result, expect), "failed to sanitize test1");

    return NULL;
}

char *all_tests()
{
    mu_suite_start();

    int rc = setup_data_store("/tmp");
    mu_assert(rc == 0, "Failed to setup the data store.");

    mu_run_test(test_path_parsing);
    mu_run_test(test_encrypt_armor_name);
    mu_run_test(test_path_sanitize_armor);
    mu_run_test(test_create);
    mu_run_test(test_sample);
    mu_run_test(test_store_load);

    return NULL;
}

RUN_TESTS(all_tests);

```

Final Review

