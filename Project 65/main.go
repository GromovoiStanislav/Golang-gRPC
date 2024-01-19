package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"

	pb "protobuf-example/person"
)

func main() {

    fname := "AddressBook.bin"
    
    p := &pb.Person{
        Id:    1234,
        Name:  "John Doe",
        Email: "jdoe@example.com",
        Phones: []*pb.Person_PhoneNumber{
                {Number: "996-4321", Type: pb.Person_HOME},
                {Number: "555-4321", Type: pb.Person_MOBILE},
        },
    }

    book := &pb.AddressBook{}
    book.People = append(book.People, p)

    // Write the new address book back to disk
    {
        out, err := proto.Marshal(book)
        if err != nil {
                log.Fatalln("Failed to encode address book:", err)
        }
        if err := ioutil.WriteFile(fname, out, 0644); err != nil {
                log.Fatalln("Failed to write address book:", err)
        }
    }


    // Read the existing address book
    {
        in, err := ioutil.ReadFile(fname)
        if err != nil {
                log.Fatalln("Error reading file:", err)
        }
        book := &pb.AddressBook{}
        if err := proto.Unmarshal(in, book); err != nil {
                log.Fatalln("Failed to parse address book:", err)
        }

        fmt.Println(book) // people:{name:"John Doe" id:1234 email:"jdoe@example.com" phones:{number:"996-4321" type:HOME} phones:{number:"555-4321"}}
    }
    
}

