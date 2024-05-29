import tkinter as tk
from ctypes import cdll, wintypes
import os

def main():
    root = tk.Tk()
    root.title("python-load-dll")
    root.geometry("400x300")

    current_directory = os.getcwd()
    libenergyPath = current_directory + "\libenergy.dll"
    print('libenergyPath: ', libenergyPath)

    # 加载dll句柄
    libenergy = cdll.LoadLibrary(libenergyPath)

    # 获取energy 导出的 api
    initCEFApplication = libenergy.initCEFApplication
    cefFormShow = libenergy.cefFormShow
    cefFormFree = libenergy.cefFormFree

    loadLibenergyBtn = tk.Button(root, text="加载libenergy.dll", command=lambda: load_dll_function())
    loadLibenergyBtn.pack()

    showWindowBtn = tk.Button(root, text="显示libenergy创建的窗口", command=lambda: show_window_function())
    showWindowBtn.pack()

    # 1.先加载dll
    def load_dll_function():
        print('call_dll_function')
        initCEFApplication()

    # 2. 显示窗口
    def show_window_function():
        print('show_window_function')
        cefFormShow()

    def on_closing():
        print('on_closing')
        cefFormFree()
        root.destroy()
    root.protocol("WM_DELETE_WINDOW", on_closing)

    # 进入主事件循环
    root.mainloop()

if __name__ == "__main__":
    main()