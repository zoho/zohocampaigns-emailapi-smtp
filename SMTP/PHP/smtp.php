<?php

function smtpSendMail($to, $subject, $htmlContent, $textContent, $headers, $server, $port, $username, $accessToken) {
    // Create a connection to the SMTP server
    $connection = fsockopen($server, $port, $errno, $errstr, 30);
    if (!$connection) {
        throw new Exception("Could not connect to SMTP server: $errstr ($errno)");
    }

    // Function to print server responses for debugging
    function debugPrint($label, $response) {
        echo "[DEBUG] $label: $response\n";
    }

    // Get server response
    function getServerResponse($connection) {
        $response = '';
        while ($str = fgets($connection, 515)) {
            $response .= $str;
            if (substr($str, 3, 1) == ' ') {
                break;
            }
        }
        return $response;
    }

    // Send a command to the SMTP server
    function sendCommand($connection, $command) {
        fwrite($connection, $command . "\r\n");
        return getServerResponse($connection);
    }

    // Start the SMTP conversation
    $response = getServerResponse($connection);
    debugPrint('Connection Response', $response);
    if (strpos($response, '220') !== 0) {
        throw new Exception("Connection error: $response");
    }

    // Send EHLO command
    $response = sendCommand($connection, 'EHLO ' . gethostname());
    debugPrint('EHLO Response', $response);
    if (strpos($response, '250') !== 0) {
        throw new Exception("EHLO error: $response");
    }

    // Start TLS encryption
    $response = sendCommand($connection, 'STARTTLS');
    debugPrint('STARTTLS Response', $response);
    if (strpos($response, '220') !== 0) {
        throw new Exception("STARTTLS error: $response");
    }

    // Enable crypto (TLS)
    stream_context_set_option($connection, 'ssl', 'verify_peer', false);
    stream_context_set_option($connection, 'ssl', 'verify_peer_name', false);
    stream_context_set_option($connection, 'ssl', 'allow_self_signed', true); // true for test environment, false for production environment
    if (!stream_socket_enable_crypto($connection, true, STREAM_CRYPTO_METHOD_TLSv1_3_CLIENT)) {
        throw new Exception("Unable to start TLS encryption");
    }

    // Re-send EHLO after TLS
    $response = sendCommand($connection, 'EHLO ' . gethostname());
    debugPrint('EHLO after TLS Response', $response);
    if (strpos($response, '250') !== 0) {
        throw new Exception("EHLO after TLS error: $response");
    }

    // Send AUTH command with access token
    $response = sendCommand($connection, 'AUTH ACCESS_TOKEN ' . $accessToken);
    debugPrint('AUTH Response', $response);
    if (strpos($response, '235') !== 0) {
        throw new Exception("Authentication error: $response");
    }

    // Send MAIL FROM command
    $response = sendCommand($connection, 'MAIL FROM:<' . $username . '>');
    debugPrint('MAIL FROM Response', $response);
    if (strpos($response, '250') !== 0) {
        throw new Exception("MAIL FROM error: $response");
    }

    // Send RCPT TO command
    foreach ($to as $recipient) {
        $response = sendCommand($connection, 'RCPT TO:<' . $recipient . '>');
        debugPrint('RCPT TO Response for ' . $recipient, $response);
        if (strpos($response, '250') !== 0 && strpos($response, '251') !== 0) {
            throw new Exception("RCPT TO error for $recipient: $response");
        }
    }

    // Send DATA command
    $response = sendCommand($connection, 'DATA');
    debugPrint('DATA Response', $response);
    if (strpos($response, '354') !== 0) {
        throw new Exception("DATA command error: $response");
    }

    // Send email headers and content
    $emailContent = $headers;
    $emailContent .= "Subject: $subject\r\n";
    $emailContent .= "To: " . implode(', ', $to) . "\r\n";
    $emailContent .= "MIME-Version: 1.0\r\n";
    $emailContent .= "Content-Type: multipart/alternative; boundary=\"boundary\"\r\n\r\n";
    $emailContent .= "--boundary\r\n";
    $emailContent .= "Content-Type: text/plain; charset=UTF-8\r\n\r\n";
    $emailContent .= $textContent . "\r\n\r\n";
    $emailContent .= "--boundary\r\n";
    $emailContent .= "Content-Type: text/html; charset=UTF-8\r\n\r\n";
    $emailContent .= $htmlContent . "\r\n\r\n";
    $emailContent .= "--boundary--\r\n";

    // End the data section
    $emailContent .= ".\r\n";

    // Send email data
    fwrite($connection, $emailContent);

    // Get server response for end of data
    $response = getServerResponse($connection);
    debugPrint('End of DATA Response', $response);
    if (strpos($response, '250') !== 0) {
        throw new Exception("End of data error: $response");
    }

    // Send QUIT command
    $response = sendCommand($connection, 'QUIT');
    debugPrint('QUIT Response', $response);

    // Close the connection
    fclose($connection);

    echo "Email sent successfully";
}

// Email details
$to = ["lucy@example.campaigns.zoho.com"];
$subject = "My first mail using Zoho Campaigns Email API SMTP";
$htmlContent = "<html><body>Welcome \$[first_name|Customer]$!<br>Summer Hot Savings, You Don't Want to Miss</body></html>";
$textContent = "Welcome \$[first_name|Customer]$! Summer Hot Savings, You Don’t Want to Miss";
$headers = "From: aron@marketing.campaigns.zoho.com\r\n";
$meta_data = [
    "campaign_name" => "Summer is here",
    "recipient_data" => [
        "lucy@example.campaigns.zoho.com" => [
            "name" => "Aaron Fletcher",
            "additional_data" => ["phone" => "+91231241444", "country" => "IN"],
            "merge_data" => ["first_name" => "Aaron"],
        ]
    ],
];
$jsonMetaData = json_encode($meta_data);
$headers .= "X-ZCEA-SMTP-DATA: " . $jsonMetaData . "\r\n";

// SMTP server configuration
$server = 'smtp.campaigns.zoho.com';
$port = 587;
$username = 'aron@marketing.campaigns.zoho.com';
$accessToken = '1000.*************************';

// Send the email
try {
    smtpSendMail($to, $subject, $htmlContent, $textContent, $headers, $server, $port, $username, $accessToken);
} catch (Exception $e) {
    echo "Error: " . $e->getMessage();
}

?>
