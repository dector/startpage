# Startpage



## What Is It

Startpage serves as a custom browser homepage with intelligent search routing:
- Enter full URLs to navigate directly to them
- Use custom shortcuts (e.g., `mail` -> Gmail, `yt` -> YouTube)
- Automatically fallback to Google search for unrecognized queries
- Lightweight and fast with minimal dependencies

## Setup

### Prerequisites

- Go 1.23.0 or later
- [templ](https://templ.guide/) (for template generation)
- [Taskfile](https://taskfile.dev/)

### TL;DR Init

To setup systemd service and sample config:

```bash
task deploy
```

and open [localhost:1111](http://localhost:1111) in the browser.

### Building

```bash
task build
```

### Configuration

Create a config file at `~/.config/startpage/config.yml` to define your custom shortcuts.
Or use sample config:

```bash
task build:config
```

[config.yml example](./assets/sample.config.yml)

### Port Configuration

The application uses port **1111** by default. You can override this with the `PORT` environment variable:

```bash
# Use default port (1111)
./bin/startpage

# Use custom port
PORT=8080 ./bin/startpage

# Development mode uses port 1110
DEV=1 ./bin/startpage
```

Port priority:
1. `PORT` environment variable (if set)
2. Port 1110 (if `DEV` environment variable is set)
3. Port 1111 (default)

## Setting up Systemd Service

### 1. Install service

```bash
task deploy:service
```

### 3. Check service status

```bash
systemctl --user status startpage
```

### 4. View logs

```bash
journalctl --user -u startpage -f
```

### Disable automatic restart on boot

By default service is set to restart on boot. To disable this behavior:

```bash
systemctl --user disable startpage
```

## Development

```bash
# Install templ
go get -tool github.com/a-h/templ/cmd/templ@latest

# Run in development mode with hot reload
task dev
```

This will start the application with automatic template regeneration and reload on port 1112 (proxying to port 1110).

## Usage

1. Set your browser's homepage to `http://localhost:1111` (or your configured port)
2. Type shortcuts (e.g., `mail`, `yt`) to navigate to configured URLs
3. Type full URLs to navigate directly
4. Type anything else to search on Google
