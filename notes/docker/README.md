# Docker Notes

## Install

```
DBUILD=stable
DARCH=x86_64
DVERSION=docker-20.10.3.tgz

curl -LO https://download.docker.com/linux/$DBUILD/$DARCH/$DVERSION
tar -xzf docker-20.10.3.tgz
cp docker/* $HOME/.local/bin

sudo groupadd docker
sudo usermod -aG docker $USER
newgrp docker
```

## Basic Workflow

`docker build . -t docker-image-name`

`docker run docker-image-name`

## system

`docker system prune` removes all stopped containers, networks not used
by any containers, dangling images, and dangling build cache. Ive often
recovered GBs worth of space using this command.

`docker system df` outputs a report that includes the type, total number
of the type, number of active type, total size, and reclaimable size.

`docker system events` outputs real time events from the Docker daemon.
Docker container events include: `attach`, `commit`, `copy`, `create`,
`destroy`, `detach`, `die`, `exec_create`, `exec_detach`, `exec_start`,
`export`, `kill`, `oom`, `pause`, `resize`, `restart`, `start`, `stop`,
`top`, `unpause`, and `update. Docker image events include: `delete`,
`import`, `load`, `pull`, `push`, `save`, `tag`, and `untag`. Docker
volume events include: `create`, `mount`, `unmount`, `destroy`. Docker
network events include: `create`, `connect`, `disconnect`, and `destroy`.

`docker system info` outputs system wide information about the Docker
installation, such as the daemon version, the kernel version, the
current operating system, the number of CPUs, the total memory, and
more.

## ps

`docker ps` is an alias for `docker container ls` and outputs containers.

`docker ps -a` || `docker ps --all` outputs all containers.

`docker ps -n <n>` || `docker ps --last=<n>` outputs `n` last created
containers.

`docker ps -l` || `docker ps --latest` outputs the latest created
container.

`docker ps -s` || `docker ps --size` outputs the total file sizes.

## images

`docker images` is an alias for `docker image ls` and outputs Docker
images.

`docker images -a` || `docker images --all` outputs all Docker images.

## container

`docker container attach <container>` attaches local stdin, stdout, and
stderr to `<container>`.

`docker container commit <container>` creates a new image using the
current state of `<container>`. 

`docker container cp <container>:<path> <destination>` copies
files/directories between `<container>:<path>` and `<destination>`.

`docker container create` creates a new container.

`docker container diff <container>` outputs a diff of the
files/directories for `<container>`.

`docker container exec <container> <command>` runs the `<command>` in the
`<container>`.

`docker container export <container>` exports the `<container>`
filesystem as a tar archive. I found this command quite funny, as if you
dont provide the `-o` flag it will output `cowardly refusing to save to a
terminal. Use the -o flag or redirect`.

`docker container inspect <container>` outputs information about
`<container>`. This includes information such as: `id`, `created`,
`state`, `image`, `name`, `platform`, and much more. More than one
`<container>` can be specified.

`docker container kill <container>` kills `<container>`. More than one
`<container>` can be specified.

`docker container logs <container>` outputs the logs of `<container>`.

`docker container ls` is the same as `docker ps`.

`docker container pause <container>` pauses all processes within
`<container>`. More than one `<container>` can be specified.

`docker container port <container>` outputs all port mappings or a
specific port mapping for `<container>`.

`docker container prune` removes all stopped containers.

`docker container rename <container> <name>` renames `<container>` to
`<name>`.

`docker container restart <container>` restarts `<container>`. More than
one `<container>` can be specified.

`docker container rm <container>` removes `<container>`. More than one
`<container>` can be specified.

`docker container run <image>` is similar to `docker container exec`,
but does not require a running container.

`docker container start <container>` starts `<container>`. More than one
`<container>` can be specified.

`docker container stats <container>` outputs the resource usage of
`<container>`. The resource information provided includes: `container
id`, `name`, `CPU %`, `mem usage / limit`, `mem %`, and more.

`docker container stop <container>` stops `<container>`. More than one
`<container>` can be specified.

`docker container top <container>` outputs the processes running on
`<container>`.

`docker container unpause <container>` unpauses all processes in
`<container>`. More than one `<container>` can be specified.

`docker container update <container>` updates the configuration of
`<container>`. More than one `<container>` can be specified.

`docker container wait <container>` blocks until `<container>` stops,
then outputs their exit codes. More than one `<container>` can be
specified.

## build

`docker build <path>` builds the Dockerfile found at `<path>`. A URL can
be used for `<path>`.

`docker build -f` || `docker build --file <path>` is the same as above.

`docker build --no-cache <path>` does not use cache when building the
image. This is `false` by default.

`docker build --pull <path>` attempts to pull the latest version of
the image.

`docker build --compress <path>` compresses the build context using gzip.

`docker build -q <path>` || `docker build --quiet <path>` suppresses any
build output and prints the image ID if successful.
