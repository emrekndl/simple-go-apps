package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func _commadLineArguments() {
	// Command-line arguments.
	// os.Args[0] is the name of the program.
	// go build main.go
	// ./main.go a b c d
	argWithProg := os.Args
	argWithoutProg := os.Args[1:]
	// arg3 := os.Args[3]

	fmt.Println("Args with prog:", argWithProg)
	fmt.Println("Args without prog:", argWithoutProg)
	// fmt.Println("Arg 3:", arg3)
	fmt.Println()

	// Command-line flags.
	// wordPtr := flag.String("word", "foo", "a string")
	// numbPtr := flag.Int("numb", 42, "an int")
	// boolPtr := flag.Bool("fork", false, "a bool")
	//
	// var svar string
	// flag.StringVar(&svar, "svar", "bar", "a string var")
	//
	// flag.Parse()
	//
	// fmt.Println("word: ", *wordPtr)
	// fmt.Println("numb: ", *numbPtr)
	// fmt.Println("fork: ", *boolPtr)
	// fmt.Println("svar: ", svar)
	// fmt.Println("tail: ", flag.Args())
	// fmt.Println()

	// Command-line Subcommands.
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	if len(os.Args) < 2 {
		fmt.Println("Expected foo or bar subcommand!")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("Subcommand: foo")
		fmt.Println("foo enable: ", *fooEnable)
		fmt.Println("foo name: ", *fooName)
		fmt.Println("foo tail: ", fooCmd.Args())

	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("Subcommand: bar")
		fmt.Println("bar level: ", *barLevel)
		fmt.Println("bar tail: ", barCmd.Args())
	default:
		fmt.Println("Unknown subcommand!")
		os.Exit(1)
	}
}

func _environmentVariables() {
	// BAR=2 go run main.go

	os.Setenv("FOO", "1")
	fmt.Println("FOO: ", os.Getenv("FOO"))
	fmt.Println("BAR: ", os.Getenv("BAR"))
	fmt.Println()

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
		fmt.Println(pair[1])
		fmt.Println()
	}

}

func _logging() {
	log.Println("Standard logger.")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("Standard logger flags with microseconds.")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Standard logger file/line.")

	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("From my:log.")

	mylog.SetPrefix("ohmy:")
	mylog.Println("From ohmy:log.")

	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)
	buflog.Println("From buf:log. Hello!")
	fmt.Print("From buflog:", buf.String())

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("Hi there! (INFO)")

	myslog.Info("hello again", "key", "val", "age", 25)
}

func _httpClient() {
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status: ", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n", req.URL.Path[1:])
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, header := range req.Header {
		for _, h := range header {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func _httpServer() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	fmt.Println("Listening on port 8090")
	http.ListenAndServe(":8090", nil)
}

func helloContext(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	fmt.Println("Server: HelloContext handler started.")
	defer fmt.Println("Server: HelloContext handler ended.")

	select {
	case <-time.After(time.Second * 10):
		fmt.Fprintf(w, "Hello, %s!\n", req.URL.Path[1:])
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("Server: ", err)

		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func _context() {
	http.HandleFunc("/helloContext", helloContext)
	fmt.Println("Listening on port 8099")
	http.ListenAndServe(":8099", nil)
}

func _spawningProcesses() {
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("Failed executing date -x: ", err)
		case *exec.ExitError:
			fmt.Println("Command exit rc = ", e.ExitCode())
		default:
			panic(err)
		}
	}

	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}

func _execingProcesses() {
	binaryPath, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binaryPath, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

func _signals() {
	signalChan := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signalChan
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("Awaiting signal.")
	<-done
	fmt.Println("Exiting.")
}

func main() {
	// Command-line arguments.
	// _commadLineArguments()
	// Environment variables.
	_environmentVariables()
	fmt.Println("-------------")
	// Logging.
	_logging()
	fmt.Println("-------------")
	// HTTP client.
	_httpClient()
	fmt.Println("-------------")
	// HTTP server.
	// _httpServer()
	fmt.Println("-------------")
	// Context.
	// _context()
	fmt.Println("-------------")
	// Spawning processes.
	_spawningProcesses()
	fmt.Println("-------------")
	// Execing processes.
	_execingProcesses()
	fmt.Println("-------------")
	// Signals.
	// _signals()
	fmt.Println("-------------")
	// os.Exit() (0: succes end status, 1-127: different error status)
	// defers will not be run when using os.Exit.
}
