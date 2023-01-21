
#  Build stage 
FROM golang:1.18.2-alpine3.16 AS builder 
WORKDIR /app 
COPY . . 
RUN go build -o main main.go 

# Run stage 
FROM alpine:3.16
RUN mkdir -p /app/env
WORKDIR /app 
COPY --from=builder /app/main /app/ 
COPY --from=builder /app/env/.env /app/env/.env

# this shoud not be run in production
# COPY app.env . 

EXPOSE 8080 
CMD ["/app/main"]
# ENTRYPOINT [ "./app/main" ]