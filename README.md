# Vesta

## Development

```sh
docker-compose -f docker-compose.dev.yaml up -d
PORT=8080 DB_PORT=9999 DB_HOST=0.0.0.0 DB_USERNAME=username DB_PASSWORD=password DB_NAME=vesta go run .
```