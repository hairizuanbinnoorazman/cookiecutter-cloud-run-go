FROM golang:1.12 as builder
WORKDIR /go/src/{{cookiecutter.golang_mod_name}}
ADD . .
RUN go get
RUN go build -o app .

FROM ubuntu:18.10 as prod
COPY --from=builder /go/src/{{cookiecutter.golang_mod_name}}/app /usr/bin/app
EXPOSE 8080
CMD ["/usr/bin/app"]
