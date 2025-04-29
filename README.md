# Go Bookstore

A simple RESTful API for managing a bookstore, built with Go.

## Features

- Add, update, delete, and retrieve books.
- Lightweight and fast API using Go's standard library.
- JSON-based communication.

## Prerequisites

- [Go](https://golang.org/dl/) (version 1.18 or higher)
- [Git](https://git-scm.com/)

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/ZaouiAmine/go-bookstore.git
    cd go-bookstore
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Run the application:

    ```bash
    go run main.go
    ```

4. The API will be available at `http://localhost:9010`.

## API Endpoints

| Method | Endpoint       | Description          |
|--------|----------------|----------------------|
| GET    | `/books`       | Get all books       |
| GET    | `/book/{id}`  | Get a book by ID    |
| POST   | `/book`       | Add a new book      |
| PUT    | `/book/{id}`  | Update a book by ID |
| DELETE | `/book/{id}`  | Delete a book by ID |


## Postman Collection

To simplify testing, you can download the Postman collection for the Bookstore API:

1. Download the [Bookstore API Postman Collection](https://github.com/ZaouiAmine/go-bookstore/blob/main/bookstore.postman_collection.json).

2. Import the collection into Postman:
    - Open Postman.
    - Click on "Import" in the top-left corner.
    - Select the downloaded `.json` file.

3. Use the imported requests to test the API endpoints.

Make sure the application is running locally at `http://localhost:9010` before testing.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
