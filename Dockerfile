FROM golang:1.10
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM scratch
COPY --from=0 /go/src/app/app ./
ENTRYPOINT [ "/app" ]