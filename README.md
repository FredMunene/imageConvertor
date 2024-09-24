# Image Converter

A program that converts an image into grayscale and ASCII art, and saves both the converted grayscale image and the ASCII art representation as separate files.

## Features

+ Converts an input image to grayscale.
+ Generates ASCII art from the grayscale image.
+ Saves the output as:
    - A PNG file for the grayscale image.
    - A TXT file for the ASCII art.

## Usage
### Prerequisites

Before running the program, ensure that you have Go installed on your system.
### How to Run

1. Clone or download the repository containing the program.

2. Make sure you are in the correct directory, where the main Go file and other source files are located.

3. To run the program, use the following command format:

    ```bash
    go run . <image-file>
    ```
    For example:

    ```bash
        go run . lion.jpg
    ```

### Example Command

```bash
go run . lion.jpg
```

Both files will be created in the same directory where the program is executed.
### Sample Output

+ Grayscale Image: lion.png
+ ASCII Art: lion.txt


## Planned Features

Here are some features that are planned to be added in the future:
+ Create a Web-based application.
+ Support for Multiple Image Formats:
    - Add support for additional image formats such as PNG, GIF.
+ Adjustable ASCII Art Resolution:
    - Provide options to adjust the resolution of the ASCII art (e.g., more or fewer characters per line).
+ Customizable ASCII Characters:
    - Allow users to specify a custom set of ASCII characters to be used in the conversion.
+ Colored ASCII Art:
    - Generate colored ASCII art by preserving the color of the original image and mapping it to ANSI color codes.
+ CLI Enhancements:
    - Add more command-line options such as setting the output directory, customizing file names, and adjusting the grayscale conversion method.
+ Batch Processing:
    - Implement support for processing multiple images at once, with output filenames generated automatically.
## Contribution
Contributions are welcome! Here's how you can contribute to the project:
1. Fork the Repository.
2. Create a new branch for your feature or bug fix:
    ```bash
    git checkout -b feature-name
    ```
3. Make changes, push them and Create a Pull Request

## Code of Conduct

Please make sure to read and follow the [Code of Conduct](https://www.contributor-covenant.org/version/2/1/code_of_conduct/) before contributing.

## License

This project is licensed under the MIT License. See the [LICENSE](License) file for details.