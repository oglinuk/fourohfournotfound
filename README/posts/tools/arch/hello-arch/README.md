---
title: Hello Arch Linux!
date: 2020-04-15T06:15:42+12:00
draft: false
tags: ["Hello World,", "Linux,", "Arch,", "Open Source"]
---

Arch Linux is one of those things that I've known about for a long time but have never really gotten around to exploring. I have used manjaro before, which was introduced to me by my friend [Sam](https://github.com/pigeonhands), but have never used/installed vanilla Arch. This blog is going to cover the steps to setup vanilla Arch using [this installation guide](https://github.com/D3S0X/arch-installation) which was suggested to me by [zer0the0ry](https://gitlab.com/zer0the0ry).

Before starting I should note that the guide suggests to be in UEFI BIOS mode, but I am going to be using legacy for the first installation.

---

Since this is my first time installing vanilla Arch, I am going to be using [virtualbox](https://www.virtualbox.org/), which can be installed via 

`sudo apt install virtualbox`

Next to download the vanilla Arch ISO which is available on the [archlinux downloads](https://www.archlinux.org/download/) page. As [rwxrob](https://gitlab.com/rwxrob) suggests, you should verify the SHA1/MD5 checksums and the [PGP](https://en.wikipedia.org/wiki/Pretty_Good_Privacy) signature of the ISO before continuing.

The MD5 checksum can be verified via the `md5sum` command and likewise the SHA1 checksum can be verified via the `sha1sum` command.

To verify the PGP signature we need to download the `archlinux-2020.04.01-x86_64.iso.sig` file which can be found in the same place as the ISO. Run the following command.

`gpg --verify archlinux-2020.04.01-x86_64.iso.sig`

Because we don't have the public key we have to run 

`gpg --recv-key 4AA4767BBC9C4B1D18AE28B77F2D434B9741E8AC`, then run the above command again.

The output should be 

```
gpg: Signature made Wed 01 Apr 2020 16:38:31 NZDT
gpg:                using RSA key 4AA4767BBC9C4B1D18AE28B77F2D434B9741E8AC
gpg: Good signature from "Pierre Schmitz <pierre@archlinux.de>" [unknown]
gpg: WARNING: This key is not certified with a trusted signature!
gpg:          There is no indication that the signature belongs to the owner.
Primary key fingerprint: 4AA4 767B BC9C 4B1D 18AE  28B7 7F2D 434B 9741 E8AC
```

---

Once archiso has booted, the first thing to do is an update && upgrade via `pacman -Sy`.

Next to change the root password via `passwd`.

I am on my desktop, which is connected via ethernet so I will be skipping the step to connect to WiFi via `wifi-menu`.

The next command the guide suggests is `timedatectl set-ntp true` to make sure the systems clock is correct.

To check the BIOS mode, enter the command `ls /sys/firmware/efi/efivars`. If the directory does not exist it means the legacy BIOS mode is active. As I stated at the start the guide suggests UEFI, but I am using legacy.

---

**Now to partition.**

To find the drive to be paritioned run the `fdisk -l` command. In my case the drive is called `/dev/sda`, so I ran the command `fdisk /dev/sda`.

The guide states that if you are using BIOS, the partition table type used should be `DOS`. To create a new partition table enter `o`. Then to create a new partition enter `n`, choose a partition number, the first sector, and finally the last sector. I left mine all default as I am not going to bother with a swap file for my first installation. One last thing is to set the bootable flag via `a`. Finally enter `w` to write and exit.

Once the partitioning is done its time to format the partition via the following commands.

`mkfs.ext4 -L ROOT /dev/sda1`

Remember to change the `/dev/sda1` value to whatever drive and partition number you are using. 

I've come across the `mkfs` command before when using `dd` to create a live USB, but haven't really explored what the command actually does. `man mkfs` shows that the `mkfs` command is used for building a Linux filesystem. It should also be noted that it states in the man page "**this mkfs frontend is deprecated in favour of filesystem specific mkfs.type utils.**", which would explain why the command is`mkfs.ext4`. Looking up `man mkfs.ext4` confused me a bit at first as it comes up with `mke2fs` which is a command to create an ext2/ext3/ext4 filesystem. It made more sense when I read the part in the description stating "If `mke2fs` is run as `mkfs.XXX` (i.e., `mkfs.ext2`, `mkfs.ext3`, or `mkfs.ext4`) the option `-t XXX` is implied..." The `-t` flag is to specify the filesystem type. `-L` sets the label of the volume, in this case `ROOT`. 

Very interesting and always fun to learn something new!

After the formatting is done, it's time to mount the filesystem via the following.

`mount /dev/sda1 /mnt`

---

**Time for the base installation.**

The first thing the guide suggests is to "rank the mirrors before for faster downloads". I didn't quite understand what this meant, but after a quick Lynx search for `archlinux-keyring reflector` I came across the [Reflector Arch-Wiki](https://wiki.archlinux.org/index.php/Reflector) page. It states "Reflector is a script which can retrieve the latest mirror list from the MirrorStatus page, filter the most up-to-date mirrors, sort them by speed and overwrite the file /etc/pacman.d/mirrorlist." 

Since I live in New Zealand this is actually pretty handy.

To rank the mirrorlist I ran the following commands.

`pacman -Sy archlinux-keyring reflector`

`reflector --country 'New Zealand' --age 15 --protocol https --sort rate --save /etc/pacman.d/mirrorlist`

Next to start the installation via the following.

`pacstrap /mnt base base-devel linux linux-firmware linux-lts sysfsutils usbutils e2fsprogs inetutils netctl vi less which man-db man-pages intel-ucode`

*Can you spot the difference to the guide in the above command (other than intel-ucode since im using an Intel CPU)?*

After the packages have finished installing the next thing to do is to create the filesystem table via 

`genfstab -U /mnt >> /mnt/etc/fstab`

Finally to change to the installed system via `arch-chroot /mnt`.

---

**Onto installing the bootloader**

Run the following commands.

`pacman -S grub os-prober`

`grub-install --target=i386-pc --recheck /dev/sda`

`grub-mkconfig -o /boot/grub/grub.cfg`

---

**Configuring the system.**

If you guessed the above question correctly you will have noticed I use vi and not nano.

First thing to do is to setup the hostname via the following.

`echo oglinuk > /etc/hostname` (replace oglinuk with your desired name)

then add the following lines to `/etc/hosts`.

`
127.0.0.1 	localhost
::1		localhost
127.0.1.1	oglinuk.localdomain	oglinuk
`

Now that the hostname is setup its time to setup the locale. First to uncomment any language relevant in `/etc/locale.gen`. In my case I am going to uncomment `en_NZ.UTF-8 UTF-8` and `en_NZ ISO-8859-1`. After the changes have been made run the command `locale-gen`.

Once the locales have been generated, the next thing to do is to set the locale via the following.

`echo LANG=en_US.UTF-8 > /etc/locale.conf`

`export LANG=en_US.UTF-8`

and set the keymap via

`echo KEYMAP=us > /etc/vconsole.conf`

After the locale has been setup, the next thing to do is to setup the time && date via the following.

`ln -sf /usr/share/zoneinfo/NZ /etc/localtime` (ensure you replace NZ with your country)

`hwclock --systohc --utc`

With time && date setup, on to setting up multilib, which is needed for gaming. Uncomment the following lines in `/etc/pacman.conf`.

`
[multilib]
Include = /etc/pacman.d/mirrorlist
`

then run `pacman -Syu`

---

**Setting up a User**

We setup the root password at the start, so I skipping that step. To create a new user execute the following.

`useradd -m -g users -G wheel,audio,video,storage,power,input,optical,sys,log,network,floppy,scanner,rfkill,lp,adm -s /bin/bash oglinuk`

Replace `oglinuk` with your desired username.

Then run `passwd oglinuk`.

Next enable sudo via `EDITOR=vi visudo`, then uncommenting the following line.

`%wheel ALL=(ALL) ALL`

---

**Installing a Desktop**

First to install a display server via the following.

`pacman -S xorg-server xorg-xinit xorg-xrandr xorg-xfontsel xorg-xkill`

Next for the desktop environment. I am going to be using `Xfce`.

`pacman -S xfce4 xfce4-goodies`

For the display manager I am going to use `SDDM`.

`pacman -S ssdm`

`systemctl enable sddm`

---

**Installing Useful Packages**

*General packages*

`pacman -S linux-headers linux-lts-headers dkms jshon expac git curl wget acpid avahi net-tools xdg-user-dirs alarcitty lynx tmux go docker`

I am not using an SSD so I am skipping the step for running `systemctl enable --now fstrim.timer`.

*Media packages*

`pacman -S gst-libav gst-plugins-base gst-plugins-good gst-plugins-bad gst-plugins-ugly gstreamer-vaapi gst-transcoder x264 x265 lame`

I am not going to bother with printer support as I have no use for it.

*Graphics drivers*

`pacman -S mesa lib32-mesa`

`pacman -S vulkan-icd-loader lib32-vulkan-icd-loader`

I have a NVIDIA 1050 ti so I am going to install the NVIDIA drivers.

`pacman -S nvidia nvidia-utils lib32-nvidia-utils`

*Gaming packages*

`pacman -S vkd3d lib32-vkd3d faudio lib32-faudio`

*Networking packages*

`pacman -S networkmanager networkmanager-openvpn networkmanager-pptp networkmanager-vpnc`

`systemctl enable NetworkManager`

*Archive and filesystem util packages*

`pacman -S p7zip unrar unarchiver unzip unace xz rsync nfs-utils cifs-utils ntfs-3g exfat-utils`

*Sound packages*

`pacman -S alsa-utils pulseaudio-alsa pulseaudio-equalizer`

Since I am using the `Xfce` desktop environment, I am going to install the control app for the GTK desktop.

`pacman -S pavucontrol`

The guide suggests to comment out the `load-module module-role-cork` line in `/etc/pulse/default.pa` file.

---

**Reboot**

`exit`

`umount -R /mnt`

`telinit 6`

`telinit` is a command I've never seen before. `man telinit` shows that the command `telinit` is used to change the SysV system runlevel. The `6` option is to reboot the machine. The man page does state that "since the concept of SysV runlevels is obsolete the runlevel requests will be transparently translated into systemd unit activation requests." Cool stuff!

---

**Post Installation**

The guide suggest to set the X11 keymap via `localectl set-x11-keyman us`. Make sure you replace `us` with your desired keymap.

Now to setup the AUR helper, which we will be using [yay](https://github.com/Jguer/yay), an AUR helper written entirely in Golang.

`git clone https://aur.archlinux.org/yay.git`

`cd yay && makepkg -rsi`

`cd .. && rm -rf yay`

---

That's all I am going to cover as I am not interested in themes or any GUI stuff. Feel free to continue following the guide if you are wanting to customize the themeing. Hope you enjoyed and now to play around with the fresh Arch install! Shout out to [D3S0X](https://github.com/D3S0X) for the awesome guide and to [zer0the0ry](https://gitlab.com/zer0the0ry) for the suggestion!
