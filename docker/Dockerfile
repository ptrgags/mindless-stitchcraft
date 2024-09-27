FROM golang:1.23.1-bookworm AS build
WORKDIR /app
COPY . .
RUN go build -o mindless-stitchcraft main.go

FROM scratch
COPY --from=build /app/mindless-stitchcraft /test
ENTRYPOINT ["/test"]