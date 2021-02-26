FROM golang:1.15-alpine as builder

# build-base is a meta-package that will install the GCC, libc-dev and binutils packages (amongst others)
RUN apk add build-base

# working directory inside container. default is /go/src/app
WORKDIR /go/src/cgo-playground

# copy project to container and build
COPY . /go/src/cgo-playground
RUN go build -o /bin/cgo-playground ./cmd

# build the smallest possible image
FROM scratch
COPY --from=builder /bin/cgo-playground /bin/cgo-playground

# run container
ENTRYPOINT ["/bin/cgo-playground"]
