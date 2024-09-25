# triffgonix

<img alt="ai image" height="300px" src="assets/triffgonix_logo_ai.png" width="300px"/>

^ this logo is ai generated, but if you got a better one i'm open for suggestions : )

## what it is about

at work i had the task to mentor an intern into making a real-time dart application for our team.
while teaching him vue and quasar among other things, i wasn't really content with the way we went about the whole project.

therefore i wanted to give it my own spin, using go instead of pocketbase and nuxt instead of quasar for the frontend.
the real-time functionality will, at least so i think, be implemented using web sockets

## setup

### tools

this project is made with the **GONUTS** stack (yeah i know, really creative right?)

GO - well for golang in the backend.

NU - nuxt for the frontend, in the end this means vue.

TS - this can mean two things, typescript which is used in nuxt, but more importantly tailwindcss as the css and ui framework.

this means to run triffgonix you need a node environment and have go installed on your system.

### run project

to run the project, you need to start the nuxt frontend and go backend.

install the client dependencies and run the client:

```bash
cd client
yarn install
yarn dev
```

if you're not using yarn as the node package manager, use npm instead. should work just fine.

then you can install the server dependencies and run the server:

```bash
cd server
go mod tidy
go run main.go
```

### server tests

the backend has unit tests, crazy right? at the moment i mainly use unit tests to develop and assure functionality of the engine package. you can run those tests using following command:

```bash
cd server
go test server/core/engine
```

you can even get a testing coverage for the package using following commands instead:

```bash
go test server/core/engine -coverprofile cover.out
go tool cover -html=cover.out
```

this will open the code coverage in your default browser.

## server code structure

the server code is structured in following "features":

- **api** - here goes the code for the http and web socket interface
- **core** - here goes the code that is common between multiple features. an example is domain data objects
- **dart** - here goes all the code relevant to the dart logic, mainly all the different engines that power different game modes
- **database** - here goes all the code needed for database interactions such as orm logic and entities
