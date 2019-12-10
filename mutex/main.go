package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func request(c chan bool) {
	resp, err := http.Get("http://google.com")
	if err != nil {
		c <- false
		return
	}
	_, _ = ioutil.ReadAll(resp.Body)
	c <- true
}

func main() {
	t1 := time.Now()
	c := make(chan bool)
	for i := 0; i < 100; i++ {
		go request(c)
	}
	for i := 0; i < 100; i++ {
		<-c
	}
	fmt.Print(time.Now().Sub(t1))

}

// 	m := SyncMap{
// 		m: make(map[int]int),
// 	}
// 	for i := 0; i < 100; i++ {
// 		wg.Add(1)
// 		go m.setValue(i, i+5)
// 	}

// 	wg.Wait()
// 	for key, value := range m.m {
// 		fmt.Printf("%d %d\n", key, value)
// 	}

// }
