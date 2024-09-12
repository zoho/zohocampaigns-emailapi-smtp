<?php
use PHPMailer\PHPMailer\PHPMailer;
use PHPMailer\PHPMailer\Exception;

// Load Composer's autoloader
require 'vendor/autoload.php';

function smtpSendMail() {
    $emailData = prepareEmailData();
    
    // Create a new PHPMailer instance
    $mail = new PHPMailer(true);

    $mail->SMTPDebug = 2;
    $mail->Debugoutput = function($str, $level) {
        echo "Debug level $level; message: $str<br>";
    };

    $mail->isSMTP();
    $mail->Host       = $emailData["host"];
    $mail->Port       = $emailData["port"];
    $mail->SMTPAuth   = true;
    $mail->Username   = $emailData["username"];
    $mail->Password   = $emailData["password"];
    $mail->SMTPSecure = PHPMailer::ENCRYPTION_STARTTLS;

    // Sender and recipient settings
    $mail->setFrom($emailData["senderAddress"]);

    foreach ($emailData["recipients"] as $recipient) {
        $mail->addAddress($recipient);
    }

    // Email content
    $mail->isHTML(true);
    $mail->Subject = $emailData["subject"];
    $mail->Body    = $emailData["htmlContent"];
    $mail->AltBody = $emailData["textContent"];

    // Custom headers
    $mail->addCustomHeader('X-ZCEA-SMTP-DATA', json_encode($emailData["metaData"]));

    // Send the email
    $mail->send();    
}

// Prepare email data
function prepareEmailData() {
    $emailData = array(
        "recipients" => [
            "sophia@zylker.com"
        ],
        "subject" => "My first mail using Zoho Campaigns Email API SMTP",
        "htmlContent" => "<html><body>Welcome \$[first_name|Customer]$!<br>Summer Hot Savings, You Don't Want to Miss</body></html>",
        "textContent" => "Welcome \$[first_name|Customer]$! Summer Hot Savings, You Don’t Want to Miss",
        "metaData" => array(
            "transmission_name" => "Summer is here",
            "recipient_data" => [
                "sophia@zylker.com" => [
                    "name" => "Sophia Alexandri",
                    "additional_data" => ["phone" => "+301234567890", "country" => "Greece"],
                    "merge_data" => ["first_name" => "Sophia"],
                ]
            ],
        ),
        "host" => "smtp-campaigns.zoho.com",
        "port" => 587,
        "senderAddress" => "aron@zylker.com",
        "username" => "apikey",
        "password" => "1000.*************************"
    );
    return $emailData;
}

// Send the email
try {
    smtpSendMail();
} catch (Exception $e) {
    echo "Error: " . $e->getMessage();
}

?>
