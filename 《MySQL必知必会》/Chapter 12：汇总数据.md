# 聚集函数

## 1. 为什么需要聚集函数

想象一下，当我们需要确定表中的总记录数、获得某一列的平均值或最值，而这些数据并不能直接从数据表中检索出来时，我们就需要用到聚集函数。聚集函数的作用就是为数据表中提供特定计算，从而达到分析数据、统计数据的目的。

## 2. 什么是聚集函数

所谓聚集函数，就是运行在SQL语句中，用于计算和返回单个值的函数。以下是常用的聚集函数：
| 函数 | 说明 |
| :--: | :--- |
| `AVG()` | 计算某列的平均值 |
| `COUNT()` | 返回（符合某种过滤条件的）某列的行数 |
| `MAX()` | 返回某列的最大值 |
| `MIN()` | 返回某列的最小值 |
| `SUM()` | 返回某列之和 |
**注意**：聚集函数通常与`SELECT`语句搭配使用，也就是说函数实际统计的是`SELECT`语句选择所得的数据，而非全部数据。

## 3. 如何使用聚集函数

- **`AVG()`**
  - `AVG()`使用须知：
    - 要求指定列名：该函数作用于数值类型的列并且必须指定列名
    - 只作用于单列：以指定的列名为入参，返回计算结果。若要计算多个列的平均值，则需要多个`AVG()`函数
    - 忽略列值为`NULL`的行：也就是说，在计算平均值时不会将对应的列值为`NULL`的行计算在其中
    - 查询计算所的的平均值只返回一行结果
  - 使用示例：
    - 查询单列的平均值：`SELECT AVG(prod_price) FROM products;`
    - 查询存在过滤条件的列平均值：`SELECT AVG(prod_price) AS price_avg FROM products WHERE prod_price > 5;`
    - 查询多列平均值：`SELECT AVG(quantity), AVG(item_price) FROM order_items;`

- **`COUNT()`**
  - `COUNT()`使用须知：
    - 不指定列名时统计的是表中的所有记录数，指定列名时统计的是特定的列数
    - `NULL`值的处理方式：当`COUNT()`中指定了列时，会忽略值为`NULL`的列；当为`COUNT(*)`时则会将其包括进来
  - 使用示例：
    - 使用方式1，统计表中的总记录数：`SELECT COUNT(*) FROM vendors;`
    - 使用方式2，统计表中特定的列总数：
      - `SELECT COUNT(vend_state) FROM vendors;`
      - `SELECT COUNT(vend_state) FROM vendors WHERE vend_state='NY';`

- **`MAX()`**
  - `MAX()`使用须知：
    - 必须指定列名
    - 作用的数据类型不局限于数值类型
    - 忽略列值为`NULL`的行
  - 使用示例：
    - 字符串类型：`SELECT MAX(prod_name) FROM products;`
    - 数值类型：`SELECT MAX(prod_price) FROM products;`

- **`MIN()`**
  用法与`MAX()`相同，不再赘述

- **`SUM()`**
  - `SUM()`使用须知：
    - 必须指定列名
    - 只作用于数值类型
  - 使用示例：
    - 统计单列总和：`SELECT SUM(item_price) FROM order_items;`
    - 多列值合计计算：`SELECT SUM(item_price * quantity) FROM order_items;`

## 聚集不同的函数

所谓聚集不同值，即通过筛选出同列不同值在使用函数处理。要达到这个目的，可以使用`DISTINCT`。在此之前，我们需要另外了解一个关键字`ALL`，我们直接从SQL入手：
`SELECT SUM(ALL item_price) FROM order_items;`
如上查询语句，在`SUM`函数中对列`item_price`加上了`ALL`约束。而实际上，`ALL`是默认的，无需特别添加。
再来看看`DISTINCT`与聚集函数的结合：`SELECT SUM(DISTINCT item_price) FROM order_items;`

## 组合聚集函数

组合聚集函数就是将不同的聚集函数组合在一起进行数据的统计分析，示例如下：
`SELECT COUNT(*) AS num_items, MIN(prod_price) AS price_min, MAX(prod_price) AS price_max, AVG(prod_price) AS price_avg FROM products;`
