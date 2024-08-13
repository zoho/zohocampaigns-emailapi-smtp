package com.example;

import java.util.Date;
import java.util.Properties;

import javax.mail.AuthenticationFailedException;
import javax.mail.Message;
import javax.mail.Session;
import javax.mail.internet.InternetAddress;
import javax.mail.internet.MimeBodyPart;
import javax.mail.internet.MimeMessage;
import javax.mail.internet.MimeMultipart;
import com.sun.mail.smtp.SMTPTransport;

import org.json.JSONObject;

public class App {
    public static void main( String[] args ) {
        /*
		 * Prepare mail data
		 */

		String senderAddress = "aaron@zylker.com";
		String[] recipients = { "aaron.test@gmail.com", "ea.test@zylker.com" };
		String subject = "My first mail using Zoho Campaigns Email API SMTP";

		JSONObject metaData = new JSONObject();

		JSONObject recipientData = new JSONObject();

		// Recipient 1
		JSONObject recipient1 = new JSONObject();
		JSONObject additionalDataR1 = new JSONObject();
		JSONObject mergeDataR1 = new JSONObject();

		additionalDataR1.put("phone", "+919876543210");
		additionalDataR1.put("country", "IN");
		mergeDataR1.put("first_name", "Aaron");

		recipient1.put("name", "Aaron Fletcher");
		recipient1.put("additional_data", additionalDataR1);
		recipient1.put("merge_data", mergeDataR1);

		// Recipient 2
		JSONObject recipient2 = new JSONObject();
		JSONObject additionalDataR2 = new JSONObject();
		JSONObject mergeDataR2 = new JSONObject();

		additionalDataR2.put("phone", "+919876543210");
		additionalDataR2.put("country", "IN");
		mergeDataR2.put("first_name", "EA");

		recipient2.put("name", "EA Test");
		recipient2.put("additional_data", additionalDataR2);
		recipient2.put("merge_data", mergeDataR2);

		recipientData.put("aaron.test@gmail.com", recipient1);
		recipientData.put("ea.test@zylker.com", recipient2);

		metaData.put("campaign_name", "Summer is here");
		metaData.put("recipient_data", recipientData);

		// Body
		String htmlContent = "<html><body>Welcome $[first_name|Customer]$!<br>Summer Hot Savings, You Don't Want to Miss</body></html>";
		String textContent = "Welcome $[first_name|Customer]$! Summer Hot Savings, You Don’t Want to Miss";

		/*
		 * Create session and send mail
		 */
		try {
			// Set up mail session properties
			Properties properties = new Properties();
			properties.put("mail.smtp.auth", "true");
			properties.put("mail.smtp.starttls.enable", "true");
			properties.put("mail.smtp.ssl.protocols", "TLSv1.2");
			properties.put("mail.smtp.host", "smtp.campaigns.zoho.com");
			properties.put("mail.smtp.port", "587");
			Session session = Session.getInstance(properties);
			session.setDebug(true);

			// Create and configure the MimeMessage
			Message message = new MimeMessage(session);
			message.setFrom(new InternetAddress());
			for (String recipient : recipients) {
				message.addRecipient(Message.RecipientType.TO, new InternetAddress(recipient));
			}
			message.setSubject(subject);
			message.setFrom(new InternetAddress(senderAddress));
			message.setHeader("X-ZCEA-SMTP-DATA", metaData.toString());

			MimeMultipart mimemultipart = null;
			MimeBodyPart mainBodyPart = new MimeBodyPart();
			if (htmlContent != null) {
				mimemultipart = new MimeMultipart("alternative");
				mainBodyPart.setContent(htmlContent, "text/html;charset=UTF-8");
				mainBodyPart.addHeader("Content-Transfer-Encoding", "quoted-printable");
			}
			if (textContent != null) {
				mimemultipart = new MimeMultipart("alternative");
				mainBodyPart.setText(textContent, "UTF-8");
				mainBodyPart.setHeader("Content-Transfer-Encoding", "quoted-printable");
			}

			mimemultipart.addBodyPart(mainBodyPart);
			message.setContent(mimemultipart);
			message.setSentDate(new Date());

			SMTPTransport transport = (SMTPTransport) session.getTransport("smtp");

			// Connect to the SMTP server
			transport.connect("", "");

			// Send the custom AUTH ACCESS_TOKEN command
			String accessToken = "1000.***************************************";
			int response = transport.simpleCommand("AUTH ACCESS_TOKEN " + accessToken);
			if (response != 235) {
				throw new AuthenticationFailedException("AUTH ACCESS_TOKEN command failed: " + response);
			}

			// Send the email
			transport.sendMessage(message, message.getAllRecipients());
			System.out.println(transport.getLastServerResponse());
			transport.close();
		} catch (Exception e) {
			System.out.println(e.getMessage());
		}
	}
}
