package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "brsr/routes"
)

func main() {
    os.Mkdir("reports", 0755)
    routes.InitRoutes()
    fmt.Println("✅ Server running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
