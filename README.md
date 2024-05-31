# THN-ex 1
REST API with simple audit features

## Context:
We want to trace all the requests made to a single endpoint `/feature`, writing an audit log with some information about the request and saving the number of requests made by each IP.

The API will have another endpoint `/metrics` to query, given a specific IP, how many requests have been made from this IP to the first endpoint.

## Additional info:
- The GET `/feature` endpoint will return anything (hello world text is enough!).
- After accessing the GET endpoint, you should write (stdout is valid) some information about the request, including a timestamp, the IP, headers, etc.
- The GET `/metrics` endpoint will receive the IP as a query parameter (mandatory) and will return a JSON with the number of requests.
- It is not necessary to have any kind of persistence; you can save the number of requests in a temporary resource, in memory, etc.

===========================================================================

## How to use?

If you are using Windows, simply run the `.exe` file (I promise it has no viruses or ill intent) and then access `localhost` on port 8080.
You can also navigate to the root of this project directory and run: `go run .`

Then you might want to visit the URL `/swagger/index.html` for documentation about the API.
There you will find `/health`, `/v1/feature`, and `/v1/metrics` endpoints documented.

### Tip on `/v1/metrics`:
We don't want to expose such important data to anyone, right? That's why I included some top-notch military security!
To access it, you will need to include an "X-API-Key" header with the value "THN_KEY". :o
