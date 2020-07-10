# FROM golang:latest AS builder

# WORKDIR /app
# COPY go.mod go.sum ./
# COPY [^client]* ./
# ENV PORT_GO $PORT
# RUN go mod download
# RUN go build .
# # Expose port 8080 to the outside world
# EXPOSE 8080
# CMD ./chat-it

# Build the Go API
FROM golang:latest AS builder

ADD . /app
WORKDIR /app
RUN rm -rf client
ENV PORT_GO $PORT
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o /main ./cmd/chat-it


# FROM node:alpine AS node_builder
# COPY --from=builder /app/client ./
# # ARG PORT
# # ARG URI
# # ENV REACT_APP_DEPLOY true
# # ENV REACT_APP_PORT $PORT
# # ENV REACT_APP_URI $URI
# ENV REACT_APP_URI "koffee.localhost"
# RUN npm install

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /main ./
COPY --from=builder /app/A_GOOGLE_CREDENTIALS.json ./
COPY --from=builder /app/.env ./
# COPY --from=node_builder /build ./build
RUN chmod +x ./main
EXPOSE 8080

CMD ./main