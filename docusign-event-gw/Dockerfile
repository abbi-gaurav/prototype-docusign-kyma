FROM golang:1.15 as builder

ENV GO111MODULE=on

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY main.go                .
COPY internal               ./internal

RUN ls /app/
RUN CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o docusign-event-gw .

FROM scratch
WORKDIR /app
COPY --from=builder /app/docusign-event-gw /app/

EXPOSE 8080
ENTRYPOINT ["/app/docusign-event-gw"]