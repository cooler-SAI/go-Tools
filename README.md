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

#### Use
Just use this code in your project:

```go
// Using the logger with default settings (text output)
zerolog.Log.Info().Msg("This is a message with default settings")

// Initialize with colored console output
zerolog.Init()

// Now logs will be in colored format
zerolog.Log.Info().Msg("This is an informational message")
zerolog.Log.Warn().Msg("This is a warning")
zerolog.Log.Error().Msg("This is an error")
```

---

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

---

### `random` - Random Value Generation

Utility functions for generating random values using math/rand/v2.

#### Installation

To include `random` in your project, you'll first need to get the package:
```go
import (
    "fmt"
    // ... other imports
    "github.com/cooler-SAI/go-Tools/random" // Import the zerolog package
)
```
### How to Use

To use any tool from this repository, you'll need to import it into your Go project and then run `go mod tidy`.

```go
// Returns random int in [min, max] range
num := random.RandRange(1, 100)
```

---

### `Cleaner` - Run and Del build

This section describes how to configure GoLand IDE to automatically compile, run, and then delete your Go project's executable file after it finishes execution. This is achieved by using a custom BAT file that manages the entire build and cleanup process.

### Key Cleaner Features:

*   **Automated Build & Run:** Automatically compiles your Go project and launches the executable with a single action (Shift + F10).
*   **Post-Execution Cleanup:** Ensures the compiled executable (.exe on Windows) is deleted immediately after the program finishes, keeping your project directory clean.
*   **Simplified Workflow:** Eliminates the need for manual cleanup commands after each test or run.
*   **IDE Integrated:** Fully configured within GoLand's Run/Debug configurations for seamless developer experience.
*   **Customizable:** The underlying BAT script can be easily modified to include additional pre-run or post-run commands if needed.

### Installation (GoLand Configuration)

Important Notes for the BAT file:

1.  PROJECT_DIR: Set this to the absolute path of the directory containing your main Go file (e.g., transactions.go).
2.  BUILD_NAME: Specify your desired executable file name (e.g., transactions.exe).
3.  go build -o "%BUILD_NAME%" transactions.go: This command compiles your transactions.go and saves the result as transactions.exe within the PROJECT_DIR.

### Configure GoLand Run/Debug Configuration:

Now, you need to configure GoLand to use this BAT file to control the build and run process.

1. Open Run/Debug Configurations:
   Go to Run | Edit Configurations... in the top menu of GoLand.

2. Select or Create a Configuration:
   Choose your existing "Go Build" or "Go Application" configuration for your project (e.g., postgres or transactions). If you don't have one, create a new one.

3. Configure the "Before launch" section:

   REMOVE the default "Build" (Go Build) step:
   In the "Before launch" list, find the step of type "Build" (Go Build). Select it and click the minus (-) button to remove it. This is crucial, as the build will now be handled by your BAT file.

4. Add your "External Tool":
   Click the plus (+) button and select Run External Tool. If you haven't configured this tool before, click + again to create a new one. Fill in the fields for the new external tool as follows:
      * Name: Run and Clean Build (or any descriptive name)
      * Program: Provide the full, absolute path to your build_and_run.bat file.
      * Example: D:\Enterprise Development\Go-projects\go-projects\base\postgres\build_and_run.bat
      * Arguments: Leave this field BLANK. The BAT file handles the paths internally and does not require arguments from GoLand.
      * Working directory: $ProjectFileDir$ (This macro points to the root directory of your Go project, e.g., D:\Enterprise Development\Go-projects\go-projects\base).

5. Apply Changes:
   Click OK to save the new external tool. Ensure that Run and Clean Build is now the only step listed in the "Before launch" section (apart from potentially "Build dependencies" or similar steps GoLand adds by default). Click Apply, then OK in the "Run/Debug Configurations" window to save all changes.

### How to Use

1. Once the GoLand configuration is set up:

2. Simply run your project using Shift + F10 (or by clicking the "Run" button in GoLand).

3. The build_and_run.bat script will automatically compile your Go application, execute it, and then delete the generated executable file upon completion.

#### Example console log:
```log
"D:\Enterprise Development\Go-projects\go-projects\base\postgres\build_and_run.bat"

--- Build Process ---
Current working directory for build: D:\Enterprise Development\Go-projects\go-projects\base\postgres
Target build path: D:\Enterprise Development\Go-projects\go-projects\base\postgres\transactions.exe
Go build successful.

--- Launching Program ---
Launching: "D:\Enterprise Development\Go-projects\go-projects\base\postgres\transactions.exe"
Program finished with exit code: 0

--- Cleanup Process ---
Deleting build: "D:\Enterprise Development\Go-projects\go-projects\base\postgres\transactions.exe"
Build successfully deleted.

Process finished with exit code 0
```
---













