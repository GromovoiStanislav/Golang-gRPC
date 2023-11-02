package main

import (
	"io/ioutil"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
	
	pb "main/proto"
)





func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage:  %s ADDRESS_BOOK_FILE\n", os.Args[0])
	}
	fname := os.Args[1]


	in := pb.AddressBook{People: []*pb.Person{
		{
			Name:  "John Doe",
			Id:    101,
			Email: "john@example.com",
		},
		{
			Name: "Jane Doe",
			Id:   102,
		},
		{
			Name:  "Jack Doe",
			Id:    201,
			Email: "jack@example.com",
			Phones: []*pb.Person_PhoneNumber{
				{Number: "555-555-5555", Type: pb.Person_WORK},
			},
		},
		{
			Name:  "Jack Buck",
			Id:    301,
			Email: "buck@example.com",
			Phones: []*pb.Person_PhoneNumber{
				{Number: "555-555-0000", Type: pb.Person_HOME},
				{Number: "555-555-0001", Type: pb.Person_MOBILE},
				{Number: "555-555-0002", Type: pb.Person_WORK},
			},
		},
		{
			Name:  "Janet Doe",
			Id:    1001,
			Email: "janet@example.com",
			Phones: []*pb.Person_PhoneNumber{
				{Number: "555-777-0000"},
				{Number: "555-777-0001", Type: pb.Person_HOME},
			},
		},
	}}


	buf, err := proto.Marshal(&in)
    if err != nil {
        log.Fatalln("Failed to encode address book:", err)
    }
	//fmt.Println(buf)

	if err := ioutil.WriteFile(fname, buf, 0644); err != nil {
        log.Fatalln("Failed to write address book:", err)
    }
}