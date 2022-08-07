package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	printCurrentDateTime()
	printCurrentUTCTime()
	printExtractDateParts()
	printSimpleDateParts()
	printParseDateTime()
	printDatetimeArithmetic()
	printDateTimeDuration()
	printDateimeLocation()
}

func printDateimeLocation() {
	names := []string{
		"Local",
		"UTC",
		"Pacific/Galapagos",
		"Europe/Budapest",
		"Europe/Moscow",
		"Asia/Vladivostok",
		"Antarctica/Vostok",
		"America/New_York",
		"Africa/Tripoli",
	}

	now := time.Now()

	for _, name := range names {

		loc, err := time.LoadLocation(name)

		if err != nil {
			log.Fatal(err)
		}

		t := now.In(loc)

		fmt.Println(loc, ": ", t)
	}
}

func printDateTimeDuration() {
	t1 := time.Date(2020, time.November, 10, 23, 0, 0, 0, time.UTC)
	t2 := time.Date(2021, time.July, 28, 16, 22, 0, 0, time.UTC)

	elapsed := t2.Sub(t1)

	fmt.Println("The elapsed in nanoseconds is: ", elapsed)
}

func printDatetimeArithmetic() {
	now := time.Now()

	t1 := now.Add(time.Hour * 27)
	fmt.Println(t1.Format(time.UnixDate))

	t2 := now.AddDate(2, 10, 11)
	fmt.Println(t2.Format(time.UnixDate))

	t3 := now.Add(-time.Hour * 6)
	fmt.Println(t3.Format(time.UnixDate))
}

func printParseDateTime() {
	// array of dates
	vals := []string{"2021-07-28", "2020-11-12", "2019-01-05"}

	for _, val := range vals {
		t, err := time.Parse("2006-01-02", val)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(t)
	}
}

func printSimpleDateParts() {
	year, month, day := time.Now().Date()
	fmt.Println(year, month, day)
	fmt.Println(year, int(month), day)
}

func printExtractDateParts() {
	now := time.Now()

	fmt.Println("Year:", now.Year())
	fmt.Println("Month:", now.Month())
	fmt.Println("Month as Int:", int(now.Month()))
	fmt.Println("Day:", now.Day())
	fmt.Println("Hour:", now.Hour())
	fmt.Println("Minute:", now.Minute())
	fmt.Println("Second:", now.Second())
	fmt.Println("Nanosecond:", now.Nanosecond())
}

func printCurrentUTCTime() {
	utc := time.Now().UTC()
	fmt.Println("Current UTC datetime: ", utc)
}

func printCurrentDateTime() {
	now := time.Now()
	fmt.Println("Current datetime:", now)
}
