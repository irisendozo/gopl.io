package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Experiment to measure the difference in running time between our potentially
// inefficient versions and the one that uses strings.Join
func main() {
	// Measure stringJoin()
	timeNow := time.Now()
	stringJoin()
	fmt.Println("stringJoin(): ", time.Since(timeNow))

	timeNow = time.Now()
	stringFor()
	fmt.Println("stringFor(): ", time.Since(timeNow))
}

func stringJoin() {
	fmt.Println(strings.Join(os.Args[0:], "\n"))
}

func stringFor() {
	for _, arg := range os.Args {
		fmt.Println(arg + "\n")
	}
}
