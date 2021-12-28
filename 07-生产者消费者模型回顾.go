package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(out chan<- int, idx int) {
	for i := 0; i < 50; i++ {
		num := rand.Intn(800)
		fmt.Printf("生产者%dth，生产：%d\n", idx, num)
		out <- num
	}
	close(out)
}

func consumer(in <-chan int, idx int) {
	for num := range in {
		fmt.Printf("-----消费者%dth，消费：%d\n", idx, num)
	}
}

//问题：channel生产者和消费者并没有体现出FIFO的原则，原因在于取出数据要打印的时候发生了cpu切换
func main() {
	product := make(chan int)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		go producer(product, i+1) // 5 生产者
	}
	for i := 0; i < 5; i++ {
		go consumer(product, i+1) // 5 个消费者
	}
	for {

	}
}
