package main

import (
	"fmt"
	"testing"
	"time"
)

/*func TestChannel(t *testing.T) {
	message := make(chan string)
	go func() {
		message <- "TestChannel success"
	}()
	msg := <- message
	fmt.Println(msg)
}*/
func TestChannelSend(t *testing.T) {
	ch := make(chan string)

	go func() {
		ch <- "goroutine send"
	}()
	ch1 := <-ch
	fmt.Println(ch1)
}
func TestRecyclingChannel(t *testing.T) {
	ch := make(chan interface{})
	go func() {
		for i := 0; i <= 3; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()
	for msg := range ch {
		fmt.Println(msg)
		if msg == 3 {
			break
		}
	}
}
func TestChannelBuffer(t *testing.T) {
	message := make(chan string, 2) //两个缓冲
	message <- "test channl buffer 1"
	message <- "test channel buffer 2"
	fmt.Println(<-message)
	fmt.Println(<-message)
}

//通道阻塞同步
func worker(done chan bool) {
	fmt.Println("working")
	time.Sleep(time.Second)
	fmt.Println("done")

	//结束阻塞
	done <- true
}
func TestChannelSync(t *testing.T) {
	done := make(chan bool)
	go worker(done)

	<-done
}

//TestProAndCus
func Printer(c chan int) {
	for {
		msg := <-c
		if msg == 0 {
			break
		}

		fmt.Println(msg)
	}
	c <- -1
}
func TestProAndCus(t *testing.T) {

	c := make(chan int)
	go Printer(c)

	for i := 1; i < 10; i++ {
		c <- i
	}

	c <- 0

	<-c
}

//在goroutine执行前程序结束运行
func TestGorou(t *testing.T) {
	go println("goroutine message")
	println("main function message")
}

//TODO channel 中值channel  make(chan <- chan bool)

// muliti channel receive
func getMessage(order string, delay time.Duration) chan string {
	c := make(chan string)
	go func() {
		for i := 1; i < 3; i++ {
			c <- fmt.Sprintf("%s %d", order, i)
			time.Sleep(delay)
		}
	}()
	return c
}
func TestMultiChannel(t *testing.T) {
	c1 := getMessage("first", 30)
	c2 := getMessage("second", 150)
	c3 := getMessage("third", 100)

	for i := 1; i < 3; i++ {
		fmt.Println(<-c1)
		fmt.Println(<-c2)
		fmt.Println(<-c3)
	}
}
