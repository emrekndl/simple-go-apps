package main

import (
	"cmp"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"sync"
	"sync/atomic"
	"time"
)

func _timers() {
	timer1 := time.NewTimer(time.Second * 2)

	<-timer1.C
	fmt.Println("Timer 1 fired.")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired.")
	}()
	stop := timer2.Stop()
	if stop {
		fmt.Println("Timer 2 stopped.")
	}

	time.Sleep(time.Second * 2)
}

func _tickers() {
	ticker := time.NewTicker(time.Millisecond * 500)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at: ", t)
			}
		}
	}()

	time.Sleep(time.Millisecond * 2000)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker Stopped.")
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker ", id, "started job ", j)
		time.Sleep(time.Second)
		fmt.Println("worker ", id, "finished job ", j)
		results <- j * 2
	}
}

func _workerPools() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for k := 1; k <= numJobs; k++ {
		<-results
	}
}

func worker2(id int) {
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d finished\n", id)
}

func _waitgroups() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker2(i)
		}()
	}

	wg.Wait()
}

func _rateLimiting() {
	requests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests {
		<-limiter
		fmt.Println("Request ", req, time.Now())
	}

	// Bursting the Limiter
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("Bursty Request ", req, time.Now())
	}

}

func _atomicCounters() {
	var ops atomic.Uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				// atomic counter
				ops.Add(1)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("ops: ", ops.Load())
}

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func _mutexes() {
	cnt := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			cnt.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()

	fmt.Println(cnt.counters)
}

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key   int
	value int
	resp  chan bool
}

func _statefulGoroutines() {
	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.value
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				// response := <-read.resp
				// fmt.Printf("Read: key: %d, resp: %d\n", read.key, response)
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:   rand.Intn(5),
					value: rand.Intn(100),
					resp:  make(chan bool),
				}
				writes <- write
				<-write.resp
				// response := <-write.resp
				// fmt.Printf("Write: key: %d, val: %d, resp: %t\n", write.key, write.value, response)
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps: ", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps: ", writeOpsFinal)
}

func _sorting() {
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings: ", strs)

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints: ", ints)

	S := slices.IsSorted(ints)
	fmt.Println("Sorted: ", S)
}

func _sortingByFunctions() {
	fruits := []string{"peach", "banana", "kiwi"}

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}
	slices.SortFunc(fruits, lenCmp)
	fmt.Println("Fruits: ", fruits)

	type Person struct {
		name string
		age  int
	}
	people := []*Person{
		&Person{"Bob", 31},
		&Person{"John", 42},
		&Person{"Michael", 17},
		&Person{"Jenny", 26},
	}
	// people := []Person{
	// 	{"Bob", 31},
	// 	{"John", 42},
	// 	{"Michael", 17},
	// 	{"Jenny", 26},
	// }

	// slices.SortFunc(people, func(a, b Person) int {
	// 	return cmp.Compare(a.age, b.age)
	// })
	slices.SortFunc(people, func(a, b *Person) int {
		return cmp.Compare(a.age, b.age)
	})

	// fmt.Println("People: ", people)
	for _, v := range people {
		fmt.Println("People: ", v.name, v.age)
	}
}

func _panic() {
	panic("PANIC, a problem occurred")

	_, err := os.Create("/tmp/non-existent-file")
	if err != nil {
		panic(err)
	}
}

func createFile(p string) *os.File {
	fmt.Println("Creating file: ", p)
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("Writing to file: ", f.Name())
	fmt.Fprint(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("Closing file: ", f.Name())
	if err := f.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func _defer() {
	f := createFile("/tmp/test")
	defer closeFile(f)
	writeFile(f)
}

func mayPanic() {
	panic("A problem occurred")
}

func _recover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic, Error: ", r)
		}
	}()

	mayPanic()
	fmt.Println("After mayPanic()")
}

func main() {
	// Timers
	_timers()
	fmt.Println("-----------------")
	// Tickers
	_tickers()
	fmt.Println("-----------------")
	// Worker Pools
	_workerPools()
	fmt.Println("-----------------")
	// Wait Groups
	_waitgroups()
	fmt.Println("-----------------")
	// Rate Limiting
	_rateLimiting()
	fmt.Println("-----------------")
	// Atomic Counters
	_atomicCounters()
	fmt.Println("-----------------")
	// Mutexes
	_mutexes()
	fmt.Println("-----------------")
	// Stateful Goroutines
	_statefulGoroutines()
	fmt.Println("-----------------")
	// Sorting
	_sorting()
	fmt.Println("-----------------")
	// Sorting by Functions
	_sortingByFunctions()
	fmt.Println("-----------------")
	// Panic
	// _panic()
	fmt.Println("-----------------")
	// Defer
	_defer()
	fmt.Println("-----------------")
	// Recover
	_recover()
	fmt.Println("-----------------")

	fmt.Println("Done")
}
