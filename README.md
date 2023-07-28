# Todo CRUD API in Go with gin (No DB)

This project implements a simple CRUD (Create, Read, Update, Delete) API for managing todo using Go. Instead of using an actual database, it uses an in-memory struct to store todo data.

## Requirements

- Go 1.15 or higher

## Getting Started

1. Clone the repository:

```
   git clone https://github.com/your-username/todo-crud-api.git
   cd todo-crud-api
```

2. Install dependencies:

```
go mod tidy
```

3. Run the server:

```
go run main.go
```

The server should now be running at http://localhost:8080.

# Endpoints

## Get All Todo

- URL: /todos
- Method: GET
- Response: List of all Todos in JSON format.

## Get a Single Todo

- URL: /todo/{id}
- Method: GET
- Response: Todo with the specified ID in JSON format.

## Create a New Todo

- URL: /create
- Method: POST
- Request Body: todo data in JSON format (ID, Title).
- Response: Newly created Todo in JSON format.

## Update a Todo

- URL: /update/{id}
- Method: PUT
- Request Body: Updated todo data in JSON format (ID, Title).
- Response: Updated todo in JSON format.

## Delete a Todo

- URL: /delete/{id}
- Method: DELETE
- Response: Deleted todo in JSON format.

## Todo Structure

A todo is represented by the following JSON structure:

```
{
  "ID": "string",
  "Title": "string",
}

```

## Sample Requests

1. Get All Todos:

```
GET /todos
```

2. Get a Single Movie

```
GET /todo/{id}

```

3. Create a New Todo:

```
POST /create
Content-Type: application/json

{
  "ID": "todo_123",
  "Title": "New Todo Title",

}

```

4. Updated A Todo

```
PUT /todo/{id}
Content-Type: application/json

{
  "ID": "todo_123",
  "Title": "Updated Todo Title",
}

```

5. Delete a Todo:

```
DELETE /delete/{id}

```

## Error Handling

- If a todo with the given ID is not found, the server will respond with a 404 Not Found error.
- For other errors, the server will respond with a 500 Internal Server Error.

## License

This project is licensed under the MIT License. Feel free to use, modify, and distribute it as per the terms of the license.
