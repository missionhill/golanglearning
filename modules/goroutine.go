package modules

import (
	"time"
	"sync"
)

func checkSimpleChannel(){
	var chx, chy chan int
	AssertTrue(chx==nil)
	AssertTrue(chy==nil)
	chx, chy = make(chan int), make(chan int)
	AssertTrue(chx!=nil)
	AssertTrue(chy!=nil)
	z := 0
	go func() {
		z++
		AssertEqual(z, 1)
		chx<-1
		y:= <-chy
		AssertEqual(y,1)
	}()
	x := <-chx // blocking receiver assuming go func is executed afterwards
	chy<-1 // blocking sender assuming go func is executed afterwards
	z++
	AssertEqual(z, 2)
	AssertEqual(x, 1)
}

func checkClosedChannel(){
	defer func() {
		r := recover()
		AssertTrue(r!=nil)
	}()
	ch := make(chan int)
	close(ch)
	ch<-1
}

func checkPipeline1(){
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 10 ; x++ {
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		for i := 0; i < 10 ; i++ {
			x := <-naturals
			squares <- x * x
		}
	}()

	for i:=0; i < 10 ; i++ {
		AssertEqual(<- squares, i*i)
	}
}

func checkPipeline2(){
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 10 ; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break // channel was closed and drained
			}
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	i := 0
	for {
		x, ok := <- squares
		if ok {
			AssertEqual(x, i*i)
		} else {
			break
		}
		i++
	}
}

func checkPipeline3(){
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	i:=0
	for x := range squares {
		AssertEqual(x, i*i)
		i++
	}
}

func checkSendReceiveOnlyChannels(){
	channel := make(chan int)
	sendOnly := (chan<- int)(channel)
	receiveOnly := (<-chan int)(channel)
	AssertTrue(sendOnly != nil)
	AssertTrue(receiveOnly != nil)
}

func checkBufferredChannels(){
	ch := make(chan string, 3)
	AssertEqual(cap(ch), 3)
	ch<-"1"
	ch<-"2"
	ch<-"3"
	AssertEqual(<-ch, "1")
	AssertEqual(<-ch, "2")
	AssertEqual(<-ch, "3")
}

func checkParallelRunTime(){
	start := time.Now()
	for i:=0 ; i < 10 ; i++ {
		go func(){
			time.Sleep(10 * time.Millisecond)
		}()
	}
	AssertTrue(int(time.Since(start)/time.Millisecond) < 10*10)
	start = time.Now()
	for i:=0 ; i < 10 ; i++ {
			time.Sleep(10 * time.Millisecond)
	}
	AssertTrue(int(time.Since(start)/time.Millisecond) >= 10*10)
}

func checkClosureGotcha(){
	numbers :=[]int{1,2,3,4,5,6,7,8,9}
	ch1 :=make(chan int,9)
	ch2 :=make(chan int,9)
	for _, num := range numbers {
		go func() {
			ch1<- num
		}()
		go func(x int){
			ch2<- x
		}(num)
	}
	sum1 := 0
	sum2 := 0
	for i:=0 ; i<9 ; i ++ {
		sum1 += <-ch1
		sum2 += <-ch2
	}
	AssertTrue(sum1!=45)
	AssertEqual(sum2, 45)
}

func checkWaitGroup(){
	ch1 := make(chan int)
	wg := sync.WaitGroup{}
	for i:= 0 ; i < 10 ; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			ch1<-num
		}(i)
	}
	go func() {
		wg.Wait()
		close(ch1)
	}()
	sum1 := 0
	for x := range ch1 {
		sum1 += x
	}
	AssertEqual(sum1, 45)
}

func checkMultiplex(){
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			AssertEqual(x, i-1)
		case ch <- i:
		}
	}
}

func checkCloseBroadcast(){
	ch :=make(chan int)
	var wg sync.WaitGroup
	wg.Add(3)
	for i:=0 ; i < 3; i++ {
		go func() {
			_, ok := <-ch
			AssertTrue(!ok)
		}()
		wg.Done()
	}
	close(ch)
	wg.Wait()
}

func GoroutineMain(){
	checkSimpleChannel()
	checkClosedChannel()
	checkPipeline1()
	checkPipeline2()
	checkPipeline3()
	checkSendReceiveOnlyChannels()
	checkBufferredChannels()
	checkParallelRunTime()
	checkClosureGotcha()
	checkWaitGroup()
	checkMultiplex()
	checkCloseBroadcast()
}