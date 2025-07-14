package routes

import (
	"net/http"

	"brsr-backend/controllers"
)

// SetupRouter registers all application routes and returns the configured
// http.Handler. Additional routes can be added here as the project grows.
func SetupRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/report/submit", controllers.ReportHandler)
	return mux
}
