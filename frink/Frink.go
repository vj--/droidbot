package frink

import (
  "fmt"
  "strconv"
  "net/http"
  "net/url"
  "encoding/json"
  "io/ioutil"
  "droidbot/slacked"
  "math/rand"
)

type Frink struct {}

func SearchAndRespond(s slacked.SlackRequest) {
  // Search with Frikiac
	parse_request_url, _ := url.Parse("https://frinkiac.com/api/search")
	add_query := make(url.Values)
  add_query.Add("q", s.CompString)
	parse_request_url.RawQuery = add_query.Encode()
	request_url := parse_request_url.String()
  fmt.Println("URL: ", request_url)
	resp, err := http.Get(request_url)
	if err != nil {
		fmt.Println("ERROR: ", err)
    // SLACK
    var response slacked.SlackResponse
    response.Response_type = slacked.MSG_USER
    response.Text = "Error talking to homer. Sorry :("
    response.SendResponse(s.Response_url)
    return
	}

  // Response
  body, _ := ioutil.ReadAll(resp.Body)
  defer resp.Body.Close()

  // JSON Parsing
  var respObjArray []interface{}
  err = json.Unmarshal(body, &respObjArray)
  if err != nil {
    fmt.Println("ERROR: ", err)
    // SLACK
    var response slacked.SlackResponse
    response.Response_type = slacked.MSG_USER
    response.Text = "Error talking to homer. SORRY :("
    response.SendResponse(s.Response_url)
    return
  }
  if len(respObjArray) < 1 {
    // No results, sorry
    fmt.Println("ERROR: no results")
    // SLACK
    var response slacked.SlackResponse
    response.Response_type = slacked.MSG_USER
    response.Text = "Sorry, no results."
    response.SendResponse(s.Response_url)
    return
  }
  // Random number.
  randNumber := randInt(0, len(respObjArray) - 1)

  respobjj := (respObjArray[randNumber]).(map[string]interface{})
  episode := respobjj["Episode"].(string)
  timestamp := respobjj["Timestamp"].(float64)
  timestampStr := strconv.FormatFloat(timestamp, 'f', 0, 32)

  // GET screencap
  screencapURL := "https://frinkiac.com/api/caption?e="+episode+"&t="+timestampStr
  fmt.Println("ScreenCapURL: ", screencapURL)
  resp, err = http.Get(screencapURL)
  if err != nil{
    fmt.Println("ERROR: ", err)
    // SLACK
    var response slacked.SlackResponse
    response.Response_type = slacked.MSG_USER
    response.Text = "Error talking to homer, getting image. Sorry :("
    response.SendResponse(s.Response_url)
    return
  }
  // Response
  body, _ = ioutil.ReadAll(resp.Body)
  defer resp.Body.Close()
  // Parse
  var sccapJObj map[string]interface{}
  err = json.Unmarshal(body, &sccapJObj)
  if err != nil {
    fmt.Println("ERROR: ", err)
    // SLACK
    var response slacked.SlackResponse
    response.Response_type = slacked.MSG_USER
    response.Text = "Error talking to homer. SORRY :("
    response.SendResponse(s.Response_url)
    return
  }
  sframe := sccapJObj["Frame"].(map[string]interface{})
  sepisode := sframe["Episode"].(string)
  stimestamp := sframe["Timestamp"].(float64)
  stimestampStr := strconv.FormatFloat(stimestamp, 'f', 0, 32)
  // get first subtitle
  ssubtitles := sccapJObj["Subtitles"].([]interface{})
  ssubtitle := (ssubtitles[0]).(map[string]interface{})
  scontent := ssubtitle["Content"].(string)

  image := "https://frinkiac.com/meme/" + sepisode + "/" + stimestampStr + ".jpg"

  // SLACK
  var response slacked.SlackResponse
  response.Response_type = slacked.MSG_CHANNEL
  response.Text = s.Command + " " +s.Text

  // Attachment
  attachment := make(map[string]string)
  attachment["title"] = scontent
  // attachment["title_link"] = "http://www.google.com"
  attachment["text"] = "- " + s.User_name
  attachment["image_url"] = image
  response.Attachments = []map[string]string{attachment}

  // Pass it off
  response.SendResponse(s.Response_url)
}

func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}
