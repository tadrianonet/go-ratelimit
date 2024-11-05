# Rate limit

This code provides a Golang project implementing rate limit.

## Install dependencies:

```bash
go mod tidy
```

## Configuration

You can adjust the rate limit settings in the code as follows:

Rate: Set the number of requests per second.
Burst: Set the maximum number of requests that can be processed at once during short bursts.

For example, to set a limit of 10 requests per second with a burst of 2:

```bash
limiter := rate.NewLimiter(10, 2)
```

## License

This project is open-source and available under the MIT License.

