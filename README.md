# sysStart-Quit

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

**1. Clone the repository:**
   ```bash
   git clone https://github.com/DishanSamuel/sysStart-Quit.git
   cd sysStart-Quit
   ```
**2. Make the script executable & install requirements:**
   ```bash
   chmod +x ./config.sh
   ./config.sh
   ```

## Usage

**1. setting up the text file listing startup & essential services required:**
<br>

   THERE ARE TWO .txt FILES <br>
   => normalser.txt (used in quiting services) <br>
   => startup.txt   (used in starting services) <br>

   (i)   **normalser.txt** --> contains essinatial services for the linux system to operate <br>
   (ii)  **startup.txt**   --> contains the services to be started during running the seript <br>
   
   These two files can be costimized according to user needs

   !! DISCLAIMER !! ( normalser.txt may wary from system to system, please review the file before running it)

   

   TO KNOW THE DEFAULT SERVICE IN THE SYSTEM YOU CAN RUN THE FOLLOWING COMMAND
   ```
   systemctl list-units --type=service --state=running
   ```


   DEFAULTS ARE 
   ```
   accounts-daemon.service
   colord.service
   cron.service
   dbus.service
   fwupd.service
   gdm.service
   haveged.service
   ModemManager.service
   NetworkManager.service
   polkit.service
   power-profiles-daemon.service
   rtkit-daemon.service
   systemd-journald.service
   systemd-logind.service
   systemd-timesyncd.service
   systemd-udevd.service
   udisks2.service
   upower.service
   user@1000.service
   wpa_supplicant.service
   bolt.service
   bluetooth.service
   ```


   
**2. Run the script to start & end essential services:**

   ```bash
   ./run_main.sh
   ```
   ![image](https://github.com/user-attachments/assets/ef247bc0-999f-4f23-b8d8-68ff067316dc)


   PRESS "1" --> to start services <br>
   PRESS "2" --> end services


**3.Working:**

   ![image](https://github.com/user-attachments/assets/ffe57c39-f520-4d3c-8e23-245895a57bec)
   

   ![image](https://github.com/user-attachments/assets/fe32faca-55d6-4ce7-8fd9-92a267561419)




## Dependencies

- Linux-based system
- `systemctl` (for managing systemd services)
- `grep`, `awk`, `sed`(for process management)
  

## Future Improvements

- Golang integration for better performance and cross-platform compatibility.
- GUI or CLI enhancements for easier usage.
- Logging functionality to track service management actions.


## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## License

MIT License. See `LICENSE` for details.
