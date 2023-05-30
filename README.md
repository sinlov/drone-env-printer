[![go-ubuntu](https://github.com/sinlov/drone-env-printer/workflows/go-ubuntu/badge.svg?branch=main)](https://github.com/sinlov/drone-env-printer/actions)
[![GoDoc](https://godoc.org/github.com/sinlov/drone-env-printer?status.png)](https://godoc.org/github.com/sinlov/drone-env-printer/)
[![GoReportCard](https://goreportcard.com/badge/github.com/sinlov/drone-env-printer)](https://goreportcard.com/report/github.com/sinlov/drone-env-printer)
[![codecov](https://codecov.io/gh/sinlov/drone-env-printer/branch/main/graph/badge.svg)](https://codecov.io/gh/sinlov/drone-env-printer)
[![docker version semver](https://img.shields.io/docker/v/sinlov/drone-env-printer?sort=semver)](https://hub.docker.com/r/sinlov/drone-env-printer/tags?page=1&ordering=last_updated)
[![docker image size](https://img.shields.io/docker/image-size/sinlov/drone-env-printer)](https://hub.docker.com/r/sinlov/drone-env-printer)
[![docker pulls](https://img.shields.io/docker/pulls/sinlov/drone-env-printer)](https://hub.docker.com/r/sinlov/drone-env-printer/tags?page=1&ordering=last_updated)
[![github release](https://img.shields.io/github/v/release/sinlov/drone-env-printer?style=social)](https://github.com/sinlov/drone-env-printer/releases)

## for what

- this project used to drone CI

## Pipeline Settings (.drone.yml)

`1.x`

```yaml
steps:
  - name: drone-env-printer
    image: sinlov/drone-env-printer:latest
    pull: if-not-exists
    settings:
      debug: false
      env_printer_print_keys:
        - GOPATH
        - GOBIN
```

- full config

```yaml
steps:
  - name: drone-env-printer
    image: sinlov/drone-env-printer:latest
    pull: if-not-exists
    settings:
      debug: false
      env_printer_print_keys:
        - GOPATH
        - GOBIN
      env_printer_padding_left_max: 42
```
- `1.x` drone-exec only support env

- download by [https://github.com/sinlov/drone-env-printer/releases](https://github.com/sinlov/drone-env-printer/releases) to get platform binary, then has local path
- binary path like `C:\Drone\drone-runner-exec\plugins\drone-env-printer.exe` can be drone run env like `EXEC_DRONE_ENV_PRINTER_PLUGIN_FULL_PATH`
- env:EXEC_DRONE_ENV_PRINTER_PLUGIN_FULL_PATH can set at file which define as [DRONE_RUNNER_ENVFILE](https://docs.drone.io/runner/exec/configuration/reference/drone-runner-envfile/) to support each platform

```yaml
steps:
  - name: drone-env-printer-exec # must has env EXEC_DRONE_ENV_PRINTER_PLUGIN_FULL_PATH and exec tools
    environment:
      PLUGIN_DEBUG: false
      PLUGIN_ENV_PRINTER_PRINT_KEYS: "GOPATH,GOBIN"
      PLUGIN_ENV_PRINTER_PADDING_LEFT_MAX: 42
```

# Features

- more see [features/README.md](features/README.md)

# dev

## depends

in go mod project

```bash
# warning use private git host must set
# global set for once
# add private git host like github.com to evn GOPRIVATE
$ go env -w GOPRIVATE='github.com'
# use ssh proxy
# set ssh-key to use ssh as http
$ git config --global url."git@github.com:".insteadOf "https://github.com/"
# or use PRIVATE-TOKEN
# set PRIVATE-TOKEN as gitlab or gitea
$ git config --global http.extraheader "PRIVATE-TOKEN: {PRIVATE-TOKEN}"
# set this rep to download ssh as https use PRIVATE-TOKEN
$ git config --global url."ssh://github.com/".insteadOf "https://github.com/"

# before above global settings
# test version info
$ git ls-remote -q https://github.com/sinlov/drone-env-printer.git

- test code

add env then test

```bash
export PLUGIN_MSG_TYPE=post \
  export PLUGIN_WEBHOOK=7138d7b3-abc
```

```bash
make test
```

- see help

```bash
make dev
```

update main.go file set env then and run

```bash
export PLUGIN_MSG_TYPE= \
  export PLUGIN_WEBHOOK= \
  export DRONE_REPO=sinlov/drone-env-printer \
  export DRONE_REPO_NAME=drone-env-printer \
  export DRONE_REPO_NAMESPACE=sinlov \
  export DRONE_REMOTE_URL=https://github.com/sinlov/drone-env-printer \
  export DRONE_REPO_OWNER=sinlov \
  export DRONE_COMMIT_AUTHOR=sinlov \
  export DRONE_COMMIT_AUTHOR_AVATAR=  \
  export DRONE_COMMIT_AUTHOR_EMAIL=sinlovgmppt@gmail.com \
  export DRONE_COMMIT_BRANCH=main \
  export DRONE_COMMIT_LINK=https://github.com/sinlov/drone-env-printer/commit/68e3d62dd69f06077a243a1db1460109377add64 \
  export DRONE_COMMIT_SHA=68e3d62dd69f06077a243a1db1460109377add64 \
  export DRONE_COMMIT_REF=refs/heads/main \
  export DRONE_COMMIT_MESSAGE="mock message commit" \
  export DRONE_STAGE_STARTED=1674531206 \
  export DRONE_STAGE_FINISHED=1674532106 \
  export DRONE_BUILD_STATUS=success \
  export DRONE_BUILD_NUMBER=1 \
  export DRONE_BUILD_LINK=https://drone.xxx.com/sinlov/drone-env-printer/1 \
  export DRONE_BUILD_EVENT=push \
  export DRONE_BUILD_STARTED=1674531206 \
  export DRONE_BUILD_FINISHED=1674532206
```

- then run

```bash
make run
```

## docker

```bash
# then test build as test/Dockerfile
$ make dockerTestRestartLatest
# if run error
# like this error
# err: missing webhook, please set webhook
#  fix env settings then test

# see run docker fast
$ make dockerTestRunLatest

# clean test build
$ make dockerTestPruneLatest

# see how to use
$ docker run --rm sinlov/drone-env-printer:latest -h
```
