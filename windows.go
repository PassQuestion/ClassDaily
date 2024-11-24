/*
川大附中IDEA课程节孵化项目：电子值日生（1.0）
开发者：
windows.go 彭思齐 约680行有效代码
自动轮换(节点2).py 符晋豪  41行有效代码 编译为changeseat.exe
基础随机分组（节点1）.py 符晋豪 25行有效代码 编译为selectseat.exe
随机抽人（节点2）.py 符晋豪（初版） 彭致远（后续UI修改） 初版共35行有效代码 编译为Random.exe
管理者： 刘芸朵 宋羿延
徽标设计： 刘芸朵
UI设计：黄嘉琪
记录员：胡懿轩 华胥睿
测试员：彭致远 许锡川
文档组：华胥睿 查祉旭 刘芸朵 卓钰轩
*/
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/beevik/etree"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
)

type hwnd struct {
	*walk.MainWindow
}

func get_rand(max int) int64 {
	rand.New(rand.NewSource(time.Now().UnixNano())) //随机数种子
	random_number := rand.Intn(max + 1)             // 生成最大为 输入max的随机数
	for random_number == 0 {                        //防0模块
		random_number = rand.Intn(max + 1)
	}
	return int64(random_number) //rand必须返回int64类型
}
func get_date() int { //查找今天星期几
	current := time.Now()
	date := current.Weekday()
	return int(date)
}
func random_window() { //随机数生成窗口，现在已被同伴的python编译程序（即第597行的Random.exe）替代
	var students_number_line *walk.LineEdit
	students_number := LineEdit{
		AssignTo:  &students_number_line,
		Text:      "请输入学生总数",
		TextColor: walk.Color(walk.RGB(100, 100, 100)),
		MaxLength: 5,
	} //使用lineedit而不是numberedit是为了兼容字符串输入
	var rand int64
	var resultpointer *walk.TextEdit
	randombutton := PushButton{
		Text: "生成",
		OnClicked: func() { //按钮点击 ——> 生成随机数（string转int进入函数get_rand，int64转回string输出）
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
		}.Create()

	}

}
func set_class(date int, class int, name string) { //set_class函数，第二阶段开发的核心内容 缺点：太多switch
	classlist := etree.NewDocument() //etree库开始使用，这两句代码是使用etree 加载class.xml文件
	if err := classlist.ReadFromFile("./class.xml"); err != nil {
		walk.MsgBox(walk.App().ActiveForm(), "无法加载课表", "课表加载失败。请保证课表文件（class.xml）加载正确。", walk.MsgBoxIconError) //！！！class.xml文件不存在就会报错，循环报错，只能kill
	}
	rootelement := classlist.SelectElement("root") //设置根节点
	dayselement := rootelement.FindElement("Days") //设置节点days，这一部分应参考class.xml文件的格式
	var dayelement *etree.Element                  //设置节点day
	switch date {                                  //将date对应星期1234567

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
	switch class { //class对应123456789节课
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
	classelement.SetText(name) //将name修改
	file, openerr := os.OpenFile("./class.xml", os.O_RDWR, 0)
	if openerr != nil {
		walk.MsgBox(walk.App().ActiveForm(), "写入错误", "修改错误。请检查class.xml是否存在。", walk.MsgBoxIconError)
		return
	}
	resultstring, err := classlist.WriteToString()
	if err != nil {
		walk.MsgBox(walk.App().ActiveForm(), "写入错误", "修改错误", walk.MsgBoxIconError)
	}
	file.WriteString(resultstring) //写入文件，多次点击会出问题
}
func setting_window() { //setting_window函数，第二阶段的成果 启动在第589行里有定义 格式：星期 x 的第 y 节课更改为 z
	wd3ptr, err := walk.NewMainWindow()
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
	a := []string{" 语文 ", " 数学 ", " 英语 ", " 物理 ", " 化学 ", " 生物 ", " 政治 ", " 通用技术 ", " 信息技术 ", " 音乐^美术 ", " 自习 ", " 校本课程 ", " 班会 "} //一共13节课，地理，历史未包含 这里字符串的空格是兼容class.xml
	namebox := ComboBox{
		Model:        []string{"语文", "数学", "英语", "物理", "化学", "生物", "政治", "通用技术", "信息技术", "音乐^美术", "自习", "校本课程", "班会"},
		CurrentIndex: 0,
		AssignTo:     &nameboxptr,
	}
	surebutton := PushButton{
		Text: "确定",
		OnClicked: func() {
			set_class(daylistboxptr.CurrentIndex()+1, classlistboxptr.CurrentIndex()+1, a[nameboxptr.CurrentIndex()]) //调用set_class函数
			wd3ptr.Close()                                                                                            //将窗口关闭以免再次点击
		},
	}
	widget3 := []Widget{
		daylabel, daylistbox, classlabel, classlistbox, changelabel, namebox, surebutton,
	}
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

/*布置作业模块 第三阶段开发的核心内容，但是控件的指针调用总出问题，希望前辈能够将指针修改正确使这一功能运行*/
type homeworkitem struct { //定义struct类型，对应每一项作业 格式：“作业项” n页 备注：备注内容（如不写第xx题等）
	checkbox    CheckBox
	checkboxptr **walk.CheckBox
	pagebox     ComboBox
	pageboxptr  **walk.ComboBox
	noteline    LineEdit
	notelineptr **walk.LineEdit
	pagelabel   Label
	widget      []Widget
	checkstatus bool
}

func create_homework_item(name string) homeworkitem {
	bl := false
	homeworkcheckptr := new(*walk.CheckBox)
	homeworkbox := CheckBox{
		AssignTo:       homeworkcheckptr,
		TextOnLeftSide: true,
		Text:           name,
		Name:           name,
		Checked:        false,
	}
	pageboxptr := new(*walk.ComboBox)
	pagebox := ComboBox{
		AssignTo:     pageboxptr,
		CurrentIndex: 1,
		Model:        []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20"}, //定义最高20页页数
	}
	notelineptr := new(*walk.LineEdit)
	noteline := LineEdit{
		AssignTo: notelineptr,
		ReadOnly: false,
	}
	widget := []Widget{
		homeworkbox, pagebox, noteline,
	}
	item := homeworkitem{
		checkbox:    homeworkbox,
		checkboxptr: homeworkcheckptr,
		pagebox:     pagebox,
		pageboxptr:  pageboxptr,
		noteline:    noteline,
		notelineptr: notelineptr,
		pagelabel:   Label{Text: "页 备注："},
		widget:      widget,
		checkstatus: bl,
	}
	return item
}

type testpaper struct {
	testpaperbox    CheckBox
	pagebox         ComboBox
	label           Label
	noteline        LineEdit
	testpaperboxptr **walk.CheckBox
	paperpageptr    **walk.ComboBox
	notelineptr     **walk.LineEdit
}

func create_testpaper() testpaper {
	testpaperboxptr := new(*walk.CheckBox)
	testpaperbox := CheckBox{
		AssignTo: testpaperboxptr,
		Text:     "试卷",
		Checked:  false,
	}
	paperpageptr := new(*walk.ComboBox)
	pagebox := ComboBox{
		AssignTo: paperpageptr,
		Model:    []string{"1", "2", "3", "4", "5", "6"},
	}
	label := Label{
		Text: "张 备注：",
	}
	notelineptr := new(*walk.LineEdit)
	noteline := LineEdit{
		AssignTo: notelineptr,
		ReadOnly: false,
	}
	item := testpaper{
		testpaperbox:    testpaperbox,
		testpaperboxptr: testpaperboxptr,
		pagebox:         pagebox,
		paperpageptr:    paperpageptr,
		label:           label,
		noteline:        noteline,
		notelineptr:     notelineptr,
	}
	return item
}
func sub_to_string(subject int) string {
	switch subject {
	case 1:
		{
			return "语文"
		}
	case 2:
		{
			return "数学"
		}
	case 3:
		{
			return "英语"
		}
	case 4:
		{
			return "物理"
		}
	case 5:
		{
			return "化学"
		}
	case 6:
		{
			return "生物"
		}
	case 7:
		{
			return "政治"
		}
	}
	return "其他"
}

var checkboxptr = new(walk.CheckBox)
var pageboxptr = new(walk.ComboBox)
var notelineptr = new(walk.LineEdit)

func homework_window(subject int) {
	wd5ptr, err := walk.NewMainWindow()
	dynaticwidget := []Widget{}
	homeworkitems := make([]homeworkitem, 0)
	var widgetitem []Widget
	var hsp HSplitter
	h := 0
	/* 添加作业项 */
	if subject == 0 {
		homeworkitems = append(homeworkitems, create_homework_item("高考调研"))
		h += 40
		homeworkitems = append(homeworkitems, create_homework_item("能力风暴"))
		h += 40
		homeworkitems = append(homeworkitems, create_homework_item("晨读晚练"))
		h += 40 //h即窗口高度，每添加一项作业加40像素
	} else if subject == 1 {
		homeworkitems = append(homeworkitems, create_homework_item("课时精练"))
		h += 40
	} else if subject == 2 {
		homeworkitems = append(homeworkitems, create_homework_item("高考英语 天天练"))
		h += 40
		homeworkitems = append(homeworkitems, create_homework_item("一线课堂"))
		h += 40
	} else if subject == 3 {
		homeworkitems = append(homeworkitems, create_homework_item("步步高 学习笔记（物理）"))
		h += 40
		homeworkitems = append(homeworkitems, create_homework_item("步步高 练透（物理）"))
		h += 40
		homeworkitems = append(homeworkitems, create_homework_item("试吧（化学）"))
		h += 40
	} else if subject == 4 {
		homeworkitems = append(homeworkitems, create_homework_item("步步高 学习笔记（化学）"))
		h += 40
		homeworkitems = append(homeworkitems, create_homework_item("步步高 练透（化学）"))
		h += 40
		homeworkitems = append(homeworkitems, create_homework_item("试吧（化学）"))
		h += 40
	} else if subject == 5 {
		homeworkitems = append(homeworkitems, create_homework_item("步步高 学习笔记（生物）"))
		h += 40
		homeworkitems = append(homeworkitems, create_homework_item("步步高 练透（生物）"))
		h += 40
	} // 添加作业结束
	for i := 0; i < len(homeworkitems); i += 1 {
		widgetitem = []Widget{
			homeworkitems[i].checkbox, homeworkitems[i].pagebox, homeworkitems[i].pagelabel, homeworkitems[i].noteline,
		}
		hsp = HSplitter{
			Children:  widgetitem,
			Alignment: AlignHCenterVCenter,
		}
		dynaticwidget = append(dynaticwidget, hsp)
		h += 40
	}
	paper := create_testpaper()
	widgetitem2 := []Widget{
		paper.testpaperbox, paper.pagebox, paper.label, paper.noteline,
	}
	hsppaper := HSplitter{
		Children: widgetitem2,
	}
	dynaticwidget = append(dynaticwidget, hsppaper)
	surebutton := PushButton{
		Text: "确定",
		OnClicked: func() {
			homeworkfile, err1 := os.OpenFile("./homework", os.O_APPEND, 0644)
			if err1 != nil {
				walk.MsgBox(walk.App().ActiveForm(), "作业文件错误", "请创建一个名为homework的文件。", walk.MsgBoxOK)
				return
			}
			defer homeworkfile.Close()
			var homeworkstr string
			var page string
			homeworkstr = sub_to_string(subject+1) + "\n"
			for j := 0; j < len(homeworkitems); j += 1 {
				checkboxptr = *homeworkitems[j].checkboxptr
				pageboxptr = *homeworkitems[j].pageboxptr
				notelineptr = *homeworkitems[j].notelineptr
				if checkboxptr.Checked() { //如果对应作业的checkbox被选中
					homeworkstr += homeworkitems[j].checkbox.Name
					page = fmt.Sprintf("%d", pageboxptr.CurrentIndex()+1)
					homeworkstr += page
					homeworkstr += "页 "
					homeworkstr += notelineptr.Text()
					homeworkstr += "\n" //则加字符串：格式为 {作业} {n}页 {备注} 换行
				}
			}
			paperptr := *paper.testpaperboxptr

			if paperptr.Checked() {
				homeworkstr += "试卷"
				paperpage := *paper.paperpageptr
				page = fmt.Sprintf("%d", paperpage.CurrentIndex()+1)
				homeworkstr += page
				homeworkstr += "张 "
				notelineptr = *paper.notelineptr
				homeworkstr += notelineptr.Text()
				homeworkstr += "\n"
			}
			writer := bufio.NewWriter(homeworkfile)
			writer.WriteString(homeworkstr)
			writer.Flush()
			wd5ptr.Close()
		},
	}
	dynaticwidget = append(dynaticwidget, surebutton)
	if err == nil {
		MainWindow{
			AssignTo: &wd5ptr,
			Title:    "作业布置",
			Size:     Size{500, h},
			Layout:   VBox{},
			Children: dynaticwidget,
		}.Run()
	}

}
func seat_window() {
	seatwinptr := new(walk.MainWindow)
	num1ptr := new(walk.NumberEdit)
	num1 := NumberEdit{
		AssignTo: &num1ptr,
	}
	labelline := Label{Text: "行"}
	num2ptr := new(walk.NumberEdit)
	num2 := NumberEdit{
		AssignTo: &num2ptr,
	}
	labelcolumn := Label{Text: "列共"}
	num3ptr := new(walk.NumberEdit)
	num3 := NumberEdit{
		AssignTo: &num3ptr,
	}
	labelpeople := Label{Text: "人"}
	callbutton := PushButton{
		Text: "确定",
		OnClicked: func() {
			line := fmt.Sprintf("%d", int64(num1ptr.Value()))
			column := fmt.Sprintf("%d", int64(num2ptr.Value()))
			people := fmt.Sprintf("%d", int64(num3ptr.Value()))
			fmt.Println(line, column, people)
			cmd2 := exec.Command("./selectseat.exe", people, line, column)
			cmd2.Run()
			seatwinptr.Close()
		},
	}
	widget6 := []Widget{
		num1, labelline, num2, labelcolumn, num3, labelpeople, callbutton,
	}
	wd6 := MainWindow{
		AssignTo: &seatwinptr,
		Title:    "创建座位",
		Size:     Size{200, 100},
		Layout:   HBox{},
		Children: widget6,
	}
	wd6.Run()
}
func get_class(date int, class int) string { //get_class函数，第一阶段开发的核心内容，和set_class函数的查找部分思路一致
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
func create_label(classes int) Label {
	return Label{
		Text:          get_class(get_date(), classes),
		TextAlignment: AlignCenter,
	}
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

func choose_homework_window() {
	chooselabel := Label{
		Text: "请选择科目：",
	}
	var subjectboxptr *walk.ComboBox
	subjectbox := ComboBox{
		AssignTo: &subjectboxptr,
		Model:    []string{"语文", "数学", "英语", "物理", "化学", "生物", "政治"},
	}
	var newwinbuttonptr *walk.PushButton
	newwinbutton := PushButton{
		AssignTo:  &newwinbuttonptr,
		Text:      "确定",
		OnClicked: func() { homework_window(subjectboxptr.CurrentIndex()) },
	}
	widget4 := []Widget{
		chooselabel, subjectbox, newwinbutton,
	}

	var homework_windowptr *walk.MainWindow
	choose_homework_window := MainWindow{
		AssignTo: &homework_windowptr,
		Size:     Size{100, 100},
		Title:    "选择课程",
		Layout:   VBox{},
		Children: widget4,
	}
	choose_homework_window.Run()

}
func backcode() { //提醒老师下课的后台函数
	for {
		nowhour, nowmin, nowsecond := time.Now().Clock()
		if nowhour == 10 && nowmin == 15 && nowsecond == 00 {
			walk.MsgBox(walk.App().ActiveForm(), "下课了", "老师，这节课下楼做操，请您尽早下课", walk.MsgBoxOK)
		}
	}

}

/*定义主窗口大小*/
const (
	WIDTH  = 320
	HEIGHT = 1020
)

func main() {
	width := win.GetSystemMetrics(win.SM_CXSCREEN) / 6
	height := win.GetSystemMetrics(win.SM_CYSCREEN) / 20 * 19
	datefilestring, err := os.ReadFile("./date")
	if err != nil {
		walk.MsgBox(walk.App().ActiveForm(), "日期文件错误", "请创建一个名为date的文件。", walk.MsgBoxOK)
		return
	}
	y, m, d := time.Now().Date()
	datestring := fmt.Sprintf("%d%d%d", y, m, d)
	if string(datefilestring) != datestring {
		os.Create("./homework")
		os.Create("./date") //若点击按钮日期等于上次点击按钮日期，则清空上次文件（create方法）

		os.WriteFile("./date", []byte(datestring), 0643)
	}

	go backcode()
	str1 := "星期"
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
		Font: Font{
			"宋体", 20, true, false, false, false,
		},
		Alignment: AlignHCenterVCenter,
	}
	label1 := create_label(1)
	label2 := create_label(2)
	label3 := create_label(3)
	label4 := create_label(4)
	label5 := create_label(5)
	label6 := create_label(6)
	label7 := create_label(7)
	label8 := create_label(8)
	label9 := create_label(9)

	/* Settingbutton := PushButton{
		Text: "设置",
		OnClicked: func() {
			fmt.Println("clicked")
			setting_window()
		},
	}
	Randombutton := PushButton{
		Text:      "随机数",
		OnClicked: func() { random_window() },
	} */
	/* homeworkbutton := PushButton{
		Text:      "作业布置",
		OnClicked: func() { choose_homework_window() },
	} */
	widget := []Widget{
		date, label1, label2, label3, label4, label5, label6, label7, label8, label9, /*Settingbutton, Randombutton,  homeworkbutton,*/
	}
	menuitem := []MenuItem{
		Menu{
			Text: "File",
			Items: []MenuItem{
				Action{
					Text:        "修改课表...",
					OnTriggered: setting_window,
				},
				Action{
					Text: "随机排座位...",
					OnTriggered: func() {
						seat_window()
					},
				},
				Action{
					Text: "轮换座位...",
					OnTriggered: func() {
						cmd3 := exec.Command("./changeseat.exe")
						cmd3.Run()
					},
				},
			},
		},
		Menu{
			Text: "Function",
			Items: []MenuItem{
				Action{
					Text:        "布置作业...",
					OnTriggered: choose_homework_window,
				},
				Action{
					Text: "随机数",
					OnTriggered: func() {
						cmd := exec.Command("./Random.exe")
						cmd.Run()
					},
				},
			},
		},
	}
	rootwindow.MainWindow = new(walk.MainWindow)
	MainWindow{
		AssignTo: &rootwindow.MainWindow,
		Title:    "电子值日生",
		Size:     Size{int(width), int(height)},
		Layout:   VBox{},
		Font: Font{
			"宋体", 16, false, false, false, false,
		},
		MenuItems: menuitem,
		Children:  widget,
		Icon:      icon,
	}.Create()
	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN)
	win.SetWindowPos(
		rootwindow.Handle(),
		win.HWND_BOTTOM,
		xScreen*5/6, 0,
		width,
		height,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(rootwindow.Handle(), win.SW_SHOW)
	rootwindow.Run()
}
