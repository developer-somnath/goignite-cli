# GoIgnite CLI

GoIgnite CLI is a command-line tool designed to help you quickly scaffold and manage GoIgnite projects. It provides commands for generating new projects, managing project components, and more.

## Features

- Quickly generate a new GoIgnite project structure
- Version tracking of the CLI
- Scaffold new controllers, models, and other components (coming soon)

## Installation

You can install the GoIgnite CLI globally on your system using the `go install` command. Ensure you have [Go](https://golang.org/doc/install) installed and set up on your system.

### Install the latest version:

To install the latest version of GoIgnite CLI, run:

```bash
go install github.com/your-username/goignite-cli@latest

```
### Install the specific version:
```bash
go install github.com/your-username/goignite-cli@v1.0.0

```

## Download Pre-built Executable (Windows)

If you're using Windows and prefer to download the executable directly rather than installing via Go, you can download the latest `goignite-cli.exe` from the releases page.

### Steps to Download and Use:

1. Go to the Releases page of this repository.
2. Download the latest version of `goignite-cli.exe` under the **Assets** section.
3. Place the downloaded `goignite-cli.exe` in a directory of your choice.
4. Add the directory where you placed the `goignite-cli.exe` to your system's PATH environment variable for easier usage. Here's how to do that:
   * Right-click on **This PC** or **My Computer**, and choose **Properties**.
   * Click on **Advanced system settings**.
   * In the System Properties window, click on the **Environment Variables** button.
   * Under **System variables**, select the **Path** variable and click **Edit**.
   * Click **New** and add the path where you placed `goignite-cli.exe`.
   * Click **OK** to close all windows.
5. Now, you can use `goignite` from any command prompt window:

   ```
   goignite new my_project
   ```

### Running Without Adding to PATH

Alternatively, you can run the CLI without adding it to the PATH by navigating to the folder where you downloaded `goignite-cli.exe` and running:

```
./goignite-cli.exe new my_project
```

## Usage

Once the CLI is installed, you can use it to generate projects and more. Below are the basic commands:

### Create a New Project

To create a new GoIgnite project, run:

```
goignite new my_project
```

This command will create a new folder `my_project` with the following structure:

```
my_project/
│
├── app/
│   ├── controllers/
│   └── models/
│
├── config/
├── routes/
├── public/
└── resources/
    └── views/
```

### Check CLI Version

To display the version of the GoIgnite CLI you have installed, run:

```
goignite version
```

This will print the current version of the CLI.

### Updating GoIgnite CLI

To update GoIgnite CLI to the latest version, simply run:

```
go install github.com/your-username/goignite-cli@latest
```

This will pull the latest version from the repository and install it on your system.

### Uninstall GoIgnite CLI

To uninstall the GoIgnite CLI, simply delete the binary file from your Go binaries path:

```
rm $(go env GOPATH)/bin/goignite
```

---

### Key Points:
- **Go Install**: Instructions on how to install the CLI using `go install` with support for the latest and specific versions.
- **Windows Executable**: Clear steps for downloading and setting up the `goignite-cli.exe` file on Windows.
- **Usage**: Basic commands like creating a new project and checking the CLI version.
- **Updating & Uninstalling**: How to update and remove the CLI.
- **Contributing & License**: Encouraging collaboration and showing the license information.

This README gives users multiple installation options, detailed usage, and all the essential information for contributing or updating the CLI. You can now push this to your GitHub repository!


