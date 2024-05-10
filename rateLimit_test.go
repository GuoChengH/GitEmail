package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func TestNoTokenRateLimit(t *testing.T) {

	var count = 1
	var resp *http.Response
	var err error
	for i := 0; i < count; i++ {
		resp, err = http.Get("https://api.github.com/users?since=100&per_page=100")

		if err != nil {
			log.Fatal(err)
		}
		// 检查速率限制信息
		remaining := resp.Header.Get("X-RateLimit-Remaining")
		reset := resp.Header.Get("X-RateLimit-Reset")
		resetInt, _ := strconv.ParseInt(reset, 10, 64)
		fmt.Println("Remaining Requests:", remaining)
		fmt.Println("Reset:", time.Unix(resetInt, 0))

	}

}
func TestWithTokenRateLimit(t *testing.T) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/users?since=100&per_page=10", nil)
	if err != nil {
		return
	}

	req.Header.Set("Authorization", "token "+"token")
	var count = 100
	for i := 0; i < count; i++ {
		resp, _ := client.Do(req)
		// 检查速率限制信息
		remaining := resp.Header.Get("X-RateLimit-Remaining")
		reset := resp.Header.Get("X-RateLimit-Reset")
		resetInt, _ := strconv.ParseInt(reset, 10, 64)
		fmt.Println("Remaining Requests:", remaining)
		fmt.Println("Reset:", time.Unix(resetInt, 0))
	}
}
