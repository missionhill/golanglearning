package modules

import (
	"sync"
)

func bufferedChannelAsMutex(){
	var wg sync.WaitGroup
	wg.Add(2)
	mutex := make(chan struct{}, 1)
	balance := 0
	go func() {
		mutex <- struct{}{}
		balance += 3
		<- mutex
		wg.Done()
	}()
	go func() {
		mutex <- struct{}{}
		balance += 2
		<- mutex
		wg.Done()
	}()
	wg.Wait()
	AssertEqual(balance, 5)
}

func SimpleMutex(){
	var wg sync.WaitGroup
	wg.Add(2)
	mutex := sync.Mutex{}
	balance := 0
	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		balance += 3
		wg.Done()
	}()
	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		balance += 2
		wg.Done()
	}()
	wg.Wait()
	AssertEqual(balance, 5)
}

var cache map[string]int

func _initSimpleMap(){
	if cache == nil {
		cache = make(map[string]int)
		cache["test"] = 0
	}
	cache["test"]++

}

func SimpleSyncOnce(){
	var initOnce sync.Once
	var wg sync.WaitGroup
	wg.Add(10)
	for i:=0 ; i < 10 ; i++ {
		go func() {
			initOnce.Do(_initSimpleMap)
			wg.Done()
		}()
	}
	wg.Wait()
	x, ok := cache["test"]
	AssertEqual(x, 1)
	AssertTrue(ok)
}

func ConcurrencyMain(){
	bufferedChannelAsMutex()
	SimpleMutex()
	SimpleSyncOnce()
}
