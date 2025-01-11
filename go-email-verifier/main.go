package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Domain -- HasMX -- HasSPF -- SPRecord -- HasDMARC -- DmarcRecord\n")

	for scanner.Scan() {
		CheckDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(fmt.Printf("Error: Could not read from input: %v\n ", err))
	}
}

// TODO: CheckDomain if valid or invalid
func CheckDomain(domain string) {
}
