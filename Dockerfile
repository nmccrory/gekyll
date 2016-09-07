FROM golang

RUN mkdir -p /app
WORKDIR /app

ADD . /app

RUN go get github.com/gorilla/pat
RUN go install github.com/gorilla/pat
RUN go build ./server.go

EXPOSE 8080

CMD ["./server"]