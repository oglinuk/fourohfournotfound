# Media Streaming

Ever wondered how to stream media without needing a proprietary site?
Enter the [nginx-rtmp-module](https://github.com/arut/nginx-rtmp-module).
I dont live stream, but I would like to start a collection of public
domain media that I make available for public viewing. This post is going
to cover how to get started using the rtmp module and how to stream/view
content.

## Install Nginx && nginx-rtmp-module

The first thing to do (after creating a droplet) is to install Nginx and
the rtmp module.

```
sudo apt-get install build-essential libpcre3 libpcre3-dev libssl-dev git zlib1g-dev -y

cd /usr
mkdir build
cd build

git clone git://github.com/arut/nginx-rtmp-module.git

curl -LO http://nginx/download/nginx-1.19.9.tar.gz
tar -xzf nginx-1.19.9.tar.gz
cd nginx-1.19.9

./configure --with-http_ssl_module --add-module=../nginx-rtmp-module
make
make install

/usr/local/nginx/sbin/nginx
```

Next to setup the live streaming, which I will [refer you to the
docs](https://github.com/arut/nginx-rtmp-module/wiki/Getting-started-with-nginx-rtmp#set-up-live-streaming)
for the `nginx.conf` content.

Once youve updated the config file, restart Nginx.

```
/usr/local/nginx/sbin/nginx -s stop
/usr/local/nginx/sbin/nginx
```

Finally to check and ensure its working head to `http://<ip
address>:8080/stat`. If everything was setup correct you should see
various information like `# clients`, `Video`, `Audio`, and others.

## Test the Streaming

To test out the streaming I will be using
[Spring](https://www.blender.org/press/spring-open-movie/) and
[ffmpeg](https://ffmpeg.org/).

To start the stream you will need to upload the video to the server, then
run `ffmpeg -re -i spring-blender-2019.mp4 -c copy -f flv
rtmp://localhost/myapp/test` on the droplet.

To view the stream, on your local machine, run `ffplay rtmp://<ip
address>/myapp/test`. If you arent using Linux, you can view the stream
using VLC by going to `Media` on the top left, select `Open Network
Stream`, click the `Network` tab, and enter the rtmp URL.

If everything was done correctly, you should see the video play on your
local machine! So freaking awesome!


## Stream with OBS

In order to publish a stream to the remote server, we must add a
directive to the `nginx.conf` to allow for a stream to be published from
my desktop (since it will be where I play the dvd from).

```
rtmp {
	server {
		listen 1935;
		ping 30s;
		notify_method get;

		application live {
			live on;
			allow publish <ip address>;
			deny publish all;
	}
}
```

The `allow publish <ip address>;` directive enables me to publish from my
desktop, and we set `deny publish all;` so that only I can publish
streams.

Need to restart Nginx again.

```
/usr/local/nginx/sbin/nginx -s stop
/usr/local/nginx/sbin/nginx
```

With that done, now to configure OBS. In `settings`, under the `stream`
tab, select `custom` for the `service`, enter `rtmp://<ip
address>:1935/myapp`, and `test` for the `stream key`.

Finally setup a screen capture source, start VLC with the desired video,
and click the `start streaming` button in OBS. Its as easy as that!
