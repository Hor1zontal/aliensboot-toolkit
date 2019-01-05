/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/10/29
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var globalFilters = make(map[string]struct{})
//
func AddFilter(name string) {
	globalFilters[name] = struct{}{}
}

func IsFilter(name string, filters []string) bool {
	for _, filter := range filters {
		if strings.Contains(name, filter) {
			return true
		}
	}

	for filter, _  := range globalFilters {
		if strings.Contains(name, filter) {
			return true
		}
	}
	return false
}

func CopyDir(srcPath string, destPath string, replaceContent map[string]string, filters ...string) error {
	//检测目录正确性
	if srcInfo, err := os.Stat(srcPath); err != nil {
		fmt.Println(err.Error())
		return err
	} else {
		if !srcInfo.IsDir() {
			e := errors.New("srcPath不是一个正确的目录！")
			fmt.Println(e.Error())
			return e
		}

		srcInfo.Name()
	}

	if destInfo, err := os.Stat(destPath); err != nil {
		err := os.MkdirAll(destPath, os.ModePerm)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	} else {
		if !destInfo.IsDir() {
			e := errors.New("destInfo不是一个正确的目录！")
			fmt.Println(e.Error())
			return e
		}
	}
	//加上拷贝时间:不用可以去掉
	//destPath = destPath + "_" + time.Now().Format("20060102150405")

	err := filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if !f.IsDir() {
			if IsFilter(path, filters) {
				fmt.Println("filter file :" + path)
				return nil
			}
			path := strings.Replace(path, "\\", "/", -1)
			srcPath = strings.Replace(srcPath, "\\", "/", -1)
			destNewPath := strings.Replace(path, srcPath, destPath, -1)
			//fmt.Println("copy file :" + path + " to " + destNewPath)
			CopyFile(path, destNewPath, replaceContent)
		}
		return nil
	})
	if err != nil {
		fmt.Printf(err.Error())
	}
	return err
}

//生成目录并拷贝文件
func CopyFile(src, dest string, replaceContent map[string]string) {
	data := ReadFile(src)
	if data == nil {
		return
	}
	//分割path目录
	destSplitPathDirs := strings.Split(dest, "/")
	//检测时候存在目录
	destSplitPath := ""
	for index, dir := range destSplitPathDirs {
		if index < len(destSplitPathDirs)-1 {
			destSplitPath = destSplitPath + dir + "/"
			b, _ := PathExists(destSplitPath)
			if b == false {
				fmt.Println("create diretory:" + destSplitPath)
				//创建目录
				err := os.Mkdir(destSplitPath, os.ModePerm)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}

	content := string(data)

	if replaceContent != nil {
		for replaceKey, replaceValue := range replaceContent {
			content = strings.Replace(content, replaceKey, replaceValue, -1)
		}
	}

	WriteFile(dest, []byte(content))
	//return io.Copy(dstFile, srcFile)
}

func WriteFile(filePath string, content []byte) {
	dstFile, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	_, err1 := dstFile.Write(content) //写入文件(字节数组)
	if err1 != nil {
		fmt.Printf(err1.Error())
	}
}

func ReadFile(filePath string) []byte {
	srcFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer srcFile.Close()
	data, err := ioutil.ReadAll(srcFile)
	if err != nil {
		fmt.Println(err.Error())
	}
	return data
}

//检测文件夹路径时候存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
