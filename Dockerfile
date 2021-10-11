
FROM golang:1.16-alpine


WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /exe

EXPOSE 2204

CMD [ "/exe" ]