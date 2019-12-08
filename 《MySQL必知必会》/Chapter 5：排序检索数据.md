**术语介绍**
  - ORDER BY子句：用于排序检索出的数据，通常与SELECT语句结合使用。*注意：`ORDER BY`后的列可以不在`SELECT`子句中*
  - ASC：升序，大多数数据库的默认排序方式，也就是说使用`ORDER BY`时可以不指定排序规则即为升序排序
  - DESC：降序，在排序时需要显式指定降序排序

<br>
**数据表(提供数据)**
![数据表](leanote://file/getImage?fileId=5de4fdd3ab6441193f006d7d)
<br>
**无序检索数据**
通常，在无序检索数据表中的单列数据时，检索出的结果是按照其插入数据表中的顺序排列的。但是，数据的顺序会因更新、删除等操作而改变。故在无序查询数据时，不该默认数据是有序的。
<br>
**有序检索数据**
对于通过`SELECT`语句检索出的数据，可以用子句`ORDER BY`进行排序。此处的排序按照排序的列数分两种情况：

  - 单列排序
    - 单列升序排序
      - `SELECT * FROM students ORDER BY class_id;`
      - `SELECT name, gender FROM students ORDER BY class_id;`
    - 单列降序排序
      - `SELECT * FROM students ORDER BY class_id DESC;`
  - 多列排序
    - 多列降序排序
      - 多列降序中需要对每一列都显式指明采用降序排序，未指明均默认为升序：`SELECT * FROM students ORDER BY class_id DESC, score DESC;`
    - 多列升序排列
      - `SELECT * FROM students ORDER BY class_id, score;`
    - 多列混合排序
      - `SELECT * FROM students ORDER BY class_id, score DESC, gender;`

<br>
**`FROM`, `ORDER BY`, `LIMIT`之间的顺序问题**
`FROM` --> `ORDER BY` --> `LIMIT`，三者的顺序若是出错则会引起SQL错误。