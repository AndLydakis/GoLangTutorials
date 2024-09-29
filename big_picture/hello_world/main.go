package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// PT1
	//fmt.Println("Hello, world")

	// PT2
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "web services are easy with Go")
	// })

	// http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "./home.html")
	// })
	// http.ListenAndServe(":3000", nil)

	// PT3
	level := flag.String("level", "CRITICAL", "log level to filter for")
	flag.Parse()

	f, err := os.Open("./log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() //defer close until main exits

	bufReader := bufio.NewReader(f)

	for line, err := bufReader.ReadString('\n'); err == nil; line, err = bufReader.ReadString('\n') {
		if strings.Contains(line, *level) {
			fmt.Println(line)
		}

	}

}
