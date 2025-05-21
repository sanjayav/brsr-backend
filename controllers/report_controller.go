package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"brsr/models"
	"brsr/services"
	"brsr/utils"
)

// ReportHandler handles validated BRSR report submission
func ReportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read JSON body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Save temporary input
	tmpFile := "tmp_input.json"
	err = ioutil.WriteFile(tmpFile, body, 0644)
	if err != nil {
		http.Error(w, "Failed to save temp file", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tmpFile)

	// Validate JSON schema
	err = utils.ValidateBRSRJSON(tmpFile, "schemas/brsr_schema.json")
	if err != nil {
		http.Error(w, fmt.Sprintf("Schema validation failed: %v", err), http.StatusBadRequest)
		return
	}

	// Parse to model
	var report models.Report
	if err := json.Unmarshal(body, &report); err != nil {
		http.Error(w, "Failed to parse JSON into model", http.StatusBadRequest)
		return
	}

	// Generate PDF
	pdfPath, err := services.GeneratePDF(report.Company, report.SectionData)
	if err != nil {
		http.Error(w, fmt.Sprintf("PDF generation error: %v", err), http.StatusInternalServerError)
		return
	}
	report.PDFPath = pdfPath

	// Generate XML
	xmlPath, err := services.GenerateXMLReport(report.Company, report.FinancialYear, report.SectionData, report.CompletedSections)
	if err != nil {
		http.Error(w, fmt.Sprintf("XML generation error: %v", err), http.StatusInternalServerError)
		return
	}
	report.XMLPath = xmlPath

	// Upload to IPFS
	ipfsHash, err := services.UploadToIPFS(report.SectionData)
	if err != nil {
		http.Error(w, fmt.Sprintf("IPFS upload failed: %v", err), http.StatusInternalServerError)
		return
	}
	report.IPFSHash = ipfsHash

	// Log to blockchain
	txHash, err := services.LogReportToBlockchain(report.Company, ipfsHash, report.FinancialYear)
	if err != nil {
		http.Error(w, fmt.Sprintf("Blockchain logging failed: %v", err), http.StatusInternalServerError)
		return
	}
	report.TxHash = txHash

	// Save to DB or local storage
	err = services.SaveReport(report)
	if err != nil {
		http.Error(w, fmt.Sprintf("Saving report failed: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":     " Report successfully processed and submitted.",
		"company":     report.Company,
		"financial":   report.FinancialYear,
		"ipfsHash":    report.IPFSHash,
		"txHash":      report.TxHash,
		"pdf":         report.PDFPath,
		"xml":         report.XMLPath,
		"submittedAt": time.Now().Format(time.RFC3339),
	})
}
