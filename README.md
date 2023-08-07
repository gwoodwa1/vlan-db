# VLAN Database Server

This project provides a simple HTTP server for managing a database of VLANs (Virtual Local Area Networks). It provides endpoints for fetching all VLANs and adding new ones. The application also prevents adding VLANs with IDs that are reserved.

## Features

- Fetch all VLANs: `GET /vlans`
- Add a new VLAN: `POST /addnew`

## Technologies

- Golang
- SQLite3
- JSON

## Directory Structure

Here is the basic directory structure of this project:
```
├── main.go
├── handlers
│ └── handler.go
├── types
│ └── types.go
├── db
│ └── db.go
├── static
│ └── (static files to be served, like HTML, CSS, JS)
└── README.md
```

- `main.go`: Entry point for the application. This file sets up the HTTP server, defines the routes, and connects to the SQLite3 database.
- `handlers/handler.go`: Defines the HTTP handlers for the various endpoints. It includes functions to fetch all VLANs from the database and add a new VLAN.
- `types/types.go`: Defines the data structures (`structs`) used in this project, such as the `Vlan` and `ReservedVlans` structs.
- `db/db.go`: Manages the database connection and related operations.
- `static`: A directory that contains static files to be served (if any), such as HTML, CSS, and JavaScript files.

## How to Run

1. Clone this repository to your local machine.
2. In the terminal, navigate to the directory of the cloned repository.
3. Run `go run main.go`. This will start the server on port 8080.
4. To seed some entries into the Database `go run seed.go`

Make sure that you have Go installed on your machine.
