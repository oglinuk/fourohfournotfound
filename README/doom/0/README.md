---
Title: Doom 0
date: 2019-09-22T11:20:58+12:00
tags: ["DOOM", "C", "C++", "Linux", "Open Source"]
draft: false
---

DOOM is easily (in my opinion), one of the most iconic games ever made. It is enjoyable to play even to this day. The best part of DOOM is the fact that it has been made [open source](https://github.com/id-Software/DOOM) under a [GPL license](https://doomwiki.org/wiki/Linux_Doom). From this there have been many [source ports](https://doomwiki.org/wiki/Source_port) of the original DOOM code. I particularly like gzdoom as it adds graphics rending features, but there are quite a few that all have their own unique additions. Id software was ahead of their time, in more ways that one. One of the amazements are at the WAD files. 
I am still new to the ports of DOOM, but everytime I play them, the fun and playability never cease.

In this blog I am going to show you how to setup your linux environment to play an entirely free version of DOOM called [freedoom](https://freedoom.github.io/). The OS I am using is ubuntu 18.04, but should be similar if not the same across multiple distributions. There are only a few things needed to play DOOM on your local machine; install 9 OS dependencies, clone a source port ([crispy-doom](https://github.com/fabiangreffrath/crispy-doom) in my case), and a `DOOM.WAD` file. You can get freely available WADS of a non-official DOOM called [freedoom1/2](https://freedoom.github.io/download.html). For this blog we are going to use [freedoom](https://freedoom.github.io).

---

The dependencies that need to be installed are as follows.

* `autoconf`
* `build-essential`
* `cmake`
* `git`
* `libsdl2-dev`
* `libsdl2-mixer-dev`
* `libsdl2-net-dev`

The source port we are going to use is [crispy-doom](https://github.com/fabiangreffrath/crispy-doom), a source port fork of another source port called [chocolate-doom](https://github.com/chocolate-doom/chocolate-doom).

First clone the repository with the following command.

`git clone https://github.com/fabiangreffath/crispy-doom`

Then to compile run the following commands.

```
cd crispy-doom
autoreconf -fiv
./configure
make
``` 

Now to get the `freedoom1.wad`, which can be downloaded for free at the [freedoom website](https://freedoom.github.io/download.html)

To run `cd` into the `src` directory and execute `./crispy-doom -iwad <PATH TO freedoom1.wad>`.

*That's it*. ***dOOOOOOOM on!***
