# Go Demo Project

## Features

- Does not require to be inside `GOPATH`
- Go runtime is packaged in Docker Compose
- Multiple Binaries are supported (see `cmd/` dir)
- Package management with vendoring using Go Modules
- Restart web server on code changes
- Obligatory [cobra](https://github.com/spf13/cobra) and [chi](https://github.com/go-chi/chi) recommendation


## Quick Start

```sh
docker-compose up -d
docker-compose exec app go mod vendor
docker-compose exec app go build -o bin/CompileDaemon github.com/githubnemo/CompileDaemon
chmod +x scripts/build.sh scripts/serve.sh
```

We are ready to build our app now (output is set to `/bin` folder):

```sh
docker-compose exec app scripts/build.sh
```

Run web server:

```sh
docker-compose exec app bin/http
```

Run console:

```sh
docker-compose exec app bin/console
```

Serve app in development mode (restart web server on file changes):

```sh
docker-compose exec app scripts/serve.sh
```

## Managing dependencies

### How to install app dependencies

1. First, run

```sh
docker-compose exec app go get github.com/rs/zerolog
```

2. Then import the dependency

```sh
import (
    _ "github.com/rs/zerolog"
)
```

3. Lastly, vendor your dependencies (again):

```sh
docker-compose exec app go mod vendor
```

If you skip steps 2-3, new dependencies will be installed to GOPATH regardless of whether your app is already vendored or not. So make sure to follow this to a letter.

### How to install tool dependencies?

> Tool dependencies are not imported by our app but are still used as part of build/test environment. CompileDaemon is an example of such dependency in this project.

Following [the official guide](https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module) we track tool dependency in a separate `tools.go` file or equivalent. See `tools/compile_daemon.go` in this project. 

After installing the tool dependency, we need to manually build it into our bin directory like we did in a Quick Start guide.

### How to prune unused dependencies?

Run

```sh
docker-compose exec app go mod tidy
```
