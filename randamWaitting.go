package main
import (
	"fmt"
	"math/rand"
	"time"
)

func compute(ch chan int) {
	randomValue := rand.Intn(10)
	time.Sleep(time.Duration(randomValue) * time.Second)
	ch <- randomValue
}
func main() {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go compute(ch)
	}
	c := 0
	for {
		c = c + 1
		fmt.Println("count", c)
		if c == 11 {
			break
		} else {
			val, _ := <-ch
			fmt.Println(val)
		}
	}

}
