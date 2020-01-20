package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// 被讀取的檔案
	fName := os.Args[1]

	file, err := os.Open(fName)
	if err != nil {
		log.Fatal(err)
	}

	// 第一種印出內容的方式 (不會分行，會保留 \n)
	// data := make([]byte, 100)
	// count, err := file.Read(data)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("read %d bytes: %q\n", count, data[:count])

	// 第二種印出內容的方式 (會分行)
	io.Copy(os.Stdout, file)
}
