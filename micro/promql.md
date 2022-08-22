# PromQL

## 什么是 PromQL？

PromQL（Prometheus Query Language）是 Prometheus 内置的数据查询语言，它能实现对事件序列数据的查询、聚合、逻辑运算等。它并且被广泛应用在 Prometheus 的日常应用当中，包括对数据查询、可视化、告警处理当中。

简单地说，PromQL 广泛存在于以 Prometheus 为核心的监控体系中。所以需要用到数据筛选的地方，就会用到 PromQL。例如：监控指标的设置、报警指标的设置等等。

## PromQL 基础用法

### increase()

increase(v range-vector) 函数获取区间向量中的第一个和最后一个样本并返回其增长量，它会在单调性发生变化时(如由于采样目标重启引起的计数器复位)自动中断。

由于这个值被外推到指定的整个时间范围，所以即使样本值都是整数，你仍然可能会得到一个非整数值，如果除以一定时间就可以获取该时间内的平均增长率。

例如，以下表达式返回区间向量中每个时间序列过去 5 分钟内 HTTP 请求数的增长数：

increase(http_requests_total{job="apiserver"}[5m])
该函数配合 counter 数据类型使用，它的返回值类型只能是计数器类型。

### rate()

rate(v range-vector) 函数可以直接计算区间向量 v 在时间窗口内平均每秒增长速率，它会在单调性发生变化时(如由于采样目标重启引起的计数器复位)自动中断。

该函数配合 counter 数据类型使用，它的返回值类型只能用计数器，在长期趋势分析或者告警中推荐使用这个函数。该函数的返回结果不带有度量指标，只有标签列表。

例如，以下表达式返回区间向量中每个时间序列过去 5 分钟内 HTTP 请求数的每秒增长率：

```prom
rate(http_requests_total[5m])

结果：
{code="200",handler="label_values",instance="120.77.65.193:9090",job="prometheus",method="get"} 0
{code="200",handler="query_range",instance="120.77.65.193:9090",job="prometheus",method="get"} 0
{code="200",handler="prometheus",instance="120.77.65.193:9090",job="prometheus",method="get"} 0.2
```

### irate()

irate(v range-vector) 函数用于计算区间向量的增长率，但是其反应出的是瞬时增长率。

irate 函数是通过区间向量中最后两个两本数据来计算区间向量的增长速率，它会在单调性发生变化时(如由于采样目标重启引起的计数器复位)自动中断。

这种方式可以避免在时间窗口范围内的“长尾问题”，并且体现出更好的灵敏度，通过 irate 函数绘制的图标能够更好的反应样本数据的瞬时变化状态。

例如，以下表达式返回区间向量中每个时间序列过去 5 分钟内最后两个样本数据的 HTTP 请求数的增长率：

irate(http_requests_total{job="api-server"}[5m])
irate 只能用于绘制快速变化的计数器，在长期趋势分析或者告警中更推荐使用 rate 函数。因为使用 irate 函数时，速率的简短变化会重置 FOR 语句，形成的图形有很多波峰，难以阅读。

[info] 注意

当将 irate() 函数与聚合运算符（例如 sum()）或随时间聚合的函数（任何以 \_over_time 结尾的函数）一起使用时，必须先执行 irate 函数，然后再进行聚合操作，

否则当采样目标重新启动时 irate() 无法检测到计数器是否被重置。
