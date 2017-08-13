# Imager.go
### An application to download random public domain images in different dimensions automatically into your disk.
![Imager.go downloading images](https://raw.githubusercontent.com/rakibtg/Imager.go/master/static/Screen%20Shot%202017-08-13%20at%209.02.30%20PM.png "Imager.go running")

You might be working on a new design or building a working prototype and you need a tons of images in different dimensions. 

Now collecting those images from the internet and saving them one after another is difficult, it did'nt work i have tried.
Okay, here is Imager.go app for y'all 😉

**Imager.go** will download as many images as many you want from Unsplash right into your disk. It also allow you to set a range 
of images width and height.

**How to run the application?**
If you do not care about the code that powers Imager.go then download the application.
- Mac OS https://github.com/rakibtg/Imager.go/blob/master/bin/ImagerLinux?raw=true
- Linux 64bit https://github.com/rakibtg/Imager.go/blob/master/bin/ImagerLinux?raw=true
- Windows 64bit https://github.com/rakibtg/Imager.go/blob/master/bin/ImagerLinux?raw=true

Open the downloaded file in a terminal.

- Open terminal
- Go to the directory where you have downloaded the application `cd /IMager`
- Run the app with this command `./ImagerMac` if you are on a mac. Or `./ImagerWindows` or `./ImagerLinux`

**If you want to run the app in a Go-lang env**
- Clone this repo
- `go run Imager.go downloader.go imageRandomUrl.go intro.go takeInput.go` run this in a terminal

**If you want to build**
- `go build Imager.go downloader.go imageRandomUrl.go intro.go takeInput.go` in a terminal.

Pull Requests are welcome ❤️💕

Tweet me your thoughts <a href="https://twitter.com/rakibtg">@rakibtg</a>
