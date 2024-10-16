package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/beevik/etree"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

const (
	WIDTH  = 320
	HEIGHT = 820
)

type hwnd struct {
	*walk.MainWindow
}

func get_rand(max int) int64 {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	random_number := rand.Intn(max)
	for random_number == 0 {
		random_number = rand.Intn(max)
	}
	return int64(random_number)
}
func get_date() int {
	current := time.Now()
	date := current.Weekday()
	return int(date)
}
func random_window() {
	type TE struct {
		*walk.TextEdit
	}
	var students_number_line *walk.LineEdit
	students_number := LineEdit{
		AssignTo:  &students_number_line,
		Text:      "请输入学生总数",
		TextColor: walk.Color(walk.RGB(100, 100, 100)),
		MaxLength: 5,
	}
	var rand int64
	var resultpointer *walk.TextEdit
	randombutton := PushButton{
		Text: "生成",
		OnClicked: func() {
			maxstudent, err := strconv.Atoi(students_number_line.Text())
			if err == nil {
				rand = get_rand(maxstudent)
				result := strconv.FormatInt(rand, 10)
				resultpointer.SetText(result)
			}
		},
	}

	resulttext := TextEdit{
		AssignTo: &resultpointer,
	}
	widget2 := []Widget{
		students_number, randombutton, resulttext,
	}
	window2pointer, e := walk.NewMainWindow()
	if e == nil {
		MainWindow{
			AssignTo: &window2pointer,
			Title:    "随机数",
			Layout:   VBox{},
			Size:     Size{300, 200},
			Children: widget2,
		}.Run()
		fmt.Println(window2pointer.Visible())
	}

}
func set_class(date int, class int, name string) {
	classlist := etree.NewDocument()
	if err := classlist.ReadFromFile("./class.xml"); err != nil {
		walk.MsgBox(walk.App().ActiveForm(), "无法加载课表", "课表加载失败。请保证课表文件（class.xml）加载正确。", walk.MsgBoxIconError)
	}
	rootelement := classlist.SelectElement("root")
	dayselement := rootelement.FindElement("Days")
	var dayelement *etree.Element
	switch date {

	case int(time.Monday):
		dayelement = dayselement.FindElement("Day[@ID='d1']")
	case int(time.Tuesday):
		dayelement = dayselement.FindElement("Day[@ID='d2']")
	case int(time.Wednesday):
		dayelement = dayselement.FindElement(`./Day[@ID="d3"]`)
	case int(time.Thursday):
		dayelement = dayselement.FindElement(`./Day[@ID="d4"]`)
	case int(time.Friday):
		dayelement = dayselement.FindElement(`./Day[@ID="d5"]`)

	}
	var classelement *etree.Element
	switch class {
	case 1:
		classelement = dayelement.FindElement("./Class[@ID='l1']/name")
	case 2:
		classelement = dayelement.FindElement("./Class[@ID='l2']/name")
	case 3:
		classelement = dayelement.FindElement("./Class[@ID='l3']/name")
	case 4:
		classelement = dayelement.FindElement("./Class[@ID='l4']/name")
	case 5:
		classelement = dayelement.FindElement("./Class[@ID='l5']/name")
	case 6:
		classelement = dayelement.FindElement("./Class[@ID='l6']/name")
	case 7:
		classelement = dayelement.FindElement("./Class[@ID='l7']/name")
	case 8:
		classelement = dayelement.FindElement("./Class[@ID='l8']/name")
	case 9:
		classelement = dayelement.FindElement("./Class[@ID='l9']/name")

	}
	classelement.SetText(name)
	file, openerr := os.OpenFile("./class.xml", os.O_RDWR, 0)
	if openerr != nil {
		walk.MsgBox(walk.App().ActiveForm(), "写入错误", "修改错误。请检查class.xml是否存在。", walk.MsgBoxIconError)
		return
	}
	resultstring, err := classlist.WriteToString()
	if err != nil {
		walk.MsgBox(walk.App().ActiveForm(), "写入错误", "修改错误", walk.MsgBoxIconError)
	}
	file.WriteString(resultstring)
}
func setting_window() {

	daylabel := Label{
		Text:      "星期",
		Alignment: AlignHCenterVCenter,
	}
	var daylistboxptr *walk.ComboBox
	daylistbox := ComboBox{
		Model:        []string{"一", "二", "三", "四", "五"},
		CurrentIndex: 0,
		AssignTo:     &daylistboxptr,
	}
	classlabel := Label{
		Text:      "的第",
		Alignment: AlignHCenterVCenter,
	}
	var classlistboxptr *walk.ComboBox
	classlistbox := ComboBox{
		Model:        []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
		CurrentIndex: 0,
		AssignTo:     &classlistboxptr,
	}
	changelabel := Label{
		Text: "节课更改为",
	}
	var nameboxptr *walk.ComboBox
	a := []string{" 语文 ", " 数学 ", " 英语 ", " 物理 ", " 化学 ", " 生物 ", " 政治 ", " 通用技术 ", " 信息技术 ", " 音乐^美术 ", " 自习 ", " 校本课程 ", " 班会 "}
	namebox := ComboBox{
		Model:        []string{"语文", "数学", "英语", "物理", "化学", "生物", "政治", "通用技术", "信息技术", "音乐^美术", "自习", "校本课程", "班会"},
		CurrentIndex: 0,
		AssignTo:     &nameboxptr,
	}
	surebutton := PushButton{
		Text: "确定",
		OnClicked: func() {
			fmt.Println(nameboxptr.DisplayMember())
			fmt.Println(nameboxptr.BindingMember())
			set_class(daylistboxptr.CurrentIndex()+1, classlistboxptr.CurrentIndex()+1, a[nameboxptr.CurrentIndex()])
		},
	}
	widget3 := []Widget{
		daylabel, daylistbox, classlabel, classlistbox, changelabel, namebox, surebutton,
	}
	wd3ptr, err := walk.NewMainWindow()
	if err == nil {
		wd3 := MainWindow{
			AssignTo: &wd3ptr,
			Title:    "设置",
			Layout:   HBox{},
			Size:     Size{500, 100},
			Children: widget3,
		}
		wd3.Run()
	}

}
func get_class(date int, class int) string {
	classlist := etree.NewDocument()
	if err := classlist.ReadFromFile("./class.xml"); err != nil {
		walk.MsgBox(walk.App().ActiveForm(), "无法加载课表", "课表加载失败。请保证课表文件（class.xml）加载正确。", walk.MsgBoxOK)
	}
	rootelement := classlist.SelectElement("root")
	rootelement = rootelement.FindElement("Days")
	var dayelement *etree.Element
	switch date {
	case int(time.Sunday):
		return "今日无课"
	case int(time.Monday):
		dayelement = rootelement.FindElement("Day[@ID='d1']")
	case int(time.Tuesday):
		dayelement = rootelement.FindElement("Day[@ID='d2']")
	case int(time.Wednesday):
		dayelement = rootelement.FindElement(`./Day[@ID="d3"]`)
	case int(time.Thursday):
		dayelement = rootelement.FindElement(`./Day[@ID="d4"]`)
	case int(time.Friday):
		dayelement = rootelement.FindElement(`./Day[@ID="d5"]`)
	case int(time.Saturday):
		return "今日无课"

	}
	var classelement *etree.Element
	switch class {
	case 1:
		classelement = dayelement.FindElement("./Class[@ID='l1']/name")
	case 2:
		classelement = dayelement.FindElement("./Class[@ID='l2']/name")
	case 3:
		classelement = dayelement.FindElement("./Class[@ID='l3']/name")
	case 4:
		classelement = dayelement.FindElement("./Class[@ID='l4']/name")
	case 5:
		classelement = dayelement.FindElement("./Class[@ID='l5']/name")
	case 6:
		classelement = dayelement.FindElement("./Class[@ID='l6']/name")
	case 7:
		classelement = dayelement.FindElement("./Class[@ID='l7']/name")
	case 8:
		classelement = dayelement.FindElement("./Class[@ID='l8']/name")
	case 9:
		classelement = dayelement.FindElement("./Class[@ID='l9']/name")

	}
	return classelement.Text()

}

