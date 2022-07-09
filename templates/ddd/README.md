# Domain Driven Design Template
Template for a Web API Project with structure that implements Domain Driven
Design ( PSA: I'm still exploring and learning DDD and hence this strucutre may
evolve as my understanding of DDD does)

This document goes over and explains the folder structure

## Project Structure
### Root
On the root of the project you can see three folders, cmd, pkg and docs, as 
you will see, the "internal" package (which is common in most go projects) has 
been omitted to simplify the structure and avoid having to make the decision
of placing code between the two packages (internal or pkg) while still having
to worry about it sticking to DDD's philosophy

#### cmd
TL;DR this package contains the entry point (main.go) and nothing much else

cmd is a very small package that usually would just contain the entry point
for the application, multiple entrypoints can be created here to have different
version of the applications (for example one application using real data while
the other using sample data)

Each entry point are to be in it's own folder within cmd

#### docs
This is an auto generated package created by the swaggo/swag package to serve
the service's swagger documentation

Follow their documentation on [Github](https://github.com/swaggo/swag) to
install the cli and regenerate the swagger document with the following command
from the root directory of the project
```
swag init -g cmd/server/main.go
```
Note the location of main.go if a custom location has been created

The GitHub has further documentation on how the cli uses comments in the code
(placed above the controllers generally) to generate the swagger file

#### pkg
This is the package that contains most of the application, you will see that
the package has three subpackages(folders? not sure what these are called)
namely app, domain and infra 

##### domain
This package contains all the business logic, validations and entities for all
domains in the application. This package *DOES NOT* contain any knowledge of
the actual data sources (Databases or external services reachable via REST APIs
, GRPCS etc for example) or how the data is persisted (Generally speaking
Databases, but could be APIs) but simply invoke them to get/save/modify data
