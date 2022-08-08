![](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202207281322806.png)

# 函数声明

函数声明使用`func`关键字，结构如下

```go
func 函数名称(参数列表) (返回类型列表) {
    // 函数体
}
```

# 函数调用

定义一个`sayHello`函数，并在main函数中调用

```go
package main

import "fmt"

func main() {
	sayHello()
}
// 定义了一个无参数，返回空类型的函数
func sayHello() {
	fmt.Println("Hello go")
}
```

# 函数的参数

标准函数参数就是在函数名后面依次排列即可，比如(参数名  参数类型, 参数名  参数类型)

```go
// 函数的参数(标准)
func param1(i int, j int) {
	fmt.Println(i, j)
}
```

i和j就是两个不同类型的参数。

# 同类型参数简写

比如上面的函数可以修改为如下，

```go
// 函数的参数(标准)，i并没有单独设置类型。而是和j公用一个类型。
func param1(i, j int) {
	fmt.Println(i, j)
}
```

# 函数返回值

定义一个返回一个int类型的函数

```go
// 函数返回值
func add(x, y int) int {
   return x+y
}
```

# 函数多返回值

比如传递两个参数，返回两个值的和以及差

```go
// 多返回值
func addSub(x, y int) (int, int) {
   return x+y, x-y
}
```

# 命名返回值

返回值类型处可以设置名称，函数体中只需要给这个名称的返回值赋值，最后return即可。

```go
// 命名返回值
func namedReturn(x, y int) (a int, b int) {
	a =  x + y
	b = x - y
	return
}
```

# 参数值传递

go中参数的传递默认是值传递。

```go
package main

import "fmt"

func main() {
	// 测试函数参数传递是否是值传递
	i := 1
	prepareChange(i) // 设置i=100
	fmt.Println(i)   // 结果还是1
}

// 函数的值传递
func prepareChange(i int) {
	i = 100
}
```

那如何做到值改变呢?这就需要传递指针（也就是引用传递）

```go
func main() {
	sayHello()
	// 测试函数参数传递是否是值传递
	i := 1
	prepareChange(i) // 设置i=100
	fmt.Println(i)   // 结果还是1

	j := 1
	prepareChange2(&j)
	fmt.Println(j) // j=100
}
// 函数的引用传递
func prepareChange2(i *int) {
	*i = 100
}
```

可以看到值变化了，这是因为prepareChange修改的是当时定义的同一内地址的数据。

针对值传递的代码，我们增加打印下执行prepareChange前后变量i的内存地址

```go
func main() {
	sayHello()
	// 测试函数参数传递是否是值传递
	i := 1
	fmt.Println(&i)// 0xc000069f00
	prepareChange(i) // 设置i=100
	fmt.Println(i)   // 结果还是1
}
// 函数的值传递
func prepareChange(i int) {
	fmt.Println(&i)// 0xc000069ee0
	i = 100
}
```

结果显示两处的内存地址是不同的。

> 结论 值传递go会拷贝一份新的数据，想用同一份数据则传递指针。

# 空标志符

当返回结果不想使用的时候，则可以使用`_`接受

```go
// 空标识符,比如下面是正常接收
//r1, r2 := addSub(1, 2)
// 如果我不关心r2的结果则可使用空标识符
r1, _ := addSub(1, 2)
fmt.Println(r1)
```

# 匿名函数

当实际开发中，我们不关心函数的名字，只关心函数的逻辑。则可以使用匿名函数

```go
package main

func main() {
	// 匿名函数
	func1 := func() {
		fmt.Println("匿名函数体")
	}
	func1()
}
```

# 匿名函数立即调用

上面的函数是通过func1()来调用的，也可以直接定义后立即调用。

```go
package main

func main() {
	// 匿名函数
	func() {
		fmt.Println("匿名函数体")
  }()
}

```

无需给函数定义变量接收，直接调用。

# 函数作为返回值

```go

package main

func main() {
	// 调用调用返回类型是函数的函数
	f := returnFunc()
	f(1)
}
func returnFunc () func(i int) int {
	return func(i int) int {
		return i + 1;
	}
}
```

# 函数作为参数

函数作为参数，对比实际业务，比如回调

```go
package main

func main() {
	// 函数作为参数
	biz(i, callback)
}
func biz(i int, c func(i int)) {
	c(i)
}

func callback(i int) {
	fmt.Println("回调成功", i)
}
```

# 闭包

闭包（Closure）是匿名函数的一个特例。当一个匿名函数所访问的变量定义在函数体的外部时，就称这样的匿名函数为闭包。

闭包是由函数及其相关引用环境组合而成的实体(即：闭包=函数+引用环境)。

“官方”的解释是：所谓“闭包”，指的是一个拥有许多变量和绑定了这些变量的环境的表达式（通常是一个函数），因而这些变量也是该表达式的一部分。

```go
package main

func main() {
	// 闭包
  f2 := cumulative()
  i2 := f2(1)
  i3 := f2(1)
  i4 := f2(1)
  fmt.Println(i2, i3, i4) // 1 2 3
}
// 闭包函数，累加功能
func cumulative() func(i int) int {
	var result int
	return func(i int) int {
		result += i
		return result
	}
}
```

f2是一个函数，是cumulative函数返回的函数（匿名函数），该匿名函数引用了result，请注意result在该函数的外部。只要该函数f2生命周期没有结束，result的值不会丢失。f2就形成了一个闭包。

那什么时候声明周期结束呢，只需要重新定义引用。

上面的实示例做下修改

```go
package main

func main() {
	// 闭包
  f2 := cumulative()
  i2 := f2(1)
  i3 := f2(1)
  i4 := f2(1)
  fmt.Println(i2, i3, i4) // 1 2 3
  
  // 此处重新定义了f2
  f2 = cumulative()
	i5 := f2(1)
	i6 := f2(1)
	i7 := f2(1)
	fmt.Println(i5, i6, i7) // 1 2 3 ，可以看到结果并没有继续增加。而是从头开始计算
}
// 闭包函数，累加功能
func cumulative() func(i int) int {
	var result int
	return func(i int) int {
		result += i
		return result
	}
}
```

