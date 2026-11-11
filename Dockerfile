FROM golang:1.26-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG GIT_COMMIT=unknown
RUN CGO_ENABLED=0 go build -ldflags "-X $(go list -m)/internal/version.GitCommit=${GIT_COMMIT}" -o /app .

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=builder /app /app
ENTRYPOINT ["/app"]
CMD ["server"]
