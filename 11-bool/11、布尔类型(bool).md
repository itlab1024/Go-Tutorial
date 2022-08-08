![](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202207281322806.png)

# 定义

布尔类型通过`bool`定义，只有两个取值，true和false。它的零值是false。

```go
// 定义布尔类型
var b bool
fmt.Print("布尔类型的零值是", b) // false
```

# 布尔运算符

布尔运算符主要是用于判断得出bool结果的运算符，比如`==`,`!=`,`>=`,`<=`,`>`,`<`。

```go
// 布尔运算符
i := 1
j := 2
fmt.Println(i == j) // false
fmt.Println(i != j) //true
fmt.Println(i >= j) // false
fmt.Println(i <= j) // true
fmt.Println(i < j) // true
fmt.Println(i > j) //false
```

# 字符串bool相互转化

字符串转化为bool规则：

以下字符串会被成功转化为true，error返回的是nil

```tex
"1", "t", "T", "true", "TRUE", "True":
```

以下字符串会被成功转化为false，error返回的是nil

```tex
"0", "f", "F", "false", "FALSE", "False"
```

除了以上情况，会返回false，但是error会有错误信息。返回的是syntaxError错误哦。

两者相互转化常使用**strconv**包的内容。

```go
// 字符串和bool相互转化
// ParseBool返回两个值，第一个是具体的值，第二个是error。
b1, _ := strconv.ParseBool("true")
fmt.Println(b1) // true
// 如果字符串不是true或者false呢
b2, _ := strconv.ParseBool("能正常转化成功吗")
fmt.Println(b2) // false, 这是因为我忽略错误。也可以不忽略错误
//处理错误
b3, err := strconv.ParseBool("能正常转化成功吗")
if err != nil {
  fmt.Println("转化有错误")
} else {
  fmt.Println(b3)
}
b4, _ := strconv.ParseBool("True")
fmt.Println(b4)

// 也可以将bool转化为字符串
fmt.Println(strconv.FormatBool(false)) // false
fmt.Println(strconv.FormatBool(true))  // true
```

