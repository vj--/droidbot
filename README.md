# Droid for Slack in Go

A Go project to cater requests from Slack to respond as necessary. This was developed for own personal project. In the current state, it interacts with <https://frinkiac.com/> to return a related but random image and post it to slack channel as response.

## User experience flow

Once setup a Slack user will invoke a command similar like this:

`/droid frink Hello, World`

Droidbot will receive the request, search Frinkiac for 'Hello, World' and return an image and post as response on the Slack channel.

## setup

1. Slack will need to be configured to have command `/droid` pass to the URL of the server this project is run on.
2. Modify `slacked/SlackConfig.go` with the parameters you get from Slack
3. Run the server by `go run main.go` or `go install` and running `droidbot` from the bin folder

--------------------------------------------------------------------------------

This is a personal project. I wish to have more commands like `frink` added, like search. The project has less documentation, but will be updated and mode more useful later.
