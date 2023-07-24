FROM golang:1.16-alpine
WORKDIR /app

COPY src/go.mod ./
RUN go mod download

COPY src/*.go ./

RUN go build -o /golang

EXPOSE 8080

CMD [ "/golang" ]