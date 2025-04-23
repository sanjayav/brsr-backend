# BRSR Backend for Multi-Company Reporting

## Features

- Accepts BRSR report data via JSON POST
- Generates a PDF report for each company
- Uploads data to simulated IPFS
- Logs report hash on a simulated blockchain
- Returns IPFS hash, tx hash, and PDF path

## API Endpoint

### POST /report/submit

**Request Body:**
```json
{
  "company": "Example Corp",
  "financialYear": "2023-24",
  "sectionData": {
    "A1": "Answer A1",
    "B1": "Answer B1",
    ...
  },
  "completedSections": [true, true, false]
}
```

**Response:**
```json
{
  "ipfsHash": "bafyfakeipfshash12345",
  "txHash": "0xFAKEBLOCKCHAINTX12345",
  "pdfReportPath": "reports/Example Corp_report.pdf"
}
```

## Run the Server
```bash
go run main.go
```
