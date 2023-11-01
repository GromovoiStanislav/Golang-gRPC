package main // Замените на имя вашего пакета, если это необходимо

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"

	pb "main/proto"
)



func main() {
	
  {
    p := pb.Person{
      Id:    1234,
      Name:  "John Doe",
      Email: "jdoe@example.com",
      Phones: []*pb.Person_PhoneNumber{
          {Number: "555-4321", Type: pb.Person_PHONE_TYPE_HOME},
      },
    }

    book := &pb.AddressBook{}
    book.People = append(book.People, &p)


    // Write the new address book back to disk.
    out, err := proto.Marshal(book)
    if err != nil {
        log.Fatalln("Failed to encode address book:", err)
    }
    //fmt.Println(out)
    if err := ioutil.WriteFile("address-book", out, 0644); err != nil {
        log.Fatalln("Failed to write address book:", err)
    }
  }

  {
    // Read the existing address book.
    in, err := ioutil.ReadFile("address-book")
    if err != nil {
        log.Fatalln("Error reading file:", err)
    }
    book := &pb.AddressBook{}
    if err := proto.Unmarshal(in, book); err != nil {
        log.Fatalln("Failed to parse address book:", err)
    }


    
    p := book.GetPeople()
    fmt.Println(p)
    fmt.Println(p[0].GetId()) // 1234
    fmt.Println(p[0].GetName()) // John Doe
    fmt.Println(p[0].GetEmail()) // jdoe@example.com
    fmt.Println(p[0].GetPhones()) // [number:"555-4321" type:PHONE_TYPE_HOME]
    fmt.Println(p[0].GetPhones()[0].GetNumber()) //555-4321
    fmt.Println(p[0].GetPhones()[0].GetType()) //PHONE_TYPE_HOME

    fmt.Println(p[0].GetLastUpdated()) // nil
    
  }
}
