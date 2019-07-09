package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

//CheckAndCreateDir check and create dir
func CheckAndCreateDir(path string) error {
	if subPathExists, err := FileExists(path); err != nil {
		return fmt.Errorf("Could not determine if subPath %s exists; will not attempt to change its permissions", path)
	} else if !subPathExists {
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to mkdir:%s", path)
		}

		if err := os.Chmod(path, 0755); err != nil {
			return err
		}
	}
	return nil
}

//DirIsEmpty 验证目录是否为空
func DirIsEmpty(dir string) bool {
	infos, err := ioutil.ReadDir(dir)
	if len(infos) == 0 || err != nil {
		return true
	}
	return false
}

//OpenOrCreateFile open or create file
func OpenOrCreateFile(filename string) (*os.File, error) {
	return os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0777)
}

//FileExists check file exist
func FileExists(filename string) (bool, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

//SearchFileBody 搜索文件中是否含有指定字符串
func SearchFileBody(filename, searchStr string) bool {
	body, _ := ioutil.ReadFile(filename)
	return strings.Contains(string(body), searchStr)
}

//IsHaveFile 指定目录是否含有文件
//.开头文件除外
func IsHaveFile(path string) bool {
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if !strings.HasPrefix(file.Name(), ".") {
			return true
		}
	}
	return false
}

func IsFileExist(f string) bool {
	return IsExist(f) && IsFile(f)
}

func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

func IsFile(f string) bool {
	fi, e := os.Stat(f)
	if e != nil {
		return false
	}
	return !fi.IsDir()
}

// get path of this file
func getCurrentPath() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("Can not get current file info")
		// 错误，默认当前目录
		return "./"
	}
	lastIndex := strings.LastIndex(file, "/")
	if lastIndex < 0 {
		return "./"
	}

	path := file[:lastIndex+1]
	return path
}