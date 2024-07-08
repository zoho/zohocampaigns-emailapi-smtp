const axios = require('axios');

async function main() {
    const accessToken = "1000.****************************";

    // Recipients
    const recipientData = [{
        address: "lucy@example.campaigns.zoho.com",
        name: "Aaron Fletcher",
        additional_data: {
            phone: "+919876543210",
            country: "IN"
        },
        merge_data: {
            first_name: "Aaron"
        }
    }];

    // Content
    const content = {
        subject: "My first mail using Zoho Campaigns Email API HTTP",
        html: "<html><body>Welcome $[first_name|Customer]$!<br>Summer Hot Savings, You Don't Want to Miss</body></html>",
        text: "Welcome $[first_name|Customer]$! Summer Hot Savings, You Don’t Want to Miss",
        from: {
            address: "aron@marketing.campaigns.zoho.com",
            name: "Aron Fletcher"
        }
    };

    // Payload
    const payload = {
        campaign_name: "hello_customer",
        recipients: recipientData,
        content: content
    };

    try {
        const url = "https://campaigns.zoho.com/emailapi/v1/transmission";
        const headers = {
            "Content-Type": "application/json",
            "Accept": "application/json",
            "Authorization": `Zoho-oauthtoken ${accessToken}`
        };

        const response = await axios.post(url, payload, { headers, timeout: 60000 });

        if (response.status === 200) {
            console.log("Response:", JSON.stringify(response.data, null, 2));
        } else {
            console.log(`Error: ${response.status}`);
            console.log(response.data);
        }
    } catch (error) {
        console.error(`Exception occurred: ${error}`);
        if (error.response) {
            console.error(`Status: ${error.response.status}`);
            console.error(error.response.data);
        }
    }
}

main();
