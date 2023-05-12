#!/bin/bash

function migrate() {
	cd ./superengine
	sea-orm-cli generate entity -o model/src --lib
}

function proto() {
	cd ./superengine
	cargo run --package control -- build proto
	cargo run --package control -- build enums
}

function swagger() {
	cd ./superserver
	rm -f ../protocol/api/swagger.yaml
	swag init --output ../protocol/api --outputTypes yaml --parseInternal
}

case "$1" in
    migrate)
        migrate
        ;;
	proto)
		proto
        ;;
	swagger)
		swagger
		;;
    *)
    info "Usage: $0 {migrate|proto|swagger}"
esac
