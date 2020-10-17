# nogobk

### description

nogobk is a no go back server. awesome.

fully log in system with profile for each user wrriten in go, using session cookies

to read from a postgres database and use middleware for authentication and session

with mux router in a go api server.

endpoints: login, profile/{user}, signup, logout

databases: users, session (related to user id), posts (related to user id)

environment: two docker containers - web api and postgres

it will extend `github.com/nadavbm/gobulenat

use shorter names and password to ease the environment setup

### roadmap

- initial: create docker-compose, directory structs, environment launch script, Makefile (Done)

- dev environment: create dev logger for environment, env package to use environment variable (db connect, static and tempalte directories)

- create database: dat packages in api, mappers for user and session, context

- create api: run main, create a new api server, api router `/login`, `/profile/{user}`, `/signup`, `/logout`

- html template and static content: attach header, login, signup and profile pages

- create signup endpoint: allow signup via curl, set a signup form that send json `{name: "", email: "", password:""}`, write to `users` table in bcrypt

- create login endpoint: allow the login first via curl command set a login form that send json `{email: "", password:""}`, authenticate by reading credentials from database

- use json: all forms should send json to the relevant endpoint (find the best way to use html template form with json)

- set profile page: set basic profile page, just `Hello {{.Name}}`

- create auth middleware: use signup and login pages to access profile

- set header: change the header content based on endpoint location - login with signup in header, signup with login in header, profile with logout in header

- use session: access profile page only with session

- oauth connection: use database structure to set oauth to github and google

- set profile acl: only authenticated user allowed to access his own page

### playlist

relapse records, allthemwitches, foo manchu, ufomammuth, melvins (add more)
