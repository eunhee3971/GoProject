package main

import (
	"fmt"
	"os"
	"testing"
)

var testFILENAME = []string{
	"ccc",
	"aaa",
	"wtmp",
}

func TestGetFileInfo(t *testing.T) {

	for i := 0; i < len(testFILENAME); i++ {
		_, err := GetFileInfo(testFILENAME[i])
		if err != nil {
			t.Error("실패한 파일 :" + testFILENAME[i])
		} else {
			fmt.Println("성공한 파일 : " + testFILENAME[i])
		}
	}

}

func TestBinaryRead(t *testing.T) {
	var utmp UTMP

	file, _ := os.Open("wtmp")
	err := BinaryRead(file, &utmp)
	if err != nil {
		t.Error("실패한 경우")
	}
}
