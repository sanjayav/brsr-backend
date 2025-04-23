package models

type User struct {
    ID       string `json:"id"`
    Name     string `json:"name"`
    Wallet   string `json:"wallet"`
    Role     string `json:"role"`     // "company", "auditor", "regulator"
    Password string `json:"password"` // plaintext for demo, hash in production
}
