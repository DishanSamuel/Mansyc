# sysStart-quit

**sysStart-quit** is a Bash script designed to efficiently manage system services on Linux.  
It allows users to **start essential services** listed in a configuration file and **terminate unnecessary services** by comparing running processes against a predefined list.  

In the future, **Golang** integration may be added to enhance performance and usability.

## Features

- Start all essential services listed in a file.
- Terminate unnecessary services that are not part of the essential list.
- Works on **Linux** systems.
- Lightweight and efficient.
- Future support for **Golang** integration.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/sysStart-quit.git
   cd sysStart-quit
   ```
2. Make the script executable:
   ```bash
   chmod +x sysStart-quit.sh
   ```

## Usage

1. Create a text file (e.g., `services.txt`) listing essential services:
   ```
   sshd
   cron
   apache2
   ```
2. Run the script to **start** essential services:
   ```bash
   ./sysStart-quit.sh --start services.txt
   ```
3. Run the script to **stop all unnecessary services**:
   ```bash
   ./sysStart-quit.sh --stop services.txt
   ```

## Dependencies

- Linux-based system
- `systemctl` (for managing systemd services)
- `grep`, `awk`, `ps` (for process management)

## Future Improvements

- Golang integration for better performance and cross-platform compatibility.
- GUI or CLI enhancements for easier usage.
- Logging functionality to track service management actions.

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## License

MIT License. See `LICENSE` for details.
