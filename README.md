# GoShort

**GoShort** is a simple Golang API URL Shortener, started as a learning project it gave me a nice time building the first version, I'll try to update it often with new ideas and functionality, one of the things I wanna do first is use Redis + MySQL to balance the load of the current MySQL database, the second thing I wanna do is restructure the files in order to make it easier to mantain. Hope you people like this and I'm open to new ideas, just remember this is just the server part, the client will be public soon.

# Setup

 - Install dependecies: `go get -u ./...`
 - Copy `.env.example` and save it as `.env`
 - Edit `.env` file with your `PORT` and `DB_URL`info
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


This URL will be used to redirect the user to the original url, to be implemented, in [TODO list](#to-do)

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

 - [ ] Make the redirection in the main route
 - [ ] Make a delete route
 - [ ] Make guest and member URL's
 - [ ] Make a route to fetch all the member URL's
 - [ ] Social login
 - [ ] Make **GoShort** client with VueJS
 - [ ] Implement Redis to search by code
 - [ ] Implement visit count
