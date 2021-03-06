- 包一级声明的名字在**整个包内每个源文件**中都可以访问，不仅仅局限在声明的源文件中

- 包一级的声明没有“先定义后使用”的规则，都是平级的，无论先后都可以相互调用

- 包级别声明的变量将会在**main执行前**完成初始化，局部变量在执行到的时候完成初始化

- 简短变量（:=）声明要求左侧至少要有1个是未声明的新变量

  - 未声明的变量：声明并赋值
  - 已声明的变量：仅赋值

- 如果变量名和外部词法域的变量同名，那么简短声明将会**重新声明一个新的变量**

- go可以返回局部变量的地址——只要还有引用，地址就一直有效

- 当使用命令行参数解析（通过flag.Parse()来获取参数值）时，普通的命令行参数需通过flag.Args()来获取

- 变量生命周期

  - 包级别：跟整个程序运行周期一致
  - 局部变量：从声明开始，直到**不再被引用**为止

- 可赋值=>可比较（==, !=）

- 类型声明

  ```go
  type typename underlyingType
  ```

  可以在相同的底层基础上表达不同的上层概念。不同“概念”间不能直接一起运算/比较

  运算仍然发生在底层类型上，底层类型支持什么运算符就可以用什么运算符

- 底层基础类型相同才允许转型操作。当然，数值等的转换不在其中（如float转int丢弃小数，string转byte[]进行拷贝等）

- 包变量按照依赖关系的初始化

- 可以在每个文件中声明**多个**init函数用于包变量的初始化，执行顺序与他们的声明顺序一致