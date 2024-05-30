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
