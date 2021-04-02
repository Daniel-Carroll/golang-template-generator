# df-efc-fulfillment

Repo for the EFC Fulfillment Service


## Package Structure

### cmd

Where the code that composes the packages and generates the application binary lives


### domain 

Structs and interfaces that define the shape and function of your application logic


### http

HTTP server implemenation using go-chi for routing

### schema

Flyway scripts that will generate the projects database(s) for you locally

### sql

The implementation of the interfaces in SQL

### Integration testing.

In order to run your integration tests locally, just type the command "./integration-test/run-integrations.sh" in the base directory. This will spin up an api container and an integration testing container
that will run all of the tests located inside of the integration-test directory. The linter is extremely sensitive so make sure everything is formatted properly. You can find all reports for the integration testing underneath the
"reports" directory. This is not to be pushed to github. Keep in mind that if there are any "Delete" or "Modify" integration tests you will most likely need to re-run the database volume as the integration tests will alter the data in the sample dataset. This is also key as each test after the initial one also matters sequential. If I do a "Delete ID 1" integration test and then I do a "Get ID 1" the Get will fail because the previous test removed that specific data row.