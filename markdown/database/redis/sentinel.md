# Redis

## [中文文档](http://redisdoc.com/topic/sentinel.html) [English Version](http://redis.io/topics/sentinel)

## 哨兵机制简介

1. Sentinel(哨兵) 进程是用于监控 Redis 集群中 Master 主服务器工作的状态
1. 在 Master 主服务器发生故障的时候，可以实现 Master 和 Slave 服务器的切换，保证系统的高可用(High Availability.)
1. 哨兵机制被集成在 Redis2.6+ 的版本中，到了 2.8 版本后就稳定下来了。

## 哨兵进程的作用

1. 监控(Monitoring)：哨兵(sentinel) 会不断地检查你的 Master 和 Slave 是否运作正常。
1. 提醒(Notification)：当被监控的某个 Redis 节点出现问题时, 哨兵(sentinel) 可以通过 API 向管理员或者其他应用程序发送通知。(使用较少)
1. 自动故障迁移(Automatic failover)：当一个 Master 不能正常工作时，哨兵(sentinel) 会开始一次自动故障迁移操作。具体操作如下：
   1. 它会将失效 Master 的其中一个 Slave 升级为新的 Master, 并让失效 Master 的其他 Slave 改为复制新的 Master。
   1. 当客户端试图连接失效的 Master 时，集群也会向客户端返回新 Master 的地址，使得集群可以使用现在的 Master 替换失效 Master。
   1. Master 和 Slave 服务器切换后，Master 的 redis.conf、Slave 的 redis.conf 和 sentinel.conf 的配置文件的内容都会发生相应的改变，即 Master 主服务器的 redis.conf 配置文件中会多一行 slaveof 的配置，sentinel.conf 的监控目标会随之调换。

## 哨兵进程的工作方式

1. 每个 Sentinel（哨兵）进程以每秒钟一次的频率向整个集群中的 Master 主服务器，Slave 从服务器以及其他 Sentinel（哨兵）进程发送一个 PING 命令。（此处我们还没有讲到集群，下一章节就会讲到，这一点并不影响我们模拟哨兵机制）
1. 如果一个实例（instance）距离最后一次有效回复 PING 命令的时间超过 down-after-milliseconds 选项所指定的值， 则这个实例会被 Sentinel（哨兵）进程标记为主观下线（SDOWN）。
1. 如果一个 Master 主服务器被标记为主观下线（SDOWN），则正在监视这个 Master 主服务器的所有 Sentinel（哨兵）进程要以每秒一次的频率确认 Master 主服务器的确进入了主观下线状态。
1. 当有足够数量的 Sentinel（哨兵）进程（大于等于配置文件指定的值）在指定的时间范围内确认 Master 主服务器进入了主观下线状态（SDOWN）， 则 Master 主服务器会被标记为客观下线（ODOWN）。
1. 在一般情况下， 每个 Sentinel（哨兵）进程会以每 10 秒一次的频率向集群中的所有 Master 主服务器、Slave 从服务器发送 INFO 命令。
1. 当 Master 主服务器被 Sentinel（哨兵）进程标记为客观下线（ODOWN）时，Sentinel（哨兵）进程向下线的 Master 主服务器的所有 Slave 从服务器发送 INFO 命令的频率会从 10 秒一次改为每秒一次。
1. 若没有足够数量的 Sentinel（哨兵）进程同意 Master 主服务器下线， Master 主服务器的客观下线状态就会被移除。若 Master 主服务器重新向 Sentinel（哨兵）进程发送 PING 命令返回有效回复，Master 主服务器的主观下线状态就会被移除。
