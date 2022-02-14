<h1 align="center">
  <img src="https://i.ibb.co/Y7sr1z5/Solution-2-Request-Header-Parser.png" alt="Logo" width="600">
</h1>

## Features

- IP Address, language, and software used by the user

## Tech Stack

**Server:** Golang (Gin)

## Run Locally

**Your server is accessible in port 3000**
Here is the steps to run it with `golang`

```bash
# Move to directory
$ cd workspace

# Clone this repository
$ git clone https://github.com/madjiebimaa/fcc-request-header-parser-ms.git

# Move to project
$ cd fcc-request-header-parser-ms

# Set gin to release mode
$ export GIN_MODE=release

# Run the application
$ go run main.go
```

Here is the steps to run it with `docker-compose`

```bash
# Move to directory
$ cd workspace

# Clone this repository
$ git clone https://github.com/madjiebimaa/fcc-request-header-parser-ms.git

# Move to project
$ cd fcc-request-header-parser-ms

# Set gin to release mode
$ export GIN_MODE=release

# Download, setup, and run the image
$ docker-compose up -d
```

## Running Tests

To run tests and get the percentage of code coverage, run the following command

```bash
  go test ./... -cover
```

## Lessons Learned

- How to create API with Golang (Gin)
- How to create middleware
- How to create gin handler
- How to create unit test for handler
- How to get data from request header

## API Reference

[Test API with REST Client Extension](https://github.com/madjiebimaa/fcc-request-header-parser-ms/tree/main/docs/apis)
