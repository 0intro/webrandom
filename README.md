[![Build Status](https://travis-ci.org/0intro/webrandom.svg?branch=master)](https://travis-ci.org/0intro/webrandom)

Webrandom
=========

Webrandom provides HTTP services analogous to /dev/null, /dev/random and /dev/zero.

This program is mostly useful as a dummy web server to measure performance of HTTP clients and proxies.

Usage
-----

```
$ webrandom [ -http localhost:8080 ]
```

The following requests are supported:

* POST /null
* GET /random/\<size in bytes>
* GET /zero/\<size in bytes>

Examples
--------

Upload a file:

```
$ curl -X POST -d @rand http://localhost/null
```

Download a 1 GB file of random data:

```
$ curl -o rand http://localhost/random/1048576
```

Download a 1 GB file of zero data:

```
$ curl -o zero http://localhost/zero/1048576
```
