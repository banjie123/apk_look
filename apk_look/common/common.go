package common

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

// 文件结果存储切片
var Files_result []string

// 用于apk文件完成计数
var count = 0

// 文件路径
var (
	Package_name string
	Call_file    string
	Save_file    string
)

func Files_in_package(package_name string, results *[]string, system string) error {
	switch {
	case system == "linux":
		rd, err := ioutil.ReadDir(package_name)
		if err != nil {
			//fmt.Println("read dir fail:", err)
			color.Red("read dir fail:", err)
			return err
		}
		for _, fi := range rd {
			// 不包含文件夹下的文件
			if !fi.IsDir() {
				if fi.Size() >= 60000000 {
					color.Cyan("%s文件大于60M，不加入解析行列", fi.Name())
					continue
				}
				//fullName := package_name  + fi.Name()
				fullName := package_name + "/" + fi.Name()
				//fmt.Println(fullName)
				*results = append(*results, fullName)
			}
		}
		return nil
	case system == "windows":
		rd, err := ioutil.ReadDir(package_name)
		if err != nil {
			//fmt.Println("read dir fail:", err)
			color.Red("read dir fail:", err)
			return err
		}
		for _, fi := range rd {
			if !fi.IsDir() {
				if fi.Size() >= 60000000 {
					color.Cyan("%s大于60M，不加入解析行列", fi.Name())
					continue
				}
				// 不包含文件夹下的文件
				fullName := package_name + "\\" + fi.Name()
				//fullName := package_name + "\" + fi.Name()
				//fmt.Println(fullName)
				*results = append(*results, fullName)
			}
		}
		return nil
	default:
		color.Red("not support system！！")
	}

	return nil
}

func RunInLinuxWithErr(call_file string, save_file string, file_path string) error {
	if strings.Contains(file_path, ".apk") {
		count = count + 1
		cmd_string := fmt.Sprintf(" python3 %s -f %s --args=\"--threads-count 5\"  -o %s", call_file, file_path, save_file)
		command := exec.Command("/bin/bash", "-c", cmd_string)
		err := command.Run()
		if err != nil {
			color.Red("please check your input paramters")
			os.Exit(1)
		}
		write_string := fmt.Sprintf("echo --------------%s-------------- >> result.txt", file_path)
		command1 := exec.Command("/bin/bash", "-c", write_string)
		command1.Run()
		//fmt.Println("--->已完成第\t",count,"\t个目标<---")
		i := len(strings.Split(file_path, "\\"))
		file_path = strings.Split(file_path, "\\")[i-1]
		color.Magenta("--->已完成 %v 目标<---", file_path)
	}
	return nil

}

func RunInWindowsWithErr(call_file string, save_file string, file_path string) error {
	if strings.Contains(file_path, ".apk") {
		count = count + 1
		cmd_string := fmt.Sprintf(" python3 %s -f %s --args=\"--threads-count 5\"  -o %s", call_file, file_path, save_file)
		command, _ := neibuCommand("cmd", "/c "+cmd_string)
		//command := exec.Command("cmd", "/c"+cmd_string)
		err := command.Run()
		if err != nil {
			color.Red("please check your input paramters")
			os.Exit(1)
		}
		write_string := fmt.Sprintf("echo --------------%s-------------- >> result.txt", file_path)
		command1, err := neibuCommand("cmd", "/c"+write_string)
		//command1 := exec.Command("cmd", "/c"+write_string)
		command1.Run()
		//fmt.Println("--->已完成第\t",count,"\t个目标<---")
		i := len(strings.Split(file_path, "\\"))
		file_path = strings.Split(file_path, "\\")[i-1]
		color.Magenta("--->已完成 %v 目标<---", file_path)
	}
	return nil

}

func neibuCommand(name, args string) (*exec.Cmd, error) {
	if filepath.Base(name) == name {
		lp, err := exec.LookPath(name)
		if err != nil {
			return nil, err
		}
		name = lp
	}
	return &exec.Cmd{
		Path:        name,
		SysProcAttr: &syscall.SysProcAttr{CmdLine: name + " " + args},
	}, nil
}
