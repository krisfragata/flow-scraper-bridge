package main

import(
	"fmt"
)

func isRelease(date int) bool{
	releasesMap := map[int]string {
		146: "Saturday, May 25",
		147: "Sunday, May 26",
		153: "Saturday, June 1",
		154: "Sunday, June 2",
		161: "Sunday, June 9",
		168: "Sunday, June 16",
		174: "Saturday, June 22",
		175: "Sunday, June 23",
		180: "Friday, June 28",
		181: "Saturday, June 29",
		188: "Saturday, July 6",
		189: "Sunday, July 7",
		195: "Saturday, July 13",
		196: "Sunday, July 14",
		201: "Friday, July 19",
		202: "Saturday, July 20",
		208: "Friday, July 26",
		210: "Sunday, July 28",
		215: "Friday, August 2",
		216: "Saturday, August 3",
		217: "Sunday, August 4",
		222: "Friday, August 9",
		223: "Saturday, August 10",
		230: "Saturday, August 17",
		231: "Sunday, August 18",
		234: "Saturday, August 24",
		238: "Sunday, August 25",
		244: "Saturday, August 31",
		245: "Sunday, September 1",
		251: "Saturday, September 7",
		252: "Sunday, September 8",
		258: "Saturday, September 14",
		259: "Sunday, September 15",
		273: "Sunday, September 29",
		279: "Saturday, October 5",
		287: "Sunday, October 13",
	}

	releaseDays := map[int]string {
		146: "1",
		147: "1",
		153: "1",
		154: "1",
		161: "1",
		168: "1",
		174: "1",
		175: "1",
		180: "1",
		181: "1",
		188: "1",
		189: "1",
		195: "1",
		196: "1",
		201: "1",
		202: "1",
		208: "1",
		210: "1",
		215: "1",
		216: "1",
		217: "1",
		222: "1",
		223: "1",
		230: "1",
		231: "1",
		234: "1",
		238: "1",
		244: "1",
		245: "1",
		251: "1",
		252: "1",
		258: "1",
		259: "1",
		273: "1",
		279: "1",
		287: "1",
	}

	fmt.Println(releasesMap[date])
	
	return releaseDays[date] == "1"

}