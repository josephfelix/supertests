DOCKERCOMPOSE ?= $(shell which docker-compose)
XDGOPEN	?= $(shell which xdg-open)

run:
	$(XDGOPEN) http://127.0.0.1
	$(DOCKERCOMPOSE) up -d
bash:
	$(DOCKERCOMPOSE) exec app bash
build:
	$(DOCKERCOMPOSE) exec app buffalo build
down:
	$(DOCKERCOMPOSE) down