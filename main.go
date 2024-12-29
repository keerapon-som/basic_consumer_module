package main

import (
	"fmt"
	"kafyim/consumer"
	"time"
)

func main() {
	// var workerza consumer.ConsumerService
	consumer.NewConsumer().ConsumerType("test").Handler(Testeiei).Open()
	// workerza.Close()
	for {
		fmt.Println("Just Sleep")
		time.Sleep(10 * time.Second)
	}

}

func Testeiei(c consumer.Consumer) {
	fmt.Println("Hello Test eiei")
	fmt.Println(c)
}
