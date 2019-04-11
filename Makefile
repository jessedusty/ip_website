build:
	docker build -t jessedusty/ip_updater .

run:
	docker run --rm -it -p 80:80 jessedusty/ip_updater

push: build
	docker push jessedusty/ip_updater
