FROM golang:1.23 AS builder
COPY . /source/
WORKDIR /source
RUN CGO_ENABLED=0 go build -o norrdns .

FROM ubuntu:24.04
EXPOSE 53/udp
COPY --from=builder /source/norrdns /
RUN apt update && apt dist-upgrade -y && apt install -y ca-certificates && rm -rf /var/cache/apt /var/lib/apt/lists/
CMD [ "/norrdns" ]