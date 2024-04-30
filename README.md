# Contextify

Contextify is a command-line interface (CLI) tool that helps you generate a markdown representation of your codebase, making it easily comprehensible for Artificial Intelligence (AI) systems.

By providing a quick and structured context of your project, Contextify empowers AI to better understand and assist in writing code.

## Inspiration

The inspiration behind Contextify stems from the need to efficiently communicate the context of a project to AI systems.

When working on a codebase, developers often seek assistance from AI to generate code, provide suggestions, or offer insights.

However, without a proper understanding of the project's structure and context, AI may struggle to provide accurate and relevant support.

Contextify aims to bridge this gap by generating a markdown file that represents your codebase in a format that is easily digestible by AI.

By feeding this contextual information to AI systems, they can better grasp the project's structure, dependencies, and conventions, enabling them to generate more precise and contextually relevant code suggestions.

## Installation

To install Contextify globally and use it as a command-line tool, follow these steps:

1. Ensure you have Go installed on your system. If not, you can download and install Go from the official website: [https://golang.org](https://golang.org)

2. Open your terminal and run the following command to install Contextify:

```bash
go install github.com/dannyfranca/contextify@latest
```

```bash
contextify version
```

You should see the version number of Contextify printed in the terminal.

## Usage

To generate a markdown representation of your codebase using Contextify, follow these steps:

1. Open your terminal and navigate to the root directory of your project.

2. Run the following command:

```bash
contextify gen -o <output_file> -e <extensions> <directory_paths>
```

- `<output_file>`: Specify the path and filename for the generated markdown file. For example, `codebase.md`.
- `<extensions>`: Provide a comma-separated list of file extensions to include in the markdown file. For example, `go,md,json,html`.
- `<directory_paths>`: Specify one or more directory paths to include in the markdown file. For example, `./src ./docs`.

Example usage:

```bash
contextify gen -o codebase.md -e go,md,json,html ./src ./docs
```

3. Contextify will process the specified directories and generate a markdown file with the provided filename. The markdown file will contain a structured representation of your codebase, including the file paths and their contents.

4. You can now feed the generated markdown file to AI systems to provide them with the necessary context of your project.

## Contributing

Contributions to Contextify are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on the [GitHub repository](https://github.com/dannyfranca/contextify).

## License

Contextify is released under the [Apache License](http://www.apache.org/licenses).
