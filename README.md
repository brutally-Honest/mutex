# Simple Go HTTP Counter Server

This is a minimal Go HTTP server that demonstrates safe vs. unsafe concurrency handling with counters.  
It exposes an endpoint `/sweet` that increments two counters:

- **Safe counter**: incremented with a `sync.Mutex` to avoid race conditions.
- **Unsafe counter**: incremented without synchronization (may race under load).

---

## Features

- Simple HTTP server written in Go.
- Demonstrates concurrent access to shared state.
- Shows difference between safe (mutex-protected) and unsafe increments.
- JSON response with both counter values.

---

## Endpoints

### GET /sweet

- Increments both the safe and unsafe counters.
- Returns JSON like:

{
"body": "Sweet consumed",
"safe": 12,
"unsafe": 12
}

---

## Running the Server

1. Clone the repository:

git clone https://github.com/your-username/simple-server.git  
cd simple-server

2. Run the server:

go run main.go

3. By default, the server runs on port 7999:

http://localhost:7999/sweet

---

## Testing Concurrency

The server was stress-tested using [`hey`](https://github.com/rakyll/hey) to simulate multiple concurrent requests.

Command used:

hey -n 200000 -c 1700 http://localhost:7999/sweet

### Results Summary

- Total requests: 200,000
- Completed: 198,146
- Errors: 754 (connection refused)
- Total time: 6.3886 seconds
- Requests/sec: 31,133.4
- Fastest response: 0.0001 secs
- Slowest response: 0.6402 secs
- Average response: 0.0524 secs

### Response Time Distribution

- 0.000 [1]
- 0.064 [160,927] | ■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
- 0.128 [29,517] | ■■■■■■■
- 0.192 [6,165] | ■■
- 0.256 [1,235]
- 0.320 [269]
- 0.384 [27]
- 0.448 [3]
- 0.512 [1]
- 0.576 [0]
- 0.640 [1]

### Latency Percentiles

- 10% in 0.0001 secs
- 25% in 0.0376 secs
- 50% in 0.0509 secs
- 75% in 0.0579 secs
- 90% in 0.1061 secs
- 95% in 0.1134 secs
- 99% in 0.1683 secs

### Status Code Distribution

- 200 → 198,146 responses

### Sample Response from curl

{
"body": "Sweet consumed",
"safe": 235,897,
"unsafe": 235,895
}

**Observations:**

- The `safe` counter remains correct under high concurrency.
- The `unsafe` counter may lag slightly due to race conditions, visible under heavy load.
- Demonstrates the difference between **mutex-protected** and **unsynchronized** shared state.

---

## Project Structure

.
├── main.go # HTTP server setup  
└── counter/  
 └── counter.go # Counter logic (safe vs unsafe)

---

## Notes

- This project is intentionally simple to demonstrate **concurrency safety** in Go.
- In real-world scenarios, prefer using **mutexes** or **atomic operations** for shared state.
- The unsafe counter is included to highlight what happens without synchronization.

---
