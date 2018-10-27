# Clamor Campaign Micro-Service.

A simple microservice that controls the campaign data.

## Auth

Access to any endpoint requires a valid JWT to be passed in the Authentication header as a Bearer token.

## Endpoints

### /about

Only handles GET requests. Returns back a simple json object that gives info like microserivce build date and version details.

### /campaign

This handles GET, POST, PUT, and DELETE. It's really basic crud around the Campaign model defined in src/models/campaign.go

GET - if there is an id query string param it'll pull that single model. If there is no it'll return all models the user has access to.

POST - Pass a new campaign model as json in the body and it'll create a new campaign.

PUT - Update an existing campaign by passing json in the body.

DELETE - Will soft delete a campaign based on the id query string. This sets a deleted flag in the db.

### /campaignTypes

This only handles GET Requests. This will always return back an array of campaign type models (src/models/campaignType.go). These aren't mutable.
