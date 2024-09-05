FROM golang:1.22.4-bullseye AS build-env
WORKDIR /mechain
COPY . .
RUN make build

FROM golang:1.22.4-bullseye
RUN apt-get update -y && apt-get install ca-certificates jq -y
COPY --from=build-env /mechain/build/mechaind /usr/bin/mechaind
CMD ["mechaind"]