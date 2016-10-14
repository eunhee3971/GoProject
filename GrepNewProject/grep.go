package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 2 {

		//파일 열기
		if file, err := os.Open(os.Args[2]); err == nil {
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				if strings.Contains(scanner.Text(), os.Args[1]) {
					fmt.Println(scanner.Text())
				}
			}

			if err = scanner.Err(); err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	} else if len(os.Args) == 2 {

		consolescanner := bufio.NewScanner(os.Stdin)

		// by default, bufio.Scanner scans newline-separated lines
		for consolescanner.Scan() {
			input := consolescanner.Text()
			if strings.Contains(input, os.Args[1]) {
				fmt.Println(input)
			}
		}

	} else {
		fmt.Println("Argument Count Error")
	}

}
