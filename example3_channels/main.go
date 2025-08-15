package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

const NumberOfPizzasToCook = 5

var pizzasMade, pizzasFailed, total int

type PizzaOrder struct {
	orderNumber int
	message     string
	success     bool
}
type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizzaOrder(orderNumber int) *PizzaOrder {
	orderNumber++
	var msg string
	var success bool

	if orderNumber <= NumberOfPizzasToCook {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order #%d! \n", orderNumber)

		rnd := rand.Intn(12) + 1

		fmt.Printf("Making pizza #%d. It will take %d seconds... \n", orderNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza order #%d", orderNumber)
			pizzasFailed++

		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza order #%d", orderNumber)
			pizzasFailed++

		} else {
			msg = fmt.Sprintf("Pizza order #%d, is ready!", orderNumber)
			success = true
			pizzasMade++
		}
		total++
	}

	return &PizzaOrder{
		orderNumber: orderNumber,
		message:     msg,
		success:     success,
	}
}

func produceOrders(producerOrders *Producer) {
	//keep track of which pizza we are making
	var orderNumber int

	//run forever or until we receive a quit notification
	for {
		//try to make a pizza
		currentPizzaOrder := makePizzaOrder(orderNumber)
		orderNumber = currentPizzaOrder.orderNumber
		select {
		case producerOrders.data <- *currentPizzaOrder:

			//Close ch's when consumer sends quit notification
		case quitChan := <-producerOrders.quit:
			close(producerOrders.data)
			close(quitChan)
			return

		}
	}
}

func consumeOrders(producerOrders *Producer) {
	for order := range producerOrders.data {
		if order.orderNumber <= NumberOfPizzasToCook {
			if order.success {
				color.Green(order.message)
				color.Green("Order #%d is out for delivery!", order.orderNumber)
			} else {
				color.Red(order.message)
				color.Red("The customer is really mad!")
			}
		} else {
			color.Cyan("Done making pizzas...")

			//Send close signal to ch quit
			err := producerOrders.Close()
			if err != nil {
				color.Red("*** Error closing channel!", err)
			}
		}
	}
}

func main() {
	//seed the random number generator
	//rand.Seed(time.Now().Unix())

	// print out a message
	color.Green("-------------------------------- \n")
	color.Green("The Pizzeria is open to business \n")
	color.Green("-------------------------------- \n")

	//create a producer
	producer := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}
	//run the producer in the background
	go produceOrders(producer)

	//create and run consumer
	consumeOrders(producer)

	color.Cyan("-------------------------------- \n")
	color.Cyan("Summary :")
	color.Cyan("-------------------------------- \n")
	color.Cyan("We made %d pizzas, but failed to make %d, with %d attempts in total", pizzasMade, pizzasFailed, total)

}
