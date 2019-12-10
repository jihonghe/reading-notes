**操作符`LIKE`和通配符`%`, `_`的使用**
`LIKE`和`%`，`_`都是用于匹配列值为字符类型的列。
<br>

**数据表（提供数据）**
![数据表](./static/images/students.png)
<br>
| 操作符/通配符 | 说明 |
| :-----------: | :--: |
| `LIKE` | 操作符，用于模糊搜索 |
| `%` | 通配符，匹配出现任意次数的任意字符 |
| `_` | 通配符，匹配单个字符 |
<br>

**示例**

- 使用`%`匹配
  - 匹配名字以"小"开头的学生：`SELECT * FROM students WHERE name LIKE '小%';`
  - 匹配名字以"明"结尾的学生：`SELECT * FROM students WHERE name LIKE '%明';`
  - 匹配名字中含有"红"的学生：`SELECT * FROM students WHERE name LIKE '%红%';`
- 使用`_`匹配
  - 匹配名字以"小"开头且名字为两个字的学生：`SELECT * FROM students WHERE name LIKE '小_';`

<br>

**通配符使用须知**

- 通配符通常用于匹配字符串类型的值
- 使用通配符匹配的内容区分大小写：`J%`与`j%`是不同的
- 通配符`%`无法匹配NULL

<br>

**通配符使用技巧**

- 尽量避免使用通配符，通配符会影响查询效率。
- 在必须使用通配符时，应尽可能地将通配符的相关子句放在SQL的最后面。将通配符放在在查询语句的搜索子句的开始处时，查询效率是最低的。
  - 示例：`SELECT * FROM students (WHERE class_id BETWEEN 1 AND 2) AND name like '%小%';`

<br>

**总结**
操作符`LIKE`与通配符`%`和`_`的结合使用可以完成简单的字符串值的简单搜索匹配工作
