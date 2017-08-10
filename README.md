# The appLariat golang API client and CLI
[![Build Status](https://travis-ci.org/applariat/go-apl.svg?branch=master)](https://travis-ci.org/applariat/go-apl)

The appLariat API Client and CLI enables users of appLariat to integrate their current application delivery
workflows into appLariat.  appLariat has been designed to provide simple and consistent 
delivery of applications to hybrid cloud environments, leveraging containers and Kubernetes.
Ideal for providing continous delivery and deployment at any stage of development and any size project team.  

The CLI is a work in progress and is currently focused on the CRUD functions related to appLariat 
release and deployment resources. We will continue to improve the features and ease of use. The scripts 
directory contains a set of sample scripts designed to help you get started using the CLI.

We are working on proper user documentation for both the API Client and the CLI.

## Installing the apl CLI

Download the latest archive for your architecture from [releases](https://github.com/applariat/go-apl/releases)

For MacOS and Linux run the following command:
`tar -xzf apl-*.tar.gz; ./apl_install.sh`
 
 or
 
 `wget -q https://github.com/applariat/go-apl/releases/download/v0.2.1/apl_install.sh; bash ./apl_install.sh`

The apl_install.sh script performs the following:
* Checks to confirm latest version
* Moves binary to /usr/local/bin
* Creates the apl access config file

For Windows:
* Extract the contents of the tar bundle
* Move the apl binary into your PATH

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


## Using the apl CLI

Simply run `apl <COMMAND>`

`apl` or `apl --help` will show usage and a list of commands:
```
Usage:
   apl [command]

Available Commands:
  audits           manage audits
  components       manage components
  credentials      manage credentials
  deployments      manage deployments
  events           manage events
  help             Help about any command
  loc-artifacts    manage loc-artifacts
  loc-deploys      manage loc-deploys
  orgs             manage orgs
  policies         manage policies
  policy-results   manage policy-results
  policy-schedules manage policy-schedules
  project-roles    manage project-roles
  projects         manage projects
  releases         manage releases
  roles            manage roles
  stack-artifacts  manage stack-artifacts
  stack-components manage stack-components
  stack-versions   manage stack-versions
  stacks           manage stacks
  types            manage types
  users            manage users
  workloads        manage workloads

Flags:
  -o, --output string   Output format: json|yaml (default "table")

Use " [command] --help" for more information about a command.
```

### Usage

#### Create/Update Commands

The create/update commands requires the -f --file flag. 

`-f, --file string   Input file for create/update: json|yaml`

Examples:

```
apl users create -f user.yaml
# or
apl users update -f user.json
```

#### Output

All commands will print out an abbreviated table of fields. To get the full record you must provide the -o --output flag.

`-o, --output string   Output format: json|yaml (default "table")`

Examples:
```
apl releases --version 1 --output json
apl stacks --name wordpress --output yaml
apl deployments -o json
```

Additional CLI command docs are available in [cmd/apl/docs](https://github.com/applariat/go-apl/tree/master/cmd/apl/docs)


