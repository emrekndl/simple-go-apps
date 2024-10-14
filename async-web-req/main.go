package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/go-ping/ping"
)

func main() {
	URL_COUNT := 2

	var wg sync.WaitGroup

	urls := get_url(URL_COUNT)
	fmt.Println("The urls: ", urls)

	wg.Add(URL_COUNT)

	for i := 0; i < URL_COUNT; i++ {
		go func(i int) {
			defer wg.Done()
			pingStats, err := ping_urls(urls[i], 4, 5*time.Second)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			print_stats(urls[i], pingStats)
		}(i)
	}

	wg.Wait()

}

func ping_urls(url string, count int, timeout time.Duration) (pingStats *ping.Statistics, err error) {
	pinger, err := ping.NewPinger(url)
	if err != nil {
		return nil, err
	}
	done := make(chan bool)
	pinger.Count = count
	go func() {
		err = pinger.Run() // blocks until finished
		done <- true
		close(done)
	}()

	select {
	case <-done:
		if err != nil {
			return nil, err
		}
		return pinger.Statistics(), nil
	case <-time.After(timeout):
		return nil, fmt.Errorf("Ping timeout after %v for %s", timeout, url)
	}
}

func get_url(url_count int) []string {
	urls := make([]string, url_count)
	for i := 0; i < url_count; i++ {
		fmt.Print("Enter the url(foo.com): ")
		fmt.Scanln(&urls[i])
		if !strings.HasPrefix(urls[i], "www.") {
			urls[i] = "www." + urls[i]
		}
	}
	return urls
}

func print_stats(url string, pingStats *ping.Statistics) {
	if pingStats == nil {
		fmt.Printf("No stats available for %s\n", url)
		return
	}
	fmt.Printf("Ping results for %s:\n", url)
	fmt.Printf("Sent: %d\n", pingStats.PacketsSent)
	fmt.Printf("Received: %d\n", pingStats.PacketsRecv)
	fmt.Printf("Lost: %d\n", pingStats.PacketsRecv-pingStats.PacketsSent)
	fmt.Printf("Approximate rtt: %v\n", pingStats.AvgRtt)
	fmt.Printf("Minimum rtt: %v\n", pingStats.MinRtt)
	fmt.Printf("Maximum rtt: %v\n", pingStats.MaxRtt)
	fmt.Printf("Packet loss: %v\n", pingStats.PacketLoss)
}
