package httpref

// Statuses represents all of the defined HTTP statuses
var Statuses = References{
	{
		Name:    "1xx",
		IsTitle: true,
		Summary: "Informational response",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Status
https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#1xx_Informational_response`,
	},
	{
		Name:    "100",
		Summary: "Continue",
		Description: `The HTTP 100 Continue informational status response code indicates that everything so far is OK and that the client should continue with the request or ignore it if it is already finished.

To have a server check the request's headers, a client must send Expect: 100-continue as a header in its initial request and receive a 100 Continue status code in response before sending the body.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/100`,
	},
	{
		Name:    "101",
		Summary: "Switching Protocols",
		Description: `The HTTP 101 Switching Protocols response code indicates the protocol the server is switching to as requested by a client which sent the message including the Upgrade request header.

The server includes in this response an Upgrade response header to indicate the protocol it switched to. The process is described in detail in the article Protocol upgrade mechanism.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/101`,
	},
	{
		Name:        "102",
		Summary:     "Processing (WebDAV)",
		Description: `This code indicates that the server has received and is processing the request, but no response is available yet.`,
	},
	{
		Name:    "103",
		Summary: "Early Hints",
		Description: `The HTTP 103 Early Hints information response status code is primarily intended to be used with the Link header to allow the user agent to start preloading resources while the server is still preparing a response.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/103`,
	},
	{
		Name:    "2xx",
		IsTitle: true,
		Summary: "Successful responses",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Status
https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#2xx_Success`,
	},
	{
		Name:    "200",
		Summary: "OK",
		Description: `The HTTP 200 OK success status response code indicates that the request has succeeded. A 200 response is cacheable by default.

The meaning of a success depends on the HTTP request method:

    GET: The resource has been fetched and is transmitted in the message body.
    HEAD: The entity headers are in the message body.
    POST: The resource describing the result of the action is transmitted in the message body.
    TRACE: The message body contains the request message as received by the server.

The successful result of a PUT or a DELETE is often not a 200 OK but a 204 No Content (or a 201 Created when the resource is uploaded for the first time).

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/200`,
	},
	{
		Name:    "201",
		Summary: "Created",
		Description: `The HTTP 201 Created success status response code indicates that the request has succeeded and has led to the creation of a resource. The new resource is effectively created before this response is sent back and the new resource is returned in the body of the message, its location being either the URL of the request, or the content of the Location header.

The common use case of this status code is as the result of a POST request.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/201`,
	},
	{
		Name:    "202",
		Summary: "Accepted",
		Description: `The HyperText Transfer Protocol (HTTP) 202 Accepted response status code indicates that the request has been received but not yet acted upon. It is non-committal, meaning that there is no way for the HTTP to later send an asynchronous response indicating the outcome of processing the request. It is intended for cases where another process or server handles the request, or for batch processing.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/202`,
	},
	{
		Name:    "203",
		Summary: "Non-Authoritative Information ",
		Description: `The HTTP 203 Non-Authoritative Information response status indicates that the request was successful but the enclosed payload has been modified by a transforming proxy from that of the origin server's 200 (OK) response .

The 203 response is similar to the value 214, meaning Transformation Applied, of the Warning header code, which has the additional advantage of being applicable to responses with any status code.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/203`,
	},
	{
		Name:    "204",
		Summary: "No Content",
		Description: `The HTTP 204 No Content success status response code indicates that the request has succeeded, but that the client doesn't need to go away from its current page. A 204 response is cacheable by default. An ETag header is included in such a response.

The common use case is to return 204 as a result of a PUT request, updating a resource, without changing the current content of the page displayed to the user. If the resource is created, 201 Created is returned instead. If the page should be changed to the newly updated page, the 200 should be used instead.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/204`,
	},
	{
		Name:    "205",
		Summary: "Reset Content",
		Description: `The HTTP 205 Reset Content response status tells the client to reset the document view, so for example to clear the content of a form, reset a canvas state, or to refresh the UI.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/205`,
	},
	{
		Name:    "206",
		Summary: "Partial Content (WebDAV)",
		Description: `The HTTP 206 Partial Content success status response code indicates that the request has succeeded and has the body contains the requested ranges of data, as described in the Range header of the request.

If there is only one range, the Content-Type of the whole response is set to the type of the document, and a Content-Range is provided.

If several ranges are sent back, the Content-Type is set to multipart/byteranges and each fragment covers one range, with Content-Range and Content-Type describing it.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/206`,
	},
	{
		Name:    "207",
		Summary: "Multi-Status (WebDAV)",
		Description: `Conveys information about multiple resources, for situations where multiple status codes might be appropriate.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status`,
	},
	{
		Name:    "208",
		Summary: "Already Reported (WebDAV)",
		Description: `Used inside a <dav:propstat> response element to avoid repeatedly enumerating the internal members of multiple bindings to the same collection.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status`,
	},
	{
		Name:    "226",
		Summary: "IM Used",
		Description: `The server has fulfilled a GET request for the resource, and the response is a representation of the result of one or more instance-manipulations applied to the current instance.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status`,
	},
	{
		Name:    "3xx",
		IsTitle: true,
		Summary: "Redirection messages",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Status
https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#3xx_Redirection`,
	},
	{
		Name:    "300",
		Summary: "Multiple Choices",
		Description: `The HTTP 300 Multiple Choices redirect status response code indicates that the request has more than one possible responses. The user-agent or the user should choose one of them. As there is no standardized way of choosing one of the responses, this response code is very rarely used.

If the server has a preferred choice, it should generate a Location header.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/300`,
	},
	{
		Name:    "301",
		Summary: "Moved Permanently",
		Description: `The HyperText Transfer Protocol (HTTP) 301 Moved Permanently redirect status response code indicates that the resource requested has been definitively moved to the URL given by the Location headers. A browser redirects to this page and search engines update their links to the resource (in 'SEO-speak', it is said that the 'link-juice' is sent to the new URL).

Even if the specification requires the method (and the body) not to be altered when the redirection is performed, not all user-agents align with it - you can still find this type of bugged software out there. It is therefore recommended to use the 301 code only as a response for GET or HEAD methods and to use the 308 Permanent Redirect for POST methods instead, as the method change is explicitly prohibited with this status.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/301`,
	},
	{
		Name:    "302",
		Summary: "Found",
		Description: `The HyperText Transfer Protocol (HTTP) 302 Found redirect status response code indicates that the resource requested has been temporarily moved to the URL given by the Location header. A browser redirects to this page but search engines don't update their links to the resource (in 'SEO-speak', it is said that the 'link-juice' is not sent to the new URL).

Even if the specification requires the method (and the body) not to be altered when the redirection is performed, not all user-agents conform here - you can still find this type of bugged software out there. It is therefore recommended to set the 302 code only as a response for GET or HEAD methods and to use 307 Temporary Redirect instead, as the method change is explicitly prohibited in that case.

In the cases where you want the method used to be changed to GET, use 303 See Other instead. This is useful when you want to give a response to a PUT method that is not the uploaded resource but a confirmation message such as: 'you successfully uploaded XYZ'.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/302`,
	},
	{
		Name:    "303",
		Summary: "See other",
		Description: `The HyperText Transfer Protocol (HTTP) 303 See Other redirect status response code indicates that the redirects don't link to the newly uploaded resources, but to another page (such as a confirmation page or an upload progress page). This response code is usually sent back as a result of PUT or POST. The method used to display this redirected page is always GET.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/303`,
	},
	{
		Name:    "304",
		Summary: "Not Modified",
		Description: `The HTTP 304 Not Modified client redirection response code indicates that there is no need to retransmit the requested resources. It is an implicit redirection to a cached resource. This happens when the request method is safe, like a GET or a HEAD request, or when the request is conditional and uses a If-None-Match or a If-Modified-Since header.

The equivalent 200 OK response would have included the headers Cache-Control, Content-Location, Date, ETag, Expires, and Vary.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/304`,
	},
	{
		Name:    "305",
		Summary: "Use proxy",
		Description: `Defined in a previous version of the HTTP specification to indicate that a requested response must be accessed by a proxy. It has been deprecated due to security concerns regarding in-band configuration of a proxy.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status`,
	},
	{
		Name:    "306",
		Summary: "unused",
		Description: `This response code is no longer used; it is just reserved. It was used in a previous version of the HTTP/1.1 specification.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status`,
	},
	{
		Name:    "307",
		Summary: "Temporary redirect",
		Description: `HTTP 307 Temporary Redirect redirect status response code indicates that the resource requested has been temporarily moved to the URL given by the Location headers.

The method and the body of the original request are reused to perform the redirected request. In the cases where you want the method used to be changed to GET, use 303 See Other instead. This is useful when you want to give an answer to a PUT method that is not the uploaded resources, but a confirmation message (like "You successfully uploaded XYZ").

The only difference between 307 and 302 is that 307 guarantees that the method and the body will not be changed when the redirected request is made. With 302, some old clients were incorrectly changing the method to GET: the behavior with non-GET methods and 302 is then unpredictable on the Web, whereas the behavior with 307 is predictable. For GET requests, their behavior is identical.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/307`,
	},
	{
		Name:    "308",
		Summary: "Permanent redirect",
		Description: `The HyperText Transfer Protocol (HTTP) 308 Permanent Redirect redirect status response code indicates that the resource requested has been definitively moved to the URL given by the Location headers. A browser redirects to this page and search engines update their links to the resource (in 'SEO-speak', it is said that the 'link-juice' is sent to the new URL).

The request method and the body will not be altered, whereas 301 may incorrectly sometimes be changed to a GET method.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/308`,
	},
	{
		Name:    "4xx",
		IsTitle: true,
		Summary: "Client error responses",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Status
https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#4xx_Client_errors`,
	},
	{
		Name:    "400",
		Summary: "Bad request",
		Description: `The HyperText Transfer Protocol (HTTP) 400 Bad Request response status code indicates that the server cannot or will not process the request due to something that is perceived to be a client error (e.g., malformed request syntax, invalid request message framing, or deceptive request routing).

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400`,
	},
	{
		Name:    "401",
		Summary: "Unauthorized",
		Description: `The HTTP 401 Unauthorized client error status response code indicates that the request has not been applied because it lacks valid authentication credentials for the target resource.

This status is sent with a WWW-Authenticate header that contains information on how to authorize correctly.

This status is similar to 403, but in this case, authentication is possible.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401`,
	},
	{
		Name:    "402",
		Summary: "Payment required",
		Description: `The HTTP 402 Payment Required is a nonstandard client error status response code that is reserved for future use.

Sometimes, this code indicates that the request can not be processed until the client makes a payment. Originally it was created to enable digital cash or (micro) payment systems and would indicate that the requested content is not available until the client makes a payment. However, no standard use convention exists and different entities use it in different contexts.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/402`,
	},
	{
		Name:    "403",
		Summary: "Forbidden",
		Description: `The HTTP 403 Forbidden client error status response code indicates that the server understood the request but refuses to authorize it.

This status is similar to 401, but in this case, re-authenticating will make no difference. The access is permanently forbidden and tied to the application logic, such as insufficient rights to a resource.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403`,
	},
	{
		Name:    "404",
		Summary: "Not found",
		Description: `The HTTP 404 Not Found client error response code indicates that the server can't find the requested resource. Links which lead to a 404 page are often called broken or dead links, and can be subject to link rot.

A 404 status code does not indicate whether the resource is temporarily or permanently missing. But if a resource is permanently removed, a 410 (Gone) should be used instead of a 404 status.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404`,
	},
	{
		Name:    "405",
		Summary: "Method not allowed",
		Description: `The HyperText Transfer Protocol (HTTP) 405 Method Not Allowed response status code indicates that the request method is known by the server but is not supported by the target resource.

The server MUST generate an Allow header field in a 405 response containing a list of the target resource's currently supported methods.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/405`,
	},
	{
		Name:    "406",
		Summary: "Not acceptable",
		Description: `The HyperText Transfer Protocol (HTTP) 406 Not Acceptable client error response code indicates that the server cannot produce a response matching the list of acceptable values defined in the request's proactive content negotiation headers, and that the server is unwilling to supply a default representation.

Proactive content negotiation headers include:

    Accept
    Accept-Charset
    Accept-Encoding
    Accept-Language

In practice, this error is very rarely used. Instead of responding using this error code, which would be cryptic for the end user and difficult to fix, servers ignore the relevant header and serve an actual page to the user. It is assumed that even if the user won't be completely happy, they will prefer this to an error code.

If a server returns such an error status, the body of the message should contain the list of the available representations of the resources, allowing the user to choose among them.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/406`,
	},
	{
		Name:    "407",
		Summary: "Proxy authentication required",
		Description: `The HTTP 407 Proxy Authentication Required client error status response code indicates that the request has not been applied because it lacks valid authentication credentials for a proxy server that is between the browser and the server that can access the requested resource.

This status is sent with a Proxy-Authenticate header that contains information on how to authorize correctly.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/407`,
	},
	{
		Name:    "408",
		Summary: "Request timeout",
		Description: `The HyperText Transfer Protocol (HTTP) 408 Request Timeout response status code means that the server would like to shut down this unused connection. It is sent on an idle connection by some servers, even without any previous request by the client.

A server should send the "close" Connection header field in the response, since 408 implies that the server has decided to close the connection rather than continue waiting.

This response is used much more since some browsers, like Chrome, Firefox 27+, and IE9, use HTTP pre-connection mechanisms to speed up surfing.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408`,
	},
	{
		Name:    "409",
		Summary: "Conflict",
		Description: `The HTTP 409 Conflict response status code indicates a request conflict with current state of the server.

Conflicts are most likely to occur in response to a PUT request. For example, you may get a 409 response when uploading a file which is older than the one already on the server resulting in a version control conflict.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/409`,
	},
	{
		Name:    "410",
		Summary: "Gone",
		Description: `The HyperText Transfer Protocol (HTTP) 410 Gone client error response code indicates that access to the target resource is no longer available at the origin server and that this condition is likely to be permanent.

If you don't know whether this condition is temporary or permanent, a 404 status code should be used instead.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/410`,
	},
	{
		Name:    "411",
		Summary: "Length required",
		Description: `The HyperText Transfer Protocol (HTTP) 411 Length Required client error response code indicates that the server refuses to accept the request without a defined Content-Length header.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/411`,
	},
	{
		Name:    "412",
		Summary: "Precondition failed",
		Description: `The HyperText Transfer Protocol (HTTP) 412 Precondition Failed client error response code indicates that access to the target resource has been denied. This happens with conditional requests on methods other than GET or HEAD when the condition defined by the If-Unmodified-Since or If-None-Match headers is not fulfilled. In that case, the request, usually an upload or a modification of a resource, cannot be made and this error response is sent back.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/412`,
	},
	{
		Name:    "413",
		Summary: "Payload too large",
		Description: `The HTTP 413 Payload Too Large response status code indicates that the request entity is larger than limits defined by server; the server might close the connection or return a Retry-After header field.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/413`,
	},
	{
		Name:    "414",
		Summary: "URI too long",
		Description: `The HTTP 414 URI Too Long response status code indicates that the URL requested by the client is longer than the server is willing to interpret.

There are a few conditions when this might occur:

    A client has improperly converted a POST request to a GET request with more than ≈2 kB of submitted data.
    A client has descended into a loop of redirection (for example, a redirected URL prefix that points to a suffix of itself, or mishandled relative URLs),
    The server is under attack by a client attempting to exploit potential security holes.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/414`,
	},
	{
		Name:    "415",
		Summary: "Unsupported media type",
		Description: `The HTTP 415 Unsupported Media Type client error response code indicates that the server refuses to accept the request because the payload format is in an unsupported format.

The format problem might be due to the request's indicated Content-Type or Content-Encoding, or as a result of inspecting the data directly.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/415`,
	},
	{
		Name:    "416",
		Summary: "Range not satisfiable",
		Description: `The HyperText Transfer Protocol (HTTP) 416 Range Not Satisfiable error response code indicates that a server cannot serve the requested ranges. The most likely reason is that the document doesn't contain such ranges, or that the Range header value, though syntactically correct, doesn't make sense.

The 416 response message contains a Content-Range indicating an unsatisfied range (that is a '*') followed by a '/' and the current length of the resource. E.g. Content-Range: bytes */12777

Faced with this error, browsers usually either abort the operation (for example, a download will be considered as non-resumable) or ask for the whole document again.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/416`,
	},
	{
		Name:    "417",
		Summary: "Expectation failed",
		Description: `The HTTP 417 Expectation Failed client error response code indicates that the expectation given in the request's Expect header could not be met.

See the Expect header for more details.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/417`,
	},
	{
		Name:    "418",
		Summary: "I'm a teapot",
		Description: `The HTTP 418 I'm a teapot client error response code indicates that the server refuses to brew coffee because it is a teapot. This error is a reference to Hyper Text Coffee Pot Control Protocol which was an April Fools' joke in 1998.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/418`,
	},
	{
		Name:    "421",
		Summary: "Misdirected Request",
		Description: `The request was directed at a server that is not able to produce a response. This can be sent by a server that is not configured to produce responses for the combination of scheme and authority that are included in the request URI.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status`,
	},
	{
		Name:    "422",
		Summary: "Unprocessable Entity (WebDAV)",
		Description: `The HyperText Transfer Protocol (HTTP) 422 Unprocessable Entity response status code indicates that the server understands the content type of the request entity, and the syntax of the request entity is correct, but it was unable to process the contained instructions.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422`,
	},
	{
		Name:    "423",
		Summary: "Locked (WebDAV)",
		Description: `The resource that is being accessed is locked.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status`,
	},
	{
		Name:    "424",
		Summary: "Failed Dependency (WebDAV)",
		Description: `The request failed due to failure of a previous request.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status`,
	},
	{
		Name:    "425",
		Summary: "Too early",
		Description: `The HyperText Transfer Protocol (HTTP) 425 Too Early response status code indicates that the server is unwilling to risk processing a request that might be replayed, which creates the potential for a replay attack.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/425`,
	},
	{
		Name:    "426",
		Summary: "Upgrade required",
		Description: `The HTTP 426 Upgrade Required client error response code indicates that the server refuses to perform the request using the current protocol but might be willing to do so after the client upgrades to a different protocol.

The server sends an Upgrade header with this response to indicate the required protocol(s).

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/426`,
	},
	{
		Name:    "428",
		Summary: "Precondition required",
		Description: `The HTTP 428 Precondition Required response status code indicates that the server requires the request to be conditional.

Typically, this means that a required precondition header, such as If-Match, is missing.

When a precondition header is not matching the server side state, the response should be 412 Precondition Failed.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/428`,
	},
	{
		Name:    "429",
		Summary: "Too many requests",
		Description: `The HTTP 429 Too Many Requests response status code indicates the user has sent too many requests in a given amount of time ("rate limiting").

A Retry-After header might be included to this response indicating how long to wait before making a new request.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429`,
	},
	{
		Name:    "431",
		Summary: "Request Header Fields Too Large",
		Description: `The HTTP 431 Request Header Fields Too Large response status code indicates that the server refuses to process the request because the request’s HTTP headers are too long. The request may be resubmitted after reducing the size of the request headers.

431 can be used when the total size of request headers is too large, or when a single header field is too large. To help those running into this error, indicate which of the two is the problem in the response body — ideally, also include which headers are too large. This lets users attempt to fix the problem, such as by clearing their cookies.

Servers will often produce this status if:

    The Referer URL is too long
    There are too many Cookies sent in the request

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/431`,
	},
	{
		Name:    "451",
		Summary: "Unavailable For Legal Reasons",
		Description: `The HyperText Transfer Protocol (HTTP) 451 Unavailable For Legal Reasons client error response code indicates that the user requested a resource that is not available due to legal reasons, such as a web page for which a legal action has been issued.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/451`,
	},
	{
		Name:    "5xx",
		IsTitle: true,
		Summary: "Server error responses",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Status
https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#5xx_Server_errors`,
	},
	{
		Name:    "500",
		Summary: "Internal server error",
		Description: `The HyperText Transfer Protocol (HTTP) 500 Internal Server Error server error response code indicates that the server encountered an unexpected condition that prevented it from fulfilling the request.

This error response is a generic "catch-all" response. Usually, this indicates the server cannot find a better 5xx error code to response. Sometimes, server administrators log error responses like the 500 status code with more details about the request to prevent the error from happening again in the future.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500`,
	},
	{
		Name:    "501",
		Summary: "Not implemented",
		Description: `The HyperText Transfer Protocol (HTTP) 501 Not Implemented server error response code means that the server does not support the functionality required to fulfill the request.

This status can also send a Retry-After header, telling the requester when to check back to see if the functionality is supported by then.

501 is the appropriate response when the server does not recognize the request method and is incapable of supporting it for any resource. The only methods that servers are required to support (and therefore that must not return 501) are GET and HEAD.

If the server does recognize the method, but intentionally does not support it, the appropriate response is 405 Method Not Allowed.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/501`,
	},
	{
		Name:    "502",
		Summary: "Bad gateway",
		Description: `The HyperText Transfer Protocol (HTTP) 502 Bad Gateway server error response code indicates that the server, while acting as a gateway or proxy, received an invalid response from the upstream server.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/502`,
	},
	{
		Name:    "503",
		Summary: "Service unavailable",
		Description: `The HyperText Transfer Protocol (HTTP) 503 Service Unavailable server error response code indicates that the server is not ready to handle the request.

Common causes are a server that is down for maintenance or that is overloaded. This response should be used for temporary conditions and the Retry-After HTTP header should, if possible, contain the estimated time for the recovery of the service.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/503`,
	},
	{
		Name:    "504",
		Summary: "Gateway timeout",
		Description: `The HyperText Transfer Protocol (HTTP) 504 Gateway Timeout server error response code indicates that the server, while acting as a gateway or proxy, did not get a response in time from the upstream server that it needed in order to complete the request.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/504`,
	},
	{
		Name:    "505",
		Summary: "HTTP version no supported",
		Description: `The HyperText Transfer Protocol (HTTP) 505 HTTP Version Not Supported response status code indicates that the HTTP version used in the request is not supported by the server.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/505`,
	},
	{
		Name:    "506",
		Summary: "Variant also negotiates",
		Description: `The HyperText Transfer Protocol (HTTP) 506 Variant Also Negotiates response status code may be given in the context of Transparent Content Negotiation (see RFC 2295). This protocol enables a client to retrieve the best variant of a given resource, where the server supports multiple variants.

The Variant Also Negotiates status code indicates an internal server configuration error in which the chosen variant is itself configured to engage in content negotiation, so is not a proper negotiation endpoint.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/506`,
	},
	{
		Name:    "507",
		Summary: "Insufficient Storage (WebDAV)",
		Description: `The HyperText Transfer Protocol (HTTP)  507 Insufficient Storage response status code may be given in the context of the Web Distributed Authoring and Versioning (WebDAV) protocol (see RFC 4918).

It indicates that a method could not be performed because the server cannot store
the representation needed to successfully complete the request.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/507`,
	},
	{
		Name:    "508",
		Summary: "Loop Detected (WebDAV)",
		Description: `The HyperText Transfer Protocol (HTTP) 508 Loop Detected response status code may be given in the context of the Web Distributed Authoring and Versioning (WebDAV) protocol.

It indicates that the server terminated an operation because it encountered an infinite loop while processing a request with "Depth: infinity". This status indicates that the entire operation failed.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/508`,
	},
	{
		Name:    "510",
		Summary: "Not extended",
		Description: `The HyperText Transfer Protocol (HTTP)  510 Not Extended response status code is sent in the context of the HTTP Extension Framework, defined in RFC 2774.

In that specification a client may send a request that contains an extension declaration, that describes the extension to be used. If the server receives such a request, but any described extensions are not supported for the request, then the server responds with the 510 status code.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/510`,
	},
	{
		Name:    "511",
		Summary: "Network Authentication Required",
		Description: `The HTTP 511 Network Authentication Required response status code indicates that the client needs to authenticate to gain network access.

This status is not generated by origin servers, but by intercepting proxies that control access to the network.

Network operators sometimes require some authentication, acceptance of terms, or other user interaction before granting access (for example in an internet café or at an airport). They often identify clients who have not done so using their Media Access Control (MAC) addresses.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/511`,
	},
}
