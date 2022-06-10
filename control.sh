#!/bin/bash

SERVICE_NAME="toolbox"
VERSION="v1.0.0"

function build() {
	go build -o $SERVICE_NAME -tags=jsoniter -ldflags '-w -s -X main.tags=$SERVICE_NAME-$VERSION'
}

function swagger() {
	cd $PWD
	$GOPATH/bin/swag init
}

function status() {
	auth_pid=$(ps -ef |grep $SERVICE_NAME|grep -v grep|awk '{print $2}')
	if [[ x$auth_pid == x"" ]]; then
		echo "$SERVICE_NAME service is not running."
		return 0
	else
		echo "$SERVICE_NAME service is running as ${auth_pid}"
	fi
}

function start() {
	auth_pid=$(ps -ef |grep $SERVICE_NAME |grep -v grep|awk '{print $2}')
	if [[ x$auth_pid == x"" ]]; then
		echo "starting auth process..."
		nohup ./$SERVICE_NAME >> runtime.log 2>&1 &
	fi
	status
}

function restart() {
	#如果auth进程存在，直接kill服务进程，如果不存在，启动auth
	auth_pid=$(ps -ef |grep $SERVICE_NAME|grep -v grep|awk '{print $2}')
	if [[ x$auth_pid != x"" ]]; then
			echo "$SERVICE_NAME service is running as ${auth_pid}"
			echo "killing ${auth_pid}"
			kill -15 ${auth_pid}
			echo "wait server exit..."
			auth_pid=$(ps -ef |grep $SERVICE_NAME|grep -v grep|awk '{print $2}')
			if [[ x$auth_pid != x"" ]]; then
			  kill -9 ${auth_pid}
			fi
	fi
	start
	echo "wait server start..."
	status
}

function stop() {
	auth_pid=$(ps -ef |grep $SERVICE_NAME|grep -v grep|awk '{print $2}')
	if [[ x$auth_pid == x"" ]]; then
		echo "$SERVICE_NAME service is not running, skipping it."
		return 0
	else
		echo "$SERVICE_NAME service is running as ${auth_pid}"
		echo "killing service #${auth_pid}"
		kill ${auth_pid}
		return 1
	fi
}

case "$1" in
    build)
        build
        ;;
	swagger)
		swagger
		;;
    start)
        start
        ;;
    stop)
        stop
        ;;
    restart)
        restart
	      ;;
    status)
        status
        ;;
    *)
    echo "Usage: $0 {start|stop|restart|status}"
esac
