package main

import(
	"fmt"
	go_say_hello "github.com/safiarazl/go-say-hello"
	"time"
)

func main()  {
	fmt.Println(go_say_hello.SayHello())
	fmt.Println(time.Now())
	fmt.Println(time.Since(time.Date(2022, time.July, 18, 0, 0, 0, 0, time.UTC)))
}