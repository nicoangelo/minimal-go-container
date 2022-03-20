FROM golang:1.17-alpine AS build

WORKDIR /app

COPY go.mod .
COPY *.go .
RUN CGO_ENABLED=0 go build  .

FROM scratch

COPY --from=build /app/go-docker /
CMD ["/go-docker"]
