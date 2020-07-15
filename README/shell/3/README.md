---
title: Shell 3 ~ Exploring the System
date: 2020-04-29T19:57:23+12:00
draft: false
tags: ["Shell,", "Linux,", "CLI,", "BASH,", "Open Source"]
---

Chapter 3 is about *exploring the system*. The commands covered are `ls` (list directory content), `file` (determine file type), and `less` (view file contents one screenful at a time). I use `ls` daily, but have to admit I've never used `file` or `less`. As the book states, `ls` is one of the most used commands. Hard to disagree since it's whats used to see a directories content is and the associated attributes. The wonderful thing about `ls` is the fact that you can use it on directories other than the current working directory. An example of this would be `ls /var/log`. One thing I did not know about `ls` is that you can list multiple directories. An example would be `ls /var/log /var/backups`. To view more detail (or long listing format according to the man page) we can specify the `-l` flag, which looks like `ls -l`.

This is where the book first mentions *options && arguments*. What I called a *flag* is what the book refers to as an *option*. Commands can often have their behavior changed by adding one or more *options*, and options can be changed even further by adding *arguments*. The structure looks like the following.

`command -option(s) argument(s)`

Options can consist of either a single character that is preceded by a dash (like `ls -l`), or they can be a word preceded by two dashes (like `ls --indicator-style=slash`). You can also combine options to allow for a more specific output, like `ls -lh`, which lists the content directory using long listing format and in human readable sizes. Like naming in Linux, command options are case-sensitive. An example of this is `ls -a` and `ls -A`. The `-a` option lists all files (including hidden ones), where as the `-A` option is the same as `-a` except it does not show hidden files.

Let's break down the output of `ls -l`. I will be using the example output from the book, which is as follows.

```
-rw-r--r-- 1 root root 3576296 2017-04-03 11:05 Experience ubuntu.ogg
```

The first field (`-rw-r--r--`) are the permissions. The first grouping of `rw-` is the owner permissions, meaning the permissions apply only to the owner of the file. The second grouping `r--` is the group permissions, which applies to the group that has been assigned to the file. The third and last grouping (also `r--`) is the permissions for all users on the system. You might have noticed that I skipped over something. I never knew what this particular *thing* was until now. The thing I am referring to is the first character in `-rw-r--r--`, or the first `-`. If you execute `ls -l` you will most likely see a `d` in place of the `-`. This refers to the `type` of the file, which in the case of `d` the file type is a directory. 

The following are types that can be found on Linux. [[source]](https://unix.stackexchange.com/a/59133)

* - (file)
* d (directory)
* c (character device)
* l (symlink)
* p (named pipe)
* s (socket)
* b (block device)
* D (door)

The second field (`1`) is the number of hard links associated with the name. The third field (`root`) is the username of the owner of the file. The fourth field (`root`) is the name of the group that owns the file. The fifth field (`3576296`) is the size in bytes. The sixth field (`2017-04-03 11:05`) is the date and time of the last modification. The seventh and final field (`Experience ubuntu.ogg` is the name of the file.

---

Now lets explore the `file` command. In Linux the file's name is not required to reflect the type of the file itself. For an example lets look at the following file.

```HTML
<!DOCTYPE html>
<html>
  <h1>Hello world!</h1>
</html>
```

Let's first save this file as the name `test.html` and run `file test.html`, we get the following output.

```
test.html: HTML document, ASCII text
```

Now let's rename the file to just `test` and run `file test`. We get the following output.

```
test: HTML document, ASCII text
```

---

As I said at the start `less` is not a command that I have used. Although I've known about it, I usually default to using the `cat` command. An interesting note that the book makes is how *ASCII* is one of the earliest and simplest representations of text. ASCII stands for *American Standard Code for Information Interchange*. Text is a one-to-one mapping of characters to numers, which means that 42 characters of text translates to 42 bytes of data. An example of `less` is `less /var/log/syslog`.

There are a list of directories that the book suggests to explore, which I will list below, but am going to leave the exploring/explanation for you to do/research.

* `/`
* `/bin`
* `/boot`
* `/dev`
* `/etc`
* `/home`
* `/lib`
* `/lost+found`
* `/media`
* `/mnt`
* `/opt`
* `/proc`
* `/root`
* `/sbin`
* `/tmp`
* `/usr`
* `/usr/bin`
* `/usr/lib`
* `/usr/local`
* `/usr/sbin`
* `/usr/share`
* `/usr/share/doc`
* `/var`
* `/var/log`

---

The last part of the chapter talks about *links*, more specifically *symbolic links*. During your exploration of the various directories you will have come across files that look like the following.

```
lrwxrwxrwx 1 root root   9 Apr 17 11:24 editor -> /bin/nano
```

If we refer back to the types of files that can be found on Linux, then we will know that this is a file of the type *symlink*, which can also be referred to as a *symbolic link* or *soft link*. Unix-like systems have the ability to have a file referenced by multiple names, which may not be obvious as to why that would be useful. My first encounter with symlinks was when I got started with [nginx](/post/nginx-0), when having to run the `ln -s` command to create a symbolic link between `fourohfournotfound.com` in `sites-available` to `sites-enabled`. It would be a pain to have to manually update `sites-enabled/fourohfournotfound.com` every time I make a change to the `sites-available/fourohfournotfound.com` file. This is where symbolic links come in. Rather than having two seperate files, we instead create a symbolic link in `sites-enabled` called `fourohfournotfound.com`, which points to the `sites-available/fourohfournotfound.com` file. Now anytime a change is made to the `sites-available/fourohfournotfound.com` file, `sites-enabled` reflects those changes.

The example that I have shown is a symbolic link located in `/etc/alternatives` for `editor` pointing to `/bin/nano`. This refers to the default editor used, which in this case, resolves to `/bin/nano`. We can change this symbolic link by executing `sudo update-alternatives --config editor`, and choosing `/usr/bin/vim.basic`. Now if we look at the symbolic link again we can see it is now the following.

```
lrwxrwxrwx 1 root root  12 April 29 19:15 editor -> /usr/bin/vim.basic
```

That's all for chapter 3!
