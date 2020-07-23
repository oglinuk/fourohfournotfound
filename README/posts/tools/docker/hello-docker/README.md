---
title: Hello Docker!
date: 2019-03-30
tags: ["Hello World,", "Docker,", "Python"]
draft: false
---

Back in 2017 I learned about a ton of new and exciting technologies that I had never used or even heard of. One of those technologies is ***docker***. If you don't know of docker, then let me give some background information. "Docker is an open platform for developing, shipping, and running applications." [[source]](https://docs.docker.com/engine/docker-overview/). Before we can dive further into what this really means we need to first come to an understanding of [*containerization*](https://en.wikipedia.org/wiki/Container_(virtualization)). According to the docker website, "a container is a standard unit of software that packages up code and all its dependencies so the application runs quickly and reliably from one computing environment to another." [[source]](https://www.docker.com/resources/what-container).

One of the biggest things to remember is that containers are *not* virtual machines. "Containers and virtual machines have similar resource isolation and allocation benefits, but function differently because containers virtualize the operating system instead of hardware. Containers are more portable and efficient." [[source]](https://www.docker.com/resources/what-container)

A question I was once asked when I gave a talk at a Python meetup was "how easy is it to dockerize an existing project?" This is a great question, as one of the biggest make or breaks with technology is its capability of integrating with existing code bases. Though I can't say for larger code bases, I can say for my projects, adding docker is simple. This is coming from the perspective of Python projects, which have whats called *requirements.txt* files that contain all the dependancies. All that needs adding is a *Dockerfile*, which an example can be found below.

---

```Dockerfile
FROM python:3.6
ADD . /code
WORKDIR /code
EXPOSE 9123
RUN pip install -r requirements.txt
CMD ["python3", "run.py"]
```

Let's break the above down line by line. 

`FROM` gets the *python 3.6* [image](https://docs.docker.com/glossary/?term=image) and uses it for the base. 

`ADD` copies the current directory (.) to the filesystem of the image at the `/code` path. 

`WORKDIR` sets the working directory for the `RUN` and `CMD` instructions that follow. 

`EXPOSE` instructs the docker container to listen on the specified port at runtime.

`RUN` executes any given command(s). In this case we are executing a pip install.

`CMD` provides defaults for an executing container. In this case the default is running `python3 run.py`.

---

With this one file, the application is now dockerized. All that needs done is executing the following commands.

`docker build . --tag=hello-world`

`docker run hello-world`

Now that's what I call ***Spicy***.
