package tests

import (
    "log"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"

    "github.com/gorilla/mux"
    "github.com/aureliomalheiros/rate-limiter/middleware"
)

func TestRateLimiter(t *testing.T) {
    r := mux.NewRouter()
    r.Use(middleware.RateLimiterMiddleware)
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World!"))
    })

    ts := httptest.NewServer(r)
    defer ts.Close()

    client := &http.Client{}

    for i := 0; i < 12; i++ {
        req, _ := http.NewRequest("GET", ts.URL, nil)
        resp, err := client.Do(req)
        if err != nil {
            t.Errorf("Error making request: %v", err)
        }
        defer resp.Body.Close()

        if i < 10 {
            if resp.StatusCode != http.StatusOK {
                t.Errorf("Expected status OK, got %v", resp.Status)
            } else {
                log.Printf("Request %d: %v", i+1, resp.Status)
            }
        } else {
            if resp.StatusCode != http.StatusTooManyRequests {
                t.Errorf("Expected status 429, got %v", resp.Status)
            } else {
                log.Printf("Request %d: %v", i+1, resp.Status)
            }
        }

        time.Sleep(time.Second)
    }
}
