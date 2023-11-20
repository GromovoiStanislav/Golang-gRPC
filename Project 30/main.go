package main // Замените на имя вашего пакета, если это необходимо

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"

	pb "protobufexample/pb"
)



func writeData() {
	hussein := &pb.Employee{
		Id:     1001,
		Name:   "Hussein",
		Salary: 3000,
	}

	ahmed := &pb.Employee{
		Id:     1002,
		Name:   "Ahmed",
		Salary: 9000,
	}

	rick := &pb.Employee{
		Id:     1003,
		Name:   "Rick",
		Salary: 5000,
	}

	employees := &pb.Employees{
		Employees: []*pb.Employee{hussein, ahmed, rick},
	}

	bytes, err := proto.Marshal(employees)
	if err != nil {
		log.Fatal("Error marshaling data: ", err)
	}

	err = ioutil.WriteFile("employeesbinary", bytes, 0644)
	if err != nil {
		log.Fatal("Error writing to file: ", err)
	}
}

func readData() {
	bytes, err := ioutil.ReadFile("employeesbinary")
	if err != nil {
		log.Fatal("Error reading from file: ", err)
	}

	employees := &pb.Employees{}
	err = proto.Unmarshal(bytes, employees)
	if err != nil {
		log.Fatal("Error unmarshaling data: ", err)
	}

	fmt.Println(employees.GetEmployees())
	fmt.Println()

	employeesList := employees.GetEmployees()
	for _, element := range employeesList {
		fmt.Println("Id:", element.GetId())
		fmt.Println("Name:", element.GetName())
		fmt.Println("Salary:", element.GetSalary())
		fmt.Println()
	}
}

func main() {
	writeData()
	readData()
}