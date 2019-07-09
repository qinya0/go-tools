package root

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "a test for go-tool",
	}
	dirCmd = &cobra.Command{
		Use: "dir",
		Short: "get now dir",
		Run: DirFunc,
	}
)

func init() {
	testCmd.AddCommand(dirCmd)
}

func DirFunc(cmd *cobra.Command, args []string) {
	printFile(-1)
	printFile(0)
	printFile(1)
	printFile(2)
	printFile(3)

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Printf("get file path err:%s \n", err.Error())
		return
	}
	fmt.Printf("dir:%s \n", dir)
}

func printFile(i int) {
	_, file, line, ok := runtime.Caller(i)
	if !ok {
		fmt.Println("Can not get current file info")
		// 错误，默认当前目录
		return
	}
	lastIndex := strings.LastIndex(file, "/")
	if lastIndex < 0 {
		return
	}

	path := file[:lastIndex+1]
	fmt.Printf("path:%s line:%d\n", path, line)
	return
}