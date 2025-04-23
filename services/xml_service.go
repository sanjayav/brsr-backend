package services

import (
    "encoding/xml"
    "fmt"
    "os"
)

type ReportXML struct {
    XMLName           xml.Name                `xml:"BRSRReport"`
    Company           string                  `xml:"Company"`
    FinancialYear     string                  `xml:"FinancialYear"`
    CompletedSections []bool                  `xml:"CompletedSections>Section"`
    SectionData       []ReportField           `xml:"SectionData>Field"`
}

type ReportField struct {
    Key   string `xml:"Key"`
    Value string `xml:"Value"`
}

func GenerateXMLReport(company string, year string, data map[string]interface{}, completed []bool) string {
    fields := []ReportField{}
    for k, v := range data {
        fields = append(fields, ReportField{Key: k, Value: fmt.Sprintf("%v", v)})
    }

    report := ReportXML{
        Company:           company,
        FinancialYear:     year,
        CompletedSections: completed,
        SectionData:       fields,
    }

    output, err := xml.MarshalIndent(report, "", "  ")
    if err != nil {
        fmt.Println("Error generating XML:", err)
        return ""
    }

    filePath := fmt.Sprintf("reports/%s_report.xml", company)
    file, err := os.Create(filePath)
    if err != nil {
        fmt.Println("Error creating XML file:", err)
        return ""
    }
    defer file.Close()

    file.Write([]byte(xml.Header))
    file.Write(output)

    return filePath
}
