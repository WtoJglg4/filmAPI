FROM golang:latest

WORKDIR /app
COPY ./ ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /main

# RUN go build -o main .

# RUN chmod +x main
CMD [ "./main" ]