# NOTICE: this build script assumes that "make apiserver" has been alread executed

FROM golang as builder
WORKDIR /builder
COPY ./go.mod ./go.mod 
COPY ./go.sum ./go.sum 
COPY ./api/ ./api/
COPY ./pkg/ ./pkg/
COPY ./cmd/ ./cmd/
RUN CGO_ENABLED=0 go build -o server cmd/finder_web/main.go

FROM alpine
WORKDIR /opt/
COPY ./assets/*.gz ./assets/
COPY ./web/ ./web/ 
EXPOSE 80/tcp
COPY --from=builder /builder/server .
ENTRYPOINT ["/opt/server", "-port=80", "-flights_data=./assets", "-web_data=./web"]
