---
title: Docker 2 ~ Multitier Architecture
date: 2019-04-28
tags: ["Docker,", "Python,", "Docker-compose"]
draft: false
---

For my SDV701 assessment 2 we have to build a tiered software. I've decided to use Flask for the presentation tier, Flask-restful for the logic tier, and mongodb (via pymongo) for the data tier. I've also decided to use docker and docker-compose since running a mongodb server in docker is as easy as pulling an image from docker-hub. Below is a snippet from the `docker-compose.yml` file, showing how easy it is.

```Yaml
mongodb:
  image: mongo:latest
  ports:
    - 27017:27017
  volumes:
    - ./data/db:/usr/data/db
  command: mongod --smallfiles --logpath=/dev/null
```

During the development of the REST API I ran into `[Errno 111] Connection refused`.

This was because when I was calling the API method from the presentation tier (using requests), I was using `localhost` for the host in the url. This was confusing for a while since I was able to execute a `curl` command, passing `'http://localhost:9124/todo'` as the url with no errors. Eventually (after a bit of duckduckgo searches and tampering) I asked [Sam](https://github.com/pigeonhands) and he pointed out that the docker containers can't use `localhost` and should instead use the containers [container_name](https://docs.docker.com/compose/compose-file/#container_name). Do note that the documentation states that "you cannot scale a service beyond 1 container if you have specified a custom name. Attempting to do so results in an error." In this case I am only needing 1 container, so it shouldn't be a problem.

In order for the `container_name` to be used, (originally) [Sam](https://github.com/pigeonhands) suggested to use [links](https://docs.docker.com/compose/compose-file/#links). After reading the documentation on `links`, it is suggested to use [networks](https://docs.docker.com/compose/compose-file/#networks), as `links` may eventually be removed. The `docker-compose.yml` file now looks like the following (with the addition of the presentation and logic tier).

```Yaml
version: '3.7'
services:
  web:
    hostname: frontend
    container_name: tiered-frontend
    build: ./frontend
    depends_on:
      - api
    ports:
      - 9123:9123
    volumes:
      - .:/usr/src/application
    environment:
      - ENV=development
      - WEB_HOST=0.0.0.0
      - WEB_PORT=9123
    networks:
      - tiered
  api:
    hostname: backend
    container_name: tiered-backend
    build: ./backend
    depends_on:
      - mongodb
    ports:
      - 9124:9124
    volumes:
      - .:/usr/src/application
    environment:
      - ENV=development
      - API_HOST=0.0.0.0
      - API_PORT=9124
      - DB=mongodb://mongodb:27017/tieredapp
    networks:
      - tiered
  mongodb:
    hostname: database
    container_name: tiered-database
    image: mongo:latest
    ports:
      - 27017:27017
    volumes:
      - ./data/db:/usr/data/db
    command: mongod --smallfiles --logpath=/dev/null
    networks:
      - tiered
networks:
  tiered:
```

Now we can call the REST API from the presentation tier (via `requests`) with the following.

```Python
r = requests.get('http://tiered-backend:9124/todo')
```

***Spicy***
