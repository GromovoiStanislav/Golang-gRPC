package main

import (
    "fmt"
    "io/ioutil"
    "log"
    
    "github.com/google/uuid"
    "google.golang.org/protobuf/proto"
    "google.golang.org/protobuf/encoding/protojson"
    "google.golang.org/protobuf/types/known/timestamppb"

    "protobuf-example/src/phonebook/phonebookpb"
)

func main() {
    person1 := createPerson("Guga Zimmermann", "gugazimmermann@gmail.com","Brazil")
    person1Phone1 := createPhone("+55 47 98870-4247", phonebookpb.Person_MOBILE)
    person1Phone2 := createPhone("+55 47 98870-4247", phonebookpb.Person_WORK)
    addPhonesToPerson(person1, person1Phone1, person1Phone2)

    person2 := createPerson("Second Person", "secondperson@email.com","Brazil")
    person2Phone1 := createPhone("+XX XX XXXXX-XXXX", phonebookpb.Person_HOME)
    addPhonesToPerson(person2, person2Phone1)

    phoneBook := &phonebookpb.PhoneBook{}
    addPersonToPhoneBook(phoneBook, person1, person2)

    jsonPhoneBook := protobufToJson(phoneBook)
    fmt.Println(jsonPhoneBook)
    fmt.Println()

    writeJsonToFile("phonebook.json", jsonPhoneBook)

    emptyPhoneBookProtobuf := &phonebookpb.PhoneBook{}
    jsonToProtobuf(jsonPhoneBook, emptyPhoneBookProtobuf)
    fmt.Println(emptyPhoneBookProtobuf)
    fmt.Println()

    writeProtobufToFile("phonebook.bin", phoneBook)

    emptyPhoneBookProtobuf2 := &phonebookpb.PhoneBook{}
    readProtobufFromFile("phonebook.bin", emptyPhoneBookProtobuf2)
    fmt.Println(emptyPhoneBookProtobuf2)
    fmt.Println()
}

func createPerson(name string, email string, address string) *phonebookpb.Person {
    pe := &phonebookpb.Person{
        Id:          uuid.New().String(),
        Name:        name,
        // Email:       email,
        Phones:      []*phonebookpb.Person_PhoneNumber{},
        LastUpdated: timestamppb.Now(),
        Address:    address,
    }
    return pe
}

func createPhone(n string, t phonebookpb.Person_PhoneType) *phonebookpb.Person_PhoneNumber {
    ph := &phonebookpb.Person_PhoneNumber{
        Number: n,
        Type:   t,
    }
    return ph
}

func addPhonesToPerson(person *phonebookpb.Person, phones ...*phonebookpb.Person_PhoneNumber) {
    phoneNumber := []*phonebookpb.Person_PhoneNumber{}
    for _, phone := range phones {
        phoneNumber = append(phoneNumber, phone)
    }
    person.Phones = phoneNumber
}

func addPersonToPhoneBook(phoneBook *phonebookpb.PhoneBook, persons ...*phonebookpb.Person) {
    people := []*phonebookpb.Person{}
    
    for _, person := range persons {
        people = append(people, person)
    }

    phoneBook.People = people
}

func protobufToJson(pb proto.Message) string {
    marshaler := protojson.MarshalOptions{
        Indent:          "  ",
        UseProtoNames:   true,
        EmitUnpopulated: true,
    }

    j, err := marshaler.Marshal(pb)
    if err != nil {
        log.Println("Can't convert Protobuf to JSON", err)
    }

    // the return type of Marshal is []byte, so we'll convert to string
    return string(j)
}

func jsonToProtobuf(json string, pb proto.Message) {
    // Unmarshal accept []byte, so we'll convert json from string to []byte
    err := protojson.Unmarshal([]byte(json), pb)
    if err != nil {
        log.Println("Can't convert from JSON to Protobuf", err)
    }
}


func writeJsonToFile(filename, data string) error {
	err := ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return err
	}

	return nil
}

func writeProtobufToFile(fileName string, pb proto.Message) error {
    out, err := proto.Marshal(pb)
    if err != nil {
        log.Println("Can't serialize to bytes", err)
        return err
    }
    if err := ioutil.WriteFile(fileName, out, 0664); err != nil {
        log.Println("Can't write to file", err)
        return err
    }
    return nil
}

func readProtobufFromFile(fileName string, pb proto.Message) error {
    in, err := ioutil.ReadFile(fileName)
    if err != nil {
        log.Println("Can't read the file", err)
        return err
    }
    if err := proto.Unmarshal(in, pb); err != nil {
        log.Println("Can't unserialize the bytes", err)
        return err
    }
    return nil
}