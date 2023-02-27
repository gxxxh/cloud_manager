package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
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

// 资源前添加上路径名称
func JoinName(packagePath, packageName, split string) string {
	dirs := strings.Split(packagePath, "/")
	flag := false
	name := ""
	for _, dirName := range dirs {
		if dirName == packageName {
			flag = true
			continue
		}
		if flag {
			if name != "" {
				name += split
			}
			name += UpperFirst(dirName)
		}
	}
	return name
}

// 递归创建多级目录
func CreateMultiDir(dirPath string) error {
	pathExists, err := PathExists(dirPath)
	if err != nil {
		return err
	}
	if !pathExists {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// this path must be a file instead of directory
func Save(data []byte, path string) {
	dirPath := filepath.Dir(path)
	err := CreateMultiDir(dirPath)
	if err != nil {
		log.Println("create dir failed, error info: ", err)
	}

	file, err := os.Create(path)
	if err != nil {
		log.Fatalln("Open Code File  error, ", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln("Close Code File error, ", err)
		}
	}(file)
	_, err = file.Write(data)
	if err != nil {
		log.Fatalln("Write code to file error, ", err)
	}
}

func CreateDirByPackagePath(basePath, packagePath, packageName string) string {
	dirs := strings.Split(packagePath, "/")
	flag := false
	curPath := basePath
	for _, dirName := range dirs {

		if dirName == packageName {
			flag = true
			continue
		}
		if flag {
			curPath = filepath.Join(curPath, dirName)
			pathExists, err := PathExists(curPath)
			if err != nil {
				log.Println(err)
			}
			if pathExists == false {
				err = os.Mkdir(curPath, 0750)
				if err != nil {
					fmt.Println(err)
				}
			}

		}
	}
	return curPath
}
