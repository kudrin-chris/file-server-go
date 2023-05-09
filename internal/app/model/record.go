package model

import "strconv"

type Record struct {
	Id         int
	N          int
	Mqtt       string
	Invid      string
	Msg_id     string
	Text       string
	Context    string
	Class      string
	Level      int
	Area       string
	Addr       string
	Block      string
	Type_      string
	Bit_       string
	Invert_bit string
}

func ConvertString(r Record) string {
	return strconv.Itoa(r.Id) + "\t\t\t" + strconv.Itoa(r.N) + "\t\t\t" + r.Mqtt + "\t\t\t" + r.Invid + "\t\t\t" + r.Msg_id + "\t\t\t" + r.Text + "\t\t\t" + r.Context + "\t\t\t" + r.Class + "\t\t\t" + strconv.Itoa(r.Level) + "\t\t\t" + r.Area + "\t\t\t" + r.Addr + "\t\t\t" + r.Block + "\t\t\t" + r.Type_ + "\t\t\t" + r.Bit_ + "\t\t\t" + r.Invert_bit
}
