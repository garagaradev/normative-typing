package main
import "fmt"

type Cat struct {
  Name string
}

type Dog struct {
  Name string
}

func main(){
  var c Cat = Cat{Name:"a cat"}
  var d Dog = Dog{Name:"a dog"} 

  fmt.Println(c.Name,"&",d.Name)
  var animal Dog = c //expected error

}
