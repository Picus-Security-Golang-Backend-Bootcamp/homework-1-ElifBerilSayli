package helper

import (
	"fmt"
	"strings"
)

func List(localBookSlice []string) {
	for _, v := range localBookSlice {
		fmt.Println(v)
	}
}

func Search(localBookSlice []string, bookName string) {
	fmt.Printf("Bulunan Kitaplar:")
	for _, v := range localBookSlice {
		tempV := strings.ToLower(v)
		tempBookName := strings.ToLower(bookName)
		if strings.Contains(tempV, tempBookName) {
			fmt.Printf(" %s \n", v)
		}
	}
}
