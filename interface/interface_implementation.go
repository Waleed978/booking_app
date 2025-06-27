package interface_implementation

import "fmt"

func SendMessage(m Message) {
	fmt.Println(m.GetMessage())
}

type Message interface {
	GetMessage() string
}
type BirthDay struct {
	Date  string
	Month string
	Year  string
}

func (b BirthDay) GetMessage() string {
	return fmt.Sprintf("Happy Birthday %s %s %s", b.Date, b.Month, b.Year)
}

type ReportMessage struct {
	Report string
}

func (r ReportMessage) GetMessage() string {
	return fmt.Sprintf("Report: %s", r.Report)
}
func Test(m Message) {
	SendMessage(m)
	fmt.Println()
}

// func main() {

// 	test(
// 		birthDay{
// 			date:  "1",
// 			month: "1",
// 			year:  "2000",
// 		},
// 	)
// 	test(
// 		reportMessage{
// 			report: "Report",
// 		},
// 	)
// }
