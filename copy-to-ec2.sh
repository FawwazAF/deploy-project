#!/bin/bash
scp -r -i D:/Documents/Alterra/AWS/aws_ec2_sing.pem C:/Users/fawwa/go/deployment-test/project-deploy copy/program/* ubuntu@13.212.54.221:/home/ubuntu/app
# scp -r -i D:/Documents/Alterra/AWS/aws_ec2_sing.pem C:/Users/fawwa/projectdump.sql ubuntu@52.77.223.67:/home/ubuntu
