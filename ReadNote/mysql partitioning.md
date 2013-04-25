参考连接：
[http://dev.mysql.com/doc/refman/5.1/zh/partitioning.html](http://dev.mysql.com/doc/refman/5.1/zh/partitioning.html)

## Mysql 分区

### 概述
* 通过使用 SHOW VARIABLES LIKE '%partition%' 命令来确定Mysql是否支持分区；
* 分区功能支持Mysql服务器的任何存储引擎；
* 同一个分区表的所有分区必须使用同一个存储引擎；

#### 分区优点
* 与单个磁盘或文件系统分区相比，可以存储更多的数据；
* 对于无意义的数据，可以删除响应的分区；
* 对于增加新数据，可以通过添加新的分区；
* 涉及sum()和count()的查询，很容易地进行并行处理；
* 通过跨多个磁盘来分散数据查询，可以获得最大的查询吞吐量

### 分区类型
* RANGE分区：基于属于一个给定连续区间的列值，把多行分配给分区；
* LIST分区：类似于按RANGE分区，区别在于LIST分区是基于列值匹配一个离散值集合中的某个值来进行选择；
* HASH分区：基于用户定义的表达式的返回值来进行选择的分区，该表达式使用将要插入到表中的这些行的列值进行计算。这个函数可以包含MySQL 中有效的、产生非负整数值的任何表达式；
* KEY分区：类似于按HASH分区，区别在于KEY分区只支持计算一列或多列，且MySQL 服务器提供其自身的哈希函数。必须有一列或多列包含整数值；
* 分区的名字基本上遵循其他MySQL 标识符应当遵循的原则，例如用于表和数据库名字的标识符。但是应当注意，分区的名字是不区分大小写的；

#### RANGE分区
`CREATE TABLE employees (
    id INT NOT NULL,
    fname VARCHAR(30),
    lname VARCHAR(30),
    hired DATE NOT NULL DEFAULT '1970-01-01',
    separated DATE NOT NULL DEFAULT '9999-12-31',
    job_code INT NOT NULL,
    store_id INT NOT NULL
)
PARTITION BY RANGE (store_id) (
    PARTITION p0 VALUES LESS THAN (6),
    PARTITION p1 VALUES LESS THAN (11),
    PARTITION p2 VALUES LESS THAN (16),
    PARTITION p3 VALUES LESS THAN MAXVALUE
)；`

#### LIST分区
`CREATE TABLE employees (
    id INT NOT NULL,
    fname VARCHAR(30),
    lname VARCHAR(30),
    hired DATE NOT NULL DEFAULT '1970-01-01',
    separated DATE NOT NULL DEFAULT '9999-12-31',
    job_code INT,
    store_id INT
)
PARTITION BY LIST(store_id)
    PARTITION pNorth VALUES IN (3,5,6,9,17),
    PARTITION pEast VALUES IN (1,2,10,11,19,20),
    PARTITION pWest VALUES IN (4,12,13,14,18),
    PARTITION pCentral VALUES IN (7,8,15,16)
)；`