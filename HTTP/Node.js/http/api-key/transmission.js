const axios = require('axios');

// Function to construct the payload
function constructPayload() {
    // Recipients
    const recipientData = [{
        address: "sophia@zylker.com",
        name: "Sophia Alexandri",
        additional_data: {
            phone: "+301234567890",
            country: "Greece"
        },
        merge_data: {
            first_name: "Sophia"
        }
    }];

    // Content
    const content = {
        subject: "My first mail using Zoho Campaigns Email API HTTP",
        html: "<html><body>Welcome $[first_name|Customer]$!<br>Summer Hot Savings, You Don't Want to Miss</body></html>",
        text: "Welcome $[first_name|Customer]$! Summer Hot Savings, You Don’t Want to Miss",
        from: {
            address: "aron@zylker.com",
            name: "Aron Fletcher"
        }
    };

    // Payload
    const payload = {
        campaign_name: "Summer is here",
        recipients: recipientData,
        content: content
    };

    return payload;
}

async function main() {
    const apiKey = "1000.****************************"; // Replace with your API key

    // Construct the payload
    const payload = constructPayload();

    try {
        const url = "https://campaigns.zoho.com/emailapi/v2/transmission";
        const headers = {
            "Content-Type": "application/json",
            "Accept": "application/json",
            "Authorization": `Zoho-zapikey ${apiKey}`
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
