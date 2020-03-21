#!/usr/bin/env bash

# Just gets the top level directory of this project. Useful for scripting within the project via relative file paths
EKUBE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

ekube () {
    # if no command given force help page
    local OPTION
	if [[ "$1" != "" ]]; then
        OPTION=$1
    else
        OPTION="help"
    fi
	# handle input options
    case "${OPTION}" in
        'help')
echo "Usage: $ ${FUNCNAME} [option] [flags]
Options:
- help: show this menu
- protoc: Generate grpc files from the proto files
- mock: Mock all services (req: gomock)
- test: run all mock tests
"
        ;;
        'protoc')
            ekubeProtoc
        ;;
        'mock')
            ekubeMockServices
        ;;
        'test')
          ekubeTestServers
        ;;
        *)
            echo -e "ERROR: invalid option. Try..\n$ ${FUNCNAME} help"
        ;;
    esac
}

ekubeTestServers () {
     go test $(go list $EKUBE_DIR/...)
}

ekubeProtoc () {
    packageName="protoc"
    package-installed "${packageName}"

    # Note that in bash, non-zero exit codes are error codes. returning 0 means success
    if [[ "$?" == "0" ]]; then
        # If installed, run protoc
        PROTO_FOLDER="servers"
        SERVER_DIR="${EKUBE_DIR}/${PROTO_FOLDER}"
        # need relative path. cd in subshell to have fine return a path relative to the proto folder
        SERVERS=$(cd "${SERVER_DIR}" && find . -maxdepth 1 -mindepth 1 -type d)
        for SERVER in ${SERVERS}; do
            # for each server found, run proto
            protoc --go_out=plugins=grpc:. "${PROTO_FOLDER}/${SERVER}"/*.proto
        done
    else
        # If protobuf missing, tell them to install it
        echo "missing required package 'protobuf'. Please run the following commands and try again:"
        echo "install protobuf, and then run..."
        echo "$ go get -u github.com/golang/protobuf/protoc-gen-go"
    fi
}

# Generate mock files for all services, putting the results in the proper file. renames some stuff for consistency.
# If you update any services, recommend running this function to update the services for the tests.
ekubeMockServices () {
    packageName="mockgen"
    which "${packageName}"

    # Note that in bash, non-zero exit codes are error codes. returning 0 means success
    if [[ "$?" == "0" ]]; then
      MOCK_FOLDER="services"
      SERVICE_DIR="${EKUBE_DIR}/${MOCK_FOLDER}"
      SERVICES=$(find "${SERVICE_DIR}" -maxdepth 1 -mindepth 1 -type d)
      for SERVICE_PATH in ${SERVICES}
      do
          if [[ -f ${SERVICE_PATH}/interface.go ]]; then
              FOLDER_NAME="${SERVICE_PATH##*/}"
              mockgen \
                  -source=${SERVICE_PATH}/interface.go \
                  -destination=mocks/${MOCK_FOLDER}_mocks/${FOLDER_NAME}_mock.go \
                  -package=${MOCK_FOLDER}_mocks \
                  -mock_names Service=Mock_${FOLDER_NAME}
              fi
      done
    else
        # If mockgen missing, tell them to install it
        echo "missing required package 'mockgen'. Please run the following commands and try again:"
        echo "install protobuf, and then run..."
        echo "$ go get -u github.com/golang/protobuf/protoc-gen-go"
    fi
}

# Check if a command exists in the environment
# Returns 0 if command found
package-installed () {
	result=$(compgen -A function -abck | grep "^$1$")
    # Note that in bash, non-zero exit codes are error codes. returning 0 means success
	if [[ "${result}" == "$1" ]]; then
		# package installed
		return 0
	else
		# package not installed
		return 1
	fi
}
