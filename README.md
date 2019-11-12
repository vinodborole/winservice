#Setup

On your BASE_DIR

git clone https://github.com/vinodborole/winservice.git


set GOPATH = $BASE_DIR/winservice


cd $BASE_DIR/winservice/src/service

dep init

go install

#Usage

usage: service <command>
       where <command> is one of
       install, remove, debug, start, stop, pause or continue.

###Step 1

service install

###Step 2

service start

Check windows service for a service with  name "my service"

In order to start service in debug mode use following cmd

service debug

 