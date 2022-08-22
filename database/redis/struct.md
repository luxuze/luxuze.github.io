# Redis 数据类型

## String

- set 和 get，做简单的 KV 缓存。
- 缓存功能：String 字符串是最常用的数据类型，不仅仅是 Redis，各个语言都是最基本类型，因此，利用 Redis 作为缓存，配合其它数据库作为存储层，利用 Redis 支持高并发的特点，可以大大加快系统的读写速度、以及降低后端数据库的压力。
- 计数器：许多系统都会使用 Redis 作为系统的实时计数器，可以快速实现计数和查询的功能。而且最终的数据结果可以按照特定的时间落地到数据库或者其它存储介质当中进行永久保存。
- 共享用户 Session：用户重新刷新一次界面，可能需要访问一下数据进行重新登录，或者访问页面缓存 Cookie，但是可以利用 Redis 将用户的 Session 集中管理，在这种模式只需要保证 Redis 的高可用，每次用户 Session 的更新和获取都可以快速完成。大大提高效率。

## Hash

- 类似 Map 的一种结构，这个一般就是可以将结构化的数据，比如一个对象（前提是这个对象没嵌套其他的对象）给缓存在 Redis 里，然后每次读写缓存的时候，可以就操作 Hash 里的某个字段。

## List

- 有序列表。
- 可以通过 List 存储一些列表型的数据结构，类似粉丝列表、文章的评论列表之类的东西。
- 可以通过 lrange 命令，读取某个闭区间内的元素，可以基于 List 实现分页查询，这个是很棒的一个功能，基于 Redis 实现简单的高性能分页，可以做类似微博那种下拉不断分页的东西，性能高，就一页一页走。
- 可以搞个简单的消息队列，从 List 头怼进去，从 List 屁股那里弄出来。

## Set

- Set 是无序集合，会自动去重的那种。
- 直接基于 Set 将系统里需要去重的数据扔进去，自动就给去重了，如果你需要对一些数据进行快速的全局去重，你当然也可以基于 JVM 内存里的 HashSet 进行去重，但是如果你的某个系统部署在多台机器上呢？得基于 Redis 进行全局的 Set 去重。
- 可以基于 Set 玩儿交集、并集、差集的操作，比如交集吧，我们可以把两个人的好友列表整一个交集，看看俩人的共同好友是谁？对吧。

## Sorted Set

- Sorted set 是排序的 Set，去重但可以排序，写进去的时候给一个分数，自动根据分数排序。
- 有序集合的使用场景与集合类似，但是 set 集合不是自动有序的，而 Sorted set 可以利用分数进行成员间的排序，而且是插入时就排序好。所以当你需要一个有序且不重复的集合列表时，就可以选择 Sorted set 数据结构作为选择方案。

- 排行榜：有序集合经典使用场景。例如视频网站需要对用户上传的视频做排行榜，榜单维护可能是多方面：按照时间、按照播放量、按照获得的赞数等。
- 用 Sorted Sets 来做带权重的队列，比如普通消息的 score 为 1，重要消息的 score 为 2，然后工作线程可以选择按 score 的倒序来获取工作任务。让重要的任务优先执行。

### 高级用法

#### **Bitmap**

位图是支持按 bit 位来存储信息，可以用来实现 **布隆过滤器（BloomFilter）**；

#### **HyperLogLog:**

供不精确的去重计数功能，比较适合用来做大规模数据的去重统计，例如统计 UV；

#### **Geospatial:**

可以用来保存地理位置，并作位置距离计算或者根据半径计算位置等。有没有想过用 Redis 来实现附近的人？或者计算最优地图路径？

#### **pub/sub：**

功能是订阅发布功能，可以用作简单的消息队列。

#### **Pipeline：**

可以批量执行一组指令，一次性返回全部结果，可以减少频繁的请求应答。

#### **Lua：**

**Redis** 支持提交 **Lua** 脚本来执行一系列的功能。

我在前电商老东家的时候，秒杀场景经常使用这个东西，讲道理有点香，利用他的原子性。

话说你们想看秒杀的设计么？我记得我面试好像每次都问啊，想看的直接**点赞**后评论秒杀吧。

#### **事务：**

最后一个功能是事务，但 **Redis** 提供的不是严格的事务，**Redis** 只保证串行执行命令，并且能保证全部执行，但是执行命令失败时并不会回滚，而是会继续执行下去。