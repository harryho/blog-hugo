+++
title = "MySql: JSON"
description="Json Support "
+++

## Json support 

As of MySQL 5.7.8, MySQL supports a native JSON data type defined by RFC 7159 that enables efficient access to data in JSON (JavaScript Object Notation) documents. The JSON data type provides these advantages over storing JSON-format strings in a string column:

* Automatic validation of JSON documents stored in JSON columns. Invalid documents produce an error.

* Optimized storage format. JSON documents stored in JSON columns are converted to an internal format that permits quick read access to document elements. When the server later must read a JSON value stored in this binary format, the value need not be parsed from a text representation. The binary format is structured to enable the server to look up subobjects or nested values directly by key or array index without reading all values before or after them in the document.



### Caveats

* The size of any JSON document stored in a JSON column is limited to the value of the max_allowed_packet system variable. 

* A JSON column cannot have a non-NULL default value.


### Json functions

> Along with the JSON data type, a set of SQL functions is available to enable operations on JSON values, such as creation, manipulation, and searching. 


Name |  Description
:--|:--|
->  | Return value from JSON column after evaluating path; equivalent to JSON_EXTRACT().
->>  | Return value from JSON column after evaluating path and unquoting the result; equivalent to JSON_UNQUOTE(JSON_EXTRACT()).
JSON_APPEND() |  (deprecated 5.7.9)  Append data to JSON document
JSON_ARRAY() |   Create JSON array
JSON_ARRAY_APPEND() |   Append data to JSON document
JSON_ARRAY_INSERT() |   Insert into JSON array
JSON_CONTAINS() |   Whether JSON document contains specific object at path
JSON_CONTAINS_PATH() |   Whether JSON document contains any data at path
JSON_DEPTH() |   Maximum depth of JSON document
JSON_EXTRACT() |   Return data from JSON document
JSON_INSERT() |   Insert data into JSON document
JSON_KEYS() |   Array of keys from JSON document
JSON_LENGTH() |   Number of elements in JSON document
JSON_MERGE() |  (deprecated 5.7.22)  Merge JSON documents, preserving duplicate keys. Deprecated synonym for JSON_MERGE_PRESERVE() | 
JSON_MERGE_PATCH() |   Merge JSON documents, replacing values of duplicate keys
JSON_MERGE_PRESERVE() |   Merge JSON documents, preserving duplicate keys
JSON_OBJECT() |   Create JSON object
JSON_PRETTY() |   Print a JSON document in human-readable format
JSON_QUOTE() |   Quote JSON document
JSON_REMOVE() |   Remove data from JSON document
JSON_REPLACE() |   Replace values in JSON document
JSON_SEARCH() |   Path to value within JSON document
JSON_SET() |   Insert data into JSON document
JSON_STORAGE_SIZE() |   Space used for storage of binary representation of a JSON document
JSON_TYPE() |   Type of JSON value
JSON_UNQUOTE() |   Unquote JSON value
JSON_VALID() |   Whether JSON value is valid


### Export table to json file

* Return json object

```sql
---  Return record as json in query
SELECT 
json_object( 
   'entityId',  entityId,
   'categoryName',  categoryName,
   'description',  description,
   'picture',  picture
 ) as json, 'Category' 
  FROM  Category 
  WHERE entityId = 1 ; 
```

* Json object result

```json
{
  "picture": null,
  "entityId": 1,
  "description": "Soft drinks, coffees, teas, beers, and ales",
  "categoryName": "Beverages"
}

```

* Return record as json array


```sql
-- Use json_array and json_object functions 
DROP TEMPORARY TABLE IF EXISTS tmp_json_data;

CREATE TEMPORARY TABLE tmp_json_data (
  jsonText TEXT
) 
SELECT 
json_object( 
   'entityId',  entityId,
   'categoryName',  categoryName,
   'description',  description,
   'picture',  picture
 ) as json 
  INTO jsonText
FROM  Category ; 
  
SET SESSION  group_concat_max_len = 9999;

SELECT concat('[', group_concat(jsonText),']') jsonArray
FROM tmp_json_data;
```

* Json array result

```json
[
    {
      "categoryName" : "Beverages",
      "description" : "Soft drinks, coffees, teas, beers, and ales",
      "entityId" : 1,
      "picture" : null
    },
    {
      "categoryName" : "Condiments",
      "description" : "Sweet and savory sauces, relishes, spreads, and seasonings",
      "entityId" : 2,
      "picture" : null
    },
    {
      "categoryName" : "Confections",
      "description" : "Desserts, candies, and sweet breads",
      "entityId" : 3,
      "picture" : null
    }
]

```
