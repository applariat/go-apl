#!/bin/bash
start=`date +%s`
usage="$(basename "$0") [-h] [-i] [-n str] [-c str] [-a str]
-- Use appLariat CLI to override an artifact on a running deployment

Where:
    -h         -show this help text
    -i		   -interactive prompts for input
    -n str     -deployment name (required unless -i)
    -c str     -component name (required unless -i)
    -a str	   -artifact to use in override (required unless -i)
    
Pre-requistes:
	apl CLI (v0.2.0+)
	jq  (1.5+)         -run to install: 
"

# Script vars
CMD_DIR="/usr/local/bin"
CONFIG_DIR="$HOME/.apl"
CONFIG_FILE="config.toml"
RESOURCE_DIR="$HOME/.apl/resources"
LOC_DEPLOY_FILE="loc_deploy.json"
LOC_ART_FILE="loc_art.json"
STACKS_FILE="stacks.json"
ENV_FILE="appenv.sh"
INTERACTIVE=false

DEPLOYMENT_NAME=""
COMPONENT_NAME=""
ARTIFACT_NAME=""

#Check the users environment
command -v apl >/dev/null 2>&1 || { echo >&2 "This script requires apl CLI, install it first;  Aborting."; exit 1; }
command -v jq >/dev/null 2>&1 || { echo >&2 "This script requires jq tool, install it first;  Aborting."; exit 1; }

#Command Line Variables - see usage
while getopts ":n:c:a:ih" opt; do
  case $opt in
    i)
      INTERACTIVE=true
      ;;
    n)
      DEPLOYMENT_NAME=$OPTARG
      ;;
    c)
      COMPONENT_NAME=$OPTARG
      ;;
    a)
      ARTIFACT_NAME=$OPTARG
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
#echo "Checking connection to applariat"

