package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Domain -- HasMX -- HasSPF -- SPFRecord -- HasDMARC -- DmarcRecord\n")

	for scanner.Scan() {
		CheckDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(fmt.Printf("Error: Could not read from input: %v\n ", err))
	}
}

// TODO: CheckDomain if valid or invalid
func CheckDomain(domain string) {
	// init the variables
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, DmarcRecord string

	// lookup for MX records
	mxRecords, err := net.LookupMX(domain)
	// check weather there is any error
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	// true the flag if there is any records
	if len(mxRecords) > 0 {
		hasMX = true
	}
}
