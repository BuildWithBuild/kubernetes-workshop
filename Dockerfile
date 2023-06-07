FROM golang:1.20-alpine

WORKDIR /usr/src/kubernetes-workshop
COPY . .

RUN go build .
EXPOSE 8080

CMD ["./kubernetes-workshop"]