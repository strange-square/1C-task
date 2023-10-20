package main

import (
	"fmt"
	"github.com/strange-square/1C-task/internal/comparator"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Необходимо указать две директории и порог сходства")
		return
	}
	dir1 := os.Args[1]
	dir2 := os.Args[2]
	thresholdStr := os.Args[3]
	threshold, err := strconv.ParseFloat(thresholdStr, 64)
	if err != nil {
		fmt.Println("Необходимо указать порог схожести в формате числа с плавающей точкой. Пример: 0.33")
		return
	}
	if threshold < 0.0 || threshold > 1.0 {
		fmt.Println("Порог должен лежать в диапазоне от 0 до 1. Пример: 0.33")
		return
	}

	result, err := comparator.CompareDirs(dir1, dir2, threshold)
	if err != nil {
		fmt.Println("Ошибка при сравнении файлов:", err)
	}

	fmt.Println("Идентичные файлы:")
	if len(result.GetIdenticalFiles()) != 0 {
		for file1, files2 := range result.GetIdenticalFiles() {
			for _, file2 := range files2 {
				fmt.Println(file1, " - ", file2)
			}
		}
	} else {
		fmt.Println("Не найдено")
	}

	fmt.Println("Похожие файлы:")
	if len(result.GetSimilarFiles()) != 0 {
		for file1, files2 := range result.GetSimilarFiles() {
			for _, file2 := range files2 {
				fmt.Println(file1, " - ", file2)
			}
		}
	} else {
		fmt.Println("Не найдено")
	}

	fmt.Printf("Файлы, найденные в %s, но не найденные в %s:\n", dir1, dir2)
	if len(result.GetNotFoundFilesFromFirstDir()) != 0 {
		for _, file := range result.GetNotFoundFilesFromFirstDir() {
			fmt.Println(file)
		}
	} else {
		fmt.Println("Не найдено")
	}

	fmt.Printf("Файлы, найденные в %s, но не найденные в %s:\n", dir2, dir1)
	if len(result.GetNotFoundFilesFromSecondDir()) != 0 {
		for _, file := range result.GetNotFoundFilesFromSecondDir() {
			fmt.Println(file)
		}
	} else {
		fmt.Println("Не найдено")
	}
}
