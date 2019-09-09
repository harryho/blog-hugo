+++
title = "MySql Note - 2"
description="MySql Stored Proc  & Function "
+++


### Function 

```sql

DROP FUNCTION IF EXISTS isAusLandLine;
DELIMITER $$
CREATE FUNCTION isAusLandLine (in_value VARCHAR(500)) RETURNS TINYINT
BEGIN
    DECLARE vv_number VARCHAR(500);

    SELECT REPLACE(in_value, '+','') INTO vv_number;

    RETURN CASE WHEN LEFT(vv_number, 2) IN ('02','03','07','08') THEN LENGTH(vv_number) = 10 ELSE 0 END;
END $$

DELIMITER ;

```


### Stored proc sample


```sql

DROP PROCEDURE IF EXISTS stored_proc_1;

DELIMITER $

CREATE PROCEDURE stored_proc_1 (

)
BEGIN

SELECT * FROM table_1;
END


DELIMITER ;

```


