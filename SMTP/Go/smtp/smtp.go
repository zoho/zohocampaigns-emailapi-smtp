package main

import (
    "crypto/tls"
    "encoding/json"
    "fmt"
    "net"
    "net/smtp"
    "net/textproto"
    "encoding/base64"
    "github.com/jordan-wright/email"
)

func main() {
    // Prepare the email data
    e, smtpServer, username, password := prepareEmailData()

    // Send the email
    if err := sendEmail(e, smtpServer, username, password); err != nil {
        fmt.Println("Failed to send email:", err)
    }
}

// prepareEmailData constructs and returns the email message and SMTP configuration
func prepareEmailData() (*email.Email, string, string, string) {
    
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
        return nil, "", "", ""
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
    username := "apikey"
    password := "1000.*********************************" // Replace with your access token

    return e, smtpServer, username, password
}

// sendEmail handles the logic for sending an email using an SMTP server
func sendEmail(e *email.Email, smtpServer string, username string, password string) error {
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
    if err = client.StartTLS(&tls.Config{ServerName: host, InsecureSkipVerify: false}); err != nil {
        return fmt.Errorf("Error starting TLS: %v", err)
    }

    // Perform authentication
    if err = sendAuthLogin(client, username, password); err != nil {
        fmt.Errorf("Error during AUTH LOGIN: %v", err)
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
    var response string
    for {
        line, err := client.Text.ReadLine()
        if err != nil {
            return fmt.Errorf("Error %v", err)
        }
        if len(line) <= 0 {
            break
        }
        response += line
    }
    fmt.Printf("Response:\n%v\n", response)

    // Close the SMTP client
    client.Quit()

    return nil
}

// sendAuthLogin sends the AUTH LOGIN command with base64-encoded credentials.
func sendAuthLogin(client *smtp.Client, username, password string) error {
    // Send the AUTH LOGIN command.
    if err := client.Text.PrintfLine("AUTH LOGIN"); err != nil {
        return fmt.Errorf("error sending AUTH LOGIN command: %v", err)
    }

    // Read the server's response.
    code, message, err := client.Text.ReadResponse(334)
    if err != nil || code != 334 {
        return fmt.Errorf("error reading server response: %v, %s", err, message)
    }

    // Send the base64-encoded username.
    encodedUsername := base64.StdEncoding.EncodeToString([]byte(username))
    if err := client.Text.PrintfLine(encodedUsername); err != nil {
        return fmt.Errorf("error sending base64-encoded username: %v", err)
    }

    // Read the server's response.
    code, message, err = client.Text.ReadResponse(334)
    if err != nil || code != 334 {
        return fmt.Errorf("error reading server response after username: %v, %s", err, message)
    }

    // Send the base64-encoded password.
    encodedPassword := base64.StdEncoding.EncodeToString([]byte(password))
    if err := client.Text.PrintfLine(encodedPassword); err != nil {
        return fmt.Errorf("error sending base64-encoded password: %v", err)
    }

    // Read the server's response to the authentication attempt.
    code, message, err = client.Text.ReadResponse(235)
    if err != nil || code != 235 {
        return fmt.Errorf("authentication failed: %v, %s", err, message)
    }

    return nil
}
