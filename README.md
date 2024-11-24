# ClassDaily——使班级管理更简单
## 安装 Installization
> 在安装之前，请确保您已经安装了Python pip 模块。
- changeseat.exe 因大小问题未上传至，请使用以下命令安装：
```cmd
pip install pandas openpyxl tkinter pyinstaller
pyinstaller 自动轮换（节点3）.py
mv .\dist\自动轮换（节点3）.exe .\changeseat.exe
```
- Project.exe 为主文件，点击即可运行。
## 使用 Usage
以下是主窗口：
![主窗口](https://github.com/PassQuestion/ClassDaily/img/Main.png)
- 可以看到主窗口中的课表以及菜单栏，菜单栏中有File，Function两个子菜单
    - File: **修改课表**,**随机排座位**，**轮换座位**功能
    - Function： **布置作业**，**随机数**功能
#### 修改课表功能
点击后，修改课表窗口弹出。
![修改课表窗口](https://github.com/PassQuestion/ClassDaily/img/Setting.png)
选择三个下拉栏，使其可以修改您所想要修改的课节。
点击确定后，程序会自动重新启动。
#### 随机排座位功能
点击后，输入行与列的信息，点击“确定”，在当前文件夹中会建立matrix.xlsx文件，以学号代替学生的座位进行编排。
可以使用外部程序（MicroSoft Office , WPS , LibreOffice）打开。
#### 轮换座位功能
点击后，可以按照某种特定的顺序轮换matrix.xlsx中储存的学生座位。
#### 布置作业功能
点击后，选择作业学科窗口弹出。
![选择作业学科窗口](https://github.com/PassQuestion/ClassDaliy/img/Choosehomework.png)
选择学科，点击“确定”。
![选择作业窗口](https://github.com/PassQuestion/ClassDaily/img/Homework.png)
通过选择对应的作业项，布置对应学科的作业。
#### 随机数功能、
点击后，随机数生成器窗口弹出，且始终位于其他程序上方，便于随机数的生成。
![随机数窗口](https://github.com/PassQuestion/ClassDaily/img/Random.png)
> 默认状态：最大值 59 最小值 1 人数 1 
