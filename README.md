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

- transaction and context: prepare for using database to handler context (not completed)

- handlers: set http handlers for `/` default redirecting, next handlers: `/login` (if not authenticated), `/signup` (if sumbit form from `/`) and `/profile/{user}` (if user authenticated redirect by user `id`)

- database: create read\write user\session from\to tables in mapper (first simple, then start use context)

- create signup endpoint: allow signup via curl, set a signup form that send json `{name: "", email: "", password:""}`, write to `users` table in bcrypt

- create auth middleware: read user and read session from datbase

- test: create user and login with curl commands (use `scripts/check.sh`) && use session to access profile page by session token

- use signup and login pages to access profile (additional database `profiles` creation)

- create login endpoint: allow the login first via curl command set a login form that send json `{email: "", password:""}`, authenticate by reading credentials from database

- html template and static content: attach header, login, signup and profile pages

- use json: all forms should send json to the relevant endpoint (find the best way to use html template form with json)

- set profile page: set basic profile page, just `Hello {{.Name}}`

- set header: change the header content based on endpoint location - login with signup in header, signup with login in header, profile with logout in header

- oauth connection: use database structure to set oauth to github and google

- set profile acl: only authenticated user allowed to access his own page

### done

- initial: create docker-compose, directory structs, environment launch script, Makefile (Done)

- dev environment: create dev logger for environment, env package to use environment variable (db connect, static and tempalte directories) (Done)

- create database: dat packages in api, mappers for user and session, context and db migrations

- create api: run main, create a new api server, api router `/login`, `/profile/{user}`, `/signup`, `/logout`

### playlist

relapse records, allthemwitches, foo manchu, ufomammuth, melvins (add more)
