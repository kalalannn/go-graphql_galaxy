package client

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const RetryTimeout = 50 * time.Millisecond

func CreateURL(host, port, path string) string {
	return fmt.Sprintf("http://%s:%s%s", host, port, path)
}

func PingServer(url string, timeout time.Duration) error {
	return pingServer(url, RetryTimeout, timeout)
}

func pingServer(url string, retryTimeout time.Duration, timeout time.Duration) error {
	client := &http.Client{
		Timeout: timeout,
	}
	for {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println("Error creating request:", err)
			time.Sleep(retryTimeout)
			continue
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Println("Server not yet available, retrying...", err)
			time.Sleep(retryTimeout)
			continue
		}

		if resp.StatusCode == http.StatusOK {
			log.Println("Server is ready!")
			resp.Body.Close()
			return nil
		}

		log.Println("Server returned:", resp.StatusCode, "Retrying...")
		resp.Body.Close()
		time.Sleep(retryTimeout)
	}
}
