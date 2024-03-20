FROM golang:1.20-alpine AS builder

WORKDIR /usr/local/src/

RUN apk --no-cache add bash git make gcc gettext musl-dev

# dependences
COPY ["go.mod", "go.sum", "./"]

ENV GO111MODULE=on

RUN go mod download

# build
COPY . ./
RUN go build -o ./bin/main cmd/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/main /
# COPY docs/ schema/ .env configs/ /
COPY docs ./docs/
COPY schema ./schema/
COPY .env /
COPY configs ./configs/

CMD [ "/main" ]