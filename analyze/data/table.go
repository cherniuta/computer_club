package data

import "time"

type Table struct {
	tableID   int
	isBusy    bool
	timeStart time.Time
	allTime   time.Time
	price     int
	revenue   int
}

func NewTable(tableID int, price int) *Table {
	return &Table{
		tableID: tableID,
		isBusy:  false,
		price:   price}
}

func (t *Table) Busy(timeStart time.Time) {
	t.isBusy = true
	t.timeStart = timeStart
}

func (t *Table) Free(timeEnd time.Time) {
	t.isBusy = false

	difference := int(timeEnd.Sub(t.timeStart).Minutes())
	intervalHour := difference / 60
	intervalMinute := difference % 60
	t.allTime = t.allTime.Add(time.Hour * time.Duration(intervalHour)).Add(time.Minute * time.Duration(intervalMinute))
	t.revenue += intervalHour * (t.price)
	if intervalMinute != 0 {
		t.revenue += t.price
	}

}

func (t *Table) GetIsBusy() bool {
	return t.isBusy
}

func (t *Table) GetRevenue() int {
	return t.revenue
}

func (t *Table) GetTime() time.Time {
	return t.allTime
}
