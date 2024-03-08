package od

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type FlightInfo struct {
	Flight string
	Number int
}

func AirportFlightScheduler() {
	var str string
	fmt.Scan(&str)

	flights := strings.Split(str, ",")
	flightInfos := make([]FlightInfo, len(flights))
	for i, flight := range flights {
		num, _ := strconv.Atoi(flight[2:])
		flightInfos[i] = FlightInfo{
			Flight: flight[:2],
			Number: num,
		}
	}

	sort.Slice(flightInfos, func(i, j int) bool {
		cmp := strings.Compare(flightInfos[i].Flight, flightInfos[j].Flight)
		if cmp == -1 {
			return true
		} else if cmp == 0 {
			return flightInfos[i].Number < flightInfos[j].Number
		}
		return false
	})

	for i, fi := range flightInfos {
		num := strconv.Itoa(fi.Number)
		if fi.Number < 1000 {
			num = fmt.Sprintf("0%v", fi.Number)
		}
		if i != len(flightInfos)-1 {
			fmt.Printf("%v%v,", fi.Flight, num)
		} else {
			fmt.Printf("%v%v", fi.Flight, num)
		}
	}
}
