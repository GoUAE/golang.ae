gen:
    templ generate

dev:
    templ generate --watch --proxy="http://localhost:8080" --cmd="go run ."

static: 
    go run main.go --static
