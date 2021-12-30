package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

func getLsDate(date string) (int, int, int) {
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	dateTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		exitf("%v\n", err)
	}

	year, month, day := dateTime.Date()
	return year, int(month), day
}

func ls(address, username, password, startDate, endDate string) {
	token, err := getToken(address, username, password)
	if err != nil {
		exitf("%v\n", err)
	}

	startYear, startMonth, startDay := getLsDate(startDate)

	var endYear, endMonth, endDay int
	if endDate == "" {
		endYear, endMonth, endDay = getLsDate(startDate)
	} else {
		endYear, endMonth, endDay = getLsDate(endDate)
	}

	body, _ := json.Marshal([]Request{
		{
			Command: "Search",
			Action:  0,
			Param: Param{
				Search: &Search{
					StreamType: "main",
					StartTime: Time{
						Year:  startYear,
						Month: startMonth,
						Day:   startDay,
					},
					EndTime: Time{
						Year:   endYear,
						Month:  endMonth,
						Day:    endDay,
						Hour:   23,
						Minute: 59,
						Second: 59,
					},
				},
			},
		},
	})

	resp, err := postRequest(
		getURL(address, "/cgi-bin/api.cgi?cmd=Search&token="+token),
		bytes.NewBuffer(body),
	)
	if err != nil {
		exitf("%v\n", err)
	}

	fmt.Printf("Time\t\t\tDuration\tName\t\n")
	for _, file := range resp.Value.SearchResult.File {
		startTime := toTime(&file.StartTime)
		endTime := toTime(&file.EndTime)

		fmt.Printf("%s\t%s\t\t%s\n",
			startTime.Format("2006-01-02 15:04:05"),
			endTime.Sub(startTime),
			file.Name,
		)
	}
}
