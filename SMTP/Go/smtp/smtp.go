package main

import (
    "crypto/tls"
    "encoding/json"
    "fmt"
    "net"
    "net/smtp"
    "net/textproto"

    "github.com/jordan-wright/email"
)

func main() {
    // Prepare the email data
    e, smtpServer, auth := prepareEmailData()

    // Send the email
    if err := sendEmail(e, smtpServer, auth); err != nil {
        fmt.Println("Failed to send email:", err)
    } else {
        fmt.Println("Email sent successfully!")
    }
}

// prepareEmailData constructs and returns the email message and SMTP configuration
func prepareEmailData() (*email.Email, string, CustomAuth) {
    
    // Email sender, recipient details, and other data
    senderAddress := "aron@zylker.com"
    recipients := []string{"sophia@zylker.com"}
    subject := "My first mail using Zoho Campaigns Email API SMTP"
    metaData := map[string]interface{}{
        "campaign_name": "Summer is here",
        "recipient_data": map[string]interface{}{
            "sophia@zylker.com": map[string]interface{}{
                "name": "Sophia Alexandri",
                "additional_data": map[string]string{
                    "phone":   "+301234567890",
                    "country": "Greece",
                },
                "merge_data": map[string]string{
                    "first_name": "Sophia",
                },
            },
        },
    }

    // Email content
    htmlContent := "<html><body>Welcome $[first_name|Customer]$!<br>Summer Hot Savings, You Don't Want to Miss</body></html>"
    textContent := "Welcome $[first_name|Customer]$! Summer Hot Savings, You Don’t Want to Miss"

    // Convert metaData to JSON string
    metaDataBytes, err := json.Marshal(metaData)
    if err != nil {
        fmt.Println("Error marshaling metaData:", err)
        return nil, "", CustomAuth{}
    }
    metaDataStr := string(metaDataBytes)

    // Create the email message
    e := email.NewEmail()
    e.From = senderAddress
    e.To = recipients
    e.Subject = subject
    e.Text = []byte(textContent)
    e.HTML = []byte(htmlContent)
    e.Headers = textproto.MIMEHeader{
        "X-ZCEA-SMTP-DATA": {metaDataStr},
    }

    // SMTP server configuration
    host := "smtp-campaigns.zoho.com"
    port := "587"
    smtpServer := net.JoinHostPort(host, port)

    // Set up custom SMTP authentication
    auth := CustomAuth{
        accessToken: "1000.***************", // Replace with your access token
        host:        host,
    }

    return e, smtpServer, auth
}

// sendEmail handles the logic for sending an email using an SMTP server
func sendEmail(e *email.Email, smtpServer string, auth CustomAuth) error {
    host, _, _ := net.SplitHostPort(smtpServer)

    // Connect to the SMTP server
    conn, err := net.Dial("tcp", smtpServer)
    if err != nil {
        return fmt.Errorf("Error connecting to SMTP server: %v", err)
    }
    defer conn.Close()

    // Create a new SMTP client
    client, err := smtp.NewClient(conn, host)
    if err != nil {
        return fmt.Errorf("Error creating SMTP client: %v", err)
    }
    defer client.Close()

    // Start TLS encryption
    if err = client.StartTLS(&tls.Config{ServerName: host, InsecureSkipVerify: true}); err != nil {
        return fmt.Errorf("Error starting TLS: %v", err)
    }

    // Perform authentication
    if err = sendAuthCommand(client, auth); err != nil {
        return fmt.Errorf("error during authentication: %v", err)
    }

    // Set the sender and recipient
    if err = client.Mail(e.From); err != nil {
        return fmt.Errorf("Error setting sender: %v", err)
    }

    for _, recipient := range e.To {
        if err = client.Rcpt(recipient); err != nil {
            return fmt.Errorf("Error setting recipient: %v", err)
        }
    }

    // Get the writer to write the email data
    wc, err := client.Data()
    if err != nil {
        return fmt.Errorf("Error getting data writer: %v", err)
    }

    // Convert email to bytes and write the email data
    emailBytes, err := e.Bytes()
    if err != nil {
        return fmt.Errorf("Error converting email to bytes: %v", err)
    }

    if _, err = wc.Write(emailBytes); err != nil {
        return fmt.Errorf("Error writing email data: %v", err)
    }

    // Close the writer
    if err = wc.Close(); err != nil {
        return fmt.Errorf("Error closing data writer: %v", err)
    }

    // Read response

    // Close the SMTP client
    client.Quit()

    return nil
}

// sendAuthCommand manually sends the AUTH command to the SMTP client
func sendAuthCommand(client *smtp.Client, auth CustomAuth) error {
    // Send the AUTH command with the unencoded access token
    command := fmt.Sprintf("AUTH ACCESS_TOKEN %s", auth.accessToken)
    if err := client.Text.PrintfLine(command); err != nil {
        return fmt.Errorf("error sending AUTH command: %v", err)
    }

    // Read the server response to confirm authentication
    _, _, err := client.Text.ReadResponse(235)
    if err != nil {
        return fmt.Errorf("error reading authentication response: %v", err)
    }

    return nil
}

// CustomAuth implements the smtp.Auth interface for AUTH ACCESS_TOKEN
type CustomAuth struct {
    accessToken string
    host        string
}