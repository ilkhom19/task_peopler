# People-Base API

People-Base API is a RESTful web service for managing and retrieving information about individuals. It provides CRUD (Create, Read, Update, Delete) operations for people, along with filtering options based on name, age, gender, and nationality.

## Table of Contents
- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
- [Usage](#usage)

## Getting Started

### Prerequisites

Before running the People-Base API, make sure you have the following prerequisites installed:

- Go (Programming language): [Installation Guide](https://golang.org/doc/install)
- SQLite3 (Database): [Installation Guide](https://www.sqlite.org/download.html)
- Goose (Database migration tool): [Installation Guide](https://github.com/pressly/goose)
- Swagger (API documentation tool): [Installation Guide](https://github.com/swaggo/swag)
### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/ilkhom19/task_peopler.git
   ```

2. Change to the project directory:

   ```bash
   cd task_peopler
   ```

3. Build and run the API:

   ```bash
   swag init --pd -g cmd/main.go
   source .env && go build -o app main.go && ./app
   ```

   The API will be accessible at `http://localhost:8080/docs/index.html`.
    ![Screenshot-20240125013852-2594x1250.png](..%2F..%2FDownloads%2FScreenshot-20240125013852-2594x1250.png)
## Usage

The API provides endpoints for managing people. You can use tools like `curl` or Postman to interact with the API. Here are some example requests:

- Create a new person:

  ```bash
  curl -X POST -H "Content-Type: application/json" -d '{"name": "John", "surname": "Doe"}' http://localhost:8080/people
  ```

- Retrieve a person by ID:

  ```bash
  curl http://localhost:8080/people/1
  ```

- Delete a person by ID:

  ```bash
  curl -X DELETE http://localhost:8080/people/1
  ```

- Get people by name:

  ```bash
  curl http://localhost:8080/people/name?name=John&page=1&pageSize=10
  ```

- Get people by age:

  ```bash
  curl http://localhost:8080/people/age?age=30&page=1&pageSize=10
  ```

- Get people by gender:

  ```bash
  curl http://localhost:8080/people/gender?gender=male&page=1&pageSize=10
  ```

- Get people by nationality:

  ```bash
  curl http://localhost:8080/people/nationality?nationality=UZ&page=1&pageSize=10
  ```