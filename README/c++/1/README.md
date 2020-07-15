---
title: C++ 1 ~ Make && CMake
date: 2019-07-21T09:30:37+12:00
tags: ["C++,", "Make,", "CMake,", "Fundamentals"]
draft: false
---

One of the things [Sam](https://github.com/pigeonhands) suggested to learn/use with C++ is [make](https://gnu.org/software/make/) and [cmake](https://cmake.org). I've come across make before, but not cmake, this blog is going to be about learning the basics of both.

The following is a simple program that I made to start with make.

```C++
// main.cpp
#include "hello.h"

int main()
{
    Hello h;
    h.printHello();

    return 0;
}
```

```C++
// hello.h
class Hello
{
public:
    void printHello();
};
```

```C++
// hello.cpp
#include <iostream> // System defined header file
#include "hello.h" // User defined header file

using namespace std;

void Hello::printHello()
{
    cout << "Hello world from Makefile example!\n";
}
```

\

Now for the `Makefile`.

```Makefile
# Makefile
output: main.o hello.o
	g++ main.o hello.o -o main

main.o: main.cpp
	g++ -c main.cpp

hello.o: hello.cpp hello.h
	g++ -c hello.cpp

clean:
	rm *.o main
```

Now all that needs to be done is run the `make` command. After running the newly compiled binary `main`, we get the expected output; `Hello world from Makefile example!`.

\

Next is to create a cmake example using the same program.

```
// CMakeLists.txt
cmake_minimum_required (VERSION 3.10)

project (hw_cmake_example)

add_executable(
    main
    main.cpp
    hello.cpp
    hello.h
)
```

With the `CMakeLists.txt` file done, now to create a build directory `mkdir build && cd build`, then to run the command `cmake ..`. This creates a few files, one of which is a `Makefile`. We can now run `make` and then run the compiled binary `main`, which gives the expected output; `Hello world from the CMake example!`.
