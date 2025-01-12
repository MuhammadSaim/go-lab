package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
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

// CheckDomain function is responsible for the
// check all the records on domain to verify the email
func CheckDomain(domain string) {
	// init the variables\
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

	// lookup for TXT records to find SPF records
	textRecords, err := net.LookupTXT(domain)
	// check weather there is any error
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	// loop through the spf records and find the spf version
	for _, record := range textRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	// lookup orm dmarc records

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	// check weather there is any error
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	// loop through the dmarc records and find the dmarc version
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			DmarcRecord = record
			break
		}
	}

	// output all the details
	fmt.Printf("%v -- %v -- %v -- %v -- %v -- %v", domain, hasMX, hasSPF, spfRecord, hasDMARC, DmarcRecord)
}
