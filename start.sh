#!/bin/bash 

goapp serve -host 0.0.0.0  -port 8080 -admin_port 3000 dispatch.yaml client/app.yaml server/app.yaml