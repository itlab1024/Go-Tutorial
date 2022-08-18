package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	// ls命令
	cmd := exec.Command("ls")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return
	}
	// 多系统
	goos := runtime.GOOS
	goarch := runtime.GOARCH
	fmt.Println("当前电脑的系统是", goos, "内核是", goarch) // 当前电脑的系统是 darwin 内核是 amd64
	switch goos {
	case "darwin":
	// todo 执行ls指令
	case "windows":
		// todo 执行dir指令
	}
}
