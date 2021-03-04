# 数据库开发原则

## 一、核心原则

1. 不在数据库做运算
    - 尽量不在数据库做运算
    - 复杂运算移到程序端CPU
    - 尽可能简单应用MYSQL

1. 控制单表数据量
    - 纯INT不超过1000W
    - 含CHAR不超过500W

1. 控制列数量
    - 单表不超50个纯INT字段
    - 单表不超20个CHAR(10)字段
    - IO高效
    - 全表遍历
    - 表修复快
    - 提高并发
    - alter table更快

1. 平衡范式与冗余
    - 第一范式：单个字段不可再分。唯一性。
    - 第二范式：不存在非主属性只依赖部分主键。消除不完全依赖。
    - 第三范式：消除传递依赖。
    - 用一句话来总结范式和冗余：
        - 冗余是以存储换取性能，
        - 范式是以性能换取存储。

1. 拒绝3B
    - 大SQL（BIG SQL）：要减少
    - 大事务（BIG Transaction）
    - 大批量（BIG Batch）

## 二、字段类原则

1. 用好数值字段类型
    - 有符号int最大可以支持到约22亿，远远大于我们的需求和MySQL单表所能支持的性能上限。对于OLTP应用来说，单表的规模一般要保持在千万级别，不会达到22亿上限。如果要加大预留量，可以把主键改为改为无符号int，上限为42亿，这个预留量已经是非常的充足了。
    - 使用bigint，会占用更大的磁盘和内存空间，内存空间毕竟有限，无效的占用会导致更多的数据换入换出，额外增加了IO的压力，对性能是不利的。

1. 将字符转化为数字
    - 更高效
    - 查询更快
    - 占用空间更小
    - 举例：
        用无符号INT存储IP，而非CHAR(15)
        INT UNSIGNED
        可以用INET_ATON()和INET_NTOA()来实现IP字符串和数值之间的转换

1. 优先使用ENUM或SET

    对于一些枚举型数据，推荐优先使用ENUM或SET，这样的场景适合：
    1. 字符串型
    1. 可能值已知且有限

    存储方面：
    1. ENUM占用1字节，转为数值运算
    1. SET视节点定，最多占用8字节
    1. 比较时需要加‘单引号（即使是数值）

    举例：
    ```SQL
    `sex` enum('F','M') COMMENT '性别';
    `c1` enum('0','1','2','3') COMMENT '审核';
    ```

1. 避免使用NULL字段

    为什么在数据库表字段设计的时候尽量都加上NOT NULL DEFAULT ''，这里面不得不说用NULL字段的弊端
    1. 很难进行查询优化
    1. NULL列加索引，需要额外空间
    1. 含NULL复合索引无效
        - 举例:
            - `a` char(32) DEFAULT NULL 【不推荐】
            - `b` int(10) NOT NULL 【不推荐】
            - `c` int(10) NOT NULL DEFAULT 0 【推荐】

1. 少用并拆分TEXT/BLOB

    - TEXT类型处理性能远低于VARCHAR
    - 强制生成硬盘临时表
    - 浪费更多空间
    - VARCHAR(65535)==>64K(注意UTF-8)
    - 尽量不用TEXT/BLOB数据类型
    - 如果业务需要必须用，建议拆分到单独的表

1. 不在数据库里存图片

    - 将图片全部存在数据库，将使得数据库体积变大，会造成读写速度变慢。
    - 对数据库的读/写的速度永远都赶不上文件系统处理的速度
    - 数据库备份变的巨大，越来越耗时间
    - 对文件的访问需要穿越你的应用层和数据库层
    - 推荐处理办法：数据库中保存图片路径

## 三、索引类原则

1. 谨慎合理添加索引
    - 添加索引是为了改善查询
    - 添加索引会减慢更新
    - 索引不是越多越好
    - 能不加的索引尽量不加（综合评估数据密度和数据分布，最好不超过字段数20%）
    - 结合核心SQL有限考虑覆盖索引

1. 字符字段必须建前缀索引
    区分度：
    - 单字母区分度：26
    - 4字母区分度：26*26*26*26 = 456,976
    - 5字母区分度：26*26*26*26*26 = 11,881,376
    - 6字母区分度：26*26*26*26*26*26 = 308,915,776
    - 字符字段必须建前缀索引，例如：
        ```sql
        `pinyin` varchar(100) DEFAULT NULL COMMENT '小区拼音',
        KEY `idx_pinyin` (`pinyin`(8)),
        ) ENGINE=InnoDB
        ```

1. 不在索引列做运算
    - 会导致无法使用索引
    - 会导致全表扫描

1. 自增列或全局ID做INNODB主键
    - 对主键建立聚簇索引
    - 二级索引存储主键值
    - 主键不应更新修改
    - 按自增顺序插入值
    - 忌用字符串做主键
    - 聚簇索引分裂
    - 推荐用独立于业务的AUTO_INCREMENT列或全局ID生成器做代理主键
    - 若不指定主键，InnoDB会用唯一且非空值索引代替

1. 尽量不用外键
    线上OLTP系统尽量不用外键：
    - 外键可节省开发量
    - 有额外开销
    - 逐行操作
    - 可“到达”其他表，意味着锁
    - 高并发时容易死锁

