package main

import (
	"fmt"
	"io/ioutil"
	"log"


	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"protobuf-example/src/addressbookpb"
	"protobuf-example/src/complexpb"
	"protobuf-example/src/enumpb"
	"protobuf-example/src/simplepb"

)

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Test",
		SampleList: []int32{1, 4, 7, 8},
	}
	return &sm
}

func doEnum(ID int32) *enumpb.EnumMessage {
	en := enumpb.EnumMessage{
		Id: ID,
	}
	return &en
}

func doDummy(i int32, n string) *complexpb.DummyMessage {
	dm := complexpb.DummyMessage{
		Id:   i,
		Name: n,
	}
	return &dm
}

func doComplex(ID int32) *complexpb.ComplexMessage {
	dm1 := doDummy(23, "Dummy 23")
	cm := complexpb.ComplexMessage{
		Id:       ID,
		OneDummy: dm1,
	}
	return &cm
}

func doPerson(
	id int32,
	name string,
	email string,
	phones []*addressbookpb.Person_PhoneNumber,
) *addressbookpb.Person {
	pe := addressbookpb.Person{
		Id:          id,
		Name:        name,
		Email:       email,
		Phones:      phones,
		LastUpdated: timestamppb.Now(),
	}
	return &pe
}

func main() {
	// simple proto

	fmt.Print("\n\n*** Simple ***\n\n")
	sm := doSimple()
	sm2 := &simplepb.SimpleMessage{}
	readAndWrite("simple.bin", sm, sm2)
	fmt.Println("")
	jsonDemo(sm)

	// enum proto

	fmt.Print("\n\n*** Enum ***\n\n")
	en := doEnum(30)
	fmt.Println(en)
	en.DayOfTheWeek = enumpb.DayOfTheWeek_FRIDAY
	fmt.Println(en)
	en2 := &enumpb.EnumMessage{}
	readAndWrite("enum.bin", en, en2)
	fmt.Println("")
	jsonDemo(en)

	// complex proto

	fmt.Print("\n\n*** Complex ***\n\n")
	cm := doComplex(30)
	fmt.Println(cm)
	dmTest := doDummy(66, "Dummy 66")
	dms := []*complexpb.DummyMessage{
		doDummy(50, "Dummy  50"),
		doDummy(51, "Dummy  51"),
		dmTest,
	}
	cm.MultipleDummy = dms
	fmt.Println(cm)
	dmTest.Name = "Renamed Dummy"
	fmt.Println(cm)

	cm2 := &complexpb.ComplexMessage{}
	readAndWrite("enum.bin", cm, cm2)
	fmt.Println("")
	jsonDemo(cm)

	// address book proto
	fmt.Print("\n\n*** Address Book ***\n\n")

	phoneWork := &addressbookpb.Person_PhoneNumber{
		Number: "+55 47 98870-4247",
		Type:   addressbookpb.Person_HOME,
	}
	phoneHome := &addressbookpb.Person_PhoneNumber{
		Number: "+55 47 99985-1681",
		Type:   addressbookpb.Person_WORK,
	}
	person1 := doPerson(
		1,
		"Guga Zimmermann",
		"gugazimmermann@gmail.com",
		[]*addressbookpb.Person_PhoneNumber{
			phoneWork,
			phoneHome,
		},
	)

	phoneMobile := &addressbookpb.Person_PhoneNumber{
		Number: "+55 47 99924-0805",
		Type:   addressbookpb.Person_MOBILE,
	}
	person2 := doPerson(
		2,
		"Renata Kauling Zimmermann de Negreiros",
		"rekauling@gmail.com",
		[]*addressbookpb.Person_PhoneNumber{
			phoneMobile,
		},
	)

	var people []*addressbookpb.Person
	people = append(people, person1, person2)

	ab := &addressbookpb.AddressBook{}
	ab.People = people
	jsonDemo(ab)
}

func readAndWrite(fileName string, pb proto.Message, pb2 proto.Message) {
	writeToFile(fileName, pb)
	readFromFile(fileName, pb2)
	fmt.Println("PB from", fileName)
	fmt.Println(pb2)
}

func jsonDemo(pb proto.Message) {
	pbAsJSONString := toJson(pb)
	fmt.Println("PB to JSON:")
	fmt.Println(string(pbAsJSONString))
	fromJson(pbAsJSONString, pb)
	fmt.Println("PB from JSON:", pb)
}

func writeToFile(fileName string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}
	if err := ioutil.WriteFile(fileName, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written to:", fileName)
	return nil
}

func readFromFile(fileName string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Can't read the file", err)
		return err
	}

	if err := proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Can't unserialize the bytes", err)
		return err
	}

	return nil
}

func toJson(pb proto.Message) []byte {
	marshaler := protojson.MarshalOptions{
		Indent:          "  ",
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}
	s, err := marshaler.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
	}
	return s
}

func fromJson(in []byte, pb proto.Message) {
	err := protojson.Unmarshal(in, pb)
	if err != nil {
		log.Fatalln("Can't convert from JSON to PB", err)
	}
}