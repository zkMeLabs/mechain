FROM golang:1.22.4-bullseye AS build-env

WORKDIR /go/src/github.com/zkmelabs

ENV CGO_CFLAGS="-O -D__BLST_PORTABLE__"
ENV CGO_CFLAGS_ALLOW="-O -D__BLST_PORTABLE__"

COPY . .

RUN make build

FROM golang:1.22.4-bullseye

RUN apt-get update -y
RUN apt-get install ca-certificates jq -y

WORKDIR /root

COPY --from=build-env /go/src/github.com/zkmelabs/build/mechaind /usr/bin/mechaind

CMD ["mechaind"]