## 四、SQL类原则

1. SQL语句尽可能简单
    - 一条SQL只能在一个CPU运算
    - 拒绝大SQL，拆解成多条简单SQL
        - 简单SQL缓存命中率更高
        - 减少锁表时间，特别是MyISAM
        - 用上多CPU

1. 保持事务（连接）短小
    - 事务/连接使用原则：即开即用，用完即关
    - 与事务无关操作都放到事务外面，减少锁资源的占用
    - 不破坏一致性前提下，使用多个短事务代替长事务
    - 举例：
        - 发帖时的图片上传等待
        - 大量的sleep连接

1. 尽可能避免使用SP/TRIG/FUNC
    - 线上OLTP系统中，我们应当：
        - 尽可能少用存储过程
        - 尽可能少用触发器
        - 减少使用MySQL函数对结果进行处理
        - 将上述这些事情都交给客户端程序负责

1. 尽量不用SELECT *
    - 用SELECT * 时，将会更多的消耗CPU、内存、IO以及网络带宽
    - 我们在写查询语句时，应当尽量不用SELECT * ,只取需要的数据列：
        - 更安全的设计：减少表变化带来的影响
        - 为使用covering index提供可能性
        - Select/JOIN 减少硬盘临时表生成，特别是有TEXT/BLOB时

1. 改写OR为IN()
    - 同一字段，将or改写为in()
        - OR效率：O(n)
        - IN效率：O(Log n）
    - 举例：
        - 不推荐：
            ```sql
            Select * from opp WHERE phone='12347856' or phone='42242233'
            ```
        - 推荐：
            ```sql
            Select * from opp WHERE phone in ('12347856' , '42242233')
            ```
1. 改写OR为UNION
    - 不同字段，将or改为union
        - 减少对不同字段进行 "or" 查询
        - Merge index往往很弱智
        - 如果有足够信心：set global optimizer_switch='index_merge=off';
    - 举例：
        - 不推荐：
            ```sql
            Select * from opp
            WHERE phone='010-88886666'
            or
            cellPhone='13800138000';
            ```
        - 推荐：
            ```sql
            Select * from opp
            WHERE phone='010-88886666'
            union
            Select * from opp
            WHERE cellPhone='13800138000';
            ```

1. 避免负向查询和%前缀模糊查询
    - 在实际开发中，我们要尽量避免负向查询，那什么是负向查询呢，主要有以下：
        - NOT、!=、<>、!<、!>、NOT EXISTS、NOT IN、NOT LIKE等
    - 同时，我们还要避免%前缀模糊查询，因为这样会使用B+ Tree，同时会造成使用不了索引，并且会导致全表扫描，性能和效率可想而知

1. LIMIT高效分页
    ```sql
    #先使用程序获取ID：
    select id from table limit 10000,10;
    #再用in获取ID对应的记录
    Select * from table WHERE id in (123,456…) ;
    ```

1. 用UNION ALL 而非UNION
    - 如果无需对结果进行去重，仅仅是对多表进行联合查询并展示，则用UNION ALL，因为UNION有去重开销
    - 举例：
        ```sql
        SELECT * FROM detail20091128 UNION ALL
        SELECT * FROM detail20110427 UNION ALL
        SELECT * FROM detail20110426 UNION ALL
        SELECT * FROM detail20110425 UNION ALL
        SELECT * FROM detail20110424 UNION ALL
        SELECT * FROM detail20110423;
        ```

1. 分解联接保证高并发
    - 高并发DB不建议进行两个表以上的JOIN
    - 适当分解联接保证高并发：
        - 可缓存大量早期数据
        - 使用了多个MyISAM表
        - 对大表的小ID IN()
        - 联接引用同一个表多次

1. GROUP BY 去除排序
    - 使用GROUP BY可以实现分组和自动排序
    - 无需排序：Order by NULL
    - 特定排序：Group by DESC/ASC

1. 打散大批量更新
    - 大批量更新尽量凌晨操作，避开高峰
    - 凌晨不限制
    - 白天上线默认为100条/秒（特殊再议）
    - 举例：
        ```SQL
        update post set tag=1 WHERE id in (1,2,3);
        sleep 0.01;
        update post set tag=1 WHERE id in (4,5,6);
        sleep 0.01;
        ……
        ```

1. Know Every SQL
    - 作为DBA乃至数据库开发人员，我们必须对数据库的每条SQL都非常了解，常见的命令有：
        - SHOW PROFILE
        - MYSQLsla
        - MySQLdumpslow
        - explain
        - Show Slow Log
        - Show Processlist
        - SHOW QUERY_RESPONSE_TIME(Percona）

## 五、约定类原则

1. 隔离线上线下
    - 构建数据库的生态环境，确保开发无线上库操作权限
    - 原则：线上连线上，线下连线下
        - 生产数据用pro库
        - 预生产环境用pre库
        - 测试用test库
        - 开发用dev库

1. 禁止未经DBA确认的子查询
    - 大部分情况优化较差
    - 特别WHERE中使用IN id的子查询
    - 一般可用JOIN改写
    - 举例：
        ```SQL
        MySQL> select * from table1 where id in (select id from table2);
        MySQL> insert into table1 (select * from table2); //可能导致复制异常
        ```
