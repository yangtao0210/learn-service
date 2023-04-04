package main

import (
	"fmt"
	"os/exec"
	"strings"
)

/* @命令执行检查
1.使用exec.Command,exec.CommandContext、syscall.StartProcess、os.StartProcess等函数时，第一个参数（path）直接取外部输入值时，应使用白名单限定可执行的命令范围，不允许传入bash、cmd、sh等命令；
2.使用exec.Command、exec.CommandContext等函数时，通过bash、cmd、sh等创建shell，-c后的参数（arg）拼接外部输入，应过滤\n $ & ; | ‘ “ ( ) `等潜在恶意字符；
*/

// checkIllegal 检查命令的合法性
func checkIllegal(cmdName string) bool {
	if strings.Contains(cmdName, "&") || strings.Contains(cmdName, "|") || strings.Contains(cmdName, ";") ||
		strings.Contains(cmdName, "$") || strings.Contains(cmdName, "'") || strings.Contains(cmdName, "`") ||
		strings.Contains(cmdName, "(") || strings.Contains(cmdName, ")") || strings.Contains(cmdName, "\"") {
		return false
	}
	return true
}

func CmdExecute() {
	userInputVal := "echo hello"
	cmdName := "ping " + userInputVal
	if checkIllegal(cmdName) {
		cmd := exec.Command("sh", "-c", cmdName)
		output, _ := cmd.CombinedOutput()
		fmt.Println(string(output))
	}
	fmt.Println(cmdName + "is illegal!")
}
