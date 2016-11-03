package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strings"
	"unsafe"
)

//const
const (
	UTUNKNOWN    = 0
	RUNLVL       = 1
	BOOTTIME     = 2
	NEWTIME      = 3
	OLDTIME      = 4
	INITPROCESS  = 5
	LOGINPROCESS = 6
	USERPROCESS  = 7
	DEADPROCESS  = 8
	ACCOUNTING   = 9

	UTLINESIZE = 32
	UTNAMESIZE = 32
	UTHOSTSIZE = 256
)

//ExitStatus struct
type ExitStatus struct {
	ETermination uint16
	EExit        uint16
}

//TIMEVAL struct
type TIMEVAL struct {
	TvSec  uint32
	TvUsec uint32
}

//UTMP struct
type UTMP struct {
	UtType uint16
	_      uint16
	UtPid  uint32
	UtLine [UTLINESIZE]byte
	UtId   [4]byte
	UtUser [UTNAMESIZE]byte
	UtHost [UTHOSTSIZE]byte
	UtExit ExitStatus

	UtSession uint32
	UtTv      TIMEVAL
	UtAddrV6  [4]uint32
	Unused    [20]byte
}

func main() {

	if len(os.Args) == 2 {

		strPath := os.Args[1]
		fileInfo, err := GetFileInfo(strPath)
		defer fileInfo.Close()

		for {
			utmp := UTMP{}
			err = binary.Read(fileInfo, binary.LittleEndian, &utmp)

			if err != nil {
				if err == io.EOF || err == io.ErrUnexpectedEOF {
					return
				}
				panic(err)
			}
			PrintLine(utmp)
		}

	} else if len(os.Args) == 3 {

		strPath := os.Args[2]

		fileInfo, err := GetFileInfo(strPath)
		defer fileInfo.Close()

		fileSize := GetFileSize(fileInfo)

		utmp := UTMP{}
		structSize := unsafe.Sizeof(utmp)

		fileInfo.Seek(fileSize-int64(structSize)*5, os.SEEK_SET)
		CheckError(err)

		for {
			err = binary.Read(fileInfo, binary.LittleEndian, &utmp)
			if err != nil {
				if err == io.EOF || err == io.ErrUnexpectedEOF {
					for {
						Aftersize := GetFileSize(fileInfo)
						if fileSize < Aftersize {
							binary.Read(fileInfo, binary.LittleEndian, &utmp)
							PrintLine(utmp)
							fileSize = Aftersize
						}
					}

				}
			}
			PrintLine(utmp)
		}

	} else {
		fmt.Println("Argument Error")
		fmt.Println("use BinaryLog PATH or BinaryLog OPTIONS PATH ")
		os.Exit(0)
	}

}

// CheckError A Simple function to verify error
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

//PrintLine A simple function to print wtmp
func PrintLine(utmp UTMP) {

	fmt.Printf("ut_type=%d, ut_pid=%d, ut_line=%s, ut_id =%s, ut_user=%s, ut_host=%s, e_termination=%d, e_exit=%d, ut_session=%d, tv_sec=%d, tv_usec=%d, ut_addr_v6=%d, ut_unused=%s\n",
		utmp.UtType,
		utmp.UtPid,
		strings.Trim(string(utmp.UtLine[:]), "\x00"),
		strings.Trim(string(utmp.UtId[:]), "\x00"),
		strings.Trim(string(utmp.UtUser[:]), "\x00"),
		strings.Trim(string(utmp.UtHost[:]), "\x00"),
		utmp.UtExit.ETermination,
		utmp.UtExit.EExit,
		utmp.UtSession,
		utmp.UtTv.TvSec,
		utmp.UtTv.TvUsec,
		utmp.UtAddrV6,
		strings.Trim(string(utmp.Unused[:]), "\x00"))
}

//GetFileInfo function
func GetFileInfo(strPath string) (*os.File, error) {

	fileInfo, err := os.OpenFile(strPath, os.O_RDONLY, os.FileMode(0644))

	return fileInfo, err
}

//GetFileSize function
func GetFileSize(FileInfo *os.File) (size int64) {

	fileStat, err := FileInfo.Stat()
	CheckError(err)
	fileSize := fileStat.Size()

	return fileSize
}

//BinaryRead function
func BinaryRead(file *os.File, utmp *UTMP) (err error) {
	err = binary.Read(file, binary.LittleEndian, utmp)
	return err
}
