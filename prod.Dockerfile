# Build AdonisJS
FROM node:16-alpine AS buildjs
# Set directory for all files
WORKDIR /app
# Copy over package.json files
COPY package*.json ./
# Install all packages
# RUN npm install
# Copy over source code
COPY . .
# Expose port to outside world
# Start server up
CMD npm install && npm run prod

FROM golang:1.18beta1-alpine AS build
RUN apk --no-cache add gcc g++ make git
WORKDIR /go/src/app
COPY . .
RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./main ./main.go

# final stage
FROM alpine:3.13
RUN apk --no-cache add ca-certificates
WORKDIR /go
COPY --from=build /go/src/app .
COPY --from=buildjs /app/public .
EXPOSE 8080

ENTRYPOINT /go/main 