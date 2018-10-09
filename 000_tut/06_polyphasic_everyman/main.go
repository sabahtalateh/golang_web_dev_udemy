package main

import (
	"fmt"
	"os"
	"strconv"
)

const defaultDuration = 480
const minutesInDay = 1440

type period struct {
	name      string
	startMins int
	endMins   int
}

func main() {
	schedule := make(map[int]period)

	duration := getDuration()

	minsLeft := float64(minutesInDay)

	// Large period
	largePeriodProportion := 180.0 / 240.0

	currentTimeMins := 0.0
	largePeriodDuration := float64(duration) * largePeriodProportion
	minsLeft -= largePeriodDuration
	schedule[0] = period{
		name:      "Big Sleep",
		startMins: 0,
		endMins:   int(largePeriodDuration),
	}
	currentTimeMins += largePeriodDuration

	// Small periods 1
	smallPeriodsCount := 3
	smallPeriodProportion := 20.0 / 240.0
	smallPeriodDuration := float64(duration) * smallPeriodProportion
	minsWillLeaveAfterAllSmallPeriods := minsLeft - (smallPeriodDuration * float64(smallPeriodsCount))
	noSleepPeriodDuration := minsWillLeaveAfterAllSmallPeriods / 4

	start := int(currentTimeMins)
	currentTimeMins += noSleepPeriodDuration
	end := int(currentTimeMins)
	schedule[1] = period{
		name:      "Super Day",
		startMins: start,
		endMins:   end,
	}

	start = int(currentTimeMins)
	currentTimeMins += smallPeriodDuration
	end = int(currentTimeMins)
	schedule[2] = period{
		name:      "Little Sleep",
		startMins: start,
		endMins:   end,
	}

	// Small periods 2
	start = int(currentTimeMins)
	currentTimeMins += noSleepPeriodDuration
	end = int(currentTimeMins)
	schedule[3] = period{
		name:      "Super Day",
		startMins: start,
		endMins:   end,
	}

	start = int(currentTimeMins)
	currentTimeMins += smallPeriodDuration
	end = int(currentTimeMins)
	schedule[4] = period{
		name:      "Little Sleep",
		startMins: start,
		endMins:   end,
	}

	// Small periods 3
	start = int(currentTimeMins)
	currentTimeMins += noSleepPeriodDuration
	end = int(currentTimeMins)
	schedule[5] = period{
		name:      "Super Day",
		startMins: start,
		endMins:   end,
	}

	start = int(currentTimeMins)
	currentTimeMins += smallPeriodDuration
	end = int(currentTimeMins)
	schedule[6] = period{
		name:      "Little Sleep",
		startMins: start,
		endMins:   end,
	}

	// Small periods 4
	start = int(currentTimeMins)
	currentTimeMins += noSleepPeriodDuration
	end = int(currentTimeMins)
	schedule[7] = period{
		name:      "Super Day",
		startMins: start,
		endMins:   end,
	}

	startTimeHours := getStartHours()

	for i := 0; i < 8; i++ {
		p := schedule[i]
		fmt.Printf("%-12s | ", p.name)

		startHours := p.startMins / 60.0
		startMins := p.startMins - (startHours * 60)

		endHours := p.endMins / 60.0
		endMins := p.endMins - (endHours * 60)

		startHours += startTimeHours
		endHours += startTimeHours

		startHours = startHours % 24
		endHours = endHours % 24

		fmt.Printf("%02d:%02d - %02d:%02d\n", startHours, startMins, endHours, endMins)
	}
	fmt.Printf("Total sleep = %2d hours\nTotal day   = %2d hours\n", (duration / 60), ((1440 - duration) / 60))
}

func getDuration() int {
	args := os.Args

	var duration int
	var err error
	if len(args) > 1 {
		duration, err = strconv.Atoi(args[1])
		duration *= 60
		if err != nil {
			fmt.Println("Wrong duration, duration set to 8 hours")
			duration = defaultDuration
		}
	} else {
		duration = defaultDuration
	}

	return duration
}

func getStartHours() int {
	args := os.Args

	var startHours int
	var err error
	if len(args) > 2 {
		startHours, err = strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Wrong start hours, start hours set to 0")
			startHours = 0
		}
	} else {
		startHours = 0
	}

	return startHours
}
