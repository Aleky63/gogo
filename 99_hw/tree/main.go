package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func dirTree(path string, showFiles bool) {
	// Читаем содержимое директории
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading directory %s: %v\n", path, err)
		return
	}

	// Сортируем по имени (гарантируем порядок, даже если ReadDir уже сортирует)
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})

	// Печатаем дерево
	printEntries(entries, path, "", showFiles)
}

func printEntries(entries []os.DirEntry, basePath string, prefix string, showFiles bool) {
	lastIndex := len(entries) - 1

	for i, entry := range entries {
		isLast := i == lastIndex
		entryPath := filepath.Join(basePath, entry.Name())

		// Определяем префикс для этого элемента
		var linePrefix string
		if isLast {
			linePrefix = prefix + "└───"
		} else {
			linePrefix = prefix + "├───"
		}

		// Выводим текущий элемент
		if entry.IsDir() {
			// Это директория — выводим только имя
			fmt.Println(linePrefix + entry.Name())

			// Рекурсивно обходим поддиректорию
			newPrefix := prefix
			if isLast {
				newPrefix += "\t" // последний — без вертикальной линии
			} else {
				newPrefix += "│\t" // не последний — с вертикальной линией
			}

			subEntries, err := os.ReadDir(entryPath)
			if err == nil {
				sort.Slice(subEntries, func(i, j int) bool {
					return subEntries[i].Name() < subEntries[j].Name()
				})
				printEntries(subEntries, entryPath, newPrefix, showFiles)
			}
		} else {
			// Это файл — выводим только если showFiles == true
			if showFiles {
				info, err := os.Stat(entryPath)
				if err != nil {
					fmt.Println(linePrefix + entry.Name())
					continue
				}
				size := info.Size()
				sizeStr := "empty"
				if size > 0 {
					sizeStr = fmt.Sprintf("%db", size)
				}
				fmt.Println(linePrefix + entry.Name() + " (" + sizeStr + ")")
			} else {
				// Не показываем файлы, если -f не указан
				// Ничего не делаем
			}
		}
	}
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <directory> [-f]\n", args[0])
		os.Exit(1)
	}

	path := args[1]
	showFiles := false

	if len(args) > 2 && args[2] == "-f" {
		showFiles = true
	}

	dirTree(path, showFiles)
}
