# SimpleGoORM
 
### 链接：
 link: [7天从零实现ORM框架](https://mp.weixin.qq.com/s?__biz=MzUxOTAwNTM2MQ==&mid=2247486563&idx=1&sn=533fae8f1a4aabadd06064de67d5d892&chksm=f98179bdcef6f0ab8ec9b45192ae04e40a209afab07653fba5482c2981f173101ca49c8b7b79&scene=178&cur_album_id=2207687429896110085#rd)

### 内容
主要内容如下表。
 | 步骤 |  用途 |
 | --   | --|
 |SQL库| 和底层数据直接打交到的库函数|
 |Session封装|对库函数的封装，更方便使用|
 |Dialect对象|完成语言基础类型对数据表类型的转换|
 |Schema对象|完成结构体和表结构的转换，并封装基础方法，主要为拼接SQL|
 |Clause封装|完成复杂SQL的拼接，主要是select/where/limit等基础语句的封装，返回sql字符串和对应的数值|
 |Schema对象|通过反射基础借助Clause完成insert和find的功能|
 |Schema对象|通过对象特定的方法返回对象本身，完成链式操作|
 |Schema对象|通过反射查询出hook增加find/insert/update方法|
 |Schema对象|对begin/commit/rollback基础方法的封装，完成事务|
 |Schema对象|通过表结构字段的对比借助事务方法，完成表的迁移|

### 个人理解
- 对SQL库的基础函数进行封装，对SQL的操作变得更加易用。
- 通过Dialect完成语言基础类型对数据库类型的转换，完成转化的第一步。
- 通过Schema完成结构体对数据表类型的转换， 完成转化的第二步。
- 使用Clause将复杂的SQL转化为多个简单的SQL，为后续增强Schema方法做好了准备。
- 通过反射基础完成ORM的基础能力。
- 技术点，使用反射时一定需要理解interface的结构信息。