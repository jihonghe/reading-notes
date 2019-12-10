# 使用数据处理函数

## MySQL中的函数

- **为何需要函数**：函数为数据处理与转换提供了便捷高效的处理方式，MySQL中的函数能够满足大部分数据处理与转换需求。

- **MySQL的函数分类**：根据操作的值的类型分为以下几类
  - 用于处理字符串的文本函数
  - 用于数值计算的数值函数
  - 用于处理日期与时间的日期与时间函数
  - 用于获取DBMS信息的特殊函数

## 使用函数

- **文本处理函数**
  - 大小写转换
    - `Upper()`：将所有字母转为大写形式
    - `Lower()`：将所有字母转为小写形式
  - 空格处理
    - `Trim()`：去除字符串两端空格
    - `LTrim()`：去除字符串左边的空格
    - `RTrim()`：去除字符串右边的空格
  - 查找与统计
    - `Length()`：计算并返回字符串长度
    - `Locate()`：找出指定子串在字符串中的位置
    - `SubString()`：返回字符串指定范围的子串
    - `Left()`：返回左边指定数目的字符
    - `Right()`：返回右边指定数目的字符
    - `Soundex()`：略，在本文档中不提，可自行查找。
<br>
  - **使用示例**
    此处仅给出部分函数的使用示例
    - `Locate()`：
      - 用法一：`LOCATE(sub_str, str)`
        - `SELECT LOCATE('toy', prod_name) AS toy_position FROM products;`，其返回位置为子串中的第一个字符在目标字符串中的位置。如：`Fish bean bag toy`中`toy_position`为15。
      - 用法二：`LOCATE(sub_str, str, pos)`，`pos`用于指定查找的开始位置，`pos >= 1`
        - `SELECT LOCATE('toy', prod_name, 4) AS toy_position FROM products;`
    - `SubString()`：
      此处仅给出一个用法，其余三种用法不在此给出
      - `SUBSTRING(str, pos)`：`SELECT SUBSTRING(prod_name, 4) AS toy_position FROM products;`
    - `Left()`：
      - `SELECT LEFT(prod_name, 4) AS left_part FROM products;`
<br>
- **日期和时间处理函数**
  - 日期函数
    在MySQL中，日期的首选格式为：`YYYY-MM-DD`，该种格式的好处在于排除了多义性，明确了年、月、日的位置，并且在SQL的查询、插入或者更新中只支持该格式。
    - `Date()`：返回日期时间中的日期部分
    - `Year(), Month(), Day()`：这三个函数的作用是从日期中分别取出年、月、日部分并返回
    - `CurDate()`：返回当前日期
      - 示例：`SELECT CURDATE();`
  - 时间函数
    在MySQL中，时间的格式为：`HH:MM:SS`
    - `Time()`：返回日期时间中的时间部分
    - `Hour(), Minute(), Second()`：这三个函数的作用是从时间中分别取出时、分、秒部分并返回
    - `CurTime()`:返回当前时间
      - 示例：`SELECT CURTIME();`
  - 其他常用函数
    - `Now()`：返回当前日期与时间
      - 示例：`SELECT NOW();`
  - 日期和时间计算函数
    - `AddDate(date: date, days: int)`：增加一个日期
    - `AddTime()`：增加一个时间
    - `Date_Add()`：略
    - `DayOfWeek()`：返回一个指定日期对应的星期几
    - `DateDiff(date1: data, date2: date)`：计算两个日期之差，`date1 - date2`
      - 示例：`SELECT DATEDIFF('2019-11-12', '2019-12-20');`，返回结果为`-38`
  - 日期时间格式化
    - `Date_Format()`：返回一个格式化的日期或时间字符串
<br>
  - **使用示例**
    - 日期查询：查询指定日期的订单，对比一下以下两种写法
      - 写法一：`SELECT * FROM orders WHERE order_date='2004-05-01';`
      - 写法二：`SELECT * FROM orders WHERE DATE(order_date)='2004-05-01';`
      以上两种写法的区别在于，写法二使用了`DATE()`函数来获取日期。一般而言，在order_date字段中，若只有日期，则两种的效果等同；若还存在时间，则只有写法二能满足预期的需求。同样的，对于过滤时间，用`TIME()`函数也是类似的效果。
    - 查询日期、时间范围：
      - 查询某个月的所有订单：同样提供集中写法以供比较
        - 写法一(使用`BETWEEN ... AND ...`)：`SELECT * FROM orders WHERE order_date BETWEEN '2004-05-01' AND '2004-05-31';`
        - 写法二(使用`Month()`和`Year()`函数)：`SELECT * FROM orders WHERE YEAR(order_date)=2004 AND MONTH(order_date)=5;`
        显然，以上两种写法中写法二更为方便，无需知道指定月的天数。而这，也恰好验证了使用函数的便捷性。

<br>

- **数值处理函数**
  数值处理函数主要用于代数计算、三角函数计算及几何计算。由于使用的相对不频繁，故此处仅给出函数及函数说明，不提供示例。

| 函数 | 说明 |
| :-: | :--- |
| Abs() | 返回一个数的绝对值 |
| Mod() | 返回除操作所得的余数 |
| Sqrt() | 返回一个数的平方根 |
| Exp() | 返回一个数的指数值 |
| Pi() | 返回圆周率 |
| Rand() | 返回一个随机数 |
| Sin() | 返回一个数的正弦 |
| Cos() | 返回一个数的余弦 |
| Tan() | 返回一个数的正切 |
