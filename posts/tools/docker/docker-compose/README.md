# Docker Compose

In this blog we are going to explore a tool thats apart of the docker ecosystem; ***docker-compose***. "Compose is a tool for defining and running multi-container Docker applications." [[source]](https://docs.docker.com/compose/) There is an awesome example found on the [getting started](https://docs.docker.com/compose/gettingstarted/) section, which shows the power of docker-compose. The example is shown below.

***app.py***

```Python
from flask import Flask
import time
import redis

app = Flask(__name__)
cache = redis.Redis(host='redis', port=6379)

def get_seen_count():
    retries = 5
    while True:
        try:
            return cache.incr('seen')
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

---

***requirements.txt***

```
flask
redis
```

---

***Dockerfile***

```Dockerfile
FROM python:3.6-alpine
ADD . /code
WORKDIR /code
EXPOSE 9001
RUN pip install -r requirements.txt
CMD ["python", "app.py"]
```

---

***docker-compose.yml***

```Yaml
version: '3'
services:
  web:
    build: .
    ports:
     - "9001:9001"
  redis:
    image: "redis:alpine"
```

Now all that we need to do is execute `docker-compose up --build` and we have a flask web application with a redis database running. 

***Spicy***
