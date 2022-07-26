include .make/help.mk
include .make/docker.mk
include .make/app.mk
include .make/test.mk

PROJECT_NAME ?= coding-dojo-api-golang
VERSION ?= latest
DOCKER_REGISTRY ?= ghcr.io
DOCKER_REPOSITORY ?= valentinlutz