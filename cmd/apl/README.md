# appLariat (apl) Command Line Interface

The appLariat Command Line Interface (CLI) is a unified tool to manage your appLariat service. You can control your appLariat services from the command line and automate them through scripts.


## Installing the apl CLI

Download the apl.sh script from [https://github.com/applariat/go-apl](https://github.com/applariat/go-apl)

The script performs the following
* Downloads the binary from the [releases page](https://github.com/applariat/go-apl/releases).
* Downloads some helper scripts we have created as examples of using the CLI
* Extracts the binary and moves it to /usr/local/bin
* Creates the apl access config file


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
