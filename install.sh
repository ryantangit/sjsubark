#!/bin/bash

echo "SJSUBarker Installation"
# Setting up project folder at home directory
mkdir $HOME/.sjsubarker/
mkdir -p $HOME/.sjsubarker/etl/ $HOME/.sjsubarker/etl/logs/ $HOME/.sjsubarker/etl/webpages/

# Copy current config over
cp ./etl/campus_close.json $HOME/.sjsubarker/etl/campus_close.json

