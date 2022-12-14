package main

import (
	"apk_leaks/common"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"os"
	"runtime"
	"time"
)

// apkleaks工具在apk文件大于60M左右就会出现问题,加入文件时对文件大小做了一个排除
// color带颜色输出时自动加了`\n`
// color支持格式体字符串
// apkleaks.py比较吃内存，就没有走多协程
func main() {
	start := time.Now()
	const system = runtime.GOOS
	// 路径最好不要包含中文
	flag.StringVar(&common.Package_name, "p", "", "The folder where apk file save")
	flag.StringVar(&common.Call_file, "a", "", "Specify the path where the calling file is located")
	flag.StringVar(&common.Save_file, "o", "result.txt", "The path where result save")
	flag.Parse()
	if common.Package_name == "" || common.Call_file == "" {
		flag.Usage()
		os.Exit(1)
	}
	// 1.获取某个目录下的所有文件的缓存,system表示操作系统类型
	common.Files_in_package(common.Package_name, &common.Files_result, system)
	// 2.遍历读取文件名，使用apkleaks进行解析写入
	color.Magenta("开始调用apkleaks.py解析数据，请稍等......")
	for _, filepath := range common.Files_result {
		if system == "windows" {
			_ = common.RunInWindowsWithErr(common.Call_file, common.Save_file, filepath)
		} else if system == "linux" {
			_ = common.RunInLinuxWithErr(common.Call_file, common.Save_file, filepath)
		} else {
			color.Red("not support system！！")
		}
	}

	color.Red("扫描结束，请查看保存的文件")
	t := time.Now().Sub(start)
	fmt.Printf("[*] 扫描结束,耗时: %s\n", t)
}
