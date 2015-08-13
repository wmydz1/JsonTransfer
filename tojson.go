package main
import (
    "encoding/json"
    "log"
    "fmt"

)

type Human struct {
    Name string `json:"name"`
    Age int      `json:"age"`
    Phone string  `json:"phone"`
}

type Employee struct {
    Human               `json:"human"`
    Speciality string `json:"spec"`
    Phone string     `json:"phone"`
}
type Worker struct {
    Human []Human
}
type Manager struct {
    Name string `json:"name"`
    Age int      `json:"age"`
    Phone string  `json:"phone"`
    Postion string `json:"position"`
}
type Boss struct {
    Human []Human
    Manager []Manager
}

func main() {
    bob := Employee{Human{"Bob", 34, "111-666-XXX"}, "BadBoy", "100PX"}
    b, err := json.Marshal(bob)
    if (err !=nil) {
        log.Fatal(err)
    }
    fmt.Println(bob)
    fmt.Println(string(b))
    fmt.Println("-----------------")
    data := []Human{{"Bob1", 34, "111-666-XXX"}, {"Bob2", 34, "111-666-XXX"}, {"Bob3", 34, "111-666-XXX"}}
    datajson, _ := json.Marshal(data)
    fmt.Println(string(datajson))
    fmt.Println("-----------------")

    mywork := Worker{data}
    workjson, _ := json.Marshal(mywork)
    fmt.Println(string(workjson))
    fmt.Println("-----------------")
    managers := []Manager{{"CoolBo", 100, "123000", "big worker"}, {"CoolBo", 100, "123000", "big worker"}}

    boss := Boss{data, managers}

    bossjson, _ := json.Marshal(boss)
    fmt.Println(string(bossjson))
}