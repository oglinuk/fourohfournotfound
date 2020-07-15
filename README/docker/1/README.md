---
title: Docker 1 ~ Docker Compose
date: 2019-04-01
tags: ["Docker,", "Python,", "Docker-compose"]
draft: false
---

In this blog we are going to explore a tool thats apart of the docker ecosystem; ***docker-compose***. "Compose is a tool for defining and running multi-container Docker applications." [[source]](https://docs.docker.com/compose/) There is an awesome example found on the [getting started](https://docs.docker.com/compose/gettingstarted/) section, which shows the power of docker-compose. The example is shown below.

***app.py***

```Python
import time

import redis
from flask import Flask


app = Flask(__name__)
cache = redis.Redis(host='redis', port=6379)


def get_hit_count():
    retries = 5
    while True:
        try:
            return cache.incr('hits')
        except redis.exceptions.ConnectionError as exc:
            if retries == 0:
                raise exc
            retries -= 1
            time.sleep(0.5)


@app.route('/')
def hello():
    count = get_hit_count()
    return 'Hello World! I have been seen {} times.\n'.format(count)

if __name__ == "__main__":
    app.run(host="0.0.0.0", debug=True)
```

***requirements.txt***

```None
flask
redis
```

***Dockerfile***

```Dockerfile
FROM python:3.4-alpine
ADD . /code
WORKDIR /code
RUN pip install -r requirements.txt
CMD ["python", "app.py"]
```

***docker-compose.yml***

```Yaml
version: '3'
services:
  web:
    build: .
    ports:
     - "5000:5000"
  redis:
    image: "redis:alpine"
```

Now all that we need to do is execute `docker-compose up --build` and we have a flask web application with a redis database running. 

***Spicy***
