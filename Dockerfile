FROM ubuntu:22.04

WORKDIR /usr/src/app

RUN apt update
RUN apt install -y curl 
RUN apt upgrade -y
RUN apt install npm -y
RUN npm install -g n
RUN n lts
RUN npm install -g typescript

COPY ./ .

RUN npm install
RUN tsc
EXPOSE 3000
CMD ["node", "dist/app.js"]