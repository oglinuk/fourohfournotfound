# Hello Shell!

I first made the switch to Linux back in 2017, and have enjoyed every moment since. Like most, I found the terminal to be a daunting thing at first, much like Linux itself. I quickly discovered how much more I enjoyed using a terminal over GUI applications. Though I have gotten by over the past years with the basic commands like `cd` or `rm`, I never really dedicated time to improving my skills with the command line. 

This is going to be the first of many blogs aimed to record my learnings and discoveries with the Linux shell. Before I go any further I just want to stop and take a moment to recognize the many who have contributed to projects that have shaped the world as we know it today. Thank you. You are amazing and I will forever remember what you have done for not only me but all of those who use the things you have helped create.

One of those wonderful humans is [William Shotts](http://www.linuxcommand.org), who is the author of the book [The Linux Command Line](http://www.linuxcommand.org/tlcl.php). This is a freely available book (in pdf form) under the Creative Commons Attribution-Noncommercial-No De-rivative Works 3.0 license. If you ever read this blog William, thank you.

Now I can't think of a better reason to learn the shell than the one that William gives on his website. 

"A few years ago we had a problem where I used to work. There was a shared drive on one of our file servers that kept getting full. I won't mention that this legacy operating system did not support user quotas; that's another story. But the server kept getting full and it stopped people from working. One of our software engineers spent the better part of a day writing a C++ program that would look through all the user's directories and add up the space they were using and make a listing of the results. Since I was forced to use the legacy OS while I was on the job, I installed a Linux-like command line environment for it. When I heard about the problem, I realized I could do all the work this engineer had done with this single line: 

`du -s * | sort -nr > $HOME/user_space_report.txt`." [[source]](http://www.linuxcommand.org/lc3_learning_the_shell.php)

Not only did this make me laugh, but this also made me realize just how *powerful* the shell can be. Lets take this command as the first to deconstruct and understand. `man du` shows that the `du` command is used to summarize the disk usage of a set of files, recursively for directories. The `-s` flag is to display only a total for each argument, which in this case is the wildcard `*`. The `|` means that we pipe the output of the first command to the command that follows after the `|`. `man sort` shows that `sort` is used to write the sorted concatenation of all file(s) to standard output. The `n` flag is to compare according to string numerical value, and the `r` flag is to reverse the result of the comparisons. The result of those two commands are then redirected via `>` to a file called `user_space_report.txt`.

Truly amazing.