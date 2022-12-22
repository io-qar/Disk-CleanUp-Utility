from golang:1.19.4

workdir /app

env SRC_DIR=/go/src/koro/

add . $SRC_DIR

run cd $SRC_DIR; go build -o clean-utility; cp clean-utility /app/

# entrypoint ["./clean-utility --config='../config.json'"]
cmd ./clean-utility --config="$SRC_DIR/config.json"