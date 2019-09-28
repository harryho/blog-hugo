+++
title = "JS & ES snippet"
description="JavaScript, ECMAScript snippets"
+++

### Setup the global node modules

* Add an environment variable to tell node where the global node module sits

```bash
# I use nvm to manage my node and node modules
export NODE_PATH=${HOME}/.nvm/versions/node/<node_version>/lib/node_modules
```

* If you use npm by default, you may have permission problem to access the node modules. I suggest you to set a customized global node module folder under your home directory. 

    * Create a new folder under your home directory
    * Install npm to new global node module
    * update environment variables in the profile

    ```bash
    mkdir $HOME/.node_modules
    npm config set prefix $HOME/.node_modules
    npm install -g npm
    echo 'export NODE_MODULES=$HOME/.node_modules' >> $HOME/.profile
    echo 'export PATH=$PATH:$NODE_MODULES/bin' >> $HOME/.profile
    source $HOME/.profile
    ```


### Parsing CSV file

#### Parse csv file with fast-csv

* Install package fast-csv globally

    npm install -g fast-csv

* Assume there is a csv file named test.csv, which contains a few contacts

##### Sample 1 

```js
const fs = require("fs")
const csv = require('fast-csv');

fs.createReadStream('test.csv')
    .pipe(csv.parse({headers:true}))
    .on('error', error => console.error(error))
    .on('data', row => console.log(`ROW=${JSON.stringify(row)}`))
    .on('end', rowCount => console.log(`Parsed ${rowCount} rows`));


```


##### Sample 2

* Validate empty field 

```js
const csv = require('fast-csv')
const fs = require('fs')
const fileStream = fs.createReadStream('test.csv')
const parser = csv.parse({ ignoreEmpty: true, headers: true, trim: true })

const invalidFilter

fileStream
  .pipe(parser)
  .validate(data => data.last_name && data.mobile.startsWith('04'))
  .on('error', error => console.error(error))
  .on('data', row => console.log(`Valid [row=${JSON.stringify(row)}]`))
  .on('data-invalid', (row, rowNumber) => console.log(`Invalid [rowNumber=${rowNumber}] [row=${JSON.stringify(row)}]`))
  .on('end', rowCount => console.log(`Parsed ${rowCount} rows`))
```

#### Sample 3

* Transform the csv's delimiter from  "," to "|" , and combine first two columns into one

```js


```

