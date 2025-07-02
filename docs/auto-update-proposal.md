# Automatic Update System Proposal

## 1. Introduction

This document proposes an approach to implement an automatic update system for the `deepl-cli` executable, similar to what is used by tools like OhMyZsh. The goal is to allow the application to check for and install new available versions on startup, thereby improving user experience and maintenance.

## 2. Operating Principles

The automatic update process will follow these general steps:

1.  **Current Version Check:** On launch, the application will determine its own version (which will be embedded during compilation).
2.  **Remote Source Query:** The application will contact a remote source (e.g., [GitHub Releases](https://docs.github.com/en/repositories/releasing-projects-on-github/managing-releases-in-a-repository)) to check if a newer version is available.
3.  **New Version Download:** If a new version is detected, the executable corresponding to the user's OS and architecture will be downloaded to a temporary location.
4.  **Executable Replacement:** The application will attempt to replace itself with the newly downloaded executable. This is the most delicate step, as a running process cannot directly overwrite its own file. This typically involves:
    *   Renaming the old executable (e.g., `deepl.old`).
    *   Moving the new executable to the original path (`deepl`).
    *   Handling potential errors (permissions, corrupted download, etc.).
    *   Cleaning up the old executable on a subsequent successful startup.

## 3. Technical Implementation (Go)

To facilitate the implementation of this system in Go, it is highly recommended to use an existing library that handles the complexities of replacing a running executable and managing versions. The `github.com/rhysd/go-github-selfupdate` library is an excellent candidate for projects hosted on GitHub.

### Implementation Steps:

1.  **Library Integration:**
    *   Add `github.com/rhysd/go-github-selfupdate` to the project dependencies (`go get github.com/rhysd/go-github-selfupdate`).

2.  **Version Definition:**
    *   The application version must be defined in a way that can be read by the code (e.g., via a `const` variable or a variable injected at compile time via `ldflags`).
    *   Example: `go build -ldflags "-X main.version=1.0.0"`

3.  **Check and Update Logic:**
    *   At the beginning of the `main` function (or in a dedicated function called by `main`), implement the check logic.
    *   Use the `selfupdate.UpdateSelf` or `selfupdate.UpdateSelfFromSource` function from the library.
    *   This function will handle version comparison, downloading, and replacement.

    ```go
    package main

    import (
    	"fmt"
    	"os"
    	"time"

    	"github.com/rhysd/go-github-selfupdate/selfupdate"
    )

    var version = "0.0.1" // This value will be replaced by ldflags during compilation

    func doSelfUpdate() {
    	latest, err := selfupdate.UpdateSelf(version, "Jsamson33/deepl-cli")
    	if err != nil {
    		fmt.Println("Self-update failed:", err)
    		return
    	}
    	if latest.Version.String() != version {
    		fmt.Println("Successfully updated to version", latest.Version)
    		fmt.Println("Release Note:\n", latest.ReleaseNotes)
    	} else {
    		fmt.Println("Current version is the latest")
    	}
    }

    func main() {
    	// Optional: Check for updates only after a certain delay or with a specific flag
    	// if time.Now().Sub(lastUpdateTime) > 24 * time.Hour {
    	//     doSelfUpdate()
    	// }

    	doSelfUpdate()

    	// Rest of the application logic
    	// ...
    }
    ```

4.  **Error Handling and User Experience:**
    *   Provide clear messages to the user regarding the update status (new version available, successful update, failure).
    *   Handle cases where the user does not have the necessary permissions for the update.
    *   Consider adding a CLI flag (`--no-self-update`) to disable automatic checking.

## 4. Important Considerations

*   **Permissions:** The update system requires write permissions to the directory where the executable is installed. This can be an issue on some systems (e.g., `/usr/bin`). One solution is to ask the user to install the executable in a directory where they have write access (e.g., `/usr/local/bin` or a personal directory).
*   **Binary Signatures:** For increased security, it is highly recommended to sign release binaries and verify these signatures during the update. The `go-github-selfupdate` library supports GPG signature verification.
*   **Check Frequency:** Avoid checking for updates on every launch to prevent overloading the GitHub API or slowing down application startup. A daily or weekly check is usually sufficient.
*   **Rollback:** In case of update failure, the old executable is usually retained (renamed). This allows for manual rollback if necessary.
*   **Compatibility:** Ensure that new versions are compatible with older configurations and user data.

## 5. Next Steps

1.  Discuss the feasibility of implementing this feature.
2.  If approved, integrate the `go-github-selfupdate` library.
3.  Set up the release process on GitHub Actions to generate binaries with versions and potentially signatures.
4.  Implement the update logic in the code.
5.  Rigorously test the update system on different platforms.
