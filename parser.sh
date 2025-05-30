#!/bin/bash

read -p "Enter host url: " host

host=${host}

config=$(cd parser && echo "$host" | go run ./main.go)

echo "$config" | sudo tee ./config.json >/dev/null