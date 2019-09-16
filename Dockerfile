# Start from the latest golang base image
FROM golang:latest as builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /main .


######## Start a new stage from scratch #######
FROM scratch
WORKDIR /
COPY --from=builder /main .
EXPOSE 8080
CMD ["./main"]