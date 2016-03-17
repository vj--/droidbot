package slacked

import (
  // "fmt"
  "errors"
  "net/http"
)

type SlackRequest struct{
  Token, Team_id, Team_domain, Channel_id, Channel_name, User_name, Text, Response_url, Command, CompString /*Comp String*/ string
}

/*
token=sumhdPO1WRXybqzxuOicPil3
team_id=T0001
team_domain=example
channel_id=C2147483705
channel_name=test
user_id=U2147483697
user_name=Steve
command=/weather
text=94070
response_url=https://hooks.slack.com/commands/1234/5678
*/

// Create new
func NewSRequest(r *http.Request) SlackRequest {
  var s SlackRequest
  s.Token = r.URL.Query().Get("token")
  s.Team_id = r.URL.Query().Get("team_id")
  s.Team_domain = r.URL.Query().Get("team_domain")
  s.Channel_id = r.URL.Query().Get("channel_id")
  s.User_name = r.URL.Query().Get("user_name")
  s.Text = r.URL.Query().Get("text")
  s.Response_url = r.URL.Query().Get("response_url")
  s.Command = r.URL.Query().Get("command")
  return s
}

// Validate request
func (s *SlackRequest) ValidateRequest() error {
  if SLACK_TOKEN != s.Token {
    return errors.New("Token does not match: " + s.Token)
  }

  if TEAM_ID != s.Team_id {
    return errors.New("Team ID does not match: " + s.Team_id)
  }

  if TEAM_DOMAIN != s.Team_domain {
    return errors.New("Team Domain does not match: " + s.Team_domain)
  }

  // no errors.
  return nil
}
