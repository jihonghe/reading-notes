# 创建计算字段

## 计算字段

- **什么是计算字段**：计算字段是运行时在`SELECT`语句内创建的，其值由数据表中的列值处理所得，计算字段并不属于数据表列。
- **为什么需要计算字段**：当存储于数据表中的数据无法直接满足需求，而是需要通过一定的数据处理，此时可以考虑创建计算字段。当然，也可以先检索出原数据列值，在通过程序进行计算。但相较程序计算而言，同样的计算结果，通过SQL计算字段更为方便。
- **MySQL中的计算字段与列值**：在DBMS中，DBMS并不能辨别出计算字段与列值，只有服务端的数据库可以辨别出来。一般而言，这并不会影响查询语句的编写，此处仅仅简单记录。

## 创建计算字段
通常，我们可以通过处理数据表列来生成新值，从而创建计算字段。例如，通过concat()函数拼接字符串类型的列值与其他字符串生成新的字段值；也可以通过字符串处理函数对列值进行格式化等生成字段。对于生成的字段值，我们需要一个字段对应，而这就是创建计算字段的过程。

**一个简单的示例**：`SELECT Concat(vend_name, '-', vendor-country) FROM vendors;`，该查询语句的返回结果生成计算字段`CONCAT(vender_name, '-', vendor-country)`，一般而言，如果不指定别名，计算字段名称即为字段计算的字符串形式。
<br>
下面介绍一些可用于字段计算的函数、关键字以及运算符。
| 名称 | 说明 | 示例 |
| :----: | :----: | :----- |
| `CONCAT()` |字符串拼接函数 | `SELECT CONCAT('(', vend_name, ')') FROM vendors;` |
| `Ttrim()` | 去除字符串两端空格 | `SELECT Trim(vend_name) FROM vendors;` |
| `LTrim()` | 去除字符串左边的空格 | `SELECT LTrim(vend_name) FROM vendors;` |
| `RTrim()` | 去除字符串右边的空格 | `SELECT RTrim(vend_name) FROM vendors;` |
| `AS` | 关键字，用于设置别名，可作用于表、列、计算字段等等 | `SELECT Trim(vend_name) AS processed_vend_name FROM vendors;` |
| `+`, `-`, `*`, `/` | 算术运算符，用于数值类型的列值计算 | `SELECT prod_id, quantity, item_price, quantity * item_price AS expanded_price FROM order_items WHERE order_num = 20005;` |
