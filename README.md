# DZ-Edge-API

[![Go Report Card](https://goreportcard.com/badge/github.com/istanleyy/dz-edge-api)](https://goreportcard.com/report/github.com/istanleyy/dz-edge-api)
[![Release](https://img.shields.io/github/v/release/istanleyy/dz-edge-api?include_prereleases)](https://github.com/istanleyy/dz-edge-api/releases)

DZ Edge API is an API server based on Gin that interacts with shop floor control systems (PLC's) to synchronize data with iShopFloor MOM solution.

## Setup

To build and run DZ-Edge-API, you need to install Go first (**version 1.11+ is required**).
Once you have setup Go working environment, you can clone and build the repo with:

```bash
# clone the project from GitHub
$ git clone https://github.com/istanleyy/dz-edge-api.git

# build the project and save executable to the build/ directory
$ cd dz-edge-api
$ go build -o build

# cross compiling for Windows can use the following command
$ GOOS=windows go build -o build
```

## Quick Start

Follow the above build procedure, you may find the executable file in the build folder located under the project root directory. Start the DZ-Edge-API server with:

```bash
# assuming at project root directory
$ ./dz-edge-api
```

This starts the API server at localhost:8080. By default, current version will look for WebAccess API server on the same machine running DZ-Edge-API. If WebAccess API server is unreachable when DZ-Edge-API launches, it will keep retrying to establish the connection. API endpoints will become available only when the connection with WebAccess is established.

## API Examples

### POST /job/create

CreateJob adds a new job with the following fields in the request body.

Sample request body:

```json
{
    "woID": "5edf6dcfc5dbda7ac93c7f19",
    "process": "gal"
}
```

- "woID" - 24-char work ID
- "process" - 'pre' for 'pre-process', and 'gal' for galvanization

Sample HTTP 200 response body:

```json
{
    "queueIndex": 0,
    "woID": "gbe732fd831d6cea6b38d256"
}
```

- "queueIndex" - position of the work in queue
- "woID" - 24-char work ID

Sample HTTP error body:

```json
{
    "message": "job exists in galvanizing process queue"
}
```

- "message" - the error message

### POST /job/update

UpdateJob updates the job status. Fields for the request body.

Sample request body:

```json
{
    "woID": "5edf6dcfc5dbda7ac93c7f19",
    "process": "gal",
    "eventID": 0
}
```

- "woID" - 24-char work ID
- "process" - 'pre' for 'pre-process', and 'gal' for galvanization
- "eventID" - the event code to update the job with, 0(finish order), 10(weight empty bucket), 11(weight full bucket), 12(weight good product)

Sample HTTP 200 response body:

```json
{
    "queueIndex": 0,
    "woID": "gbe732fd831d6cea6b38d256"
}
```

- "queueIndex" - position of the work in queue
- "woID" - 24-char work ID

Sample HTTP error body:

```json
{
    "code": "u001",
    "message": "unknown EventID for the pre-treatment process"
}
```

- "code" - the error code, 'u001' refers to DZ-Edge-API error, 'u002' refers to iShopFloor API error, 'u003' refers to WebAccess API error
- "message" - the error message

### GET /job/:jobId/:processType

GetJob responds with the data in the queue identified by the uri parameters.

Sample request:

```bash
curl http://localhost:8080/api/v1/job/gbe732fd831d6cea6b38d254/pre
```

- "jobId" - 24-char work ID
- "processType" - 'pre' for 'pre-process', and 'gal' for galvanization

Sample HTTP 200 response body for 'pre':

```json
{
    "process": "Pre-treatment",
    "woID": "gbe732fd831d6cea6b38d254",
    "emptyWeight": 110.2,
    "fullWeight": 250.7,
    "skimTime1": 10.1,
    "skimTime2": 20.2,
    "skimTime3": 30.3,
    "picklingTime1": 50.5,
    "picklingTime2": 70.3,
    "picklingTime3": 89.1,
    "fluxTime1": 500.5
}
```

Sample HTTP 200 response body for 'gal':

```json
{
    "process": "Galvanized",
    "woID": "gbe732fd831d6cea6b38d256",
    "finishEmptyWeight": 110.2,
    "finishFullWeight": 250.7,
    "littleBlueWeight": [
        1.2,
        1.3,
        1.4,
        1.5,
        1.6,
        1.7
    ],
    "littleBlueTime": [
        3.1,
        3.2,
        3.3,
        3.4,
        3.5,
        3.6
    ],
    "littleBlueTemp": [
        2.1,
        2.2,
        2.3,
        2.4,
        2.5,
        2.6
    ]
}
```

Sample HTTP error body:

```json
{
    "message": "cannot find job in pre-treatment process"
}
```

- "message" - the error message

### GET /process/:operation

GetProcMetrics responds with the constant read register containing the values for the metrics of the given operations.

Sample request:

```bash
curl http://localhost:8080/api/v1/process/1
```

- "operation" - the operation code of the process, '1' for degreasing #1, '2' for degreasing #2, '3' for pickling #1, '4' for pickling #2, '5' for pickling #3, '6' for galvanization

Sample HTTP 200 response body for process '1' to '5':

```json
{
    "process": "Pre-treatment",
    "woID": "",
    "operation": 1,
    "temp": 0,
    "ph": 0,
    "conductivity": 0
}
```

Sample HTTP 200 response body for process '6':

```json
{
    "process": "Galvanized",
    "woID": "",
    "operation": 6,
    "setTemp": 0,
    "actTemp": 0
}
```

Sample HTTP error body:

```json
{
    "message": "unknown operation"
}
```

- "message" - the error message

### DELETE /job/all

[DEV ONLY] Remove all jobs from the queues by setting their status to '0' in the PLC.

Sample request:

```bash
curl -X DELETE http://localhost:8080/api/v1/job/all
```

Sample HTTP 200 response body:

```json
{
    "message": "removed all jobs"
}
```

- "message" - the response message

Sample HTTP error body:

```json
{
    "message": "failed to remove all jobs"
}
```

- "message" - the error message
