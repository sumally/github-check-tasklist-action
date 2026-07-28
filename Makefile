NAME = github-check-tasklist-action
TAG = 1.1.0

.PHONY: build
build:
	docker buildx build --push --platform linux/amd64,linux/arm64 -t "ghcr.io/sumally/${NAME}:${TAG}" .
