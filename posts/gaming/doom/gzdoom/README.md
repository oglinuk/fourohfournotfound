# GZDoom

GZDoom is a sourceport which incorporates modern GPU capabilities.

`mkdir -pv ~/gzdoom_build`

## **[ZMusic Steps](https://forum.zdoom.org/viewtopic.php?t=67489)**

There is one dependency of GZDoom that is also maintained by `coelckers`, which
is [`ZMusic`](https://github.com/coelckers/ZMusic.git).

```
git clone https://github.com/coelckers/ZMusic.git
cd ZMusic
mkdir build && cd build
cmake -DCMAKE_BUILD_TYPE=Release \
    -DCMAKE_INSTALL_PREFIX=`pwd`/../build_install ..
make install
```

## **[GZDoom Steps](https://zdoom.org/wiki/Compile_GZDoom_on_Linux#Compiling)**

```
cd ~/gzdoom_build &&
git clone git://github.com/coelckers/gzdoom.git &&
mkdir -pv gzdoom/build
```

```
cd gzdoom
git config --local --add remote.origin.fetch +refs/tags/*:refs/tags/*
git pull
```

```
a='' && [ "$(uname -m)" = x86_64 ] && a=64
c="$(lscpu -p | grep -v '#' | sort -u -t , -k 2,4 | wc -l)" ; [ "$c" -eq 0 ] && c=1
cd ~/gzdoom_build/gzdoom/build &&
rm -f output_sdl/liboutput_sdl.so &&
if [ -d ../fmodapi44464linux ]; then
f="-DFMOD_LIBRARY=../fmodapi44464linux/api/lib/libfmodex${a}-4.44.64.so \
-DFMOD_INCLUDE_DIR=../fmodapi44464linux/api/inc"; else
f='-UFMOD_LIBRARY -UFMOD_INCLUDE_DIR'; fi &&
cmake .. -DCMAKE_BUILD_TYPE=Release $f \
    -DZMUSIC_INCLUDE_DIR=~/gzdoom_build/ZMusic/build_install/include \
    -DZMUSIC_LIBRARIES=~/gzdoom_build/ZMusic/build_install/lib/libzmusic.so &&
make -j$c
```
