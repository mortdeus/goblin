package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	daystr = [...]string{
		"##", "01", "02", "03", "04", "05", "06", "07",
		"08", "09", "10", "11", "12", "13", "14",
		"15", "16", "17", "18", "19", "20", "21",
		"22", "23", "24", "25", "26", "27", "28",
		"29", "30", "31",
	}
	monthstr = [...]string{"",
		"January  %-d", "February  %-d", "March  %-d", "April  %-d",
		"May  %-d", "June  %-d", "July  %-d", " August    %-d", "September  %-d",
		"October  %-d", "November  %-d", "December  %-d",
	}
	daysInMonth = [...]int{0,
		31, 29, 31, 30,
		31, 30, 31, 31,
		30, 31, 30, 31,
	}

	today = time.Now()
)

type year int
type month struct {
	days   []byte
	first  int
	offset int
}

const calTemplate = "\n" +
	"  ]====================================[   \n" +
	"  |  Su | Mo | Tu | We | Th | Fr | Sa  |   \n" +
	"  ]====================================[   \n" +
	"   | %s | %s | %s | %s | %s | %s | %s |    \n" +
	"   | %s | %s | %s | %s | %s | %s | %s |    \n" +
	"   | %s | %s | %s | %s | %s | %s | %s |    \n" +
	"   | %s | %s | %s | %s | %s | %s | %s |    \n" +
	"   | %s | %s | %s | %s | %s | %s | %s |    \n" +
	"   | %s | %s | %s | %s | %s | %s | %s |    \n" +
	"  ]====================================[   \n" +
	" <           %-s   	        >\n" +
	"  ]====================================[   \n"

func main() {
	cal(procFlags())
}

func cal(m int, y year) {
	y.numDaysPerMonth()
	var b = make([]byte, 0)

	janPrefix := [...][]byte{
		[]byte{},
		[]byte{31},
		[]byte{30, 31},
		[]byte{29, 30, 31},
		[]byte{28, 29, 30, 31},
		[]byte{27, 28, 29, 30, 31},
		[]byte{26, 27, 28, 29, 30, 31},
	}
	first := y.jan1()
	b = append(b, janPrefix[first]...)

	offset := first
	months := make([]month, 12)
	for mo, ndays := range daysInMonth {
		if mo == 0 {
			continue
		}
		for i := 0; i < ndays; i++ {
			b = append(b, byte(i+1))
		}
		s := make([]byte, ndays+first)
		v := copy(s, b[offset-first:offset+ndays])
		for i := 0; i < 42-v; i++ {

			s = append(s, byte(i+1))

		}

		offset += ndays
		mon := month{s, first, offset}
		first = (first + ndays) % 7
		months[mo-1] = mon
	}
	tyear, tmonth, tday := today.Date()
	for i, mon := range months {
		if tyear == int(y) && i+1 == int(tmonth) {
			mon.days[(mon.first+tday)-1] = 0
		}
		if i+1 == m || m == 0 {
			fmt.Printf(calTemplate,
				daystr[mon.days[0]], daystr[mon.days[1]], daystr[mon.days[2]], daystr[mon.days[3]], daystr[mon.days[4]], daystr[mon.days[5]], daystr[mon.days[6]],
				daystr[mon.days[7]], daystr[mon.days[8]], daystr[mon.days[9]], daystr[mon.days[10]], daystr[mon.days[11]], daystr[mon.days[12]], daystr[mon.days[13]],
				daystr[mon.days[14]], daystr[mon.days[15]], daystr[mon.days[16]], daystr[mon.days[17]], daystr[mon.days[18]], daystr[mon.days[19]], daystr[mon.days[20]],
				daystr[mon.days[21]], daystr[mon.days[22]], daystr[mon.days[23]], daystr[mon.days[24]], daystr[mon.days[25]], daystr[mon.days[26]], daystr[mon.days[27]],
				daystr[mon.days[28]], daystr[mon.days[29]], daystr[mon.days[30]], daystr[mon.days[31]], daystr[mon.days[32]], daystr[mon.days[33]], daystr[mon.days[34]],
				daystr[mon.days[35]], daystr[mon.days[36]], daystr[mon.days[37]], daystr[mon.days[38]], daystr[mon.days[39]], daystr[mon.days[40]], daystr[mon.days[41]],
				fmt.Sprintf(monthstr[i+1], y))
		}
	}

}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "%s:\t%s\n", "cal", err.Error())
	os.Exit(2)
}

func procFlags() (m int, y year) {
	args := os.Args[1:]
	y = year(today.Year())

	if len(args) == 0 {
		m = int(today.Month())
		return
	}

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
