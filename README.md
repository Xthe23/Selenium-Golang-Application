# Selenium Golang Application

This repository contains a Selenium WebDriver application written in Go. The application is designed to interact with dynamic JavaScript websites, automating the process of navigating and extracting information from these sites.

## Features

- Utilizes Selenium WebDriver for browser automation.
- Written in the Go programming language.
- Configured to run in a Debian/Ubuntu environment within WSL (Windows Subsystem for Linux).
- Customizable Chrome binary path set as an environment variable.

  
## Getting Started

Follow these steps to set up the Selenium Golang Application on a Debian-based system:

### 1. Download Chrome & ChromeDriver

- [Download Chrome & ChromeDriver (Ensure Both are the Same Version)](https://googlechromelabs.github.io/chrome-for-testing/)

### 2. Extract the Downloaded Files

- Extract both the downloaded Chrome and ChromeDriver archives to a directory of your choice.

### 3. Set the `CHROME_BINARY_PATH` Environment Variable

  - Find the path where you decided to extract the Chrome binary.
  - Open your `~/.zshrc` file (or `~/.bashrc` if using bash) in a text editor.
    ```bash
      vim ~/.zshrc 
    ```
  - Add the following line to the end of the file, replacing /path/to/chrome with the actual path to the Chrome binary:

    ```bash
      export CHROME_BINARY_PATH="/path/to/chrome" 
      ```
   
  - Save the file and reload your ~/.zshrc configuration.

    ```bash
    source ~/.zshrc
    ```
### 4. Set the `chromedriver` Environment Variable
  - Find the path where you decided to extract the Chrome Driver binary.
  - Open your `~/.zshrc` file (or `~/.bashrc` if using bash) in a text editor.
    ```bash
    vim ~/.zshrc 
    ```
  - Add the following line to the end of the file, replacing /path/to/chrome with the actual path to the Chrome binary:
    ```bash
     export PATH="$PATH:/path/to/chromedriver"
    ```
  - Save the file and reload your ~/.zshrc configuration.
    ```bash
      source ~/.zshrc
    ```
### 5. Verify the Configuration
  - Verify that the CHROME_BINARY_PATH environment variable is set correctly:
    ```bash
      echo $CHROME_BINARY_PATH
    ```
  - Verify that chromedriver is in your PATH and the chrome driver version:
    ```bash
      which chromedriver && chromedriver --version
    ```
    
Now, you are ready to run the Selenium Golang Application!



  



