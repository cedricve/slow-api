# slow-api

A configurable API service that introduces artificial delays to responses, useful for testing timeout handling, loading states, and performance optimization in client applications.

## Concept

`slow-api` is a simple HTTP server built with Go and Gin that intentionally slows down API responses by introducing configurable delays. This is particularly useful for:

- **Testing timeout behavior**: Verify that your client applications properly handle slow or unresponsive APIs
- **Simulating network latency**: Mimic real-world network conditions in development environments
- **Testing loading states**: Ensure loading indicators and user feedback mechanisms work correctly
- **Performance testing**: Validate retry logic, circuit breakers, and other resilience patterns
- **Development and debugging**: Slow down responses to better observe and debug asynchronous behavior

The server provides a `/health` endpoint that:
1. Accepts incoming requests
2. Waits for a random delay between configurable minimum and maximum timeouts
3. Returns a JSON response with the health status, timestamp, and actual delay applied

## Configuration

The delay behavior is controlled via environment variables:

- `MIN_TIMEOUT`: Minimum delay in milliseconds (default: 4000ms)
- `MAX_TIMEOUT`: Maximum delay in milliseconds (default: 4000ms)
- `PORT`: Server port (default: 8081)

Configuration can be set in a `.env.local` file or via standard environment variables.

## Usage

### Starting the Server

```bash
go run main.go
```

### Example Response

```json
{
  "status": "ok",
  "timestamp": "2025-12-03T10:30:45Z",
  "delay_ms": 4000
}
```

The `delay_ms` field indicates the actual delay that was applied to that specific request.
