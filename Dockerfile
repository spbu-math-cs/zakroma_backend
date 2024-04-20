FROM golang:alpine

ARG PORT

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

EXPOSE ${SERVER_PORT}

RUN --mount=type=cache,target="/root/.cache/go-build" go build -o server_entrypoint .
CMD ["./server_entrypoint"]
