# ![WebSiteScreenshot](http://image.prntscr.com/image/ca355a85a09e42ee893cf95027ce877c.png)

# WebSiteScreenshot ![Language Badge](https://img.shields.io/badge/Language-Go-blue.svg) ![Go Report](https://goreportcard.com/badge/github.com/DiSiqueira/WebSiteScreenshot) ![License Badge](https://img.shields.io/badge/License-MIT-blue.svg) ![Status Badge](https://img.shields.io/badge/Status-Beta-brightgreen.svg)

The WebSiteScreenshot's goal is to be a perfect tool providing a stupidly easy-to-use and fast program to take screenshots of website using only the terminal.

The Program work using the thumbnail.ws API its free to take up to 1000 screenshots per month. 

## Project Status

WebSiteScreenshot is on beta. Pull Requests [are welcome](https://github.com/DiSiqueira/WebSiteScreenshot#social-coding)

## Features

- MORE THAN 60 DEFAULT EXTENSIONS!!!
- It's perfect to save images of your favorites websites.
- Instantly download and save your images.
- STUPIDLY [EASY TO USE](https://github.com/DiSiqueira/WebSiteScreenshot#usage)
- Very fast start up and response time
- Uses native libs
- Free to use (up to 1000 screenshots per mouth)

## Installation

### Option 1: Go Get

```bash
$ go get github.com/DiSiqueira/WebSiteScreenshot
$ WebSiteScreenshot -h
```

### Option 2: From source

```bash
$ git clone https://github.com/DiSiqueira/WebSiteScreenshot.git
$ cd WebSiteScreenshot/
$ go build *.go
```

## Usage

### Basic usage

```bash
# Download a screenshot from google
$ WebSiteScreenshot -api=YOURAPIKEY -website=google.com
```

### Show help

```bash
$ WebSiteScreenshot -h
```

## Getting a API Key

Sign up [here](https://thumbnail.ws/sign-up.html) to get an API KEY and make up to 1000 screenshots per mouth for free. You only need an email (no confirmation).


## Program Help

![](http://image.prntscr.com/image/2f86884d7915469bb9b4822fc1c18083.png)

## Contributing

### Bug Reports & Feature Requests

Please use the [issue tracker](https://github.com/DiSiqueira/WebSiteScreenshot/issues) to report any bugs or file feature requests.

### Developing

PRs are welcome. To begin developing, do this:

```bash
$ go get gopkg.in/ini.v1
$ git clone --recursive git@github.com:DiSiqueira/WebSiteScreenshot.git
$ cd WebSiteScreenshot/
$ go run *.go
```

## Social Coding

1. Create an issue to discuss about your idea
2. [Fork it](https://github.com/DiSiqueira/WebSiteScreenshot/fork)
3. Create your feature branch (`git checkout -b my-new-feature`)
4. Commit your changes (`git commit -am 'Add some feature'`)
5. Push to the branch (`git push origin my-new-feature`)
6. Create a new Pull Request
7. Profit! :white_check_mark:

## License

The MIT License (MIT)

Copyright (c) 2013-2017 Diego Siqueira

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.