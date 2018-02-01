package main

import (
    "github.com/golang/protobuf/proto"
    "addr"
	"fmt"
	"os"
)

func main() {
	userFile := "log"
    newTest := &addr.Book{}

	fin,err := os.Open(userFile)
    defer fin.Close()
    if err != nil {
   		fmt.Println(userFile,err)
        return
    }
    buf := make([]byte, 1024)
    n,_ := fin.Read(buf)

    err = proto.Unmarshal(buf[:n], newTest)
    if err != nil {
        fmt.Printf("unmarshaling error: ", err)
    }else{
		fmt.Println(newTest)
	}

}
