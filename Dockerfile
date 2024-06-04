# Build the multicluster-mesh-addon binary
FROM golang:1.19 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY apis/ apis/
COPY pkg/ pkg/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o multicluster-mesh-addon main.go

FROM alpine:3.13

WORKDIR /
COPY --from=builder /workspace/multicluster-mesh-addon .
USER 65532:65532

CMD ["/multicluster-mesh-addon"]
