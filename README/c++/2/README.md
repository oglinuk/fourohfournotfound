---
title: C++ 2 ~ OpenCV
date: 2019-07-21T15:22:42+12:00
tags: ["C++,", "OpenCV"]
draft: false
---

[OpenCV](https://opencv.org) is a library that I will be working with in C++, so this blog is going to be dedicated to installing OpenCV and setting it up for C++. The steps I will be following are the ones found on the [linux tutorial](https://docs.opencv.org/4.0.0/d7/d9f/tutorial_linux_install.html) section of the OpenCV website.

To start, we need to install the following dependancies.

```Bash
sudo apt install build-essential \
                 cmake \
                 git \
                 libgtk2.0-dev \
                 pkg-config \
                 libavcodec-dev \
                 libavformat-dev \
                 libswscale-dev \
                 python3-dev \
                 python3-numpy \
                 libtbb2 \
                 libtbb-dev \
                 libjasper-dev \
                 libjpeg-dev \
                 libpng-dev \
                 libtiff-dev \
                 libdc1394-22-dev
```

---

Next to clone the OpenCV and OpenCV contrib git repositories via 

`git clone https://github.com/opencv/opencv.git` and 

`git clone https://github.com/opencv/opencv_contrib.git`. 

To simplfy things I have made a directory called `opencv_dev`, and cloned the repositories inside of it. 

---

Now to create a directory called `build` inside of the `opencv` directory and `cd` into that new directory.

Within `build` we now run the following command.

```
cmake -D CMAKE_BUILD_TYPE=RELEASE \
      -D OPENCV_GENERATE_PKGCONFIG=YES \
      -D CMAKE_INSTALL_PREFIX=/usr/local \
      -D OPENCV_EXTRA_MODULES_PATH=../../opencv_contrib/modules 
      -D INSTALL_C_EXAMPLES=ON \
      -D INSTALL_PYTHON_EXAMPLES=ON \
      -D WITH_V4L=ON \
      -D WITH_QT=OFF \
      -D WITH_TBB=ON \
      -D WITH_CUDA=ON \
      -D WITH_OPENGL=ON \
      -D BUILD_EXAMPLES=ON \
      -D BUILD_OPENCV_PYTHON2=ON \
      -D BUILD_OPENCV_PYTHON3=ON \
      -D BUILD_SHARED_LIBS=ON ..
```

The next command to run is `make -j7`, then `sudo make install`.

---

With that I now have the ability to use `#include <opencv2/opencv.hpp>` in the following example.

```C++
// main.cpp
#include <opencv2/opencv.hpp>

using namespace cv;

int main(int argc, char const *argv[])
{
    if (argc < 2) {
        printf("Usage: ./main <image_path>\n");
        return -1;
    }

    Mat img;
    img = imread(argv[1], 1);

    if (!img.data) {
        printf("No image data\n");
        return -1;
    }

    namedWindow("Image", WINDOW_AUTOSIZE);
    imshow("Image", img);

    waitKey(0);

    return 0;
}
```

Then to make the `CMakeLists.txt` file, which is as follows.

```
cmake_minimum_required(VERSION 3.10)
project( main )
find_package( OpenCV REQUIRED )
include_directories( ${OpenCV_INCLUDE_DIRS} )
add_executable( main main.cpp )
target_link_libraries( main ${OpenCV_LIBS} )
```

Because the above is taken from the [gcc and cmake tutorial](https://docs.opencv.org/3.2.0/db/df5/tutorial_linux_gcc_cmake.html), I want to cover each line of the `CMakeLists.txt` file. This is mainly for my personal learning, but will be beneficial if the reader happens to also be learning. 

---

The first line `cmake_minimum_required(VERSION 3.10)`, sets the minimum required version of `cmake` for a project and update policy settings to match the version given. In my case, `3.10` is specified since that is the installed version of `cmake` on my machine. [[source]](https://cmake.org/cmake/help/v3.10/command/cmake_minimum_required.html?highlight=cmake_minimum#command:cmake_minimum_required)

The second line `project( main )`, sets a name, version, and enables languages for the entire project. [[source]](https://cmake.org/cmake/help/v3.10/command/project.html?highlight=project)

The third line `find_package( OpenCV REQUIRED )`, loads settings for an external project. [[source]](https://cmake.org/cmake/help/v3.10/command/find_package.html?highlight=find_package)

The fourth line `include_directories( ${OpenCV_INCLUDE_DIRS} )`, adds include directories to the build. [[source]](https://cmake.org/cmake/help/v3.10/command/include_directories.html?highlight=include_directories#command:include_directories)

The fifth line `add_executable( main main.cpp )`, adds an executable to the project using the specified source files. [[source]](https://cmake.org/cmake/help/v3.10/command/add_executable.html?highlight=add_executable)

The sixth and final line `target_link_libraries( main ${OpenCV_LIBS} )`, specify libraries or flags to use when linking a given target and/or its dependents. [[source]](https://cmake.org/cmake/help/v3.10/command/target_link_libraries.html?highlight=target_link_libraries)

---

Finally to build/compile via `mkdir build && cd build`, `cmake ..`, `make`, then `./main os.jpeg`.

We got the `os.jpeg` picture to open; OpenCV with C++ works!
