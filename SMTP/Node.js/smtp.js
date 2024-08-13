const nodemailer = require("nodemailer");

const PORT = 587;
const HOST = "smtp.campaigns.zoho.com";

const senderAddress = "aaron@zylker.com";
const recipients = ["aaron.test@gmail.com", "ea.test@zylker.com"];
const subject = "My first mail using Zoho Campaigns Email API SMTP";

const recipientData = {
  "aaron.test@gmail.com": {
    name: "Aaron Fletcher",
    additionalData: {
      phone: "+91231241444",
      country: "IN",
    },
    mergeData: {
      first_name: "Aaron",
    },
  },
  "ea.test@zylker.com": {
    name: "EA Test",
    additionalData: {
      phone: "+9198762313210",
      country: "IN",
    },
    mergeData: {
      first_name: "EA",
    },
  },
};

const metaData = {
  campaign_name: "Summer is here",
  recipient_data: recipientData,
};

const htmlContent =
  "<html><body>Welcome $[first_name|Customer]$!<br>Summer Hot Savings, You Don't Want to Miss</body></html>";
const textContent =
  "Welcome $[first_name|Customer]$! Summer Hot Savings, You Don’t Want to Miss";

const accessToken =
  "1000.***************************************"; // Replace with your access token

async function sendEmail() {
  // Create a transporter
  let transporter = nodemailer.createTransport({
    host: HOST,
    port: PORT,
    secure: false, // use TLS
    auth: {
      type: "custom",
      method: "ACCESS_TOKEN",
      user: "username",
      pass: "password",
      options: {
        accessToken,
      },
    },
    authMethod: "ACCESS_TOKEN",
    customAuth: {
      ACCESS_TOKEN: async (ctx) => {
        let cmd = await ctx.sendCommand(
          "AUTH ACCESS_TOKEN " + ctx.auth.credentials.options.accessToken.trim()
        );
      },
    },
    logger: true,
    debug: true,
  });

  // Prepare the email options
  let mailOptions = {
    from: senderAddress,
    to: recipients.join(", "),
    subject: subject,
    text: textContent,
    html: htmlContent,
    headers: {
      "X-ZCEA-SMTP-DATA": JSON.stringify(metaData),
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

sendEmail();
