FROM golang

WORKDIR /app
COPY . /app
RUN go build -o api.exe

EXPOSE 8080

ENTRYPOINT ["./api.exe"]
