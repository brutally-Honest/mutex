package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Counter struct {
	Count int
	mu    sync.Mutex
}

func (ctr *Counter) increment() {
	ctr.mu.Lock()
	ctr.Count++
	ctr.mu.Unlock()
}

func (ctr *Counter) GetCount() int {
	ctr.mu.Lock()
	defer ctr.mu.Unlock()
	return ctr.Count
}

func (ctr *Counter) GetCountWithoutLock() int {
	return ctr.Count
}

var ctr = Counter{Count: 0}

func main() {
	http.HandleFunc("/sweet", testHandler)
	err := http.ListenAndServe(":7999", nil)
	if err != nil {
		fmt.Println("Server error:", err)
		panic("Stopping Server")
	}
}

type Message struct {
	Body  string `json:"body"`
	Count int    `json:"count"`
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	ctr.increment()
	fmt.Println("Api Hit", ctr.GetCount())
	msg := Message{
		Body:  "Sweet consumed",
		Count: ctr.GetCount(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(msg)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
