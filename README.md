## publicip

**publicip** is a simple command-line tool written in Go that displays your public IP address.

### Features

* Shows your public IP address.
* Outputs JSON format with `-json` flag.
* Supports Linux, macOS, and Windows.

### Installation

**Pre-built binaries:**

Download the pre-built binary for your operating system from [link to releases page](https://github.com/munnep/publicip/releases).

**From source:**

- Clone the repository:

```bash
git clone https://github.com/munnep/publicip.git
```
- Build the tool:
```Bash
cd publicip
go build -o publicip
```

### Usage

```bash
publicip -h

Flags:

  -json
        Specify if output should be JSON
  -version
        Show the version
```

Examples:

Show public IP:
```bash
publicip
```

Show public IP in JSON:
```Bash
publicip -json
```

Output:

Without the -json flag, publicip will display your public IP address as a simple string.
```bash
123.456.78.90
```

With the -json flag, the output will be formatted as follows:

JSON
```
{
  "ip": "123.456.78.90"
}
```

### License
This project is licensed under the MIT License. See the LICENSE file for details.
