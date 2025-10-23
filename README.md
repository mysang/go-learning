# 🐹 Go Learning

A structured Go learning repository with lessons, exercises, and reusable utilities — designed to help developers learn Go step by step through hands-on coding and clean examples.

---

## 📘 Overview

This repository aims to provide a practical learning path for mastering the Go programming language.
It includes **lessons** that explain concepts with examples, **exercises** to test your understanding, and **utility functions** that you can reuse in your own projects.

---

## 📂 Repository Structure

```plaintext
go-learning/
│
├── lessons/          # Step-by-step theory with examples
│   ├── 01_variables/
│   │   ├── README.md
│   │   └── main.go
│   ├── 02_structs/
│   │   ├── README.md
│   │   └── main.go
│   └── ...
│
├── exercises/        # Coding exercises with questions and solutions
│   ├── 01_student_struct/
│   │   ├── question.md
│   │   ├── solution.go
│   │   └── explanation.md
│   └── ...
│
└── utils/            # Reusable helper code
    ├── math_utils.go
    ├── string_utils.go
    └── file_utils.go
```

---

## 🐳 Run with Docker

You don’t need to install Go locally — this repository includes a ready-to-use Docker environment for learning Go.

### 1️⃣ Start the container

```bash
docker compose up
```

The first time you run this command, Docker will automatically build the image and start the container.
Once it’s running, you’ll have a Go environment ready for use.

---

### 2️⃣ Open an interactive shell inside the container

In another terminal window, run:

```bash
docker exec -it go-learning bash
```

Now you can execute Go commands inside the container.

---

### 3️⃣ Run any Go file manually

Inside the container, choose the lesson or exercise you want to run:

```bash
# Run a lesson
go run ./lessons/01_variables

# Run an exercise
go run ./exercises/01_student_struct/solution.go

# Format or test your code
go fmt ./...
go test ./...
```

Each folder is independent — you can explore, modify, and run each lesson or exercise as you learn.

---

## 🧠 Learning Goals

By following this repository, you will:

- Understand Go fundamentals (variables, functions, structs, interfaces, concurrency)
- Practice real-world problem solving with Go
- Learn to write clean, idiomatic Go code
- Build a reusable codebase for your future Go projects

---

## 💡 Contribution

Contributions are welcome!
If you’d like to add new lessons or exercises, please:

1. Fork the repository
2. Follow the existing folder and naming conventions
3. Submit a pull request

---

## 📄 License

This project is licensed under the **MIT License** — you’re free to learn, use, and share it.
