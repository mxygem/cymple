# cymple

A simple **Cy**press exa**mple**.

This project consists of a simple server written in Go and some initial example tests. It's not
intended as an example of "the best way" to do things, but rather a way things _could_ be done.

The cypress specs are found in [./cypress/integration/*](/cypress/integration/)

The API spec uses a local data object for its assertions and the UI spec illustrates two different ways of utilizing test fixtures for its assertions.

1. import path
2. cypress' `fixture` command

## Running

Prerequisites:

* [Go](https://go.dev/dl/)
* [Cypress](https://docs.cypress.io/guides/getting-started/installing-cypress)
* [Just*](https://github.com/casey/just#installation)

*Just isn't technically required though the justfile has a recipe that makes orchestrating
everything easier.

`just cypressall` - This command will:
    1. Build the server and start it as a background process
    2. Install node's dependencies
    3. Run all of Cypress' tests in headless mode
    4. Stop the server.

Example output:

```text
╭─11:52 gemini ~/code/mxygem/cymple
╰─> just cypressall
2021/12/19 11:53:02 employees api started on port 3000

up to date in 297ms

> cymple@0.0.1 cy:run
> cypress run --headless --quiet



  Employees API
    ✓ can returns all employees (28ms)
    ✓ can return individual employees (27ms)


  2 passing (64ms)



  homepage initial information
    ✓ displays the expected count of employees (91ms)
    ✓ displays the expected details for each employee (93ms)


  2 passing (199ms)
```

`just dev` - This will utilize [reflex](https://github.com/cespare/reflex) to run the go server and
will watch for any of its go or html files to change. Upon change it will restart the server making
local dev easier.

`just cypress` - Opens the cypress test runner GUI.
