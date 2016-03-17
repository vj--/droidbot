package main

/*
  Author: Vijay Raj
*/

/*
TODO:

1. validate input and token - DONE
2. pass on to correct component - DONE
3. Create JSON back - DONE
4. Test if you can send using repsonse url instead - DONE
5. provide help - DONE
6. pass objects around - DONE
7. Logging
8. Interaction with Frinkiac - KIND OF DONE
9. Randomization
*/

import (
  "fmt"
  "strings"
  "net/http"
  "encoding/json"
  "io/ioutil"
  "github.com/gorilla/mux"
  "droidbot/slacked"
  "droidbot/frink"
)

// MAIN ENTRY
func main() {
  r := mux.NewRouter()

  // Routes
  r.HandleFunc("/", MainHandler)

  // Test Route
  r.HandleFunc("/test", TestHandler)

  // Bind to a port
  http.ListenAndServe(":4000", r)
}

// Main Handler
func MainHandler(w http.ResponseWriter, r *http.Request){

  slackRequest := slacked.NewSRequest(r)

  // Validate request
  err := slackRequest.ValidateRequest()
  if err != nil {
    fmt.Println("Validation failed on slack config: ", err)
    respondHelper(w, slacked.MSG_USER, "Sorry, validation error. Contact @vj")
    return
  }

  // Validate text input - text should be <type> <text-for-type>
  splitArray := strings.Split(slackRequest.Text, " ")
  if len(splitArray) < 1 {
    fmt.Println("Validation failed on text: ", slackRequest.Text)
    respondHelper(w, slacked.MSG_USER, "I can't understand you " + slackRequest.User_name + ". Please use *"+ slackRequest.Command +" help*")
    return
  }

  // Switch on component
  compType := strings.ToLower(splitArray[0])
  switch compType {

    case "frink": // FRINK

      // TODO There got to be a better way in golang :(
      strForComp := ""
      for i:=1; i < len(splitArray); i++ {
        strForComp = strForComp + splitArray[i] + " "
      }

      slackRequest.CompString = strForComp

      // defer to not hold program.
      defer frink.SearchAndRespond(slackRequest)

    case "help": // HELP
      fmt.Println("Help Component asked for")
      respondHelper(w, slacked.MSG_USER, "*/droid help*\n\n_/droid frink <TextForFrink>_ : To get a related image, kind of, from Frinkiac.\n\nInteractions with Frinkiac are loosely based off github.com/gesteves/frinkiac-slack/blob/master/app.rb \n\n*/droid* would ideally be able to recognize natural language and help you. But for now limited. Long way to go...")

    case "about":
      fmt.Println("About Component asked for")
      respondHelper(w, slacked.MSG_USER, "*/droid about*\n\nInteractions with Frinkiac are loosely based off github.com/gesteves/frinkiac-slack/blob/master/app.rb \n\n*/droid* would ideally be able to recognize natural language and help you. But for now limited. Long way to go...")

    // What ?
    default:
      fmt.Println("Unavailable component type: ", compType)
      respondHelper(w, slacked.MSG_USER, "I can't understand you " + slackRequest.User_name + ". Please use *"+ slackRequest.Command +" help*")
      return
  }
}

func respondHelper(w http.ResponseWriter, t string, m string) {
  // Temporary 'I am looking' response. To send actual response to response_url
  w.Header().Set("Content-type", "application/json")
  data := map[string]string { "response_type" : t,
                              "text" : m }
  json_bytes,_ := json.Marshal(data)
  fmt.Fprintf(w, "%s\n", json_bytes)
}

// For testing, to just print the output
func TestHandler(w http.ResponseWriter, r *http.Request){

  raw, err := ioutil.ReadAll(r.Body)
  defer r.Body.Close()
  if err != nil {
    fmt.Println("ERROR in TestHandler: ", err)
    return
  }

  // no Error
  fmt.Printf("TestHandler Raw: %q\n", raw)
}
