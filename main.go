package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
)

// определение структуры файла
type Files struct {
	name      string
	extension string
	size      int64
}

var s []Files

// определение функции для ввода информации классы Files в консоль
func (ob *Files) print() {
	fmt.Println("Name:", ob.name, "Type:", ob.extension, "FileSize/byte", ob.size)
}

// определение функции для получения строк через консоль
func getFilePathFromCommand(root string, sort string) (string, string, error) {
	if root == "None" || sort == "None" {
		fmt.Println("->Введите правильную командную строку:(--root=/pathfile  --sort=Desc) or --root=/pathfile")
	}
	var sourcepath *string
	var sortflag *string
	sourcepath = flag.String(root, "None", "")
	sortflag = flag.String(sort, "None", "")
	flag.Parse()
	return *sourcepath, *sortflag, nil
}

// функция для проверкаи попки
func rootExist(root string) (bool, error) {
	_, err := os.Stat(root)
	if os.IsNotExist(err) {
		fmt.Println("Root не существует...!")
	}
	return true, nil
}

func getRoot(root string) (string, error) {
	var rootflag *string
	rootflag = flag.String(root, "None", "")
	flag.Parse()
	_, err := rootExist(*rootflag)
	if err != nil {
		panic(err)
	}
	return *rootflag, nil
}

// метод для получения значения size класса
func (ob *Files) getSize() int64 {
	return ob.size
}

// метод для получения значения name класса
func (ob *Files) getName() string {
	return ob.name
}

// метод для получения значения Extension класса
func (ob *Files) getExtension() string {
	return ob.extension
}

// метод для получение информации о файлах
func getFilesRecurvise(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	return true, nil
}

// метод для получение информации католога файлы
func getFileLocation(root string, filename string) (string, error) {
	if root == "" {
		return "", errors.New("Root  пуст!")
	}
	return root + "/" + filename, nil
}

// Получение все файл из котолога
func getAllFromDir(path string) ([]Files, error) {
	err := filepath.Walk(path, func(p string, inf os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		size, err := getsize(p)
		if err != nil {
			fmt.Println(err)
		}
		Ext, err := getFileExtension2(p)
		if err != nil {
			fmt.Println(err)
		}
		element := Files{name: p, extension: Ext, size: size}
		s = append(s, element)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return s, nil
}

// функция для получения значения  size
func getsize(filename string) (int64, error) {
	f, err := os.Stat(filename)
	if err != nil {
		fmt.Println(err)
	}
	return f.Size(), nil
}

// функция для получения значения  Extension
func getFileExtension(root string, filename string) (string, error) {
	f, err := getFileLocation(root, filename)
	if err != nil {
		fmt.Println(err)
	}
	st, err := os.Stat(f)
	if err != nil {
		fmt.Println(err)
	}
	if st.IsDir() {
		return "Каталог", nil
	}
	return "файл", nil
}

// функция для получения значения  Extension2
func getFileExtension2(filename string) (string, error) {
	f, err := os.Stat(filename)
	if err != nil {
		fmt.Println(err)
	}
	if f.IsDir() {
		return "Каталог", nil
	}
	return "файл", nil
}

// функция для получения значения  файлы из католога
func getFilesFromDirectory(pathName string) ([]Files, error) {
	fi, err := os.Open(pathName)
	if err != nil {
		log.Fatal(err, fi.Name())
	}
	defer fi.Close()
	files, err := os.ReadDir(pathName)
	if err != nil {
		fmt.Print("Невозможно прочитать из каталога!", err)
	}
	for _, item := range files {
		p, err := getFileLocation(pathName, item.Name())
		f, err := os.Stat(p)
		if err != nil {
			panic(err)
		}
		Ext, err := getFileExtension(pathName, item.Name())
		name := pathName + "/" + f.Name()
		element := Files{name: name, extension: Ext, size: f.Size()}
		s = append(s, element)
	}
	return s, nil
}

// функция для Обработки сортировки по Убывающий
func sortAsc(arr []Files) {
	if len(arr) < 0 {
		fmt.Println("Массив пуст!")
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].size < arr[j].size
	})
}

// функция для Обработки сортировки по возврастающий
func sortDesc(arr []Files) {
	if len(arr) < 0 {
		fmt.Println("Массив пуст!")
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].size > arr[j].size
	})
}

// выборка сортировки
func selectSord(struc []Files, root string, sortMode string) error {
	if len(struc) < 0 {
		log.Panic("Нет элементов в массиве!")
	}
	if root != "None" && sortMode == "None" {
		sortAsc(struc)
		for i := 0; i < len(struc); i++ {
			struc[i].print()
		}
	} else if sortMode == "Desc" && root != "None" {
		sortDesc(struc)
		for i := 0; i < len(struc); i++ {
			struc[i].print()
		}
	}
	return nil
}

func main() {
	rootflag := "root"
	sortflag := "sort"
	root, sort, err := getFilePathFromCommand(rootflag, sortflag)
	files, err := getFilesFromDirectory(root)
	if err != nil {
		fmt.Println(err)
	}
	selectSord(files, root, sort)
}
