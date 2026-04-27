# containers
docker in go! and from scratch

uses apples native Virtualization API.

(no, I don't make the bindings, im not crazy)
(courtesy of github.com/Code-Hex/vz)


Built for Apple Silicon (M4):

steps so far:

download this: https://old-releases.ubuntu.com/releases/20.04.1/ubuntu-20.04-live-server-arm64.iso

extract with unarchiver. 

go into casper folder
rename vmlinuz to vmlinuz.gz and extract
also get the initramfs

those 2 files are the only things you need from the iso

