import random as r
import tkinter as tk
from tkinter import messagebox

# 创建主窗口
root = tk.Tk()
root.attributes('-topmost', True)   # 设置窗口在任务栏最上面弹出

num_people = int(input("请输入人数："))   # 获取用户输入的人数
elements = [str(i) for i in range(1, num_people + 1)]

def select_random_element():
    element = r.choice(elements)  # 随机选择一个元素
    messagebox.showinfo('随机选人', f"随机选择的人是：{element}")  # 弹出窗口显示随机元素
    elements.remove(element)  # 从列表中移除已选择的元素
    if len(elements) == 0:
        messagebox.showinfo('提示', '所有人已被抽过')
        button.config(state=tk.DISABLED)# 将按钮的状态设置为禁用

def on_mouse_down(event):
    root._offsetx = event.x
    root._offsety = event.y

def on_mouse_move(event):
    x = root.winfo_pointerx() - root._offsetx
    y = root.winfo_pointery() - root._offsety
    root.geometry(f"+{x}+{y}")

# 创建按钮
button = tk.Button(root, text="选择随机元素", command=select_random_element)
button.pack()

# 绑定鼠标事件
root.bind('<Button-1>', on_mouse_down)
root.bind('<B1-Motion>', on_mouse_move)

root.mainloop()  # 运行主循环
