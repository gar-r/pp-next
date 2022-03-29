FROM golang:latest

WORKDIR /pp-next

COPY . .

RUN go build

CMD [ "./ppnext" ]
