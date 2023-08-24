SHELL=/bin/zsh

generate:
	# @cd frontend && npm run build
	@rm -rf backend/core/static backend/core/statik
	@cp -r frontend/dist backend/core/static
	@cd backend/core && make generate

print:
	@echo $(GOPATH)
	@echo $(PATH)