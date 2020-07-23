---
title: Hello Alacritty!
date: 2020-04-11T10:36:28+12:00
tags: ["Hello World,", "Alacrity,", "Linux,", "Rust,", "Open Source"]
draft: false
---

If you use Linux then most likely you also use a terminal. One of the things that I still have not gotten used to after switching from windows is how *customizable* Linux really is. In this case the terminal. The first thing I asked myself when I came across [Alacritty](https://github.com/alacritty/alacritty) is "why?". It never clicked until [rwxrob](https://twitch.tv/rwxrob) made a comment about how he was using Alacritty as his terminal. 

Alacritty is a GPU-accelerated terminal built in [Rust](https://www.rust-lang.org/). In this blog I am going to cover how to install and configure Alacritty. I am using ubuntu 18.04, but the [installation](https://github.com/alacritty/alacritty#Installation) section covers other distributions.

First to install the required debian dependencies

```Bash
sudo apt install cmake \
                 pkg-config \
                 libfreetype6-dev \
                 libfontconfig1-dev \
                 libxcb-xfixes0-dev \
                 python3
```

We also need rust which can be easily installed by the following [cURL](https://github.com/curl/curl) command. 

`curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh`

The installation page of Alacritty suggests using the stable compiler which can be enabled via the following.

`rustup override set stable` && `rustup update stable`

---

Once all of the dependencies are installed the next thing to do is to clone the Alacritty repository via the following git command.

`git clone https://github.com/alacritty/alacritty` && `cd alacritty`

If you're on a debian distribution, then you can run the following commands.

`cargo install cargo-deb` && `cargo deb --install -p alacritty`

If not you can simply run `cargo build --release`.

Then to set the default terminal run 

`gsettings set org.gnome.desktop.default-applications.terminal exec 'alacritty'`

---

Open and enjoy the awesome terminal!
