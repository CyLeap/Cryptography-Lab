// Defines core data models such as Users, Transactions, and Certificates
package models

type User struct {
    ID       string
    Name     string
    Email    string
    // Additional fields for user info and credentials
}

type Transaction struct {
    ID            string
    Amount        float64
    EncryptedData []byte
    SenderID      string
    ReceiverID    string
    // Additional transaction metadata
}

type Certificate struct {
    SerialNumber string
    PublicKey    []byte
    Issuer       string
    // Additional certificate properties
}
