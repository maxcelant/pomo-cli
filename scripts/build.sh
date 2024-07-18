#!/bin/bash
go build -o pomo ./cmd/pomo
sudo mv pomo /usr/local/bin/
echo "Pomo cli added to local bin directory!"
