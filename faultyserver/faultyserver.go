package faultyserver

import "strings"

// n is number of server
// logs is log ==> patter "s1 error" or "s2 error"
// if there consecutive logs error for 1 server more than 3
// you have to replace that server

// out is count of replace server

var (
	maxErrorsStatus int32  = 3
	logErrorStatus  string = "error"
)

func countFaults(n int32, logs []string) int32 {
	// read log
	var countServerErrorLog map[string]int32 = make(map[string]int32, n)
	var replaceServer int32 = 0
	for _, log := range logs {
		data := strings.Split(log, " ")
		server := data[0]
		status := data[1]

		if countServerErrorLog[server] >= int32(0) && countServerErrorLog[server] < maxErrorsStatus {
			if status == logErrorStatus {
				countServerErrorLog[server] = countServerErrorLog[server] + int32(1)
			} else {
				countServerErrorLog[server] = 0
			}
		}

		if countServerErrorLog[server] == maxErrorsStatus {
			replaceServer++
			countServerErrorLog[server] = 0
		}
	}
	return replaceServer

	// check replase server
	// var replaceServer int32 = 0
	// for _, errors := range countServerErrorLog {
	// 	if errors >= maxErrorsStatus {
	// 		replaceServer++
	// 	}
	// }
	// return replaceServer
}
