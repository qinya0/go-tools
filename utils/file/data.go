package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 从文件按行读取，保存到[]string返回
func GetDataFromFile(fileName string) ([]string, error) {
	fileName = GetCurrentPath() + fileName
	resStr := []string{}
	if !IsFileExist(fileName) {
		err := fmt.Errorf("not found file %s because it not exist", fileName)
		return nil, err
	}

	file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}

	rd := bufio.NewReader(file)
	for {
		// 按行读取
		line, err := rd.ReadString('\n')
		if err != nil {
			// 文件终止符，默认读取完成
			if io.EOF == err {
				resStr = append(resStr, line)
				break
			} else {
				return nil, err
			}
		}
		resStr = append(resStr, line)
	}

	return resStr, nil
}

func SaveByte(fileName string, data []byte) error {
	return saveDataToFile(fileName, string(data))
}

func SaveString(fileName string, data string) error {
	return saveDataToFile(fileName, data)
}

// 将data保存到文件
func saveDataToFile(fileName string, data string) error {
	fileName = GetCurrentPath() + fileName
	addEnter := true
	//fmt.Printf("[test] fielName:%s\n", fileName)
	if !IsFileExist(fileName) {
		if _, err := os.Create(fileName); err != nil {
			return err
		}
		addEnter = false
	}

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0640)
	if err != nil {
		return err
	}

	addStr := data
	if addEnter {
		addStr = "\n"+addStr
	}

	if _, err := file.Write([]byte(addStr)); err != nil {
		return err
	}

	return nil
}
