#! /bin/bash

trap "exit 1" TERM
export TOP_PID=$$

NC='\033[0m'
RED='\033[0;31m'
GREEN='\033[0;32m'

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

copmose_dir=$DIR/../../../docker-compose.test.yml
reload_copmose_each_test=false
tests_list=( test_ping.py test_auth.py )

for i in "$@" ; do
    case $i in
        -r=*|--reload_copmose_each_test=*)
            reload_copmose_each_test="${i#*=}"
            shift
        ;;
        -d=*|--docker-compose=*)
            copmose_dir="${i#*=}"
            shift
        ;;
        -t=*|--tests=*)
            tests_list="${i#*=}"
            shift
        ;;
        *)
            # unknown option
        ;;
    esac
done

wait_healty() {
    echo "Waiting for server to be healthy"
    container_id=$(docker ps -qf "name=$1")
    while true; do
        if [ $( docker container inspect -f '{{.State.Running}}' $container_id ) = true ]; then
            echo "Server is healthy"
            break
        fi
        sleep 1
    done
    sleep 5
}

run_docker_compose() {
    echo "Starting server"

    if $(docker-compose -f $copmose_dir up --build -d); then
        wait_healty "server"
        echo "Server started"
    else
        echo "Server failed to start"
        kill -s TERM $TOP_PID
    fi
}

stop_docker_compose() {
    echo "Stopping server"

    if $( docker-compose -f $copmose_dir down ); then
        echo "Server stopped"
    else
        echo "Server failed to stop"
        kill -s TERM $TOP_PID
    fi
}

# Run the server
if [ ! $reload_copmose_each_test = true ]; then
    echo "$( run_docker_compose )"
fi

passed=true

echo "Running tests: [${tests_list[@]}]"

# Run the tests
for test in "${tests_list[@]}"; do
    echo "Running test $test"

    if [ $reload_copmose_each_test = true ]; then
        echo "$( run_docker_compose )"
    fi

    cmd=$( python3 -m pytest $DIR/$test --color=yes --tb=short )
    if [ $? -eq 0 ]; then
        echo "${GREEN}Test $test passed${NC}"
    else
        echo "${RED}Test $test failed${NC}"
        echo "$cmd"
        passed=false
        break
    fi

    if [ $reload_copmose_each_test = true ]; then
        echo "$( stop_docker_compose )"
    fi

done

# Stop the server
if [[ ! $reload_copmose_each_test = true || $passed = false ]]; then
    echo "$( stop_docker_compose )"
fi

if [ $passed = true ]; then
    echo "All tests passed"
    exit 0
else
    echo "Some tests failed"
    exit 1
fi
