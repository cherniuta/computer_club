package analyze

import (
	"bufio"
	"computer_club/analyze/data"
	"computer_club/collection"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	NotOpenYet       = "13 NotOpenYet"
	ICanWaitNoLonger = "13 ICanWaitNoLonger!"
	PlaceIsBusy      = "13 PlaceIsBusy"
	YouShallNotPass  = "13 YouShallNotPass"
	ClientUnknown    = "13 ClientUnknown"
)

func validateTime(time string) bool {
	timeFormat := regexp.MustCompile(`^([01]?[0-9]|2[0-3]):([0-5][0-9])$`)
	if timeFormat.MatchString(time) {
		return true
	} else {
		return false
	}
}

func validateInstruction(instructionID int) bool {
	singleDigitRegex := regexp.MustCompile(`^[1234]$`)
	if singleDigitRegex.MatchString(strconv.Itoa(instructionID)) {
		return true
	} else {
		return false
	}
}

func validateClient(client string) bool {
	alphanumericRegex := regexp.MustCompile(`^[\w-]+$`)
	if alphanumericRegex.MatchString(client) {
		return true
	} else {
		return false
	}
}

func validateTable(table int, countTable int) bool {
	naturalNumberRegex := regexp.MustCompile(`^[1-9]\d*$`)
	if naturalNumberRegex.MatchString(strconv.Itoa(table)) {
		if table > 0 && table <= countTable {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func Sort(clients []*data.Client) {
	n := len(clients)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if clients[i].GetClient() > clients[j].GetClient() {
				clients[i], clients[j] = clients[j], clients[i]
			}
		}
	}
}

func Analyze(file *os.File) {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	naturalNumberRegex := regexp.MustCompile(`^[1-9]\d*$`)
	if naturalNumberRegex.MatchString(scanner.Text()) == false {
		fmt.Println(scanner.Text())
		return
	}
	countTable, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	times := strings.Split(scanner.Text(), " ")
	if validateTime(times[0]) == false || validateTime(times[1]) == false {
		fmt.Println(strings.Join(times, " "))
		return
	}
	timeStart, _ := time.Parse("15:04", times[0])
	timeEnd, _ := time.Parse("15:04", times[1])

	scanner.Scan()
	if naturalNumberRegex.MatchString(scanner.Text()) == false {
		fmt.Println(scanner.Text())
		return
	}
	price, _ := strconv.Atoi(scanner.Text())

	tables := make([]*data.Table, 0)
	for index := 0; index < countTable; index++ {
		tables = append(tables, data.NewTable(index+1, price))
	}
	clients := make(map[string]*data.Client)

	queueClients := collection.Queue{}

	fmt.Println(timeStart.Format("15:04"))

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if len(line) < 3 {
			fmt.Println(strings.Join(line, " "))
			return
		}
		if validateTime(line[0]) == false {
			fmt.Println(strings.Join(line, " "))
			return
		}
		time, _ := time.Parse("15:04", line[0])

		instructionID, _ := strconv.Atoi(line[1])
		if validateInstruction(instructionID) == false {
			fmt.Println(strings.Join(line, " "))
			return
		}

		client := line[2]
		if validateClient(client) == false {
			fmt.Println(strings.Join(line, " "))
			return
		}

		var table int

		if len(line) == 4 {
			table, _ = strconv.Atoi(line[3])

			if validateTable(table, countTable) == false {
				fmt.Println(strings.Join(line, " "))
				return
			}

		} else if len(line) > 4 {
			fmt.Println(strings.Join(line, " "))
			return
		}

		fmt.Println(strings.Join(line, " "))
		switch instructionID {
		case 1:
			if len(line) > 3 {
				return
			}
			if int(time.Sub(timeStart).Minutes()) < 0 {
				fmt.Println(time.Format("15:04"), NotOpenYet)
				break
			}

			value, ok := clients[client]
			if ok {
				if value.GetIsInClub() == true {
					fmt.Println(time.Format("15:04"), YouShallNotPass)
				} else {
					value.Came()
				}
			} else {
				clients[client] = data.NewClient(client)
			}
		case 2:
			if len(line) != 4 {
				return
			}

			value, ok := clients[client]
			if ok {
				if value.GetIsInClub() == false {
					fmt.Println(time.Format("15:04"), ClientUnknown)
				} else {
					if tables[table-1].GetIsBusy() == false {
						if table != value.GetTable()-1 {
							if value.GetTable() != 0 {
								tables[value.GetTable()-1].Free(time)
							}
							value.TakeTable(table)
							tables[table-1].Busy(time)
						} else {
							fmt.Println(time.Format("15:04"), PlaceIsBusy)
						}

					} else {
						fmt.Println(time.Format("15:04"), PlaceIsBusy)
					}

				}

			} else {
				fmt.Println(time.Format("15:04"), ClientUnknown)
			}
		case 3:
			if len(line) != 3 {
				return
			}

			flag := true
			for index, _ := range tables {
				if tables[index].GetIsBusy() == false {
					fmt.Println(time.Format("15:04"), ICanWaitNoLonger)
					flag = false
					break
				}

			}

			if flag == false {
				break
			}
			if queueClients.Length() >= countTable {
				clients[client].Gone()
				fmt.Println(time.Format("15:04"), "11", client)
			} else {
				value, ok := clients[client]
				if ok {
					if value.GetIsInClub() == true {
						queueClients.Enqueue(value)
					} else {
						fmt.Println(time.Format("15:04"), ClientUnknown)
					}
				} else {
					fmt.Println(time.Format("15:04"), ClientUnknown)
				}

			}
		case 4:
			if len(line) != 3 {
				return
			}

			value, ok := clients[client]
			if ok {
				if value.GetIsInClub() == true {
					if value.GetTable() != 0 {
						//освобождаем стол
						tables[value.GetTable()-1].Free(time)
						//сажаем за него из очереди
						if queueClients.Length() != 0 {
							nextClient := queueClients.Dequeue()
							nextClient.TakeTable(value.GetTable())
							tables[value.GetTable()-1].Busy(time)
							fmt.Println(time.Format("15:04"), "12", nextClient.GetClient(), value.GetTable())
						}

					}
					value.Gone()
				} else {
					fmt.Println(time.Format("15:04"), ClientUnknown)
				}
			} else {
				fmt.Println(time.Format("15:04"), ClientUnknown)
			}
		}

	}

	notLeftClients := make([]*data.Client, 0)
	for index, _ := range clients {
		if clients[index].GetIsInClub() == true {
			notLeftClients = append(notLeftClients, clients[index])
		}
	}
	if len(notLeftClients) != 0 {
		Sort(notLeftClients)

		for index, _ := range notLeftClients {
			if notLeftClients[index].GetTable() != 0 {
				tables[notLeftClients[index].GetTable()-1].Free(timeEnd)
			}
			fmt.Println(timeEnd.Format("15:04"), "11", notLeftClients[index].GetClient())
		}
	}

	fmt.Println(timeEnd.Format("15:04"))
	for index := 0; index < countTable; index++ {
		fmt.Println(index+1, tables[index].GetRevenue(), tables[index].GetTime().Format("15:04"))
	}

}
