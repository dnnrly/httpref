# HTTP Reference

This is a handy little helper that puts HTTP reference at your fingertips, when you're on the CLI, when you need it.

[![ko-fi](https://www.ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/W7W414S4U)


[![codecov](https://codecov.io/gh/dnnrly/httpref/branch/master/graph/badge.svg?token=HW5W5HEEEX)](https://codecov.io/gh/dnnrly/httpref)
[![godoc](https://godoc.org/github.com/dnnrly/httpref?status.svg)](http://godoc.org/github.com/dnnrly/httpref)
[![report card](https://goreportcard.com/badge/github.com/dnnrly/httpref)](https://goreportcard.com/report/github.com/dnnrly/httpref)

![GitHub watchers](https://img.shields.io/github/watchers/dnnrly/httpref?style=social)
![GitHub stars](https://img.shields.io/github/stars/dnnrly/httpref?style=social)

[![Twitter URL](https://img.shields.io/twitter/url?style=social&url=https%3A%2F%2Fgithub.com%2Fdnnrly%2Fhttpref)](https://twitter.com/intent/tweet?url=https://github.com/dnnrly/httpref)

## But why?

Because I can **never** remember what the bloody http status codes mean, or the details of methods, or whatever. It's annoying!

## Why not just wikipedia?

Well it boils down to 2 things:

1. If I use the browser then it means that I have to use the mouse, or a track pad or something.
2. I'm a little bored on a Friday night. There's some rubbish film on. Quite frankly I'm a little bored.
3. If I make it good enough then people will give me stars. This, as everyone knows, is a proxy for love and it will make up for many deficiencies in my life.
4. My attention to detail is incredible. I therefore invite you to sponsor me using the links at the top of the page. :)

Before we get any further, I'm really interested in how you feel about this tool. Please take the time to fill in this short survey: https://forms.gle/mHh6idwwUvdfYZM67

# Installation

```shell
$ go install github.com/dnnrly/httpref/cmd/httpref@latest
```

# Usage

## Filter by Title

```
$ httpref 1
  1xx  Informational response
  100  Continue
  101  Switching
  102  Processing
  103  Early hints


$ httpref 200
200 - OK

The HTTP 200 OK success status response code indicates that the request has succeeded. A 200 response is cacheable by default.

The meaning of a success depends on the HTTP request method:

    GET: The resource has been fetched and is transmitted in the message body.
    HEAD: The entity headers are in the message body.
    POST: The resource describing the result of the action is transmitted in the message body.
    TRACE: The message body contains the request message as received by the server.

The successful result of a PUT or a DELETE is often not a 200 OK but a 204 No Content (or a 201 Created when the resource is uploaded for the first time).

https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/200


$ httpref <div>
<div>
  The Content Division element

The <div> HTML element is the generic container for flow content. It has no effect on the content or
layout until styled in some way using "CSS" (e.g. styling is directly applied to it, or some kind of
layout model like Flexbox is applied to its parent element).

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/div
```

## Full-text Search

The `--search` option accepts a term to full-text search. This option is available for header, html, method, and status references.

```
$ httpref --search clear
Clear-Site-Data
  Clears browsing data (e.g. cookies, storage, cache) associated with the requesting
   website.
205
  Reset Content
431
  Request Header Fields Too Large
```

## HTML Elements Reference

The `html` subcommand looks up 100+ active and deprecated HTML elements.

```
$ httpref html <abbr>
<abbr>
  The Abbreviation element

The <abbr> HTML element represents an abbreviation or acronym.

When including an abbreviation or acronym, provide a full expansion of the term in plain text on
first use, along with the "<abbr>" to mark up the abbreviation. This informs the user what the
abbreviation or acronym means.

The optional "title" attribute can provide an expansion for the abbreviation or acronym when a full
expansion is not present. This provides a hint to user agents on how to announce/display the content
while informing all users what the abbreviation means. If present, "title" must contain this full
description and nothing else.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/abbr


$ httpref html --search anchor
<a>
  The Anchor element

The <a> HTML element (or _anchor_ element), with its "href" attribute, creates a hyperlink to web
pages, files, email addresses, locations in the same page, or anything else a URL can address.

Content within each "<a>" _should_ indicate the link's destination. If the "href" attribute is
present, pressing the enter key while focused on the "<a>" element will activate it.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a
```

## Important `make` targets

* `deps` - downloads all of the deps you need to build, test, and release
* `build` - builds your application
* `test` - runs unit tests
* `ci-test` - run tests for CI validation
* `acceptance-test` - run the acceptance tests
* `lint` -  run linting
* `update` - update Go dependencies
* `clean` - clean project dependencies
* `clean-deps` - remove all of the build dependencies too
