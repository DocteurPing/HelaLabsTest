# Hela Labs Technical Assessment

This Go program is designed to find words that can be constructed from a given set of letters. The program takes a set of letters as a command-line argument and reads a list of English words from a file named "words.txt". It then identifies and displays the words that can be constructed from the provided letters.

## Instructions

To run the program, follow these steps:

1. Ensure you have Go installed on your machine.

2. Clone or download the repository.

3. Open a terminal and navigate to the project directory.

4. Run the program using the following command:

    ```bash
    go run main.go <letters>
    ```

   Replace `<letters>` with the set of letters you want to use for finding matching words.

5. View the output, which will display the words that can be constructed from the provided letters.

## Example Usage

```bash
go run main.go yxmiasdaegw
```

This command will find and display the words that can be constructed from the letters "yxmiasdaegw".
## Additional Information

- The program reads words from a file named "words.txt." Ensure this file exists in the same directory as the executable or update the file path accordingly.

- The runtime of the program is measured, excluding any file reading or preparation time, to focus on the word matching performance.

- For testing purposes, you may replace the contents of "words.txt" with your own list of English words.

Feel free to reach out if you have any questions or need further clarification. Happy coding!