# The appLariat golang api client

[![Build Status](https://travis-ci.org/applariat/go-apl.svg?branch=master)](https://travis-ci.org/applariat/go-apl)

This project is not complete. Do not use.

# Requires the following env vars

* export APL_API="http://localhost:8080/v1/"
* export APL_SVC_USERNAME="yourusername"
* export APL_SVC_PASSWORD="passwd"

# Install

go get github.com/applariat/go-apl/apl


## Configuration via Config File

Create a file called `"$HOME/.apl/config.toml"`

Example:
```
$ cat ~/.apl/config.toml 
api = "https://api.applariat.io/v1/"
svc_username = "your@login-email.com"
svc_password = "passwd"
```


## Configuration via ENV Vars

You can use a configuration file and/or ENV vars. The environment variables will override the config file settings.

```
export APL_API=https://api.applariat.io/v1/
export APL_SVC_USERNAME=your@login-email.com
export APL_SVC_PASSWORD=passwd
```
