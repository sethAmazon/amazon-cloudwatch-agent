#!/bin/sh

while sudo fuser /var/{lib/{dpkg,apt/lists},cache/apt/archives}/lock >/dev/null 2>&1; do
   echo "Waiting for lock"
   sleep 1
done
echo "Lock is empty"
sudo dpkg -i -E ./amazon-cloudwatch-agent.deb