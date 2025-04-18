# Go Basic Server App

This is a simple Go-based web server application that serves static files and handles basic HTTP requests.

## Features

- Serves static HTML files from the `static/` directory.
- Handles a `/hello` endpoint that responds with "Hello, World!".
- Handles a `/form` endpoint to process form submissions.

## Project Structure

## Endpoints

### `/`

Serves static files from the `static/` directory. For example:

- `/` serves `index.html`.
- `/hello.html` serves `hello.html`.

### `/hello`

Responds with a simple "Hello, World!" message.

### `/form`

Handles form submissions. Accepts `POST` requests with the following fields:

- `name`
- `address`

Responds with the submitted data.

## How to Run

1. Make sure you have Go installed (version 1.24.2 or later).
2. Navigate to the `basic-server-app` directory.
3. Run the server:
   ```sh
   go run main.go
   ```
