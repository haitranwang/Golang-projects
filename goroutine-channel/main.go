package main

import (
	"fmt"
	"time"
)

type Message struct {
	OrderId string
	Title   string
	Price   int
}

func buyTicket(channel chan<- Message, orders []Message) {
	for _, order := range orders {
		time.Sleep(time.Second * 1)
		fmt.Printf("Pub:::%s\n", &order.OrderId)
		channel <- order // send order message to channel
	}
	fmt.Println("Buying ticket...")
	close(channel)

}

func cancelTicket(channel chan<- string, cancelOrder []string) {
	for _, orderId := range cancelOrder {
		time.Sleep(time.Second * 1)
		fmt.Printf("Cancel ticket:::%s\n", &orderId)
		channel <- orderId // send order message to channel
	}
	fmt.Println("Canceling ticket...")
	close(channel)
}

func handlerOrder(orderId string) {

}

func main() {
	buyChannel := make(chan Message)
	cancelChannel := make(chan string)

	// Simulator
	buyOrders := []Message{
		{OrderId: "Order-01", Title: "Title-01", Price: 30},
		{OrderId: "Order-02", Title: "Title-02", Price: 40},
		{OrderId: "Order-03", Title: "Title-03", Price: 50},
	}

	cancelOrders := []string{"Order-01", "Order-03"}

	go buyTicket(buyChannel, buyOrders)
	go cancelTicket(cancelChannel, cancelOrders)

	time.Sleep(time.Second * 10)
	fmt.Printf("End buying and canceling..")
}

// // pub/sub use channel and go routines

// type Message struct {
// 	OrderId string
// 	Title   string
// 	Price   int
// }

// func publisher(channel chan<- Message, orders []Message) {

// 	for _, order := range orders {
// 		fmt.Printf("Pub::: %s\n", order.OrderId)
// 		channel <- order
// 		time.Sleep(time.Second * 1)
// 	}

// 	close(channel)

// }

// func subscriber(channel <-chan Message, username string) {
// 	for msg := range channel {
// 		fmt.Printf("userName: %s ::Order %s ::Title %s:: Price::%d\n", username, msg.OrderId, msg.Title, msg.Price)
// 		time.Sleep(time.Second * 1)
// 	}
// }

// func main() {
// 	// 1. channel order
// 	orderChannel := make(chan Message)

// 	// 2. Simulate orders
// 	order := []Message{
// 		{OrderId: "Order-01", Title: "Title-01", Price: 30},
// 		{OrderId: "Order-02", Title: "Title-02", Price: 40},
// 		{OrderId: "Order-03", Title: "Title-03", Price: 50},
// 	}

// 	// send order to pub
// 	go publisher(orderChannel, order)
// 	go subscriber(orderChannel, "Hai Tran Quang")

// 	// sleep
// 	time.Sleep(time.Second * 5)
// 	fmt.Sprintln("End pub sub..")
// }

// type Course struct {
// 	Title string
// 	Price int
// }

// func main() {
// 	ch := make(chan Course)

// 	go func() {
// 		course := Course{
// 			Title: "Go Channel",
// 			Price: 30,
// 		}

// 		ch <- course // Transfer data to channel
// 	}()

// 	c := <-ch // Receive data from channel

// 	fmt.Printf("Received course: title %s, price: %d", c.Title, c.Price)
// }
