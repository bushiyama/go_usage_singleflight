package main

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

var group singleflight.Group

func main() {
	client("hoge")
	clientSingle("hoge")

	fmt.Println("===== goroutine START =====")

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			client("hoge")
			clientSingle("hoge")
			wg.Done()
		}()
	}
	wg.Wait()
}

func clientSingle(param string) (int, error) {
	// param string only
	v, err, _ := group.Do(param, func() (interface{}, error) {
		time.Sleep(time.Second * 1)
		fmt.Println("call single!")
		if param != "" {
			return "1", nil
		}
		return nil, errors.New("hoge")
	})

	if err != nil {
		return 0, err
	}
	i, err := strconv.Atoi(v.(string))
	return i, err
}

func client(param string) (int, error) {
	time.Sleep(time.Second * 1)
	fmt.Println("call Not single!")
	if param != "" {
		return 2, nil
	}
	return 0, errors.New("hoge")
}
