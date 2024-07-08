//$Id$
package com.example;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.net.URL;
import java.net.URLConnection;
import java.nio.charset.StandardCharsets;

import javax.net.ssl.HttpsURLConnection;

import org.json.JSONArray;
import org.json.JSONObject;

public class App {
	public static void main(String args[]) {
        String accessToken = "1000.****************************";

		// Recipients
        JSONArray recipientData = new JSONArray();
		JSONObject recipient1 = new JSONObject();
		JSONObject additionalDataR1 = new JSONObject();
		JSONObject mergeDataR1 = new JSONObject();

		additionalDataR1.put("phone", "+919876543210");
		additionalDataR1.put("country", "IN");
		mergeDataR1.put("first_name", "Aaron");

		recipient1.put("address", "lucy@example.campaigns.zoho.com");
		recipient1.put("name", "Aaron Fletcher");
		recipient1.put("additional_data", additionalDataR1);
		recipient1.put("merge_data", mergeDataR1);

		recipientData.put(recipient1);
		
		// Content
		JSONObject content = new JSONObject();
		content.put("subject", "My first mail using Zoho Campaigns Email API HTTP");
		content.put("html", "<html><body>Welcome $[first_name|Customer]$!<br>Summer Hot Savings, You Don't Want to Miss</body></html>");
		content.put("text", "Welcome $[first_name|Customer]$! Summer Hot Savings, You Don’t Want to Miss");
		
		JSONObject fromData = new JSONObject();
		fromData.put("address", "aron@marketing.campaigns.zoho.com");
		fromData.put("name", "Aron Fletcher");
		content.put("from", fromData);

		// Payload
		JSONObject payload = new JSONObject();
		payload.put("campaign_name", "hello_customer");
		payload.put("recipients", recipientData);
		payload.put("content", content);

        try {
        	URL url = new URL("https://campaigns.zoho.com/emailapi/v1/transmission");
        	URLConnection urlConnection = (HttpsURLConnection) url.openConnection();
			((HttpsURLConnection) urlConnection).setRequestMethod("GET");
			urlConnection.setConnectTimeout(60000);
			urlConnection.setReadTimeout(60000);
			urlConnection.setDoInput(true);
			urlConnection.setDoOutput(true);
			urlConnection.setRequestProperty("Content-Type", "application/json");
            urlConnection.setRequestProperty("Accept", "application/json");
            urlConnection.setRequestProperty("Authorization", "Zoho-oauthtoken " + accessToken);
			if (payload != null) {
				OutputStreamWriter outputStreamWriter = new OutputStreamWriter(((HttpsURLConnection) urlConnection).getOutputStream(), "UTF-8");// No I18N
				outputStreamWriter.write(payload.toString());
				outputStreamWriter.flush();
			}
			int status = ((HttpsURLConnection) urlConnection).getResponseCode();
			
            // Read the response
            try (BufferedReader br = new BufferedReader(new InputStreamReader(((HttpsURLConnection) urlConnection).getInputStream(), StandardCharsets.UTF_8))) {
                StringBuilder response = new StringBuilder();
                String responseLine;
                while ((responseLine = br.readLine()) != null) {
                    response.append(responseLine.trim());
                }
                // Parse the response
                JSONObject jsonResponse = new JSONObject(response.toString());
                System.out.println("Response: " + jsonResponse.toString(2));
            }
            ((HttpsURLConnection) urlConnection).disconnect();
        } catch (Exception e) {
            e.printStackTrace();
        }
	}
}
