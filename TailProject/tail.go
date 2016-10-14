package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		// 오류를 처리
		return
	}
	Beforesize, err := file.Seek(0, os.SEEK_END)
	if err != nil {
		return
	}

	for {
		//파일 변경 위치 값 CHECK
		Aftersize, err := file.Seek(0, os.SEEK_END)
		if err != nil {
			return
		}

		if Beforesize < Aftersize {
			//변경 전 위치 까지 이동
			reSize, err := file.Seek(Beforesize, 0)
			if err != nil {
				return
			}
			//슬라이스 선언
			s := make([]byte, Aftersize-reSize)
			//파일로 부터 정해진 크기만큼 데이터 읽기
			file.Read(s)

			fmt.Printf("%s", string(s))

			Beforesize = Aftersize
		}

	}

}
