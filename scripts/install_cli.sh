#!/bin/bash
set -e # Exit with nonzero exit code if anything fails
usage="$(basename "$0") [-h] [-u str] [-p str] [-o] -- installer for the appLariat CLI

Where:
    -h        -show this help text
    -u str    -applariat username (use quotes on str), required when -o set or first time setup
    -p str    -applariat password (use quotes on str), required when -o set or first time setup
    -o        -overwrite an existing config file
"

# Script vars
export APL_API="https://api.applariat.io/v1/"
APL_CLI_RELEASE=${APL_CLI_RELEASE:-v0.2.0}
CMD_DIR="/usr/local/bin"
SCRIPT_DIR="$HOME/applariat"
CONFIG_DIR="$HOME/.apl"
CONFIG_FILE="config.toml"
DOWNLOAD_URL="https://github.com/applariat/go-apl/releases/download"
INSTALL=false


#Command Line Variables - see usage
while getopts ":u:p:oh" opt; do
  case $opt in
    u)
      APL_USER=$OPTARG
      ;;
    p)
      APL_PASS=$OPTARG
      ;;
    o)
      OVERWRITE="true"
      ;;
    h)
      echo
      echo "$usage"
      exit
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      echo 
      echo "Usage: $usage" >&2
      exit 1
      ;;
    :)
      echo "Option -$OPTARG requires an argument." >&2
      echo
      echo "Usage: $usage" >&2
      exit 1
      ;;
  esac
done

set +e

#Set the config for APL CLI
if [ -f "${CONFIG_DIR}/${CONFIG_FILE}" ] && [ -z $OVERWRITE ]; then
    echo "CLI configuration exists, proceeding"
elif [ ! -f "${CONFIG_DIR}/${CONFIG_FILE}" ] || [ $OVERWRITE == "true" ]; then
    if [ -z $APL_USER ] || [ -z $APL_PASS ]; then
        echo "CLI configuration not found or overwriting"
        echo "Username and password required to create config"
        echo
        echo "Usage: $usage" >&2
        exit 1
    fi
    echo "Setting APL API value: $APL_API"
    echo "Setting APL User/Password: $APL_USER : *********"
    
    if [ ! -d "$CONFIG_DIR" ]; then
        mkdir -p $CONFIG_DIR
    fi
    
    #Create the config file
    cat >$CONFIG_FILE <<EOL
api = "${APL_API}"
svc_username = "${APL_USER}"
svc_password = "${APL_PASS}"
EOL
    
    #Move and secure the config file
    mv -f $CONFIG_FILE $CONFIG_DIR
    chmod 600 ${CONFIG_DIR}/${CONFIG_FILE}
fi

APL_FILE=apl-${APL_CLI_RELEASE}-linux_amd64.tgz
if [[ "$OSTYPE" == "darwin"* ]]; then
    APL_FILE=apl-${APL_CLI_RELEASE}-darwin_amd64.tgz
fi

echo "Checking apl CLI"
if [ `command -v apl` ]; then
    #check the version
    CUR_VERSION=( `apl version` )
    if [[ "${CUR_VERSION[2]}" != "${APL_CLI_RELEASE}" ]]; then
        echo "Older version of apl CLI installed, upgrading!"
        INSTALL=true
    else
        echo "$CUR_VERSION is the current version of apl CLI"
    fi
else
    INSTALL=true
fi
echo

if [[ ${INSTALL} == true ]]; then
	if [ ! -f ${APL_FILE} ]; then
    	echo "Downloading cli package ${APL_FILE}"
    	wget -q ${DOWNLOAD_URL}/${APL_CLI_RELEASE}/${APL_FILE}
	fi
    echo "Installing apl command"
    tar zxf ${APL_FILE}
    mv -f ./apl $CMD_DIR
    rm -f ${APL_FILE}
fi

echo

# Testing to make sure it works
echo "Let's get a list of applications to make sure it works"
APL_TEST=$(apl stacks)

echo
echo "Results:"
echo "${APL_TEST}"
if [ $? -ne 0 ]
then
    echo "Error: Darn! Something went wrong, contact appLariat support for assistance, exiting"
    exit 1
fi

echo
echo "You are ready to use the APL CLI"
echo "See the helper scripts downloaded to ${SCRIPT_DIR}/${SCRIPTS} for examples of usage"
echo "Run apl -h to see command options"
#apl -h

#Cleaning up
rm -f ${APL_FILE}

echo "APL CLI Installation Complete"