package parallel

import (
	"fmt"
	"testing"
	"time"
)

func TestFanout(t *testing.T) {
	values := []int{1, 2, 3, 4}
	source := make(chan int)
	go func() {
		for _, v := range values {
			source <- v
		}
	}()
	results := Split(source, 2)
	go func ()  {
		for value := range results[0] {
			fmt.Println("first goroutine", value)
		}
		
	}()

	go func ()  {
		for value := range results[1] {
			fmt.Println("second goroutine", value)
		}
		
	}()
	time.Sleep(1 * time.Second)
}
