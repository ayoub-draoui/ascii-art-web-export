
FROM golang:1.22.2
LABEL key="value"
WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .
# RUN go build -o main main.go
RUN go build -o main .

EXPOSE 8083

CMD [ "./main" ]