FROM golang:alpine

ARG PORT

WORKDIR /build

COPY ./ /build

EXPOSE ${SERVER_PORT}

RUN go build -o server_entrypoint .
CMD ["./server_entrypoint"]