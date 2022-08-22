# 常见八股

## map 的扩容原理

```go


// A header for a Go map.
type hmap struct {
	// Note: the format of the hmap is also encoded in cmd/compile/internal/reflectdata/reflect.go.
	// Make sure this stays in sync with the compiler's definition.
	count     int // # live cells == size of map.  Must be first (used by len() builtin)
	flags     uint8
	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
	hash0     uint32 // hash seed

	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

	extra *mapextra // optional fields
}

```

## slice 的扩容原理

```go

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

```

1. 切片是对底层数组的一个抽象，描述了它的一个片段。
1. 切片实际上是一个结构体，它有三个字段：长度，容量，底层数据的地址。
1. 多个切片可能共享同一个底层数组，这种情况下，对其中一个切片或者底层数组的更改，会影响到其他切片。
1. append 函数会在切片容量不够的情况下，调用 growslice 函数获取所需要的内存，这称为扩容，扩容会改变元素原来的位置。
1. 扩容策略并不是简单的扩为原切片容量的 2 倍或 1.25 倍，还有内存对齐的操作。扩容后的容量 >= 原容量的 2 倍或 1.25 倍。
1. 当直接用切片作为函数参数时，可以改变切片的元素，不能改变切片本身；想要改变切片本身，可以将改变后的切片返回，函数调用者接收改变后的切片或者将切片指针作为函数参数

## chan 的原理

```go

type hchan struct {
	qcount   uint           // total data in the queue
	dataqsiz uint           // size of the circular queue
	buf      unsafe.Pointer // points to an array of dataqsiz elements
	elemsize uint16
	closed   uint32
	elemtype *_type // element type
	sendx    uint   // send index
	recvx    uint   // receive index
	recvq    waitq  // list of recv waiters
	sendq    waitq  // list of send waiters

	// lock protects all fields in hchan, as well as several
	// fields in sudogs blocked on this channel.
	//
	// Do not change another G's status while holding this lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking.
	lock mutex
}
```

## 并发控制技术有哪些

1. waitgroup
1. chan select
1. context
1. mutex

## gmp 调度 [link](https://learnku.com/articles/41728)

![](https://cdn.learnku.com/uploads/images/202003/11/58489/Ugu3C2WSpM.jpeg!large)

1. 全局队列（Global Queue）：存放等待运行的 G。
1. P 的本地队列：同全局队列类似，存放的也是等待运行的 G，存的数量有限，不超过 256 个。新建 G’时，G’优先加入到 P 的本地队列，如果队列满了，则会把本地队列中一半的 G 移动到全局队列。
1. P 列表：所有的 P 都在程序启动时创建，并保存在数组中，最多有 GOMAXPROCS(可配置) 个。
1. M：线程想运行任务就得获取 P，从 P 的本地队列获取 G，P 队列为空时，M 也会尝试从全局队列拿一批 G 放到 P 的本地队列，或从其他 P 的本地队列偷一半放到自己 P 的本地队列。M 运行 G，G 执行之后，M 会从 P 获取下一个 G，不断重复下去。

## 内存分配和垃圾回收 [likn](https://zhuanlan.zhihu.com/p/264789260)

GC 负责回收堆内存

主流的垃圾回收算法有两大类，分别是追踪式垃圾回收算法和引用计数法（ Reference counting ）。而 Go 语言现在用的三色标记法就属于追踪式垃圾回收算法的一种。

追踪式算法的核心思想是判断一个对象是否可达，一旦这个对象不可达就可以在垃圾回收的控制循环里被 GC 回收了。那么我们怎么判断一个对象是否可达呢？很简单，第一步找出所有的全局变量和当前函数栈里的变量，标记为可达。第二步，从已经标记的数据开始，进一步标记它们可访问的变量，以此类推。

![](https://pic2.zhimg.com/80/v2-9488a6193e540afc7a4bcd591afa6db5_1440w.jpg)
这个算法最大的问题是 GC 执行期间需要把整个程序完全暂停，不能异步地进行垃圾回收，对实时性要求高的系统来说，这种需要长时间挂起的标记清扫法是不可接受的。所以就需要一个算法来解决 GC 运行时程序长时间挂起的问题。

### 三色标记

三色标记算法将程序中的对象分成白色、黑色和灰色三类：

白色对象 — 潜在的垃圾，其内存可能会被垃圾收集器回收；
黑色对象 — 活跃的对象，包括不存在任何引用外部指针的对象以及从根对象可达的对象，垃圾回收器不会扫描这些对象的子对象；
灰色对象 — 活跃的对象，因为存在指向白色对象的外部指针，垃圾收集器会扫描这些对象的子对象；

写屏障主要做一件事情，修改原先的写逻辑，然后在对象新增的同时给它着色，并且着色为灰色。
一次完整的垃圾回收会分为四个阶段，分别是标记准备、标记、结束标记以及清理。在标记准备和标记结束阶段会需要 STW，标记阶段会减少程序的性能，而清理阶段是不会对程序有影响的。

gc 时机

1. gcTriggerHeap：针对于对内存的分配达到控制计算的触发堆的大小。
1. gcTriggerTime ：针对在一定时间内没有触发就会触发 GC。
1. gcTriggerCycle：针对没有达到指定的轮询次数触发 GC
