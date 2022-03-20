.PHONY: all

build:
	@docker build . -t go-docker

run:
	@docker run -it --rm go-docker

size:
	@docker images go-docker --format="{{.Size}}"