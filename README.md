# Cursor-Based Pagination API

This repository provides an example implementation of cursor-based pagination using Golang, Gin, and Entgo. Cursor-based pagination is a technique used to paginate through large datasets efficiently by using a pointer (cursor) to indicate the position in the dataset, rather than using offset and limit.

## Features

- **Cursor-Based Pagination**: Efficiently paginate through large datasets.
- **Compound Sorting**: Supports sorting by multiple fields to handle tie-breaking scenarios.
- **Dynamic Sorting**: Allows specifying sort fields dynamically via query parameters.
- **Gin Framework**: Uses the Gin web framework for handling HTTP requests.
- **Entgo ORM**: Utilizes Entgo for database operations.

## API Endpoints

### Get Users

Retrieves a paginated list of users.

- **URL**: `/users`
- **Method**: `GET`
- **Query Parameters**:
  - `cursor` (optional): The base64-encoded cursor to retrieve the next set of results.
  - `limit` (optional): The number of users to retrieve. Defaults to 10 if not specified.
  - `sort` (optional): The fields to sort by, separated by commas. Supports compound sorting. Default is `id`.

#### Example Request

```sh
curl -X GET "http://localhost:8080/users?sort=salary,id&limit=10"

{
  "data": [
    {
      "id": 1,
      "name": "Alice",
      "salary": 35000
    },
    {
      "id": 2,
      "name": "Bob",
      "salary": 40000
    },
    // ... more users
  ],
  "next_cursor": "aWQ9MTJ8c2FsYXJ5PTgwMDAw"
}
```