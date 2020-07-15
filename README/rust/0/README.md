---
title: Rust 0 ~ Hello Rust!
date: 2019-03-14
tags: ["Hello World,", "Rust"]
draft: false
---

The first question anyone might ask themselves when they read the title. This is an extremely valid question, since why if one could learn something as simple as Python, would one want to learn a language like Rust? I know this because that was the first question that *I* asked. The answer is not entirely formulated as I still learn new and exciting things everyday. 

What I can say is that for one Rust is *fast*. To understand how fast I cant directly compare a language like *Python* with it. But I can compare how fast [Python vs Go](https://benchmarksgame-team.pages.debian.net/benchmarksgame/fastest/python3-go.html) is, and I can also compare how fast [Rust vs Go](https://benchmarksgame-team.pages.debian.net/benchmarksgame/fastest/rust-go.html) is. Then to understand how *much faster* Rust is than go (even though the gap is small on some benchmarks) one can look at the [Rust vs C](https://benchmarksgame-team.pages.debian.net/benchmarksgame/fastest/rust.html) and [Go vs C](https://benchmarksgame-team.pages.debian.net/benchmarksgame/fastest/go-gcc.html). A co-worker of mine told me you can't really rely on benchmarks because people can always tweak and do small optimizations to make anything faster, but the fact that it competes with C is insane. It is also important to note that these are only benchmarks from one source and the benchmarks are quite old.

One of the other things I can say is that it is actively developed (and by a big player in the open source space ~ [Mozilla](https://en.wikipedia.org/wiki/Rust_(programming_language)#cite_note-future-tense-35)). There are websites called *arewe<name>yet*, some examples are [arewegameyet](http://arewegameyet.com/) or [arewewebyet](http://www.arewewebyet.org/), which serve as awesome sources of information as to the current status of a certain aspect in Rust. There is also a rich [awesome-rust](https://github.com/rust-unofficial/awesome-rust#table-of-contents) repository.

A big reason why I started to learn Go was its ability to produce a fully self-contained binary/executable that works on any operating system. As someone who dislikes the need for virtual environments when using Python in order to manage dependancies, this is like a breath of fresh air. This was before I came to know of *docker*. Rust also has this ability with `cargo build`. 

The one that got me interested in Rust was none other than [Sam](https://github.com/pigeonhands), the one who also introduced me to Python, Go, and many other of the skills that I am currently developing. One of the things that got *him* interested was an [emulator project](http://koute.github.io/pinky-web/) built in Rust that runs in the browser. Its kinda hard to understand why this particular example is quite exciting, since I myself don't yet fully understand [*web-assembly*](https://webassembly.org/), but I do know that being able to an emulator in a browser is freaking awesome.

What better way to learn a language than to dive right into it with a project. After spending a few hours of reading various [crate documentation](https://crates.io)/[Rust documentation](https://docs.rs/) and hacking, Sam and I made the basis for a Rust ported version of my concurrent crawler written in Go ([goccer](https://github.com/OGLinuk/goccer)). The source code for [cruster](https://github.com/OGLinuk/cruster) is also publicly available. Its hard to compare the crawlers at this stage since the way that Sam and I designed the Rust crawler is different.
