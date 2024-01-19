package main


import (
    "fmt"
    "io/ioutil"
    "log"

    "github.com/golang/protobuf/proto"

    pb "protobuf-example/person"
)

func main() {
    person := &pb.Person{Name: "XXX"}
    fmt.Printf("Person's name is %s\n", person.GetName())

    //Now lets write this person object to file
    out, err := proto.Marshal(person)
    if err != nil {
        log.Fatalf("Serialization error: %s", err.Error())
    }
    if err := ioutil.WriteFile("person.bin", out, 0644); err != nil {
        log.Fatalf("Write File Error: %s ", err.Error())
    }
    fmt.Println("Write Success")


    //Read from file
    in, err := ioutil.ReadFile("person.bin")
    if err != nil {
        log.Fatalf("Read File Error: %s ", err.Error())
    }
    person2 := &pb.Person{}
    err2 := proto.Unmarshal(in, person2)
    if err2 != nil {
        log.Fatalf("DeSerialization error: %s", err.Error())
    }
    fmt.Println("Read Success")

    fmt.Printf("Person2's name is %s\n", person2.GetName())
}

// Output:
// Person's name is XXX
// Write Success        
// Read Success
// Person2's name is XXX