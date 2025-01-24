import requests
import json

def main():
    access_token = "1000.****************************" # Replace with your access token
    payload = constructPayload()
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
        print(response.text)

    except Exception as e:
        print(f"Error: {e}")

def constructPayload():

    # Recipients
    recipient_data = []
    recipient1 = {
        "address": "sophia@zylker.com",
        "name": "Sophia Alexandri",
        "additional_data": {
            "phone": "+301234567890",
            "country": "Greece"
        },
        "merge_data": {
            "first_name": "Sophia"
        }
    }

    recipient_data.append(recipient1)

    # Content
    content = {
        "subject": "My first mail using Zoho Campaigns Email API HTTP",
        "html": "<html><body>Welcome $[first_name|Customer]$!<br>Summer Hot Savings, You Don't Want to Miss</body></html>",
        "text": "Welcome $[first_name|Customer]$! Summer Hot Savings, You Don’t Want to Miss",
        "from": {
            "address": "aron@zylker.com",
            "name": "Aron Fletcher"
        }
    }

    # Payload
    payload = {
        "campaign_name": "Summer is here",
        "recipients": recipient_data,
        "content": content
    }

    return payload

if __name__ == "__main__":
    main()
