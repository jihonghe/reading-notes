# 使用正则表达式进行搜索

```sql
/*
REGEXP与LIKE的对比：
    1. LIKE匹配整个字符串，REGEXP匹配子串
    2. REGEXP用于更为复杂的字符串匹配
*/

# 1. 匹配基本字符串
# 匹配 prod_name 字段值中含有'inch'的列
SELECT prod_name FROM products WHERE prod_name REGEXP 'inch' ORDER BY prod_name;

# 匹配 prod_name 字段值中任意字符后跟'ear'的列
SELECT prod_name FROM products WHERE prod_name REGEXP '.oy' ORDER BY prod_name;

# 2. 进行 OR 匹配
# 注意：在使用 '|' 进行 OR 匹配时，'|' 两边的空格也会匹配，也就是'doll | bag | inch'与'doll|bag|inch'是不同的。
# 匹配 prod_name 字段值中含有'doll'或'inch'或'bag'的列
SELECT prod_name FROM products WHERE prod_name REGEXP 'doll|bag|inch';
# 匹配 prod_name 字段值中含有'doll '或' inch'或' bag '的列
SELECT prod_name FROM products WHERE prod_name REGEXP 'doll | bag | inch';

# 3. 匹配几个字符之一
# 使用'[]'进行多值匹配，可将其视为另一种形式的 OR 语句，但又有所不同
SELECT prod_name FROM products WHERE prod_name REGEXP '[0123456789]{1,3} inch';
# 替换成 OR 表达
SELECT prod_name FROM products WHERE prod_name REGEXP '(0|1|2|3|4|5|6|7|8|9){1,3} inch';

# 4. 匹配范围
# 使用'[]'进行字符的范围匹配
SELECT prod_name FROM products WHERE prod_name REGEXP '[0-99] inch';
# 匹配 prod_name 中含有字母范围'o-q'的列
SELECT prod_name FROM products WHERE prod_name REGEXP '[o-q]';

# 5. 匹配转义字符
# 所谓转义字符，即含有特殊意义的字符，需要转义才能匹配到该类字符本身。
# 如在正则表达式中含有特殊意义的字符：'.', '|', '[]', '-'等
# 具有特殊意义的字符，如换行符('\n')、回车('\r')、制表符('\t')、纵向制表符('\v')、换页符('\f')等
# 注意：匹配特殊意义字符是需要使用双反斜杠'\\'，这是MySQL要求的，MySQL自己解析一个'\'，正则表达式自己解析一个'\'

# 匹配含有'.'字符的 vend_name
SELECT vend_name FROM vendors WHERE vend_name REGEXP '\\.';
# 匹配含有换行符的 vend_name
SELECT vend_name FROM vendors WHERE vend_name REGEXP '\\n';

# 6. 匹配一类字符
# 使用MySQL中的预定义字符集，可以方便的地匹配某一类字符。如字母、数字等。

    # a. 匹配任意字母和数字：'[:alnum:]'，与'[a-zA-Z0-9]'有相同效果
SELECT prod_name FROM products WHERE prod_name REGEXP '[:alnum:]';

    # b. 匹配任意字母：'[:alpha:]'，与'[a-zA-Z]'有相同效果
SELECT prod_name FROM products WHERE prod_name REGEXP '[:alpha:]';

    # c. 匹配空格和制表符：'[:blank:]'，与'[\\t]'效果相同
SELECT prod_name FROM products WHERE prod_name REGEXP '[:blank:]';

    # d. 匹配ASCII控制字符：'[cntrl]'，匹配字符ASCII值为[0, 37]和127的控制字符
SELECT prod_name FROM products WHERE prod_name REGEXP '[:cntrl:]';

    # e. 匹配任意数字：'[:digit:]'，与'[0-9]'效果相同
SELECT prod_name FROM products WHERE prod_name REGEXP '[:digit:]';

    # f. 匹配任意小写字母：'[:lower:]'，与'[a-z]'效果相同
SELECT prod_name FROM products WHERE prod_name REGEXP '[:lower:]';

    # g. 匹配任意大写字母：'[:upper:]'，效果与'[A-Z]'相同
SELECT prod_name FROM products WHERE prod_name REGEXP '[:upper:]';

    # h. 匹配任意可打印字符：'[:print:]'，匹配任意可打印字符
SELECT prod_name FROM products WHERE prod_name REGEXP '[:print:]';

    # i. 匹配除空格外，满足'[:print:]'的字符：'[:graph:]'
SELECT prod_name FROM products WHERE prod_name REGEXP '[:graph:]';

    # j. 匹配包括空格在内的任意空白字符：'[:space:]'，效果与'[\\f\\r\\n\\t\\v]'相同
SELECT prod_name FROM products WHERE prod_name REGEXP '[:space:]';

    # k. 匹配既不在'[:alnum:]'也不在'[:cntrl:]'的任意字符：'[:punct:]'
SELECT prod_name FROM products WHERE prod_name REGEXP '[:punct:]';

    # l. 匹配任意十六进制数字：'[:xdigit:]'
SELECT prod_name FROM products WHERE prod_name REGEXP '[:xdigit:]';

# 7. 匹配多个实例
# 所谓匹配多个实例，就是基于正则表达式，设置正则匹配的次数，使其能够找出指定数目的匹配的字符串（如果有）。
# 实现匹配多个实例，可以使用正则表达式配合正则表达式重复元字符完成。
/*    重复元字符
    元字符           说明
    '*'         0个或多个匹配
    '+'         1个或多个匹配
    '?'         0个或1个匹配
   '{n}'        n个匹配
   '{m, n}'     m到n个匹配(最大值为255`)

