FROM registry.cn-hangzhou.aliyuncs.com/knative-sample/golang:1.13-alpine3.10 as builder
WORKDIR /go/src/github.com/knative-sample/serving-controller
COPY cmd/ cmd
COPY pkg/ pkg
COPY vendor/ vendor
RUN CGO_ENABLED=0 go build -o controller cmd/main.go

FROM registry.cn-hangzhou.aliyuncs.com/knative-sample/alpine-sh:3.10.3
COPY --from=builder /go/src/github.com/knative-sample/serving-controller /app/bin/
ENTRYPOINT ["/app/bin/controller"]
