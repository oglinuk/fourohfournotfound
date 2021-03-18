# Hello Tensorflow!

There are a lot of frameworks that you could choose from when it comes to AI. The following are the steps to install and setup [Tensorflow](https://github.com/Tensorflow/Tensorflow) (more specifically setting up Tensorflow to use the [object detection API](https://github.com/tensorflow/models/tree/master/research/object_detection)). It is important to note that I am using a GPU, so my setup involves installing Tensorflow with GPU support.

To start, the following dependencies are required.

```Bash
sudo apt install cmake \
                 curl \
                 git
```

---

Next we need the proper drivers. 

```Bash
sudo add-apt repository ppa:graphics-drivers/ppa
sudo apt update
sudo apt install nvidia-driver-410
```

---

The next thing is to install CUDA. You will need to install the [CUDA 10.0](https://developer.nvidia.com/cuda-10.0-download-archive?target_os=Linux&target_arch=x86_64&target_distro=Ubuntu&target_version=1804&target_type=deblocal) file. Note that I am using ubuntu 18.04.

```Bash
sudo dpkg -i cuda-repo-ubuntu1804-10-0-local-10.0.130-410.48_1.0-1_amd64.deb
sudo apt-key add /var/cuda-repo-10-0-local-10.0.130-410.48/7fa2af80.pub
sudo apt update
sudo apt install cuda
```

It's important to note that the CUDA version should be the same as your nvidia driver, which in my case is `410`.

Then add the following to the end of your `bashrc` file.

```Bash
export PATH=/usr/local/cuda-10.0/bin:/usr/local/cuda-10.0/NsightCompute-1.0${PATH:+:${PATH}}

export LD_LIBRARY_PATH=/usr/local/cuda-10.0/lib64${LDInstall_LIBRARY_PATH:+:${LD_LIBRARY_PATH}}
```

Run `source ~/.bashrc` then cd into `/usr/local/cuda-10.0/samples` and run `sudo make`.

---

Next to install [cuDNN](https://developer.nvidia.com/rdp/cudnn-download), which you will need to create an account for. Download the cuDNN `runtime`, `developer`, and `code samples` files under cuDNN v7.6.3 for CUDA 10.0.

```Bash
sudo dpkg -i libcudnn7_7.6.3.30-1+cuda10.0_amd64.deb
sudo dpkg -i libcudnn7-dev_7.6.3.30-1+cuda10.0_amd64.deb
sudo dpkg -i libcudnn7-doc_7.6.3.30-1+cuda10.0_amd64.deb

cd /usr/src/cudnn_samples_v7/mnistCUDNN
sudo make clean && sudo make
./mnistCUDNN
```

Like with the CUDA version, make sure the version of cuDNN corresponds to the version of CUDA (`10.0` in this case).

---

Now to install OpenCV(4) from source.

```Bash
sudo apt install build-essential \
                 libgtk2.0-dev \
                 pkg-config \
                 libavcodec-dev \
                 libavformat-dev \
                 libswscale-dev \
                 python3-dev \
                 libtbb2 \
                 libtbb-dev \
                 libjpeg-dev \
                 libpng-dev \
                 libtiff-dev \
                 libjasper-dev \
                 libdc1394-22-dev

mkdir opencv_build && cd opencv_build
git clone https://github.com/opencv/opencv.git
git clone https://github.com/opencv/opencv_contrib.git
cd opencv
mkdir build && cd build

cmake -D CMAKE_BUILD_TYPE=RELEASE \
      -D OPENCV_GENERATE_PKGCONFIG=YES \
      -D CMAKE_INSTALL_PREFIX=/usr/local \
      -D OPENCV_EXTRA_MODULES_PATH=../../opencv_contrib/modules \
      -D WITH_OPENGL=ON \
      -D WITH_QT=OFF \
      -D WITH_CUDA=ON \
      -D WITH_GTK=ON \
      -D WITH_GTK_2_X=ON \
      -D WITH_TBB=ON \
      -D WITH_V4L=ON \
      -D BUILD_OPENCV_PYTHON3=ON \
      -D BUILD_SHARED_LIBS=ON ..

make -j7
sudo make install
```

---

Finally to install Tensorflow with GPU support.

```Bash
sudo apt install python3-pip

python3 -m pip install --user six \
                              wheel \
                              setuptools \
                              mock \
                              future \
                              matplotlib \
                              opencv-python \
                              lxml \
                              pillow \
                              Cython \
                              jupyter \
                              numpy \
                              pandas \
                              scipy \
                              scikit-learn \
                              tensorflow-gpu \

python3 -m pip install --user keras_applications==1.0.6 --no-deps
python3 -m pip install --user keras_preprocessing==1.0.5 --no-deps

echo "deb [arch=amd64] https://storage.googleapis.com/bazel-apt stable jdk1.8" | sudo tee /etc/apt/sources.list.d/bazel.list
curl https://bazel.build/bazel-release.pub.gpg | sudo apt-key add -
sudo apt update
sudo apt install bazel
sudo apt install --only-upgrade bazel
bazel version
```

---

All you need to do now is restart your machine and start playing! My suggestion is to check out the [tensorflow models](https://github.com/tensorflow/models) repository and play around with some pre-trained models.
