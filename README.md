# Sonar

A developer tool for inspecting localhost ports. Resolves Docker container names, shows clickable URLs, and provides filtering, sorting, and real-time monitoring.

```
$ sonar
PORT   PROCESS                      CONTAINER                    IMAGE             CPORT   URL
1780   proxy (traefik:3.0)          my-app-proxy-1               traefik:3.0       80      http://localhost:1780
3000   next-server (v16.1.6)                                                               http://localhost:3000
5432   db (postgres:17)             my-app-db-1                  postgres:17       5432    http://localhost:5432
6873   frontend (frontend:latest)   my-app-frontend-1            frontend:latest   5173    http://localhost:6873
9700   backend (backend:latest)     my-app-backend-1             backend:latest    8000    http://localhost:9700

5 ports (4 docker, 1 user)
```

## Features

- Docker-aware — resolves container names, images, compose services, and port mappings
- Colored output with clickable URLs
- Hides desktop apps (Figma, Discord, etc.) by default
- Configurable columns, filtering, and sorting
- JSON output for scripting
- Port forwarding, real-time watch mode, and more

## Install

Requires [Go](https://go.dev/dl/).

```sh
git clone https://github.com/rkrebs/sonar.git
cd sonar
./install.sh
```

This builds the binary and adds `~/.sonar/bin` to your PATH. Restart your terminal or run `source ~/.zshrc` to start using it.

To customize the install location:

```sh
SONAR_INSTALL_DIR=/usr/local/bin ./install.sh
```

## Usage

### List ports (default)

```sh
sonar                          # show all ports
sonar --filter docker          # only Docker ports
sonar --sort name              # sort by process name
sonar --json                   # JSON output
sonar -a                       # include desktop apps
sonar -c port,compose,image,url  # custom columns
```

Available columns: `port`, `process`, `pid`, `type`, `url`, `container`, `image`, `containerport`, `compose`, `project`, `user`, `bind`, `ip`

### Inspect a port

```sh
sonar info 6873
```

Shows full details: command, user, bind address, Docker container/image/compose info.

### Open in browser

```sh
sonar open 3000
```

### Watch for changes

```sh
sonar watch              # poll every 2s
sonar watch -i 500ms     # poll every 500ms
```

Shows the initial table, then prints diffs as ports come and go.

### Port mapping

```sh
sonar map 6873 3002
```

Makes the service on port 6873 available on port 3002. Useful for accessing Docker services on a friendlier port.

### Kill a process

```sh
sonar kill 3000           # send SIGTERM
sonar kill 3000 -f        # send SIGKILL
```

Warns if the port belongs to a Docker container and suggests `docker stop` instead.

### Global flags

```sh
sonar --no-color           # disable colored output (also respects NO_COLOR env)
```

## Supported platforms

- macOS (uses `lsof`)
- Linux (uses `ss`)
