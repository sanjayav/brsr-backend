package routes

import (
    "net/http"
    "brsr/controllers"
)

func InitRoutes() {
    http.HandleFunc("/report/submit", controllers.ReportHandler)
}
