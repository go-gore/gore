## Changelog

### v0.4.1

- modified os.Exit calls to log.Fatal calls with stderr messages passed to the user
- refactored executable path definition, eliminated unnecessary string assignments
- improved source code comment documentation

### v0.4.0

- refactored execution approach to background compile with `go build`, followed by execution of the compiled binary (and away from `go run`)
- refactored temporary Go source file cleanup to a new private function
- refactored temporary executable binary cleanup to a new private function

### v0.3.0

- added support for the use of arguments to the Go source executed with gor

### v0.2.2

- fix: test for the presence of a shebang line before removal of the first line of the file read

### v.0.2.1

- minor refactor of in-application version string
- minor source code refactor

### v0.2.0

- added support for multi-source file main package builds

### v0.1.0

- initial release