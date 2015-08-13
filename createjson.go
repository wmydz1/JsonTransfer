package main
import (

    "encoding/json"
    "fmt"
)

type Fruits map[string]int
type Vegetables map[string]int

type Data struct {
    Fruits Fruits
    Veggies Vegetables
}
type PayLoad struct {
    Stuff Data
}
func main() {


    fruits := make(map[string]int)
    fruits["Apples"] =25
    fruits["Orange"] =11
    b, _ := json.Marshal(fruits)
    //{"Apples":25,"Orange":11}
    fmt.Println(string(b))

    vegetables := make(map[string]int)
    vegetables["Carrots"] =21
    vegetables["Peppers"] =5

    d := Data{fruits, vegetables}
    data, _ := json.Marshal(d)
    //{"Fruits":{"Apples":25,"Orange":11},"Veggies":{"Carrots":21,"Peppers":5}}
    fmt.Println(string(data))

    p := PayLoad{d}

    pjson, _ := json.Marshal(p)
    //{"Stuff":{"Fruits":{"Apples":25,"Orange":11},"Veggies":{"Carrots":21,"Peppers":5}}}
    fmt.Println(string(pjson))


}
