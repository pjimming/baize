SHELL=/bin/zsh

api:
	goctl api go -api core.api -dir . --home ../template
	goctl api format --dir ./apis -declare

generate:
	@rm -f core-yaml.go baize
	go generate
	go build -o baize
	@mv baize ../../