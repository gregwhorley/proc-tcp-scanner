# Q & A

Questions taken from [here](https://github.com/gravitational/careers/blob/main/challenges/cloud/automation.md)

### How would you prove the code is correct?
Unit and integration testing can prove each function--given a predetermined set of inputs--operates as
expected, but system or end-to-end testing can provide use cases that were never accounted for.
I would build then run the binary on a Linux system and then generate new connections by writing a small
shell script that: runs curl to open new http connections to a site like https://google.com, runs dig for
a small sample of domain names, and other command line operations that will open tcp connections.

### How would you make this solution better?
Adding an interface for the extract, transform, and load phases can allow for processing of any file. Maybe
using a channel to read in and write out data would be a better option than a hard sleep? More testing and
updates to make the new connection reporting and port scanner detection more robust by storing connection
data persistently to statefully report instead of printing out possible duplicates.

### Is it possible for this program to miss a connection?
Yes, a connection could be opened and closed numerous times between the time it takes for the initial read
of the file and the end of the hard sleep timer.

### If you weren't following these requirements, how would you solve the problem of logging every new connection?
Instead of sleeping for 10 seconds between reads, I would try integrating a scanner from a package like bufio with
a channel to increase the response time and miss fewer connections. I would take every unique combo of remote ip:port
and local ip:port and use a hashing algorithm to create a short hash of the unique line then store the hash in a DB.
Hashes that don't match are reported so no duplicates are printed. Same for the port scan.

### Why did you choose x to write the build automation?
I chose make because it was just a simple, quick way to perform the task. I would rely on the DSL for
whatever CI/CD is available for build/test automation and deployment to a production environment. A better
choice for local or dev builds is something like [Skaffold](https://skaffold.dev/) or [Telepresence](https://www.telepresence.io/) that will build and deploy
a local build for the purpose of testing and debugging.

### Is there anything else you would test if you had more time?
More negative test cases for invalid input, more integration tests, and I would build a test harness
and automate an end-to-end test suite to run on a real build of the app in a real Linux environment.

### What is the most important tool, script, or technique you have for solving problems in production? Explain why this tool/script/technique is the most important.
My current set of responsibilities include navigating around Kubernetes and troubleshooting problems
with first party code or its interactions with Kubernetes resources. The most important tool for me
is [Lens](https://k8slens.dev/). It provides manifest data, metrics, logs, and other useful data at
a glance. It allows me to build EKS connection profiles for several AWS accounts so I can switch between
clusters at the click of a button. It's generally a very user friendly tool to troubleshoot issues.

### If you had to deploy this program to hundreds of servers, what would be your preferred method? Why?
I would create a Helm template for a Kubernetes daemonset and then use CI/CD to deploy helm releases
to the clusters. Helm is the de facto package manager for Kubernetes and daemonset is the perfect resource
to ensure that the app will be running on every node. If the task were to deploy to hundreds or bare metal
or VM servers, I would probably use a config management tool like Ansible or Salt. I would need to do
some research on that since I don't really deploy to that type of environment.

### What is the hardest technical problem or outage you've had to solve in your career? Explain what made it so difficult?
The hardest technical problem I've had to solve would be the time where a cloud infrastructure orchestration
service would constantly fail to provision an instance of an on-prem product that my team was tasked with
lifting and shifting into the cloud. We built the orchestration service to put all of the necessary
AWS services into place and completely rollback if any part of the provisioning failed. The service 
inconsistently logged the failures and I had to rely on the end-to-end test suite to uncover the root
cause, which was a AWS limitation for the number of network load balancers that can be created per account.
I then discovered that the design of the orchestration service and how it spins up new infrastructure
would result in a complete inability to provision more than X number of customers. This led to several
tweaks, including a complete change to the Kubernetes ingress controller used for the internal communications
via NLB so that only one AWS resource is needed instead of one NLB per custoemr.