#Run Interactive Mode
######################
if [[ $INTERACTIVE == true ]]; then
	echo "Starting interactive mode"
	echo "Looking up available deployments"
	DEPLOY_LIST=$(apl deployments -o json)
	if [[ $(echo $DEPLOY_LIST | jq -r '.[0] | has("name")') != true  ]]; then
		echo "There was a problem connecting to appLariat"
		exit 1
	fi
	dlist=( `echo ${DEPLOY_LIST} | jq -rc '.[].name'` )
	echo
	echo "Available apps: ${dlist[@]}"
	echo
	read -p "Enter the deployment to update: " deploy
	DEPLOYMENT_NAME=$deploy
	DEPLOY_REC=$(echo ${DEPLOY_LIST} | jq -r --arg dname $DEPLOYMENT_NAME '.[] | select(.name == $dname)')
	DEPLOYMENT_ID=$(echo ${DEPLOY_REC} | jq -r '.id')
	#echo $DEPLOYMENT_ID
	echo "Looking up components for ${DEPLOYMENT_NAME}"
	clist=( `echo ${DEPLOY_REC} | jq -r '.release.components[].name'` )
	echo
	echo "Available components: ${clist[@]}"
	echo
	read -p "Type in the component to override: " comp
	COMPONENT_NAME=$comp
	sa_id=$(echo ${DEPLOY_REC} | jq -r --arg cn $COMPONENT_NAME '.release.components[] | 
	select(.name == $cn) | .services[0].build.artifacts |  if has("code") then .code elif has("builder") then .builder else .image end')
	SA_REC=$(apl stack-artifacts get $sa_id -o json)
	LOC_ARTIFACT_ID=$(echo ${SA_REC} | jq -r '.loc_artifact_id')
	echo
	echo "Current Stack Artifact: $(echo ${SA_REC} | jq -r '.artifact_name' )"
	echo
	read -p "Enter the artifact name for the new artifact: " aname
	ARTIFACT_NAME=$aname
  
	echo "Values Entered: "
	echo "	Deployment Name: $DEPLOYMENT_NAME"
	echo "	Component Name: $COMPONENT_NAME"
	echo "	Artifact Name: $ARTIFACT_NAME"
fi
################
#End of Interactive

#Check for Required Fields
if [ -z $DEPLOYMENT_NAME ] || [ -z $COMPONENT_NAME ] || [ -z $ARTIFACT_NAME ]; then
	echo "Deployment Name and component name and artifact name required, 
	use -n <name> -c <component name> -a <artifact name> or -i for interactive, exiting"
	echo
	echo "Usage: $usage" >&2
	exit 1
fi


#Non-interactive query based on passed in information
if [[ $INTERACTIVE == false ]]; then
	DEPLOY_LIST=$(apl deployments --name $DEPLOYMENT_NAME -o json)
	DEPLOY_REC=$(echo ${DEPLOY_LIST} | jq -r '.[]')
	DEPLOYMENT_ID=$(echo ${DEPLOY_REC} | jq -rc '.id')
fi

#Get stack component id
STACK_ID=$(echo ${DEPLOY_REC} | jq -r '.stack.id')
#echo $STACK_ID
COMP_REC=$(echo ${DEPLOY_REC} | \
  jq -c --arg cname $COMPONENT_NAME '.release.components[] | select(.name == $cname)')
STACK_COMPONENT_ID=$(echo ${COMP_REC} | jq -r '.stack_component_id')
COMP_SERVICE_NAME=$(echo ${COMP_REC} | jq -r '.services[0].name')
#echo $STACK_COMPONENT_ID
#Get the artifact type for the component
ARTIFACT_TYPE=$(echo ${COMP_REC} | jq -r '.services[0].build.artifacts |  
  if has("code") then "code" elif has("builder") then "builder" else "image" end')
#if [ -z LOC_ARTIFACT_ID ]; then
	sa_id=$(echo ${COMP_REC} | jq -r '.services[0].build.artifacts |  
	  if has("code") then .code elif has("builder") then .builder else .image end')
	LOC_ARTIFACT_ID=$(apl stack-artifacts get $sa_id -o json | jq -r '.loc_artifact_id')
	#echo $LOC_ARTIFACT_ID
#fi
   

#Submit the override
#First register the new artifact
SA_CMD="apl stack-artifacts create --stack-id $STACK_ID --loc-artifact-id ${LOC_ARTIFACT_ID} --name ${ARTIFACT_NAME} --stack-artifact-type ${ARTIFACT_TYPE} --artifact-name ${ARTIFACT_NAME} -o json"
echo "Adding Artifact"
SA_CREATE=$(${SA_CMD})

if [[ $? -ne 0 ]]; then
	echo $SA_CREATE
	exit 1
elif [[ $(echo $SA_CREATE | jq -r '. | has("data")') == "true" ]]; then
	STACK_ARTIFACT_ID=$(echo $SA_CREATE | jq -r '.data')
	#echo $STACK_ARTIFACT_ID
else
	echo "ERROR: ${SA_CREATE}"
	exit 1
fi

#Error in CLI for this structure
#Construct the component string for the override
#artifacts=$(echo ${COMP_REC} | jq -c '.services[0].build.artifacts')
#comp=(StackComponentID=${STACK_COMPONENT_ID})
#comp+=(ServiceName=${COMP_SERVICE_NAME})
#comp+=(StackArtifactID=${STACK_ARTIFACT_ID})
#artifacts=$(echo ${artifacts} | jq -c --arg art ${ARTIFACT_TYPE} 'with_entries(select(.key != $art))')
#comp+=( `echo ${artifacts} | jq -r 'map_values("StackArtifactID=" + .) |to_entries|.[].value'` )

#comp=$(IFS=, ; echo "${comp[*]}")

#using file based override
cat >update.yaml <<EOL
command: override
components:
- stack_component_id: ${STACK_COMPONENT_ID}
  services:
  - name: ${COMP_SERVICE_NAME}
    build:
      artifacts:
        code: ${STACK_ARTIFACT_ID}
EOL


OVERRIDE_COMMAND="apl deployments override $DEPLOYMENT_ID -f update.yaml -o json"  

echo "Submitting the override:"
echo "$OVERRIDE_COMMAND"

APL_DEPLOY_OVERRIDE=$(${OVERRIDE_COMMAND})
echo
if [[ $(echo $APL_DEPLOY_OVERRIDE | jq -r '. | has("message")') == "true" ]]; then
     echo $APL_DEPLOY_OVERRIDE | jq -r '.message'
elif [[ $(echo $APL_DEPLOY_OVERRIDE | jq -r '. | has("unchanged")') == "true" ]]; then
     echo "Component Update Pending"
else
   echo "RESULT: $APL_DEPLOY_OVERRIDE"
   exit 1
fi
#echo $DEPLOYMENT_ID

echo "Deployment ID: $DEPLOYMENT_ID"
echo "Waiting for the override to complete"
sleep 30
#state=$(apl deployments get $APL_DEPLOYMENT_ID -o json | jq -r '.status.state')
while [[ $(apl deployments get $DEPLOYMENT_ID -o json | jq -r --arg cn $COMPONENT_NAME '.status.components[] | select(.name == $cn) | .state') != "running" ]]; do
  echo "Component is updating"
  sleep 10
done
echo "Deployment override completed with the following info:"
echo
apl deployments get $DEPLOYMENT_ID -o json | \
jq '.status | { name: .namespace, state: .state, description: .description, services: .components[].services[]}'

end=`date +%s`
runtime=$((end-start))
echo
echo "APL Override Successfully"