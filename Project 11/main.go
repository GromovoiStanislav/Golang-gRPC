package main // Замените на имя вашего пакета, если это необходимо

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"

	pb "main/proto"
)



func main() {
	elliot := &pb.Person{
        Name: "Elliot",
        Age:  24,
        SocialFollowers: &pb.SocialFollowers{
          Youtube: 2500,
          Twitter: 1400,
      },
    }

    data, err := proto.Marshal(elliot)
    if err != nil {
        log.Fatal("marshaling error: ", err)
    }

  // printing out our raw protobuf object
    fmt.Println(data)

  // let's go the other way and unmarshal
  // our byte array into an object we can modify
  // and use
    newElliot := &pb.Person{}
    err = proto.Unmarshal(data, newElliot)
    if err != nil {
        log.Fatal("unmarshaling error: ", err)
  }

  // print out our `newElliot` object
  // for good measure
  fmt.Println(newElliot.GetAge())
  fmt.Println(newElliot.GetName())
  fmt.Println(newElliot.SocialFollowers.GetTwitter())
  fmt.Println(newElliot.SocialFollowers.GetYoutube())
}
