#!/bin/bash
start=`date +%s`
usage="$(basename "$0") [-h] [-i] [-s str] [-r str] [-n str] [-w int] [-o] [-c str] [-l str] [-a str] [-e]
-- Use appLariat CLI to deploy an application

Where:
    -h         -show this help text
    -i		   -interactive prompts for input
    -s str     -application name (required unless using -i or -e)
    -r str     -release version (required unless using -i or -e)
    -n str     -deployment name (optional, defaults to app_name-<random string>
    -w [1..6]  -workload level (optional, defaults to development)
    -o         -override a component (optional)
    -c str     -component name (required with -o option)
    -a str	   -artifact to use in override (required with -o option)
    -e 		   -use environment variables for input, all other flags ignored
    			variables written in $HOME/.apl/resources/appenv.sh (optional)
    
Pre-requistes:
	apl CLI (v0.2.0+)
	jq  (1.5+)         -run to install: 
"

# Script vars
wl=(1 2 3 4 5 6)
CMD_DIR="/usr/local/bin"
CONFIG_DIR="$HOME/.apl"
CONFIG_FILE="config.toml"
RESOURCE_DIR="$HOME/.apl/resources"
ENV_FILE="appenv.sh"
INTERACTIVE=false
OVERRIDE=false
ENV=false

STACK_NAME=""
RELEASE=""
DEPLOY_NAME=""
WORKLOAD="level1"
COMPONENT_NAME=""
LOC_ARTIFACT_NAME=""
ARTIFACT_NAME=""
DEP_TAG=$(date |md5 | head -c6)

#Check the users environment
command -v apl >/dev/null 2>&1 || { echo >&2 "This script requires apl CLI, install it first;  Aborting."; exit 1; }
command -v jq >/dev/null 2>&1 || { echo >&2 "This script requires jq tool, install it first;  Aborting."; exit 1; }

#Command Line Variables - see usage
while getopts ":s:r:n:w:c:l:a:oieh" opt; do
  case $opt in
    s)
      STACK_NAME=$OPTARG
      ;;
    r)
      RELEASE=$OPTARG
      ;;
    n)
      DEPLOY_NAME=$OPTARG
      ;;
    w)
      if [[ " ${wl[@]} " =~ " ${OPTARG} " ]]; then
        WORKLOAD="level$OPTARG"
      else
        echo "Invalid workload level, using default of 1"
      fi
      ;;
    i)
      echo "Using Interactive"
      INTERACTIVE=true
      ;;
    o)
      echo "Overriding a component"
      OVERRIDE=true
      ;;
    c)
      COMPONENT_NAME=$OPTARG
      ;;
    l)
      LOC_ARTIFACT_NAME=$OPTARG
      ;;
    a)
      ARTIFACT_NAME=$OPTARG
      ;;
    e)
      echo "Setting deployment fields from environment variables"
      ENV=true
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

#First Collect Info
if [[ $ENV == true ]]; then
   if [ -f ${RESOURCE_DIR}/${ENV_FILE} ]; then
       #make sure the env vars are set
       . ${RESOURCE_DIR}/${ENV_FILE}
   else
       echo "No environment file, trying to see if variables are already set"
   fi
   
   set -e
   STACK_NAME=${APL_STACK_NAME:?Missing required env var}
   RELEASE=${APL_RELEASE_VER:?Missing required env var}
   LOC_DEPLOY_ID=${APL_LOC_DEPLOY_ID:?Missing required env var}
   RELEASE_ID=${APL_RELEASE_ID:?Missing required env var}
   STACK_ID=${APL_STACK_ID:?Missing required env var}
   WORKLOAD=${APL_WORKLOAD:-1}
   if [[ $OVERRIDE == true ]]; then
       LOC_ARTIFACT_ID=${APL_LOC_ARTIFACT_ID:?Missing required env var}
   	   STACK_COMPONENT_ID=${APL_STACK_COMPONENT_ID:?Missing required env var}
   	   COMPONENT_NAME=${APL_STACK_COMPONENT_NAME:?Missing required env var}
   	   COMP_SERVICE_NAME=${APL_STACK_COMPONENT_NAME:?Missing required env var}
   	   ARTIFACT_NAME=${APL_ARTIFACT_NAME:?Missing required env var}
   	   ARTIFACT_TYPE=${APL_ARTIFACT_TYPE:?Missing required env var}
   fi
   set +e
