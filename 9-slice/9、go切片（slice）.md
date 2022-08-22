![](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202207281322806.png)
# 什么是切片？
切片是数组的引用，对切片的操作实际是操作切片指针指向的数组，数组使用比较局限，比如不能动态扩容，而切片(slice)就支持，并且提供了很多好用的API。
# 切片的定义

切片的定义使用如下方式:

var 变量名 []类型

```go
// 定义切片
var slice1 []int
// slice1[0] = 1 // 编译通过，但是没有初始化所以执行报错，触发panic
fmt.Println("切片的零值是nil?", slice1 == nil) //切片的零值是nil? true

// 这根数组定义有些类似，千万不要搞混
// 比如如下是数组定义
// var array1 [1]int
```

# 切片的初始化

使用切片必须初始化，否则运行出错。切片的初始化有多方式

第一种：使用[]int{}初始化

```go
// 切片的初始化
// 第一种
slice1 = []int{1, 2, 3}
fmt.Println("切片slice1的长度是", len(slice1), "容量是", cap(slice1)) //切片slice1的长度是 3 容量是 3
```

这里提到了切片的长度和容量，数组的长度和容量是相等的，切片确不一定，之后单独说。

第二种：使用new关键字初始化

```go
// 第二种使用new, new通常用来创建结构体指针，所以以下代码虽然没有问题，但是却不常用，常使用make函数来创建切片
var s1 = new([]int)
*s1 = append(*s1, 1, 2, 3)
fmt.Println(*s1) // [1 2 3]
```

new函数返回的是指针。

第三种：使用make函数创建切片,格式如下

```go
make([]类型, length, capacity)
```

代码示例如下：

```go
// 使用make函数初始化切片
// 指定类型和长度
var s2 = make([]int, 3)
fmt.Println("s2的长度是", len(s2), "容量是", cap(s2)) //s2的长度是 3 容量是 3

// 指定类型，长度=0，容量=3
s3 := make([]int, 0, 3)
fmt.Println("s3的长度是", len(s3), "容量是", cap(s3), "结果是：") //s2的长度是 0 容量是 3

```

第四种：通过重新切片初始化切片

```go
//第四种：通过重新切片初始化切片
s4 := make([]int, 2, 8)
fmt.Println("s4的长度是", len(s4), "容量是", cap(s4)) //s4的长度是 2 容量是 8
// s4已经是切片了，还可以对其进行再次切片
//s5 := s4[2:9] // 9是不允许的，因为s4的cap=8，也就是说再次切片的最大容量不能超过8
s5 := s4[2:8]
fmt.Println("s5的长度是", len(s5), "容量是", cap(s5)) //s5的长度是 6 容量是 6
```

切片再次切片的用法与通过数组初始化切片类似都是使用[strartIndex:endIndex]。

第五种：使用数组初始化切片

```go
// 第五种使用数组初始化切片
var array [3]int = [3]int{1, 2, 3}
// 通过array数组初始化切片
var sl1 = array[1:2]
for _, v := range sl1 {
  fmt.Println(v) // 结果是3，数组索引从1，到2（不包含2，左闭右开）结果就是2
}
```

针对这种[strartIndex:endIndex]切片方式，两个index是可以省略的，strartIndex的默认值就是0，endIndex默认值就是最大索引，如果省略就是使用默认值。

```go
//针对这种[strartIndex:endIndex]切片方式，两个index是可以省略的，strartIndex的默认值就是0，endIndex默认值就是最大索引，如果省略就是使用默认值。
sl2 := array[:]
sl3 := array[0:]
sl4 := array[:2]
fmt.Println(sl2, sl3, sl4) //[1 2 3] [1 2 3] [1 2]
```



# 长度(len)和容量(cap)的区别

len是实际使用的内容，cap是预分配的内存，也就是切片可以重新切片的最大长度，容量是可能大于长度的，配一张网络图。

![61c8747cde259895194a2bf806035bed](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202208071229182.jpeg)

# 追加元素

使用append方法进行元素的追加，其实上面已经使用过，需要注意的是append方法返回的是新的切片，原始切片是不会变化的。请看面下的结果。

```go
// 追加元素
var ap = []int{1, 2, 3}
after := append(ap, 4)
fmt.Println(ap)    //[1 2 3]
fmt.Println(after) // [1 2 3 4]
```

append后ap并没有变化，返回的新切片是append后的结果。

# 循环切片

使用range关键字

```go
// 循环
for index, value := range after {
  fmt.Print(index, "=", value, ",") // 0=1,1=2,2=3,3=4,
}
fmt.Println()
//一个接收默认是index
for index := range after {
  fmt.Print(index, ",") // 0,1,2,3,
}
fmt.Println()
// 可以使用_忽略某个值
for _, value := range after {
  fmt.Print(value, ",") //1,2,3,4,
}
```

# 拷贝函数

拷贝按照长度最短的进行拷贝。

```go
//拷贝
cp1 := []int{1, 2, 3, 4, 5}
cp2 := []int{5, 4, 3}
copy(cp2, cp1)   // 只会复制cp1的前3个元素到cp2中,cp1结果还是
fmt.Println(cp1) //[1 2 3 4 5]
fmt.Println(cp2) //[1 2 3]
fmt.Println("------")

copy(cp1, cp2)   // 只会复制cp2的3个元素到cp1的前3个位置
fmt.Println(cp1) //[1 2 3 4 5]
fmt.Println(cp2) //[1 2 3]
```

