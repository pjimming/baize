SHELL=/bin/zsh

api:
	goctl api go -api core.api -dir . --home ../template
	goctl api format --dir ./apis -declare

generate:
	@rm -f core-yaml.go
	go generate
	rm -f baize && go build -o baize