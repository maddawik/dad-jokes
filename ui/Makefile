.PHONY: tailwind templ dev build clean

tailwind:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --minify

templ:
	templ generate

dev:
	air

build: tailwind templ
	go build -o ./bin/main ./cmd/main.go

clean:
	rm -r ./tmp ./bin
