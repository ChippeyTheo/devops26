FROM golang:1.26-alpine
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o server .
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

USER appuser

EXPOSE 9090

CMD ["./server"]