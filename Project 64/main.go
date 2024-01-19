package main


import (
    "fmt"
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
    fmt.Println(pb.Status_STOPPED) // STOPPED
    fmt.Println(pb.Status_value["RUNNING"]) // 0
    fmt.Println(pb.Status_value["STOPPED"]) // 1
    fmt.Println(pb.Status_name[0]) // RUNNING
    fmt.Println(pb.Status_name[1]) // STOPPED


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


    technicalReview := &pb.TechnicalReview{}
    // or:
    sections := []*pb.Section{}
    sections = append(sections, section1 )
    sections = append(sections, section2 )
    technicalReview.Sections = sections
    // or:
    // technicalReview.Sections = append(technicalReview.Sections, section1 )
    // technicalReview.Sections = append(technicalReview.Sections, section1 )
    fmt.Println(technicalReview) // sections:{title:"some title-1" content:"some content" type:BEST_PRACTICES} sections:{title:"some title-1" content:"some content" type:BEST_PRACTICES}



    ////////////////////////
    newUser := &pb.User{
        Name:  "Alice",
        Age:   30,
        Email: "alice@example.com",
    }

    fmt.Println(newUser) // name:"Alice" age:30 email:"alice@example.com"

    // сериализация данных
    serialized_data , err := proto.Marshal(newUser)
    if err != nil {
        log.Fatal("Marshaling error: ", err)
    }

    // десериализация данных
    newUser2 := &pb.User{}
    err = proto.Unmarshal(serialized_data, newUser2)
    if err != nil {
        log.Fatal("Unmarshaling error: ", err)
    }
    fmt.Println(newUser2.GetName(), newUser2.GetAge(), newUser2.GetEmail()) // Alice 30 alice@example.com


}

