
# GoShort
[![Go Report Card](https://goreportcard.com/badge/github.com/garaekz/goshort)](https://goreportcard.com/report/github.com/garaekz/goshort)

**GoShort** is a simple Golang API URL Shortener, started as a learning project it gave me a nice time building the first version, I'll try to update it often with new ideas and functionality, one of the things I wanna do first is use Redis + MySQL to balance the load of the current MySQL database, the second thing I wanna do is restructure the files in order to make it easier to mantain. Hope you people like this and I'm open to new ideas, ~~just remember this is just the server part, the client will be public soon~~, some tutorial will be public soon on my blog.

# Setup

 - Install dependencies: `go get -u ./...`
 - Copy `.env.example` and save it as `.env`
 - Edit `.env` file with your `PORT` and `DATABASE_URL` info
 - Install **VueJS** dependencies: `npm install`
 - Compile frontend: `npm run build`
 - Build: `go build`
 - Run the server: `./goshort`

# API

Currently has just two basic functions, search the original URL by a given code and store new URL's randomly generating unique codes.

## Get the original URL

    GET /:code

**Returns:**

    {
	    "url": {
	        "id": "5",
	        "code": "dBA8",
	        "original_url": "https://www.example.com/this-is-a-long-url"
	    }
    }


This code is used to redirect to original URL, if not found shows a 404 Page.

## Submit a URL to shorten

    POST /api/v1/shorten
**Body:**

    {
	    "original_url": "https://www.example.com/lets-see-if-this-works"
	}
**Returns:**

      {
        "url": {
            "code": "pJ7V",
            "original_url": "https://www.example.com/lets-see-if-this-works"
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
 - [ ] Implement QR Code
 - [ ] Implement local storage history to guests
 - [ ] Implement throttle by IP
 - [x] Add LICENSE
 - [x] Add CHANGELOG.md

# Changelog
You can see all our realeases in the[ changelog](CHANGELOG.md)
```
