# sysStart-Quit 🚀

**sysStart-Quit** is a **Bash script** designed to efficiently manage system services on Linux.  
It helps users **start essential services** 🟢 listed in a configuration file and **terminate unnecessary services** 🔴 by comparing running processes against a predefined list.

**Future Updates:**  
Integration with **Golang** 🖥️ is planned to enhance performance and usability.

---

## Features 🌟

- **Start essential services** 🟢 listed in a configuration file.
- **Terminate unnecessary services** 🔴 that are not part of the essential list.
- Works seamlessly on **Linux** systems 🐧.
- **Lightweight** and **efficient** for system resource management ⚡.
- **Future support** for **Golang** integration 🚀.

---

## Installation 🔧

### 1. Clone the repository:

    git clone https://github.com/DishanSamuel/Mansyc.git<br>
    cd sysStart-Quit

### 2. Install Go (if you haven't already) 🖥️:

Follow the instructions here: [Install Go](https://go.dev/doc/install)

### 2. Make the script executable & install requirements:

    chmod +x config.sh
    ./config.sh

## Usage 🛠️
Setting up the text files for startup and essential services:

There are two configuration .txt files:

### normalser.txt 📝

   -Contains essential services for the Linux system to operate.<br>
   -This file may vary from system to system, so please review it before running ⚠️.<br>

### startup.txt 📝

   -Contains the services that should be started during script execution.

### You can customize these files according to your needs.
   
   To know the default services on your system, run the following command:
      
    systemctl list-units --type=service --state=running

### Default services include:

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

## Run the script to start or end essential services:

## Golang Integration 🖥️

As of the latest update, Golang has been integrated into the project to improve performance and manage system resources more efficiently. The Go implementation handles service management and system interaction more effectively than Bash alone.
Running the Go version:

To use the Go version of the script, follow these steps:

Build the Go program:

      go build -o mansyc main.go

Run the Go application:

       ./mansyc

The Go version will offer a more responsive and efficient management of services.

 

When the script runs, you will be prompted with two options:

   use arrow keys to select what you want to do
    🟢 to start services. (ENTER)
    🔴 to terminate unnecessary services.(DOWN-ARRWO + ENTER)

## Example Working Screenshots 📸
![image](https://github.com/user-attachments/assets/8e8f6bc6-cde8-4aac-9282-0cec58a563a4)

![image](https://github.com/user-attachments/assets/1519cd65-cfca-4105-87da-3206c61cc34d)

![image](https://github.com/user-attachments/assets/2560e5e8-8476-410c-81bc-ff230a69a17d)





## Dependencies 📦

   Linux-based system 🐧

   Go (for Golang integration 🖥️)
   Install from: Go installation

   systemctl (for managing systemd services 🔧)

   Utilities like grep, awk, sed (for process management 🔍)

## Future Improvements 🔮

   Golang Enhancements 🖥️: Further optimization and added features in Golang.

   GUI/CLI Enhancements 🎨: To simplify the usage experience.

   Logging Functionality 📊: To track service management actions.

## Contributing 🤝

   Contributions are welcome!
   Feel free to open an issue or submit a pull request.
   License 📜
   
   This project is licensed under the MIT License. See LICENSE for details.

