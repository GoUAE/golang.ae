default: run
gen: templates

run:
    go run main.go

static:
    go run main.go -static

dev:
    templ generate --watch --proxy="http://localhost:8080" --cmd="go run ."

deploy:
    templ generate && go build && sudo systemctl restart gouae.service

templates:
    templ generate

