package main

import (
    "crypto/tls"
    "encoding/base64"
    "encoding/json"
    "fmt"
    "net"
    "net/smtp"
    "net/textproto"
    "strings"

    "github.com/jordan-wright/email"
)

func main() {
    // Prepare mail data
    senderAddress := "aaron@zylker.com"
    recipients := []string{"aaron.test@gmail.com", "ea.test@zylker.com"}
    subject := "My first mail using Zoho Campaigns Email API SMTP"

    metaData := map[string]interface{}{
        "campaign_name": "Summer is here",
        "recipient_data": map[string]interface{}{
            "aaron.test@gmail.com": map[string]interface{}{
                "name": "Aaron Fletcher",
                "additional_data": map[string]string{
                    "phone":   "+919876543210",
                    "country": "IN",
                },
                "merge_data": map[string]string{
                    "first_name": "Aaron",
                },
            },
            "ea.test@zylker.com": map[string]interface{}{
                "name": "EA Test",
                "additional_data": map[string]string{
                    "phone":   "+919876543210",
                    "country": "IN",
                },
                "merge_data": map[string]string{
                    "first_name": "EA",
                },
            },
        },
    }

    htmlContent := "<html><body>Welcome $[first_name|Customer]$!<br>Summer Hot Savings, You Don't Want to Miss</body></html>"
    textContent := "Welcome $[first_name|Customer]$! Summer Hot Savings, You Don’t Want to Miss"

    // Convert metaData to JSON string
    metaDataBytes, err := json.Marshal(metaData)
    if err != nil {
        fmt.Println("Error marshaling metaData:", err)
        return
    }
    metaDataStr := string(metaDataBytes)

    // Create email message
    e := email.NewEmail()
    e.From = senderAddress
    e.To = recipients
    e.Subject = subject
    e.Text = []byte(textContent)
    e.HTML = []byte(htmlContent)
    e.Headers = textproto.MIMEHeader{
        "X-ZCEA-SMTP-DATA": {metaDataStr},
    }

    // Set up SMTP client
    host := "campaigns.zoho.com"
    port := "587"
    address := net.JoinHostPort(host, port)

    auth := CustomAuth{
        accessToken: "1000.****************************", // Replace with your access token
        host: host,
    }

    // Connect to the SMTP server
    conn, err := net.Dial("tcp", address)
    if err != nil {
        fmt.Println("Error connecting to SMTP server:", err)
        return
    }
    defer conn.Close()

    // Create a new SMTP client
    client, err := smtp.NewClient(conn, host)
    if err != nil {
        fmt.Println("Error creating SMTP client:", err)
        return
    }
    defer client.Close()

    // Start TLS encryption
    if err = client.StartTLS(&tls.Config{ServerName: host, InsecureSkipVerify: true}); err != nil {
        fmt.Println("Error starting TLS:", err)
        return
    }

    // Authenticate with the server
    if err = client.Auth(auth); err != nil {
        fmt.Println("Error authenticating:", err)
        return
    }

    // Set the sender and recipient first
    if err = client.Mail(senderAddress); err != nil {
        fmt.Println("Error setting sender:", err)
        return
    }

    for _, recipient := range recipients {
        if err = client.Rcpt(recipient); err != nil {
            fmt.Println("Error setting recipient:", err)
            return
        }
    }

    // Get the writer to write the email data
    wc, err := client.Data()
    if err != nil {
        fmt.Println("Error getting data writer:", err)
        return
    }

    // Convert email to bytes
    emailBytes, err := e.Bytes()
    if err != nil {
        fmt.Println("Error converting email to bytes:", err)
        return
    }

    // Write the email data
    if _, err = wc.Write(emailBytes); err != nil {
        fmt.Println("Error writing email data:", err)
        return
    }

    // Close the writer
    if err = wc.Close(); err != nil {
        fmt.Println("Error closing data writer:", err)
        return
    }

    // Close the SMTP client
    if err = client.Quit(); err != nil {
        fmt.Println("Error closing SMTP client:", err)
    } else {
        fmt.Println("Email sent successfully!")
    }
}

// CustomAuth implements the smtp.Auth interface for AUTH ACCESS_TOKEN
type CustomAuth struct {
    accessToken string
    host        string
}

// Start begins the authentication process.
func (a CustomAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
    if !strings.HasSuffix(server.Name, a.host) {
        return "", nil, fmt.Errorf("wrong host name")
    }
    return "ACCESS_TOKEN", []byte(base64.StdEncoding.EncodeToString([]byte("\x00"+a.accessToken))), nil
}

// Next continues the authentication process.
func (a CustomAuth) Next(fromServer []byte, more bool) ([]byte, error) {
    if more {
        return nil, nil
    }
    return nil, nil
}
