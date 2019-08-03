#!/usr/bin/env bash

bin/CompileDaemon -build="go build -o bin/http cmd/http/main.go" -command="./bin/http"
