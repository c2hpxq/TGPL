- 4种复合数据结构

  - 数组：同构，定长
  - slice：同构，变长
  - map：变长
  - 结构体：异构，定长

- 对数组取range，可接收的迭代返回值是：下标、元素值

- 字面值初始化

  ```go
  var a [2]int{1, 2}
  ```

  索引初始化

  ```go
  var a [10]int{8:10}
  ```

  未赋值的都会初始化为0值

  长度由字面值确定

  ```go
  var a [...]int{20:10}
  ```

  

- slice的cap为slice起始位置 ～ 数组结束位置的容量大小。切片超出len会扩展slice，超出cap则panic

- 由于动态扩容的精确策略不是标准的一部分，所以无法保证append后的slice和原slice用的是相同底层空间。所以通常将append结果赋给原slice

  ```go
  s = append(s, c)
  ```

- 更一般地，涉及长度、容量、底层数组更新的操作，**都需要显式更新slice**

- map的key必须支持==

- map查找失败将返回0值，或者也可以接受2值返回

  ```go
  a, ok = m["key"]
  ```

  ok指明这个key是否存在，i.e.，区分key不存在返回0值/key存在对应value就是0值

- map类型的0值是nil，可以查、删、len、range，但不能存——这需要底层真的有存储，就算在第一次存入时分配，那也至少需要传入指向map的指针

- map的value、slice中的值都不能取地址，因为底层结构可能重新分配；结构体的成员可以取地址

- 结构体中大写字母开头的成员是导出的

- Go的一个特殊语法：结构体嵌入。结构体成员只写类型，不写成员名，这样就可以缩短通过.来访问的层数

  ```go
  type Point struct {
    X, Y int
  }
  type Circle struct {
    Point
    Radius int
  }
  ```

  然后，就可以只用一个.从Circle变量访问到Point的成员变量

  ```go
  var c Circle
  c.X = 1
  c.Y = 2
  ```

  但初始化不支持嵌入

  ```go
  c = Circle{
    Point: Point{X: 1, Y: 2},
    Radius: 10,
  }
  ```

- 书上写如果嵌入类型名小写开头（如不是Point，而是point），包外就不能访问嵌入的成员。我测试（Go1.13）c.point.X不行，但c.X仍然可以。实际使用时还是编译器测试一下