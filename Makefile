NPM ?= $(shell which npm)
YARN ?= $(shell which yarn)
PKG_MANAGER ?= $(if $(YARN),$(YARN),$(NPM) run)
BUFFALO ?= $(shell which buffalo)
DOCKERCOMPOSE ?= $(shell which docker-compose)
XDGOPEN	?= $(shell which xdg-open)

setup:
	$(DOCKERCOMPOSE) up -d
	$(BUFFALO) db migrate
	$(PKG_MANAGER) install
dev:
	$(DOCKERCOMPOSE) up -d
	$(XDGOPEN) http://127.0.0.1:3000
	$(BUFFALO) dev
build:
	$(BUFFALO) build
down:
	$(DOCKERCOMPOSE) down