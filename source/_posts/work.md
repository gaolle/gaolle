##### redis 

* string（字符串）

  ```c
  struct __attribute__ ((__packed__)) sdshdr64 {
  	// 记录当前已使用的字节数（不包括'\0'），获取SDS长度的复杂度为O(1)
      uint64_t len; 
      // 记录当前字节数组总共分配的字节数量
      uint64_t alloc; 
      // 标记当前字节数组的属性，是sdshdr8还是sdshdr16等
      unsigned char flags; 
      // 字节数组，用于保存字符串，包括结尾空白字符'\0'
      char buf[];
  };
  ```

* map（哈希）

* list（列表）

* set（集合）

  * 哈希

    * 开放地址法
      * 线性探测
    * 拉链法

  * inset（整数集合）->只包含整数值元素，且元素数量不多时，会使用整数集合作为底层实现

    ```c
    typedef struct intset {
        // 编码方式
        uint32_t encoding;
        // 记录整数集合的元素数量，即contents数组长度
        uint32_t length;
        // 整数集合的每个元素在数组中按值的大小从小到大排序，且不包含重复项
        int8_t contents[];
    } intset;
    ```

    * 整数集合升级
      - 根据新元素类型，扩展整数集合底层数组的空间大小，并为新元素分配空间
      - 把数组现有的元素都转换成新元素的类型，并将转换后的元素放到正确的位置，且要保持数组的有序性
      - 添加新元素到底层数组

* sortSet（有序集合）

##### go

* RWMutex

  * 读优先、写优先（go）、不指定优先级

    ```go
    type RWMutex struct {
    	w           Mutex  // 多个writer的竞争
    	writerSem   uint32 // 写信号
    	readerSem   uint32 // 读信号
    	readerCount int32  // 负数时，表明存在writer等待请求锁
    	readerWait  int32  // writer等待完成的reader的数量
    }
    
    const rwmutexMaxReaders = 1 << 30
    ```

  * 坑点
    * 声明即用，不可复制
  * 死锁条件
    * 互斥
    * 请求并等待
    * 不可抢夺
    * 循环等待

* Context

  ```go
  type Context interface {
  	Deadline() (deadline time.Time, ok bool)
  	Done() <-chan struct{}
  	Value(key interface{}) interface{}
  }
  ```

  * 场景

    * 上下文消息传递

    * 控制 goroutine 的运行

      ```go
      	cancel, cancelFunc := context.WithCancel(context.Background())
      	go func() {
      		for {
      			select {
      			case <-cancel.Done():
      				return
      			}
      		}
      	}()
      	time.Sleep(time.Second * 2)
      	cancelFunc()
      	time.Sleep(time.Second * 2)
      ```

    * 超时控制的方法调用

    * 可以取消的方法调用

  * 约定

    * 函数第一个参数
    * nil 不能为 Context 类型的参数值
    * 不能持久化
    * 上下文消息传递不应该是内建类型，可能会产生冲突，建议自己定义的类型

* sync.Map

  * 原生map 并发

    * 加锁

      ```go
      type SyncMap struct {
      	mu   sync.RWMutex
      	data map[int]interface{}
      }
      ```

    * 分片

  * 场景

    * 当给定键的条目只被写入一次但读取多次时，如在只增长的缓存中
    * 当多个 goroutine 读取、写入和覆盖不相交键集的条目时

  ```go
  type Map struct {
  	mu Mutex
  	// 安全只读map
  	read atomic.Value
  	// 加锁才能访问->包含所有read字段但未删除的元素以及新增加的元素
  	dirty map[interface{}]*entry
  	// 从read读不命中时，+1  等于dirty长度时，dirty提升为read,重新分配内存给dirty
  	misses int
  }
  ```

  * 注意
    * 双重检测，读写操作时，先访问read，不命中时访问dirty
    * delete时打标记，真正删除时为 missess == dirty长度，直接将dirty提升为read，原read丢弃被回收

* atomic

  cpu提供lock指令，

  * 类型：int32、int64、uint32、uint64 、pointer(add方法不支持)
  * 方法：add、cas、swap、load、store，操作为地址值

* Channel

  循环队列，异步安全，多生产者多消费者

  * 数据传递

    ```
    
    ```

    

  * 信号通知

    ```go
    	go func() {
    		ch := make(chan struct{}, 1)
    		for {
    			select {
    			case <-ch:
    				close(ch)
    				return
    			default:
    				println("first")
    				ch <- struct{}{}
    			}
    		}
    	}()
    ```

  * panic

    * close为nil的chan
    * close已经close的chan
    * send已经close的chan

* gc时机

  * 主动触发：runtime.GC()
  * 被动触发
    * 系统监控：超过两分钟没有gc，强制触发gc
    * 步调算法：控制内存增长的比例





mvcc  版本链+视图       表隐藏的三个字段     事务id，回滚指针，主键   写锁保证顺序

bin log 逻辑日志 redo log 物理日志





逃逸分析（内存分配由栈分配为到堆）    

* 指针
* interface
* 栈内存空间不够
* 闭包



引用类型

* 指针
* 切片
* map
* 函数
* channel
* interface



检验和、序列号、超时重传、流量控制、拥塞避免（慢开始，拥塞避免、快重传、快恢复）