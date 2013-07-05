package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	daystr = [...]string{
		"", "01", "02", "03", "04", "05", "06", "07",
		"08", "09", "10", "11", "12", "13", "14",
		"15", "16", "17", "18", "19", "20", "21",
		"22", "23", "24", "25", "26", "27", "28",
		"29", "30", "31",
	}
	monthstr = [...]string{"",
		"January", "February", "March", "April",
		"May", "June", "July", "August", "September",
		"October", "November", "December",
	}
	daysInMonth = [...]int{0,
		31, 29, 31, 30,
		31, 30, 31, 31,
		30, 31, 30, 31,
	}

	today = time.Now()
)

type year int

const calTemplate = "\n" +
	"  ]====================================[   \n" +
	"  |  Su | Mo | Tu | We | Th | Fr | Sa  |   \n" +
	"  ]====================================[   \n" +
	"   | %v | %v | %v | %v | %v | %v | %v |    \n" +
	"   | %v | %v | %v | %v | %v | %v | %v |    \n" +
	"   | %v | %v | %v | %v | %v | %v | %v |    \n" +
	"   | %v | %v | %v | %v | %v | %v | %v |    \n" +
	"   | %v | %v | %v | %v | %v | %v | %v |    \n" +
	"   | %v | %v | %v | %v | %v | %v | %v |    \n" +
	"  ]====================================[   \n" +
	" <                 %v                   >  \n" +
	"  ]====================================[   \n"

func main() {
	cal(procFlags())

}

func cal(m int, y year) {
	y.numDaysPerMonth()
	fmt.Println(daysInMonth)
	var (
		b          []byte
		d          int
		dayOfFirst = y.jan1()
	)
	for _, v := range daysInMonth {
		d += v
	}

	b = make([]byte, d+dayOfFirst)

	var offsetDays []byte

	switch dayOfFirst {
	case 0:
		offsetDays = []byte{0}
	case 1:
		offsetDays = []byte{31}
	case 2:
		offsetDays = []byte{30, 31}
	case 3:
		offsetDays = []byte{29, 30, 31}
	case 4:
		offsetDays = []byte{28, 29, 30, 31}
	case 5:
		offsetDays = []byte{27, 28, 29, 30, 31}
	case 6:
		offsetDays = []byte{26, 27, 28, 29, 30, 31}
	}
	copy(b, offsetDays)

	var offs = len(offsetDays) - 1
	for mo, ndays := range daysInMonth {
		var _ = mo
		for i := 0; i < ndays; i++ {
			b[offs] = byte(i + 1)
			offs++
		}
	}
	/*
		fmt.Printf(calTemplate,
			daystr[b[offset+0]], daystr[b[offset+1]], daystr[b[offset+2]], daystr[b[offset+3]], daystr[b[offset+4]], daystr[b[offset+5]], daystr[b[offset+6]],
			daystr[b[offset+7]], daystr[b[offset+8]], daystr[b[offset+9]], daystr[b[offset+10]], daystr[b[offset+11]], daystr[b[offset+12]], daystr[b[offset+13]],
			daystr[b[offset+14]], daystr[b[offset+15]], daystr[b[offset+16]], daystr[b[offset+17]], daystr[b[offset+18]], daystr[b[offset+19]], daystr[b[offset+20]],
			daystr[b[offset+21]], daystr[b[offset+22]], daystr[b[offset+23]], daystr[b[offset+24]], daystr[b[offset+25]], daystr[b[offset+26]], daystr[b[offset+27]],
			daystr[b[offset+28]], daystr[b[offset+29]], daystr[b[offset+30]], daystr[b[offset+31]], daystr[b[offset+32]], daystr[b[offset+33]], daystr[b[offset+34]],
			daystr[b[offset+35]], daystr[b[offset+36]], daystr[b[offset+37]], daystr[b[offset+38]], daystr[b[offset+39]], daystr[b[offset+40]], daystr[b[offset+41]],
			daystr[b[offset+42]], daystr[b[offset+43]], daystr[b[offset+44]], daystr[b[offset+45]], daystr[b[offset+46]], daystr[b[offset+47]], daystr[b[offset+48]],
			monthstr[mo])
	}*/

}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "%s:\t%s\n", "awk", err.Error())
	os.Exit(2)
}

func procFlags() (m int, y year) {
	args := os.Args[1:]
	y = year(today.Year())
	for _, a := range args {
		if m == 0 &&
			(a[0] >= 'A' && a[0] <= 'z' ||
				(len(a) <= 2 && a[0] >= '0' && a[0] <= '9')) {

			if a[0] < 'a' {
				a = string(a[0]+'a'-'A') + a[1:]
			}
			if len(a) > 1 && a[0] < 'A' {
				switch a[1] {
				case 0:
				case 1:
				case 2:
				default:
					goto YEAR
				}
			}
			switch a {
			case "jan", "january", "1":
				m = 1
			case "feb", "february", "2":
				m = 2
			case "mar", "march", "3":
				m = 3
			case "apr", "april", "4":
				m = 4
			case "may", "5":
				m = 5
			case "jun", "june", "6":
				m = 6
			case "jul", "july", "7":
				m = 7
			case "aug", "august", "8":
				m = 8
			case "sep", "september", "9":
				m = 9
			case "oct", "october", "10":
				m = 10
			case "nov", "november", "11":
				m = 11
			case "dec", "december", "12":
				m = 12
			default:
				fatal(fmt.Errorf("Invalid month argument value: %s", a))
			}
			continue
		}
	YEAR:
		if len(a) <= 4 && a[0] >= '1' && a[0] <= '9' {

			for _, s := range a {
				if s < '0' || s > '9' {
					fatal(fmt.Errorf("Invalid year argument value: %s", a))
				}
			}
			if tmp, err := strconv.Atoi(a); err != nil {
				fatal(err)
			} else {
				y = year(tmp)
			}

		} else {

			fatal(fmt.Errorf("Invalid argument value: %s", a))
		}
	}
	return
}

func (yr year) numDaysPerMonth() {
	d := yr.jan1()
	switch (year(int(yr)+1).jan1() + (7 - d)) % 7 {
	case 2:
	case 1:
		daysInMonth[2] = 28
	//1752
	default:
		daysInMonth[9] = 19
	}
}

//	return day of the week
//	of jan 1 of given year
func (yr year) jan1() int {

	//	normal gregorian calendar
	//	one extra day per four years
	y := int(yr)
	d := 4 + y + (y+3)/4

	// 	julian calendar
	// 	regular gregorian
	// 	less three days per 400

	if y > 1800 {
		d -= (y - 1701) / 100
		d += (y - 1601) / 400
	}

	// 	great calendar changeover instant

	if y > 1752 {
		d += 3
	}

	return d % 7
}
