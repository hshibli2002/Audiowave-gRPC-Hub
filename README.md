# MBPlayer gRPC Service

MBPlayer is a music data management service implemented using Go, gRPC, PostgreSQL, and Elasticsearch for advanced searching capabilities. It provides functionalities for managing songs, artists, and playlists.

## Features

- CRUD operations for songs, artists, and playlists.
- Search functionality for songs and playlists.
- User management, including follow/unfollow artists and like songs.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go 1.16 or later installed on your machine.
- PostgreSQL installed and running.
- (Optional) Elasticsearch installed and running for search functionalities.

## Setting Up

To set up the MBPlayer service for development, follow these steps:

### Clone the repository

```bash
git clone https://github.com/yourusername/MBPlayer.git
cd MBPlayer
```

## Configure the environment

Create a `.env` file in the root directory of the project and add the following environment variables:

```env
- DB_HOST=localhost
- DB_PORT=5432
- DB_USER=your_db_user
- DB_PASSWORD=your_db_password
- DB_NAME=your_db_name
- GRPC_PORT=50051
```

## Install dependencies

```bash
    go mod download        
```

## Running the service

To run the MBPlayer service, use the following command:

```bash
    go run cmd/main/main.go
```
