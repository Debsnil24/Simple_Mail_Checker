package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"github.com/Debsnil24/Simple_Mail_Checker.git/controller"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain,hasMX,hasSPF,SPFrecord,hasDMARC,DMARCrecord\n")

	for scanner.Scan() {
		controller.CheckDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: Couldn't read input: %v\n", err)
	}
}


