package utils

import (
	"errors"
	"io/ioutil"
	"os"
)

// 判断所给路径文件/文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	//isnotexist来判断，是不是不存在的错误
	if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
		return false, nil
	}
	return false, err //如果有错误了，但是不是不存在的错误，所以把这个错误原封不动的返回
}

func Readfile(path string) (string, error) {
	exist, err := PathExists(path)
	if err != nil {
		return "", err
	}
	if !exist {
		return "", errors.New("file not exist")
	}
	content, err := ioutil.ReadFile(path)
	return string(content), err
}

func GetPackageName(name string) string {
	packageName := ""
	for _, ch := range name {
		if ch == '.' {
			break
		}
		packageName += string(ch)
	}
	return packageName
}
