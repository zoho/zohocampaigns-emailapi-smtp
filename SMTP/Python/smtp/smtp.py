import smtplib
import ssl
from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText
import json

def prepareEmailData():
    emailData = {}
    emailData["host"] = "smtp-campaigns.zoho.com"
    emailData["port"] = 587
    emailData["username"] = "apikey"
    emailData["password"] = "1000.****************************" # replace with your API key

    emailData["subject"] = "My first mail using Zoho Campaigns Email API SMTP"
    emailData["sender_address"] = "aaron@zylker.com"
    emailData["recipients"] = ["sophia@zylker.com"]
    emailData["meta_data"] = {
        "transmission_name": "Summer is here",
        "recipient_data": {
            "sophia@zylker.com": {
                "name": "Sophia Alexandri",
                "additional_data": {"phone": "+301234567890", "country": "Greece"},
                "merge_data": {"first_name": "Sophia"},
            }
        },
    }
    emailData["html_content"] = "<html><body>Welcome $[first_name|Customer]$!<br>Summer Hot Savings, You Don't Want to Miss</body></html>"
    emailData["text_content"] = "Welcome $[first_name|Customer]$! Summer Hot Savings, You Don’t Want to Miss"
    return emailData

# Create session and send mail
try:
    # Create and configure the MIME message
    emailData = prepareEmailData()
    message = MIMEMultipart("alternative")
    message["From"] = emailData["sender_address"]
    message["To"] = ", ".join(emailData["recipients"])
    message["Subject"] = emailData["subject"]
    message.add_header("X-ZCEA-SMTP-DATA", json.dumps(emailData["meta_data"]))

    # Attach both plain text and HTML parts
    part1 = MIMEText(emailData["text_content"], "plain")
    part2 = MIMEText(emailData["html_content"], "html")
    message.attach(part1)
    message.attach(part2)

    # Send the email
    mailserver = smtplib.SMTP(emailData["host"], emailData["port"])
    # set_debuglevel() is intentionally not enabled: the protocol trace includes the Base64 AUTH
    # exchange (the API key) and message data, which would leak into logs.
    # Use it only for debugging in a local setup, never in production.
    
    # identify ourselves to smtp gmail client
    code, response = mailserver.ehlo()
    print(f"EHLO response: {code} - {response.decode()}")
    
    # secure our email with tls encryption
    # Verify the server certificate and hostname so the connection cannot be
    # intercepted or downgraded before credentials are sent.
    code, response = mailserver.starttls(context=ssl.create_default_context())
    print(f"STARTTLS response: {code} - {response.decode()}")
    
    # re-identify ourselves as an encrypted connection
    code, response = mailserver.ehlo()
    print(f"EHLO (after STARTTLS) response: {code} - {response.decode()}")
    
    # Authenticate using AUTH LOGIN
    mailserver.user = emailData["username"]
    mailserver.password = emailData["password"]
    code, response = mailserver.auth("LOGIN", mailserver.auth_login, initial_response_ok=False)
    print(f"AUTH LOGIN response: {code} - {response}")
    
    # Send the email
    mailserver.sendmail(emailData["sender_address"], emailData["recipients"], message.as_string())

    # Close the connection
    mailserver.quit()
    print("Email sent successfully")
except Exception as e:
    print(f"Error: {e}")
