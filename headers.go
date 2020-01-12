package httpref

var Headers = References{
	{
		Name:    "Headers",
		IsTitle: true,
		Summary: "Guidance about headers",
		Description: `HTTP headers let the client and the server pass additional information with an HTTP request or response. An HTTP header consists of its case-insensitive name followed by a colon (:), then by its value. Whitespace before the value is ignored.

Custom proprietary headers have historically been used with an X- prefix, but this convention was deprecated in June 2012 because of the inconveniences it caused when nonstandard fields became standard in RFC 6648; others are listed in an IANA registry, whose original content was defined in RFC 4229. IANA also maintains a registry of proposed new HTTP headers.

Headers can be grouped according to their contexts:

    General headers apply to both requests and responses, but with no relation to the data transmitted in the body.
    Request headers contain more information about the resource to be fetched, or about the client requesting the resource.
    Response headers hold additional information about the response, like its location or about the server providing it.
    Entity headers contain information about the body of the resource, like its content length or MIME type.

Headers can also be grouped according to how proxies handle them:

    Connection
    Keep-Alive
    Proxy-Authenticate
    Proxy-Authorization
    TE
    Trailer
    Transfer-Encoding
    Upgrade.

End-to-end headers
    These headers must be transmitted to the final recipient of the message: the server for a request, or the client for a response. Intermediate proxies must retransmit these headers unmodified and caches must store them.
Hop-by-hop headers
    These headers are meaningful only for a single transport-level connection, and must not be retransmitted by proxies or cached. Note that only hop-by-hop headers may be set using the Connection general header.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers`,
	},
	{
		Name:        "Authentication",
		IsTitle:     true,
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Authentication",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Authentication`,
	},
	{
		Name:    "WWW-Authenticate",
		Summary: "Defines the authentication method that should be used to access a resource.",
		Description: `The HTTP WWW-Authenticate response header defines the authentication method that should be used to gain access to a resource.

The WWW-Authenticate header is sent along with a 401 Unauthorized response.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/WWW-Authenticate`,
	},
	{
		Name:    "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Authorization",
		Summary: "Contains the credentials to authenticate a user-agent with a server.",
		Description: `The HTTP Authorization request header contains the credentials to authenticate a user agent with a server, usually after the server has responded with a 401 Unauthorized status and the WWW-Authenticate header.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Authorization`,
	},
	{
		Name:    "Proxy-Authenticate",
		Summary: "Defines the authentication method that should be used to access a resource behind a proxy server.",
		Description: `The HTTP Proxy-Authenticate response header defines the authentication method that should be used to gain access to a resource behind a proxy server. It authenticates the request to the proxy server, allowing it to transmit the request further.

The Proxy-Authenticate header is sent along with a 407 Proxy Authentication Required.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Proxy-Authenticate`,
	},
	{
		Name:        "Caching",
		IsTitle:     true,
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Caching",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Caching`,
	},
	{
		Name:    "Age",
		Summary: "The time, in seconds, that the object has been in a proxy cache.",
		Description: `The Age header contains the time in seconds the object has been in a proxy cache.

The Age header is usually close to zero. If it is Age: 0, it was probably just fetched from the origin server; otherwise It is usually calculated as a difference between the proxy's current date and the Date general header included in the HTTP response.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Age`,
	},
	{
		Name:    "Cache-Control",
		Summary: "Directives for caching mechanisms in both requests and responses.",
		Description: `The Cache-Control HTTP header holds directives (instructions) for caching in both requests and responses. A given directive in a request does not mean the same directive should be in the response.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cache-Control`,
	},
	{
		Name:    "Clear-Site-Data",
		Summary: "Clears browsing data (e.g. cookies, storage, cache) associated with the requesting website.",
		Description: `The Clear-Site-Data header clears browsing data (cookies, storage, cache) associated with the requesting website. It allows web developers to have more control over the data stored locally by a browser for their origins.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Clear-Site-Data`,
	},
	{
		Name:    "Expires",
		Summary: "The date/time after which the response is considered stale.",
		Description: `The Expires header contains the date/time after which the response is considered stale.

Invalid dates, like the value 0, represent a date in the past and mean that the resource is already expired.
If there is a Cache-Control header with the max-age or s-maxage directive in the response, the Expires header is ignored.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Expires`,
	},
	{
		Name:    "Pragma",
		Summary: "Implementation-specific header that may have various effects anywhere along the request-response chain. Used for backwards compatibility with HTTP/1.0 caches where the Cache-Control header is not yet present.",
		Description: `The Pragma HTTP/1.0 general header is an implementation-specific header that may have various effects along the request-response chain. It is used for backwards compatibility with HTTP/1.0 caches where the Cache-Control HTTP/1.1 header is not yet present.

Note: Pragma is not specified for HTTP responses and is therefore not a reliable replacement for the general HTTP/1.1 Cache-Control header, although it does behave the same as Cache-Control: no-cache, if the Cache-Control header field is omitted in a request. Use Pragma only for backwards compatibility with HTTP/1.0 clients.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Pragma`,
	},
	{
		Name:    "Warning",
		Summary: "General warning information about possible problems.",
		Description: `

Note: The Warning header is soon to be deprecated; see Warning (https://github.com/httpwg/http-core/issues/139) and Warning: header & stale-while-revalidate (https://github.com/whatwg/fetch/issues/913) for more details.

The Warning general HTTP header contains information about possible problems with the status of the message. More than one Warning header may appear in a response.

Warning header fields can in general be applied to any message, however some warn-codes are specific to caches and can only be applied to response messages.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Warning`,
	},
	{
		Name:        "Client hints",
		IsTitle:     true,
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Client_hints",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Client_hints`,
	},
	{
		Name:    "Accept-CH",
		Summary: "Servers can advertise support for Client Hints using the Accept-CH header field or an equivalent HTML <meta> element with http-equiv attribute ([HTML5]).",
		Description: `Secure context
This feature is available only in secure contexts (HTTPS), in some or all supporting browsers.

This is an experimental technology
Check the Browser compatibility table carefully before using this in production.

The Accept-CH header is set by the server to specify which Client Hints headers client should include in subsequent requests.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-CH`,
	},
	{
		Name:    "Accept-CH-Lifetime",
		Summary: "Servers can ask the client to remember the set of Client Hints that the server supports for a specified period of time, to enable delivery of Client Hints on subsequent requests to the server’s origin ([RFC6454]).",
		Description: `The Accept-CH-Lifetime header is set by the server to specify the persistence of Accept-CH header value that specifies for which Client Hints headers client should include in subsequent requests.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-CH-Lifetime`,
	},
	{
		Name:    "Early-Data",
		Summary: "Indicates that the request has been conveyed in early data.",
		Description: `The Early-Data header is set by an intermediate to indicate that the request has been conveyed in TLS early data, and additionally indicates that an intermediary understands the 425 (Too Early) status code.  The Early-Data header is not set by the originator of the request (i.e., a browser).

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Early-Data`,
	},
	{
		Name:        "Content-DPR",
		Summary:     "A number that indicates the ratio between physical pixels over CSS pixels of the selected image response.",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Client_hints`,
	},
	{
		Name:    "DPR",
		Summary: "A number that indicates the client’s current Device Pixel Ratio (DPR), which is the ratio of physical pixels over CSS pixels (Section 5.2 of [CSSVAL]) of the layout viewport (Section 9.1.1 of [CSS2]) on the device.",
		Description: `The DPR header is a Client Hints headers which represents the client device pixel ratio (DPR), which is the the number of physical device pixels corresponding to every CSS pixel.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/DPR`,
	},
	{
		Name:    "Device-Memory",
		Summary: "Technically a part of Device Memory API, this header represents an approximate amount of RAM client has.",
		Description: `The Device-Memory header is a Device Memory API header that works like Client Hints header which represents the approximate amount of RAM client device has.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Device-Memory`,
	},
	{
		Name:    "Save-Data",
		Summary: "A boolean that indicates the user agent's preference for reduced data usage.",
		Description: `The Save-Data header field is a boolean which, in requests, indicates the client's preference for reduced data usage. This could be for reasons such as high transfer costs, slow connection speeds, etc.

A value of On indicates explicit user opt-in into a reduced data usage mode on the client, and when communicated to origins allows them to deliver alternative content to reduce the data downloaded such as smaller image and video resources, different markup and styling, disabled polling and automatic updates, and so on.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Save-Data`,
	},
	{
		Name:        "Viewport-Width",
		Summary:     "A number that indicates the layout viewport width in CSS pixels. The provided pixel value is a number rounded to the smallest following integer (i.e. ceiling value).",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Client_hints`,
	},
	{
		Name:        "Viewport-Width",
		Summary:     "A number that indicates the layout viewport width in CSS pixels. The provided pixel value is a number rounded to the smallest following integer (i.e. ceiling value).",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Client_hints`,
	},
	{
		Name:    "Width",
		Summary: "The Width request header field is a number that indicates the desired resource width in physical pixels (i.e. intrinsic size of an image).",
		Description: `

    The Width request header field is a number that indicates the desired resource width in physical pixels (i.e. intrinsic size of an image). The provided pixel value is a number rounded to the smallest following integer (i.e. ceiling value).

    If the desired resource width is not known at the time of the request or the resource does not have a display width, the Width header field can be omitted. If Width occurs in a message more than once, the last value overrides all previous occurrences

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Client_hints`,
	},
	{
		Name:        "Conditionals",
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Conditionals",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Conditionals`,
	},
	{
		Name:    "Last-Modified",
		Summary: "The last modification date of the resource, used to compare several versions of the same resource. It is less accurate than ETag, but easier to calculate in some environments. Conditional requests using If-Modified-Since and If-Unmodified-Since use this value to change the behavior of the request.",
		Description: `The Last-Modified response HTTP header contains the date and time at which the origin server believes the resource was last modified. It is used as a validator to determine if a resource received or stored is the same. Less accurate than an ETag header, it is a fallback mechanism. Conditional requests containing If-Modified-Since or If-Unmodified-Since headers make use of this field.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Last-Modified`,
	},
	{
		Name:    "ETag",
		Summary: "A unique string identifying the version of the resource. Conditional requests using If-Match and If-None-Match use this value to change the behavior of the request.",
		Description: `The ETag HTTP response header is an identifier for a specific version of a resource. It lets caches be more efficient and save bandwidth, as a web server does not need to resend a full response if the content has not changed. Additionally, etags help prevent simultaneous updates of a resource from overwriting each other ("mid-air collisions").

If the resource at a given URL changes, a new Etag value must be generated. A comparison of them can determine whether two representations of a resource are the same. Etags are therefore similar to fingerprints, and might also be used for tracking purposes by some servers. They might also be set to persist indefinitely by a tracking server.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/ETag`,
	},
	{
		Name:    "If-Match",
		Summary: "Makes the request conditional, and applies the method only if the stored resource matches one of the given ETags.",
		Description: `The If-Match HTTP request header makes the request conditional. For GET and HEAD methods, the server will send back the requested resource only if it matches one of the listed ETags. For PUT and other non-safe methods, it will only upload the resource in this case.

The comparison with the stored ETag uses the strong comparison algorithm, meaning two files are considered identical byte to byte only. If a listed ETag has the W/ prefix indicating a weak entity tag, it will never match under this comparison algorithm.

There are two common use cases:

    For GET and HEAD methods, used in combination with a Range header, it can guarantee that the new ranges requested comes from the same resource than the previous one. If it doesn't match, then a 416 (Range Not Satisfiable) response is returned.
    For other methods, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match`,
	},
	{
		Name:    "If-None-Match",
		Summary: "Makes the request conditional, and applies the method only if the stored resource doesn't match any of the given ETags. This is used to update caches (for safe requests), or to prevent to upload a new resource when one already exists.",
		Description: `The If-None-Match HTTP request header makes the request conditional. For GET and HEAD methods, the server will send back the requested resource, with a 200 status, only if it doesn't have an ETag matching the given ones. For other methods, the request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed.

When the condition fails for GET and HEAD methods, then the server must return HTTP status code 304 (Not Modified). For methods that apply server-side changes, the status code 412 (Precondition Failed) is used. Note that the server generating a 304 response MUST generate any of the following header fields that would have been sent in a 200 (OK) response to the same request: Cache-Control, Content-Location, Date, ETag, Expires, and Vary.

The comparison with the stored ETag uses the weak comparison algorithm, meaning two files are considered identical if the content is equivalent — they don't have to be identical byte for byte. For example, two pages that differ by the date of generation in the footer would still be considered as identical.

When used in combination with If-Modified-Since, If-None-Match has precedence (if the server supports it).

There are two common use cases:

    For GET and HEAD methods, to update a cached entity that has an associated ETag.
    For other methods, and in particular for PUT, If-None-Match used with the * value can be used to save a file not known to exist, guaranteeing that another upload didn't happen before, losing the data of the previous put; this problem is a variation of the lost update problem.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-None-Match`,
	},
	{
		Name:    "If-Modified-Since",
		Summary: "Makes the request conditional, and expects the entity to be transmitted only if it has been modified after the given date. This is used to transmit data only when the cache is out of date.",
		Description: `The If-Modified-Since request HTTP header makes the request conditional: the server will send back the requested resource, with a 200 status, only if it has been last modified after the given date. If the request has not been modified since, the response will be a 304 without any body; the Last-Modified response header of a previous request will contain the date of last modification. Unlike If-Unmodified-Since, If-Modified-Since can only be used with a GET or HEAD.

When used in combination with If-None-Match, it is ignored, unless the server doesn't support If-None-Match.

The most common use case is to update a cached entity that has no associated ETag.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Modified-Since`,
	},
	{
		Name:    "If-Unmodified-Since",
		Summary: "Makes the request conditional, and expects the entity to be transmitted only if it has not been modified after the given date. This ensures the coherence of a new fragment of a specific range with previous ones, or to implement an optimistic concurrency control system when modifying existing documents.",
		Description: `The If-Unmodified-Since request HTTP header makes the request conditional: the server will send back the requested resource, or accept it in the case of a POST or another non-safe method, only if it has not been last modified after the given date. If the resource has been modified after the given date, the response will be a 412 (Precondition Failed) error.

There are two common use cases:

    In conjunction with non-safe methods, like POST, it can be used to implement an optimistic concurrency control, like done by some wikis: editions are rejected if the stored document has been modified since the original has been retrieved.
    In conjunction with a range request with a If-Range header, it can be used to ensure that the new fragment requested comes from an unmodified document.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Unmodified-Since`,
	},
	{
		Name:    "Vary",
		Summary: "Determines how to match request headers to decide whether a cached response can be used rather than requesting a fresh one from the origin server.",
		Description: `The Vary HTTP response header determines how to match future request headers to decide whether a cached response can be used rather than requesting a fresh one from the origin server. It is used by the server to indicate which headers it used when selecting a representation of a resource in a content negotiation algorithm.

The Vary header should be set on a 304 Not Modified response exactly like it would have been set on an equivalent 200 OK response.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Vary`,
	},
	{
		Name:        "Connection management",
		IsTitle:     true,
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Connection_management",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Connection_management`,
	},
	{
		Name:    "Connection",
		Summary: "Controls whether the network connection stays open after the current transaction finishes.",
		Description: `The Connection general header controls whether or not the network connection stays open after the current transaction finishes. If the value sent is keep-alive, the connection is persistent and not closed, allowing for subsequent requests to the same server to be done.
Connection-specific header fields such as Connection must not be used with HTTP/2.

Except for the standard hop-by-hop headers (Keep-Alive, Transfer-Encoding, TE, Connection, Trailer, Upgrade, Proxy-Authorization and Proxy-Authenticate), any hop-by-hop headers used by the message must be listed in the Connection header, so that the first proxy knows it has to consume them and not forward them further. Standard hop-by-hop headers can be listed too (it is often the case of Keep-Alive, but this is not mandatory).

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Connection`,
	},
	{
		Name:    "Keep-Alive",
		Summary: "Controls how long a persistent connection should stay open.",
		Description: `The Keep-Alive general header allows the sender to hint about how the connection may be used to set a timeout and a maximum amount of requests.

The Connection header needs to be set to "keep-alive" for this header to have any meaning. Also, Connection and Keep-Alive are ignored in HTTP/2; connection management is handled by other mechanisms there.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Keep-Alive`,
	},
	{
		Name:    "Content negotiation",
		IsTitle: true,
		Summary: "https://developer.mozilla.org/en-US/docs/Web/HTTP/Content_negotiation",
		Description: `In HTTP, content negotiation is the mechanism that is used for serving different representations of a resource at the same URI, so that the user agent can specify which is best suited for the user (for example, which language of a document, which image format, or which content encoding).

https://developer.mozilla.org/en-US/docs/Web/HTTP/Content_negotiation`,
	},
	{
		Name:    "Accept",
		Summary: "Informs the server about the types of data that can be sent back.",
		Description: `The Accept request HTTP header advertises which content types, expressed as MIME types, the client is able to understand. Using content negotiation, the server then selects one of the proposals, uses it and informs the client of its choice with the Content-Type response header. Browsers set adequate values for this header depending on the context where the request is done: when fetching a CSS stylesheet a different value is set for the request than when fetching an image, video or a script.


https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept`,
	},
	{
		Name:    "Accept-Charset",
		Summary: "Which character encodings the client understands.",
		Description: `The Accept-Charset request HTTP header advertises which character encodings the client understands. Using content negotiation, the server selects one of the encodings, uses it, and informs the client of its choice within the Content-Type response header, usually in a charset= parameter. Browsers usually don't send this header, as the default value for each resource is usually correct and transmitting it would allow fingerprinting.

If the server cannot serve any character encoding from this request header, it can theoretically send back a 406 Not Acceptable error code. But for a better user experience, this is rarely done and the Accept-Charset header is ignored.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Charset`,
	},
	{
		Name:    "Accept-Encoding",
		Summary: "The encoding algorithm, usually a compression algorithm, that can be used on the resource sent back.",
		Description: `The Accept-Encoding request HTTP header advertises which content encoding, usually a compression algorithm, the client is able to understand. Using content negotiation, the server selects one of the proposals, uses it and informs the client of its choice with the Content-Encoding response header.

Even if both the client and the server supports the same compression algorithms, the server may choose not to compress the body of a response, if the identity value is also acceptable. Two common cases lead to this:

    The data to be sent is already compressed and a second compression won't lead to smaller data to be transmitted. This may be the case with some image formats;
    The server is overloaded and cannot afford the computational overhead induced by the compression requirement. Typically, Microsoft recommends not to compress if a server uses more than 80% of its computational power.

As long as the identity value, meaning no encoding, is not explicitly forbidden, by an identity;q=0 or a *;q=0 without another explicitly set value for identity, the server must never send back a 406 Not Acceptable error.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Encoding`,
	},
	{
		Name:    "Accept-Language",
		Summary: "Informs the server about the human language the server is expected to send back. This is a hint and is not necessarily under the full control of the user: the server should always pay attention not to override an explicit user choice (like selecting a language from a dropdown).",
		Description: `The Accept-Language request HTTP header advertises which languages the client is able to understand, and which locale variant is preferred. (By languages, we mean natural languages, such as English, and not programming languages.) Using content negotiation, the server then selects one of the proposals, uses it and informs the client of its choice with the Content-Language response header. Browsers set adequate values for this header according to their user interface language and even if a user can change it, this happens rarely (and is frowned upon as it leads to fingerprinting).

This header is a hint to be used when the server has no way of determining the language via another way, like a specific URL, that is controlled by an explicit user decision. It is recommended that the server never overrides an explicit decision. The content of the Accept-Language is often out of the control of the user (like when traveling and using an Internet Cafe in a different country); the user may also want to visit a page in another language than the locale of their user interface.

If the server cannot serve any matching language, it can theoretically send back a 406 (Not Acceptable) error code. But, for a better user experience, this is rarely done and more common way is to ignore the Accept-Language header in this case.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Language`,
	},
	{
		Name:        "Controls",
		IsTitle:     true,
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Controls",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Controls`,
	},
	{
		Name:    "Expect",
		Summary: "Indicates expectations that need to be fulfilled by the server to properly handle the request.",
		Description: `The Expect HTTP request header indicates expectations that need to be fulfilled by the server in order to properly handle the request.

The only expectation defined in the specification is Expect: 100-continue, to which the server shall respond with:

    100 if the information contained in the header is sufficient to cause an immediate success,
    417 (Expectation Failed) if it cannot meet the expectation; or any other 4xx status otherwise.

For example, the server may reject a request if its Content-Length is too large.

No common browsers send the Expect header, but some other clients such as cURL do so by default.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Expect`,
	},
	{
		Name:        "Max-Forwards",
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Controls",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Controls`,
	},
	{
		Name:        "Cookies",
		IsTitle:     true,
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Cookies",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Cookies`,
	},
	{
		Name:    "Cookie",
		Summary: "Contains stored HTTP cookies previously sent by the server with the Set-Cookie header.",
		Description: `The Cookie HTTP request header contains stored HTTP cookies previously sent by the server with the Set-Cookie header.

The Cookie header is optional and may be omitted if, for example, the browser's privacy settings block cookies.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cookie`,
	},
	{
		Name:    "Set-Cookie",
		Summary: "Send cookies from the server to the user-agent.",
		Description: `The Set-Cookie HTTP response header is used to send cookies from the server to the user agent, so the user agent can send them back to the server later.

For more information, see the guide on HTTP cookies.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie`,
	},
	{
		Name:    "CORS",
		IsTitle: true,
		Summary: "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#CORS",
		Description: `Cross-Origin Resource Sharing (CORS) is a mechanism that uses additional HTTP headers to tell browsers to give a web application running at one origin, access to selected resources from a different origin. A web application executes a cross-origin HTTP request when it requests a resource that has a different origin (domain, protocol, or port) from its own.

An example of a cross-origin request: the front-end JavaScript code served from https://domain-a.com uses XMLHttpRequest to make a request for https://domain-b.com/data.json.

For security reasons, browsers restrict cross-origin HTTP requests initiated from scripts. For example, XMLHttpRequest and the Fetch API follow the same-origin policy. This means that a web application using those APIs can only request resources from the same origin the application was loaded from, unless the response from other origins includes the right CORS headers.

https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS`,
	},
	{
		Name:    "Access-Control-Allow-Origin",
		Summary: "Indicates whether the response can be shared.",
		Description: `The Access-Control-Allow-Origin response header indicates whether the response can be shared with requesting code from the given origin.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin`,
	},
	{
		Name:    "Access-Control-Allow-Credentials",
		Summary: "Indicates whether the response to the request can be exposed when the credentials flag is true.",
		Description: `The Access-Control-Allow-Credentials response header tells browsers whether to expose the response to frontend JavaScript code when the request's credentials mode (Request.credentials) is include.

When a request's credentials mode (Request.credentials) is include, browsers will only expose the response to frontend JavaScript code if the Access-Control-Allow-Credentials value is true.

Credentials are cookies, authorization headers or TLS client certificates.

When used as part of a response to a preflight request, this indicates whether or not the actual request can be made using credentials. Note that simple GET requests are not preflighted, and so if a request is made for a resource with credentials, if this header is not returned with the resource, the response is ignored by the browser and not returned to web content.

The Access-Control-Allow-Credentials header works in conjunction with the XMLHttpRequest.withCredentials property or with the credentials option in the Request() constructor of the Fetch API. For a CORS request with credentials, in order for browsers to expose the response to frontend JavaScript code, both the server (using the Access-Control-Allow-Credentials header) and the client (by setting the credentials mode for the XHR, Fetch, or Ajax request) must indicate that they’re opting in to including credentials.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Credentials`,
	},
	{
		Name:    "Access-Control-Allow-Headers",
		Summary: "Used in response to a preflight request to indicate which HTTP headers can be used when making the actual request.",
		Description: `The Access-Control-Allow-Headers response header is used in response to a preflight request which includes the Access-Control-Request-Headers to indicate which HTTP headers can be used during the actual request.

This header is required if the request has an Access-Control-Request-Headers header.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Headers`,
	},
	{
		Name:    "Access-Control-Allow-Methods",
		Summary: "Specifies the methods allowed when accessing the resource in response to a preflight request.",
		Description: `The Access-Control-Allow-Methods response header specifies the method or methods allowed when accessing the resource in response to a preflight request.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Methods`,
	},
	{
		Name:    "Access-Control-Expose-Headers",
		Summary: "Indicates which headers can be exposed as part of the response by listing their names.",
		Description: `The Access-Control-Expose-Headers response header indicates which headers can be exposed as part of the response by listing their names.

By default, only the 6 CORS-safelisted response headers are exposed:

    Cache-Control
    Content-Language
    Content-Type
    Expires
    Last-Modified
    Pragma

If you want clients to be able to access other headers, you have to list them using the Access-Control-Expose-Headers header.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Expose-Headers`,
	},
	{
		Name:    "Access-Control-Max-Age",
		Summary: "Indicates how long the results of a preflight request can be cached.",
		Description: `The Access-Control-Max-Age response header indicates how long the results of a preflight request (that is the information contained in the Access-Control-Allow-Methods and Access-Control-Allow-Headers headers) can be cached.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Max-Age`,
	},
	{
		Name:    "Access-Control-Request-Headers",
		Summary: "Used when issuing a preflight request to let the server know which HTTP headers will be used when the actual request is made.",
		Description: `The Access-Control-Request-Headers request header is used by browsers when issuing a preflight request, to let the server know which HTTP headers the client might send when the actual request is made.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Request-Headers`,
	},
	{
		Name:    "Origin",
		Summary: "Indicates where a fetch originates from.",
		Description: `The Origin request header indicates where a fetch originates from. It doesn't include any path information, but only the server name. It is sent with CORS requests, as well as with POST requests. It is similar to the Referer header, but, unlike this header, it doesn't disclose the whole path.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Origin`,
	},
	{
		Name:    "Timing-Allow-Origin",
		Summary: "Specifies origins that are allowed to see values of attributes retrieved via features of the Resource Timing API, which would otherwise be reported as zero due to cross-origin restrictions.",
		Description: `The Timing-Allow-Origin response header specifies origins that are allowed to see values of attributes retrieved via features of the Resource Timing API, which would otherwise be reported as zero due to cross-origin restrictions.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Timing-Allow-Origin`,
	},
	{
		Name:        "Do Not Track",
		IsTitle:     true,
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Do_Not_Track",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Do_Not_Track`,
	},
	{
		Name:    "DNT",
		Summary: "Expresses the user's tracking preference.",
		Description: `The DNT (Do Not Track) request header indicates the user's tracking preference. It lets users indicate whether they would prefer privacy rather than personalized content.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/DNT`,
	},
	{
		Name:    "Tk",
		Summary: "Indicates the tracking status of the corresponding response.",
		Description: `The Tk response header indicates the tracking status that applied to the corresponding request.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Tk`,
	},
	{
		Name:        "Downloads",
		IsTitle:     true,
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Downloads",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Downloads`,
	},
	{
		Name:    "Content-Disposition",
		Summary: "Indicates if the resource transmitted should be displayed inline (default behavior without the header), or if it should be handled like a download and the browser should present a “Save As” dialog.",
		Description: `In a regular HTTP response, the Content-Disposition response header is a header indicating if the content is expected to be displayed inline in the browser, that is, as a Web page or as part of a Web page, or as an attachment, that is downloaded and saved locally.

In a multipart/form-data body, the HTTP Content-Disposition general header is a header that can be used on the subpart of a multipart body to give information about the field it applies to. The subpart is delimited by the boundary defined in the Content-Type header. Used on the body itself, Content-Disposition has no effect.

The Content-Disposition header is defined in the larger context of MIME messages for e-mail, but only a subset of the possible parameters apply to HTTP forms and POST requests. Only the value form-data, as well as the optional directive name and filename, can be used in the HTTP context.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition`,
	},
	{
		Name:        "Message body information",
		IsTitle:     true,
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Message_body_information",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Message_body_information`,
	},
	{
		Name:    "Content-Length",
		Summary: "The size of the resource, in decimal number of bytes.",
		Description: `The Content-Length entity header indicates the size of the entity-body, in bytes, sent to the recipient.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Length`,
	},
	{
		Name:    "Content-Type",
		Summary: "Indicates the media type of the resource.",
		Description: `The Content-Type entity header is used to indicate the media type of the resource.

In responses, a Content-Type header tells the client what the content type of the returned content actually is. Browsers will do MIME sniffing in some cases and will not necessarily follow the value of this header; to prevent this behavior, the header X-Content-Type-Options can be set to nosniff.

In requests, (such as POST or PUT), the client tells the server what type of data is actually sent.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Type`,
	},
	{
		Name:    "Content-Encoding",
		Summary: "Used to specify the compression algorithm.",
		Description: `The Content-Encoding entity header is used to compress the media-type. When present, its value indicates which encodings were applied to the entity-body. It lets the client know how to decode in order to obtain the media-type referenced by the Content-Type header.

The recommendation is to compress data as much as possible and therefore to use this field, but some types of resources, such as jpeg images, are already compressed. Sometimes, using additional compression doesn't reduce payload size and can even make the payload longer.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Encoding`,
	},
	{
		Name:    "Content-Language",
		Summary: "Describes the human language(s) intended for the audience, so that it allows a user to differentiate according to the users' own preferred language.",
		Description: `The Content-Language entity header is used to describe the language(s) intended for the audience, so that it allows a user to differentiate according to the users' own preferred language.

For example, if "Content-Language: de-DE" is set, it says that the document is intended for German language speakers (however, it doesn't indicate the document is written in German. For example, it might be written in English as part of a language course for German speakers. If you want to indicate which language the document is written in, use the lang attribute instead).

If no Content-Language is specified, the default is that the content is intended for all language audiences. Multiple language tags are also possible, as well as applying the Content-Language header to various media types and not only to textual documents.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Language`,
	},
	{
		Name:    "Content-Location",
		Summary: "Indicates an alternate location for the returned data.",
		Description: `The Content-Location header indicates an alternate location for the returned data. The principal use is to indicate the URL of a resource transmitted as the result of content negotiation.

Location and Content-Location are different. Location indicates the URL of a redirect, while Content-Location indicates the direct URL to use to access the resource, without further content negotiation in the future. Location is a header associated with the response, while Content-Location is associated with the data returned. This distinction may seem abstract without examples.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Location`,
	},
	{
		Name:        "Proxies",
		IsTitle:     true,
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Proxies",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Proxies`,
	},
	{
		Name:    "Forwarded",
		Summary: "Contains information from the client-facing side of proxy servers that is altered or lost when a proxy is involved in the path of the request.",
		Description: `The Forwarded header contains information from the client-facing side of proxy servers that is altered or lost when a proxy is involved in the path of the request.

The alternative and de-facto standard versions of this header are the X-Forwarded-For, X-Forwarded-Host and X-Forwarded-Proto headers.

This header is used for debugging, statistics, and generating location-dependent content and by design it exposes privacy sensitive information, such as the IP address of the client. Therefore the user's privacy must be kept in mind when deploying this header.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Forwarded`,
	},
	{
		Name:    "X-Forwarded-For",
		Summary: "Identifies the originating IP addresses of a client connecting to a web server through an HTTP proxy or a load balancer.",
		Description: `The X-Forwarded-For (XFF) header is a de-facto standard header for identifying the originating IP address of a client connecting to a web server through an HTTP proxy or a load balancer. When traffic is intercepted between clients and servers, server access logs contain the IP address of the proxy or load balancer only. To see the original IP address of the client, the X-Forwarded-For request header is used.

This header is used for debugging, statistics, and generating location-dependent content and by design it exposes privacy sensitive information, such as the IP address of the client. Therefore the user's privacy must be kept in mind when deploying this header.

A standardized version of this header is the HTTP Forwarded header.

X-Forwarded-For is also an email-header indicating that an email-message was forwarded from another account.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-For`,
	},
	{
		Name:    "X-Forwarded-Host",
		Summary: "Identifies the original host requested that a client used to connect to your proxy or load balancer.",
		Description: `The X-Forwarded-Host (XFH) header is a de-facto standard header for identifying the original host requested by the client in the Host HTTP request header.

Host names and ports of reverse proxies (load balancers, CDNs) may differ from the origin server handling the request, in that case the X-Forwarded-Host header is useful to determine which Host was originally used.

This header is used for debugging, statistics, and generating location-dependent content and by design it exposes privacy sensitive information, such as the IP address of the client. Therefore the user's privacy must be kept in mind when deploying this header.

A standardized version of this header is the HTTP Forwarded header.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-Host`,
	},
	{
		Name:    "X-Forwarded-Proto",
		Summary: "Identifies the protocol (HTTP or HTTPS) that a client used to connect to your proxy or load balancer.",
		Description: `The X-Forwarded-Proto (XFP) header is a de-facto standard header for identifying the protocol (HTTP or HTTPS) that a client used to connect to your proxy or load balancer. Your server access logs contain the protocol used between the server and the load balancer, but not the protocol used between the client and the load balancer. To determine the protocol used between the client and the load balancer, the X-Forwarded-Proto request header can be used.

A standardized version of this header is the HTTP Forwarded header.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-Proto`,
	},
	{
		Name:    "Via",
		Summary: "Added by proxies, both forward and reverse proxies, and can appear in the request headers and the response headers.",
		Description: `The Via general header is added by proxies, both forward and reverse proxies, and can appear in the request headers and the response headers. It is used for tracking message forwards, avoiding request loops, and identifying the protocol capabilities of senders along the request/response chain.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Via`,
	},
	{
		Name:        "Redirects",
		IsTitle:     true,
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Redirects",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Redirects`,
	},
	{
		Name:    "Location",
		Summary: "Indicates the URL to redirect a page to.",
		Description: `The Location response header indicates the URL to redirect a page to. It only provides a meaning when served with a 3xx (redirection) or 201 (created) status response.

In cases of redirection, the HTTP method used to make the new request to fetch the page pointed to by Location depends of the original method and of the kind of redirection:

    If 303 (See Also) responses always lead to the use of a GET method, 307 (Temporary Redirect) and 308 (Permanent Redirect) don't change the method used in the original request;
    301 (Permanent Redirect) and 302 (Found) doesn't change the method most of the time, though older user-agents may (so you basically don't know).

All responses with one of these status codes send a Location header.

In cases of resource creation, it indicates the URL to the newly created resource.

Location and Content-Location are different: Location indicates the target of a redirection (or the URL of a newly created resource), while Content-Location indicates the direct URL to use to access the resource when content negotiation happened, without the need of further content negotiation. Location is a header associated with the response, while Content-Location is associated with the entity returned.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Location`,
	},
	{
		Name:        "Request context",
		IsTitle:     true,
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Request_context",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Request_context`,
	},
	{
		Name:    "From",
		Summary: "Contains an Internet email address for a human user who controls the requesting user agent.",
		Description: `The From request header contains an Internet email address for a human user who controls the requesting user agent.

If you are running a robotic user agent (e.g. a crawler), the From header should be sent, so you can be contacted if problems occur on servers, such as if the robot is sending excessive, unwanted, or invalid requests.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/From`,
	},
	{
		Name:    "Host",
		Summary: "Specifies the domain name of the server (for virtual hosting), and (optionally) the TCP port number on which the server is listening.",
		Description: `The Host request header specifies the domain name of the server (for virtual hosting), and (optionally) the TCP port number on which the server is listening.

If no port is given, the default port for the service requested (e.g., "80" for an HTTP URL) is implied.

A Host header field must be sent in all HTTP/1.1 request messages. A 400 (Bad Request) status code will be sent to any HTTP/1.1 request message that lacks a Host header field or contains more than one.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Host`,
	},
	{
		Name:    "Referer",
		Summary: "The address of the previous web page from which a link to the currently requested page was followed.",
		Description: `The Referer request header contains the address of the previous web page from which a link to the currently requested page was followed. The Referer header allows servers to identify where people are visiting them from and may use that data for analytics, logging, or optimized caching, for example.

Important: Although this header has many innocent uses it can have undesirable consequences for user security and privacy. See Referer header: privacy and security concerns for more information and mitigations.

Note that referer is actually a misspelling of the word "referrer". See HTTP referer on Wikipedia for more details.

A Referer header is not sent by browsers if:

    The referring resource is a local "file" or "data" URI.
    An unsecured HTTP request is used and the referring page was received with a secure protocol (HTTPS).

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referer`,
	},
	{
		Name:    "Referrer",
		Summary: "See the Referer header",
		Description: `This header was spelt incorrectly in the original implementation. See Referer for further details.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referer`,
	},
	{
		Name:    "Referrer-Policy",
		Summary: "Governs which referrer information sent in the Referer header should be included with requests made.",
		Description: `The Referrer-Policy HTTP header controls how much referrer information (sent via the Referer header) should be included with requests.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy`,
	},
	{
		Name:    "User-Agent",
		Summary: "Contains a characteristic string that allows the network protocol peers to identify the application type, operating system, software vendor or software version of the requesting software user agent. See also the Firefox user agent string reference.",
		Description: `The User-Agent request header is a characteristic string that lets servers and network peers identify the application, operating system, vendor, and/or version of the requesting user agent.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/User-Agent`,
	},
	{
		Name:        "Response context",
		IsTitle:     true,
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Response_context",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Response_context`,
	},
	{
		Name:    "Allow",
		Summary: "Lists the set of HTTP request methods support by a resource.",
		Description: `The Allow header lists the set of methods supported by a resource.

This header must be sent if the server responds with a 405 Method Not Allowed status code to indicate which request methods can be used. An empty Allow header indicates that the resource allows no request methods, which might occur temporarily for a given resource, for example.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Allow`,
	},
	{
		Name:    "Server",
		Summary: "Contains information about the software used by the origin server to handle the request.",
		Description: `The Server header contains information about the software used by the origin server to handle the request.

Overly long and detailed Server values should be avoided as they potentially reveal internal implementation details that might make it (slightly) easier for attackers to find and exploit known security holes.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Server`,
	},
	{
		Name:        "Range requests",
		IsTitle:     true,
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Range_requests",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Range_requests`,
	},
	{
		Name:    "Accept-Ranges",
		Summary: "Indicates if the server supports range requests, and if so in which unit the range can be expressed.",
		Description: `The Accept-Ranges response HTTP header is a marker used by the server to advertise its support of partial requests. The value of this field indicates the unit that can be used to define a range.

In presence of an Accept-Ranges header, the browser may try to resume an interrupted download, rather than to start it from the start again.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Ranges`,
	},
	{
		Name:    "Range",
		Summary: "Indicates the part of a document that the server should return.",
		Description: `The Range HTTP request header indicates the part of a document that the server should return. Several parts can be requested with one Range header at once, and the server may send back these ranges in a multipart document. If the server sends back ranges, it uses the 206 Partial Content for the response. If the ranges are invalid, the server returns the 416 Range Not Satisfiable error. The server can also ignore the Range header and return the whole document with a 200 status code.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Range`,
	},
	{
		Name:    "If-Range",
		Summary: "Creates a conditional range request that is only fulfilled if the given etag or date matches the remote resource. Used to prevent downloading two ranges from incompatible version of the resource.",
		Description: `The If-Range HTTP request header makes a range request conditional: if the condition is fulfilled, the range request will be issued and the server sends back a 206 Partial Content answer with the appropriate body. If the condition is not fulfilled, the full resource is sent back, with a 200 OK status.

This header can be used either with a Last-Modified validator, or with an ETag, but not with both.

The most common use case is to resume a download, to guarantee that the stored resource has not been modified since the last fragment has been received.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Range`,
	},
	{
		Name:    "Content-Range",
		Summary: "Indicates where in a full body message a partial message belongs.",
		Description: `The Content-Range response HTTP header indicates where in a full body message a partial message belongs.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Range`,
	},
	{
		Name:        "Security",
		IsTitle:     true,
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Security",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Security`,
	},
	{
		Name:        "Cross-Origin-Embedder-Policy (COEP)",
		Summary:     "Allows a server to declare an embedder policy for a given document.",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Security`,
	},
	{
		Name:        "Cross-Origin-Opener-Policy (COOP)",
		Summary:     "Prevents other domains from opening/controlling a window.",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Security`,
	},
	{
		Name:    "Cross-Origin-Resource-Policy (CORP)",
		Summary: "Prevents other domains from reading the response of the resources to which this header is applied.",
		Description: `The HTTP Cross-Origin-Resource-Policy response header conveys a desire that the browser blocks no-cors cross-origin/cross-site requests to the given resource.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cross-Origin-Resource-Policy`,
	},
	{
		Name:    "Content-Security-Policy (CSP)",
		Summary: "Controls resources the user agent is allowed to load for a given page.",
		Description: `The HTTP Content-Security-Policy response header allows web site administrators to control resources the user agent is allowed to load for a given page. With a few exceptions, policies mostly involve specifying server origins and script endpoints. This helps guard against cross-site scripting attacks (XSS).

For more information, see the introductory article on Content Security Policy (CSP).

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy`,
	},
	{
		Name:    "Content-Security-Policy-Report-Only",
		Summary: "Allows web developers to experiment with policies by monitoring, but not enforcing, their effects. These violation reports consist of JSON documents sent via an HTTP POST request to the specified URI.",
		Description: `The HTTP Content-Security-Policy-Report-Only response header allows web developers to experiment with policies by monitoring (but not enforcing) their effects. These violation reports consist of JSON documents sent via an HTTP POST request to the specified URI.

For more information, see also this article on Content Security Policy (CSP).

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy-Report-Only`,
	},
	{
		Name:    "Expect-CT",
		Summary: "Allows sites to opt in to reporting and/or enforcement of Certificate Transparency requirements, which prevents the use of misissued certificates for that site from going unnoticed. When a site enables the Expect-CT header, they are requesting that Chrome check that any certificate for that site appears in public CT logs.",
		Description: `The Expect-CT header allows sites to opt in to reporting and/or enforcement of Certificate Transparency requirements, which prevents the use of misissued certificates for that site from going unnoticed.

CT requirements can be satisfied by servers via any one of the following mechanisms:

    X.509v3 certificate extension to allow embedding of signed certificate timestamps issued by individual logs
    A TLS extension of type signed_certificate_timestamp sent during the handshake
    Supporting OCSP stapling (that is, the status_request TLS extension) and providing a SignedCertificateTimestampList

When a site enables the Expect-CT header, they are requesting that the browser check that any certificate for that site appears in public CT logs.

Browsers ignore the Expect-CT header when sent over HTTP, the header only has effect on HTTPS connections.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Expect-CT`,
	},
	{
		Name:    "Feature-Policy",
		Summary: "Provides a mechanism to allow and deny the use of browser features in its own frame, and in iframes that it embeds.",
		Description: `The HTTP Feature-Policy header provides a mechanism to allow and deny the use of browser features in its own frame, and in content within any <iframe> elements in the document.

For more information, see the main Feature Policy article.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Feature-Policy`,
	},
	{
		Name:    "Public-Key-Pins (HPKP)",
		Summary: "Associates a specific cryptographic public key with a certain web server to decrease the risk of MITM attacks with forged certificates.",
		Description: `Note: Public Key Pinning mechanism was deprecated in favor of Certificate Transparency and Expect-CT header.

The HTTP Public-Key-Pins response header associates a specific cryptographic public key with a certain web server to decrease the risk of MITM attacks with forged certificates. If one or several keys are pinned and none of them are used by the server, the browser will not accept the response as legitimate, and will not display it.

For more information, see the HTTP Public Key Pinning article.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Public-Key-Pins`,
	},
	{
		Name:    "Public-Key-Pins-Report-Only",
		Summary: "Sends reports to the report-uri specified in the header and does still allow clients to connect to the server even if the pinning is violated.",
		Description: `The HTTP Public-Key-Pins-Report-Only response header sends reports of pinning violation to the report-uri specified in the header but, unlike Public-Key-Pins still allows browsers to connect to the server if the pinning is violated.

For more information, see the Public-Key-Pins header reference page and the HTTP Public Key Pinning article.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Public-Key-Pins-Report-Only`,
	},
	{
		Name:    "Strict-Transport-Security (HSTS)",
		Summary: "Force communication using HTTPS instead of HTTP.",
		Description: `The HTTP Strict-Transport-Security response header (often abbreviated as HSTS) lets a web site tell browsers that it should only be accessed using HTTPS, instead of using HTTP.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security`,
	},
	{
		Name:    "Upgrade-Insecure-Requests",
		Summary: "Sends a signal to the server expressing the client’s preference for an encrypted and authenticated response, and that it can successfully handle the upgrade-insecure-requests directive.",
		Description: `The HTTP Upgrade-Insecure-Requests request header sends a signal to the server expressing the client’s preference for an encrypted and authenticated response, and that it can successfully handle the upgrade-insecure-requests CSP directive.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Upgrade-Insecure-Requests`,
	},
	{
		Name:    "X-Content-Type-Options",
		Summary: "Disables MIME sniffing and forces browser to use the type given in Content-Type.",
		Description: `The X-Content-Type-Options response HTTP header is a marker used by the server to indicate that the MIME types advertised in the Content-Type headers should not be changed and be followed. This allows to opt-out of MIME type sniffing, or, in other words, it is a way to say that the webmasters knew what they were doing.

This header was introduced by Microsoft in IE 8 as a way for webmasters to block content sniffing that was happening and could transform non-executable MIME types into executable MIME types. Since then, other browsers have introduced it, even if their MIME sniffing algorithms were less aggressive.

Starting with Firefox 72, the opting out of MIME sniffing is also applied to top-level documents if a Content-type is provided. This can cause HTML web pages to be downloaded instead of being rendered when they are served with a MIME type other than text/html. Make sure to set both headers correctly.

Site security testers usually expect this header to be set.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options`,
	},
	{
		Name:    "X-Content-Type-Options",
		Summary: "Disables MIME sniffing and forces browser to use the type given in Content-Type.",
		Description: `The X-Content-Type-Options response HTTP header is a marker used by the server to indicate that the MIME types advertised in the Content-Type headers should not be changed and be followed. This allows to opt-out of MIME type sniffing, or, in other words, it is a way to say that the webmasters knew what they were doing.

This header was introduced by Microsoft in IE 8 as a way for webmasters to block content sniffing that was happening and could transform non-executable MIME types into executable MIME types. Since then, other browsers have introduced it, even if their MIME sniffing algorithms were less aggressive.

Starting with Firefox 72, the opting out of MIME sniffing is also applied to top-level documents if a Content-type is provided. This can cause HTML web pages to be downloaded instead of being rendered when they are served with a MIME type other than text/html. Make sure to set both headers correctly.

Site security testers usually expect this header to be set.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options`,
	},
	{
		Name:        "X-Download-Options",
		Summary:     "The X-Download-Options HTTP header indicates that the browser (Internet Explorer) should not display the option to \"Open\" a file that has been downloaded from an application, to prevent phishing attacks as the file otherwise would gain access to execute in the context of the application. (Note: related MS Edge bug).",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Security`,
	},
	{
		Name:    "X-Frame-Options (XFO)",
		Summary: "Indicates whether a browser should be allowed to render a page in a <frame>, <iframe>, <embed> or <object>.",
		Description: `The X-Frame-Options HTTP response header can be used to indicate whether or not a browser should be allowed to render a page in a <frame>, <iframe>, <embed> or <object>. Sites can use this to avoid clickjacking attacks, by ensuring that their content is not embedded into other sites.

The added security is only provided if the user accessing the document is using a browser supporting X-Frame-Options.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options`,
	},
	{
		Name:        "X-Permitted-Cross-Domain-Policies",
		Summary:     "Specifies if a cross-domain policy file (crossdomain.xml) is allowed. The file may define a policy to grant clients, such as Adobe's Flash Player, Adobe Acrobat, Microsoft Silverlight, or Apache Flex, permission to handle data across domains that would otherwise be restricted due to the Same-Origin Policy. See the Cross-domain Policy File Specification for more information.",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Security`,
	},
	{
		Name:        "X-Powered-By",
		Summary:     "May be set by hosting environments or other frameworks and contains information about them while not providing any usefulness to the application or its visitors. Unset this header to avoid exposing potential vulnerabilities.",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Security`,
	},
	{
		Name:    "X-XSS-Protection",
		Summary: "Enables cross-site scripting filtering.",
		Description: `The HTTP X-XSS-Protection response header is a feature of Internet Explorer, Chrome and Safari that stops pages from loading when they detect reflected cross-site scripting (XSS) attacks. Although these protections are largely unnecessary in modern browsers when sites implement a strong Content-Security-Policy that disables the use of inline JavaScript ('unsafe-inline'), they can still provide protections for users of older web browsers that don't yet support CSP.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection`,
	},
	{
		Name:        "Server-sent events",
		IsTitle:     true,
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Server-sent_events",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Server-sent_events`,
	},
	{
		Name:        "Last-Event-ID",
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Server-sent_events",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Server-sent_events`,
	},
	{
		Name:        "NEL",
		Summary:     "Defines a mechanism that enables developers to declare a network error reporting policy.",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Server-sent_events`,
	},
	{
		Name:        "Ping-From",
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Server-sent_events",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Server-sent_events`,
	},
	{
		Name:        "Ping-To",
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Server-sent_events",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Server-sent_events`,
	},
	{
		Name:        "Report-To",
		Summary:     "Used to specify a server endpoint for the browser to send warning and error reports to.",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Server-sent_events`,
	},
	{
		Name:        "Transfer coding",
		IsTitle:     true,
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Transfer_coding",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Transfer_coding`,
	},
	{
		Name:    "Transfer-Encoding",
		Summary: "Specifies the form of encoding used to safely transfer the entity to the user.",
		Description: `The Transfer-Encoding header specifies the form of encoding used to safely transfer the payload body to the user.
HTTP/2 doesn't support HTTP 1.1's chunked transfer encoding mechanism, as it provides its own, more efficient, mechanisms for data streaming.

Transfer-Encoding is a hop-by-hop header, that is applied to a message between two nodes, not to a resource itself. Each segment of a multi-node connection can use different Transfer-Encoding values. If you want to compress data over the whole connection, use the end-to-end Content-Encoding header instead.

When present on a response to a HEAD request that has no body, it indicates the value that would have applied to the corresponding GET message.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Transfer-Encoding`,
	},
	{
		Name:    "TE",
		Summary: "Specifies the transfer encodings the user agent is willing to accept.",
		Description: `The TE request header specifies the transfer encodings the user agent is willing to accept. (you could informally call it Accept-Transfer-Encoding, which would be more intuitive).
In HTTP/2 - the TE header field is only accepted if the trailers value is set.

See also the Transfer-Encoding response header for more details on transfer encodings. Note that chunked is always acceptable for HTTP/1.1 recipients and you don't have to specify "chunked" using the TE header. However, it is useful for setting if the client is accepting trailer fields in a chunked transfer coding using the "trailers" value.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/TE`,
	},
	{
		Name:    "Trailer",
		Summary: "Allows the sender to include additional fields at the end of chunked message.",
		Description: `The Trailer response header allows the sender to include additional fields at the end of chunked messages in order to supply metadata that might be dynamically generated while the message body is sent, such as a message integrity check, digital signature, or post-processing status.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Trailer`,
	},
	{
		Name:        "WebSockets",
		IsTitle:     true,
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#WebSockets",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#WebSockets`,
	},
	{
		Name:    "Sec-WebSocket-Accept",
		Summary: "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-WebSocket-Accept",
		Description: `The Sec-WebSocket-Accept header is used in the websocket opening handshake. It would appear in the response headers. That is, this is header is sent from server to client to inform that server is willing to initiate a websocket connection.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-WebSocket-Accept`,
	},
	{
		Name:        "Sec-WebSocket-Key",
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#WebSockets",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#WebSockets`,
	},
	{
		Name:        "Sec-WebSocket-Extensions",
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#WebSockets",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#WebSockets`,
	},
	{
		Name:        "Sec-WebSocket-Protocol",
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#WebSockets",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#WebSockets`,
	},
	{
		Name:        "Sec-WebSocket-Version",
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#WebSockets",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#WebSockets`,
	},
	{
		Name:        "Other",
		IsTitle:     true,
		Summary:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Other",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Other`,
	},
	{
		Name:        "Accept-Push-Policy",
		Summary:     "A client can express the desired push policy for a request by sending an Accept-Push-Policy header field in the request.",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Other`,
	},
	{
		Name:        "Accept-Signature",
		Summary:     "A client can send the Accept-Signature header field to indicate intention to take advantage of any available signatures and to indicate what kinds of signatures it supports.",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Other`,
	},
	{
		Name:    "Alt-Svc",
		Summary: "Used to list alternate ways to reach this service.",
		Description: `The Alt-Svc header is used to list alternate ways to reach this website.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Alt-Svc`,
	},
	{
		Name:    "Date",
		Summary: "Contains the date and time at which the message was originated.",
		Description: `The Date general HTTP header contains the date and time at which the message was originated.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Date`,
	},
	{
		Name:    "Large-Allocation",
		Summary: "Tells the browser that the page being loaded is going to want to perform a large allocation.",
		Description: `The non-standard Large-Allocation response header tells the browser that the page being loaded is going to want to perform a large allocation. It is currently only implemented in Firefox, but is harmless to send to every browser.

WebAssembly or asm.js applications can use large contiguous blocks of allocated memory. For complex games, for example, these allocations can be quite large, sometimes as large as 1GB. The Large-Allocation tells the browser that the web content in the to-be-loaded page is going to want to perform a large contiguous memory allocation and the browser can react to this header by starting a dedicated process for the to-be-loaded document, for example.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Large-Allocation`,
	},
	{
		Name:    "Link",
		Summary: "The Link entity-header field provides a means for serialising one or more links in HTTP headers. It is semantically equivalent to the HTML <link> element.",
		Description: `The HTTP Link entity-header field provides a means for serialising one or more links in HTTP headers. It is semantically equivalent to the HTML <link> element.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Link`,
	},
	{
		Name:        "Push-Policy",
		Summary:     "A Push-Policy defines the server behaviour regarding push when processing a request.",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Other`,
	},
	{
		Name:    "Retry-After",
		Summary: "Indicates how long the user agent should wait before making a follow-up request.",
		Description: `The Retry-After response HTTP header indicates how long the user agent should wait before making a follow-up request. There are three main cases this header is used:

    When sent with a 503 (Service Unavailable) response, this indicates how long the service is expected to be unavailable.
    When sent with a 429 (Too Many Requests) response, this indicates how long to wait before making a new request.
    When sent with a redirect response, such as 301 (Moved Permanently), this indicates the minimum time that the user agent is asked to wait before issuing the redirected request.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Retry-After`,
	},
	{
		Name:        "Signature",
		Summary:     "The Signature header field conveys a list of signatures for an exchange, each one accompanied by information about how to determine the authority of and refresh that signature.",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Other`,
	},
	{
		Name:        "Signed-Headers",
		Summary:     "The Signed-Headers header field identifies an ordered list of response header fields to include in a signature.",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Other`,
	},
	{
		Name:    "Server-Timing",
		Summary: "Communicates one or more metrics and descriptions for the given request-response cycle.",
		Description: `The Server-Timing header communicates one or more metrics and descriptions for a given request-response cycle. It is used to surface any backend server timing metrics (e.g. database read/write, CPU time, file system access, etc.) in the developer tools in the user's browser or in the PerformanceServerTiming interface.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Server-Timing`,
	},
	{
		Name:        "Service-Worker-Allowed",
		Summary:     "Used to remove the path restriction by including this header in the response of the Service Worker script.",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Other`,
	},
	{
		Name:    "SourceMap",
		Summary: "Links generated code to a source map.",
		Description: `The SourceMap HTTP response header links generated code to a source map, enabling the browser to reconstruct the original source and present the reconstructed original in the debugger.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/SourceMap`,
	},
	{
		Name:    "Upgrade",
		Summary: "The relevant RFC document for the Upgrade header field is RFC 7230, section 6.7. The standard establishes rules for upgrading or changing to a different protocol on the current client, server, transport protocol connection.",
		Description: `The relevant RFC document for the Upgrade header field is RFC 7230, section 6.7. The standard establishes rules for upgrading or changing to a different protocol on the current client, server, transport protocol connection. For example, this header standard allows a client to change from HTTP 1.1 to HTTP 2.0, assuming the server decides to acknowledge and implement the Upgrade header field. Neither party is required to accept the terms specified in the Upgrade header field. It can be used in both client and server headers. If the Upgrade header field is specified, then the sender MUST also send the Connection header field with the upgrade option specified. For details on the Connection header field please see section 6.1 of the aforementioned RFC.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Other`,
	},
	{
		Name:    "X-DNS-Prefetch-Control",
		Summary: "Controls DNS prefetching, a feature by which browsers proactively perform domain name resolution on both links that the user may choose to follow as well as URLs for items referenced by the document, including images, CSS, JavaScript, and so forth.",
		Description: `The X-DNS-Prefetch-Control HTTP response header controls DNS prefetching, a feature by which browsers proactively perform domain name resolution on both links that the user may choose to follow as well as URLs for items referenced by the document, including images, CSS, JavaScript, and so forth.

This prefetching is performed in the background, so that the DNS is likely to have been resolved by the time the referenced items are needed. This reduces latency when the user clicks a link.

https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-DNS-Prefetch-Control`,
	},
	{
		Name:        "X-Robots-Tag",
		Summary:     "The X-Robots-Tag HTTP header is used to indicate how a web page is to be indexed within public search engine results. The header is effectively equivalent to <meta name=\"robots\" content=\"...\">.",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Other`,
	},
	{
		Name:        "X-UA-Compatible",
		Summary:     "Used by Internet Explorer to signal which document mode to use.",
		Description: `https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#Other`,
	},
}
