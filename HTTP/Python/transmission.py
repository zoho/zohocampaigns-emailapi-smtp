import requests
import json

def main():
    access_token = "1000.****************************"

    # Recipients
    recipient_data = []
    recipient1 = {
        "address": "lucy@example.campaigns.zoho.com",
        "name": "Aaron Fletcher",
        "additional_data": {
            "phone": "+919876543210",
            "country": "IN"
        },
        "merge_data": {
            "first_name": "Aaron"
        }
    }

    recipient_data.append(recipient1)

    # Content
    content = {
        "subject": "My first mail using Zoho Campaigns Email API HTTP",
        "html": "<html><body>Welcome $[first_name|Customer]$!<br>Summer Hot Savings, You Don't Want to Miss</body></html>",
        "text": "Welcome $[first_name|Customer]$! Summer Hot Savings, You Don’t Want to Miss",
        "from": {
            "address": "aron@marketing.campaigns.zoho.com",
            "name": "Aron Fletcher"
        }
    }

    # Payload
    payload = {
        "campaign_name": "hello_customer",
        "recipients": recipient_data,
        "content": content
    }

    try:
        url = "https://campaigns.zoho.com/emailapi/v1/transmission"
        headers = {
            "Content-Type": "application/json",
            "Accept": "application/json",
            "Authorization": f"Zoho-oauthtoken {access_token}"
        }

        response = requests.post(url, headers=headers, data=json.dumps(payload), timeout=60)

        # Check the status code of the response
        status = response.status_code
        print("Status: ", status)

        # Read the response
        if response.ok:
            json_response = response.json()
            print("Response:", json.dumps(json_response, indent=2))
        else:
            print(f"Error: {response.status_code}")
            print(response.text)

    except Exception as e:
        print(f"Error: {e}")

if __name__ == "__main__":
    main()
