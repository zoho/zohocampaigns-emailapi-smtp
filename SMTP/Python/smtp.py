import smtplib
from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText
import json

# Prepare mail data
sender_address = "aaron@zylker.com"
recipients = ["aaron.test@gmail.com", "ea.test@zylker.com"]
subject = "My first mail using Zoho Campaigns Email API SMTP"

meta_data = {
    "campaign_name": "Summer is here",
    "recipient_data": {
        "aaron.test@gmail.com": {
            "name": "Aaron Fletcher",
            "additional_data": {"phone": "+919876543210", "country": "IN"},
            "merge_data": {"first_name": "Aaron"},
        },
        "ea.test@zylker.com": {
            "name": "EA Test",
            "additional_data": {"phone": "+919876543210", "country": "IN"},
            "merge_data": {"first_name": "EA"},
        },
    },
}

# Body
html_content = "<html><body>Welcome $[first_name|Customer]$!<br>Summer Hot Savings, You Don't Want to Miss</body></html>"
text_content = "Welcome $[first_name|Customer]$! Summer Hot Savings, You Don’t Want to Miss"

# Create session and send mail
try:
    access_token = "1000.****************************"

    # Create and configure the MIME message
    message = MIMEMultipart("alternative")
    message["From"] = sender_address
    message["To"] = ", ".join(recipients)
    message["Subject"] = subject
    message.add_header("X-ZCEA-SMTP-DATA", json.dumps(meta_data))

    # Attach both plain text and HTML parts
    part1 = MIMEText(text_content, "plain")
    part2 = MIMEText(html_content, "html")
    message.attach(part1)
    message.attach(part2)

    # Send the email
    mailserver = smtplib.SMTP('smtp.campaigns.zoho.com', 587)
    mailserver.set_debuglevel(1)
    # identify ourselves to smtp gmail client
    mailserver.ehlo()
    # secure our email with tls encryption
    mailserver.starttls()
    # re-identify ourselves as an encrypted connection
    mailserver.ehlo()
    code, response = mailserver.docmd("AUTH ACCESS_TOKEN " + access_token)
    if code != 235:
        raise smtplib.SMTPAuthenticationError(code, response)

    mailserver.sendmail(sender_address, recipients, message.as_string())

    mailserver.quit()
    print("Email sent successfully")
except Exception as e:
    print(f"Error: {e}")
