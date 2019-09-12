build:
	GOOS=linux go build -o app
	docker build -t gcr.io/river-wave-252713/say .
	rm -f app
	# GOOS=linux go build -o app
	# docker build -t grpc-speech .
	# rm -f app

push:
	 docker push gcr.io/river-wave-252713/say