package main

import (
	"fmt"
	"time"
)

func WriteData(intChan chan int) {
	for i := 1; i <= 10; i++ {
		intChan <- i
		time.Sleep(time.Second * 1)
		fmt.Println("writeData %v", i)
	}
	close(intChan) // 关闭 方便以后循环读取
}

func ReadData(intChan chan int, exitChan chan bool) {

	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		time.Sleep(time.Second * 1)
		fmt.Println("readData %v", v)
	}
	exitChan <- true
}

func main() {

	//	创建两个管道
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)
	go WriteData(intChan)
	go ReadData(intChan, exitChan)
	for {
		if <-exitChan {
			break
		}
	}
}
