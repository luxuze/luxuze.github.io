# MySQL 主从复制

## 线程

主要涉及三个线程：binlog 线程、I/O 线程和 SQL 线程。

1. binlog 线程 ：负责将主服务器上的数据更改写入二进制日志（Binary log）中。
2. I/O 线程 ：负责从主服务器上读取- 二进制日志，并写入从服务器的中继日志（Relay log）。
3. SQL 线程 ：负责读取中继日志，解析出主服务器已经执行的数据更改并在从服务器中重放（Replay）。

## 读写分离

主服务器处理写操作以及实时性要求比较高的读操作，而从服务器处理读操作。 读写分离能提高性能的原因在于：

1. 主从服务器负责各自的读和写，极大程度缓解了锁的争用；
2. 从服务器可以使用 MyISAM，提升查询性能以及节约系统开销；
3. 增加冗余，提高可用性。

读写分离常用代理方式来实现，代理服务器接收应用层传来的读写请求，然后决定转发到哪个服务器。

## 1. master node

1. 开启二进制日志 binlog
2. 配置唯一的 server-id
3. 获得 master 二进制文件名及位置
4. 创建一个用于 slave 和 master 通信的用户账号
5. /etc/my.cnf

   ```text
   [mysqld]
   # 开启二进制日志
   log-bin=mysql-bin
   # 设置server-id，建议使用ip最后3位
   server-id=140
   ```

## 2. slave node

1. 配置唯一的 server-id
2. 使用 master 分配的用户账号读取 master 二进制日志
3. 启动 slave 服务
4. /etc/my.cnf

   ```text
   #开启中继日志
   relay-log=mysql-relay
   #设置server-id，建议使用ip最后3位
   server-id=141
   ```

## 3.重启 mysql 服务

* systemctl restart mysqld.service
* docker restart ${container id}

## 4. 在主机上建立账户并授权 slave

```sql
GRANT REPLICATION SLAVE ON *.* TO 'mysql141'@'192.168.131.141' IDENTIFIED BY 'mysql141';

flush privileges;

--查询master的状态，输出File和Position
show master status\G
```

## 5. 告知从服务器二进制文件名与位置

```sql
CHANGE MASTER TO master_host = '192.168.131.140',
 master_user = 'mysql141',
 master_password = 'mysql141',
 master_log_file = 'mysql-bin.000001',
 master_log_pos = 120;
```

## 6. 查看从服务器状态

```sql
--开启复制
start slave;

--查看主从复制是否配置成功, 当看到Slave_IO_State:Waiting for master ot send event 、Slave_IO_Running: YES、Slave_SQL_Running: YES才表明状态正常。
SHOW SLAVE STATUS\G
```

## 读写分离中间件

* Maxsacl
* Mycat
* ...

