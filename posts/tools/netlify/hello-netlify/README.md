# Hello Netlify!

Ever wondered what the cheapest and easiest way to make/host a website
is? A Markdown/Git based website hosted with
[Netlify](https://netlify.com). This is also known as a
[Jamstack](https://jamstack.org/what-is-jamstack/) site. The total cost
for this setup (aside from your time) is $0, unless you want to use a
custom domain name, which then it only totals $8.88/m (Namecheap).
Netlify has a free plan which has everything you need, and common hosted
git services (Github, Gitlab, Sourcehut, Gitea, ...) are free. That is
how [this website](https://github.com/oglinuk/fourohfournotfound) is
made/hosted. In this post I am going to demonstrate how easy and quick it
is to create a website and start hosting it on Netlify.

## Create a Website

The first thing to do is to create the website, which as stated above
will be a Markdown based site. If you arent familiar with Markdown, then
I suggest having a look at [the basic
syntax](https://rwx.gg/lang/md/basic/). You will also need an account on
one of the above mentioned hosted Git services sites. I will be using
Github for this post.

Lets start by creating a new git repository for our website.

```
mkdir hello-jamstack
cd hello-jamstack
git init
git remote add origin git@github.com:oglinuk/hello-jamstack.git
```

Next lets create a `README.md` file, which will be our home page, via `vi
README.md`. Now to add some content.

```
# Hello world!

This is a Jamstack website using Markdown, Git, and Netlify!
```

Finally to add the `README.md` file and push it to the remote repository.

```
git add README.md
git commit -m 'Initial commit'
git push origin master
```

## Hosting on Netlify

Once you have created a Netlify account, click on the `New site from Git`
button under the `sites` section. Select `Github` and authorize Netlify.
Next pick the repository, in this case its `oglinuk/hello-jamstack`.
Finally click the `Deploy site` button. By default Netlify will generate
a random URL for your site, but you can use a custom domain name if you
have one. I will cover how to setup the custom domain name a bit later.

If we click on the link to our newly hosted site we will get `Page Not
Found`, this is because Netlify doesnt handle Markdown files. To resolve
this we will need to render the Markdown file as HTML, and we will
achieve this using a build script.

## Add a Build Script

There are various ways we can write the build script, but to keep it
simple, we're just going to write a quick shell script. In the website
repository create a file called `build` and add the following.

```
#!/bin/sh

pandoc README -o index.html
```

If you have never heard of the tool [Pandoc](https://pandoc.org), then I
highly recommend checking it out. Fortunately you dont need to install it
on your machine. Lastly you will need to make the script executable via
`chmod +x build`. Before you commit and push the change, click on the
`Site settings` button of your site and then click on the `Build &
deploy` tab on the left. Click the `Edit settings` button in the `Build
settings` section, add `./build` in the `Build command` section, and
click the `Save` button. Once that is done, commit and push the change.

Now go back to your site and you should see your new website!

The best thing about Netlify is it automatically updates the site and
triggers the build command.

## Add a Custom Domain

If you have a custom domain you can use it by going to the `Domain
Management` section on the left of the `Site settings` section. Then
click the `Add custom domain` button and type in your domain name. You
will need to point to the name servers of your domain name registrar, and
it will take a while to propogate.

Enjoy your new Jamstack site!
