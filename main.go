package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Укажите полный путь до файла вторым аргументом")
	}

	filePth := os.Args[1]

	var fileName, fileExt, fileStr string

	// Напишите код, который выведет следующее
	// filename: <name>
	// extension: <extension>

	// Подсказка. Возможно вам понадобится функция strings.LastIndex
	// Для проверки своего решения используйте функции filepath.Base() filepath.Ext(
	// ) Они могут помочь для проверки решения

	// fileName = filepath.Base(filePth)
	// fileExt = filepath.Ext(fileName)

	idx := strings.LastIndex(filePth, "/")
	if idx > 0 {
		fileStr = filePth[idx+1:]
	} else {
		fileStr = filePth
	}

	idx = strings.LastIndex(fileStr, ".")
	if idx > 0 {
		fileName = fileStr[:idx]
		fileExt = fileStr[idx+1:]
	} else {
		fileName = fileStr
	}

	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)
}
