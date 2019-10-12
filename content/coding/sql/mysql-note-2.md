+++
title = "MySql Note - 2"
description="MySql Stored Proc  & Function "
+++

## Function - UDF

>  For the UDF mechanism to work, functions must be written in C or C++ and your operating system must support dynamic loading. MySQL source distributions include a file sql/udf_example.cc that defines five UDF functions. Consult this file to see how UDF calling conventions work. The include/mysql_com.h header file defines UDF-related symbols and data structures, although you need not include this header file directly; it is included by mysql.h.

> A UDF contains code that becomes part of the running server, so when you write a UDF, you are bound by any and all constraints that apply to writing server code. For example, you may have problems if you attempt to use functions from the libstdc++ library. T

### Function sample 

* Check the input string is OZ land line 

```sql
DROP FUNCTION IF EXISTS isAusLandLine;
DELIMITER $$
CREATE FUNCTION isAusLandLine (in_value VARCHAR(500)) RETURNS TINYINT
BEGIN
    DECLARE vv_number VARCHAR(500);

    SELECT REPLACE(in_value, '+','') INTO vv_number;

    RETURN 
      CASE 
          WHEN LEFT(vv_number, 2) IN ('02','03','07','08') ss
          THEN LENGTH(vv_number) = 10 ELSE 0 
      END;
END $$

DELIMITER ;

-- SELECT isAusLandLine('03123322') from dual; -- TRUE

-- SELECT isAusLandLine('06123322') from dual; -- FALSE

```

* EXtract the Json to String


```sql
DROP FUNCTION IF EXISTS json_extract_string;

DELIMITER $

CREATE FUNCTION `json_extract_string`(
  p_json text,
  p_key text
) RETURNS varchar(40) CHARSET latin1
  BEGIN
    SET p_json = replace(p_json, '\\"', '"');
    SET p_json = replace(p_json, '" :', '":');
    SET p_json = replace(p_json, ': "', ':"');
    SET p_json = replace(p_json, ': [', ':[');

    SET @pattern_start_type = '"';
    SET @pattern_end_type = '"';
    SET @pattern = CONCAT('"', p_key, '":',@pattern_start_type);

    IF LOCATE(@pattern, p_json) > 0 THEN
      SET @start_i = LOCATE(@pattern, p_json) + CHAR_LENGTH(@pattern);
    ELSE
      SET @pattern_start_type = '[';
      SET @pattern_end_type = ']';
      SET @pattern = CONCAT('"', p_key, '":',@pattern_start_type);
      SET @start_i = LOCATE(@pattern, p_json) + CHAR_LENGTH(@pattern);
    END IF;

    IF @start_i = CHAR_LENGTH(@pattern) THEN
      SET @end_i = 0;
    ELSE
      SET @end_i = LOCATE(@pattern_end_type, p_json, @start_i) - @start_i;
    END IF;
    RETURN SUBSTR(p_json, @start_i, @end_i);
  END $

DELIMITER ;

-- SELECT json_extract_string('{"key": 123}' ) from dual; -- 123

```

## Stord procedure

> MySQL supports stored routines (procedures and functions). A stored routine is a set of SQL statements that can be stored in the server. Once this has been done, clients don't need to keep reissuing the individual statements but can refer to the stored routine instead. 

 Stored routines can be particularly useful in certain situations:

* When multiple client applications are written in different languages or work on different platforms, but need to perform the same database operations.

* When security is paramount. Banks, for example, use stored procedures and functions for all common operations. This provides a consistent and secure environment, and routines can ensure that each operation is properly logged. In such a setup, applications and users would have no access to the database tables directly, but can only execute specific stored routines. 




### Stored proc sample


* Create a stored proc to execute dynamic SQL script based on the previous execution result 

```sql
DROP PROCEDURE IF EXISTS RunIf;

DELIMITER $$

CREATE PROCEDURE RunIf(ifExpr MEDIUMTEXT, execStmt MEDIUMTEXT)
  BEGIN
    SET @sql = concat('select @result := (', ifExpr, ')');
    PREPARE stmt from @sql;
    EXECUTE stmt;
    DEALLOCATE prepare stmt;
    IF (@result = true) THEN
      SET @sql = execStmt;
      PREPARE stmt FROM @sql;
      EXECUTE stmt;
      DEALLOCATE prepare stmt;
    END IF;
  END
$$

DELIMITER ;

-- CALL( CALL RunIf('EXISTS (SELECT * FROM information_schema.columns 
-- WHERE table_schema = DATABASE() AND table_name = \'TargeTable\' 
-- AND column_name = \'description\')',
--  'ALTER TABLE TargeTable DROP COLUMN description');)


```


