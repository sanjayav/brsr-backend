package services

import (
    "fmt"
    "brsr/models"
)

var reports []models.Report

func SaveReport(r models.Report) {
    reports = append(reports, r)
    fmt.Println("Report saved for:", r.Company)
}

func GetAllReports() []models.Report {
    return reports
}
