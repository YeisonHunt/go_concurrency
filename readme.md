# Golang PostgreSQL API with Concurrency Testing

This project is a simple Golang API that connects to a PostgreSQL database. The API includes an endpoint to query data from the database, and it includes test scripts to measure the latency of requests with and without concurrency.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)
  - [Sequential Latency Test](#sequential-latency-test)
  - [Concurrent Latency Test](#concurrent-latency-test)

## Features

- Simple API to query data from PostgreSQL.
- Measure request latency with and without concurrency.
- Easy-to-follow structure and installation process.

## Installation

### Prerequisites

- [Go](https://golang.org/dl/) (1.16 or higher)
- [PostgreSQL](https://www.postgresql.org/) (Ensure PostgreSQL is installed and running)
- Basic knowledge of Golang and PostgreSQL

### Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/YeisonHunt/go_concurrency
   cd your-repo-name
   ```

2. Initialize go modules:

    ```bash
    go mod init concurrency
    go mod tidy
    ```

3. Install the required dependencies:

    ```bash
    Copy code
    go get github.com/lib/pq
    ```
4. Configure the database connection in main.go:

    ```go
    connStr := "user=youruser dbname=yourdb sslmode=disable password=yourpassword"
    ```