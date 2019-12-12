# 分组数据

## 一、为什么需要数据分组

先来看几个问题：以`products`表为例，
![products表](./static/images/products.png)

1. 假设我们需要找出`products`表中的所有供应商，我们可以怎样做？
    1. 非分组方式(用`DISTINCT`关键字实现)：`SELECT DISTINCT vend_id FROM products;`
    2. 分组实现：`SELECT vend_id FROM products GROUP BY vend_id;`
2. 假设我们需要统计某一个供应商提供的产品数，我们可以这样写SQL：
`SELECT COUNT(*) AS prod_num FROM products WHERE vend_id='DLL01';`
3. 现在，我们需要统计`products`表中各个供应商提供的产品数。那么，用2的方法可行吗？显然不可行，接下来看看如何使用分组来达到统计目的：
`SELECT vend_id, COUNT(*) AS num_prods FROM products GROUP BY vend_id;`
查询结果如下(`GROUP BY`即为分组关键字)：
![查询结果1](static/images/group_by_result01.png)

## 二、什么是数据分组

在SQL查询中，分组就是将具有某种共性的数据聚集在一起，以便对数据进行统计分析。

## 三、分组操作

### 1. 创建分组

分组是通过`GROUP BY`关键字完成的，下面是一个简单的例子：
统计`products`表中各个供应商提供的产品数：`SELECT vend_id, COUNT(*) AS num_prods FROM products GROUP BY vend_id;`
查询结果如下(`GROUP BY`即为分组关键字)：
![查询结果1](static/images/group_by_result01.png)
**查询语句分析**：通过查询结果图可以看到，有别于普通的`SELECT`查询，查询所得结果是

### 2. 过滤分组

### 3. 分组与排序

### 4. `SELECT`子句排序
