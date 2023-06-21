# Property Manager
A simple key-value store using an HTTP API.

This is intended to be used in place of a "secrets manager" or database table that manages key/values. This is a long-term storage, values get written to a file and can be shared across sessions or services.

See the `PropertyManagerAPI.yml` file for API details.

## Service

### Environment Variables
The `.env_template` file provides the list of environment variables that can be used. Below will be a quick description of each

#### STORAGE_FILE
Path to a txt file to use a storage location for properties. This must be defined.

#### API_PORT
Port to listen on. Default: 8081

### Running the project
Run with `go run main.go`.

## SDK

## UI
