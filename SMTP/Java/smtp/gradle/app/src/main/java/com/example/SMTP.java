package com.example;

import java.util.Date;
import java.util.Properties;

import javax.mail.Message;
import javax.mail.Session;
import javax.mail.internet.InternetAddress;
import javax.mail.internet.MimeBodyPart;
import javax.mail.internet.MimeMessage;
import javax.mail.internet.MimeMultipart;

import org.json.JSONObject;

import com.sun.mail.smtp.SMTPTransport;

import org.json.JSONArray;

public class SMTP {
    public static void main( String[] args ) {
		/*
		 * Create session and send mail
		 */
		try {
			JSONObject emailData = prepareEmailData();

			// Set up mail session properties
			Properties properties = new Properties();
			properties.put("mail.smtp.auth", "true");
			properties.put("mail.smtp.starttls.enable", "true");
			properties.put("mail.smtp.ssl.protocols", "TLSv1.3");
			properties.put("mail.smtp.host", emailData.getString("host"));
			properties.put("mail.smtp.port", String.valueOf(emailData.getInt("port")));
			Session session = Session.getInstance(properties);
			session.setDebug(true);

			// Create and configure the MimeMessage
			Message message = new MimeMessage(session);
			message.setFrom(new InternetAddress());
			JSONArray recipients = emailData.getJSONArray("recipients");
			for (int i = 0; i < recipients.length(); i++) {
				String recipient = recipients.getString(i);
				message.addRecipient(Message.RecipientType.TO, new InternetAddress(recipient));
			}
			message.setSubject(emailData.getString("subject"));
			message.setFrom(new InternetAddress(emailData.getString("senderAddress")));
			message.setHeader("X-ZCEA-SMTP-DATA", emailData.getJSONObject("metaData").toString());

			MimeMultipart mimemultipart = null;
			MimeBodyPart mainBodyPart = new MimeBodyPart();
			if (!emailData.isNull("htmlContent")) {
				mimemultipart = new MimeMultipart("alternative");
				mainBodyPart.setContent(emailData.getString("htmlContent"), "text/html;charset=UTF-8");
				mainBodyPart.addHeader("Content-Transfer-Encoding", "quoted-printable");
			}
			if (!emailData.isNull("textContent")) {
				mimemultipart = new MimeMultipart("alternative");
				mainBodyPart.setText(emailData.getString("textContent"), "UTF-8");
				mainBodyPart.setHeader("Content-Transfer-Encoding", "quoted-printable");
			}

			mimemultipart.addBodyPart(mainBodyPart);
			message.setContent(mimemultipart);
			message.setSentDate(new Date());

			SMTPTransport transport = (SMTPTransport) session.getTransport("smtp");

			// Connect to the SMTP server
			transport.connect(emailData.getString("username"), emailData.getString("password"));

			// Send the email
			transport.sendMessage(message, message.getAllRecipients());
			System.out.println(transport.getLastServerResponse());
			transport.close();
		} catch (Exception e) {
			System.out.println(e.getMessage());
		}
	}

	private static JSONObject prepareEmailData() throws Exception {
		JSONObject emailData = new JSONObject();

		String host = "smtp-campaigns.zoho.com";
		int port = 587;
		String username = "apikey";
		String password = "1000.***************************************"; // Replace with your access token

		String senderAddress = "aaron@zylker.com";
		String subject = "My first mail using Zoho Campaigns Email API SMTP";

		// Recipients and meta data
		JSONArray recipients = new JSONArray();
		recipients.put("sophia@zylker.com");
		
		JSONObject metaData = new JSONObject();
		JSONObject recipientData = new JSONObject();

		JSONObject recipient1 = new JSONObject();
		JSONObject additionalDataR1 = new JSONObject();
		JSONObject mergeDataR1 = new JSONObject();

		additionalDataR1.put("phone", "+301234567890");
		additionalDataR1.put("country", "Greece");
		mergeDataR1.put("first_name", "Sophia");

		recipient1.put("name", "Sophia Alexandri");
		recipient1.put("additional_data", additionalDataR1);
		recipient1.put("merge_data", mergeDataR1);

		recipientData.put("sophia@zylker.com", recipient1);

		metaData.put("transmission_name", "Summer is here");
		metaData.put("recipient_data", recipientData);

		// Body
		String htmlContent = "<html><body>Welcome $[first_name|Customer]$!<br>Summer Hot Savings, You Don't Want to Miss</body></html>";
		String textContent = "Welcome $[first_name|Customer]$! Summer Hot Savings, You Don’t Want to Miss";

		emailData.put("host", host);
		emailData.put("port", port);
		emailData.put("username", username);
		emailData.put("password", password);
		emailData.put("senderAddress", senderAddress);
		emailData.put("subject", subject);
		emailData.put("recipients", recipients);
		emailData.put("metaData", metaData);
		emailData.put("htmlContent", htmlContent);
		emailData.put("textContent", textContent);

		return emailData;
	}
}
