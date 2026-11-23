# Zoho Campaigns Email API

Welcome to the official repository of Zoho Campaigns Email API! This repository provides sample code for sending bulk emails using our HTTP and SMTP APIs. We've included examples in various programming languages to help you get started quickly.

## Table of Contents

- [Introduction](#introduction)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Security Notes](#security-notes)
- [HTTP API Clients](#http-api-clients)
  - [Java](#java-http-client)
  - [Python](#python-http-client)
  - [Go](#go-http-client)
  - [Node.js](#nodejs-http-client)
  - [PHP](#php-http-client)
- [SMTP API Clients](#smtp-api-clients)
  - [Java](#java-smtp-client)
  - [Python](#python-smtp-client)
  - [Go](#go-smtp-client)
  - [Node.js](#nodejs-smtp-client)
  - [PHP](#php-smtp-client)

## Introduction

This repository contains sample code for interacting with our service. Whether you prefer using HTTP or SMTP protocols, we have you covered with examples in multiple programming languages.

## Getting Started

### Prerequisites

To use the sample code provided, you'll need to replace placeholder access token with your own. You can follow the steps provided [here](https://www.zoho.com/campaigns/help/emailapi/authentication.html) to generate an access token.

### Security Notes

The samples are starting points. Before using them in production:

- **Do not enable protocol tracing or debug logging in production.** It is intended only for debugging in a local setup. SMTP debug output (for example, `set_debuglevel` in Python, `session.setDebug` in Java, `SMTPDebug` in PHPMailer, `debug`/`logger` in Nodemailer) records the authentication exchange, which contains your API key, along with recipients and message content. Anyone who can read your logs could then use the key.
- **Keep the API key out of source code and logs.** Load it from an environment variable or a secrets manager instead of hardcoding it, and never commit it.
- **Rotate the key** if it may have appeared in logs, and restrict who can read application logs.

## HTTP API Clients

### Java HTTP Client

[Link to Java HTTP Client Instructions](HTTP/Java/http)

### Python HTTP Client

[Link to Python HTTP Client Instructions](HTTP/Python/http)

### Go HTTP Client

[Link to Go HTTP Client Instructions](HTTP/Go/http)

### Node.js HTTP Client

[Link to Node.js HTTP Client Instructions](HTTP/Node.js/http)

### PHP HTTP Client

[Link to PHP HTTP Client Instructions](HTTP/PHP/http)

## SMTP API Clients

### Java SMTP Client

[Link to Java SMTP Client Instructions](SMTP/Java/smtp)

### Python SMTP Client

[Link to Python SMTP Client Instructions](SMTP/Python/smtp)

### Go SMTP Client

[Link to Go SMTP Client Instructions](SMTP/Go/smtp)

### Node.js SMTP Client

[Link to Node.js SMTP Client Instructions](SMTP/Node.js/smtp)

### PHP SMTP Client

[Link to PHP SMTP Client Instructions](SMTP/PHP/smtp)
