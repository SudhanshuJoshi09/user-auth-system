# User Authentication System

This project implements a simple User Authentication System that integrates MySQL and Redis as databases. It provides functionalities like user creation, login, logout, and refresh.

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

### 1. **Signup**
Create a new user by providing `name`, `email`, and `password`.

```bash
curl --location 'localhost:8082/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "sudhanshu12",
    "email": "asdfasdfasdfasdf@gmail.com",
    "password": "something12"
}'
```

### Response
- Status Code: 201 Created
- Body
```json
{
    "UserId": 21,
    "message": "User created successfuly"
}
```

### 2. **Login**
Login with `email` and `password` to get user Authorization header.
```bash
curl --location 'localhost:8082/signin' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "asdfasdfasdasdf@gmail.com",
    "password": "something12"
}'
```

### Response
- Status Code: 200 OK
- Body
```json
{
    "message": "User logged in successfuly",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFzZGZhc2RmYXNkYXNkZkBnbWFpbC5jb20iLCJleHAiOjE3MzMwNjcwMjF9.IHUMo2t62Mh9PuHnMCRazXbvdjLf5_MKDV7he6XbqpE"
}
```

### 3. **Logout/Revoke**
Logout from account using authorization header.
- Header {Authorization: Login Token}
```bash
curl --location 'localhost:8082/signout' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFzZGZhc2RmYXNkYXNkZkBnbWFpbC5jb20iLCJleHAiOjE3MzMwNjcwMDF9.uM4yf06UufB9W9lo1V5TPT0u2NDHifYk5VJqunIJdlU' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "asdfasdfasdasdf@gmail.com",
    "password": "something12"

}'
```

### Response
- Status Code: 200 OK
- Body:
```json
{
    "message": "User has logged out successfuly"
}
```
### 4. **Refresh Token**
Refresh token using current authorization header.
- Header {Authorization: Login Token}
```bash
curl --location --request POST 'localhost:8082/refresh' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFzZGZhc2RmYXNkYXNkZkBnbWFpbC5jb20iLCJleHAiOjE3MzMwNjY5MzB9.yg8_tJliDHyedCMFd-PHcqRnWmO2TvT7RXCYyZfEwl4'
```

### Response
- Status Code: 200 OK
- Body:
```json
{
    "message": "Your token has been refreshed.",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFzZGZhc2RmYXNkYXNkZkBnbWFpbC5jb20iLCJleHAiOjE3MzMwNjcwMDF9.uM4yf06UufB9W9lo1V5TPT0u2NDHifYk5VJqunIJdlU"
}
```
