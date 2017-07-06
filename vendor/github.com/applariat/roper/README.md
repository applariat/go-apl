# roper
[![Build Status](https://travis-ci.org/applariat/roper.svg?branch=master)](https://travis-ci.org/applariat/roper)

Read input from a file, URL or STDIN and unmarshal it

# why

There are a lot of libraries that read config files or read files into an interface{}. They are typically buried inside a larger project and can't be used independently. I created this to use in the cli for updates/creates, not for loading config files. Similar to what is found in Kubernetes kubectl. Right now it only supports pipes, URLs and JSON/YAML files.

