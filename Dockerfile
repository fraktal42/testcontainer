FROM golang:alpine AS build
WORKDIR /build
COPY . .

RUN go mod download
RUN go build -o app

FROM scratch
COPY --from=build /build/app app

ENTRYPOINT ["/app"]
