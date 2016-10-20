package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

// CheckError A Simple function to verify error
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: host : %s port : %s", os.Args[1], os.Args[2])
		os.Exit(1)
	}

	Port := os.Args[1]
	strPath := os.Args[2]
	bCurrNow := time.Now()
	bCurrDate := bCurrNow.Format("20060102150405")

	ServerAddr, err := net.ResolveUDPAddr("udp", ":"+Port)
	CheckError(err)

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	ServerConn.SetReadBuffer(1000000)
	defer ServerConn.Close()

	fileInfo, err := os.OpenFile(strPath+"/"+bCurrDate+".txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	CheckError(err)

	buf := make([]byte, 1024)

	for {

		aCurrNow := time.Now()
		aCurrDate := aCurrNow.Format("20060102150405")

		ServerConn.SetReadDeadline(time.Now().Add(5 * time.Second))

		n, addr, err := ServerConn.ReadFromUDP(buf)
		CheckError(err)

		sAddr := strings.Split(addr.String(), ":")

		sinput := fmt.Sprintf("[%s] [%s] [%s] \r\n", sAddr[0], aCurrDate, string(buf[0:n]))

		w := bufio.NewWriter(fileInfo)
		w.WriteString(sinput)
		w.Flush()

		if bCurrDate[:12] != aCurrDate[:12] {
			defer fileInfo.Close()
			fileInfo, err = os.OpenFile(strPath+"/"+aCurrDate+".txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
			CheckError(err)
			bCurrDate = aCurrDate
		}
	}
}
