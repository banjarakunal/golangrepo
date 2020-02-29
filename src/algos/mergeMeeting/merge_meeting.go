package main

import ( 
	"fmt"
	"sort"
)

type meeting struct{
	startTime int
	endTime int
}


func mergedMeetings(meetings []meeting) []meeting {

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].startTime < meetings[j].startTime
	})
	fmt.Println("val..", meetings)
	out := []meeting{}
	for i := range meetings{
		if(i == 0 ){
			out = append(out, meetings[i])
			continue
		}
		if(out[len(out)-1].endTime >= meetings[i].startTime){
			var maxEnd int	
			if(out[len(out)-1].endTime >= meetings[i].endTime){
				maxEnd = out[len(out)-1].endTime
			}else{
				maxEnd = meetings[i].endTime
			}
			out[len(out)-1].endTime = maxEnd 
		}else{
			out = append(out, meetings[i])
		}
	}
	return out
}