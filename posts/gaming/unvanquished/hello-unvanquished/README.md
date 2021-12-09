# Hello Unvanquished!

## Dependencies

* `curl`
* `cmake`
* `python2`
* `python-yaml`
* `python-jinja`
* `zlib`
* `libgmp`
* `libpng`
* `libjpeg` >= 8.0.0
* `libwebp` >= 0.2.0
* `libnettle`
* `libogg`
* `libvorbis`
* `libopus`
* `libopusfile`
* `libtheora`
* `freetype`
* `sdl2`
* `glew`
* `lua`
* `openal`
* `ncurses`
* `geoip`

## Get the Source Code

```
git clone https://github.com/Unvanquished/Unvanquished.git`
cd Unvanquished && git submodule update --init --recursive
```

## Download Game Assets

`./download-paks`

## Compiling

```
mkdir build && cd build
cmake .. && make -j$(nproc)
```

## Playing

`./daemon`
