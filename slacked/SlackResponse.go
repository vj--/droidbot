package slacked

import (
  "fmt"
  "bytes"
  "net/http"
  "encoding/json"
)

type SlackResponse struct {
  Response_type string `json:"response_type"`
  Text string `json:"text"`
  Attachments []map[string]string `json:"attachments"`
}

func (s *SlackResponse) SendResponse(url string){

  newData := s
  json_bytes,_ := json.Marshal(newData)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json_bytes))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("ERROR: ") // TODO Show error details
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
  fmt.Println("------------------------------")
}

// Sample response
//   {
//     "response_type": "in_channel",
//     "text": "It's 80 degrees right now.",
//     "attachments": [
//         {
//             "text":"Partly cloudy today and tomorrow"
//         }
//     ]
// }
