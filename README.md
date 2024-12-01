# User Authentication System

This project implements a simple User Authentication System that integrates MySQL and Redis as databases. It provides functionalities like user creation, retrieval, updating, and deletion, along with a connection to MySQL and Redis for persistent storage and caching. 

## Quick Start

To get started with this project, follow the steps below:

Clone the repository:

```bash
git clone https://github.com/SudhanshuJoshi09/user-auth-system.git

cd user-auth-system
```

Install Dependencies:
```bash
go mod tidy
```

Run the Application:
```bash
go run main.go
```

### Prerequisites

1. **Install MySQL**: Make sure MySQL is running locally or on a remote server. If you're running it locally, ensure the port `3307` is open and accessible.

2. **Install Redis**: Redis should be installed and running locally or on a remote server.

3. **Go**: Ensure you have Go installed on your machine (version 1.18+ is recommended).

4. **Docker**: If you want to run MySQL and Redis via Docker, you can use the following commands to set them up quickly.

### Running MySQL and Redis using Docker

To run MySQL with Docker:

```bash
docker run -d --name mysql-container -e MYSQL_ALLOW_EMPTY_PASSWORD=yes -p 3307:3306 mysql:latest
```

To run Redis with Docker:

```bash
docker run -d --name redis-container -p 6379:6379 redis:latest
```

## Usage


