<p align="center"><h1 align="center">CAR-ZONE</h1></p>
<p align="center">
	<img src="https://img.shields.io/github/last-commit/diegobbrito/car-zone?style=default&logo=git&logoColor=white&color=0080ff" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/diegobbrito/car-zone?style=default&color=0080ff" alt="repo-top-language">
	<img src="https://img.shields.io/github/languages/count/diegobbrito/car-zone?style=default&color=0080ff" alt="repo-language-count">
</p>
<br>

##  Table of Contents

- [ Overview](#-overview)
- [ Project Structure](#-project-structure)
  - [ Project Index](#-project-index)
- [ Getting Started](#-getting-started)
  - [ Prerequisites](#-prerequisites)
  - [ Installation](#-installation)
  - [ Usage](#-usage)
  - [ Testing](#-testing)

---

##  Overview

<code>❯ Car Management System</code>

---


##  Project Structure

```sh
└── car-zone/
    ├── Dockerfile
    ├── README.md
    ├── db
    │   └── Dockerfile
    ├── docker-compose.yaml
    ├── driver
    │   └── postgress.go
    ├── go.mod
    ├── go.sum
    ├── go1.24.5.linux-amd64.tar.gz
    ├── handler
    │   ├── car
    │   ├── engine
    │   └── login
    ├── main.go
    ├── middleware
    │   ├── auth_middleware.go
    │   └── metrices_middleware.go
    ├── models
    │   ├── car.go
    │   ├── engine.go
    │   └── login.go
    ├── prometheus.yml
    ├── service
    │   ├── car
    │   ├── engine
    │   └── interface.go
    └── store
        ├── car
        ├── engine
        ├── interface.go
        └── schema.sql
```


###  Project Index
<details open>
	<summary><b><code>CAR-ZONE/</code></b></summary>
	<details> <!-- __root__ Submodule -->
		<summary><b>__root__</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/car-zone/blob/master/prometheus.yml'>prometheus.yml</a></b></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/car-zone/blob/master/docker-compose.yaml'>docker-compose.yaml</a></b></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/car-zone/blob/master/main.go'>main.go</a></b></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/car-zone/blob/master/go.mod'>go.mod</a></b></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/car-zone/blob/master/go.sum'>go.sum</a></b></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/car-zone/blob/master/Dockerfile'>Dockerfile</a></b></td>
			</tr>
			</table>
		</blockquote>
	</details>
	<details> <!-- store Submodule -->
		<summary><b>store</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/car-zone/blob/master/store/interface.go'>interface.go</a></b></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/car-zone/blob/master/store/schema.sql'>schema.sql</a></b></td>
			</tr>
			</table>
			<details>
				<summary><b>engine</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/diegobbrito/car-zone/blob/master/store/engine/engine.go'>engine.go</a></b></td>
					</tr>
					</table>
				</blockquote>
			</details>
			<details>
				<summary><b>car</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/diegobbrito/car-zone/blob/master/store/car/car.go'>car.go</a></b></td>
					</tr>
					</table>
				</blockquote>
			</details>
		</blockquote>
	</details>
	<details> <!-- models Submodule -->
		<summary><b>models</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/car-zone/blob/master/models/login.go'>login.go</a></b></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/car-zone/blob/master/models/engine.go'>engine.go</a></b></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/car-zone/blob/master/models/car.go'>car.go</a></b></td>
			</tr>
			</table>
		</blockquote>
	</details>
	<details> <!-- middleware Submodule -->
		<summary><b>middleware</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/car-zone/blob/master/middleware/auth_middleware.go'>auth_middleware.go</a></b></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/car-zone/blob/master/middleware/metrices_middleware.go'>metrices_middleware.go</a></b></td>
			</tr>
			</table>
		</blockquote>
	</details>
	<details> <!-- driver Submodule -->
		<summary><b>driver</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/car-zone/blob/master/driver/postgress.go'>postgress.go</a></b></td>
			</tr>
			</table>
		</blockquote>
	</details>
	<details> <!-- service Submodule -->
		<summary><b>service</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/car-zone/blob/master/service/interface.go'>interface.go</a></b></td>
			</tr>
			</table>
			<details>
				<summary><b>engine</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/diegobbrito/car-zone/blob/master/service/engine/engine.go'>engine.go</a></b></td>
					</tr>
					</table>
				</blockquote>
			</details>
			<details>
				<summary><b>car</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/diegobbrito/car-zone/blob/master/service/car/car.go'>car.go</a></b></td>
					</tr>
					</table>
				</blockquote>
			</details>
		</blockquote>
	</details>
	<details> <!-- handler Submodule -->
		<summary><b>handler</b></summary>
		<blockquote>
			<details>
				<summary><b>engine</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/diegobbrito/car-zone/blob/master/handler/engine/engine.go'>engine.go</a></b></td>
					</tr>
					</table>
				</blockquote>
			</details>
			<details>
				<summary><b>login</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/diegobbrito/car-zone/blob/master/handler/login/login.go'>login.go</a></b></td>
					</tr>
					</table>
				</blockquote>
			</details>
			<details>
				<summary><b>car</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/diegobbrito/car-zone/blob/master/handler/car/car.go'>car.go</a></b></td>
					</tr>
					</table>
				</blockquote>
			</details>
		</blockquote>
	</details>
	<details> <!-- db Submodule -->
		<summary><b>db</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/diegobbrito/car-zone/blob/master/db/Dockerfile'>Dockerfile</a></b></td>
			</tr>
			</table>
		</blockquote>
	</details>
</details>

---
##  Getting Started

###  Prerequisites

Before getting started with car-zone, ensure your runtime environment meets the following requirements:

- **Programming Language:** Go
- **Package Manager:** Go modules
- **Container Runtime:** Docker


###  Installation

Install car-zone using one of the following methods:

**Build from source:**

1. Clone the car-zone repository:
```sh
❯ git clone https://github.com/diegobbrito/car-zone
```

2. Navigate to the project directory:
```sh
❯ cd car-zone
```

3. Install the project dependencies:


**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go build
```


**Using `docker`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Docker-2CA5E0.svg?style={badge_style}&logo=docker&logoColor=white" />](https://www.docker.com/)

```sh
❯ docker build -t diegobbrito/car-zone .
```




###  Usage
Run car-zone using the following command:
**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go run {entrypoint}
```


**Using `docker`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Docker-2CA5E0.svg?style={badge_style}&logo=docker&logoColor=white" />](https://www.docker.com/)

```sh
❯ docker run -it {image_name}
```


###  Testing
Run the test suite using the following command:
**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go test ./...
```