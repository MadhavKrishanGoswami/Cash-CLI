
FROM golang:1.23.1-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o cash-cli
CMD [ "./cash-cli" ]







