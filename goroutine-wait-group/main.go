package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func main() {
	fmt.Println("Starting...")
	ids := []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup

	start := time.Now()

	for _, id := range ids {
		wg.Add(1)
		go getProductByIdAPI(id, &wg)
	}

	// time.Sleep(2 * time.Second)
	wg.Wait()

	fmt.Println("Finish...", time.Since(start))
}

func getProductByIdAPI(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// fmt.Println(">>> Data Product id: ", id)
	// time.Sleep(time.Second * 1)

	url := fmt.Sprintf("https://fakestoreapi.com/products/%d", id)
	resp, err := http.Get(url)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll((resp.Body))

	if err != nil {
		return
	}

	fmt.Printf(">>> Data ProductId: %d: %s\n", id, string(body))
}
