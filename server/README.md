# Titly API Documentation

Base URL: `http://localhost:4000`

## `GET /`

Returns a welcome message.

### Response — `200 OK`

```json
{
  "message": "Welcome to Titly!"
}
````

---

## `POST /create-short-url`

Creates a short URL for a given long URL.

If the long URL already exists, the existing short URL is returned.

### Request

```json
{
  "long_url": "https://example.com"
}
```

### Response — `200 OK`

```json
{
  "id": 1,
  "long_url": "https://example.com",
  "short_url": "abc12345"
}
```

If the URL already exists:

```json
{
  "short_url": "abc12345"
}
```

### Response — `400 Bad Request`

Returned when the request body is invalid or `long_url` is missing.

```json
{
  "error": "..."
}
```

---

## `GET /:short-url`

Redirects a short URL to its original long URL.

### Example

```http
GET /abc12345
```

### Response — `302 Found`

Redirects to:

```text
https://example.com
```

The short URL is first looked up in Redis. On a cache miss, SQLite is checked and the mapping is stored in Redis for 24 hours.

### Response — `404 Not Found`

```json
{
  "error": "Short URL not found"
}
```

### Response — `500 Internal Server Error`

```json
{
  "error": "Failed to store mapping in cache"
}
```