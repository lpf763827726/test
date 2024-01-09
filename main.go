package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const maxLinesPerFile = 10

func main() {
	fileDir := "logs" // 文件夹路径
	stringsArray := []string{
		"Hello",
		"World",
		"Foo",
		"Bar",
		"Baz",
		"Lorem",
		"Ipsum",
		"Test",
		"Data",
		"Golang",
		"File",
		"Hello",
		"World",
		"Foo",
		"Bar",
		"Baz",
		"Lorem",
		"Ipsum",
		"Test",
		"Data",
		"Golang",
		"File",
		"Hello",
		"World",
		"Foo",
		"Bar",
		"Baz",
		"Lorem",
		"Ipsum",
		"Test",
		"Data",
		"Golang",
		"File",
		"Hello",
		"World",
		"Foo",
		"Bar",
		"Baz",
		"Lorem",
		"Ipsum",
		"Test",
		"Data",
		"Golang",
		"File",
		"Hello",
		"World",
		"Foo",
		"Bar",
		"Baz",
		"Lorem",
		"Ipsum",
		"Test",
		"Data",
		"Golang",
		"File",
		"Hello",
		"World",
		"Foo",
		"Bar",
		"Baz",
		"Lorem",
		"Ipsum",
		"Test",
		"Data",
		"Golang",
		"File",
		"Hello",
		"World",
		"Foo",
		"Bar",
		"Baz",
		"Lorem",
		"Ipsum",
		"Test",
		"Data",
		"Golang",
		"File",
		"Hello",
		"World",
		"Foo",
		"Bar",
		"Baz",
		"Lorem",
		"Ipsum",
		"Test",
		"Data",
		"Golang",
		"File",
		"Hello",
		"World",
		"Foo",
		"Bar",
		"Baz",
		"Lorem",
		"Ipsum",
		"Test",
		"Data",
		"Golang",
		"File",
		"Hello",
		"World",
		"Foo",
		"Bar",
		"Baz",
		"Lorem",
		"Ipsum",
		"Test",
		"Data",
		"Golang",
		"File",
		"Hello",
		"World",
		"Foo",
		"Bar",
		"Baz",
		"Lorem",
		"Ipsum",
		"Test",
		"Data",
		"Golang",
		"File",
		"Hello",
		"World",
		"Foo",
		"Bar",
		"Baz",
		"Lorem",
		"Ipsum",
		"Test",
		"Data",
		"Golang",
		"File",
		"Hello",
		"World",
		"Foo",
		"Bar",
		"Baz",
		"Lorem",
		"Ipsum",
		"Test",
		"Data",
		"Golang",
		"File",
		"Hello",
		"World",
		"Foo",
		"Bar",
		"Baz",
		"Lorem",
		"Ipsum",
		"Test",
		"Data",
		"Golang",
		"File",
		"Hello",
		"World",
		"Foo",
		"Bar",
		"Baz",
		"Lorem",
		"Ipsum",
		"Test",
		"Data",
		"Golang",
		"File",
		"Hello",
		"World",
		"Foo",
		"Bar",
		"Baz",
		"Lorem",
		"Ipsum",
		"Test",
		"Data",
		"Golang",
		"File",
		"Hello",
		"World",
		"Foo",
		"Bar",
		"Baz",
		"Lorem",
		"Ipsum",
		"Test",
		"Data",
		"Golang",
		"File",
		"Hello",
		"World",
		"Foo",
		"Bar",
		"Baz",
		"Lorem",
		"Ipsum",
		"Test",
		"Data",
		"Golang",
		"File",
		"Hello",
		"World",
		"Foo",
		"Bar",
		"Baz",
		"Lorem",
		"Ipsum",
		"Test",
		"Data",
		"Golang",
		"File",
		"Hello",
		"World",
		"Foo",
		"Bar",
		"Baz",
		"Lorem",
		"Ipsum",
		"Test",
		"Data",
		"Golang",
		"File",
	}

	err := writeFile(stringsArray, fileDir)
	if err != nil {
		fmt.Println("写入文件失败:", err)
		return
	}

	err = printFilesContent(fileDir)
	if err != nil {
		fmt.Println("打印文件内容失败:", err)
		return
	}
}

func writeFile(stringsArray []string, fileDir string) error {
	lineCount := 0
	fileCount := 1
	filePath := getFilePath(fileDir, fileCount)

	for _, str := range stringsArray {
		err := appendToFile(filePath, str)
		if err != nil {
			return err
		}

		lineCount++

		if lineCount >= maxLinesPerFile {
			fileCount++
			filePath = getFilePath(fileDir, fileCount)
			lineCount = 0
		}
	}

	return nil
}

func appendToFile(filePath string, str string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(str + "\n")
	if err != nil {
		return err
	}

	return writer.Flush()
}

func getFilePath(fileDir string, fileCount int) string {
	fileName := fmt.Sprintf("file%d.txt", fileCount)
	return filepath.Join(fileDir, fileName)
}

func printFilesContent(fileDir string) error {
	files, err := ioutil.ReadDir(fileDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(fileDir, file.Name())
			fmt.Printf("=== File: %s ===\n", filePath)
			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				return err
			}
			fmt.Println(string(content))
			fmt.Println()
		}
	}

	return nil
}
