- 接收多参数的函数可以直接接收多返回值的函数的结果

  ```go
  func main() {
  	fmt.Printf("%d", add(f()))
  }
  
  func f() (int, int) {
  	return 1, 1
  }
  
  func add(x, y int) int {
  	return x+y
  }
  ```

- 所有返回值都有名称时，可以直接return。未在函数体中赋值的返回值会默认初始化为0值

- go的错误处理：
  - panic应仅在bug时出现
  - 预料中的错误处理，以error为返回值（之一）返回给调用者

- log包中的函数会给所有没有换行符的串加换行
- go中函数可以作为值，有类型。
  - 0值为nil
  - 函数类型（在这里，参数+返回值是函数类型）跟变量类型一致，就可以赋值，不关函数体
  - 不可比较，因此无法作为map的key
  - 函数作为值传入，那么就可以按行为抽象函数（map、reduce的那一套）

- "%*s"格式化输出：在串前面填充空格，个数由用户传入

- 匿名函数可以引用包围它的lexical environment中的变量

- slice append slice

  ```go
  a = append(a, b...)
  
  ```

- 可变参数：最后一个变量声明加...。相关变量可以视为该类型的slice

- defer

  - 执行顺序与声明顺序相反
  - 执行时机在return之后

- panic后，defer的函数仍会执行（加入的倒序），并且在堆栈释放之前（也因此可以在这个时机打印堆栈信息）

- panic后仍有机会恢复执行

  - 在defer函数中可以通过recover获取panic信息，可以据此来决定要恢复哪些panic
  - 如果没有panic，那么recover返回的值就是nil
  - 不要无脑恢复。对于不可恢复的错误，在defer的函数中调用panic传递错误信息