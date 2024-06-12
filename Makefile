SHELL:=/bin/bash

docker-network: 
	docker network create checkout


up: 
	docker-compose up -d --build