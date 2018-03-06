package main

import (
	"consul-checks-slack/config"
	"consul-checks-slack/integrations"
	"consul-checks-slack/models"
	"fmt"
	"github.com/hashicorp/consul/api"
	"time"
)

func main() {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		fmt.Println(err.Error())
	}
	qp := api.QueryOptions{
		AllowStale: true,
	}
	var status = make(map[string]int64)

	for {
		hcs, _, err := client.Health().State(api.HealthCritical, &qp)

		if err != nil {
			fmt.Println(err.Error())
		}

		attachments := []models.Attachment{}
		if len(hcs) > 0 {
			for _, hc := range hcs {
				if val, ok := status[hc.CheckID]; ok {
					diff := time.Now().Unix() - val
					//Send alert only if it is failing for more than wait interval
					if diff > *config.WaitInterval {
						aField := []models.AttachmentField{}
						fmt.Printf("Service Name %s with Check Name %s in Node %s "+
							"is %s with Output : %s", hc.ServiceName, hc.Name, hc.Node, hc.Status, hc.Output)
						aField = append(aField, models.AttachmentField{Title: hc.Status + " in Node " + hc.Node, Short: false, Value: "Output : " + hc.Output})
						attachments = append(attachments, models.Attachment{Fields: aField,
							Color: "#C24B36", Title: "Tags : " + fmt.Sprintf("%v", hc.ServiceTags),
							Pretext: "Service name : " + hc.ServiceName + " Check Name : " + hc.Name + " is " + hc.Status})
						integrations.SendMessage(models.SlackMessage{Attachments: attachments})
						//fmt.Println("\n Service ID ", hc.ServiceID, "\n Check ID --> ", hc.CheckID, "\n --Name ", hc.Name, "\n Service name ", hc.ServiceName)

					}
				} else {
					status[hc.CheckID] = time.Now().Unix()
				}

			}
		} else {
			status = make(map[string]int64)
		}

		sleepTime := time.Duration(*config.CheckInterval) * time.Second
		time.Sleep(sleepTime)
	}

}
