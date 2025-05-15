# GoTemp - Your Go-Powered Temperature Conversion Tool

## A Little Backstory (Just a Quick One!)

Born out of a sudden urge to quickly convert temperatures without fumbling through online converters, GoTemp was whipped up as a simple command-line tool. The goal? Efficiency and ease of use right in your terminal.

## Introduction

GoTemp is a lightweight and straightforward command-line application written in Go that allows you to convert temperatures between Celsius, Fahrenheit, and Kelvin. It's designed to be quick to use and integrate into your workflow whenever you need a fast temperature conversion.

## How to Build & Use

### Prerequisites

* **Go installed:** Make sure you have Go (version 1.16 or later) installed on your system. You can download it from [https://go.dev/dl/](https://go.dev/dl/).

### Building

1.  **Clone the repository (if you haven't already):**
    ```
    git clone https://github.com/amirparsadd/GoTemp.git
    cd GoTemp
    ```

2.  **Build the Go application:**
    ```
    go build .
    ```
    This will create an executable file named `gotemp` (or `gotemp.exe` on Windows) in your current directory.

### Usage

Open your terminal and navigate to the directory where the `gotemp` executable is located. Then, run the command with the temperature value, the original unit, and the target unit:

```
./gotemp <value> <from_unit> <to_unit>
````

**Arguments:**

  * `<value>`: The numerical temperature value you want to convert.
  * `<from_unit>`: The original unit of the temperature (C for Celsius, F for Fahrenheit, K for Kelvin).
  * `<to_unit>`: The unit you want to convert the temperature to (C, F, or K).

**Examples:**

  * Convert 25 Celsius to Fahrenheit:

    ```
    ./gotemp 25 C F
    ```

    Output:

    ```
    25C -> 77F
    ```

  * Convert 68 Fahrenheit to Celsius:

    ```
    ./gotemp 68 F C
    ```

    Output:

    ```
    68F -> 20C
    ```

  * Convert 300 Kelvin to Celsius:

    ```
    ./gotemp 300 K C
    ```

    Output:

    ```
    300K -> 26.85C
    ```

  * Invalid arguments will result in an error message and instructions on how to use the tool.

## Plans and To-Dos

This is just the beginning\! Here are some potential future enhancements:

  * **More Unit Support:** Explore adding support for other temperature scales like Rankine or Delisle.
  * **Interactive Mode:** Implement a mode where the user can input values and units without command-line arguments.
  * **Error Handling Improvements:** Provide more specific and user-friendly error messages.
  * **Configuration Options:** Allow users to potentially configure default units or output formats.
  * **Testing:** Add unit tests to ensure the accuracy and reliability of the conversions.

## Contribution

Feel free to contribute to GoTemp\! If you have ideas for improvements, find bugs, or want to add new features, please:

1.  **Fork the repository.**
2.  **Create a new branch** for your changes (`git checkout -b feature/your-feature` or `git checkout -b bugfix/your-fix`).
3.  **Make your changes and commit them** (`git commit -am 'Add some feature'`).
4.  **Push to the branch** (`git push origin feature/your-feature`).
5.  **Create a new Pull Request** on GitHub.

We appreciate any contributions that make GoTemp even better\!