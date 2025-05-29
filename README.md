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