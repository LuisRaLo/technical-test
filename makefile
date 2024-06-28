DOCKER_COMPOSE = docker-compose

.PHONY: help
help: ## Shows the help of the commands available
	@echo "Uso: make [comando]"
	@echo ""
	@echo "Comandos disponibles:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_%-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: build
build: ## Build de los contenedores
	$(DOCKER_COMPOSE) build

.PHONY: build-frontend
build-frontend: ## Build de los contenedores
	$(DOCKER_COMPOSE) build frontend

.PHONY: build-backend
build-backend: ## Build de los contenedores
	$(DOCKER_COMPOSE) build backend

.PHONY: up
up: ## Start containers in the background
	$(DOCKER_COMPOSE) up -d

.PHONY: down
down: ## Stop and delete containers
	$(DOCKER_COMPOSE) down

.PHONY: logs
logs: ## Muestra los logs de los contenedores
	$(DOCKER_COMPOSE) logs -f

.PHONY: restart
restart: down up ## Restart containers

.PHONY: clean
clean: down ## Stop and delete containers and clean volumes and networks
	$(DOCKER_COMPOSE) rm -v -s -f
	$(DOCKER_COMPOSE) down --rmi all
	$(DOCKER_COMPOSE) down --volumes --remove-orphans


# Set the default goal to be 'help'
.DEFAULT_GOAL := help
