FROM golang:1.22 as Builder
WORKDIR /app
COPY go.mod ./
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app 


# RUNTIME stage

FROM scratch
COPY --from=Builder /app/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]