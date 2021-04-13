
# GoShort

[![Build Status](https://travis-ci.com/garaekz/goshort.svg?branch=master)](https://travis-ci.com/garaekz/goshort)
[![Go Report Card](https://goreportcard.com/badge/github.com/garaekz/goshort)](https://goreportcard.com/report/github.com/garaekz/goshort)

**GoShort** is a simple Golang API URL Shortener, started as a learning project it gave me a nice time building the first version, I'll try to update it often with new ideas and functionality, ~~one of the things I wanna do first is use Redis + MySQL to balance the load of the current MySQL database, the second thing I wanna do is restructure the files in order to make it easier to mantain. Hope you people like this and I'm open to new ideas, just remember this is just the server part, the client will be public soon, some tutorial will be public soon on my blog~~ this project's now on the second version, I now restructured all the code and I laid the foundations to new cool features, I´ll try to add user registration and IP throttle in order to expand this tool use and functionality.

# Setup

 - Install dependencies: `go get -u ./...`
 - Go to `config/local.yml` and set up your required info
   - `db_type: "postgres" `if you don't set this up it will fall back to **mysql**
   - `dsn: "user:pass@tcp(127.0.0.1)/goshort?charset=utf8&parseTime=true&loc=Local" ` it needs to a be a valid DSN, this example is based on a **mysql** connection where **127.0.0.1** is the host.
   - `jwt_signing_key: "RandomHS256-HMAC_JWTKey" `this is needed in order to work.
 - ~~Copy `.env.example` and save it as `.env`~~
  - ~~Edit `.env` file with your `PORT` and `DATABASE_URL` info`~~

 - Install **VueJS** dependencies: `npm install`
 - Compile frontend: `npm run build`
 - Build: `go build cmd/server/main.go -o goshort`you can user whichever name you want, remember using .exe if you're building this for Windows systems.
 - Run the server: `./goshort`Again, .exe if you're using Windows.

# API

Currently has just two basic functions, search the original URL by a given code and store new URL's randomly generating unique codes.

## Get the original URL

    GET /:code

This code is used to redirect to original URL, if not found shows a 404 Page.

## Submit a URL to shorten
	
	
    POST /v1/links
**Body:**

    {
	    "url": "https://www.example.com/lets-see-if-this-works"
	}
**Returns:**

      {
        "url": {
            "code": "pJ7V",
            "original_url": "https://www.example.com/lets-see-if-this-works",
            "created_at":  "2021-04-13T06:33:47.112707Z"
        }
    }

  
# TO DO

 - [x] Make the redirection in the main route
 - [ ] Make a delete route
 - [ ] Make guest and member URL's
 - [ ] Make a route to fetch all the member URL's
 - [ ] Social login
 - [x] Make **GoShort** client with VueJS
 - [ ] Implement Redis to search by code
 - [ ] Implement visit count
 - [x] Implement QR Code
 - [x] Visual changes (add brand)
 - [x] Implement local storage history to guests
 - [ ] Implement throttle by IP
 - [x] Add LICENSE
 - [x] Add CHANGELOG.md
 - [x] Add CI/CD to the workflow

# Compatibility
The project frontend is built with VueJS, that means ES6 syntax is the standard used, some browsers may not be compliant to this standard yet, thus some issues may arise, we've been notified about a bug happening in iPhone 6s with Safari, following that trace we find the same issue in IE11, we've managed to fix IE11 issue but couldn't reproduce the issue in iPhone 6s yet.

This issue may be found (or not) in browsers not fully compliant with ECMAScript 2015 (ES6), we added [babel-polyfill](https://github.com/babel/babel) to handle this issue but we cannot guarantee that the problem will not arise again in older browsers.

Thanks to **Yan Edy Chota Castillo** who first encountered this bug in production!

# Changelog
You can see all our realeases in the[ changelog](CHANGELOG.md)
