#!/bin/bash
usage="$(basename "$0")  --Shell caller for install_cli.sh, prompts for apl username and password"

#Determine Linux or OSX
if [[ "$OSTYPE" == "darwin"* ]]; then
    OS_TYPE=darwin
else
	OS_TYPE=linux
fi

FIND_LATEST="https://api.github.com/repos/applariat/go-apl/releases/latest"
CMD_DIR=/usr/local/bin

DOWNLOAD_URL=$(wget -qO- ${FIND_LATEST} | grep browser_download_url | grep ${OS_TYPE} | head -n 1 | cut -d '"' -f 4)
export APL_CLI_VERSION=$(echo $DOWNLOAD_URL | awk -F"/" '{print $(NF - 1)}')
APL_FILE=$(echo $DOWNLOAD_URL | awk -F"/" '{print $NF}')

echo "Downloading appLariat CLI package at $DOWNLOAD_URL"
wget -q $DOWNLOAD_URL
tar zxf ${APL_FILE}
mv -f ./apl $CMD_DIR
cp -r ./scripts $HOME/apl_scripts
#mv ./scripts $HOME/apl_scripts
echo
read -p "Type in your appLariat username: " user
echo
read -p "Type in your appLariat password: " pass

#Configuring APL CLI
echo "Configuring APL CLI"
. $HOME/apl_scripts/install_cli.sh -u $user -p $pass



