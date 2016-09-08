# Gekyll
Unlike the static site generator Jekyll, Gekyll uses Go and JSON to make an ultra-fast dynamic blog. Without creating dozens of static sites or relying on a database backend. 

### Installation
```
~$ git clone https://github.com/nmccrory/minimal-blog.git
```
```
$ cd minimal-blog
$ go build .
$ ./minimal-blog
```
## Overview
Blogs are saved in a folder as JSON objects which are then read, stored, and formatted by Go. 

Gekyll provides functions for reading JSON objects and converting them into Go types. As well as functions for formatting type attributes - such as Go's time.Time https://golang.org/pkg/time/#Time - automatic!
[time.Time](https://golang.org/pkg/time/#Time) type.
##