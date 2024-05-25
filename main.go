package main

import (
	"fmt"

	"github.com/lipesalin/go-my-own-queue/queue"

)

func main() {
	var intQueue queue.Queue = queue.Queue{}

	intQueue.Enqueue(100)
	intQueue.Enqueue(200)
	intQueue.Dequeue()                          // front is 100 and remove him
	fmt.Println("Front is: ", intQueue.Front()) // 200
	intQueue.Enqueue(300)
	intQueue.Dequeue()                          // front is 200 and remove him
	fmt.Println("Front is: ", intQueue.Front()) // 300
}
