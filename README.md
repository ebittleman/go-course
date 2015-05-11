# Go Course

## Live Site

[Go Course 2015](http://gocamp-2015.appspot.com/ "Go Course 2015") hosted on Google Appengine

## Required Software

After installing make sure the programs are available in the environment PATH

* [Git](http://git-scm.com/)
* [Mercurial](http://mercurial.selenic.com/)
* [Python](https://www.python.org/)
* [Go](https://golang.org/dl/)
* [Go AppEngine SDK](https://cloud.google.com/appengine/downloads)

## Recommended Text Editor

* [Sublime Text](http://www.sublimetext.com/)
* [GoSublime](https://packagecontrol.io/packages/GoSublime)
* [goimports](https://github.com/bradfitz/goimports)

### Setting up goimports for GoSublime

1. Install Sublime Text

    Download from: http://www.sublimetext.com/

2. Install the Sublime Text Package Manager for your version

   Follow directions for your version (2 or 3) here: https://packagecontrol.io/installation

3. Install GoSublime

    Follow installation instructions here: https://packagecontrol.io/packages/GoSublime

4. Get goimports, run the following

    go get golang.org/x/tools/cmd/goimports

5. Set go imports to run when GoSublime does formatting

    Follow instructions here: http://blog.campoy.cat/2013/12/integrating-goimports-with-gosublime-on.html

## Install the presenter tool locally

    go get github.com/ebittleman/go-course
    go get golang.org/x/tools/cmd/present

## Run Slides Locally Without Appengine SDK

    cd $GOPATH/github.com/ebittleman/go-course
    present

## Run Slides Locally With Appenging

    goapp serve github.com/ebittleman/go-course
