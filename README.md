# Audiowave gRPC Hub Service

Audiowave gRPC Hub Service is an educational project designed to deepen understanding of the Go programming language, particularly in the context of building scalable services with gRPC. It will incorporate Elasticsearch to explore SEO tools and already provides insights into structuring a robust Go project suitable for a microservices' architecture. This project also serves as a practical guide on hosting and securing such services.

## Learning Objectives

- Grasp the fundamentals and advanced topics of Go in the realm of backend development.
- Delve into gRPC for efficient server-client communication.
- Integrate Elasticsearch for powerful search capabilities and learning SEO tools.
- Comprehend the structure and components of a scalable Go project with gRPC.
- Master connecting and managing microservices in a distributed system.
- Gain hands-on experience in hosting and securing backend services.

## Features

- CRUD operations for songs, artists, and playlists. [X]
- Elastic search integration for enhanced searching. (Ongoing) []
- User management with follow and like features for artists and songs in addition to their metadata. []

## Prerequisites

Ensure the following tools and technologies are installed and properly configured:

- Go (1.16 or later).
- PostgreSQL for database management.
- Elasticsearch for advanced search features (optional but recommended for learning SEO).

## Installation and Setup

### Clone the Repository

```bash
git clone https://github.com/hshibli2002/Audiowave-gRPC-Hub.git
cd Audiowave-gRPC-Hub
```

### Environment Configuration

Create a `.env` file at the root of the project with the necessary configurations:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
GRPC_PORT=50051
```

### Install the Dependencies

```bash
go mod download
```

### Running the Services

```bash
go run cmd/main/main.go
```

## Additional Information
For those interested in testing with a live database, I provide access to a hosted database instance.
To request access, please email me at [hshibli2002@gmail.com]() with your inquiry.
