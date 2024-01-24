# People-Base API

People-Base API is a RESTful web service for managing and retrieving information about individuals. It provides CRUD (Create, Read, Update, Delete) operations for people, along with filtering options based on name, age, gender, and nationality.

## Table of Contents
- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
- [Usage](#usage)
- [API Documentation](#api-documentation)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

### Prerequisites

Before running the People-Base API, make sure you have the following prerequisites installed:

- Go (Programming language): [Installation Guide](https://golang.org/doc/install)
- SQLite3 (Database): [Installation Guide](https://www.sqlite.org/download.html)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/ilkhom19/task_peopler.git
   ```

2. Change to the project directory:

   ```bash
   cd People-Base-API
   ```

3. Build and run the API:

   ```bash
   source .env && go build -o app main.go && ./app
   ```

   The API will be accessible at `http://localhost:{port}/docs/index.html`.

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
  curl http://localhost:8080/people/gender?gender=Male&page=1&pageSize=10
  ```

- Get people by nationality:

  ```bash
  curl http://localhost:8080/people/nationality?nationality=USA&page=1&pageSize=10
  ```