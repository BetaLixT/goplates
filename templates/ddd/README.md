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

Follow their documentation on [GitHub](https://github.com/swaggo/swag) to
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
domains in the application. This package **DOES NOT** contain any knowledge of
the actual data sources (Databases or external services reachable via REST APIs
, GRPCS etc for example) or how the data is persisted (Generally speaking
Databases, but could be APIs) but simply invoke them to get/save/modify data

Each domain in the application will have it's own subpackage within the domain
package and one domain should not import another domain, in the template
"forecast" has been provided as an example domain.

Individual domain packages will typically contain a service structure that 
contains commands and queries that can be executed on the domain, for example
let's say the business requirement with forecast was to List the forecast, hence
a list function was provided by the service structure, suppose more requirements
like creating a forecast was present, this would mean a new function called
something like CreateForecast would be present as part of this structure

As mentioned earlier the details of how the actual data is fetched/stored is
irrelevant to the domain, hence one or more Repository interfaces to query
and/or modify data will be present as part of the individual domain package. 
Additionally a service provider interface must be defined which will be
used to help resolve the repository interfaces, these service providers must
be passed down as parameters to every service function. For example the forecast
domain contains the IForecastRepository that defines a List function and
IServiceProvider that defines a GetForecastRepository function, said
IServiceProvider is passed down as a parameter for the List function in the 
ForecastService struct

The individual domain package must also contain the relavent entity, if the
entity is to have other data which may be part of another domain (one to many or
many to many relationships) these are also to be defined here but only
containing the information relavent to the main domain, for example if our main
domain was Users, but a user has a set of roles, the combination of user and
role is to be part of the User domain, but only the Id of the role would be
required and no other info should be part of the User domain (those should be
part of the role domain)
[this](https://github.com/BetaLixT/goplates/tree/dev/samples/ddd-01)
sample that is part of the repo contains the above mentioned relationship. This
is based on my limited understanding on Aggrigates in DDD,
[linked](https://www.jamesmichaelhickey.com/domain-driven-design-aggregates/)
 is the article I used as a reference for this.

 In addition to what's mentioned above, the individual domain package may also
 contain error codes, validations, dtos etc as per the requirements, a slightly
 more complex domain is present in the
 [sample](https://github.com/BetaLixT/goplates/tree/dev/samples/ddd-01)
 project mentioned above
