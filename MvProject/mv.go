package main

import (
	"log"
	"os"
	"strings"
)

func main() {

	//Directory가 없는 경우 Directory 생성
	splitStr := strings.Split(os.Args[2], "/")

	for i := 0; i < len(splitStr)-1; i++ {
		if _, err := os.Stat(splitStr[i]); os.IsNotExist(err) {
			os.Mkdir(splitStr[i], 0700)
		}
	}

	//HardLink생성
	err := os.Link(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}

	//파일 삭제
	os.Remove(os.Args[1])

}
