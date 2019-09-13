#!/bin/bash

kind get nodes --name=iptables | xargs -n1 -I {} docker exec {} sysctl -w net.ipv4.conf.all.rp_filter=0
