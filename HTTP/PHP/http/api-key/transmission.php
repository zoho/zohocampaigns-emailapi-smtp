<?php

// Function to construct the payload
function constructPayload() {
    // Recipients
    $recipient_data = [];

    $recipient1 = [
        "address" => "sophia@zylker.com",
        "name" => "Sophia Alexandri",
        "additional_data" => [
            "phone" => "+301234567890",
            "country" => "Greece"
        ],
        "merge_data" => [
            "first_name" => "Sophia"
        ]
    ];

    $recipient_data[] = $recipient1;

    // Content
    $content = [
        "subject" => "My first mail using Zoho Campaigns Email API HTTP",
        "html" => "<html><body>Welcome \$[first_name|Customer]\$!<br>Summer Hot Savings, You Don't Want to Miss</body></html>",
        "text" => "Welcome \$[first_name|Customer]\$! Summer Hot Savings, You Don’t Want to Miss",
        "from" => [
            "address" => "aron@zylker.com",
            "name" => "Aron Fletcher"
        ]
    ];

    // Payload
    $payload = [
        "transmission_name" => "Summer is here",
        "recipients" => $recipient_data,
        "content" => $content
    ];

    return $payload;
}

function main() {
    $api_key = "1000.****************"; // Replace with your API key

    // Construct the payload
    $payload = constructPayload();

    $url = "https://campaigns.zoho.com/emailapi/v2/transmission";
    $headers = [
        "Content-Type: application/json",
        "Accept: application/json",
        "Authorization: Zoho-zapikey $api_key"
    ];

    $ch = curl_init($url);
    curl_setopt($ch, CURLOPT_HTTPHEADER, $headers);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
    curl_setopt($ch, CURLOPT_POST, true);
    curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($payload));
    curl_setopt($ch, CURLOPT_TIMEOUT, 60);

    $response = curl_exec($ch);
    $http_status = curl_getinfo($ch, CURLINFO_HTTP_CODE);

    if ($response === false) {
        echo "Error: " . curl_error($ch);
    } else {
        echo "Status: " . $http_status . "\n";
        echo "Response: " . $response;
    }

    curl_close($ch);
}

main();

?>
