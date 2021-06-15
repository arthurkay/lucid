# Lucid Framework

[![license](http://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/arthurkay/lucid-framework/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/arthurkay/lucid-framework)](https://goreportcard.com/report/github.com/arthurkay/lucid-framework)

# Introduction

LUCID is a golang backend framework. The frameowrk is meant to address the issue of having applications that all have different code structures.

The lucid cli tool is there to make the generation of the boiler plate code super easy.

The application architecture will take the form of:

* db
* handlers
* middleware
* models
* routes
* utils

## DB

The `db` directory is the place where all db related configurations will be kept.
These include access to different supported databases as well as initialising these instances of databases.

## Handlers

The `handlers` directory is meant for keeping the business logic of routes.
Each route in mapped to a handler, and that handler is responsible to managing the applications business logic.

## Middleware

The `middleware` directory is for storing data that is meant to be processed between requests.

## Models

The `models` directory is for storing data access API's. These apis include everything involved in CRUD operations on data and databases.

## Routes

The `routes` directory is for http routes access. These routes can be of any number, but its recommended to have them in separate files if they handle different kinds of backend access.
For example, the backend uses `api.go` for API routes and `web.go` for web based routes.

## Utils

The `utils` directory is meant for storing helper function. These functions can be as simple as a function to handle error responses to more comples logic of data mangling.


# Lucid CLI

The cli tool for generation and managing of Lucid app implementations.


# Project Status

The application for both the CLI and the actual framework are still in active development.

* [Lucid CLI](https://github.com/arthurkay/lucid)

* [Lucid Framework](https://github.com/arthurkay/lucid-framework)
