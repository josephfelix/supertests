DOCKERCOMPOSE ?= $(shell which docker-compose)
XDGOPEN	?= $(shell which xdg-open)
PYTHON ?= $(shell which python)
NAVIGATOR ?= $(if $(PYTHON),$(PYTHON) -m webbrowser,$(XDGOPEN))

setup: reset run
reset:
	sudo rm -rf node_modules
run:
	$(NAVIGATOR) http://127.0.0.1:1337
	$(DOCKERCOMPOSE) up --build
bash:
	$(DOCKERCOMPOSE) exec app bash
build:
	$(DOCKERCOMPOSE) exec app buffalo build
db:
	$(DOCKERCOMPOSE) exec db psql supertests supertests
down:
	$(DOCKERCOMPOSE) down