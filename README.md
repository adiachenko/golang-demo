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

1. Require dependency

```sh
docker-compose exec app go get github.com/rs/zerolog
```

2. Import it somewhere

```sh
import (
    _ "github.com/rs/zerolog"
)
```

3. Vendor all your dependencies:

```sh
docker-compose exec app go mod vendor
```

Note that if you skip steps 2-3, new dependencies will be installed to GOPATH (even when your app is already vendored). Also, vendoring still relies on the copy of modules in `$GOPATH/pkg/mod`. If it's empty, it will start redownloading packages anew. Keep note of this when using containers.

Interestingly, you can totally skip step 1 unless you need to specify version constraints for a package. `go mod vendor` does the download for you.

There is a [pending proposal](https://github.com/golang/go/issues/30240) to allow go mod vendor accept module patterns which may or may not solve some of the issues with this workflow.

### How to install tool dependencies?

> Tool dependencies are packages not directly imported by our app but that are still used as part of overall build/test environment. CompileDaemon is an example of such dependency in this project.

Following [the official guide](https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module) we track tool dependency in a separate `tools.go` file or equivalent. See `tools/compile_daemon.go` in this project.

After installing the tool dependency, we need to manually build it into our bin directory like we did in a Quick Start guide.

### How to prune unused dependencies?

Run

```sh
docker-compose exec app go mod tidy
```

Note that this command still doesn't cleanup `go.sum` file properly. Sandly,for now there is no way to tidy it up outside of nuclear option (deleting the file and recreating it with something like `go mod download`).
