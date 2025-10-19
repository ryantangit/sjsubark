#!/bin/bash
echo "SJSUBarker Installation Script"
# Setting up project folder at home directory

INSTALL_DIR=$HOME

if [ ! -d "$INSTALL_DIR/.sjsubark" ]; then 
	mkdir $INSTALL_DIR/.sjsubark/
fi

etl_dirs=($INSTALL_DIR/.sjsubark/etl/ $INSTALL_DIR/.sjsubark/etl/logs/ $INSTALL_DIR/.sjsubark/etl/webpages/)
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
cp ./etl/campus_close.json $INSTALL_DIR/.sjsubark/etl/campus_close.json

# Create the master.csv for the records
if [ ! -f "$INSTALL_DIR/.sjsubark/etl/master.csv" ]; 
then
	if touch $INSTALL_DIR/.sjsubark/etl/master.csv; then
		echo "Created $INSTALL_DIR/.sjsubark/etl/master.csv"
	fi
else
	echo "$INSTALL_DIR/.sjsubark/etl/master.csv already exists"
fi

# Compile the newest ETL scrape binary
if go build -o "$INSTALL_DIR/.sjsubark/etl/run_etl" ./etl/main.go; then
	echo "Compiled newest ETL scrape binary successfully"
fi

