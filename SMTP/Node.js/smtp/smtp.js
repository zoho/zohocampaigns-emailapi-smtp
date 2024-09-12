const nodemailer = require("nodemailer");

async function sendEmail() {
  // Prepare email data
  emailData = prepareEmailData();

  // Create a transporter
  let transporter = nodemailer.createTransport({
    host: emailData.host,
    port: emailData.port,
    secure: false, // set true to use SSL, false to use TLS
    auth: {
      type: "login",
      user: emailData.username,
      pass: emailData.password,
    },
    authMethod: "LOGIN",
    logger: true,
    debug: true
  });

  // Prepare the email options
  let mailOptions = {
    from: emailData.senderAddress,
    to: emailData.recipients.join(", "),
    subject: emailData.subject,
    text: emailData.textContent,
    html: emailData.htmlContent,
    headers: {
      "X-ZCEA-SMTP-DATA": {
        prepared: true,
        value: JSON.stringify(emailData.metaData)
      }
    },
  };

  // Send the email
  try {
    let info = await transporter.sendMail(mailOptions);
    console.log("Message sent: %s", info.messageId);
  } catch (error) {
    console.error("Error:", error);
  }
}

function prepareEmailData() {
  const emailData = {};
  emailData.port = 587;
  emailData.host = "smtp-campaigns.zoho.com";

  emailData.subject = "My first mail using Zoho Campaigns Email API SMTP";
  emailData.senderAddress = "aaron@zylker.com";
  emailData.recipients = ["sophia@zylker.com"];
  emailData.metaData = {
    transmission_name: "Summer is here",
    recipient_data: {
      "sophia@zylker.com": {
        name: "Sophia Alexandri",
        additional_data: {
          phone: "+301234567890",
          country: "Greece",
        },
        merge_data: {
          first_name: "Sophia",
        },
      },
    },
  };

  emailData.htmlContent =
    "<html><body>Welcome $[first_name|Customer]$!<br>Summer Hot Savings, You Don't Want to Miss</body></html>";
  emailData.textContent =
    "Welcome $[first_name|Customer]$! Summer Hot Savings, You Don’t Want to Miss";

  emailData.username = "apikey";
  emailData.password = "1000.****************************"; // Replace with your access token

  return emailData;
}

sendEmail();
