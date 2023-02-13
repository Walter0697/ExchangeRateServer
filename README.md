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

Before building the app, you should seed the database will the following command
```
go run main seed
```

You can also run the above command inside Docker container if you are running with deteched mode

Run the following command to build the image
```
docker built -t server/backend:<tag>
```

Then, use the following command to run this image, you will need to map the config file to docker in order to successfully run this application
```
docker run -v ${pwd}/config.toml:/app/config.toml --rm -p 80:80 server/backend:<tag>
```

## 🧳 API documentation <a name = "api"></a>

For this service, you will need an API key to fetch most prices exchange rate from our database, you can generate an API key using `/apikey` after logging in as a default user

### /auth/login [POST]

Post Body:
|Parameter|Optional|Description|Example|
|-|-|-|-|
|username|N|Username for login user|admin|
|password|N|Password for login user|chaos|

Default account can be set in `config.toml`

### /apikey [POST]
With jwt authentication: `Authentication: Bearer {token}` from `/auth/login`

Post Body:
|Parameter|Optional|Description|Example|
|-|-|-|-|
|identifier|N|An identifier for API key|testing_key|

For the below APIs, we are currently fetching BTC to USD only with our database, so only BTC to USD will actually give valid result.

### /price/last [GET]
With apikey authentication: `Authentication: {token}` from `/apikey`

URL Query:
|Parameter|Optional|Description|Default|
|-|-|-|-|
|crypto|Y|Crypto Currency that you want to exchange to |BTC|
|currency|Y|The currency that you want to exchange to|USD|

### /bytime/{time} [GET]
With apikey authentication: `Authentication: {token}` from `/apikey`

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
With apikey authentication: `Authentication: {token}` from `/apikey`

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