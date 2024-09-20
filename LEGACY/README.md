# LocalWebServices
A collection of containers meant to imitate common cloud services that may be used as supporting containers in early development stages.

The ideal use case would be to use these containers alongside another project, via docker compose.

## Running Locally
Each service has it's own directory in this repository. Each directory has a `service/` and a `ui/` child directory.   
The service directory holds the code to run each service. See the README for that particular service for instructions. 
### Running the Testing UI for each service
1. Build the sdk in the `sdk/` directory
```sh
cd sdk && npm install && npm run build
```
2. Then in the `ui/` folder for the service you want to test, start the UI.
```sh
npm install && npm run dev
```

## Services
### Serverless Functions (Stateless Functions)
* all in js/ts
* Runtime:
  * node
  * deno
* Use express for request / response
* mount a "functions" folder with scripts, or subdirectories with index files
* Need to be able to define common env variables for all functions

### Serverless Cloud Map
* map urls to serverless functions from above service

### Secrets Manager
* TODO: makr this as completed and give more details
* Key-Value store defined in config file
* retrieved with https requests?
  * make an sdk to wrap this?
* Needs a UI to manage

### Object Storage
* Needs a UI to manage
* Mount storage as volume
* Challenge: mimic S3 api?

### Queue Service
* Locked polling 
  * i.e make sure only one client can detect an event
* Needs a UI to poll / send / delete events
* make a JS sdk?
