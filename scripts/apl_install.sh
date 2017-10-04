#!/bin/bash
usage="$(basename "$0")  --Installing the apl CLI, prompts for apl username and password"
echo "Running $usage"

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
APL_ENV="apl.bashrc"

#DOWNLOAD_URL=$(curl -s ${FIND_LATEST} | grep browser_download_url | grep ${OS_TYPE} | head -n 1 | cut -d '"' -f 4)
#APL_CLI_VERSION=$(echo $DOWNLOAD_URL | awk -F"/" '{print $(NF - 1)}')
#APL_FILE=$(echo $DOWNLOAD_URL | awk -F"/" '{print $NF}')

if [ ! -f bin/apl ]  || [ ! -d scripts ]; then
	echo "APL CLI files not found, run this script from the local directory - ./$(basename "$0")"
	exit 1
fi

#Place bundle files
echo "Moving apl to /usr/local/bin"
mv -f bin/apl $CMD_DIR
rm -rf bin 
if [ -d "$SCRIPT_DIR" ]; then
	echo "Detected an existing ${SCRIPT_DIR}, moving in case of user changes"
	[ -d "${SCRIPT_DIR}.old" ] && rm -rf "${SCRIPT_DIR}.old"
    mv ${SCRIPT_DIR} "${SCRIPT_DIR}.old"
    echo "Previous versions of apl scripts are in ${SCRIPT_DIR}.old"
fi
echo "Moving scripts to ${SCRIPT_DIR}"
mkdir -p ${SCRIPT_DIR}
mv ./scripts/* ${SCRIPT_DIR}
rm -rf ./scripts
message=""
#Configuring APL CLI
echo "Configuring APL CLI"
if [ ! -d "$CONFIG_DIR" ]; then
    mkdir -p $CONFIG_DIR
fi

if [[ -f ${CONFIG_DIR}/${APL_ENV} ]]; then
	echo "Found existing ${APL_ENV} file"
else
	echo "Creating an applariat environment file"
	cat >${CONFIG_DIR}/${APL_ENV} <<EOL
#!/bin/bash
export APL_SCRIPT_HOME=${SCRIPT_DIR}
export PATH="${PATH}:${SCRIPT_DIR}"
export APL_LOC_DEPLOY_ID=$LOC_DEPLOY_ID
EOL
	
	if [[ "$OS_TYPE" == "darwin" ]]; then
		echo "Adding applariat environment to .bash_profile"
		cat >>${HOME}/.bash_profile <<EOL2
# added by applariat CLI installer
if [ -f ${CONFIG_DIR}/${APL_ENV} ]; then
    source ${CONFIG_DIR}/${APL_ENV}
fi
EOL2
        message="Run source ~/.bash_profile to update your path"
	else
		echo "Adding applariat environment to .bashrc"
		cat >>${HOME}/.bashrc <<EOL2
# added by applariat CLI installer
if [ -f ${CONFIG_DIR}/${APL_ENV} ]; then
    source ${CONFIG_DIR}/${APL_ENV}
fi
EOL2
	    message="Run source ~/.bashrc to update your path"
	fi

fi

config="n"
echo
read -p "Do you want to configure access to appLariat now [y/n]?: " config

if [[ $config = "y" ]]; then	
    echo
    echo "Let's gather the info we need to configure apl"
    echo
	read -r -p "Type in your appLariat username(email address): " -e user
	echo
	read -rs -p "Type in your appLariat password: " -e pass
    echo
    echo "Setting APL API value: $APL_API"
    echo "Setting APL User/Password: $user : *********"
    echo
elif [ ! -f ${CONFIG_DIR}/${CONFIG_FILE} ]; then
	echo
	echo "To manually configure access to applariat, edit ${CONFIG_DIR}/${CONFIG_FILE}
and provide values for svc_username and svc_password"
	user=""
	pass=""
else
	echo "Skipping configuration"
fi
    
if [[ $config = "y" ]] || [ ! -f ${CONFIG_DIR}/${CONFIG_FILE} ]; then
	#Create the config file
	echo "Writing the config file to ${CONFIG_DIR}/${CONFIG_FILE}"
	cat >${CONFIG_DIR}/${CONFIG_FILE} <<EOL3
api = "${APL_API}"
svc_username = "${user}"
svc_password = "${pass}"
EOL3
    
	#Secure the config file
	chmod 600 ${CONFIG_DIR}/${CONFIG_FILE}
fi

echo
echo "You are ready to use the APL CLI"
echo "See the helper scripts downloaded to ${SCRIPT_DIR} for examples"
echo "Run apl -h to see command options"
#apl -h

#Try to Clean up
#rm -f ${APL_FILE}

echo
echo "APL CLI Installation Complete, $message"


