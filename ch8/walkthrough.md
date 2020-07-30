- 没有**直接**的方法让一个goroutine去打断另一个goroutine

- go语言奇怪的时间format layout

  ```go
  Mon Jan 2 15:04:05 -0700 MST 2006
  ```

  1月2日下午3点4分5秒06年，UTC-0700

- 对已经close的channel接收，仍可以接收到之前成功发送到数据。channel中没有数据了，则返回一个0值

- channel可以不带缓冲（发送一次，值没被读取之前，阻塞；读取同理）

  ```go
  ch = make(chan int)
  ch = make(chan int, 0)
  ```

  带缓冲

  ```go
  ch = make(chan int, 3)
  ```

- 无缓冲channel读写：go保证接收数据发生在唤醒发送者之前

  - 用于goroutine等待

    ```go
    done := make(chan struct{})
    
    go func() {
      //...
      done <- struct{}{}
    }()
    
    <-done
    ```

    这里使用struct{}，空的结构体，因为我们不关心发送的具体值，只关心发送的时刻/同步作用。

- channel close了并且数据读取完后，会返回0值。这时候使用其二值接收形式，第一个变量接收值，第二个接收是否close

  - 简略写法，可以直接range channel进行循环，会在无数据后退出循环

- 没有引用后，channel会被gc（更具体地，因为channel是引用类型，应该说底层的结构没有channel指向了）

- 重复关闭channel会panic

- 单向channel

  - 只读：

    ```go
    <-chan int
    ```

  - 只写

    ```go
    chan<- int
    ```

  可以隐式从双向转为单向，但单向不可再转为双向

- select的default分支，在其他channel都无法**立即**获取的情况下触发。用于非阻塞select
- buffered channel可以当信号量用
- 利用close channel读取返回nil，关闭一个从不写入的channel可以让**一群**goroutine接收到这个消息，从而达到“群发通知goroutine的效果”
  - 光是指定一个用于消息通知的channel是不行的，因为一条消息只能被读到一次，加之一般很难准确知道目前有几个goroutine在等这个channel的消息。
  - 这个close channel的bypass相对最合理，也不需要加锁进出之类来知道具体的goroutine数目