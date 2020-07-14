#!/bin/bash
function start(){
    chmod +x ./businessCode ./static/http
    ./businessCode &
    cd static
    ./http &
}

function stop(){
	ps -ef|grep "businessCode"|grep -v grep |awk -F" " '{print $2}'|xargs -i kill {};
}

case $1 in
    start)
        start
        ;;
    stop)
        stop
        ;;
    *)
        echo "use start|stop"
        ;;
esac
