FROM docker.io/golang:1.16 AS build
COPY ./go.* /src/
WORKDIR /src
RUN go mod download

COPY . /src

RUN CGO_ENABLED=0 go build -o /rcon-cli


FROM scratch
COPY --from=build /rcon-cli /

ENTRYPOINT [ "/rcon-cli" ]