func date_to_string(get_date int) string {
	if get_date == int(time.Sunday) {
		return "日"
	} else if get_date == int(time.Monday) {
		return "一"
	} else if get_date == int(time.Tuesday) {
		return "二"
	} else if get_date == int(time.Wednesday) {
		return "三"
	} else if get_date == int(time.Thursday) {
		return "四"
	} else if get_date == int(time.Friday) {
		return "五"
	} else if get_date == int(time.Saturday) {
		return "六"
	}
	return "时间错误"
}
func backcode() {
	for {
		nowhour, nowmin, nowsecond := time.Now().Clock()
		if nowhour == 10 && nowmin == 15 && nowsecond == 00 {
			walk.MsgBox(walk.App().ActiveForm(), "下课了", "老师，这节课下楼做操，请您尽早下课", walk.MsgBoxOK)
		}
	}

}
func main() {
	go backcode()
	str1 := "今天星期"
	str2 := date_to_string(get_date())
	var stringBuilder bytes.Buffer
	stringBuilder.WriteString(str1)
	stringBuilder.WriteString(str2)
	rootwindow := new(hwnd)
	icon, iconerr := walk.NewIconFromFile("./icon.ico")
	if iconerr != nil {
		icon = walk.IconApplication()
	}
	a := stringBuilder.String()
	date := Label{
		Text: a,
	}
	label1 := Label{
		Text:          get_class(get_date(), 1),
		TextAlignment: AlignCenter,
	}

	label2 := Label{
		Text:          get_class(get_date(), 2),
		TextAlignment: AlignCenter,
	}

	label3 := Label{
		Text:          get_class(get_date(), 3),
		TextAlignment: AlignCenter,
	}

	label4 := Label{
		Text:          get_class(get_date(), 4),
		TextAlignment: AlignCenter,
	}

	label5 := Label{
		Text:          get_class(get_date(), 5),
		TextAlignment: AlignCenter,
	}

	label6 := Label{
		Text:          get_class(get_date(), 6),
		TextAlignment: AlignCenter,
	}

	label7 := Label{
		Text:          get_class(get_date(), 7),
		TextAlignment: AlignCenter,
	}

	label8 := Label{
		Text:          get_class(get_date(), 8),
		TextAlignment: AlignCenter,
	}

	label9 := Label{
		Text:          get_class(get_date(), 9),
		TextAlignment: AlignCenter,
	}

	Settingbutton := PushButton{
		Text: "设置",
		OnClicked: func() {
			fmt.Println("clicked")
			setting_window()
		},
	}
	Randombutton := PushButton{
		Text:      "随机数",
		OnClicked: func() { random_window() },
	}
	widget := []Widget{
		date, label1, label2, label3, label4, label5, label6, label7, label8, label9, Settingbutton, Randombutton,
	}
	wd1 := MainWindow{
		AssignTo: &rootwindow.MainWindow,
		Title:    "电子值日生",
		Size:     Size{WIDTH, HEIGHT},
		Layout:   VBox{},
		Font: Font{
			"微软雅黑", 20, false, false, false, false,
		},
		Children: widget,
		Icon:     icon,
	}
	wd1.Run()
}
