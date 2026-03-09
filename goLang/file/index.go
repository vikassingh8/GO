package main

import (
	"fmt"
	"os"
)



func main() {

	// file, err := os.Open("./vikas.txt")
	// if err !=nil{
	// 	panic(err)
	// }
	// fileInfo,err:=file.Stat()
	// if err !=nil{
	// 	panic(err)
	// }
	
	// fmt.Println("File Name:",fileInfo.IsDir())

	// defer file.Close()

	// fileContent := make([]byte, fileInfo.Size())
	// file.Read(fileContent)
	// fmt.Println("File Content:", string(fileContent))

	dir,err :=os.Open("../")
	if err !=nil{
		panic(err)
	}
	defer dir.Close()
	f,err:=dir.ReadDir(-1)
	if err !=nil{
		panic(err)
	}

	for _,file:=range f{
		fmt.Println(file.Name())
	}
}