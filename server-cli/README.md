# Server CLI

## Overview
The Server CLI is a command-line interface tool designed to manage and monitor server environments efficiently. It provides commands to check server statuses and monitor system performance in real-time.

## Commands

### check-servers
The `check-servers` command checks the status of servers in the project. It utilizes the Docker Compose file to verify if all servers are running smoothly.

**Usage:**
```
server-cli check-servers
```

### monitor
The `monitor` command provides real-time insights into system performance and service health. It includes functionalities for:

- System resource monitoring
- Docker container statistics
- Service health checks
- Uptime and load average
- Log analysis for errors

**Usage:**
```
server-cli monitor
```

## Installation
To install the Server CLI, clone the repository and build the project:

```bash
git clone <repository-url>
cd server-cli
go build
```

## Usage
After building the project, you can run the CLI commands directly from the terminal:

```bash
./server-cli check-servers
./server-cli monitor
```

## Contributing
Contributions are welcome! Please submit a pull request or open an issue for any enhancements or bug fixes.

## License
This project is licensed under the MIT License. See the LICENSE file for details.