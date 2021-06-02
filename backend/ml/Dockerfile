FROM python:3.8-slim

# set the working directory in the container
WORKDIR /ml

# copy the dependencies file to the working directory
COPY . .

RUN ls 
# install dependencies
RUN pip3 install -r requirements.txt

ENV ENV "prod"
EXPOSE 7000

# command to run on container starti
CMD [ "python", "./app.py" ]