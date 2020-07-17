- Go数据类型

  - 基础：数字、字符串、bool
  - 复合：数组、结构体
  - 引用：指针、slice、字典、函数、channel
  - 接口

- Go中定义了一个运算符&^（AND NOT）

  ```go
  z = x&^ y
  ```

  y中为1的位，z中为1

  y中为0的位，z中为x中相应位

  对应的集合操作 x-y

- 浮点数转为整数将会向0截断

- 字符串是不可变的字节序列，这样，“复制”字符串、字符串切片操作开销都很低，因为可以共享底层结构

- len将会返回字符串中**字节的数目**，而不是其中的rune数目；相应的，s[i]将会返回第i个字节

- 反引号包含的内容是字符串字面值。除了去除所有'\r'之外，会原样保留所有反引号括住的内容。因此，在

  - HTML模版
  - JSON值
  - 命令行信息中

  使用很方便

- rune ~ int32

- UTF8为变长编码。

- 有时切换为统一长度的[]rune会更方便（如支持random index），Go支持从string直接转为[]rune

  ```go
  r := []rune(s)
  ```

- 可以直接从整数构造一个string

  ```go
  fmt.Println(string(65)) //"A"
  ```

- string转[]byte，将会分配一个新的byte数组来保存string的拷贝，然后[]byte引用这个拷贝。对于新的[]byte的修改不会影响原string

- 数字转字符串：strconv.Itoa，fmt.Sprintf，strconv.FormatInt

- 字符串转数字：strconv.Atoi，strconv.ParseInt

- iota用于常量声明，每过一个字段自动加1