# innodb 事务

## ACID 特性

1. Atomicity：原子性 - 事务被视为不可分割的最小单元，事务的所有操作要么全部成功，要么全部失败回滚。
1. Consistency：一致性 - 数据库在事务执行前后都保持一致性状态，在一致性状态下，所有事务对一个数据的读取结果都是相同的。
1. Isolation：隔离性 - 一个事务所做的修改在最终提交以前，对其他事务是不可见的。
1. Durability：持久性 - 一旦事务提交，则其所做的修改将会永远保存到数据库中。即使系统发生崩溃，事务执行的结果也不能丢。

## 事务的四种隔离级别

在数据库操作中，为了有效保证并发读取数据的正确性，提出的事务隔离级别。我们的数据库锁，也是为了构建这些隔离级别存在的

- **未提交读(Read Uncommitted)**：事务中的修改即使没有提交，对其他事务也是可见的。
- **提交读(Read Committed)**：只能读取到已经提交的数据。Oracle 等多数数据库默认都是该级别 (不重复读)
- **可重复读(Repeated Read)**：保证在同一个事务中多次读取同样数据的结果是一样的。(mysql innodb 默认)
- **串行读(Serializable)**：强制事务串行执行，每次读都需要获得表级共享锁，读写相互都会阻塞。

| 隔离级别                  | 脏读 Dirty Read | 不可重复读 NonRepeatable Read | 幻读 Phantom Read |
| ------------------------- | --------------- | ----------------------------- | ----------------- |
| 未提交读 Read uncommitted | 可能            | 可能                          | 可能              |
| 已提交读 Read committed   | 不可能          | 可能                          | 可能              |
| 可重复读 Repeatable read  | 不可能          | 不可能                        | 可能              |
| 可串行化 Serializable     | 不可能          | 不可能                        | 不可能            |

## MySQL 中锁的种类

- 常见的表锁和行锁
- 新加入的 Metadata Lock 等等,

- 表锁是对一整张表加锁，虽然可分为读锁和写锁，但毕竟是锁住整张表，会导致并发能力下降，一般是做 ddl 处理时使用。
- 行锁则是锁住数据行，这种加锁方法比较复杂，但是由于只锁住有限的数据，对于其它数据不加限制，所以并发能力强。

## 脏读和幻读

- **脏读**：不同事务下，当前事务可以读取到另外事务未提交的数据。

- **幻读**：当某个事务在读取某个范围的记录的时候，另外一个事务又在该范围插入了新的记录，当前事务再次读取这个范围的记录，会产生幻行（Phantom Data）。