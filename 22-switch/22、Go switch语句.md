# 什么是switch语句?
如果你不是初级入门人员，switch你应该很熟悉，比如在java中就有switch，他是条件判断的简化方式，使用switch的判断语句，基本都可使用IF ELSE来实现。
# 使用
Go中的switch语句相比Java的更加灵活（JDK17貌似也可以很灵活了）。
# 简单实用
使用switch对一个加单的值进行判断，打印一些语句。
```go
package main

import "fmt"

func main() {
	// 简单使用
	i := 1
	switch i {
	case 1:
		fmt.Println("结果是1") // 结果是1
	default:
		fmt.Println("结果不是1")
	}
}
```
# 一case多条件
一个case可以有多个条件，比如判断是1或者2
```go
// 简单实用二，多条件
switch i {
case 1, 2:
    fmt.Println("结果是1或者2") // 结果是1或者2
default:
    fmt.Println("结果不是1也不是2")
}
```
# break关键字
go中的case后默认有break，不用显示地声明break关键字，很友好（java中就不行）。当然如果你手动写生break页没有问题。
```go
// break
switch i {
case 1, 2:
    fmt.Println("结果是1或者2") // 结果是1或者2
    break
default:
    fmt.Println("结果不是1也不是2")
    break
}
```
# fallthrough
使用fallthrough，那么当前case执行完毕后会执行下一个case，不会break；
```go
// fallthrough
switch i {
case 1, 2:
    fmt.Println("结果是1或者2") // 结果是1或者2
    fallthrough
default:
    fmt.Println("结果不是1也不是2") // 结果不是1也不是2
}
```
上面两处都会打印。

# 变量初始化
switch后可以定义变量并初始化，这个变量在switch语句块中都有小。这很有用。
```go
//变量初始化
switch v := 1; v {
case 1:
    fmt.Println("值等于1")
default:
}
```
# 条件case语句
上面的switch都是使用switch i {}这样的格式。也可以使用switch {}这样的方式来实现case条件表示语句，这就灵活了很多。
```go
// 条件case语句
switch {
case i > 0:
    fmt.Println("i > 0") // i > 0
case i <= 0:
    fmt.Println("i <= 0")
}
```