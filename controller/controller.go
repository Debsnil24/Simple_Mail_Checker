package controller

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func CheckDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var SPFrecord, DMARCrecord string

	MXrecord, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if len(MXrecord) > 0 {
		hasMX = true
	}

	TXTrecord, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	for _, record := range TXTrecord {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			SPFrecord = record
			break
		}
	}

	dmarcRecord, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecord {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			DMARCrecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v\n", domain, hasMX, hasSPF, SPFrecord, hasDMARC, DMARCrecord)

}