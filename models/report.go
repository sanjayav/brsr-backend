package models

type Signatory struct {
    Name   string `json:"name"`
    Role   string `json:"role"` // e.g., "CFO", "Compliance Officer"
    Date   string `json:"date"`
    Signed bool   `json:"signed"`
}

type Report struct {
    Company           string                 `json:"company"`
    FinancialYear     string                 `json:"financialYear"`
    SectionData       map[string]interface{} `json:"sectionData"`
    CompletedSections []bool                 `json:"completedSections"`
    IPFSHash          string                 `json:"ipfsHash"`
    TxHash            string                 `json:"txHash"`
    PDFPath           string                 `json:"pdfPath"`
    XMLPath           string                 `json:"xmlPath"`
    Signatories       []Signatory            `json:"signatories"`
}
