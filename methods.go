package httpref

// Methods represents all of the defined HTTP methods
var Methods = References{
	{
		Name:    "Methods",
		IsTitle: true,
		Summary: "https://developer.mozilla.org/en-US/docs/Web/HTTP/MEthods",
		Description: `HTTP defines a set of request methods to indicate the desired action to be performed for a given resource. Although they can also be nouns, these request methods are sometimes referred as HTTP verbs. Each of them implements a different semantic, but some common features are shared by a group of them: e.g. a request method can be safe, idempotent, or cacheable.

https://developer.mozilla.org/en-US/docs/Web/HTTP/MEthods`,
	},
	{
		Name:    "GET",
		Summary: "The GET method requests a representation of the specified resource. Requests using GET should only retrieve data.",
		Description: `The HTTP GET method requests a representation of the specified resource. Requests using GET should only retrieve data.

Request has body              No
Successful response has body  Yes
Safe                          Yes
Idempotent                    Yes
Cacheable                     Yes
Allowed in HTML forms         Yes

https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/GET`,
	},
	{
		Name:    "HEAD",
		Summary: "The HEAD method asks for a response identical to that of a GET request, but without the response body.",
		Description: `The HTTP HEAD method requests the headers that are returned if the specified resource would be requested with an HTTP GET method. Such a request can be done before deciding to download a large resource to save bandwidth, for example.

A response to a HEAD method should not have a body. If so, it must be ignored. Even so, entity headers describing the content of the body, like Content-Length may be included in the response. They don't relate to the body of the HEAD response, which should be empty, but to the body of similar request using the GET method would have returned as a response.

If the result of a HEAD request shows that a cached resource after a GET request is now outdated, the cache is invalidated, even if no GET request has been made.

Request has body              No
Successful response has body  No
Safe                          Yes
Idempotent                    Yes
Cacheable                     Yes
Allowed in HTML forms         No

https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/HEAD`,
	},
	{
		Name:    "POST",
		Summary: "The POST method is used to submit an entity to the specified resource, often causing a change in state or side effects on the server.",
		Description: `The HTTP POST method sends data to the server. The type of the body of the request is indicated by the Content-Type header.

The difference between PUT and POST is that PUT is idempotent: calling it once or several times successively has the same effect (that is no side effect), where successive identical POST may have additional effects, like passing an order several times.

A POST request is typically sent via an HTML form and results in a change on the server. In this case, the content type is selected by putting the adequate string in the enctype attribute of the <form> element or the formenctype attribute of the <input> or <button> elements:

    application/x-www-form-urlencoded: the keys and values are encoded in key-value tuples separated by '&', with a '=' between the key and the value. Non-alphanumeric characters in both keys and values are percent encoded: this is the reason why this type is not suitable to use with binary data (use multipart/form-data instead)
    multipart/form-data: each value is sent as a block of data ("body part"), with a user agent-defined delimiter ("boundary") separating each part. The keys are given in the Content-Disposition header of each part.
    text/plain

When the POST request is sent via a method other than an HTML form — like via an XMLHttpRequest — the body can take any type. As described in the HTTP 1.1 specification, POST is designed to allow a uniform method to cover the following functions:

    Annotation of existing resources
    Posting a message to a bulletin board, newsgroup, mailing list, or similar group of articles;
    Adding a new user through a signup modal;
    Providing a block of data, such as the result of submitting a form, to a data-handling process;
    Extending a database through an append operation.

Request has body              Yes
Successful response has body  Yes
Safe                          No
Idempotent                    No
Cacheable                     Only if freshness information is included
Allowed in HTML forms         Yes

https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/POST`,
	},
	{
		Name:    "PUT",
		Summary: "The PUT method replaces all current representations of the target resource with the request payload.",
		Description: `The HTTP PUT request method creates a new resource or replaces a representation of the target resource with the request payload.

The difference between PUT and POST is that PUT is idempotent: calling it once or several times successively has the same effect (that is no side effect), where successive identical POST may have additional effects, like passing an order several times.

Request has body              Yes
Successful response has body  No
Safe                          No
Idempotent                    Yes
Cacheable                     No
Allowed in HTML forms         No

https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/PUT`,
	},
	{
		Name:    "DELETE",
		Summary: "The DELETE method deletes the specified resource.",
		Description: `The HTTP DELETE request method deletes the specified resource.

Request has body              May
Successful response has body  May
Safe                          No
Idempotent                    Yes
Cacheable                     No
Allowed in HTML forms         No

https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/DELETE`,
	},
	{
		Name:    "CONNECT",
		Summary: "The CONNECT method establishes a tunnel to the server identified by the target resource.",
		Description: `The HTTP CONNECT method starts two-way communications with the requested resource. It can be used to open a tunnel.

For example, the CONNECT method can be used to access websites that use SSL (HTTPS). The client asks an HTTP Proxy server to tunnel the TCP connection to the desired destination. The server then proceeds to make the connection on behalf of the client. Once the connection has been established by the server, the Proxy server continues to proxy the TCP stream to and from the client.

CONNECT is a hop-by-hop method.

Request has body              No
Successful response has body  Yes
Safe                          No
Idempotent                    No
Cacheable                     No
Allowed in HTML forms         No

https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/CONNECT`,
	},
	{
		Name:    "OPTIONS",
		Summary: "The OPTIONS method is used to describe the communication options for the target resource.",
		Description: `The HTTP OPTIONS method is used to describe the communication options for the target resource. The client can specify a URL for the OPTIONS method, or an asterisk (*) to refer to the entire server.

Request has body              No
Successful response has body  Yes
Safe                          Yes
Idempotent                    Yes
Cacheable                     No
Allowed in HTML forms         No

https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/OPTIONS`,
	},
	{
		Name:    "TRACE",
		Summary: "The TRACE method performs a message loop-back test along the path to the target resource.",
		Description: `The HTTP TRACE method performs a message loop-back test along the path to the target resource, providing a useful debugging mechanism.

The final recipient of the request should reflect the message received, excluding some fields described below, back to the client as the message body of a 200 (OK) response with a Content-Type of message/http. The final recipient is either the origin server or the first server to receive a Max-Forwards value of 0 in the request.

Request has body              No
Successful response has body  No
Safe                          No
Idempotent                    Yes
Cacheable                     No
Allowed in HTML forms         No

https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/TRACE`,
	},
	{
		Name:    "PATCH",
		Summary: "The PATCH method is used to apply partial modifications to a resource.",
		Description: `The HTTP PATCH request method applies partial modifications to a resource.

The HTTP PUT method only allows complete replacement of a document. Unlike PUT, PATCH is not idempotent, meaning successive identical patch requests may have different effects. However, it is possible to issue PATCH requests in such a way as to be idempotent.

PATCH (like PUT) may have side-effects on other resources.

To find out whether a server supports PATCH, a server can advertise its support by adding it to the list in the Allow or Access-Control-Allow-Methods (for CORS) response headers.

Another (implicit) indication that PATCH is allowed, is the presence of the Accept-Patch header, which specifies the patch document formats accepted by the server.

Request has body              Yes
Successful response has body  Yes
Safe                          No
Idempotent                    No
Cacheable                     No
Allowed in HTML forms         No

https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/PATCH`,
	},
}
