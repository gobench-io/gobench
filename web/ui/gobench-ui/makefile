APP_NAME := gobench-ui
DOCKER_HUB=thoverik
DOCKER_PASS=*****

dev:
	yarn run dev
build-nc: ## Build the container without caching
	docker build --no-cache -t $(APP_NAME) .
build-yarn:
	yarn build
publish:
	docker login --username $(DOCKER_HUB) --password $(DOCKER_PASS)
	docker tag $(APP_NAME) $(DOCKER_HUB)/$(APP_NAME):latest
	docker push $(DOCKER_HUB)/$(APP_NAME):latest
dockerio:	build-yarn build-nc publish