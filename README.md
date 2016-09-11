# chaos-wave

```chaos-wave``` will randomly kill and remove running container. Inspirted by chaos monkey form netflix.

## Usage

```
$ docker run -v /var/run/docker.sock:/var/run/docker.sock furikuri/chaos-wave --interval 15m --duration 2h
```

Default **interval** is 10 minutes.
Default **duration** is 1 hour.

### Example usage in docker swarm

First start a service within docker swarm.

```
$ docker service create --replicas=3 --name hello-docker furikuri/hello-docker
```

Let's observe the running containers with ```watch 'docker ps'```.

```
Every 2,0s: docker ps                                                                                                                                                Sun Sep 11 12:31:50 2016

CONTAINER ID        IMAGE                          COMMAND                  CREATED             STATUS              PORTS               NAMES
c51676f6731f        furikuri/hello-docker:latest   "/bin/sh -c 'node ser"   8 seconds ago       Up 6 seconds        3000/tcp            hello-docker.1.8q22jjf5i7trubngv8wo4sh1v
9e6114c94252        furikuri/hello-docker:latest   "/bin/sh -c 'node ser"   8 seconds ago       Up 6 seconds        3000/tcp            hello-docker.2.5fwbshxr1p88mcszrs1u6zdxw
e0f33516014e        furikuri/hello-docker:latest   "/bin/sh -c 'node ser"   8 seconds ago       Up 6 seconds        3000/tcp            hello-docker.3.3yyalv0viof2x3runw9nb4yyr
```

Now let's start the ```chaos-wave```. After 30 seconds we should see the first output, which tells us that the first container has been killed.

```
$ docker run -v /var/run/docker.sock:/var/run/docker.sock furikuri/chaos-wave --interval 30s --duration 2h
Duration: 2h0m0s.
Interval: 30s.
Stopping container with name '/hello-docker.2.5fwbshxr1p88mcszrs1u6zdxw'
```

This container will be not existing in our ```docker ps``` anymore.

```
Every 2,0s: docker ps                                                                                                                                                Sun Sep 11 12:32:54 2016

CONTAINER ID        IMAGE                          COMMAND                  CREATED              STATUS              PORTS               NAMES
44864f1ef4f6        furikuri/chaos-wave            "/app/main --interval"   33 seconds ago       Up 31 seconds                           high_hugle
c51676f6731f        furikuri/hello-docker:latest   "/bin/sh -c 'node ser"   About a minute ago   Up About a minute   3000/tcp            hello-docker.1.8q22jjf5i7trubngv8wo4sh1v
e0f33516014e        furikuri/hello-docker:latest   "/bin/sh -c 'node ser"   About a minute ago   Up About a minute   3000/tcp            hello-docker.3.3yyalv0viof2x3runw9nb4yyr


```

```
Every 2,0s: docker ps                                                                                                                                                Sun Sep 11 12:33:01 2016

CONTAINER ID        IMAGE                          COMMAND                  CREATED              STATUS              PORTS               NAMES
1bdad8b5d987        furikuri/hello-docker:latest   "/bin/sh -c 'node ser"   4 seconds ago        Up 1 seconds        3000/tcp            hello-docker.2.9qp4ld0rn5y102u84kqcjldrm
44864f1ef4f6        furikuri/chaos-wave            "/app/main --interval"   39 seconds ago       Up 37 seconds                           high_hugle
c51676f6731f        furikuri/hello-docker:latest   "/bin/sh -c 'node ser"   About a minute ago   Up About a minute   3000/tcp            hello-docker.1.8q22jjf5i7trubngv8wo4sh1v
e0f33516014e        furikuri/hello-docker:latest   "/bin/sh -c 'node ser"   About a minute ago   Up About a minute   3000/tcp            hello-docker.3.3yyalv0viof2x3runw9nb4yyr

```
## Todos

- [ ] Regex argument for container names to kill

## License