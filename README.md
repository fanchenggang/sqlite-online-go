### 项目介绍

**通过golang 启动http服务 支持本地sqlite数据库文件操作**

```shell
#先编译前端
npm install
npm run build

#编译后端 window 
go build -o sqlite-online-go.exe
#交叉编译 linux
env GOOS=linux GOARCH=amd64 go build -o sqlite-online-go
# 启动服务 端口号 数据库文件路径
./sqlite-online-go.exe 86 ./my2.db

```
##### 创建demo数据
```sql

CREATE TABLE demo (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    age INTEGER DEFAULT 18 CHECK(age BETWEEN 0 AND 120),
    email TEXT NOT NULL UNIQUE,
    department TEXT CHECK(department IN ('IT', 'HR', 'Finance')),
    joined_date DATE DEFAULT (DATE('now')),
    salary REAL CHECK(salary > 0)
);
  
INSERT INTO demo (name, age, email, salary) 
VALUES ('Alice', 28, 'alice@example.com', 8500.50);
```

### 以下是原项目的README.md文件内容

<div align="center">

<a href="https://sqlite3.online/" target="_blank">
  <img src="https://github.com/user-attachments/assets/aef749bf-df08-4a84-8148-d34b796449d8" alt="SQLite Online Logo" width="128">
</a>

# SQLite Online

A fast, secure, and client-side SQLite editor that runs entirely in your browser—no server required. Built with WebAssembly, sql.js, and ReactJS, brings the full power of SQLite database management right to your browser.

[![License](https://img.shields.io/github/license/vwh/sqlite-online?label=License)](https://github.com/vwh/sqlite-online/blob/main/LICENSE)
[![Format Check](https://github.com/vwh/sqlite-online/actions/workflows/format.yml/badge.svg)](https://github.com/vwh/sqlite-online/actions/workflows/format.yml)
[![Lint](https://github.com/vwh/sqlite-online/actions/workflows/lint.yml/badge.svg)](https://github.com/vwh/sqlite-online/actions/workflows/lint.yml)

</div>

## Overview

- **Create database files**
- **Define, modify, and delete tables and indexes**
- **Browse, edit, add, and delete records**
- **Sort and search through data effortlessly**

SQLite Online provides a powerful yet accessible platform that requires zero installation. All processing is done within your client browser, ensuring your data remains private and secure.

## Features

- **Client Browser Processing:**  
  • No installation required – work directly in your browser  
  • All operations run on the client side – your data never leaves your device

- **Lightning Fast:**  
  • Powered by WebAssembly and enhanced with Web Workers for threading  
  • Efficient handling of large databases and heavy operations

- **Offline & PWA Support:**  
  • Fully functional without an internet connection  
  • Install as a desktop app on your mobile or desktop device

- **Comprehensive Database Management:**  
  • Create, compact, and manage database files  
  • Define, modify, and delete tables, indexes, and records  
  • Sort and search records with ease  
  • Import and export data (CSV, SQL dumps)  
  • Full SQL query support with command logging

## Keyboard Shortcuts (Hotkeys)

Streamline your workflow with these built-in hotkeys:

| Shortcut              | Action                  |
| --------------------- | ----------------------- |
| **Ctrl + s**          | Download the database   |
| **Ctrl + ArrowRight** | Go to the next page     |
| **Ctrl + ArrowUp**    | Jump to the first page  |
| **Ctrl + ArrowDown**  | Jump to the last page   |
| **Ctrl + ArrowLeft**  | Go to the previous page |
| **Ctrl + `**          | Close the edit panel    |
| **Ctrl + i**          | Trigger insert panel    |
| **Ctrl + I**          | Submit an insert edit   |
| **Ctrl + u**          | Submit an update edit   |
| **Ctrl + d**          | Submit a delete edit    |
| **Ctrl + q**          | Execute the SQL query   |

## License

SQLite Online is released under the [MIT License](https://github.com/vwh/sqlite-online/blob/main/LICENSE).


