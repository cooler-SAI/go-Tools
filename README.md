# go-Tools

This repository contains a collection of useful Go tools and utility packages designed to be easily integrated into your projects.

## How to Use

To use any tool from this repository, you'll need to import it into your Go project and then run `go mod tidy`.

### Current Tools:

---

### `stoper` - Graceful Shutdown Handler

The `stoper` package provides a simple way to handle graceful application shutdown upon receiving OS signals like `Ctrl+C` (SIGINT) or `SIGTERM`. It ensures your application can perform clean-up actions before exiting.

#### Installation

To include `stoper` in your project, simply import it:

```go
import (
    "fmt"
    // ... other imports
    "github.com/cooler-SAI/go-Tools/stoper" // Import the stoper package
)
```

---

### `zerolog` - Structured, Leveled Logging

The `zerolog` package provides a fast and flexible logger that produces structured (JSON by default) and leveled logs. It's designed for high performance and offers a simple API for rich logging capabilities.

#### Installation

To include `zerolog` in your project, you'll first need to get the package:
```go
import (
    "fmt"
    // ... other imports
    "github.com/cooler-SAI/go-Tools/zerolog" // Import the zerolog package
)
```

### `Dockerfile` - Docker Support

This repository now includes a `Dockerfile` located in the `docker/` directory. This Dockerfile provides a standardized way to build a minimal and secure Docker image for Go applications that might utilize the tools from this repository.

### Key Dockerfile Features:

*   **Multi-stage Builds:** Utilizes a builder stage for compilation and a separate, minimal runtime stage (`alpine` based) to keep the final image size small.
*   **Security Focused:** Creates a non-root user (`appuser`) to run the application, enhancing security.
*   **Optimized for Go:**
    *   Disables CGO by default (`CGO_ENABLED=0`) for static binaries.
    *   Targets Linux (`GOOS=linux`).
    *   Strips debugging symbols (`-ldflags="-s -w"`) to reduce binary size.
*   **Dependency Caching:** Leverages Docker's layer caching for Go module dependencies to speed up subsequent builds.
*   **Configurable:** Designed to build a Go project located in the parent directory relative to the `docker/` directory (i.e., your main Go project should be one level up from where the `Dockerfile` resides).

### Building an Image

To build a Docker image for your Go project using this Dockerfile:

1.  Ensure your Go project is structured such that the `docker/` directory (containing the `Dockerfile`) is a subdirectory of your main project.
2.  Navigate to the `docker/` directory in your terminal.
3.  Run the Docker build command or use config in your IDE
    