注意：需要注意的是，这些重复元字符作用的是其前面的匹配字符，如果需要整体匹配多次，则需要花括号包住作为一个整体。
例如：正则'abc?'可以匹配到'ab', 'abc'。因为元字符'?'作用的是其前面的字符'c'。其他的元字符类似。
*/
# 假设数据库中有'TNT 1 stick'和'TNT 5 sticks'，则以下SQL查询均会获取到。
SELECT prod_name FROM products WHERE prod_name REGEXP '\\([0-9] sticks?\\)';

# 用'{}'设置匹配次数
SELECT prod_name FROM products WHERE prod_name REGEXP '[[:digit:]]{4}';
# 以普通方式实现上方的正则匹配
SELECT prod_name FROM products WHERE prod_name REGEXP '[0-9][0-9][0-9][0-9]';

# 8. 定位符
# 所谓定位符，就是用于匹配特定位置的文本是否满足正则表达式
# 常见的定位符有：'^'(文本的开始), '$'(文本的结尾), '[[:<:]]'(词的开始), '[[:>:]](词的结尾)'

    # ATTENTION: '^'具有两个作用，当它在集合'[]'中时表示否定(但在MySQL5.6中不起作用)，其余情况表示定位。
    # a. '^'使用示例：查找products表中的列prod_name以数字开头的列值
SELECT * FROM products WHERE prod_name REGEXP '^[0-9|0-9\\.]';
    # (在MySQL5.6中不起作用)'^'的否定示例：查找products表中的列prod_name不以数字开头的列值
SELECT * FROM products WHERE prod_name REGEXP '[^0-9]';

    # b. '$'使用示例：查找products表中prod_name以'toy'结尾的列值
SELECT * FROM products WHERE prod_name REGEXP '(toy)$' AND prod_name REGEXP '^[a-zA-Z]';

    # c. '[[:<:]]'使用示例：查找products表中的列prod_name的值中的单词存在单词开头是数字的列值
SELECT * FROM products WHERE prod_name REGEXP '[[:<:]][0-9|0-9\\.]';

    # d. '[[:>:]]'使用示例：查找products表中的列prod_name的值中的单词存在单词结尾是'ean'的列值
SELECT * FROM products WHERE prod_name REGEXP '(ean)[[:>:]]';

```