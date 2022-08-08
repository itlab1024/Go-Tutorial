![](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202207281322806.png)

> 字符串

字符串在每种语言中都是非常重要的，go中的字符串使用`string`类型定义

# 定义

简单定义

```go
package main

import "fmt"

func main() {
	// 字符串定义
	var s string = "I am learning go language"
	fmt.Println(s) // I am learning go language
}
```

# 索引

go中的字符串可以使用索引访问

```go
// 按照索引访问
fmt.Printf("%c\n", s[0]) // I
fmt.Printf("%x\n", s[0]) // 49
```

# 长度

可以通过len函数获取字符串的长度

```go
// 字符串的长度
fmt.Println(len(s)) // 25
fmt.Println(len("哈哈")) // 6
```

**第一行好像没有错，但是第二行为什么是6呢？这是因为go中字符串是以UTF-8编码存储的，一个汉字代表三个字节，所以是6**

# 比较

字符串比较使用`==`,`>`,`<`运算符

```go
// 字符串比较
a := "go"
b := "golang"
fmt.Println(a == b) // false
fmt.Println(a >= b) // false
fmt.Println(a <= b) // true
fmt.Println(a > b) // false
fmt.Println(a < b) // true
fmt.Println(a != b) // true
```

# 遍历字符串

```go
// 遍历字符串
var s1 = "我在学习go语言"
for index, value := range s1 {
  fmt.Println("索引", index, "的值是", value)
}
```

运行结果是：

![image-20220807175830405](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202208071758538.png)

索引值为啥不连续，而且值是数字，这数字是什么东西？

> Go的字符串表示方法

**go中字符串是以UTF-8形式存储的，go中的字符分为两种，一种是是uint8类型，等价于byte类型，代表了ASCII码的一个字符，占用一个字节，另外一种是rune类型，代表一个UTF-8类型，处理中文、日文还有一些希腊文（好像是希腊文，不管了。）使用，rune等价于int32，用于处理大于1个自己小于四个字节的情况。通常情况下一个字符就是一个字节，在 `Go` 中用 `byte` 关键字来表示字节，而 `UTF-8` 编码的字符，可能会存在一个字符用三个字节表示。如果使用 `byte` 类型来存储 `UTF-8` 编码的字符串，就会导致读取单个字节时值没有意义的情况。**

正确输出方式：

需要将字符串转化为rune切片。

```go
// 使用rune输出
runeSlice := []rune(s1)
for index, value := range runeSlice {
  fmt.Printf("索引%d的值是%c,", index, value) // 索引0的值是我,索引1的值是在,索引2的值是学,索引3的值是习,索引4的值是g,索引5的值是o,索引6的值是语,索引7的值是言,
}
```

这就正确输出了！长度也是没有问题的。

当我们单独获取长度呢？可以使用utf8.RuneCountInString()方法。

```go
// 打印字符串长度
fmt.Println()
fmt.Println(utf8.RuneCountInString(s1)) // 8
```

# 字符串读取

字符串读取使用变量[index]的方法，上面已经使用到了。但是还是那个问题，要注意读取的内容是否是你想要的。

# 不可变性

字符串是不可变的，通过索引更改字符串内容是不允许的。

```go
// 字符串是不可变的
//s1[0] = "1" // 编译就不通过。
```

# 构造字符串

字符串是有字节构成，如何使用字节切片构造字符串呢

```go
// 构造字符串
byteSlice := []byte{72, 101, 108, 108, 111}
s3 := string(byteSlice)
fmt.Println(s3) // "Hello"
// 也可以使用unicode
runeSlice2 := []rune{0x0048, 0x0065, 0x006C, 0x006C, 0x006F}
s4 := string(runeSlice2)
fmt.Println(s4) // Hello
```
