package main

import (
    "github.com/golang/protobuf/proto"
    "addr"
	"fmt"
	"os"
)

func main() {
    test := &addr.Book{
		Id : proto.Int32(11),
		Str : proto.String("testMsg11"),
    }

    data, err := proto.Marshal(test)
    if err != nil {
        fmt.Printf("marshaling error: ", err)
    }else{
		//fmt.Println(data)
	    f,_ := os.Create("log")
    	defer f.Close()
		_, _ = f.Write(data)
	}

}
