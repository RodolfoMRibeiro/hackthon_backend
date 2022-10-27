FROM golang:1.19.1

WORKDIR /hackthon/app

COPY . . 

RUN go mod tidy
RUN go build

EXPOSE 8080

CMD [ "./hackthon" ]