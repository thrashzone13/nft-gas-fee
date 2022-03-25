FROM golang:1.17-alpine AS server
ENV HTTPS_PROXY="https://free.shecan.ir/dns-query"
WORKDIR /app
COPY . ./
# Install dependencies
RUN go mod download && \
  # Build the app
  GOOS=linux GOARCH=amd64 go build -o main && \
  # Make the final output executable
  chmod +x ./main

FROM node:17-alpine AS app
ENV NODE_ENV production
WORKDIR /app
COPY ./app/package.json .
COPY ./app/package-lock.json .
RUN npm ci
COPY ./app .
RUN npm run build

FROM alpine:latest
RUN apk --no-cache add bash
WORKDIR /app
COPY --from=server /app/main .
COPY --from=app /app/build ./app/build/
CMD ["./main"]
EXPOSE 8000