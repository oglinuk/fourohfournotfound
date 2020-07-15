---
title: Nginx 2 ~ Enabling HTTPS
date: 2020-04-23T17:51:34+12:00
draft: false
tags: ["Nginx", "Digital Ocean", "HTTPS", "Web Development"]
---

In the [previous nginx blog](/post/nginx-1) I covered how to host a static website on the droplet that also serves my simple Go web application. This blog is going to cover how to enable HTTPS via [LetsEncrypt](https://letsencrypt.org/) for the site. I am going to be following a medium blog by [jgefroh](https://medium.com/@jgefroh/a-guide-to-using-nginx-for-static-website-d96a9d034940). Since I already have the site up and running I am going to skip most of the content leading up to enabling HTTPS.

The first thing to do is to install [certbot](https://github.com/certbot/certbot) via the following commands.

```BASH
sudo apt install software-properties-common
sudo add-apt-repository ppa:certbot/certbot
sudo apt update
sudo apt install python-certbot-nginx
```

Then to run certbot via `sudo certbot --nginx` and follow the setup.

One reason I like jgefroh's blog in particular is because he covers how to setup auto-renewal via [cron](https://en.wikipedia.org/wiki/Cron). According to the [Let's Encrypt FAQ page](https://letsencrypt.org/docs/faq/#a-id-technical-technical-questions-a), certificates are valid for 90 days, but they recommend renewing certificates every 60 days. To add a new job to for cron execute `sudo crontab -e`. Then add the following.

```
0 0 1 * * certbot renew --post-hook "systemctl reload nginx"
```

The above will renew the certificate every month.

---

Now that the certificate is ready and the renewal is in place, it's time configure nginx. Inside the server block add the following.

```
server {
	listen 443 default_server;
	listen [::]:443 default_server;

	# content ...

	ssl on;
	ssl_certificate /etc/letsencrypt/live/fourohfournotfound.com/fullchain.pem;
	ssl_certificate_key /etc/letsencrypt/live/fourohfournotfound.com/privkey.pem;
}
```

Then to redirect any HTTP requests to HTTPS by adding another server block below the server block with the ssl configuration. This server block will look like the following.

```
server {
	listen 80;

	server_name fourohfournotfound.com www.fourohfournotfound.com;

	rewrite ^ https://$host$request_uri? permanent;
}
```

Finally to run `sudo nginx -s reload` and to test by going to [http://fourohfournotfound.com](http://fourohfournotfound.com).

---

Since I have my blog as a subdomain there is one last thing to do, which is adding the ssl configuration to the `blog.fourohfournotfound.com` nginx configuration file. Everything is the same except for the `listen` line. The only difference is not having `default_server`. The file looks like the following.

```
server {
	listen 443; # only difference
	listen [::]:443; # only difference

	# content ...

	ssl on;
	ssl_certificate /etc/letsencrypt/live/fourohfournotfound.com/fullchain.pem;
	ssl_certificate_key /etc/letsencrypt/live/fourohfournotfound.com/privkey.pem;
}

server {
	listen 80;

	server_name blog.fourohfournotfound.com www.blog.fourohfournotfound.com;

	rewrite ^ https://$host$request_uri? permanent;
}
```

***Spicy***. 
