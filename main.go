package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-server/counter"
	"time"
)

type Message struct {
	Body   string `json:"body"`
	Safe   int    `json:"safe"`
	Unsafe int    `json:"unsafe"`
}

func main() {
	ctr := counter.GetCounter()

	http.HandleFunc("/sweet", func(w http.ResponseWriter, r *http.Request) {
		ctr.IncrementSafe()
		ctr.IncrementUnsafe()
		now := time.Now().Format(time.RFC3339Nano)

		fmt.Println(now, "Endpoint Hit", r.Method)
		msg := Message{
			Body:   "Sweet consumed",
			Safe:   ctr.GetCountSafe(),
			Unsafe: ctr.GetCountUnsafe(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(msg)
		if err != nil {
			fmt.Println("Error:", err)
		}
	})
	port := 7999
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("server running on port %d", port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Server error:", err)
		panic("Stopping Server")
	}
}
