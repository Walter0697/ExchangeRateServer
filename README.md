<h3 align="center">Exchange Rate Server</h3>

<div align="center">
</div>

---

<p align="center"> A simple server to retrieve Exchange rate between BTC and USD
    <br> 
</p>

## 📝 Table of Contents

- [Getting Started](#getting_started)
- [Running Test Case](#tests)
- [API Documentation](#api)
- [Author](#authors)

## 🏁 Getting Started <a name = "getting_started"></a>

Here is a simple instruction for running this server

### Prerequisites

- Docker
- Postgres Database

First, you will need to copy `config.example.toml` into a preferrable name, in the following example, we will use `config.toml`

Change the configuration to match your preferred database inside `config.toml`

Please be aware that the serving port of this application is `80`, if you want to change it, you should do it in both `config.toml` and `Dockerfile`

Run the following command to build the image
```
docker build -t server/backend:<tag> .
```

You will need to map the config file to docker in order to successfully run this application
Copy `config.toml` to `${directory}`, and then run the following command
```
docker run -v ${directory}/config.toml:/app/config.toml --rm -p 80:80 server/backend:<tag>
```

## 🧳 API documentation <a name = "api"></a>

For the below APIs, we are currently fetching BTC to USD only with our database, so only BTC to USD will actually give valid result.

### /price/last [GET]

URL Query:
|Parameter|Optional|Description|Default|
|-|-|-|-|
|crypto|Y|Crypto Currency that you want to exchange to |BTC|
|currency|Y|The currency that you want to exchange to|USD|

### /bytime/{time} [GET]

URL Query:
|Parameter|Optional|Description|Default|
|-|-|-|-|
|crypto|Y|Crypto Currency that you want to exchange to |BTC|
|currency|Y|The currency that you want to exchange to|USD|

URL Parameter:
|Parameter|Optional|Description|Default|
|-|-|-|-|
|time|N|Time Format RFC3339|2023-02-12T14:25:26+08:00|

### /range/{start}/{end} [GET]

URL Query:
|Parameter|Optional|Description|Default|
|-|-|-|-|
|crypto|Y|Crypto Currency that you want to exchange to |BTC|
|currency|Y|The currency that you want to exchange to|USD|

URL Parameter:
|Parameter|Optional|Description|Default|
|-|-|-|-|
|start|N|Starting time for fetching the range, Time Format RFC3339|2023-02-12T14:25:26+08:00|
|end|N|Ending time for fetching the range, Time Format RFC3339|2023-02-12T14:25:26+08:00|

## 🔧 Running the tests <a name = "tests"></a>

Here is a basic instruction for running test case

Run the following command
```
go test
```
if you want to change anything about the test case, we can change the files that ends with `_test.go`

## ✍️ Author <a name = "authors"></a>

- [@Walter0697](https://github.com/Walter0697)