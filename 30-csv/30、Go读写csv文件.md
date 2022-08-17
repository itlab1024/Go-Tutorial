# 说明
在Go中使用csv需要导入encoding/csv包。
同时需要配合os包打开csv文件。
# 读取csv
准备一个csv文件read.csv:
![](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202208171556955.png)
使用csv.NewReader读取文件
```go
package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("data/read.csv")
	if err != nil {
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	fmt.Println(records) // 返回的是二维数组切片。[[姓名 职位] [诸葛亮 军师] [刘备 左将军] [曹操 丞相]]
}
```
# 写文件
写csv文件使用csv.NewWriter
```go
// 写文件
file, err = os.Create("data/write.csv")
defer file.Close()
if err != nil {
    return
}
writer := csv.NewWriter(file)
// 准备数据
data := [][]string{
    {"姓名", "职位"},
    {"庞统", "军师"},
}
writer.WriteAll(data)
```
write.csv文件内容如下：
![](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202208171603018.png)