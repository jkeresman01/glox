package main

//NOTE: defined on pages 40 - 41

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var hadError = false

func main() {
    args := os.Args[1:]

    if len(args) > 1 {
        fmt.Println("Usage: glox [script]")
		os.Exit(64) //NOTE: 64 -> Command line usage error, as defined in <sysexits.h> on BSD systems
    } else if len(args) == 1 {
        runFile(args[0])
    } else {
        runPrompt()
    }
}

func runFile(path string) {
    bytes, err := ioutil.ReadFile(path)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Could not read file: %v\n", err)
        os.Exit(74)
    }
    run(string(bytes))
}

func runPrompt() {
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("> ")
        if !scanner.Scan() {
            break
        }
        line := scanner.Text()
        run(line)
    }
}

func run(source string) {
	scaner := NewScanner(source)
	tokens := scaner.ScanTokens()

	for _, token := range tokens {
		fmt.Println(token)
	}

    fmt.Println("You entered:", source)
}

func report(line int, where,  message string) {
	fmt.Printf("[line %d] Error%s: %s\n", line, where, message)
	hadError = true
}

func errorAtLine(line int, message string) {
    report(line, "", message)
}

