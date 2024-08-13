<?php

function main() {
    $access_token = "1000.****************";

    // Recipients
    $recipient_data = [];

    $recipient1 = [
        "address" => "lucy@example.campaigns.zoho.com",
        "name" => "Aaron Fletcher",
        "additional_data" => [
            "phone" => "+919876543210",
            "country" => "IN"
        ],
        "merge_data" => [
            "first_name" => "Aaron"
        ]
    ];

    $recipient_data[] = $recipient1;

    // Content
    $content = [
        "subject" => "My first mail using Zoho Campaigns Email API HTTP",
        "html" => "<html><body>Welcome \$[first_name|Customer]\$!<br>Summer Hot Savings, You Don't Want to Miss</body></html>",
        "text" => "Welcome \$[first_name|Customer]\$! Summer Hot Savings, You Don’t Want to Miss",
        "from" => [
            "address" => "aron@marketing.campaigns.zoho.com",
            "name" => "Aron Fletcher"
        ]
    ];

    // Payload
    $payload = [
        "campaign_name" => "hello_customer",
        "recipients" => $recipient_data,
        "content" => $content
    ];

    $url = "https://campaigns.zoho.com/emailapi/v1/transmission";
    $headers = [
        "Content-Type: application/json",
        "Accept: application/json",
        "Authorization: Zoho-oauthtoken $access_token"
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

        if ($http_status >= 200 && $http_status < 300) {
            $json_response = json_decode($response, true);
            echo "Response: " . json_encode($json_response, JSON_PRETTY_PRINT);
        } else {
            echo "Error: " . $http_status . "\n";
            echo $response;
        }
    }

    curl_close($ch);
}

main();

?>
