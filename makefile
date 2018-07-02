default:
	@tmux_send "make build"

build: clean
	@go build main.go
	@./main

clean:
	@rm -rf frames
	@mkdir frames
	@mkdir -p out
