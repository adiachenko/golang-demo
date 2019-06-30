#!/usr/bin/env bash

CompileDaemon -build="go build -o bin/http cmd/http/main.go" -command="./bin/http"
