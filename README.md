# Question Site

A web app made with Gin and Vue that allows users to ask and answer questions.

## Installation

#### Requirements

- [Go](https://go.dev/doc/install)
- [Node.js](https://nodejs.org/en/download/)

#### Download and Setup

```bash
git clone https://github.com/esgameco/question-site.git

# Install node dependencies
cd questionclient
npm i

# Set secret key
export SECRET="insert encryption key" # Linux
$Env:SECRET = "insert encryption key" # Windows
```

## Run

#### Development

```bash
Linux: ./devrun.sh
Windows: .\devrun.bat
```

## TODO

1. Prototype
    - [x] Question and Answer Model
    - [x] Database Setup (sqlite for development)
    - [x] CRUD
    - [x] User Authentication
    - [x] Add API to Vue
    - [ ] Site-wide styling
    - [ ] Style components