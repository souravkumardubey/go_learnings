// package main

// import (
// 	"fmt"
// 	"time"
// )

// // * --- Order struct ---
// type order struct {
// 	id		   string
// 	amount     float32
// 	status     string
// 	createdAt  time.Time
// }

// // * constructor
// func newOrder(id string, amount float32, status string) *order {
// 	myOrder := order{
// 		id: id,
// 		amount: amount,
// 		status: status,
// 	}
// 	return &myOrder
// }

// // * Receiver type
// func (o *order) changeStatus(status string) {
// 	o.status = status
// }

// func (o order) getAmout() float32 {
// 	return o.amount
// }

// func main()  {

// 	// ? inline struct
// 	// language := struct {
// 	// 	name	string
// 	// 	isGood bool
// 	// } {"Golang", true}
// 	// fmt.Println(language)

// 	// myOrder := order{
// 	// 	id: "123",
// 	// 	amount: 100,
// 	// 	status: "received",
// 	// 	createdAt: time.Now(),
// 	// }
// 	myOrder := newOrder("123", 100, "pending")
// 	myOrder.changeStatus("received")
// 	fmt.Println(myOrder.status)
// 	fmt.Println(myOrder.getAmout())

// }