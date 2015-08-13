package main

import (
    "encoding/json"
    "fmt"

    "log"

    "os"
)

type ConfigStruct struct {
    Host              string   `json:"host"`
    Port              int      `json:"port"`
    AnalyticsFile     string   `json:"analytics_file"`
    StaticFileVersion int      `json:"static_file_version"`
    StaticDir         string   `json:"static_dir"`
    TemplatesDir      string   `json:"templates_dir"`
    SerTcpSocketHost  string   `json:"serTcpSocketHost"`
    SerTcpSocketPort  int      `json:"serTcpSocketPort"`
    Fruits            []string `json:"fruits"`
}

type Other struct {
    SerTcpSocketHost string   `json:"serTcpSocketHost"`
    SerTcpSocketPort int      `json:"serTcpSocketPort"`
    Fruits           []string `json:"fruits"`
}

func main() {
    jsonStr := `{"host": "http://localhost:9090","port": 9090,"analytics_file": "","static_file_version": 1,"static_dir": "E:/Project/goTest/src/","templates_dir": "E:/Project/goTest/src/templates/","serTcpSocketHost": ":12340","serTcpSocketPort": 12340,"fruits": ["apple", "peach"]}`

    // json -> map
    var dat map[string]interface{}
    err := json.Unmarshal([]byte(jsonStr), &dat)
    if err!= nil {
        log.Fatal(err)
    }
    fmt.Println("json -> map")
    fmt.Println(dat)
    fmt.Println(dat["host"])
    fmt.Println("--------------------------------------------")
    // json -> struct
    var config ConfigStruct

    err =json.Unmarshal([]byte(jsonStr), &config)
    if err !=nil {
        log.Fatal(err)
    }
    fmt.Println("json -> struct")
    fmt.Println(config)
    fmt.Println(config.Host)
    fmt.Println("--------------------------------------------")
    // json -> struct 部分字段
    var part Other

    err =json.Unmarshal([]byte(jsonStr), &part)
    if err !=nil {
        log.Fatal(err)
    }
    fmt.Println("json -> struct 部分转")
    fmt.Println(part)
    fmt.Println(part.SerTcpSocketPort)
    fmt.Println("-----------------------------------------------")

    // struct -> json
    fmt.Println("struct -> json")
   fmt.Println(config)
    b, err := json.Marshal(config)
    if err !=nil {
        log.Fatal(err)
    }
    fmt.Println(string(b))
    fmt.Println("-----------------------------------------------")

    // map -> json
    fmt.Println("map -> json")
    enc := json.NewEncoder(os.Stdout)
    enc.Encode(dat)
    fmt.Println("-----------------------------------------------")

    //array -> json
    fmt.Println("array -> json")
    arr :=[]string{"hello", "apple", "python", "golang", "base", "peach", "pear"}
     lang,err :=json.Marshal(arr)
     if err !=nil{
         log.Fatal(err)
     }
    fmt.Println(string(lang))
    fmt.Println("-----------------------------------------------")
    // json -> []string
    fmt.Println("json -> []string")
    var wo []string
    err = json.Unmarshal(lang,&wo)
    if err!=nil{
        log.Fatal(err)
    }
    fmt.Println(wo)
}
