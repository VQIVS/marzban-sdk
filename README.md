# Marzban SDK

A well-structured HTTP client for the Marzban API, following Go best practices and clean architecture principles.

## Features

- ✅ **Clean Architecture**: Separation of concerns with internal packages and service handlers
- ✅ **Multiple Authentication Methods**: Support for username/password and client credentials
- ✅ **Error Handling**: Comprehensive error handling with custom API error types
- ✅ **Type Safety**: Strongly typed request/response models
- ✅ **HTTP Best Practices**: Proper use of HTTP methods, status codes, and headers
- ✅ **Configurable**: Flexible client configuration with functional options
- ✅ **Complete API Coverage**: All Marzban API endpoints available as constants

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    "github.com/VQIVS/marzban-sdk/internal/models"
    "github.com/VQIVS/marzban-sdk/service/handlers"
)

func main() {
    // Create client
    client := handlers.NewMarzbanClient("https://api.marzban.example.com")

    // Method 1: Login with username/password
    loginReq := models.UserLoginReq{
        Username: "admin",
        Password: "password123",
    }

    loginResp, err := client.LoginWithUsername(loginReq)
    if err != nil {
        log.Fatalf("Login failed: %v", err)
    }

    fmt.Printf("Logged in successfully! Token: %s\n", loginResp.Token)

    // Method 2: Login with client credentials
    loginResp2, err := client.LoginWithClientID("client_id", "client_secret")
    if err != nil {
        log.Fatalf("Client login failed: %v", err)
    }

    fmt.Printf("Client login successful! Token: %s\n", loginResp2.Token)
}
```

## Architecture Overview

The SDK follows a clean architecture pattern with clear separation of concerns:

```
├── internal/                # Internal packages (not exposed to consumers)
│   ├── client/             # Core HTTP client implementation
│   │   ├── clinet.go      # Client struct and configuration
│   │   └── endpoints.go   # API endpoint constants
│   └── models/            # Request/Response models
│       ├── req.go        # Request structures
│       └── res.go        # Response structures
├── service/               # Service layer (public API)
│   ├── handlers/         # API handlers and business logic
│   │   ├── admin.go     # Admin-related operations
│   │   └── client.go    # Client wrapper
│   └── utils/           # Utility functions
├── README.md
└── go.mod
```

## API Coverage

### Authentication
- `EndpointAdminToken` - `/api/admin/token`

### Admin Management
- `EndpointAdmin` - `/api/admin`
- `EndpointAdminByUsername` - `/api/admin/{username}`
- `EndpointAdmins` - `/api/admins`
- `EndpointAdminUsersDisable` - `/api/admin/{username}/users/disable`
- `EndpointAdminUsersActivate` - `/api/admin/{username}/users/activate`
- `EndpointAdminUsageReset` - `/api/admin/usage/reset/{username}`
- `EndpointAdminUsage` - `/api/admin/usage/{username}`

### Core Management
- `EndpointCore` - `/api/core`
- `EndpointCoreRestart` - `/api/core/restart`
- `EndpointCoreConfig` - `/api/core/config`

### Node Management
- `EndpointNodeSettings` - `/api/node/settings`
- `EndpointNode` - `/api/node`
- `EndpointNodeByID` - `/api/node/{node_id}`
- `EndpointNodes` - `/api/nodes`
- `EndpointNodeReconnect` - `/api/node/{node_id}/reconnect`
- `EndpointNodesUsage` - `/api/nodes/usage`

### User Management
- `EndpointUser` - `/api/user`
- `EndpointUserByUsername` - `/api/user/{username}`
- `EndpointUsers` - `/api/users`
- `EndpointUserReset` - `/api/user/{username}/reset`
- `EndpointUserUsage` - `/api/user/{username}/usage`
- `EndpointUsersExpired` - `/api/users/expired`

### Subscription
- `EndpointSubscription` - `/sub/{token}/`
- `EndpointSubscriptionInfo` - `/sub/{token}/info`
- `EndpointSubscriptionUsage` - `/sub/{token}/usage`
- `EndpointSubscriptionClientType` - `/sub/{token}/{client_type}`

### System
- `EndpointSystem` - `/api/system`
- `EndpointInbounds` - `/api/inbounds`
- `EndpointHosts` - `/api/hosts`

## Client Configuration

The client supports various configuration options:

```go
import (
    "time"
    "github.com/VQIVS/marzban-sdk/internal/client"
    "github.com/VQIVS/marzban-sdk/service/handlers"
)

// Basic client
client := handlers.NewMarzbanClient("https://api.marzban.example.com")

// Client with custom timeout
client := handlers.NewMarzbanClient(
    "https://api.marzban.example.com",
    client.WithTimeout(30*time.Second),
)

// Client with custom HTTP client
customHTTPClient := &http.Client{
    Timeout: 60 * time.Second,
}
client := handlers.NewMarzbanClient(
    "https://api.marzban.example.com",
    client.WithHTTPClient(customHTTPClient),
)
```

## Error Handling

The SDK provides comprehensive error handling:

```go
loginResp, err := client.LoginWithUsername(loginReq)
if err != nil {
    if apiErr, ok := err.(*models.ErrorResponse); ok {
        fmt.Printf("API Error: %s\n", apiErr.Error())
        fmt.Printf("Details: %s\n", apiErr.Detail)
    } else {
        fmt.Printf("Network/Other Error: %v\n", err)
    }
    return
}
```

## Authentication Methods

### 1. Username/Password Authentication
```go
loginReq := models.UserLoginReq{
    Username: "admin",
    Password: "password123",
}
loginResp, err := client.LoginWithUsername(loginReq)
```

### 2. Client Credentials Authentication
```go
loginResp, err := client.LoginWithClientID("client_id", "client_secret")
```

Both methods automatically set the authentication token for subsequent requests.

## Working with Endpoints

You can use the predefined endpoint constants with helper functions:

```go
import "github.com/VQIVS/marzban-sdk/internal/client"

// Get endpoint for specific user
userEndpoint := client.GetUserByUsernameEndpoint("john_doe")
// Result: /api/user/john_doe

// Get endpoint for specific node
nodeEndpoint := client.GetNodeByIDEndpoint(123)
// Result: /api/node/123

// Get subscription endpoint
subEndpoint := client.GetSubscriptionEndpoint("abc123token")
// Result: /sub/abc123token/
```

## Design Principles

1. **Clean Architecture**: Internal packages hide implementation details
2. **Type Safety**: All requests/responses are strongly typed
3. **Error Transparency**: Custom error types preserve HTTP details
4. **Resource Management**: Proper cleanup of HTTP resources
5. **Extensibility**: Easy to add new endpoints and handlers

## Contributing

When adding new functionality:

1. Define request/response models in `internal/models/`
2. Add endpoint constants to `internal/client/endpoints.go`
3. Implement handlers in `service/handlers/`
4. Follow the established error handling patterns
5. Add comprehensive examples in documentation

## License

[Add your license information here]

## License

[Add your license information here]
