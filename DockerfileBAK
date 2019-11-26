# Build the Go API
FROM golang:latest AS builder
RUN mkdir /app
ADD . /app
# Copy go mod and sum files
COPY ./backend/go.mod ./backend/go.sum ./app/backend/
WORKDIR /app/backend

RUN go mod download
#RUN go build -o main .
RUN CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-w" -a -o /main .
# Build the React application
FROM node:alpine AS node_builder
COPY --from=builder /app/frontend ./
WORKDIR /app/frontend
COPY frontend/package.json ./
RUN npm install
COPY frontend/ ./
RUN npm run build
# Final stage build, this will be the container
# that we will deploy to production
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /main ./
COPY --from=node_builder /app/frontend/build ./web
RUN chmod +x ./main
EXPOSE 8080
CMD ./main