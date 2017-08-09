#!/bin/bash
usage="$(basename "$0")  --Shell caller for install_cli.sh, prompts for apl username and password"
echo "Running the appLariat CLI installer"

#Determine Linux or OSX
if [[ "$OSTYPE" == "darwin"* ]]; then
    OS_TYPE=darwin
else
	OS_TYPE=linux
fi

APL_API="https://api.applariat.io/v1/"
FIND_LATEST="https://api.github.com/repos/applariat/go-apl/releases/latest"
CMD_DIR=/usr/local/bin
SCRIPT_DIR="$HOME/apl_scripts"
CONFIG_DIR="$HOME/.apl"
CONFIG_FILE="config.toml"

DOWNLOAD_URL=$(wget -qO- ${FIND_LATEST} | grep browser_download_url | grep ${OS_TYPE} | head -n 1 | cut -d '"' -f 4)
APL_CLI_VERSION=$(echo $DOWNLOAD_URL | awk -F"/" '{print $(NF - 1)}')
APL_FILE=$(echo $DOWNLOAD_URL | awk -F"/" '{print $NF}')

if [ ! -f ${APL_FILE} ]; then
    echo "Downloading cli package ${APL_FILE}"
    wget -q $DOWNLOAD_URL
fi

if [ ! -f "bin/apl" ]  || [ ! -d scripts ]; then
	tar zxf ${APL_FILE}  --exclude=$(basename "$0")
fi
#Place bundle files
echo "Moving apl to /usr/local/bin"
mv -f bin/apl $CMD_DIR
rm -rf bin 
if [ ! -d "$CONFIG_DIR" ]; then
    mv ./scripts ${SCRIPT_DIR}
fi

#Configuring APL CLI
echo "Configuring APL CLI"
if [ ! -d "$CONFIG_DIR" ]; then
    mkdir -p $CONFIG_DIR
fi
config="n"
echo
read -p "Do you want to configure access to appLariat now [y/n]?: " config

if [[ $config = "y" ]]; then	
    echo "Let's gather the info we need to configure apl"
    echo
	read -p "Type in your appLariat username: " user
	echo
	read -p "Type in your appLariat password: " pass
    
    echo "Setting APL API value: $APL_API"
    echo "Setting APL User/Password: $user : *********"
else
	echo "To manually configure access update the file ${CONFIG_DIR}/${CONFIG_FILE}
	and provide values for svc_username and svc_password"
	user=""
	pass=""
fi
    
if [[ $config = "y" ]] || [ ! -f ${CONFIG_DIR}/$CONFIG_FILE} ]; then
	#Create the config file
	cat >$CONFIG_FILE <<EOL
api = "${APL_API}"
svc_username = "${user}"
svc_password = "${pass}"
EOL
    
	#Move and secure the config file
	mv -f $CONFIG_FILE $CONFIG_DIR
	chmod 600 ${CONFIG_DIR}/${CONFIG_FILE}
fi

echo
echo "You are ready to use the APL CLI"
echo "See the helper scripts downloaded to ${SCRIPT_DIR} for examples"
echo "Run apl -h to see command options"
#apl -h

#Cleaning up
rm -f ${APL_FILE}
rm -f apl-*.tgz

echo "APL CLI Installation Complete"



