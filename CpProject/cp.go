package main

import (
	"io/ioutil"
	"os"
)

func main() {

	bytes, err := ioutil.ReadFile(os.Args[1])
	//파일 읽기
	if err != nil {
		panic(err)
	}

	//파일 쓰기
	err = ioutil.WriteFile(os.Args[2], bytes, 0)
	if err != nil {
		panic(err)
	}
}
