package main
import (
    "fmt"
)
func main() {
    var name string
    var n1,n2,n3 int
    fmt.Scan(&name)
    fmt.Scan(&n1,&n2,&n3)
    fmt.Printf("The word %s containg %d number of alphabets.",  
               name, n1) 

}
