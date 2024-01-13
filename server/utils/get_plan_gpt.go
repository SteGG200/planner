package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// func format_plan(plan string) string {

// }
func GetPlanGpt(usergoal string, time string, information []Query) (string, error) {
	api_key := os.Getenv("OPENAI_API_KEY")
	information_json, _ := json.Marshal(information)

	body_req := CreateBody(
		"You are a planner assistant. User has a goal is: \"" + Escape(usergoal) +
			"\". They have exact " + time + " to achieve the goal no more no less and there are some extra information about user is provided in this json format: " +
			string(information_json) + ".Divide " + time + " into maximum 10 periods. " +
			"Each period should be from a month to another. For example from month 1 - month 4. " +
			"All the period do not need to have the same time interval. " +
			"Tell user what they should do in each period and " +
			"list the books and websites they could use to study in each period" +
			"(answer in a json format the exact same as this example [{\"time\": \"...\", \"plan\": \"plan in that period...\", \"resourses\":" +
			" \" 1) ...  2) ...  \"}...] .",
	)
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(body_req)

	client := &http.Client{}
	req, err := http.NewRequest("POST", endpoint, buffer)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api_key)

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	bodyBytes, _ := io.ReadAll(resp.Body)
	body_resp := make(map[string]interface{})
	_ = json.Unmarshal(bodyBytes, &body_resp)

	plan_result := body_resp["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	log.Println(plan_result)

	return plan_result, nil
}
