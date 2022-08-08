![](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202207281322806.png)

# 定义Map

go中map定义格式如下

```go
var 变量名 map[key类型]值类型
```

比如定义一个key是int类型，value是string类型的map

```go
var intMap map[int]string
```

# 初始化

map必须初始化才能使用，可以使用make函数初始化，也可以使用kv数据进行初始化

```go
// 定义map
var intMap map[int]string
fmt.Println(intMap == nil) // true
// 1. 使用make函数初始化
intMap = make(map[int]string)
fmt.Println(intMap == nil) // false

// 2. 使用kv值初始化
var m1 = map[int]string{} // {}代表空
fmt.Println(m1)           // map[]
// 也可以指定具体的kv值
var m2 = map[int]string{
  0: "a", // 最后这个逗号不能省略
}
fmt.Println(m2) //map[0:a]
```



# 基本操作

map支持插入、查询、循环等基本操作

## 插入元素

使用变量p[key] = value的方式

```go
//插入元素
intMap[0] = "a"
fmt.Println(intMap) // map[0:a]
```

## 查询元素

查询元素时需要注意如果没有对应的key会返回值类型的零值。

```go
// 查询元素
s := intMap[0]
s1 := intMap[1] // 当前map中没有key=1
fmt.Println(s)  // a
fmt.Println(s1) // s1 == "" 因为没有找到key=1的值，所以返回零值，但是如何判断是没有这个key还是这个key的值就是空字符串呢？
```

## 判断key是否存在

获取元素的方法有两个返回值，会告诉使用者该key是否存在。

```go
//检查key是否存在
//获取元素的方法有两个返回值，会告诉使用者该key是否存在。
s2, exists := intMap[2]
if exists {
  fmt.Println(s2)
} else {
  fmt.Println("key=2 不存在")// key=2 不存在
}
```

## 通过key删除

使用delete方法

```go
// 删除使用delete方法
delete(intMap, 0)
delete(intMap, 1)
```

## 循环map

使用range关键字循环

```go
// 使用range遍历
for index, value := range intMap {
  fmt.Println("index=", index, "value=", value)
}
```

