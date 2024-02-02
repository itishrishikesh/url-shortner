![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)![Redis](https://img.shields.io/badge/redis-%23DD0031.svg?style=for-the-badge&logo=redis&logoColor=white)![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
# URL-Shortner

URL-Shortner is a simple and fast service that allows you to shorten long URLs and access them with a custom alias. It is written in Go and uses Redis as the database. It also provides an API endpoint for creating and retrieving shortened URLs.

![Untitled Diagram drawio (3)](https://github.com/itishrishikesh/url-shortner/assets/24785679/2fe36927-b9c1-4a61-be08-595c6a133919)

## Features

- Shorten any valid URL with a custom alias or a randomly generated one
- Set an expiry time for the shortened URL in minutes
- Get the original URL from the shortened one
- View the rate limit and remaining requests for the API
- Configure the service using environment variables
- Run the service using Docker Compose

## Usage

### API Endpoint

The API endpoint of the service at `http://localhost:3000/shorten`. It accepts a JSON body with the following fields:

- `url`: The URL you want to shorten. It must be a valid URL starting with `http://` or `https://`.
- `short`: The custom alias you want to use for the shortened URL. It must be alphanumeric and not longer than 10 characters. If you do not provide this field, the service will generate one for you using the UUID library and taking the first 6 characters.
- `expiry`: The expiry time for the shortened URL in minutes. It must be a positive integer. If you do not provide this field, the service will use 0, which means the URL will never expire.

The API will respond with a JSON body with the following fields:

- `url`: The original URL that was shortened.
- `short`: The shortened URL that was created. It will have the format `http://localhost:3000/<alias>`, where `<alias>` is either the custom alias you provided or the randomly generated one.
- `expiry`: The expiry time for the shortened URL in minutes. It will be the same as the one you provided or 0 if you did not provide one.
- `x_rate_remaining`: The number of remaining requests you can make to the API in the current minute.
- `x_rate_limit_rest`: The number of seconds until the rate limit is reset.

For example, if you send the following request:

```
curl -X POST -H "Content-Type: application/json" -d '{"url":"https://google.com", "short":"gg", "expiry":20}' http://localhost:3000/shorten
```

You will get the following response:

```
{
  "url": "https://google.com",
  "short": "http://localhost:3000/search",
  "expiry": 20,
  "x_rate_remaining": 9,
  "x_rate_limit_rest": 57
}
```

You can then use the shortened URL `http://localhost:3000/search` to access the original URL `https://google.com` for the next 20 minutes. After that, the URL will expire and you will see a message saying that the URL is not found.

If you send an invalid request, such as an empty body, an invalid URL, an invalid alias, or an invalid expiry time, you will get an error response with the status code 400 and a message explaining the error.

For example, if you send the following request:

```
curl -X POST -H "Content-Type: application/json" -d '{"url":"google.com", "short":"gg", "expiry":20}' http://localhost:3000/shorten
```

You will get the following response:

```
{
  "error": "Invalid URL: google.com"
}
```

The API has a rate limit of 10 requests per minute per IP address. If you exceed the limit, you will get an error response with the status code 429 and a message saying that you have reached the limit. You will also see the `x_rate_remaining` and `x_rate_limit_rest` fields in the response header.

For example, if you send the 11th request in the same minute, you will get the following response:

```
{
  "error": "You have reached the limit of 10 requests per minute"
}
```

## Configuration

You can configure the service using the following environment variables:

- `DB_ADDR`: The address of the Redis database. It defaults to `db:6379`.
- `DB_PASS`: The password of the Redis database. It defaults to an empty string.
- `APP_PORT`: The port on which the service will run. It defaults to `:3000`.
- `DOMAIN`: The domain name of the service. It defaults to `localhost:3000`. This is the prefix with which shortened url will be stored.
- `API_QUOTA`: The rate limit for the API in requests per minute. It defaults to 10.

You can set these variables in the `.env` file in the root directory of the project. The file has the following format:

```
DB_ADDR="db:6379"  
DB_PASS=""  
APP_PORT=":3000"  
DOMAIN="localhost:3000"  
API_QUOTA=10
```

## Installation

You can run the service using Docker Compose, which will create a container for the service and a container for the Redis database. To do so, you need to have Docker and Docker Compose installed on your machine.

First, clone the repository and change the current directory to the project root:

```
git clone https://github.com/<your-username>/url-shortner.git
cd url-shortner
```

Then, build and run the containers using the following command:

```
docker-compose up -d
```

This will create two containers: `url-shortner_app_1` and `url-shortner_db_1`. You can check the status of the containers using the following command:

```
docker-compose ps
```

You should see something like this:

```
       Name                     Command               State           Ports         
-----------------------------------------------------------------------------------
url-shortner_app_1   /bin/sh -c go run main.go        Up      0.0.0.0:3000->3000/tcp
url-shortner_db_1    docker-entrypoint.sh redis ...   Up      6379/tcp             
```

You can also view the logs of the containers using the following command:

```
docker-compose logs
```

You can stop the containers using the following command:

```
docker-compose down
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
