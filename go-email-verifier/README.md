# Email Verification Tool

This Go program verifies the presence of MX, SPF, and DMARC records for a given domain. It checks for these DNS records to help assess a domain's email configuration and security.

## Features
- Checks if a domain has MX records (Mail Exchanger)
- Verifies the presence of SPF records (Sender Policy Framework)
- Validates the existence of DMARC records (Domain-based Message Authentication, Reporting, and Conformance)

## Requirements
- Go 1.23 or later
- Internet connection for DNS lookups

## How to Use
1. Clone the repository or copy the code into a Go project.
2. Build and run the program using:
```shell
go run main.go
```
3. Enter the domain names you want to check, one per line.
4. The tool will output the results in the following format:
```shell
Domain -- HasMX -- HasSPF -- SPFRecord -- HasDMARC -- DmarcRecord
example.com -- true -- true -- v=spf1 include:_spf.example.com ~all -- true -- v=DMARC1; p=none
```

## Example Output
```shell
example.com -- true -- true -- v=spf1 include:_spf.example.com ~all -- true -- v=DMARC1; p=none
example.net -- false -- false --  -- false --
```
