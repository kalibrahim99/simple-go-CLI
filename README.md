# Simple Go Calculator CLI

A simple command-line calculator written in Go. This project allows you to perform basic arithmetic operations (+, -, *, /) interactively from your terminal.

<p align="center">
  <img src="https://media0.giphy.com/media/v1.Y2lkPTc5MGI3NjExaWE5MTM4cHFvdnVkdHp1MmkyOWU5dDlodmt4amhibGNwam9hOW5jciZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/1F8FRRud3GiQnSKVLn/giphy.gif" alt="CLI Demo" width="700">
</p>

## Features

- Interactive CLI for calculations
- Supports addition, subtraction, multiplication, and division
- Maintains a running total
- Input validation and error handling
- Exit anytime with the `exit` command

## Usage

1. **Build the project:**
   ```sh
   go build -o calculator
   ```

2. **Run the calculator:**
   ```sh
   ./calculator
   ```
3. **if you want run with not build**
    ```sh
    go run main.go
    ```
4. **Follow the prompts:**
   - Enter operations in the format:  
     ```
     + 5
     - 2
     * 3
     / 4
     ```
   - Type `exit` to quit.

## Example

```
------ simple CLI calculator ------
current number : 0
Enter operation (e.g., + 5, * 2, / 4, or 'exit'): + 10
current number : 10
Enter operation (e.g., + 5, * 2, / 4, or 'exit'): * 2
current number : 20
Enter operation (e.g., + 5, * 2, / 4, or 'exit'): / 4
current number : 5
Enter operation (e.g., + 5, * 2, / 4, or 'exit'): exit
good day
5
------ End ------
```

## Project Structure

- `main.go` - Entry point for the CLI application
- `calc/calc.go` - Calculator logic

## Author

This is my first real project in Go. Feedback and contributions are welcome!
