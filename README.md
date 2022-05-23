# Project Title

## Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)
- [Todo](#todo)

## About <a name = "about"></a>

It is a simple quiz game written in Go. This CLI based game was written to learn the Go language and play around with the language. This project uses some basic Go packages, and demonstrates the simplicity of implementing relatively complex logic.

Packages like `os` and `fmt` help with standard general operations, while packages like `flag`, `encoding/csv` and `strings` help with specific operations, like parsing the CSV file and adding flags to the executable.

## Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [deployment](#deployment) for notes on how to deploy the project on a live system.

### Prerequisites

Make sure you have **GO** installed locally. Click [here](https://go.dev/doc/install) to get started with Go.

### Installing

A step by step series of examples that tell you how to get a development env running.

Clone the repository & change current the directory

```
git clone

cd
```

Make the executable

```
go build .
```

For help

```
quiz-app --help
```

## Usage <a name = "usage"></a>

The executable, by default looks for the `problems.csv` file. If you have a different file

```
quiz-app -csv="<FILE_NAME.csv>"
```

To start the quiz

```
quiz-app
```

After answering a question, press `enter` to continue to the next question. Once you have completed all the questions, you will be provided your result

```
You scored 12 out of 13.
```

## TODO <a name = "todo"></a>

- Implement time function via time package
- Add Go routines
