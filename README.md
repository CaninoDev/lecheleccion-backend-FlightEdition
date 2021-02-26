# API Server for Lecheleccion
_as coded in Go_

## Concept
Initially this backend was written in Ruby on Rails (See the original [project](https://github.com/CaninoDev/lecheleccion)). As an exercise, the backend has been re-implemented in Go.

## Introduction
This repository is intended to be used with Lecheleccion's front end client. Postgrewsql is used to store the data persistently. 

## Setup & Installation
* Create a database in your local PostgreSQL instance ensuring that you have access to it. 
* Clone this repository `git clone https://github.com/CaninoDev/lecheleccion-backend-golang`
* `cd lechelleccion-backend-golang`
* Edit `config.yaml` with the relevant information to your local setup.
* Run the backend `go run main.go`

## Usage and Details
This backend provides the following API endpoints for consumption:

GET
`/api/articles` <--- Returns all the articles in the database.
`/api/article/:id` <--- Returns one article with the matching `id`.
`/api/bias/:id` <--- Returns a bias instance with the the matching `id`.
`/api/user/:id`<--- Returns a user instance with the matching `id`.

### Notes
Initially (and there is a branch for it), I attempted to set up a websocket server to serve data to the frontend. However this proved to be overkill for the scope of the project. Instead, a simpler RESTful server is implemented. 