else
  echo "Checking connection to applariat"
  #qstart=`date +%s`
  LOC_DEPLOY_LIST=$(apl loc-deploys -o json)
  if [[ $(echo $LOC_DEPLOY_LIST | jq -r '.[] | has("name")') != true  ]]; then
	echo "There was a problem connecting to appLariat"
    exit 1
  fi
  
  #STACK_LIST=$(apl stacks -o json)
  #RELEASE_LIST=$(apl releases -o json)
  #if [[ $OVERRIDE -eq 1 ]]; then
    #LOC_ARTIFACT_LIST=$(apl loc-artifacts -o json)
  #fi
  #qend=`date +%s`
  #runtime=$((qend-qstart))
  #echo "apl query took $runtime sec"
fi

#set Deploy location id
if [[ $ENV == false ]] || [ -z $LOC_DEPLOY_ID ]; then
  LOC_DEPLOY_ID=$(echo $LOC_DEPLOY_LIST | jq -r '.[0].id')
  echo "Using Deployment Location $(echo $LOC_DEPLOY_LIST | jq -r '.[].name')"   
fi

#Run Interactive Mode
######################
if [[ $INTERACTIVE == true ]]; then
  echo "Starting interactive mode"
  echo "Looking up available applications"
  STACK_LIST=$(apl stacks -o json)
  slist=( `echo ${STACK_LIST} | jq -rc '.[].name'` )
  echo
  echo "Available apps: ${slist[@]}"
  echo
  read -p "Enter the app to deploy: " app
  STACK_NAME=$app
  STACK_ID=$(echo ${STACK_LIST} | jq -rc --arg sname $app '.[] | select(.name == $sname) | .id')
  #echo $STACK_ID
  echo "Looking up releases for ${STACK_NAME}"
  RELEASE_LIST=$(apl releases --stack-id $STACK_ID -o json)
  rlist=( `echo ${RELEASE_LIST} | jq -r '.[] | .version'` )
  echo
  echo "Available release versions: ${rlist[@]}"
  echo
  read -p "Enter the release version to deploy: " ver
  RELEASE=$ver
  RELEASE_REC=$(echo ${RELEASE_LIST} | jq -rc --argjson rel $RELEASE '.[] | select( .version == $rel)')
  RELEASE_ID=$(echo ${RELEASE_REC} | jq -r '.id')
  echo #$RELEASE_ID
  read -p "Enter the name of your deploy [$STACK_NAME-$RELEASE-$DEP_TAG]: " dname
  DEPLOY_NAME=${dname:-${STACK_NAME}-${RELEASE}-${DEP_TAG}}
  echo
  read -p "Enter a value from 1(Development) - 6(Production) for the workload level [1]: " wl
  WORKLOAD="level${wl:-1}"
  if [[ $OVERRIDE == true ]]; then
    clist=( `echo ${RELEASE_REC} | jq -r '.components[].name'` )
    echo
    echo "Available components: ${clist[@]}"
    echo
    read -p "Type in the component to override: " comp
    COMPONENT_NAME=$comp
    #cobj=$(echo ${RELEASE_LIST} | \
    #jq -rc --arg rid $RELEASE_ID --arg cn $COMPONENT_NAME '.[] | select(.id == $rid) | .components[] | select(.name == $cn)')
    #lalist=( `echo ${LOC_ARTIFACT_LIST} | jq -rc '.[].name'` )
    sa_id=$(echo ${RELEASE_REC} | jq -r --arg cn $COMPONENT_NAME '.components[] | 
    select(.name == $cn) | .services[0].build.artifacts |  if has("code") then .code elif has("builder") then .builder else .image end')
    #echo $sa_id
    SA_REC=$(apl stack-artifacts get $sa_id -o json)
    #echo
  	#echo "Available Artifact Locations: ${lalist[@]}"
  	#echo
    #read -p "Type in the artifact location for the new artifact: " la
    #LOC_ARTIFACT_NAME=$la
    LOC_ARTIFACT_ID=$(echo ${SA_REC} | jq -r '.loc_artifact_id')
    echo
  	echo "Current Stack Artifact: $(echo ${SA_REC} | jq -r '.artifact_name' )"
    echo
    read -p "Enter the artifact name for the new artifact: " aname
    ARTIFACT_NAME=$aname
  fi
  
  echo "Values Entered: "
  echo "	Application: $STACK_NAME"
  echo "	Release Version: $RELEASE"
  echo "	Deployment Name: $DEPLOY_NAME"
  echo "	Workload Type: $WORKLOAD"
  if [[ $OVERRIDE == true ]]; then
    echo "	Component Name: $COMPONENT_NAME"
    #echo "	Artifact Location Name: $LOC_ARTIFACT_NAME"
    echo "	Artifact Name: $ARTIFACT_NAME"
  fi
fi
################
#End of Interactive

#Check for App and Release
if [ -z $STACK_NAME ] || [ -z $RELEASE ]; then
	echo "Stack Name and release version required, use -s <name> -r <int> or -i for interactive, exiting"
	echo
	echo "Usage: $usage" >&2
	exit 1
fi

#Non-interactive query based on passed in information
if [ -z $STACK_ID ]; then
	STACK_REC=$(apl stacks --name $STACK_NAME -o json)
	STACK_ID=$(echo ${STACK_REC} | jq -rc '.[0].id')
	#echo $STACK_ID
fi
#Time to pull together all of the information
#Get Release Info
if [ -z $RELEASE_ID ]; then
    RELEASE_REC=$(apl releases --stack-id $STACK_ID -o json | jq -c --argjson rel $RELEASE '.[] | select(.version == $rel)')
    RELEASE_ID=$(echo ${RELEASE_REC} | jq -r '.id')
    #echo $RELEASE_ID
fi

if [[ $ENV == false ]] && [[ $OVERRIDE == true ]]; then
    #Get stack component id
    STACK_COMPONENT_ID=$(echo ${RELEASE_REC} | 
      jq -rc --arg cname $COMPONENT_NAME '.components[] | select(.name == $cname) | .stack_component_id')
    COMP_SERVICE_NAME=$(echo ${RELEASE_REC} | 
      jq -rc --arg cname $COMPONENT_NAME '.components[] | select(.name == $cname) | .services[0].name')
    #echo $STACK_COMPONENT_ID
    #Get the artifact type for the component
    ARTIFACT_TYPE=$(echo ${RELEASE_REC} | \
      jq -r --arg cn $COMPONENT_NAME '.components[] | select(.name == $cn) | .services[0].build.artifacts |  
      if has("code") then "code" elif has("builder") then "builder" else "image" end')
    #if [ -z LOC_ARTIFACT_ID ]; then
		sa_id=$(echo ${RELEASE_REC} | \
		  jq -r --arg cn $COMPONENT_NAME '.components[] | select(.name == $cn) | .services[0].build.artifacts |  
		  if has("code") then .code elif has("builder") then .builder else .image end')
		LOC_ARTIFACT_ID=$(apl stack-artifacts get $sa_id -o json | jq -r '.loc_artifact_id')
		#echo $LOC_ARTIFACT_ID
	#fi
fi
   
#Verify and set deployment name
if [ -z $DEPLOY_NAME ]; then
    DEPLOY_NAME=$STACK_NAME-$RELEASE-$DEP_TAG
    echo "No Deployment name provided using default - $DEPLOY_NAME"
else
    if [[ $(apl deployments --name $DEPLOY_NAME -o json | jq '. | length') != 0 ]]; then
       echo "Deployment with that name already exists, exiting"
       exit 1
    fi
    echo
    echo "Using deployment name - $DEPLOY_NAME"
fi

#Create a environment file for later use
if [[ $ENV == false ]]; then
	echo
	echo "Adding ids to an environment file ${RESOURCE_DIR}/${ENV_FILE}, run this script with the -e option to recreate this deployment"
	echo
	if [ ! -d ${RESOURCE_DIR} ]; then
		mkdir -p ${RESOURCE_DIR}
	fi

	cat >${RESOURCE_DIR}/${ENV_FILE} <<EOL
export APL_STACK_NAME=$STACK_NAME
export APL_RELEASE_VER=$RELEASE
export APL_LOC_DEPLOY_ID=$LOC_DEPLOY_ID
export APL_LOC_ARTIFACT_ID=$LOC_ARTIFACT_ID
export APL_STACK_ID=$STACK_ID
export APL_RELEASE_ID=$RELEASE_ID
export APL_STACK_COMPONENT_ID=$STACK_COMPONENT_ID
export APL_STACK_COMPONENT_NAME=$COMPONENT_NAME
export APL_COMPONENT_SERVICE_NAME=$COMP_SERVICE_NAME
export APL_ARTIFACT_NAME=$ARTIFACT_NAME
export APL_ARTIFACT_TYPE=$ARTIFACT_TYPE
export APL_WORKLOAD=$WORKLOAD
EOL
fi

#Submit the deployment
if [[ $OVERRIDE == false ]]; then
    DEPLOY_COMMAND="apl deployments create --loc-deploy-id $LOC_DEPLOY_ID --release-id $RELEASE_ID --name $DEPLOY_NAME --workload-type $WORKLOAD -o json"
else
    #A little more work to do with the override
    #First register the new artifact
    SA_CMD="apl stack-artifacts create --stack-id $STACK_ID --loc-artifact-id ${LOC_ARTIFACT_ID} --name ${ARTIFACT_NAME} --stack-artifact-type ${ARTIFACT_TYPE} --artifact-name ${ARTIFACT_NAME} -o json"
    #echo "Adding Artifact $SA_CMD"
    SA_CREATE=$(${SA_CMD})

    if [[ $(echo $SA_CREATE | jq -r '. | has("message")') == "true"  ]]; then
        echo $SA_CREATE | jq -r '.message'
        exit 1
    elif [[ $(echo $SA_CREATE | jq -r '. | has("data")') == "true" ]]; then
        STACK_ARTIFACT_ID=$(echo $SA_CREATE | jq -r '.data')
        #echo $STACK_ARTIFACT_ID
    else
        echo "ERROR: ${SA_CREATE}"
        exit 1
    fi
    
    #Now we are going to create a yaml file for the deployment
    cat >deploy.yaml <<EOL
name: ${DEPLOY_NAME}
release_id: ${RELEASE_ID}
loc_deploy_id: ${LOC_DEPLOY_ID}
workload_type: ${WORKLOAD}
components:
- stack_component_id: ${STACK_COMPONENT_ID}
  services:
  - name: $COMP_SERVICE_NAME
    overrides:
      build:
        artifacts:
          ${ARTIFACT_TYPE}: ${STACK_ARTIFACT_ID}
EOL

    DEPLOY_COMMAND="apl deployments create -f deploy.yaml -o json"  
fi
echo "Submitting the deployment:"
echo "$DEPLOY_COMMAND"

APL_DEPLOY_CREATE=$(${DEPLOY_COMMAND})
echo
if [[ $(echo $APL_DEPLOY_CREATE | jq -r '. | has("message")') == "true" ]]; then
     echo $APL_DEPLOY_CREATE | jq -r '.message'
     exit 1
elif [[ $(echo $APL_DEPLOY_CREATE | jq -r '. | has("data")') == "true" ]]; then
    APL_DEPLOYMENT_ID=$(echo $APL_DEPLOY_CREATE | jq -r '.data.deployment_id')
else
   echo "ERROR: $APL_DEPLOY_CREATE"
   exit 1
fi
#echo $APL_DEPLOYMENT_ID

if [ -z $APL_DEPLOYMENT_ID ]; then
  echo "Failed to get deployment id, you can try apl deployments command to return a list of deployments"
  exit
else
  echo "Deployment ID: $APL_DEPLOYMENT_ID"
  echo "Waiting for the deployment to complete"
  #state=$(apl deployments get $APL_DEPLOYMENT_ID -o json | jq -r '.status.state')
  while [[ $(apl deployments get $APL_DEPLOYMENT_ID -o json | jq -r '.status.state') =~ ^(queued|deploying|pending)$ ]]; do
      echo "App is deploying"
      sleep 15
  done
  echo "Deployment completed with the following info:"
  echo "Details:"
  echo
  apl deployments get $APL_DEPLOYMENT_ID -o json | 
    jq '.status | { name: .namespace, state: .state, description: .description, services: .components[].services[]}'
fi

end=`date +%s`
runtime=$((end-start))
echo
echo "APL Deployed Successfully"