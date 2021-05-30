#!/bin/bash

terraform init
if [ "$?" != "0" ]; then
    # If terraform init fails, don't start the container
    exit
fi

/server