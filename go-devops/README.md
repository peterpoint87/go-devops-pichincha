# DevOps Technical Assessment - Pedro Mendez

## Requirements
- Docker
- Docker Compose
- Go 1.22+

## Running the Application

```bash
docker-compose up --build
```

The app will be available through Kong on:

```
http://localhost:8000/DevOps
```

## Testing the Endpoint

```bash
curl -X POST \
-H "X-Parse-REST-API-Key: 2f5ae96c-b558-4c7b-a590-a501ae1c3f6c" \
-H "Content-Type: application/json" \
-d '{"message":"This is a test","to":"Juan Perez","from":"Rita Asturia","timeToLifeSec":45}' \
http://localhost:8000/DevOps
```

## Output

```
{
  "message": "Hello Juan Perez your message will be send"
}
```

Other HTTP methods will return "ERROR".

---
Developed by Pedro Mendez


## Running Tests

```bash
go test ./...
```

## GitHub Actions

Every push to main or develop triggers:

- Build
- Lint
- Test
- Binary compilation