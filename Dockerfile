#FROM golang
FROM golang:latest as builder

#ENV GO111MODULE=on

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main .

EXPOSE 8200
CMD ["./main"]