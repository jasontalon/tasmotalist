# Tasmotalist

A CLI program for searching tasmota devices over the connected network. Built for Windows and Mac.
![img_1.png](img_1.png)

![img_2.png](img_2.png)
## Motivation
I have a number of tasmota devices in my  house. There are times that I need to check if they are online. Previous way of checking is going through the phone and open an app called Fing to scan my network and see the connected devices.

Once it identifies the devices, then its up to me to guess which devices are Tasmotas, which is quite cumbersome and time-consuming. This is the reason why I created this small program to search for tasmota devices on the connected network. 

## Requirements

The program is using another utility called [nmap](https://nmap.org/) to scan available devices over the connected network.

See below for the installation procedure for respective platform

* Windows:
    * [Official Installer](https://nmap.org/download.html#windows)
    * [Chocolatey](https://community.chocolatey.org/packages/nmap#install)
* Mac:
    * [Official Installer](https://nmap.org/download.html#macosx)
    * [Homebrew](https://formulae.brew.sh/formula/nmap)