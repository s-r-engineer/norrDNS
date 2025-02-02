FROM golang:1.23 AS builder
COPY . /source/
WORKDIR /source
RUN CGO_ENABLED=0 go build -o norrdns .

FROM ubuntu
EXPOSE 53/udp
COPY --from=builder /source/norrdns /
CMD [ "/norddns" ]
CMD [ "/norrdns" ]