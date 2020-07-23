---
title: "Hello Quake3!"
date: 2020-04-13T08:31:38+12:00
draft: false
tags: ["Hello World,", "Quake3,", "C,", "Linux,", "Open Source"]
---

In this blog I am going to cover how to get Quake 3 arena running. The source port I am going to use is [ioquake3](https://github.com/ioquake/ioq3). Unlike [DOOM](/games/doom), there is no entirely free version of Quake 3. Quake 3 requires [.pak files](https://quakewiki.org/wiki/.pak), which can found in the game files after purchasing Quake 3 on steam. 

The following dependencies are required.

* `build-essential`
* `make`
* `git`
* `libsdl2-2.0-0`

---

First to download the [ioquake3](https://github.com/ioquake/ioq3) via

`git clone https://github.com/ioquake/ioq3`

`cd` into `ioq3` and execute `make`

Once the compilation process has finished, then `cd` into `build/release-linux-x86_64`

Finally execute `./ioquake3.x86_64`

***Quake on!***
