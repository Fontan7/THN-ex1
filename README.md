# THN-ex 1
REST API with simple audit features <br>

Context: <br>
  We want to trace all the requests made to a single endpoint /feature, writing an audit
log with some information about the request and saving the number of requests made
by each IP.<br>
  The API will have another endpoint /metrics to query, given a specific IP, how many
requests have been made from this IP to the first endpoint.
<br>
<br>

Additional info: <br>
  ● The GET /feature endpoint will return anything (hello world text is enough!). <br>
  ● After accessing the GET endpoint, you should write (stdout is valid) some
  information about the request, including a timestamp, the IP, headers, etc. <br>
  ● The GET /metrics endpoint will receive the IP as a query parameter (mandatory)
  and will return a JSON with the number of requests. <br>
  ● It is not necessary to have any kind of persistence,you can save the number of
  requests in a temporary resource, in memory, etc. <br>

=====================================================================================================

##How to use?

If you are rocking windows simply run the .exe file and then hit localhost on port 8080. <br>
You can also navigate to the root of this project directory and run: go run .. <br>

Then you might want to hit url /swagger/index.html for docs about the api w.w  <br>
There you will find /health, v1/feature and v1/metrics endpoints documented, <br>
Tip on v1/metrics: <br>
we don't want to expose that much important data to anyone right? thats why I included some top notch millitary security! <br>
to access it you will need to include a "X-API-Key" header with the value "THN_KEY" :o <br>
