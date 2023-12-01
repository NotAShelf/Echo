# Echo

Janus allows you to setup a simple file server for local testing.

### Requirements

- Go (any recent version should work, native to Go 1.21)

## Usage

1. Clone the project
2. Create a `.env` file

- Add `SERVER_PORT={PORT}`
- Add `BASE_PATH={PATH}`

3. Run the program with `go run .`

## Notes

Echo is intended for localhost testing ONLY and designed as a convenient alternative to spinning up a nginx webserver for serving files quickly.
Do not use in a production environment, as no security features are implemented.
