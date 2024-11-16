# Mail Service

A robust mail service built with Go, Gin framework, and RabbitMQ. This service handles sending emails using SMTP and provides an easy interface for sending, validating, and publishing emails through RabbitMQ.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [API Endpoints](#api-endpoints)
- [Testing](#testing)

## Features

- **Send Emails:** Send emails through the SMTP Gmail server.
- **Email Validation:** Ensure that email addresses, subjects, and bodies are properly formatted.
- **Queue System:** Publish email requests to a RabbitMQ queue for asynchronous handling.
- **Error Handling:** Handle validation and service errors, ensuring clear feedback.
- **Logging:** Comprehensive logging for debugging and monitoring purposes.

## Prerequisites

- **Go:** Version 1.22 or higher.
- **RabbitMQ:** A running instance of RabbitMQ for email queuing.
- **Environment Variables:** Use a `.env` file to set required environment variables.

## Installation

### Clone the Repository

    ```bash
        git clone https://github.com/Renan-Parise/mail.git
        cd mail
    ```

### Install Dependencies

```bash
    go mod tidy
```

## Configuration

### Environment Variables

Create a `.env` file in the project root and add the following environment variables:

```env
    RABBITMQ_URL=
    GMAIL_USERNAME=
    GMAIL_PASSWORD=
```

### Run RabbitMQ

Ensure RabbitMQ is running on the specified URL from your environment file. We already have a docker-compose file that will run RabbitMQ for you.

```bash
    docker-compose up -d
```

## API Endpoints

### Public Routes:

- **POST /mail/send:** Send an email by providing address, subject, and body in the request payload.

## Testing

### Run Tests

You can run unit tests to ensure that the service behaves as expected.

```bash
    go test ./...
```

## Logging

This service uses structured logging for tracing and error management. Logs will be output to the console and will provide insights into any issues with sending emails or RabbitMQ communication failures.