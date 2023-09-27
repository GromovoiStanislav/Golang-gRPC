// Этот код будет в файле mydata.pb.go

package main // Замените на имя вашего пакета, если это необходимо

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

type MyData struct {
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=Age,proto3" json:"Age,omitempty"`
}

func (m *MyData) Reset()         { *m = MyData{} }
func (m *MyData) String() string { return proto.CompactTextString(m) }
func (*MyData) ProtoMessage()    {}

func main() {
	// Создаем объект данных
	data := &MyData{
		Name: "John",
		Age:  30,
	}

	// Сериализуем данные в бинарный протобуфер
	buffer, err := proto.Marshal(data)
	if err != nil {
		log.Fatal("Ошибка при сериализации:", err)
	}

	// Выводим бинарный протобуфер (это можно отправить или сохранить)
	fmt.Println(buffer)

	// Десериализуем бинарный протобуфер обратно в объект
	decodedData := &MyData{}
	err = proto.Unmarshal(buffer, decodedData)
	if err != nil {
		log.Fatal("Ошибка при десериализации:", err)
	}

	// Выводим десериализованный объект
	fmt.Printf("%+v\n", decodedData)
}
