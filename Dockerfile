FROM golang

COPY . /app
RUN cd /app
RUN make

EXPOSE 8080

ENTRYPOINT ["./api.exe"]
