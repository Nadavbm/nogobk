rock-my-dudale:
	docker-compose -f docker-compose.yml down && \
	docker-compose -f docker-compose.yml up -d database && \
	sleep 5 && \
	echo "i think we are ready dude..." && \
	docker-compose -f docker-compose.yml up -d nogobk
