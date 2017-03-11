package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/PagerDuty/go-pagerduty"
)

/*
// actual incidentsMapper is defined in credentials.go
// along with userID, serviceID and  authtoken strings
var incidentsMapper = map[string]string{
	"Alert closed":  "resolved",
	"Alert open":    "acknowledged",
}
*/

func main() {

	opts := pagerduty.ListIncidentsOptions{
		UserIDs:    []string{userID},
		Statuses:   []string{"triggered"},
		ServiceIDs: []string{serviceID},
	}
	client := pagerduty.NewClient(authtoken)

	if resp, err := client.ListIncidents(opts); err != nil {
		panic(err)
	} else {
		count := len(resp.Incidents)
		if count <= 0 {
			fmt.Println("No incidents")
			return
		}
		var wg sync.WaitGroup
		for _, incident := range resp.Incidents {
			fmt.Println(incident.Summary)
			for summaryEntry, status := range incidentsMapper {
				if strings.Contains(incident.Summary, summaryEntry) {
					fmt.Println("   |----> will be ", status)
					updatedIncident := pagerduty.Incident{
						APIObject: pagerduty.APIObject{
							ID:   incident.ID,
							Type: "incident_reference",
						},
						Status: status,
					}
					wg.Add(1)
					go func(incident pagerduty.Incident) {
						defer wg.Done()
						e := client.ManageIncidents(actorEmail, []pagerduty.Incident{incident})
						if e != nil {
							fmt.Println(e)
						}
					}(updatedIncident)
					break
				}
			}
		}
		wg.Wait()
	}
}
