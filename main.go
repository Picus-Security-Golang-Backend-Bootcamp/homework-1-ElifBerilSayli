package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-1-ElifBerilSayli/helper"
)

func main() {
	args := os.Args
	bookSlice := make([]string, 5)
	var bookNameSlice []string
	var bookName string

	bookSlice[0] = "In Search of Lost Time"
	bookSlice[1] = "Ulysses"
	bookSlice[2] = "Don Quixote"
	bookSlice[3] = "The Great Gatsby"
	bookSlice[4] = "The Great Gatsby Second"

	if len(args) == 1 {
		projectName := path.Base(args[0])
		fmt.Printf("%s uygulamasında kullanabileceğiniz komutlar : \n search => arama işlemi için \n list => listeleme işlemi için\n", projectName)
		return
	}
	if args[1] == "list" {
		helper.List(bookSlice)
		return
	}
	if args[1] == "search" {

		if len(args) == 2 {
			fmt.Println("Arama işlemi için kitap ismi yazılmalı !! \n  go run main.go search The Great Gatsby gibi")
			return
		}

		for i := 2; i < len(args); i++ {
			bookNameSlice = append(bookNameSlice, args[i])
		}
		bookName = strings.Join(bookNameSlice, " ")
		helper.Search(bookSlice, bookName)
	}
}
