package httpref

// Html represents all defined HTML elements published by the Mozilla Developer Network (MDN)
var Html = References{
	{
		Name:    "HTML Elements",
		IsTitle: true,
		Summary: "https://developer.mozilla.org/en-US/docs/Web/HTML/Element",
		Description: `HTML (HyperText Markup Language) is the most basic building block of the Web. It defines the meaning and structure of web content. Other technologies besides HTML are generally used to describe a web page's appearance/presentation (CSS) or functionality/behavior (JavaScript).

HTML uses "markup" to annotate text, images, and other content for display in a Web browser. HTML markup includes special "elements" such as <head>, <title>, <body>, <header>, <footer>, <article>, <section>, <p>, <div>, <span>, <img>, <aside>, <audio>, <canvas>, <datalist>, <details>, <embed>, <nav>, <output>, <progress>, <video>, <ul>, <ol>, <li> and many others.

An HTML element is set off from other text in a document by "tags", which consist of the element name surrounded by "<" and ">". The name of an element inside a tag is case insensitive. That is, it can be written in uppercase, lowercase, or a mixture. For example, the <title> tag can be written as <Title>, <TITLE>, or in any other way. However, the convention and recommended practice is to write tags in lowercase.`,
	},
	{
		Name:    "<a>",
		Summary: "The Anchor element",
		Description: `The <a> HTML element (or _anchor_ element), with its "href" attribute, creates a hyperlink to web pages, files, email addresses, locations in the same page, or anything else a URL can address.

Content within each "<a>" _should_ indicate the link's destination. If the "href" attribute is present, pressing the enter key while focused on the "<a>" element will activate it.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a`,
	},
	{
		Name:    "<abbr>",
		Summary: "The Abbreviation element",
		Description: `The <abbr> HTML element represents an abbreviation or acronym.

When including an abbreviation or acronym, provide a full expansion of the term in plain text on first use, along with the "<abbr>" to mark up the abbreviation. This informs the user what the abbreviation or acronym means.

The optional "title" attribute can provide an expansion for the abbreviation or acronym when a full expansion is not present. This provides a hint to user agents on how to announce/display the content while informing all users what the abbreviation means. If present, "title" must contain this full description and nothing else.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/abbr`,
	},
	{
		Name:    "<acronym>",
		Summary: "",
		Description: `The <acronym> HTML element allows authors to clearly indicate a sequence of characters that compose an acronym or abbreviation for a word.

  **Warning:** Don't use this element. Use the <abbr> element instead.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/acronym`,
	},
	{
		Name:    "<address>",
		Summary: "The Contact Address element",
		Description: `The <address> HTML element indicates that the enclosed HTML provides contact information for a person or people, or for an organization.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/address`,
	},
	{
		Name:    "<applet>",
		Summary: "The Embed Java Applet element",
		Description: `The obsolete **HTML Applet Element** (<applet>) embeds a Java applet into the document; this element has been deprecated in favor of <object>.

Use of Java applets on the Web is deprecated; most browsers no longer support use of plug-ins, including the Java plug-in.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/applet`,
	},
	{
		Name:    "<area>",
		Summary: "The Image Map Area element",
		Description: `The <area> HTML element defines an area inside an image map that has predefined clickable areas. An _image map_ allows geometric areas on an image to be associated with "hypertext link".

This element is used only within a <map> element.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/area`,
	},
	{
		Name:    "<article>",
		Summary: "The Article Contents element",
		Description: `The <article> HTML element represents a self-contained composition in a document, page, application, or site, which is intended to be independently distributable or reusable (e.g., in syndication). Examples include: a forum post, a magazine or newspaper article, or a blog entry, a product card, a user-submitted comment, an interactive widget or gadget, or any other independent item of content.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/article`,
	},
	{
		Name:    "<aside>",
		Summary: "The Aside element",
		Description: `The <aside> HTML element represents a portion of a document whose content is only indirectly related to the document's main content. Asides are frequently presented as sidebars or call-out boxes.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/aside`,
	},
	{
		Name:    "<audio>",
		Summary: "The Embed Audio element",
		Description: `The <audio> HTML element is used to embed sound content in documents. It may contain one or more audio sources, represented using the "src" attribute or the <source> element: the browser will choose the most suitable one. It can also be the destination for streamed media, using a "MediaStream".

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/audio`,
	},
	{
		Name:    "<b>",
		Summary: "The Bring Attention To element",
		Description: `The <b> HTML element is used to draw the reader's attention to the element's contents, which are not otherwise granted special importance. This was formerly known as the Boldface element, and most browsers still draw the text in boldface. However, you should not use "<b>" for styling text; instead, you should use the CSS "font-weight" property to create boldface text, or the <strong> element to indicate that text is of special importance.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/b`,
	},
	{
		Name:    "<base>",
		Summary: "The Document Base URL element",
		Description: `The <base> HTML element specifies the base URL to use for all _relative_ URLs in a document. There can be only one "<base>" element in a document.

A document's used base URL can be accessed by scripts with "Node.baseURI". If the document has no "<base>" elements, then "baseURI" defaults to "location.href".

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/base`,
	},
	{
		Name:    "<bdi>",
		Summary: "The Bidirectional Isolate element",
		Description: `The <bdi> HTML element tells the browser's bidirectional algorithm to treat the text it contains in isolation from its surrounding text. It's particularly useful when a website dynamically inserts some text and doesn't know the directionality of the text being inserted.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/bdi`,
	},
	{
		Name:    "<bdo>",
		Summary: "The Bidirectional Text Override element",
		Description: `The <bdo> HTML element overrides the current directionality of text, so that the text within is rendered in a different direction.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/bdo`,
	},
	{
		Name:    "<bgsound>",
		Summary: "The Background Sound element",
		Description: `The <bgsound> HTML element is deprecated. It sets up a sound file to play in the background while the page is used; use <audio> instead.

  **Warning:** Do not use this! In order to embed audio in a Web page, you should be using the <audio> element.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/bgsound`,
	},
	{
		Name:    "<big>",
		Summary: "The Bigger Text element",
		Description: `The <big> HTML deprecated element renders the enclosed text at a font size one level larger than the surrounding text ("medium" becomes "large", for example). The size is capped at the browser's maximum permitted font size.

  **Warning:** This element has been removed from the specification and shouldn't be used any more. Use the CSS "font-size" property to adjust the font size.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/big`,
	},
	{
		Name:    "<blink>",
		Summary: "The Blinking Text element",
		Description: `The <blink> HTML element is a non-standard element which causes the enclosed text to flash slowly.

  **Warning:** Do not use this element as it is **obsolete** and is bad design practice. Blinking text is frowned upon by several accessibility standards and the CSS specification allows browsers to ignore the "<blink>" element.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/blink`,
	},
	{
		Name:    "<blockquote>",
		Summary: "The Block Quotation element",
		Description: `The <blockquote> HTML element indicates that the enclosed text is an extended quotation. Usually, this is rendered visually by indentation (see Notes for how to change it). A URL for the source of the quotation may be given using the **cite** attribute, while a text representation of the source can be given using the <cite> element.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/blockquote`,
	},
	{
		Name:    "<body>",
		Summary: "The Document Body element",
		Description: `The <body> HTML element represents the content of an HTML document. There can be only one "<body>" element in a document.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/body`,
	},
	{
		Name:    "<br>",
		Summary: "The Line Break element",
		Description: `The <br> HTML element produces a line break in text (carriage-return). It is useful for writing a poem or an address, where the division of lines is significant.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/br`,
	},
	{
		Name:    "<button>",
		Summary: "The Button element",
		Description: `The <button> HTML element is an interactive element activated by a user with a mouse, keyboard, finger, voice command, or other assistive technology. Once activated, it then performs a programmable action, such as submitting a form or opening a dialog.

By default, HTML buttons are presented in a style resembling the platform the "user agent" runs on, but you can change buttons' appearance with CSS.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/button`,
	},
	{
		Name:    "<canvas>",
		Summary: "The Graphics Canvas element",
		Description: `Use the **HTML "<canvas>" element** with either the canvas scripting API or the WebGL API to draw graphics and animations.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/canvas`,
	},
	{
		Name:    "<caption>",
		Summary: "The Table Caption element",
		Description: `The <caption> HTML element specifies the caption (or title) of a table.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/caption`,
	},
	{
		Name:    "<center>",
		Summary: "The Centered Text element",
		Description: `The <center> HTML element is a block-level element that displays its block-level or inline contents centered horizontally within its containing element. The container is usually, but isn't required to be, <body>.

This tag has been deprecated in HTML 4 (and XHTML 1) in favor of the CSS "text-align" property, which can be applied to the <div> element or to an individual <p>. For centering blocks, use other CSS properties like "margin-left" and "margin-right" and set them to "auto" (or set "margin" to "0 auto").

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/center`,
	},
	{
		Name:    "<cite>",
		Summary: "The Citation element",
		Description: `The <cite> HTML element is used to describe a reference to a cited creative work, and must include the title of that work. The reference may be in an abbreviated form according to context-appropriate conventions related to citation metadata.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/cite`,
	},
	{
		Name:    "<code>",
		Summary: "The Inline Code element",
		Description: `The <code> HTML element displays its contents styled in a fashion intended to indicate that the text is a short fragment of computer code. By default, the content text is displayed using the "user agent's" default monospace font.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/code`,
	},
	{
		Name:    "<col>",
		Summary: "The Table Column element",
		Description: `The <col> HTML element defines a column within a table and is used for defining common semantics on all common cells. It is generally found within a <colgroup> element.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/col`,
	},
	{
		Name:    "<colgroup>",
		Summary: "The Table Column Group element",
		Description: `The <colgroup> HTML element defines a group of columns within a table.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/colgroup`,
	},
	{
		Name:    "<content>",
		Summary: "The Shadow DOM Content Placeholder element",
		Description: `The <content> HTML element-an obsolete part of the Web Components suite of technologies-was used inside of Shadow DOM as an "insertion point", and wasn't meant to be used in ordinary HTML. It has now been replaced by the <slot> element, which creates a point in the DOM at which a shadow DOM can be inserted.

  **Note:** Though present in early draft of the specifications and implemented in several browsers, this element has been removed in later versions of the spec, and should not be used. It is documented here to assist in adapting code written during the time it was included in the spec to work with newer versions of the specification.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/content`,
	},
	{
		Name:    "<data>",
		Summary: "The Data element",
		Description: `The <data> HTML element links a given piece of content with a machine-readable translation. If the content is time- or date-related, the <time> element must be used.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/data`,
	},
	{
		Name:    "<datalist>",
		Summary: "The HTML Data List element",
		Description: `The <datalist> HTML element contains a set of <option> elements that represent the permissible or recommended options available to choose from within other controls.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/datalist`,
	},
	{
		Name:    "<dd>",
		Summary: "The Description Details element",
		Description: `The <dd> HTML element provides the description, definition, or value for the preceding term (<dt>) in a description list (<dl>).

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dd`,
	},
	{
		Name:    "<del>",
		Summary: "The Deleted Text element",
		Description: `The <del> HTML element represents a range of text that has been deleted from a document. This can be used when rendering "track changes" or source code diff information, for example. The <ins> element can be used for the opposite purpose: to indicate text that has been added to the document.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/del`,
	},
	{
		Name:    "<details>",
		Summary: "The Details disclosure element",
		Description: `The <details> HTML element creates a disclosure widget in which information is visible only when the widget is toggled into an "open" state. A summary or label must be provided using the <summary> element.

A disclosure widget is typically presented onscreen using a small triangle which rotates (or twists) to indicate open/closed status, with a label next to the triangle. The contents of the "<summary>" element are used as the label for the disclosure widget.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/details`,
	},
	{
		Name:    "<dfn>",
		Summary: "The Definition element",
		Description: `The <dfn> HTML element is used to indicate the term being defined within the context of a definition phrase or sentence. The <p> element, the <dt>/<dd> pairing, or the <section> element which is the nearest ancestor of the "<dfn>" is considered to be the definition of the term.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dfn`,
	},
	{
		Name:    "<dialog>",
		Summary: "The Dialog element",
		Description: `The <dialog> HTML element represents a dialog box or other interactive component, such as a dismissible alert, inspector, or subwindow.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dialog`,
	},
	{
		Name:    "<dir>",
		Summary: "The Directory element",
		Description: `The <dir> HTML element is used as a container for a directory of files and/or folders, potentially with styles and icons applied by the "user agent". Do not use this obsolete element; instead, you should use the <ul> element for lists, including lists of files.

  **Warning:** Do not use this element. Though present in early HTML specifications, it has been deprecated in HTML 4, and has since been removed entirely. **No major browsers support this element.**

## DOM interface

This element implements the "HTMLDirectoryElement" interface.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dir`,
	},
	{
		Name:    "<div>",
		Summary: "The Content Division element",
		Description: `The <div> HTML element is the generic container for flow content. It has no effect on the content or layout until styled in some way using "CSS" (e.g. styling is directly applied to it, or some kind of layout model like Flexbox is applied to its parent element).

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/div`,
	},
	{
		Name:    "<dl>",
		Summary: "The Description List element",
		Description: `The <dl> HTML element represents a description list. The element encloses a list of groups of terms (specified using the <dt> element) and descriptions (provided by <dd> elements). Common uses for this element are to implement a glossary or to display metadata (a list of key-value pairs).

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dl`,
	},
	{
		Name:    "<dt>",
		Summary: "The Description Term element",
		Description: `The <dt> HTML element specifies a term in a description or definition list, and as such must be used inside a <dl> element. It is usually followed by a <dd> element; however, multiple "<dt>" elements in a row indicate several terms that are all defined by the immediate next <dd> element.

The subsequent <dd> (**Description Details**) element provides the definition or other related text associated with the term specified using "<dt>".

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dt`,
	},
	{
		Name:    "<em>",
		Summary: "The Emphasis element",
		Description: `The <em> HTML element marks text that has stress emphasis. The "<em>" element can be nested, with each level of nesting indicating a greater degree of emphasis.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/em`,
	},
	{
		Name:    "<embed>",
		Summary: "The Embed External Content element",
		Description: `The <embed> HTML element embeds external content at the specified point in the document. This content is provided by an external application or other source of interactive content such as a browser plug-in.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/embed`,
	},
	{
		Name:    "<fieldset>",
		Summary: "The Field Set element",
		Description: `The <fieldset> HTML element is used to group several controls as well as labels (<label>) within a web form.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/fieldset`,
	},
	{
		Name:    "<figcaption>",
		Summary: "The Figure Caption element",
		Description: `The <figcaption> HTML element represents a caption or legend describing the rest of the contents of its parent <figure> element.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figcaption`,
	},
	{
		Name:    "<figure>",
		Summary: "The Figure with Optional Caption element",
		Description: `The <figure> HTML element represents self-contained content, potentially with an optional caption, which is specified using the <figcaption> element. The figure, its caption, and its contents are referenced as a single unit.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figure`,
	},
	{
		Name:    "<font>",
		Summary: "",
		Description: `The <font> HTML element defines the font size, color and face for its content.

  **Warning:** Do not use this element. Use the CSS Fonts properties to style text.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/font`,
	},
	{
		Name:    "<footer>",
		Summary: "",
		Description: `The <footer> HTML element represents a footer for its nearest ancestor sectioning content or sectioning root element. A "<footer>" typically contains information about the author of the section, copyright data or links to related documents.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/footer`,
	},
	{
		Name:    "<form>",
		Summary: "The Form element",
		Description: `The <form> HTML element represents a document section containing interactive controls for submitting information.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form`,
	},
	{
		Name:    "<frame>",
		Summary: "",
		Description: `The <frame> HTML element defines a particular area in which another HTML document can be displayed. A frame should be used within a <frameset>.

Using the "<frame>" element is not encouraged because of certain disadvantages such as performance problems and lack of accessibility for users with screen readers. Instead of the "<frame>" element, <iframe> may be preferred.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/frame`,
	},
	{
		Name:    "<frameset>",
		Summary: "",
		Description: `The <frameset> HTML element is used to contain <frame> elements.

  **Note:** Because the use of frames is now discouraged in favor of using <iframe>, this element is not typically used by modern web sites.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/frameset`,
	},
	{
		Name:    "<head>",
		Summary: "The Document Metadata (Header) element",
		Description: `The <head> HTML element contains machine-readable information ("metadata") about the document, like its title, scripts, and style sheets.

  **Note:** "<head>" primarily holds information for machine processing, not human-readability. For human-visible information, like top-level headings and listed authors, see the <header> element.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/head`,
	},
	{
		Name:    "<header>",
		Summary: "",
		Description: `The <header> HTML element represents introductory content, typically a group of introductory or navigational aids. It may contain some heading elements but also a logo, a search form, an author name, and other elements.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/header`,
	},
	{
		Name:    "<h1>-<h6>",
		Summary: "The HTML Section Heading elements",
		Description: `The <h1> to <h6> HTML elements represent six levels of section headings. "<h1>" is the highest section level and "<h6>" is the lowest.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/h1`,
	},
	{
		Name:    "<hgroup>",
		Summary: "",
		Description: `The <hgroup> HTML element represents a heading and related content. It groups a single "<h1>-<h6>" element with one or more "<p>".

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/hgroup`,
	},
	{
		Name:    "<hr>",
		Summary: "The Thematic Break (Horizontal Rule) element",
		Description: `The <hr> HTML element represents a thematic break between paragraph-level elements: for example, a change of scene in a story, or a shift of topic within a section.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/hr`,
	},
	{
		Name:    "<html>",
		Summary: "The HTML Document / Root element",
		Description: `The <html> HTML element represents the root (top-level element) of an HTML document, so it is also referred to as the _root element_. All other elements must be descendants of this element.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/html`,
	},
	{
		Name:    "<i>",
		Summary: "The Idiomatic Text element",
		Description: `The <i> HTML element represents a range of text that is set off from the normal text for some reason, such as idiomatic text, technical terms, taxonomical designations, among others. Historically, these have been presented using italicized type, which is the original source of the "<i>" naming of this element.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/i`,
	},
	{
		Name:    "<iframe>",
		Summary: "The Inline Frame element",
		Description: `The <iframe> HTML element represents a nested "browsing context", embedding another HTML page into the current one.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/iframe`,
	},
	{
		Name:    "<image>",
		Summary: "The Image element",
		Description: `The <image> HTML element is an ancient and poorly supported precursor to the <img> element.\n**It should not be used**.

Some browsers will attempt to automatically convert this into an <img> element, and may succeed if the "src" attribute is specified as well.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/image`,
	},
	{
		Name:    "<img>",
		Summary: "The Image Embed element",
		Description: `The <img> HTML element embeds an image into the document.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/img`,
	},
	{
		Name:    "<input>",
		Summary: "The Input (Form Input) element",
		Description: `The <input> HTML element is used to create interactive controls for web-based forms in order to accept data from the user; a wide variety of types of input data and control widgets are available, depending on the device and "user agent". The "<input>" element is one of the most powerful and complex in all of HTML due to the sheer number of combinations of input types and attributes.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input`,
	},
	{
		Name:    "<ins>",
		Summary: "",
		Description: `The <ins> HTML element represents a range of text that has been added to a document. You can use the <del> element to similarly represent a range of text that has been deleted from the document.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ins`,
	},
	{
		Name:    "<kbd>",
		Summary: "The Keyboard Input element",
		Description: `The <kbd> HTML element represents a span of inline text denoting textual user input from a keyboard, voice input, or any other text entry device. By convention, the "user agent" defaults to rendering the contents of a "<kbd>" element using its default monospace font, although this is not mandated by the HTML standard.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/kbd`,
	},
	{
		Name:    "<keygen>",
		Summary: "",
		Description: `The <keygen> HTML element exists to facilitate generation of key material, and submission of the public key as part of an HTML form. This mechanism is designed for use with Web-based certificate management systems. It is expected that the "<keygen>" element will be used in an HTML form along with other information needed to construct a certificate request, and that the result of the process will be a signed certificate.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/keygen`,
	},
	{
		Name:    "<label>",
		Summary: "The Input Label element",
		Description: `The <label> HTML element represents a caption for an item in a user interface.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label`,
	},
	{
		Name:    "<legend>",
		Summary: "The Field Set Legend element",
		Description: `The <legend> HTML element represents a caption for the content of its parent <fieldset>.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/legend`,
	},
	{
		Name:    "<li>",
		Summary: "The List Item element",
		Description: `The <li> HTML element is used to represent an item in a list. It must be contained in a parent element: an ordered list (<ol>), an unordered list (<ul>), or a menu (<menu>). In menus and unordered lists, list items are usually displayed using bullet points. In ordered lists, they are usually displayed with an ascending counter on the left, such as a number or letter.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/li`,
	},
	{
		Name:    "<link>",
		Summary: "The External Resource Link element",
		Description: `The <link> HTML element specifies relationships between the current document and an external resource.\nThis element is most commonly used to link to "stylesheets", but is also used to establish site icons (both "favicon" style icons and icons for the home screen and apps on mobile devices) among other things.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/link`,
	},
	{
		Name:    "<main>",
		Summary: "",
		Description: `The <main> HTML element represents the dominant content of the <body> of a document. The main content area consists of content that is directly related to or expands upon the central topic of a document, or the central functionality of an application.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/main`,
	},
	{
		Name:    "<map>",
		Summary: "The Image Map element",
		Description: `The <map> HTML element is used with <area> elements to define an image map (a clickable link area).

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/map`,
	},
	{
		Name:    "<mark>",
		Summary: "The Mark Text element",
		Description: `The <mark> HTML element represents text which is **marked** or **highlighted** for reference or notation purposes, due to the marked passage's relevance or importance in the enclosing context.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/mark`,
	},
	{
		Name:    "<marquee>",
		Summary: "The Marquee element",
		Description: `The <marquee> HTML element is used to insert a scrolling area of text. You can control what happens when the text reaches the edges of its content area using its attributes.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/marquee`,
	},
	{
		Name:    "<menu>",
		Summary: "The Menu element",
		Description: `The <menu> HTML element is described in the HTML specification as a semantic alternative to <ul>, but treated by browsers (and exposed through the accessibility tree) as no different than <ul>. It represents an unordered list of items (which are represented by <li> elements).

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/menu`,
	},
	{
		Name:    "<menuitem>",
		Summary: "",
		Description: `The <menuitem> HTML element represents a command that a user is able to invoke through a popup menu. This includes context menus, as well as menus that might be attached to a menu button.

A command can either be defined explicitly, with a textual label and optional icon to describe its appearance, or alternatively as an _indirect command_ whose behavior is defined by a separate element. Commands can also optionally include a checkbox or be grouped to share radio buttons. (Menu items for indirect commands gain checkboxes or radio buttons when defined against elements "<input type="checkbox">" and "<input type="radio">".)

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/menuitem`,
	},
	{
		Name:    "<meta>",
		Summary: "The metadata element",
		Description: `The <meta> HTML element represents "metadata" that cannot be represented by other HTML meta-related elements, like <base>, <link>, <script>, <style> or <title>.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meta`,
	},
	{
		Name:    "<meter>",
		Summary: "The HTML Meter element",
		Description: `The <meter> HTML element represents either a scalar value within a known range or a fractional value.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meter`,
	},
	{
		Name:    "<nav>",
		Summary: "The Navigation Section element",
		Description: `The <nav> HTML element represents a section of a page whose purpose is to provide navigation links, either within the current document or to other documents. Common examples of navigation sections are menus, tables of contents, and indexes.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/nav`,
	},
	{
		Name:    "<nobr>",
		Summary: "The Non-Breaking Text element",
		Description: `The <nobr> HTML element prevents the text it contains from automatically wrapping across multiple lines, potentially resulting in the user having to scroll horizontally to see the entire width of the text.

  **Warning:** Although this element is widely supported, it was _never_ standard HTML, so you shouldn't use it. Instead, use the CSS property "white-space" like this:

<span style="white-space: nowrap;">Long line with no breaks</span>

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/nobr`,
	},
	{
		Name:    "<noembed>",
		Summary: "The Embed Fallback element",
		Description: `The <noembed> HTML element is an obsolete, non-standard way to provide alternative, or "fallback", content for browsers that do not support the <embed> element or do not support the type of embedded content an author wishes to use. This element was deprecated in HTML 4.01 and above in favor of placing fallback content between the opening and closing tags of an <object> element.

  **Note:** While this element currently still works in many browsers, it is obsolete and should not be used. Use <object> instead, with fallback content between the opening and closing tags of the element.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/noembed`,
	},
	{
		Name:    "<noframes>",
		Summary: "The Frame Fallback element",
		Description: `The <noframes> HTML element provides content to be presented in browsers that don't support (or have disabled support for) the <frame> element. Although most commonly-used browsers support frames, there are exceptions, including certain special-use browsers including some mobile browsers, as well as text-mode browsers.

A "<noframes>" element can contain any HTML elements that are allowed within the body of an HTML document, with the exception of the <frameset> and <frame> elements, since using frames when they aren't supported doesn't make sense.

"<noframes>" can be used to present a message explaining that the user's browser doesn't support frames, but ideally should be used to present an alternate form of the site that doesn't use frames but still offers the same or similar functionality.

  **Note:** This element is obsolete and shouldn't be used, since the <frame> and <frameset> elements are also obsolete. When frames are needed at all, they should be presented using the <iframe> element.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/noframes`,
	},
	{
		Name:    "<noscript>",
		Summary: "The Noscript element",
		Description: `The <noscript> HTML element defines a section of HTML to be inserted if a script type on the page is unsupported or if scripting is currently turned off in the browser.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/noscript`,
	},
	{
		Name:    "<object>",
		Summary: "The External Object element",
		Description: `The <object> HTML element represents an external resource, which can be treated as an image, a nested browsing context, or a resource to be handled by a plugin.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/object`,
	},
	{
		Name:    "<ol>",
		Summary: "The Ordered List element",
		Description: `The <ol> HTML element represents an ordered list of items - typically rendered as a numbered list.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ol`,
	},
	{
		Name:    "<optgroup>",
		Summary: "The Option Group element",
		Description: `The <optgroup> HTML element creates a grouping of options within a <select> element.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/optgroup`,
	},
	{
		Name:    "<option>",
		Summary: "The HTML Option element",
		Description: `The <option> HTML element is used to define an item contained in a <select>, an <optgroup>, or a <datalist> element. As such, "<option>" can represent menu items in popups and other lists of items in an HTML document.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/option`,
	},
	{
		Name:    "<output>",
		Summary: "The Output element",
		Description: `The <output> HTML element is a container element into which a site or app can inject the results of a calculation or the outcome of a user action.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/output`,
	},
	{
		Name:    "<p>",
		Summary: "The Paragraph element",
		Description: `The <p> HTML element represents a paragraph. Paragraphs are usually represented in visual media as blocks of text separated from adjacent blocks by blank lines and/or first-line indentation, but HTML paragraphs can be any structural grouping of related content, such as images or form fields.

Paragraphs are block-level elements, and notably will automatically close if another block-level element is parsed before the closing "</p>" tag. See "Tag omission" below.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/p`,
	},
	{
		Name:    "<param>",
		Summary: "The Object Parameter element",
		Description: `The <param> HTML element defines parameters for an <object> element.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/param`,
	},
	{
		Name:    "<picture>",
		Summary: "The Picture element",
		Description: `The <picture> HTML element contains zero or more <source> elements and one <img> element to offer alternative versions of an image for different display/device scenarios.

The browser will consider each child "<source>" element and choose the best match among them. If no matches are found-or the browser doesn't support the "<picture>" element-the URL of the "<img>" element's "src" attribute is selected. The selected image is then presented in the space occupied by the "<img>" element.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/picture`,
	},
	{
		Name:    "<plaintext>",
		Summary: "The Plain Text element (Deprecated)",
		Description: `The <plaintext> HTML element renders everything following the start tag as raw text, ignoring any following HTML. There is no closing tag, since everything after it is considered raw text.

  **Warning:** Do not use this element.\n>\n> - "<plaintext>" is deprecated since HTML 2, and not all browsers implemented it. Browsers that did implement it didn't do so consistently.\n> - "<plaintext>" is obsolete; browsers that accept it may instead treat it as a <pre> element that still interprets HTML within.\n> - If "<plaintext>" is the first element on the page (other than any non-displayed elements, like <head>), do not use HTML at all. Instead serve a text file with the "text/plain" MIME-type.\n> - Instead of "<plaintext>", use the <pre> element or, if semantically accurate (such as for inline text), the <code> element. Escape any "<", ">" and "&" characters, to prevent browsers inadvertently parsing content the element content as HTML.\n> - A monospaced font can be applied to any HTML element via a CSS "font-family" style with the "monospace" generic value.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/plaintext`,
	},
	{
		Name:    "<portal>",
		Summary: "The Portal element",
		Description: `The <portal> HTML element enables the embedding of another HTML page into the current one for the purposes of allowing smoother navigation into new pages.

A "<portal>" is similar to an "<iframe>". An "<iframe>" allows a separate "browsing context" to be embedded. However, the embedded content of a "<portal>" is more limited than that of an "<iframe>". It cannot be interacted with, and therefore is not suitable for embedding widgets into a document. Instead, the "<portal>" acts as a preview of the content of another page. It can be navigated into therefore allowing for seamless transition to the embedded content.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/portal`,
	},
	{
		Name:    "<pre>",
		Summary: "The Preformatted Text element",
		Description: `The <pre> HTML element represents preformatted text which is to be presented exactly as written in the HTML file. The text is typically rendered using a non-proportional, or monospaced, font. Whitespace inside this element is displayed as written.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/pre`,
	},
	{
		Name:    "<progress>",
		Summary: "The Progress Indicator element",
		Description: `The <progress> HTML element displays an indicator showing the completion progress of a task, typically displayed as a progress bar.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/progress`,
	},
	{
		Name:    "<q>",
		Summary: "The Inline Quotation element",
		Description: `The <q> HTML element indicates that the enclosed text is a short inline quotation. Most modern browsers implement this by surrounding the text in quotation marks. This element is intended for short quotations that don't require paragraph breaks; for long quotations use the <blockquote> element.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/q`,
	},
	{
		Name:    "<rb>",
		Summary: "The Ruby Base element",
		Description: `The <rb> HTML element is used to delimit the base text component of a <ruby> annotation, i.e. the text that is being annotated. One "<rb>" element should wrap each separate atomic segment of the base text.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rb`,
	},
	{
		Name:    "<rp>",
		Summary: "The Ruby Fallback Parenthesis element",
		Description: `The <rp> HTML element is used to provide fall-back parentheses for browsers that do not support display of ruby annotations using the <ruby> element. One "<rp>" element should enclose each of the opening and closing parentheses that wrap the <rt> element that contains the annotation's text.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rp`,
	},
	{
		Name:    "<rt>",
		Summary: "The Ruby Text element",
		Description: `The <rt> HTML element specifies the ruby text component of a ruby annotation, which is used to provide pronunciation, translation, or transliteration information for East Asian typography. The "<rt>" element must always be contained within a <ruby> element.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rt`,
	},
	{
		Name:    "<rtc>",
		Summary: "The Ruby Text Container element",
		Description: `The <rtc> HTML element embraces semantic annotations of characters presented in a ruby of <rb> elements used inside of <ruby> element. <rb> elements can have both pronunciation (<rt>) and semantic (<rtc>) annotations.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rtc`,
	},
	{
		Name:    "<ruby>",
		Summary: "The Ruby Annotation element",
		Description: `The <ruby> HTML element represents small annotations that are rendered above, below, or next to base text, usually used for showing the pronunciation of East Asian characters. It can also be used for annotating other kinds of text, but this usage is less common.

The term _ruby_ originated as a unit of measurement used by typesetters>), representing the smallest size that text can be printed on newsprint while remaining legible.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ruby`,
	},
	{
		Name:    "<s>",
		Summary: "The Strikethrough element",
		Description: `The <s> HTML element renders text with a strikethrough, or a line through it. Use the "<s>" element to represent things that are no longer relevant or no longer accurate. However, "<s>" is not appropriate when indicating document edits; for that, use the <del> and <ins> elements, as appropriate.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/s`,
	},
	{
		Name:    "<samp>",
		Summary: "The Sample Output element",
		Description: `The <samp> HTML element is used to enclose inline text which represents sample (or quoted) output from a computer program. Its contents are typically rendered using the browser's default monospaced font (such as Courier>) or Lucida Console).

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/samp`,
	},
	{
		Name:    "<script>",
		Summary: "The Script element",
		Description: `The <script> HTML element is used to embed executable code or data; this is typically used to embed or refer to JavaScript code. The "<script>" element can also be used with other languages, such as WebGL's GLSL shader programming language and JSON.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script`,
	},
	{
		Name:    "<section>",
		Summary: "The Generic Section element",
		Description: `The <section> HTML element represents a generic standalone section of a document, which doesn't have a more specific semantic element to represent it. Sections should always have a heading, with very few exceptions.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/section`,
	},
	{
		Name:    "<select>",
		Summary: "The HTML Select element",
		Description: `The <select> HTML element represents a control that provides a menu of options.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/select`,
	},
	{
		Name:    "<shadow>",
		Summary: "The Shadow Root element",
		Description: `The <shadow> HTML element-an obsolete part of the Web Components technology suite-was intended to be used as a shadow DOM "insertion point". You might have used it if you have created multiple shadow roots under a shadow host. It is not useful in ordinary HTML.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/shadow`,
	},
	{
		Name:    "<slot>",
		Summary: "The Web Component Slot element",
		Description: `The <slot> HTML element-part of the Web Components technology suite-is a placeholder inside a web component that you can fill with your own markup, which lets you create separate DOM trees and present them together.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/slot`,
	},
	{
		Name:    "<small>",
		Summary: "the side comment element",
		Description: `The <small> HTML element represents side-comments and small print, like copyright and legal text, independent of its styled presentation. By default, it renders text within it one font-size smaller, such as from "small" to "x-small".

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/small`,
	},
	{
		Name:    "<source>",
		Summary: "The Media or Image Source element",
		Description: `The <source> HTML element specifies multiple media resources for the <picture>, the <audio> element, or the <video> element. It is an empty element, meaning that it has no content and does not have a closing tag. It is commonly used to offer the same media content in multiple file formats in order to provide compatibility with a broad range of browsers given their differing support for image file formats and media file formats.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/source`,
	},
	{
		Name:    "<spacer>",
		Summary: "",
		Description: `The <spacer> HTML element is an obsolete HTML element which allowed insertion of empty spaces on pages. It was devised by Netscape to accomplish the same effect as a single-pixel layout image, which was something web designers used to use to add white spaces to web pages without actually using an image. However, "<spacer>" no longer supported by any major browser and the same effects can now be achieved using simple CSS.

Firefox, which is the descendant of Netscape's browsers, removed support for "<spacer>" in version 4.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/spacer`,
	},
	{
		Name:    "<span>",
		Summary: "The Content Span element",
		Description: `The <span> HTML element is a generic inline container for phrasing content, which does not inherently represent anything. It can be used to group elements for styling purposes (using the "class" or "id" attributes), or because they share attribute values, such as "lang". It should be used only when no other semantic element is appropriate. "<span>" is very much like a <div> element, but <div> is a block-level element whereas a "<span>" is an inline element.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/span`,
	},
	{
		Name:    "<strike>",
		Summary: "",
		Description: `The <strike> HTML element places a strikethrough (horizontal line) over text.

  **Warning:** This element is deprecated in HTML 4 and XHTML 1, and obsoleted in the HTML Living Standard. If semantically appropriate, i.e., if it represents _deleted_ content, use <del> instead. In all other cases use <s>.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/strike`,
	},
	{
		Name:    "<strong>",
		Summary: "The Strong Importance element",
		Description: `The <strong> HTML element indicates that its contents have strong importance, seriousness, or urgency. Browsers typically render the contents in bold type.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/strong`,
	},
	{
		Name:    "<style>",
		Summary: "The Style Information element",
		Description: `The <style> HTML element contains style information for a document, or part of a document. It contains CSS, which is applied to the contents of the document containing the "<style>" element.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/style`,
	},
	{
		Name:    "<sub>",
		Summary: "The Subscript element",
		Description: `The <sub> HTML element specifies inline text which should be displayed as subscript for solely typographical reasons. Subscripts are typically rendered with a lowered baseline using smaller text.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/sub`,
	},
	{
		Name:    "<summary>",
		Summary: "The Disclosure Summary element",
		Description: `The <summary> HTML element specifies a summary, caption, or legend for a <details> element's disclosure box. Clicking the "<summary>" element toggles the state of the parent "<details>" element open and closed.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/summary`,
	},
	{
		Name:    "<sup>",
		Summary: "The Superscript element",
		Description: `The <sup> HTML element specifies inline text which is to be displayed as superscript for solely typographical reasons. Superscripts are usually rendered with a raised baseline using smaller text.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/sup`,
	},
	{
		Name:        "<table>",
		Summary:     "The Table element",
		Description: `The https://developer.mozilla.org/en-US/docs/Web/HTML/Element/table`,
	},
	{
		Name:    "<tbody>",
		Summary: "The Table Body element",
		Description: `The <tbody> HTML element encapsulates a set of table rows (<tr> elements), indicating that they comprise the body of the table (<table>).

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tbody`,
	},
	{
		Name:    "<td>",
		Summary: "The Table Data Cell element",
		Description: `The <td> HTML element defines a cell of a table that contains data. It participates in the _table model_.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/td`,
	},
	{
		Name:    "<template>",
		Summary: "The Content Template element",
		Description: `The <template> HTML element is a mechanism for holding "HTML" that is not to be rendered immediately when a page is loaded but may be instantiated subsequently during runtime using JavaScript.

Think of a template as a content fragment that is being stored for subsequent use in the document. While the parser does process the contents of the <template> element while loading the page, it does so only to ensure that those contents are valid; the element's contents are not rendered, however.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/template`,
	},
	{
		Name:    "<textarea>",
		Summary: "The Textarea element",
		Description: `The <textarea> HTML element represents a multi-line plain-text editing control, useful when you want to allow users to enter a sizeable amount of free-form text, for example a comment on a review or feedback form.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/textarea`,
	},
	{
		Name:    "<tfoot>",
		Summary: "The Table Foot element",
		Description: `The <tfoot> HTML element defines a set of rows summarizing the columns of the table.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tfoot`,
	},
	{
		Name:    "<th>",
		Summary: "The Table Header element",
		Description: `The <th> HTML element defines a cell as header of a group of table cells. The exact nature of this group is defined by the "scope" and "headers" attributes.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/th`,
	},
	{
		Name:    "<thead>",
		Summary: "The Table Head element",
		Description: `The <thead> HTML element defines a set of rows defining the head of the columns of the table.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/thead`,
	},
	{
		Name:    "<time>",
		Summary: "The (Date) Time element",
		Description: `The <time> HTML element represents a specific period in time. It may include the "datetime" attribute to translate dates into machine-readable format, allowing for better search engine results or custom features such as reminders.

It may represent one of the following:

- A time on a 24-hour clock.\n- A precise date in the Gregorian calendar (with optional time and timezone information).\n- A valid time duration.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/time`,
	},
	{
		Name:    "<title>",
		Summary: "The Document Title element",
		Description: `The <title> HTML element defines the document's title that is shown in a "browser"'s title bar or a page's tab. It only contains text; tags within the element are ignored.

"""html\n<title>Grandma's Heavy Metal Festival Journal</title>\n"""

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/title`,
	},
	{
		Name:    "<tr>",
		Summary: "The Table Row element",
		Description: `The <tr> HTML element defines a row of cells in a table. The row's cells can then be established using a mix of <td> (data cell) and <th> (header cell) elements.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tr`,
	},
	{
		Name:    "<track>",
		Summary: "The Embed Text Track element",
		Description: `The <track> HTML element is used as a child of the media elements, <audio> and <video>. It lets you specify timed text tracks (or time-based data), for example to automatically handle subtitles. The tracks are formatted in WebVTT format (".vtt" files) - Web Video Text Tracks.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/track`,
	},
	{
		Name:    "<tt>",
		Summary: "The Teletype Text element",
		Description: `The <tt> HTML element creates inline text which is presented using the "user agent's" default monospace font face. This element was created for the purpose of rendering text as it would be displayed on a fixed-width display such as a teletype, text-only screen, or line printer.

The terms **non-proportional**, **monotype**, and **monospace** are used interchangeably and have the same general meaning: they describe a typeface whose characters are all the same number of pixels wide.

This element is obsolete, however. You should use the more semantically helpful <code>, <kbd>, <samp>, or <var> elements for inline text that needs to be presented in monospace type, or the <pre> tag for content that should be presented as a separate block.

  **Note:** If none of the semantic elements are appropriate for your use case (for example, if you need to show some content in a non-proportional font), you should consider using the <span> element, styling it as desired using CSS. The "font-family" property is a good place to start.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tt`,
	},
	{
		Name:    "<u>",
		Summary: "The Unarticulated Annotation (Underline) element",
		Description: `The <u> HTML element represents a span of inline text which should be rendered in a way that indicates that it has a non-textual annotation. This is rendered by default as a simple solid underline, but may be altered using CSS.

  **Warning:** This element used to be called the "Underline" element in older versions of HTML, and is still sometimes misused in this way. To underline text, you should instead apply a style that includes the CSS "text-decoration" property set to "underline".

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/u`,
	},
	{
		Name:    "<ul>",
		Summary: "The Unordered List element",
		Description: `The <ul> HTML element represents an unordered list of items, typically rendered as a bulleted list.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ul`,
	},
	{
		Name:    "<var>",
		Summary: "The Variable element",
		Description: `The <var> HTML element represents the name of a variable in a mathematical expression or a programming context. It's typically presented using an italicized version of the current typeface, although that behavior is browser-dependent.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/var`,
	},
	{
		Name:    "<video>",
		Summary: "The Video Embed element",
		Description: `The <video> HTML element embeds a media player which supports video playback into the document. You can use "<video>" for audio content as well, but the <audio> element may provide a more appropriate user experience.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/video`,
	},
	{
		Name:    "<wbr>",
		Summary: "The Line Break Opportunity element",
		Description: `The <wbr> HTML element represents a word break opportunity-a position within text where the browser may optionally break a line, though its line-breaking rules would not otherwise create a break at that location.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/wbr`,
	},
	{
		Name:    "<xmp>",
		Summary: "",
		Description: `The <xmp> HTML element renders text between the start and end tags without interpreting the HTML in between and using a monospaced font. The HTML2 specification recommended that it should be rendered wide enough to allow 80 characters per line.

  **Note:** Do not use this element.\n>\n> - It has been deprecated since HTML3.2 and was not implemented in a consistent way. It was completely removed from current HTML.\n> - Use the <pre> element or, if semantically adequate, the <code> element instead. Note that you will need to escape the '"<"' character as '"&lt;"' to make sure it is not interpreted as markup.\n> - A monospaced font can also be obtained on any element, by applying an adequate CSS style using "monospace" as the generic-font value for the "font-family" property.

https://developer.mozilla.org/en-US/docs/Web/HTML/Element/xmp`,
	},
}
