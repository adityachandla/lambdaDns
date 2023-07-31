# DNS simulation using AWS lambda
This repository contains the backend code that simulates how a DNS server works. The project can be found [here](https://adityachandla.github.io/dnsProject).

## Code overview
Each DNS node is simulated using a separate lambda function. The `cmd` directory contains a separate folder for each lambda. utils package contains some common
utilities that are used by all the lambda functions.
