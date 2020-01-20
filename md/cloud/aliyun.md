# aliyun

* 启动Logtail容器，替换启动模板中的3个参数：your_region_name，your_aliyun_user_id，your_machine_group_user_defined_id

* [阿里云链接](https://www.alibabacloud.com/help/zh/doc-detail/66659.htm)

```sh
docker run -d -v /:/logtail_host:ro -v /var/run/docker.sock:/var/run/docker.sock --env ALIYUN_LOGTAIL_CONFIG=/etc/ilogtail/conf/${your_region_name}/ilogtail_config.json --env ALIYUN_LOGTAIL_USER_ID=${your_aliyun_user_id} --env ALIYUN_LOGTAIL_USER_DEFINED_ID=${your_machine_group_user_defined_id} registry.cn-hangzhou.aliyuncs.com/log-service/logtail

docker run -d -v /:/logtail_host:ro -v /var/run/docker.sock:/var/run/docker.sock --env ALIYUN_LOGTAIL_CONFIG=/etc/ilogtail/conf/cn-hangzhou-internet/ilogtail_config.json --env ALIYUN_LOGTAIL_USER_ID=1652220740479945 --env ALIYUN_LOGTAIL_USER_DEFINED_ID=aliyun-lxz registry.cn-hangzhou.aliyuncs.com/log-service/logtail
```
