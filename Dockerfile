FROM golang:1.16-alpine

WORKDIR /app/todos

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux go build -o /todoGo

EXPOSE 8080

CMD [ "/todoGo" ]