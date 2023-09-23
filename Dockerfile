FROM golang:1.21-bookworm AS builder

# # install nodejs
ENV NODE_MAJOR=20
RUN apt-get update && apt-get install -y ca-certificates curl gnupg
RUN mkdir -p /etc/apt/keyrings
RUN curl -fsSL https://deb.nodesource.com/gpgkey/nodesource-repo.gpg.key | gpg --dearmor -o /etc/apt/keyrings/nodesource.gpg
RUN echo "deb [signed-by=/etc/apt/keyrings/nodesource.gpg] https://deb.nodesource.com/node_$NODE_MAJOR.x nodistro main" | tee /etc/apt/sources.list.d/nodesource.list
RUN apt-get update && apt-get install nodejs -y

WORKDIR /app
COPY . .
RUN make deps
RUN make

FROM python:3-bookworm AS runner

# install nodejs
ENV NODE_MAJOR=20
RUN apt-get update && apt-get install -y ca-certificates curl gnupg
RUN mkdir -p /etc/apt/keyrings
RUN curl -fsSL https://deb.nodesource.com/gpgkey/nodesource-repo.gpg.key | gpg --dearmor -o /etc/apt/keyrings/nodesource.gpg
RUN echo "deb [signed-by=/etc/apt/keyrings/nodesource.gpg] https://deb.nodesource.com/node_$NODE_MAJOR.x nodistro main" | tee /etc/apt/sources.list.d/nodesource.list
RUN apt-get update && apt-get install nodejs -y
RUN npm i -g concurrently

WORKDIR /app
COPY --from=builder /app/EnvMonitoringDashboard.exe /app/EnvMonitoringDashboard.exe
COPY db.sqlite /app/db.sqlite
COPY requirements.txt /app/requirements.txt
COPY model /app/model
COPY main.py /app/main.py
RUN pip install -r requirements.txt

CMD ["concurrently", "--kill-others", "\"./EnvMonitoringDashboard.exe\"", "\"uvicorn main:app --port 5000 --host 0.0.0.0\""]
