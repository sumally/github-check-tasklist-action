NAME = github-check-tasklist-action
TAG = 1.0.1

.PHONY: build
build:
	docker buildx build --push --platform linux/amd64,linux/arm64 -t "ghcr.io/sumally/${NAME}:${TAG}" .
