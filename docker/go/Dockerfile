FROM golang:1.25.0

RUN useradd -m appuser

WORKDIR /app

USER appuser

CMD ["go", "run", "main.go"]
