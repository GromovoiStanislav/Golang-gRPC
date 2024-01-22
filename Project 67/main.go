package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "protobuf-example/testpackage"
)

func main() {

    // Write to file
    {
        data := &pb.TestMessage{
                Message: "Helo from Glang!!!",
                CreatedAt: timestamppb.Now(),
        }
        fmt.Println(data.CreatedAt.AsTime()) // 2024-01-22 11:00:27.1095287 +0000 UTC

        out, err := proto.Marshal(data)
        if err != nil {
                log.Fatalln("Failed to encode address book:", err)
        }
        if err := ioutil.WriteFile("from-golang.bin", out, 0644); err != nil {
                log.Fatalln("Failed to write address book:", err)
        }
    }


    // Read from file
    {
        in, err := ioutil.ReadFile("from-nodejs.bin")
        if err != nil {
                log.Fatalln("Error reading file:", err)
        }
        data := &pb.TestMessage{}
        if err := proto.Unmarshal(in, data); err != nil {
                log.Fatalln("Failed to parse address book:", err)
        }

        fmt.Println(data) // createdAt:{seconds:1705920062 nanos:978000000} message:"A rose by any other name would smell as sweet"
        fmt.Println(data.CreatedAt.AsTime()) // 2024-01-22 10:41:02.978 +0000 UTC
    }
    
}

