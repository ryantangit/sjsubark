#!/bin/bash

echo "SJSUBarker Installation Script"
# Setting up project folder at home directory

INSTALL_DIR=$HOME

if [ ! -d "$INSTALL_DIR/.sjsubarker" ]; then 
	mkdir $INSTALL_DIR/.sjsubarker/
fi

etl_dirs=($INSTALL_DIR/.sjsubarker/etl/ $INSTALL_DIR/.sjsubarker/etl/logs/ $INSTALL_DIR/.sjsubarker/etl/webpages/)
for dir in "${etl_dirs[@]}"; do 
	if [ ! -d $dir ]; 
	then
		if mkdir $dir; then
			echo "Created $dir"
		fi
	else
		echo "$dir already exists"
	fi
done

# Copy current config over
cp ./etl/campus_close.json $INSTALL_DIR/.sjsubarker/etl/campus_close.json

# Create the master.csv for the records
if [ ! -f "$INSTALL_DIR/.sjsubarker/etl/master.csv" ]; 
then
	if touch $INSTALL_DIR/.sjsubarker/etl/master.csv; then
		echo "Created $INSTALL_DIR/.sjsubarker/etl/master.csv"
	fi
else
	echo "$INSTALL_DIR/.sjsubarker/etl/master.csv already exists"
fi

# Compile the newest ETL scrape binary
go build -o "$INSTALL_DIR/.sjsubarker/etl/run_etl" ./etl/main.go
