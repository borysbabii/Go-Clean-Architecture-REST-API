# Initial stage: download modules
FROM golang:1.20.5-alpine as builder

ENV config=docker

WORKDIR /app

COPY ./ /app

RUN go mod download


# Intermediate stage: Build the binary
FROM golang:1.20.5-alpine as runner

# Add Open Container labels
LABEL org.opencontainers.image.version="1.0.0"
LABEL org.opencontainers.image.title="BudgetWhiz"
LABEL org.opencontainers.image.description="Personal Accounting Application"
LABEL org.opencontainers.image.source="https://github.com/borysbabii/Go-Clean-Architecture-REST-API"
LABEL org.opencontainers.image.vendor="Borys Babii"
LABEL org.opencontainers.image.licenses="GPL-3.0"

COPY --from=builder ./app ./app

RUN go install github.com/githubnemo/CompileDaemon@latest

WORKDIR /app
ENV config=docker

EXPOSE 5000
EXPOSE 5555
EXPOSE 7070

ENTRYPOINT CompileDaemon --build="go build cmd/api/main.go" --command=./main




