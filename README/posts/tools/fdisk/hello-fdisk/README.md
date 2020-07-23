---
title: Hello fdisk!
date: 2018-11-20
tags: ["Hello World,", "fdisk,", "Linux,", "Partitioning,", "Open Source"]
draft: false
---
[Partitioning?](https://en.wikipedia.org/wiki/Partition) No not mathematics partitions, computer partitions. More specifically hardware partitions (in my case a microSD card). I am officially put off GUIs and am taking a hiatus from them. My raspberry pi's microSD card isn't being used at the moment, so I am formatting it to be used as storage for pcap files that I am capturing on my network. Perfect time to learn how to partition drives from the terminal! After a quick ddg search, the first link was about `fdisk`. Since the microSD card was accessible from my omega2 I wanted to see if I could install fdisk using opkg on the omega2 and format the microSD card from there. `opkg install fdisk` ...

It works! Eureka!

First thing to do is list the partitions via `fdisk -l`. I don't need to use `sudo` as the omega2 operates as root.

Next to enter command mode via `fdisk <partition>`

From the help argument (`h`) we can see the available commands. 

`p` to print the partition table.

`d` to delete the partition(s).

`n` to create a new partition table

`w` to write the changes to the drive

Pretty straight-forward. 
