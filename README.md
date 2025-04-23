# 🏢 Blockchain-based BRSR Reporting Backend

A robust, blockchain-integrated backend system for Business Responsibility and Sustainability Reporting (BRSR), designed to meet enterprise compliance and ESG transparency requirements.

---

## 📌 Project Overview

This system allows **companies** to submit their structured BRSR reports, which are:
- Converted into standardized **PDF** and **XML**
- Uploaded to **IPFS** for decentralized storage
- Logged on **Polygon blockchain** for auditability
- Optionally **signed by stakeholders**
- Accessible by **auditors** and **regulators**

---

## 🛠️ Features

✅ Multi-role access (Company, Auditor, Regulator)  
✅ PDF + XML generation  
✅ IPFS Upload  
✅ Polygon Smart Contract Integration  
✅ Signatory System (CFO, ESG Head, etc.)  
✅ JSON → IPFS → PDF/XML + Blockchain  

---

## 🔁 Input → Output Flow

1. **Input** (via Frontend or API):
```json
{
  "company": "Tata Motors Ltd",
  "financialYear": "2023-24",
  "sectionData": {
    "A1": "Yes",
    "B1": "Environment policy adopted",
    "C1_P1_Q1": "Compliant"
  },
  "completedSections": [true, true, true],
  "signatories": [
    {
      "name": "John Doe",
      "role": "CFO",
      "date": "2024-04-21",
      "signed": true
    }
  ]
}
```

2. **Output** (Returned JSON):
```json
{
  "ipfsHash": "bafybeigdexample...",
  "txHash": "0x123abc...",
  "pdfPath": "reports/Tata Motors Ltd_report.pdf",
  "xmlPath": "reports/Tata Motors Ltd_report.xml"
}
```

---

## 🚀 Running Locally

### 1. Clone and Setup
```bash
git clone https://github.com/YOUR_ORG/brsr-backend.git
cd brsr-backend
go mod tidy
```

### 2. Configure `.env`
```
PORT=8080
WEB3_STORAGE_TOKEN=your_web3_storage_token
PRIVATE_KEY=your_wallet_private_key
INFURA_URL=https://polygon-mumbai.infura.io/v3/YOUR_PROJECT_ID
CONTRACT_ADDRESS=0xYourDeployedSmartContract
```

### 3. Run the Server
```bash
go run main.go
```

---

## 📬 API Endpoint

### `POST /report/submit`
Submits a report and returns hashes and file links.

---

## 📄 Smart Contract

See [`BRSRReportRegistry.sol`](./BRSRReportRegistry.sol) for the contract deployed on Polygon Mumbai Testnet.

---

## 🧠 Architecture

- Language: Go (Golang)
- Blockchain: Polygon (Mumbai / Mainnet)
- Storage: IPFS via Web3.Storage
- PDF/XML: `gofpdf` + built-in encoding
- Auth: Role-based access control
- Deployment: Ready for Railway / Render / AWS

---

## 🔐 Roles

| Role       | Capabilities                            |
|------------|------------------------------------------|
| Company    | Submit reports, manage signatories       |
| Auditor    | View and verify submissions              |
| Regulator  | Final approval, oversight                |

---

## 📜 Whitepaper Summary

BRSR (Business Responsibility and Sustainability Reporting) is a SEBI-mandated compliance structure to ensure ESG accountability. This backend:

- Digitizes the BRSR form into JSON
- Converts it into printable, official PDF and XML formats
- Logs the IPFS hash and metadata on-chain
- Provides immutable, auditable reporting for regulators

This makes ESG reporting:
- ✅ Tamper-proof
- ✅ Transparent
- ✅ Blockchain-verifiable

---

## ✨ Future Scope

- ✅ Dashboard UI
- ✅ JWT or MetaMask login
- ✅ MongoDB/PostgreSQL integration
- ✅ Full signatory approval workflow
- ✅ Email notification system

---

## 📄 License

This project is licensed under MIT — see [LICENSE](./LICENSE)

---

## 💬 Questions or Help?

Open an [issue](https://github.com/YOUR_ORG/brsr-backend/issues) or contact us at `team@yourcompany.com`.

Built with ❤️ for ESG + Blockchain innovation.
