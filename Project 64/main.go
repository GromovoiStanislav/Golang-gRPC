package main


import (
    "fmt"
    // "io/ioutil"
    "log"

    "github.com/golang/protobuf/proto"

    pb "protobuf-example/samples"
)

func main() {
    
    cmd1 := &pb.Command{
        CommandType: &pb.Command_Text{Text: "Hello, Protobuf!"},
    }
    fmt.Println(cmd1) // text:"Hello, Protobuf!"

    cmd2 := &pb.Command{
        CommandType: &pb.Command_Number{Number: 3000},
    }
    fmt.Println(cmd2) // number:3000

    env := &pb.Environment{
        Variables: map[string]string{"HOME": "/home/user", "PATH": "/usr/bin"},
    }
    fmt.Println(env) // variables:{key:"HOME"  value:"/home/user"}  variables:{key:"PATH"  value:"/usr/bin"}
    fmt.Println(env.Variables) // map[HOME:/home/user PATH:/usr/bin]
    fmt.Println(env.Variables["HOME"]) ///home/user

    var currentStatus pb.Status = pb.Status_RUNNING
    fmt.Println(currentStatus) // RUNNING


    section1 := &pb.Section{
        Title: "some title-1",
        Content: "some content",
        Type: pb.ContentType_BEST_PRACTICES,
    }
    fmt.Println(section1)// title:"some title-1"  content:"some content"  type:BEST_PRACTICES

    section2 := &pb.Section{
        Title: "some title-2",
        Content: "some content",
        Type: *pb.ContentType_INTRODUCTION.Enum(),
    }
    fmt.Println(section2) // title:"some title-2"  content:"some content"  type:INTRODUCTION


    technicalReview := &pb.TechnicalReview{
    Sections: []*pb.Section{},
    }
    sections := []*pb.Section{}
    sections = append(technicalReview.Sections, section1 )
    sections = append(technicalReview.Sections, section2 )
    technicalReview.Sections = sections
    fmt.Println(technicalReview) // sections:{title:"some title-2"  content:"some content"  type:INTRODUCTION}




    ////////////////////////
    newUser := &pb.User{
        Name:  "Alice",
        Age:   30,
        Email: "alice@example.com",
    }

    // сериализация данных
    data, err := proto.Marshal(newUser)
    if err != nil {
        log.Fatal("Marshaling error: ", err)
    }

    // десериализация данных
    newUser2 := &pb.User{}
    err = proto.Unmarshal(data, newUser2)
    if err != nil {
        log.Fatal("Unmarshaling error: ", err)
    }
    fmt.Println(newUser2.GetName(), newUser2.GetAge(), newUser2.GetEmail()) // Alice 30 alice@example.com
}

