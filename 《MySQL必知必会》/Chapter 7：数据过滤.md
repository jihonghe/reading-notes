**`WHERE`与`AND`, `OR`, `IN`和`NOT IN`的组合用法**
<br>
**数据表(数据操作)**
![数据表](leanote://file/getImage?fileId=5de65860ab644155d30030b8)
<br>
**操作符**
| 操作符 | 说明 |
| :----: | :--: |
| `AND` | '与'操作符，指前后两个条件均要满足 |
| `OR` | '或'操作符，指前后两个条件至少有一个需要满足 |
| `IN` | 范围操作符，后跟列值集合 |
| `NOT` | 用于否定其后的任何条件 |
<br>
以上操作符通常和`WHERE`组合使用，以下是相关示例：
**`WHERE`与操作符的简单组合**

  - `WHERE`与`AND`
    - `SELECT * FROM students WHERE score > 90 AND class_id = 1;`
    - `SELECT * FROM students WHERE score BETWEEN 90 AND 100;`
  - `WHERE`与`OR`
    - `SELECT * FROM students WHERE class_id = 1 OR class_id = 2;`
  - `WHERE`与`IN`
    - `SELECT * FROM students WHERE class_id IN (1, 2);`
    - `SELECT * FROM students WHERE class_id IN (SELECT class_id FROM students WHERE score BETWEEN 90 AND 100);`
  - `WHERE`与`NOT`
    - `SELECT * FROM students WHERE class_id NOT IN (1, 2);`

<br>
**`WHERE`与操作符的混合组合**
混合组合不再举例，混合组合能够进行复杂的查询，可以自由发挥。