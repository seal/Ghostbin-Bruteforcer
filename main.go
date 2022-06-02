package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/remeh/sizedwaitgroup"
)

var (
	valid     []string
	invalid   int
	timeout   string
	url       string
	total     int
	errors    []string
	generated string
)

func main() {
	CallClear()
	fmt.Println("github.com/seal")
	fmt.Println("Linear or Randon(l/r)")
	var choice string
	fmt.Scan(&choice)
	switch strings.ToLower(choice) {
	case "l":
		linear()
	case "r":
		Random()
	}

}
func linear() {
	url = "https://pst.klgrth.io/paste/"
	fmt.Println("Please note linear does not display invalid URL's or errors due to most being invalid")
	first, second, third, fourth, fifth := 0, 0, 0, 0, 0
	fmt.Println("Start from custom number?(y/n)") // Very broken, doesn't work for characters above a value of 9 (Anything over jjjjj), will fix soon
	var start string
	fmt.Scan(&start)
	if strings.ToLower(start) == "y" {
		fmt.Println("Please enter your 5 value checked:")
		var digits string
		fmt.Scan(&digits)
		split := strings.Split(digits, "") // split[o]
		first, _ = strconv.Atoi(split[0])
		second, _ = strconv.Atoi(split[1])
		third, _ = strconv.Atoi(split[2])
		fourth, _ = strconv.Atoi(split[3])
		fifth, _ = strconv.Atoi(split[4])
		//fmt.Printf("%s%s%s%s", first, second, third, fourth)

	}
	fmt.Println("Input # of times to run")
	var number int
	fmt.Scan(&number)
	fmt.Println("Input # of threads")
	var threads int
	fmt.Scan(&threads)
	swg := sizedwaitgroup.New(threads)
	chars := map[int]string{
		0:  "a",
		1:  "b",
		2:  "c",
		3:  "d",
		4:  "e",
		5:  "f",
		6:  "g",
		7:  "h",
		8:  "i",
		9:  "j",
		10: "k",
		11: "l",
		12: "m",
		13: "n",
		14: "o",
		15: "p",
		16: "q",
		17: "r",
		18: "s",
		19: "t",
		20: "u",
		21: "v",
		22: "w",
		23: "x",
		24: "y",
		25: "z",
		26: "0",
		27: "1",
		28: "2",
		29: "3",
		30: "4",
		31: "5",
		32: "6",
		33: "7",
		34: "8",
		35: "9",
	}

	for i := 0; i < number; i++ {
		swg.Add()
		go func(i int) {
			defer swg.Done()
			var generated string
			generated = chars[fifth] + chars[fourth] + chars[third] + chars[second] + chars[first]
			if chars[fourth] == "9" {
				fifth++
				fourth = 1
			} else if chars[third] == "9" {
				fourth++
				third = 1
			} else if chars[second] == "9" {
				third++
				second = 1
			} else if chars[first] == "9" {
				second++
				first = 1
			} else {
				first++
			}
			finalurl := url + generated
			resp, err := http.DefaultClient.Get(finalurl)
			if err != nil {
				errors = append(errors, err.Error())
				return
			}
			defer resp.Body.Close()
			switch resp.Status {
			case "404 Not Found":
				invalid++
			case "200 OK":
				total++
				valid = append(valid, generated)
				fmt.Printf("ValidURL %s\n", finalurl)

			}

		}(i)
	}
	swg.Wait()
	CallClear()
	fmt.Printf("Total checked: %d\n", total)
	fmt.Printf("Invalid: %d\n", invalid)
	fmt.Printf("Valid: %s\n", valid)
	fmt.Printf("Last checked: %s\n", generated)
	fmt.Printf("Last checked number value: %d%d%d%d%d\n", fifth, fourth, third, second, first)
	fmt.Printf("Errors: %s", errors)

}
func Random() {
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
				total++
				valid = append(valid, generated)
				fmt.Printf("ValidURL %s\n", finalurl)

			}

		}(i)
	}
	swg.Wait()
	CallClear()
	fmt.Printf("Total checked: %d\n", total)
	fmt.Printf("Invalid: %d\n", invalid)
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
