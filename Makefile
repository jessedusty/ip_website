
all:
	docker build -t jessedusty/ip_updater .
	docker push jessedusty/ip_updater
