---
title: Nginx 0 ~ Hello Nginx!
date: 2020-04-12T11:12:33+12:00
draft: false
tags: ["Hello World,", "Nginx,", "Golang,", "Digital Ocean,", "Systemd,", "Web Development"]
---

Nginx is another one of those technologies that I've known of for a while, but haven't had the chance to dig into. This blog will cover how to deploy a basic Golang web application on a $5/month ubuntu 18.04 vm hosted by [Digital Ocean](https://www.digitalocean.com). I've found a really good blog by [Michael Okoh](https://www.digitalocean.com/community/tutorials/how-to-deploy-a-go-web-application-using-nginx-on-ubuntu-18-04) which is what I am going to learn from.

The first thing needed is a Go web application. I've already done this step with a simple web app that returns the IPv4 address of the user and a quote from one of my favorite movies. 

---

Next is to create a systemd unit file. Mine looks like the following.

```
[Unit]
Description=iseeyou

[Service]
Type=simple
Restart=always
RestartSec=5s
ExecStart=/root/iseeyou/iseeyou
StandardOutput=syslog
StandardError=syslog

[Install]
WantedBy=multi-user.target
```

There was an issue with `tpl = template.Must(template.ParseGlob("templates/*"))`, and I didn't dig into it much, but fixed my changing the path to `"/root/iseeyou/templates/*"`. My assumption is the service looks for a `templates` dir within `/lib/systemd/system` where the `iseeyou.service` file is. 

---

With the service and application running, now for Nginx. A simple `sudo apt install nginx` and it's installed. Next to create a file in `/etc/nginx/sites-available` called `fourohfournotfound.com`. In this file we want to specify the following.

```
server {
	server_name fourohfournotfound.com www.fourohfournotfound.com;

	location / {
		proxy_pass http://localhost:<YOUR SERVERS PORT>;
		proxy_set_header Host $host;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
	}
}
```

Finally run

`sudo ln -s /etc/nginx/sites-available/fourohfournotfound.com /etc/nginx/sites-enabled/fourohfournotfound.com`

&&

`sudo nginx -s reload`

Make sure you replace fourohfournotfound.com with your own domain.

---

Goto [fourohfournotfound.com](http://fourohfournotfound.com) and enjoy!
