package controllers

import (
    "encoding/json"
    "net/http"
    "brsr/models"
    "brsr/services"
)

func ReportHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
        return
    }

    var report models.Report
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&report); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Generate PDF
    pdfPath := services.GeneratePDF(report.Company, report.SectionData)
    report.PDFPath = pdfPath

    // Generate XML
    xmlPath := services.GenerateXMLReport(report.Company, report.FinancialYear, report.SectionData, report.CompletedSections)
    report.XMLPath = xmlPath

    // Upload to IPFS
    ipfsHash, _ := services.UploadToIPFS(report.SectionData)
    report.IPFSHash = ipfsHash

    // Log on blockchain
    txHash := services.LogReportToBlockchain(report.Company, ipfsHash, report.FinancialYear)
    report.TxHash = txHash

    // Save report
    services.SaveReport(report)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(report)
}
