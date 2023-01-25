# Disk-CleanUp-Utility
Утилита, очищающая папки, указанные в вашем `config.json`.

# Установка и запуск
## Docker
1. `docker build -t cleaning .`
2. `docker run --name cleaning -it cleaning`
3. Переименуйте `example.config.json` в `config.json`
4. Отредактируйте `config.json`
## Локально
1. Переименуйте `example.config.json` в `config.json`
2. Отредактируйте `config.json`
3. `cd internal/app`
4. `go build -o clean-utility`
5. `./clean-utility --config="<путь к конфиг. файлу>"`