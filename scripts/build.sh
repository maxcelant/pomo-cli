#!/bin/bash
go build -o pomo ./cmd/pomo
sudo mv pomo /usr/local/bin/
mkdir -p ~/.pomo
cp pomo.yaml ~/.pomo/
echo "Pomo cli added to local bin directory!"
