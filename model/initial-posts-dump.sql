--
-- PostgreSQL database dump
--

SET statement_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: post; Type: TABLE; Schema: public; Owner: evantbyrne; Tablespace: 
--

CREATE TABLE post (
    id integer NOT NULL,
    url text NOT NULL
);


ALTER TABLE public.post OWNER TO evantbyrne;

--
-- Name: post_id_seq; Type: SEQUENCE; Schema: public; Owner: evantbyrne
--

CREATE SEQUENCE post_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.post_id_seq OWNER TO evantbyrne;

--
-- Name: post_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: evantbyrne
--

ALTER SEQUENCE post_id_seq OWNED BY post.id;


--
-- Name: post_meta; Type: TABLE; Schema: public; Owner: evantbyrne; Tablespace: 
--

CREATE TABLE post_meta (
    post_id integer NOT NULL,
    key text NOT NULL,
    value text NOT NULL
);


ALTER TABLE public.post_meta OWNER TO evantbyrne;

--
-- Name: id; Type: DEFAULT; Schema: public; Owner: evantbyrne
--

ALTER TABLE ONLY post ALTER COLUMN id SET DEFAULT nextval('post_id_seq'::regclass);


--
-- Data for Name: post; Type: TABLE DATA; Schema: public; Owner: evantbyrne
--

COPY post (id, url) FROM stdin;
3	/blog/stop-using-em
4	/blog/css-animations-with-javascript
5	/blog/stripe-for-android-prototype
6	/blog/learning-go
7	/blog/go-production-server-ubuntu-nginx
1	/
2	/blog
\.


--
-- Name: post_id_seq; Type: SEQUENCE SET; Schema: public; Owner: evantbyrne
--

SELECT pg_catalog.setval('post_id_seq', 7, true);


--
-- Data for Name: post_meta; Type: TABLE DATA; Schema: public; Owner: evantbyrne
--

