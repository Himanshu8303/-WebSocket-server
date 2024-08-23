# WebSocket App

This project is a WebSocket application with a Go server and a React client. The server reverses incoming messages and maintains a history of the last 5 messages. The client provides an interface to send messages and view message history.


## Prerequisites

- [Go](https://golang.org/dl/) (1.20 or later)
- [Node.js](https://nodejs.org/) (14 or later)
- [npm](https://www.npmjs.com/) (for managing client-side dependencies)

## Setup

### Server

1. **Navigate to the `server` directory:**

   ```bash
   cd server
   
2. Initialize the Go module
   `go mod init websocket-app/server`

3.Install dependencies:

   `go mod tidy`

4. Run the server:
    `go run main.go`
The server will start on http://localhost:8080.

### Client

1. **Navigate to the `client` directory:**

     `cd client`

2. Install dependencies:
     `npm install`

3. Start the client application:
     `npm start`
The client application will be available at http://localhost:3000.



