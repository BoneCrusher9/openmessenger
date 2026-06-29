# OpenMessenger

Modern open-source messenger written in Go.

> A fast, secure and extensible messenger inspired by classic ICQ and modern messaging platforms.

## Features

- User registration
- Authentication (JWT)
- Personal chats
- Group chats (planned)
- File uploads
- WebSocket
- Message history
- Redis cache
- PostgreSQL
- MinIO storage

## Tech Stack

### Backend

- Go
- Gin
- PostgreSQL
- Redis
- MinIO
- pgx
- Docker

### Planned clients

- Web (React)
- Desktop
- Mobile

## Project Structure

```
OpenMessenger
??? backend
??? deploy
??? docs
??? scripts
??? web
??? mobile
??? desktop
```

## Quick Start

```bash
docker compose up -d
```

```bash
cd backend
go run ./cmd/server
```

## Status

?? Under active development