# PostgreSQL with C++

One of the fundamental things to know when learning any programming language is how to connect and interact with a database. I am a big fan of [PostgreSQL](https://www.postgresql.org/) since its open source and relational databases (MySQL) were the first form I learned about. This blog is going to cover how to connect and interact with a postgres database via C++. After a quick duckduckgo search, I found the official C++ client API for postgres called [libpqxx](https://github.com/jtv/libpqxx). Let the learning begin.

The first thing to do is to install libpqxx, which I am going to build from source. To start, we need to download the source code via `git clone https://github.com/jtv/libpqxx`. The README says the master branch is the current development version, so we need to checkout the latest stable version (`7.2.0` at the time of this blog) via `git checkout 7.2.0`. Before we build, there is one thing needed which is `libpq`. To build, since I am using ubuntu 20.04, we will need to run `./configure && make && sudo make install` (according to [building using configure](https://github.com/jtv/libpqxx/blob/master/BUILDING-configure.md)).

```
*****************************************************

xmlto not found.
Install it, or configure with --disable-documentation

*****************************************************
make[1]: *** [Makefile:506: tutorial-stamp] Error 1
make[1]: Leaving directory '/home/no/Repos/github/oglinuk/ptp/c++/hw-postgres/libpqxx/doc'
make: *** [Makefile:639: all-recursive] Error 1
```

I ran into the error above when running `make`, which I fixed by running `./configure --disable-documentation` and running `make && sudo make install` again.

---

With the library installed, time to make a "hello world" example, which in this case will just be a simple program to connect to a postgres database running in a docker container. Before I make the program, I need to spin up dockerized postgres database. This is easily done since there is a docker image (`postgres:latest`) I can pull down and spin up. Im using docker-compose to make spinning up the container even easier via `docker-compose up --build`.

***docker-compose.yml***

```
version: '3'
services:
  db:
		image: postgres:latest
    environment:
      - POSTGRES_PASSWORD=secret
    ports:
     - 5432:5432
    restart: always
```

Now lets write some C++.

```C++
#include <iostream>
#include <pqxx/pqxx>

int main()
{
	try {
		pqxx::connection conn("host=127.0.0.1 user=postgres password=secret dbname=postgres");
		std::cout << "Successfully connected to database " << conn.dbname() << std::endl;
	} catch (std::exception const &e) {
		std::cerr << e.what() << std::endl;
		return 1;
	}

	return 0;
}
```

According to the documentation for [pqxx::connection](https://libpqxx.readthedocs.io/en/7.2.0/a00818.html#details), "when creating a connection, you can pass a connection URI or a postgres connection string, to specify the database server's address, a login username, and so on. If none is given, the connection will try to obtain them from certain environment variables. If those are not set either, the default is to try and connect to the local system's port 5432." In this example I chose to use `keyword=value` instead of a URI, as I find it easier to read.

---

Now to compile the program via `g++ -std=c++17 main.cpp -o main -lpqxx -pq`.

One thing that caught me up when I was compiling the program was the version of C++. The repository for libpqxx [specifies](https://github.com/jtv/libpqxx#upgrade-notes) that when using `7.x` versions, C++ 17 is required. This is fixed by simply adding `-std=c++17` when running `g++` or `clang++`.

If all goes well you should get `Successfully connected to database postgres`.
