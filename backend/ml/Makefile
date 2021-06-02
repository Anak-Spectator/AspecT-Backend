build:
	sudo docker build -t aspect_ml .

run:
	sudo docker run -d -it -p 7000:7000 --name aspect_ml_apps aspect_ml

stop:
	sudo docker container stop aspect_ml_apps

remove:
	sudo docker image rm -f aspect_ml

