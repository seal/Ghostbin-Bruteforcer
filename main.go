package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/remeh/sizedwaitgroup"
)

var (
	valid   []string
	invalid int
	timeout string
	url     string
	total   int
)

/*
To do:
Linear checking
Add proxy support ( create lib for connecting)
*/
func main() {
	fmt.Println("github.com/seal")
	Default()
}

func Default() {
	url = "https://pst.klgrth.io/paste/"
	fmt.Println("Input # of times to run")
	var number int
	fmt.Scan(&number)
	fmt.Println("Input # of threads")
	var threads int
	fmt.Scan(&threads)
	swg := sizedwaitgroup.New(threads)

	for i := 0; i < number; i++ {
		swg.Add()
		go func(i int) {
			defer swg.Done()
			generated := generate(5)
			finalurl := url + generated
			resp, err := http.DefaultClient.Get(finalurl)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()
			switch resp.Status {
			case "404 Not Found":
				fmt.Printf("InvalidURL: %s\n", finalurl)
				invalid++
			case "200 OK":
				valid = append(valid, generated)
				fmt.Printf("ValidURL %s\n", finalurl)

			}

		}(i)
	}
	swg.Wait()
	CallClear()
	fmt.Printf("Total checked: %d\n", total)
	fmt.Printf("Invalid: %d", invalid)
	fmt.Printf("Valid: %s", valid)
}

// This code is stolen from https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go
var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func generate(n int) string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyz0987654321")
	str := make([]rune, n)
	for i := range str {
		str[i] = chars[rand.Intn(len(chars))]
	}
	return string(str)
}
