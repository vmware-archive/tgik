#!/bin/bash

kind get nodes --name=ipvs | xargs -n1 -I {} docker exec {} sysctl -w net.ipv4.conf.all.rp_filter=0