COPY post_meta (post_id, key, value) FROM stdin;
5		
5	markdown	date  : December 4, 2013\r\ntitle : Stripe For Android Prototype\r\nurl   : /blog/stripe-for-android-prototype\r\n\r\nPublished an Android app today as part of a group school project. The app basically just allows users to view and edit charges made to their [Stripe](https://stripe.com/) accounts.\r\n\r\nBeware though, this is a prototype. I know of at least one bug. Also, due to this being a school group project I do not have the ability to publish updates, because the university controls the account used to publish the app. So, **do not use it**. :)\r\n\r\n[Google Play store link](https://play.google.com/store/apps/details?id=edu.cmich.cps396m.stripemanager&hl=en).
5	template	post
5	date	December 4, 2013
5	title	Stripe For Android Prototype
5	content	<p>Published an Android app today as part of a group school project. The app basically just allows users to view and edit charges made to their <a href="https://stripe.com/">Stripe</a> accounts.</p>\n\n<p>Beware though, this is a prototype. I know of at least one bug. Also, due to this being a school group project I do not have the ability to publish updates, because the university controls the account used to publish the app. So, <strong>do not use it</strong>. :)</p>\n\n<p><a href="https://play.google.com/store/apps/details?id=edu.cmich.cps396m.stripemanager&amp;hl=en">Google Play store link</a>.</p>\n
1		
1	markdown	url : /\r\ndesign : home\r\n\r\n<p class="intro">I am a software developer that strives to come up with simple solutions to complex problems. Contributor to many <a href="https://github.com/evantbyrne?tab=repositories">open-source projects</a>. Currently a part-time employee at <a href="http://workwithgrid.com">Grid</a> and full-time college senior.</p>\r\n\r\n## Featured Work\r\n\r\n* [Burner CMS](https://burnercms.com) - All\r\n* [Jim Leszczynski's Website](http://jimleszczynski.com) - Development\r\n* [Valkyrie](https://github.com/evantbyrne/valkyrie) - All\r\n\r\n<!--split-->\r\n\r\n## Contact Information\r\n\r\n* Phone: <a class="tel" href="tel:+1-810-728-1572">+1 (810) 728-1572</a>\r\n* E-Mail: [evantbyrne@gmail.com](mailto:evantbyrne@gmail.com)\r\n* Social: [Twitter](https://twitter.com/evantbyrne), [Github](https://github.com/evantbyrne)
1	template	post
1	design	home
1	content	<p class="intro">I am a software developer that strives to come up with simple solutions to complex problems. Contributor to many <a href="https://github.com/evantbyrne?tab=repositories">open-source projects</a>. Currently a part-time employee at <a href="http://workwithgrid.com">Grid</a> and full-time college senior.</p>\n\n<h2>Featured Work</h2>\n\n<ul>\n<li><a href="https://burnercms.com">Burner CMS</a> - All</li>\n<li><a href="http://jimleszczynski.com">Jim Leszczynski's Website</a> - Development</li>\n<li><a href="https://github.com/evantbyrne/valkyrie">Valkyrie</a> - All</li>\n</ul>\n\n
1	content_2	\n\n<h2>Contact Information</h2>\n\n<ul>\n<li>Phone: <a class="tel" href="tel:+1-810-728-1572">+1 (810) 728-1572</a></li>\n<li>E-Mail: <a href="mailto:evantbyrne@gmail.com">evantbyrne@gmail.com</a></li>\n<li>Social: <a href="https://twitter.com/evantbyrne">Twitter</a>, <a href="https://github.com/evantbyrne">Github</a></li>\n</ul>\n
2		
2	markdown	design : blog\r\ntitle : Blog\r\nurl   : /blog\r\n\r\n## All Writings
2	template	post
2	design	blog
2	title	Blog
2	content	<h2>All Writings</h2>\n
3		
3	markdown	date  : August 8, 2013\r\ntitle : Stop Using em\r\nurl   : /blog/stop-using-em\r\n\r\nUse of the `em` unit of measurement has been pretty heavily promoted by the web development community. Typically, it is used in replacement of `px` for specifying font sizes. Unlike px, the em unit scales up and down with whatever default font size is set for the web browser. This makes it more user-friendly than px for visually-impaired users who configure their web browsers' to use larger default font sizes.\r\n\r\n### How em Works\r\n\r\nThe em is a relative unit of measurement. One em is equal to the parent element's font size. Two em is twice the parent element's font size. So, if the initial font size set by the browser is 16px, then 2em = 32px. Unless, of course, another parent element changes the font size. This is where things get confusing. Consider the following:\r\n\r\n    .e1 {\r\n        font-size: 2em;\r\n    }\r\n    \r\n    .e2 {\r\n        font-size: 1.5em;\r\n    }\r\n\r\nSo all elements with the `e1` class should be 32px and all elements with the `e2` class should be 24px. But what if one of these elements is nested inside the other? Since em uses the parent element's font size to determine the current element's font size, nesting e2 inside of e1 will produce an element with text the size of 48px. As any web developer may imagine, this can cause a lot of confusion when there's a lot of nesting going on in the DOM.\r\n\r\n### Introducing rem: A Better Solution\r\n\r\nLuckily, we also have the `rem` unit. It's exactly the same as em, but it's always relative to the root element's font size.\r\n\r\n    .e1 {\r\n        font-size: 2rem;\r\n    }\r\n    \r\n    .e2 {\r\n        font-size: 1.5rem;\r\n    }\r\n\r\nNow nesting e1 and e2 shouldn't cause their font sizes to become increasingly large.\r\n\r\n### Compatibility\r\n\r\nAlthough the rem unit works with [most browsers](http://caniuse.com/rem), old versions of IE don't like it. If support of IE 8 or lower is needed, then just add the px size right before specifying the rem size.\r\n\r\n    .e1 {\r\n        font-size: 32px;\r\n        font-size: 2rem;\r\n    }\r\n    \r\n    .e2 {\r\n        font-size: 24px;\r\n        font-size: 1.5rem;\r\n    }\r\n\r\nModern web browsers will use rem and old browsers will fallback on the set px value. This also adds a little readability to the CSS by providing the px value for easy reference.
3	template	post
3	date	August 8, 2013
3	title	Stop Using em
4		
7	markdown	date  : March 16, 2014\r\ntitle : Setting Up A Production Go Web Server On Ubuntu With Nginx\r\nurl   : /blog/go-production-server-ubuntu-nginx\r\n\r\n### Overview\r\n\r\nGuide to setting up a simple production environment for Go 1.2 web applications on Ubuntu 12.04 with Nginx as a reverse proxy. This tutorial assumes that you have a basic understanding of the Go programming language and know how to use the command line.\r\n\r\n### Install Go\r\n\r\nFirst, install dependencies for [GVM](https://github.com/moovweb/gvm).\r\n\r\n    sudo apt-get update\r\n    sudo apt-get install curl git mercurial make binutils bison gcc\r\n\r\nNext, install GVM.\r\n\r\n    bash < <(curl -s https://raw.github.com/moovweb/gvm/master/binscripts/gvm-installer)\r\n    source $HOME/.gvm/scripts/gvm\r\n\r\nInstall Go version 1.2 via GVM.\r\n\r\n    gvm install go1.2\r\n    gvm use go1.2 --default\r\n\r\nAppend the following to `~/.bashrc`:\r\n\r\n    export GOPATH=$HOME/go\r\n    export GOBIN=$GOPATH/bin\r\n    export GOMAXPROCS=2\r\n\r\nReload .bashrc with the `source ~/.bashrc` command.\r\n\r\n### Basic Web App\r\n\r\nRun the `mkdir ~/go/src/hello` command to create a directory for your app. Create `~/go/src/hello/hello.go` and run with `go run ~/go/src/hello/hello.go`.\r\n\r\n    package main\r\n    \r\n    import (\r\n        "fmt"\r\n        "net/http"\r\n        "time"\r\n    )\r\n    \r\n    func handler(response http.ResponseWriter, request *http.Request) {\r\n        time.Sleep(time.Second * 5)\r\n        fmt.Fprintf(response, "Hello, %s!", request.URL.Path[1:])\r\n    }\r\n    \r\n    func main() {\r\n        http.HandleFunc("/", handler)\r\n        http.ListenAndServe(":8001", nil)\r\n    }\r\n\r\nEach request will have a five second delay, but requests should be handled concurrently. Try accessing URLs like these:\r\n\r\n    http://myserverip:8001/world\r\n    http://myserverip:8001/Evan\r\n    http://myserverip:8001/Bob\r\n\r\nIf requests aren't being handled concurrently, then you most likely didn't set `GOMAXPROCS` to a value of `2` or higher. Also note that Chrome may load two URLs synchronously if they point to the exact same location.\r\n\r\n### Create Service\r\n\r\nFirst, compile the hello program with the `go install ~/go/src/hello/hello.go` command. The binary will be located at `~/go/bin/hello`.\r\n\r\nNext, create a file at `/etc/init.d/hello.sh` for your startup script.\r\n\r\n    #!/bin/bash\r\n    \r\n    case $1 in\r\n        start)\r\n            echo "Starting hello web app."\r\n            /root/go/bin/hello &\r\n            ;;\r\n        stop)\r\n            echo "Stopping hello web app."\r\n            sudo kill $(sudo lsof -t -i:8001)\r\n            ;;\r\n        *)\r\n            echo "Hello web app service."\r\n            echo $"Usage $0 {start|stop}"\r\n            exit 1\r\n    esac\r\n    exit 0\r\n\r\nNext, give it the correct permissions by running `sudo chmod 755 /etc/init.d/hello.sh`. The service can now be started and stopped by running `sudo /etc/init.d/hello.sh start` and `sudo /etc/init.d/hello.sh stop` respectively.\r\n\r\nLastly, you may also have the hello service run on system boot by running `sudo update-rc.d hello.sh defaults`. For more information on update-rc.d, please see [the manual page](http://manpages.ubuntu.com/manpages/natty/man8/update-rc.d.8.html).\r\n\r\n### Setup Nginx\r\n\r\nFirst, install Nginx.\r\n\r\n    sudo apt-get install nginx\r\n\r\nNext, configure Nginx by editing `/etc/nginx/sites-enabled/default`. Modify `location /` to look like this:\r\n\r\n    location / {\r\n        proxy_set_header X-Real-IP $remote_addr;\r\n        proxy_set_header X-Forwarded-For $remote_addr;\r\n        proxy_set_header Host $host;\r\n        proxy_pass http://127.0.0.1:8001;\r\n    }\r\n\r\nIt's also a good idea to comment out any other locations. E.g., `location /doc/`.\r\n\r\nLastly, restart Nginx.\r\n\r\n    sudo service nginx restart\r\n\r\nAll requests to port 80 (the default HTTP port) will now be forwarded to the Go server listening on port 8001. That's all there is to it! As long as they are running on different ports and Nginx is properly configured, then additional Go web apps can be run as well.
7	template	post
7	date	March 16, 2014
7	title	Setting Up A Production Go Web Server On Ubuntu With Nginx
7	content	<h3>Overview</h3>\n\n<p>Guide to setting up a simple production environment for Go 1.2 web applications on Ubuntu 12.04 with Nginx as a reverse proxy. This tutorial assumes that you have a basic understanding of the Go programming language and know how to use the command line.</p>\n\n<h3>Install Go</h3>\n\n<p>First, install dependencies for <a href="https://github.com/moovweb/gvm">GVM</a>.</p>\n\n<pre>sudo apt-get update\nsudo apt-get install curl git mercurial make binutils bison gcc</pre>\n\n<p>Next, install GVM.</p>\n\n<pre>bash &lt; &lt;(curl -s https://raw.github.com/moovweb/gvm/master/binscripts/gvm-installer)\nsource $HOME/.gvm/scripts/gvm</pre>\n\n<p>Install Go version 1.2 via GVM.</p>\n\n<pre>gvm install go1.2\ngvm use go1.2 --default</pre>\n\n<p>Append the following to <code>~/.bashrc</code>:</p>\n\n<pre>export GOPATH=$HOME/go\nexport GOBIN=$GOPATH/bin\nexport GOMAXPROCS=2</pre>\n\n<p>Reload .bashrc with the <code>source ~/.bashrc</code> command.</p>\n\n<h3>Basic Web App</h3>\n\n<p>Run the <code>mkdir ~/go/src/hello</code> command to create a directory for your app. Create <code>~/go/src/hello/hello.go</code> and run with <code>go run ~/go/src/hello/hello.go</code>.</p>\n\n<pre>package main\n\nimport (\n    &quot;fmt&quot;\n    &quot;net/http&quot;\n    &quot;time&quot;\n)\n\nfunc handler(response http.ResponseWriter, request *http.Request) {\n    time.Sleep(time.Second * 5)\n    fmt.Fprintf(response, &quot;Hello, %s!&quot;, request.URL.Path[1:])\n}\n\nfunc main() {\n    http.HandleFunc(&quot;/&quot;, handler)\n    http.ListenAndServe(&quot;:8001&quot;, nil)\n}</pre>\n\n<p>Each request will have a five second delay, but requests should be handled concurrently. Try accessing URLs like these:</p>\n\n<pre>http://myserverip:8001/world\nhttp://myserverip:8001/Evan\nhttp://myserverip:8001/Bob</pre>\n\n<p>If requests aren't being handled concurrently, then you most likely didn't set <code>GOMAXPROCS</code> to a value of <code>2</code> or higher. Also note that Chrome may load two URLs synchronously if they point to the exact same location.</p>\n\n<h3>Create Service</h3>\n\n<p>First, compile the hello program with the <code>go install ~/go/src/hello/hello.go</code> command. The binary will be located at <code>~/go/bin/hello</code>.</p>\n\n<p>Next, create a file at <code>/etc/init.d/hello.sh</code> for your startup script.</p>\n\n<pre>#!/bin/bash\n\ncase $1 in\n    start)\n        echo &quot;Starting hello web app.&quot;\n        /root/go/bin/hello &amp;\n        ;;\n    stop)\n        echo &quot;Stopping hello web app.&quot;\n        sudo kill $(sudo lsof -t -i:8001)\n        ;;\n    *)\n        echo &quot;Hello web app service.&quot;\n        echo $&quot;Usage $0 {start|stop}&quot;\n        exit 1\nesac\nexit 0</pre>\n\n<p>Next, give it the correct permissions by running <code>sudo chmod 755 /etc/init.d/hello.sh</code>. The service can now be started and stopped by running <code>sudo /etc/init.d/hello.sh start</code> and <code>sudo /etc/init.d/hello.sh stop</code> respectively.</p>\n\n<p>Lastly, you may also have the hello service run on system boot by running <code>sudo update-rc.d hello.sh defaults</code>. For more information on update-rc.d, please see <a href="http://manpages.ubuntu.com/manpages/natty/man8/update-rc.d.8.html">the manual page</a>.</p>\n\n<h3>Setup Nginx</h3>\n\n<p>First, install Nginx.</p>\n\n<pre>sudo apt-get install nginx</pre>\n\n<p>Next, configure Nginx by editing <code>/etc/nginx/sites-enabled/default</code>. Modify <code>location /</code> to look like this:</p>\n\n<pre>location / {\n    proxy_set_header X-Real-IP $remote_addr;\n    proxy_set_header X-Forwarded-For $remote_addr;\n    proxy_set_header Host $host;\n    proxy_pass http://127.0.0.1:8001;\n}</pre>\n\n<p>It's also a good idea to comment out any other locations. E.g., <code>location /doc/</code>.</p>\n\n<p>Lastly, restart Nginx.</p>\n\n<pre>sudo service nginx restart</pre>\n\n<p>All requests to port 80 (the default HTTP port) will now be forwarded to the Go server listening on port 8001. That's all there is to it! As long as they are running on different ports and Nginx is properly configured, then additional Go web apps can be run as well.</p>\n
6		
6	markdown	date  : December 15, 2013\r\ntitle : Learning Go\r\nurl   : /blog/learning-go\r\n\r\nI'm currently in the process of learning the [Go programming language](http://golang.org/). I'm not only doing this because I believe that it's important to keep my tool belt up-to-date by learning new technologies, but also because there are a number of things about Go that I find appealing. Here are the big two:\r\n\r\n### 1. It's Compiled\r\n\r\nFirst of all, Go is a compiled language. I just love the idea of compiling code before running it. Plus the Go compiler does a fantastic job of detecting and explaining syntax errors and warnings. In a way, the compilation process feels like a unit test to me. Also, unlike with other compiled languages, Go makes the development process relatively painless. Just run `go run my_program.go` from the terminal and Go compiles and runs your source code.\r\n\r\n### 2. Built-In Web Server\r\n\r\nGo actually has a pretty decent web server built right into it, which supports HTTPS and multiprocessing right out of the box. Writing a basic HTTP file server only takes a [few lines of code](https://github.com/evantbyrne/go-fun/blob/master/http_server/file.go). I'm primarily a web developer, so the relative simplicity and completeness of the HTTP library is a huge bonus for me. Being able to compile and run a web application without needing to setup and maintain complex and bloated software like Apache is a dream come true.\r\n\r\n### Blog Experiment\r\n\r\nRight now I'm creating a [simple blog application](https://github.com/evantbyrne/go-blog) with Go. At the moment it only has basic CRUD capabilities, but perhaps I'll add authentication at some point. My main goal of this project is to get a good understanding of the strengths and weaknesses of the language. Here's a short list of my impressions so far:\r\n\r\n* Imports behave in a way that seems slightly odd to me.\r\n* Templates work well.\r\n* Reading form data is easy.\r\n* Go provides SQL interfaces, but developers have to use 3rd-party database drivers. I'm using [lib/pq](https://github.com/lib/pq).\r\n* Running SQL select queries is a pretty laborious process; but exec queries are easy.\r\n* The default URL router is overly simplistic, but it is usable and very extendable.\r\n\r\nGo contains all of the components needed to create web applications, but I found myself having to create a lot of helper functions to simplify basic tasks. The most notable example of this was running SQL select queries. So, I'll probably be creating a ___small___ framework at some point to make things easier for myself and others.
6	template	post
6	date	December 15, 2013
6	title	Learning Go
6	content	<p>I'm currently in the process of learning the <a href="http://golang.org/">Go programming language</a>. I'm not only doing this because I believe that it's important to keep my tool belt up-to-date by learning new technologies, but also because there are a number of things about Go that I find appealing. Here are the big two:</p>\n\n<h3>1. It's Compiled</h3>\n\n<p>First of all, Go is a compiled language. I just love the idea of compiling code before running it. Plus the Go compiler does a fantastic job of detecting and explaining syntax errors and warnings. In a way, the compilation process feels like a unit test to me. Also, unlike with other compiled languages, Go makes the development process relatively painless. Just run <code>go run my_program.go</code> from the terminal and Go compiles and runs your source code.</p>\n\n<h3>2. Built-In Web Server</h3>\n\n<p>Go actually has a pretty decent web server built right into it, which supports HTTPS and multiprocessing right out of the box. Writing a basic HTTP file server only takes a <a href="https://github.com/evantbyrne/go-fun/blob/master/http_server/file.go">few lines of code</a>. I'm primarily a web developer, so the relative simplicity and completeness of the HTTP library is a huge bonus for me. Being able to compile and run a web application without needing to setup and maintain complex and bloated software like Apache is a dream come true.</p>\n\n<h3>Blog Experiment</h3>\n\n<p>Right now I'm creating a <a href="https://github.com/evantbyrne/go-blog">simple blog application</a> with Go. At the moment it only has basic CRUD capabilities, but perhaps I'll add authentication at some point. My main goal of this project is to get a good understanding of the strengths and weaknesses of the language. Here's a short list of my impressions so far:</p>\n\n<ul>\n<li>Imports behave in a way that seems slightly odd to me.</li>\n<li>Templates work well.</li>\n<li>Reading form data is easy.</li>\n<li>Go provides SQL interfaces, but developers have to use 3rd-party database drivers. I'm using <a href="https://github.com/lib/pq">lib/pq</a>.</li>\n<li>Running SQL select queries is a pretty laborious process; but exec queries are easy.</li>\n<li>The default URL router is overly simplistic, but it is usable and very extendable.</li>\n</ul>\n\n<p>Go contains all of the components needed to create web applications, but I found myself having to create a lot of helper functions to simplify basic tasks. The most notable example of this was running SQL select queries. So, I'll probably be creating a <strong><em>small</em></strong> framework at some point to make things easier for myself and others.</p>\n
3	content	<p>Use of the <code>em</code> unit of measurement has been pretty heavily promoted by the web development community. Typically, it is used in replacement of <code>px</code> for specifying font sizes. Unlike px, the em unit scales up and down with whatever default font size is set for the web browser. This makes it more user-friendly than px for visually-impaired users who configure their web browsers' to use larger default font sizes.</p>\n\n<h3>How em Works</h3>\n\n<p>The em is a relative unit of measurement. One em is equal to the parent element's font size. Two em is twice the parent element's font size. So, if the initial font size set by the browser is 16px, then 2em = 32px. Unless, of course, another parent element changes the font size. This is where things get confusing. Consider the following:</p>\n\n<pre>.e1 {\n    font-size: 2em;\n}\n\n.e2 {\n    font-size: 1.5em;\n}</pre>\n\n<p>So all elements with the <code>e1</code> class should be 32px and all elements with the <code>e2</code> class should be 24px. But what if one of these elements is nested inside the other? Since em uses the parent element's font size to determine the current element's font size, nesting e2 inside of e1 will produce an element with text the size of 48px. As any web developer may imagine, this can cause a lot of confusion when there's a lot of nesting going on in the DOM.</p>\n\n<h3>Introducing rem: A Better Solution</h3>\n\n<p>Luckily, we also have the <code>rem</code> unit. It's exactly the same as em, but it's always relative to the root element's font size.</p>\n\n<pre>.e1 {\n    font-size: 2rem;\n}\n\n.e2 {\n    font-size: 1.5rem;\n}</pre>\n\n<p>Now nesting e1 and e2 shouldn't cause their font sizes to become increasingly large.</p>\n\n<h3>Compatibility</h3>\n\n<p>Although the rem unit works with <a href="http://caniuse.com/rem">most browsers</a>, old versions of IE don't like it. If support of IE 8 or lower is needed, then just add the px size right before specifying the rem size.</p>\n\n<pre>.e1 {\n    font-size: 32px;\n    font-size: 2rem;\n}\n\n.e2 {\n    font-size: 24px;\n    font-size: 1.5rem;\n}</pre>\n\n<p>Modern web browsers will use rem and old browsers will fallback on the set px value. This also adds a little readability to the CSS by providing the px value for easy reference.</p>\n
4	markdown	date  : October 17, 2013\r\ntitle : CSS Animations With Javascript\r\nurl   : /blog/css-animations-with-javascript\r\n\r\nAnimations should be done with CSS instead of Javascript when possible for a couple reasons. Firstly, CSS animations often use the GPU. Secondly, style definitions naturally belong in the CSS stylesheet and not in Javascript. So why do people use Javascript to animate web elements? Well, Javascript allows our websites to respond to browser events in ways that CSS doesn't. This means that we can't just do away with Javascript entirely.\r\n\r\n### Meeting In The Middle\r\n\r\nTo apply animations first define classes in CSS for the animations and then add/remove those classes from elements using Javascript. Here's the obligatory jQuery example:\r\n\r\n    .foo {\r\n        color: black;\r\n        transition: color 0.5s;\r\n    }\r\n    \r\n    .animate-red {\r\n        color: red;\r\n    }\r\n\r\n<!-- ... -->\r\n\r\n    $('.foo').click(function(e) {\r\n        e.preventDefault();\r\n        $(this).toggleClass('animate-red');\r\n    });\r\n\r\nPretty simple, right? Clicking on the `.foo` element will cause its color to transition between black and red. I imagine that not everything can be animated this way though, so keep an open mind.\r\n\r\n**Note:** It is also possible to move `transition: color 0.5s;` over to the `.animate-red` definition, but that would cause the animation to only run when the class is added. The transition back to black would then be instant when the class is removed.
4	template	post
4	date	October 17, 2013
4	title	CSS Animations With Javascript
4	content	<p>Animations should be done with CSS instead of Javascript when possible for a couple reasons. Firstly, CSS animations often use the GPU. Secondly, style definitions naturally belong in the CSS stylesheet and not in Javascript. So why do people use Javascript to animate web elements? Well, Javascript allows our websites to respond to browser events in ways that CSS doesn't. This means that we can't just do away with Javascript entirely.</p>\n\n<h3>Meeting In The Middle</h3>\n\n<p>To apply animations first define classes in CSS for the animations and then add/remove those classes from elements using Javascript. Here's the obligatory jQuery example:</p>\n\n<pre>.foo {\n    color: black;\n    transition: color 0.5s;\n}\n\n.animate-red {\n    color: red;\n}</pre>\n\n<!-- ... -->\n\n<pre>$('.foo').click(function(e) {\n    e.preventDefault();\n    $(this).toggleClass('animate-red');\n});</pre>\n\n<p>Pretty simple, right? Clicking on the <code>.foo</code> element will cause its color to transition between black and red. I imagine that not everything can be animated this way though, so keep an open mind.</p>\n\n<p><strong>Note:</strong> It is also possible to move <code>transition: color 0.5s;</code> over to the <code>.animate-red</code> definition, but that would cause the animation to only run when the class is added. The transition back to black would then be instant when the class is removed.</p>\n
\.


--
-- Name: post__id; Type: CONSTRAINT; Schema: public; Owner: evantbyrne; Tablespace: 
--

ALTER TABLE ONLY post
    ADD CONSTRAINT post__id PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

