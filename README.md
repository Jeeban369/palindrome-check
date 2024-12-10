# 🛠️ Palindrome Checker

A simple and interactive Go program to check whether a given string is a **Palindrome**. A palindrome is a word, phrase, or sequence that reads the same backward as forward, ignoring spaces, punctuation, and case.

## Features
- Normalizes the input by removing non-letter characters and converting to lowercase.
- Checks for palindrome status with a simple and efficient algorithm.
- Provides a user-friendly interface with clear feedback.

## How It Works
1. The user enters a string.
2. The program:
   - Normalizes the input (removes non-letter characters and converts to lowercase).
   - Compares characters from the beginning and the end of the string.
3. Outputs whether the string is a palindrome.

## Example Usage
```plaintext
🛠️  CHECK YOUR PALINDROME HERE
▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️

🗯️  Type some string: Madam

It is a Palindrome. ✅
```

Another example:
```
🛠️  CHECK YOUR PALINDROME HERE
▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️▫️

🗯️  Type some string: Hello

it is not a palindrome. ❎
```

## Installation and Execution
### Prerequisites
- Install [Go](https://golang.org/dl/) on your system.

### Steps to Run
1. Clone this repository:
   ```bash
   git clone https://github.com/yourusername/Palindrome-Checker.git
   ```
2. Navigate to the project directory:
   ```bash
   cd Palindrome-Checker
   ```
3. Run the program:
   ```bash
   go run main.go
   ```

## Code Highlights
- **Normalization**: Non-alphabet characters are ignored, and the comparison is case-insensitive.
- **Efficiency**: Only the first half of the string is compared with the reversed second half.
- **Clear Output**: Uses emojis and concise messages for a user-friendly experience.

## Contributing
Feel free to fork this repository and submit pull requests with improvements or new features.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```

This `README.md` ensures clarity and professionalism while explaining the project comprehensively.
