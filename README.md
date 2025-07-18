# Go URL Shortener

This project is a simple URL shortener built using Go. It allows users to shorten long URLs and retrieve the original URLs using the generated short links.

## Features

- Generate short URLs for long links.
- Redirect to the original URL using the short link.
- Lightweight and fast implementation.
- Easy to deploy and extend.

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/redisPlayground.git
    cd redisPlayground
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Run the application:
    ```bash
    go run main.go
    ```

## Configuration

The application uses Redis for storing URL mappings. Ensure Redis is installed and running on your system. You can configure the Redis connection in the `config.go` file.

```bash
docker pull redis
```

```bash
docker run --name redis-test-instance -p 6379:6379 -d redis
```

## API Endpoints

### Shorten URL
**POST** `/shorten`

Request:
```text
?url={originalUrl}
```

Response (path parameter):
```text
abc123
```

### Redirect to Original URL
**GET** `/redirect/abc123`

Redirects to the original URL.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Author

Frank Schweitzer