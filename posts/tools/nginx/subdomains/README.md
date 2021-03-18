# Subdomains

In the [previous blog](/tools/nginx/hello-nginx/index.html) we explored serving a Go web application with [Nginx](https://nginx.com) on [Digital Ocean](https://digitalocean.com). This blog is going to cover how to host a static website on a subdomain. One of the things I did not cover in the last blog is how to point to your domain name with Digital Ocean. For more on this please refer to [this blog](https://www.digitalocean.com/community/tutorials/how-to-point-to-digitalocean-nameservers-from-common-domain-registrars) that I found/used originally.

The first thing to do is copy over your static site using an `scp` or `rsync` command. If you're wondering what the difference between the two like I did, you can refer to [this stackexchange answer](https://unix.stackexchange.com/a/39719). I copied my static website to the `/var/www` directory, but you can choose to store it wherever you would like, just make sure you use that same path in the Nginx configuration.

The next thing to do is create another Nginx server block file in `/etc/nginx/sites-available`. Add the following to the file. 

```
server {
	root /var/www/fourohfournotfound;

	server_name blog.fourohfournotfound.com www.blog.fourohfournotfound.com;
}
```

Finally, to add a subdomain to point to the droplet. In my case, since my static website is my blog I am adding `blog` and `www.blog` as the hostname directing to the droplet with the static website.

Thanks it! Check out [blog.fourohfournotfound.com](http://blog.fourohfournotfound.com) and learn on!
