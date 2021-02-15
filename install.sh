#!/bin/bash 

echo " -------- Downloading pwnscan -------- "
wget https://github.com/nolimitcarter/pwnscan/raw/master/bin/64/pwnscan
chmod +x pwnscan
go get github.com/nolimitcarter/pwnscan
go get github.com/fatih/color
go get github.com/likexian/whois-go
echo " -------- pwnscan is now installed -------- "
