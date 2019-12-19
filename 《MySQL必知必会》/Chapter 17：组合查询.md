# 组合查询

## 什么是组合查询

顾名思义，组合查询就是将多条`SELECT`语句的查询结果作为单个查询结果集返回。这样的组合查询通常称为并查询或复合查询。

## 为什么需要组合查询

实际上，在单表查询中，组合查询通常可以使用`WHERE`子句替代，但是当`WHERE`子句条件过于复杂时，使用组合查询更能够清晰地表达。在多表查询中，如果各个表查询的数据结构类似，同样可以使用组合查询，一次性获取多条`SELECT`语句的查询结果集。总结起来，有两种基本情况需要使用组合查询：

- 在单个查询中需要从不同的表中返回类似结构的数据
- 对单个表执行多个查询，按照单个查询结果返回数据

## 如何进行组合查询

创建组合查询，我们可以使用`UNION`或这`UNION ALL`来创建组合查询。

### 使用`UNION`创建组合查询

我们可以使用`UNION`关键字来连接多个`SELECT`查询语句的结果集，下面给出一个简单的示例：

```sql
SELECT vend_id, prod_id, prod_price
FROM products
WHERE prod_price < 5
UNION
SELECT vend_id, prod_id, prod_price
FROM products
WHERE vend_id IN (1001, 1002);
```

事实上，正如前面所说，`UNION`可以使用`WHERE`进行替代，如下：

```sql
SELECT vend_id, prod_id, prod_price
FROM products
WHERE prod_price < 5 OR vend_id IN (1001, 1002);
```

但是需要注意的是，当`WEHRE`条件较为复杂，或者需要从多张表中检索相同数据结构类型的数据时，推荐使用组合查询。

### 使用`UNION ALL`创建组合查询

我们知道，可以使用`WHERE`替代`UNION`达到相同的查询效果，但实际上并不总是这样。对于`UNION ALL`的查询结果，我们就不能使用`WHERE`完成。因为通过`UNION ALL`组合查询所得的结果中，可能会存在重复的行。

### 我们可以在组合查询中做些什么

实际上，组合查询所得的结果集与单条`SELECT`语句所得的结果集并无太多差别。基于此，我们可以对组合查询所得的结果集进行排序，分组，数据统计的操作。
