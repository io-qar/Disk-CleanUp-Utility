# Disk-CleanUp-Utility
This utility cleans up folders written in `config.json`

# Setup & Run
## Docker (skip for local setup)
1. `docker build -t cleaning .`
2. `docker run --name cleaning -it cleaning`
3. Rename `example.config.json` to `config.json`
4. Edit `config.json` as you need to
## Local
1. Rename `example.config.json` to `config.json`
2. Edit `config.json` as you need to
3. `go build -o clean-utility`
4. `./clean-utility --config="./config.json"` (you can edit the path to your config file with the `--config` flag)


## Сборка проекта

> примеры сборки проекта 

### из windows для macOS через докер

запуск докер контейнера (cmd):
```
docker run -v %cd%:/builder -w /builder -i -t golang:1.19 bash
```

в контенере выполняем команду:
```
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o diskclean.app ./app.go
```

в итоге получаем файл `diskclean` который необходимо перенести на целевую машуну вместе с конфигом