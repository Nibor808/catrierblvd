### Using hot reload with air

```
go install github.com/cosmtrek/air@latest

docker run -it --rm \
    -w $(pwd) \
    -e "air_wd=$(pwd)" \
    -v $(pwd) \
    -p 3000:3000 \
    cosmtrek/air

air init

air -c .air.toml
```
