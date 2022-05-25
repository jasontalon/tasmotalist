# Tasmotalist

A CLI program for searching tasmota devices within the local network. Built for Windows and Mac.
![img_1.png](img_1.png)

![img_2.png](img_2.png)
## Motivation
I have a number of tasmota devices in my  house. There are times that I need to check if they are online. My previous way of checking is going through the phone and open an app called Fing to find connected devices within the local network.

Once it identifies the devices, then it's up to me to guess which devices are tasmotas, which is quite cumbersome and time-consuming. This is the reason why I created this small program to search for tasmota devices within the local network. 

## Requirements

The program is using another utility called [nmap](https://nmap.org/) to scan the network.

See below for the installation procedure for respective platform

* Windows:
    * [Official Installer](https://nmap.org/download.html#windows)
    * [Chocolatey](https://community.chocolatey.org/packages/nmap#install)
* Mac:
    * [Official Installer](https://nmap.org/download.html#macosx)
    * [Homebrew](https://formulae.brew.sh/formula/nmap)

## Installation

Once you have installed the mentioned requirements above, you are now ready to use Tasmotalist.

You can find the executable file in the [release page](https://github.com/jasontalon/tasmotalist/releases/tag/v1.0.0)

Or, you may build your own by downloading the [Go SDK](https://go.dev/dl/) and by cloning the repository then run the following 
```
make build
```

## Issues and Suggestions

Feel free to drop any questions, suggestions, and issues to improving the utility.

You may report by using the [Github Issues](https://github.com/jasontalon/tasmotalist/issues)


 
