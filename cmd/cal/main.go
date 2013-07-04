package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	curmonth, curyear int
	daystr            = [...]string{
		"", "01", "02", "03", "04", "05", "06", "07",
		"08", "09", "10", "11", "12", "13", "14",
		"15", "16", "17", "18", "19", "20", "21",
		"22", "23", "24", "25", "26", "27", "28",
		"29", "30", "31",
	}
	month = [...]int{0,
		31, 29, 31, 30,
		31, 30, 31, 31,
		30, 31, 30, 31,
	}
	monthstr = [...]string{"",
		"January", "February", "March", "April",
		"May", "June", "July", "August", "September",
		"October", "November", "December",
	}
	today = time.Now()
)

const calander = "\n" +
	"  ]====================================[   \n" +
	"  |  Su | Mo | Tu | We | Th | Fr | Sa  |   \n" +
	"  ]====================================[   \n" +
	"   | %S | %S | %S | %S | %S | %S | %S |    \n" +
	"   | %S | %S | %S | %S | %S | %S | %S |    \n" +
	"   | %S | %S | %S | %S | %S | %S | %S |    \n" +
	"   | %S | %S | %S | %S | %S | %S | %S |    \n" +
	"   | %S | %S | %S | %S | %S | %S | %S |    \n" +
	"   | %S | %S | %S | %S | %S | %S | %S |    \n" +
	"  ]====================================[   \n" +
	" <                 %S                   >  \n" +
	"  ]====================================[   \n"

func main() {
	procFlags()
	var p []byte
	cal(curmonth, curyear, &p, 24)
	fmt.Println(p)
}

//	return day of the week
//	of jan 1 of given year
func jan1(yr int) int {

	//	normal gregorian calendar
	//	one extra day per four years

	d := 4 + yr + (yr+3)/4

	// 	julian calendar
	// 	regular gregorian
	// 	less three days per 400

	if yr > 1800 {
		d -= (yr - 1701) / 100
		d += (yr - 1601) / 400
	}

	// 	great calendar changeover instant

	if yr > 1752 {
		d += 3
	}

	return d % 7
}

func cal(m, y int, p *[]byte, w int) {
	d := jan1(y)

	switch jan1((y+1)+7-d) % 7 {

	//non-leap year
	case 1:
		month[2] = 28
		break
	//1752
	default:
		month[9] = 19
		break
	//leap year
	case 2:
	}

	for i := 1; i < m; i++ {
		d += month[i]
	}
	d %= 7
	*p = make([]byte, 128)
	s := (*p)[3*d:]

	for i := 1; i <= month[m]; i++ {

		if i == 3 && month[m] == 19 {
			i += 11
			month[m] += 11
		}
		if i > 9 {
			s[0] = byte(i/10 + 48)
		}
		s = s[1:]
		s[0] = byte(i%10 + 48)
		s = s[2:]
		d++
		if d == 7 {
			d = 0
			s = (*p)[w:]
			p = &s
		}
	}

}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "%s:\t%s\n", "awk", err.Error())
	os.Exit(2)
}

func procFlags() {
	args := os.Args[1:]
	for _, a := range args {
		if curmonth == 0 &&
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
				curmonth = 1
			case "feb", "february", "2":
				curmonth = 2
			case "mar", "march", "3":
				curmonth = 3
			case "apr", "april", "4":
				curmonth = 4
			case "may", "5":
				curmonth = 5
			case "jun", "june", "6":
				curmonth = 6
			case "jul", "july", "7":
				curmonth = 7
			case "aug", "august", "8":
				curmonth = 8
			case "sep", "september", "9":
				curmonth = 9
			case "oct", "october", "10":
				curmonth = 10
			case "nov", "november", "11":
				curmonth = 11
			case "dec", "december", "12":
				curmonth = 12
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

			var err error
			curyear, err = strconv.Atoi(a)
			if err != nil {
				fatal(err)
			}
		} else {

			fatal(fmt.Errorf("Invalid argument value: %s", a))
		}

	}
}
