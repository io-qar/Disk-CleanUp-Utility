# Disk-CleanUp-Utility
Утилита, очищающая папки, указанные в вашем `config.json`.

# Сборка и запуск
## Docker
1. `docker build -t cleaning .`
2. `docker run --name cleaning -it cleaning`
3. Переименуйте `example.config.json` в `config.json`
4. Отредактируйте `config.json`
## Из-под windows для macOS через Docker
1. Запуск Docker контейнера (cmd): `docker run -v %cd%:/builder -w /builder -i -t golang:1.19 bash`
2. `CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o diskclean.app ./app.go`
3. Файл `diskclean` перенести на целевую машуну вместе с конфиг. файлом
4. Переименуйте `example.config.json` в `config.json`
5. Отредактируйте `config.json`
## Локально
1. Переименуйте `example.config.json` в `config.json`
2. Отредактируйте `config.json`
3. `go build -o clean-utility`
4. `./clean-utility --config="<путь до конфиг. файла>"`