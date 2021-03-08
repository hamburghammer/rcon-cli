# rcon-cli
A little RCON cli written in golang.

This is a fork from [itzg/rcon-cli](https://github.com/itzg/rcon-cli) with following extra features:
- Provied a smaller binary (ca. 50% smaller).
- Replace base RCON lib from [james4k/rcon](https://github.com/james4k/rcon) with [hamburghammer/rcon](https://github.com/hamburghammer/rcon).
- Remove `config.yml` support.
- Change the `Dockerfile` to have build support.

## Installation
### From Source
Clone the repository and install it with `go install` (requires working `go` installation)

### Docker/Podman
Clone the repository and use docker/podman to build a image with the executable `docker build -t hamburghammer/rcon-cli .`
Start the image `docker run hamburghammer/rcon-cli -h`


## Usage

```text
rcon-cli is a CLI to interact with a RCON server.
It can be run in an interactive mode or to execute a single command.

USAGE:
	rcon-cli [FLAGS] [RCON command ...]
	
FLAGS:
  -h, --help              Prints this help message and exits.
      --host string       RCON server's hostname. (default "localhost")
      --password string   RCON server's password.
      --port string       RCON server's port. (default "25575")

ENVIRONMENT VARIABLE:
	All flags can be set through the flag name in capslock with the RCON_CLI_ prefix (see examples).
	Flags have allways priority over env vars!

EXAMPLES:
	rcon-cli --host 127.0.0.1 --port 25575
	rcon-cli --password admin123 stop
	RCON_CLI_PORT=25575 rcon-cli stop
```
