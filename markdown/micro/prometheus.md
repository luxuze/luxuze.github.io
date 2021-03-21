# Prometheus

1. “监控的覆盖率”，所有系统所有层面，不一定一定要有指标，也不是具体信息 label 分得越细越好，最后搞出几千个监控项，不仅搞得眼花缭乱还让 Prometheus 变慢了；

2. 而”警报的覆盖率”，事无巨细必有要有警报，人人有责全体收警报（有些警报会发送给几十个人）。最后当然你也能预想到了，告警风暴让大家都对警报疲劳了；

3. 这些事情乍看起来都是在努力工作，但其实一开始的方向就错了，监控的目标绝对不是为了达到 xxx 个指标，xxx 条警报规则，这些东西有什么意义？依我看，负责监控的开发就算不是 SRE 也要有 SRE 的心态和视野，不要为监控系统的功能或覆盖面负责（这样很可让导致开发在监控里堆砌功能和内容，变得越来越臃肿越来越不可靠），而要为整个业务的稳定性负责，同时站在稳定性的投入产出比角度去考虑每件事情的性质和意义，不要忘记我们因何而来。

## Docker-Compose

* 这里由于官网就是v2的文件，暂时不做升级
* [prom & grafana 配置](https://github.com/luxuze/Prometheus-Demo)

```yml
version: '2'
services:
  exporter:
    container_name: mysql-exporter-dev
    image: prom/mysqld-exporter
    environment:
      # 这里指定的是要监控的 MySQL 数据库，这里我们以启动的 test-mysql 容器为示范。
      # 实际应用当中，应该配置为具体的数据库实例。
      - DATA_SOURCE_NAME=root:root@(test-mysql:3306)/

  prometheus:
    container_name: prometheus
    image: prom/prometheus
    ports:
      - "20001:9090"
    # 映射普罗米修斯的配置文件，用于配置 Exporter，这里的文件应该在后面创建好，具体
    # 路径以实际为准。
    volumes:
      - /root/Docker/Volumes/Prometheus/prometheus.yml:/etc/prometheus/prometheus.yml

  grafana:
    container_name: grafana
    image: grafana/grafana
    environment:
      # 配置 Grafana 的默认根 URL。
      - GF_SERVER_ROOT_URL=http://192.168.100.107:20002
      # 配置 Grafana 的默认 admin 密码。
      - GF_SECURITY_ADMIN_PASSWORD=admin
    ports:
      - "20002:3000"
    # 映射 Grafana 的数据文件，方便后面进行更改。
    volumes:
      - /root/Docker/Volumes/Grafana:/var/lib/grafana

  # 本服务只是用于演示，实际使用请注释掉本服务。
  mysql:
    container_name: test-mysql
    image: mysql
    environment:
      - MYSQL_ROOT_PASSWORD=root

# 这里如果需要连接外部 MySQL 就需要处在同一个网络。
networks:
  default:
    external:
      name: mysql-monitor
```
