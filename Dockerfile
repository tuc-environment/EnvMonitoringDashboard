FROM golang

WORKDIR /app
COPY . /app
RUN make

EXPOSE 8080

ENTRYPOINT ["./api.exe"]
