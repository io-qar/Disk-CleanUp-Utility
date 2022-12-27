from golang:1.19.4-alpine

run mkdir /app

copy . /app

workdir /app

run go build -o cleaning

cmd ./cleaning --config="/app/config.json"