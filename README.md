# 🧙 harry-potter

**harry-potter** (`hp`, `harry`, `potter`) is a command-line tool for managing FreeBSD jails. It provides a unified interface for the full lifecycle of jails: building and distributing images, running and inspecting jails, managing virtual networks and persistent volumes, and monitoring system-wide resource usage.

## ✨ Features

- **Familiar CLI** — identical command structure and flags to what experienced system operators already know
- **Jail lifecycle** — create, start, stop, restart, pause, kill, remove jails and inspect their state
- **Image management** — build custom jail images from a `Jailfile`, pull from registries, tag, push, save, and load
- **Virtual networking** — define virtual networks backed by `epair(4)` and `bridge(4)`, connect jails at runtime
- **Persistent volumes** — manage named volumes via `nullfs` mounts or ZFS datasets
- **Resource limits** — enforce CPU and memory quotas through FreeBSD's `rctl(8)` subsystem
- **Port forwarding** — expose jail services to the host network through `pf(4)` rules
- **Structured output** — format any listing as a table, JSON, or Go template string
- **System overview** — disk usage, live event stream, and system-wide prune

## 📋 Requirements

- FreeBSD 13.0 or later
- Root or appropriate `jail(8)` privileges
- ZFS recommended (`kern.features.zfs=1`) for image layering and snapshot support
- `pf(4)` enabled for port forwarding

## 📦 Installation

```sh
git clone https://codeberg.org/zatarain/harry-potter.git
cd harry-potter
make build
make install   # installs harry-potter and creates hp, harry, potter symlinks in /usr/local/bin
```

### 🔗 Command aliases

The binary is installed as `harry-potter`. Three symlinks pointing to the same binary are created at install time:

| Alias    | Resolves to    |
|----------|----------------|
| `hp`     | `harry-potter` |
| `harry`  | `harry-potter` |
| `potter` | `harry-potter` |

## 🚀 Quick start

```sh
# Pull a base image
hp pull freebsd:14.4-RELEASE

# Run a jail interactively, remove it on exit
hp run --rm -it freebsd:14.4-RELEASE /bin/sh

# Run a jail in the background with a name
hp run -d --name web01 --ip 10.0.0.10 freebsd:14.4-RELEASE /usr/local/sbin/nginx

# List running jails
hp ps

# Tail logs from a running jail
hp logs -f web01

# Execute a command inside a running jail
hp exec web01 pkg install -y curl

# Stop and remove the jail
hp stop web01
hp rm web01

# Build a custom image from a Jailfile
hp build -t myapp:latest .

# Show system-wide disk usage
hp system df
```

## 🛠️ Management commands

### 🔒 `jail` — manage jails

Aliases: `container`, `pot`

```
hp jail attach    Attach to a running jail's stdio
hp jail commit    Create an image from a jail's current state
hp jail cp        Copy files between host and jail
hp jail create    Create a jail without starting it
hp jail diff      Show filesystem changes in a jail
hp jail exec      Run a command in a running jail
hp jail export    Export a jail filesystem as a tar archive
hp jail inspect   Show detailed jail information
hp jail kill      Send a signal to a running jail
hp jail logs      Fetch jail log output
hp jail ls        List jails
hp jail pause     Suspend all processes in a jail
hp jail port      List port mappings for a jail
hp jail prune     Remove all stopped jails
hp jail rename    Rename a jail
hp jail restart   Restart a jail
hp jail rm        Remove one or more jails
hp jail run       Create and start a jail from an image
hp jail start     Start a stopped jail
hp jail stats     Live resource usage statistics
hp jail stop      Stop a running jail
hp jail top       Show processes running inside a jail
hp jail unpause   Resume a paused jail
hp jail update    Update resource limits of a running jail
hp jail wait      Wait for a jail to exit
```

### 🖼️ `image` — manage images

```
hp image build         Build an image from a Jailfile
hp image history       Show layer history of an image
hp image import        Import a tarball as a new image
hp image inspect       Show detailed image information
hp image load          Load an image from a tar archive or stdin
hp image ls            List locally available images
hp image prune         Remove unused images
hp image pull          Pull an image from a registry
hp image push          Push an image to a registry
hp image rm            Remove one or more images
hp image save          Save images to a tar archive
hp image tag           Create a new tag for an image
```

### 🌐 `network` — manage networks

```
hp network connect     Connect a jail to a network
hp network create      Create a virtual network
hp network disconnect  Disconnect a jail from a network
hp network inspect     Show detailed network information
hp network ls          List networks
hp network prune       Remove unused networks
hp network rm          Remove one or more networks
```

### 💾 `volume` — manage volumes

```
hp volume create       Create a named volume
hp volume inspect      Show detailed volume information
hp volume ls           List volumes
hp volume prune        Remove unused volumes
hp volume rm           Remove one or more volumes
```

### ⚙️ `system` — system-wide operations

```
hp system df           Show disk usage by images, jails, and volumes
hp system events       Stream real-time jail events
hp system info         Display host and runtime information
hp system prune        Remove all unused resources
```

## ⚡ Top-level shorthands

Every frequently used command is also available directly at the top level, without the management group prefix:

```
hp attach      hp commit      hp cp          hp create
hp diff        hp events      hp exec        hp export
hp history     hp images      hp import      hp info
hp inspect     hp kill        hp load        hp login
hp logout      hp logs        hp pause       hp port
hp ps          hp pull        hp push        hp rename
hp restart     hp rm          hp rmi         hp run
hp save        hp search      hp start       hp stats
hp stop        hp tag         hp top         hp unpause
hp update      hp wait        hp build
```

## 📄 Jailfile

A `Jailfile` defines the steps to build a jail image, analogous to a `Dockerfile`:

```dockerfile
FROM freebsd:14.4-RELEASE

RUN pkg install -y nginx

COPY nginx.conf /usr/local/etc/nginx/nginx.conf

EXPOSE 80/tcp

CMD ["/usr/local/sbin/nginx", "-g", "daemon off;"]
```

## 🔧 Configuration

harry-potter reads `~/.harry-potter/config.yaml` on startup. A minimal example:

```yaml
data_dir: /var/db/harry-potter
log_dir: /var/log/harry-potter

zfs:
  enabled: true
  pool: zroot
  dataset: zroot/jails

network:
  default_interface: em0
  default_subnet: 10.0.0.0/24

templates:
  dir: /var/db/harry-potter/templates
  mirror: https://download.freebsd.org/releases
```

Pass a custom path with the global `--config` flag.

## 🧑‍💻 Development

```sh
make build   # compile to bin/harry-potter
make tidy    # go mod tidy
make fmt     # gofmt
make vet     # go vet
make lint    # golangci-lint (requires golangci-lint in PATH)
make test    # go test ./...
```

## 📜 License

See [LICENSE](LICENSE).
