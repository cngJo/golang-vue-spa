# Golang VueJS SPA

This repository serves as an example for using a golang API backend, serving a vuejs application.

## Building

1. Build the SPA
  a. `cd internal/frontend`
  b. `npm clean-install`
  c. `npm run bulid`
2. Build the golang cmd
  a. `go build ./cmd/...`

You'll end up with a `golang-vue-spa` binary containing everything you need.

## Local development

I'd suggest opening this repository and `internal/frontend` in a separate IDE instance, since most IDEs are confused with vuejs when its not at the root directory of the project.

You can use the local DEV Server of vite as you know it with `npm run dev`. Make sure the `VITE_API_BASE_URL` matches the URL your golang API is available and make sure its running.

You than access the local dev server of vite in the browser and use `import.meta.env.VITE_API_BASE_URL` in TypeScript to access the base url of the API.
