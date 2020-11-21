# 实现MySQL主从复制需要进行的配置

## 1. master node

1. 开启二进制日志 binlog
2. 配置唯一的server-id
3. 获得master二进制文件名及位置
4. 创建一个用于slave和master通信的用户账号

5. /etc/my.cnf

    ```ini
    [mysqld]
    # 开启二进制日志
    log-bin=mysql-bin
    # 设置server-id，建议使用ip最后3位
    server-id=140
    ```

## 2. slave node

1. 配置唯一的server-id
2. 使用master分配的用户账号读取master二进制日志
3. 启动slave服务

4. /etc/my.cnf

    ```ini
    #开启中继日志
    relay-log=mysql-relay
    #设置server-id，建议使用ip最后3位
    server-id=141
    ```

## 3.重启mysql服务

* systemctl restart mysqld.service
* docker restart ${container id}

## 4. 在主机上建立账户并授权slave

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
