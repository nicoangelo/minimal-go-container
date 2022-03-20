# Small container images

Small container images are a reality. This repo shows that without extra effort you can

- run go binaries in an empty image, built [`FROM sratch`](https://hub.docker.com/_/scratch/)
- have an image size as low as 6MB

## How build

### Prerequisites:

- Docker
- Make

### Build

```sh
make build
```

The go binary is built with `CGO_ENABLED=0` to force a statically linked binary.

### Get image size

```sh
make size
```

will yield output similar to

```sh
6.07MB
```
