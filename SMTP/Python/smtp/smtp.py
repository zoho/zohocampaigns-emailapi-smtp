import smtplib
from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText
import json

def prepareEmailData():
    emailData = {}
    emailData["host"] = "smtp-campaigns.zoho.com"
    emailData["port"] = 587
    emailData["username"] = "apikey"
    emailData["password"] = "1000.****************************" # replace with your access token

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
    mailserver.set_debuglevel(1)
    
    # identify ourselves to smtp gmail client
    code, response = mailserver.ehlo()
    print(f"EHLO response: {code} - {response.decode()}")
    
    # secure our email with tls encryption
    code, response = mailserver.starttls()
    print(f"STARTTLS response: {code} - {response.decode()}")
    
    # re-identify ourselves as an encrypted connection
    code, response = mailserver.ehlo()
    print(f"EHLO (after STARTTLS) response: {code} - {response.decode()}")
    
    # Authenticate using AUTH LOGIN
    code, response = mailserver.login(emailData["username"], emailData["password"])
    print(f"AUTH LOGIN response: {code} - {response}")
    
    # Send the email
    mailserver.sendmail(emailData["sender_address"], emailData["recipients"], message.as_string())

    # Close the connection
    mailserver.quit()
    print("Email sent successfully")
except Exception as e:
    print(f"Error: {e}")
