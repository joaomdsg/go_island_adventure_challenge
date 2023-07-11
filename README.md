# Island Adventure Challenge

Welcome to the Island Adventure Challenge project! This repository hosts a set of programming challenges in the style of Advent of Code. Each challenge simulates a situation that an intrepid explorer faces while discovering new islands, and your task is to help them by solving these challenges.

## Getting Started

These instructions will guide you on how to setup your local environment for the Island Adventure Challenge project.

### Prerequisites

Ensure the following are installed on your system:

- [Go](https://golang.org/dl/)
- [Git](https://git-scm.com/downloads)
- [VS Code](https://code.visualstudio.com/download) or your preferred IDE
- [inotify-tools](https://github.com/inotify-tools/inotify-tools/wiki) (Linux only)

### Installing and Running

1. Clone the repository: `git clone https://github.com/your_username/island-adventure-challenge.git`
2. Open the project in your preferred IDE. If you're using VS Code, simply run `code .` in the terminal from the root directory of the project.

### Test && Commit || Revert Workflow

To solve the challenges in the Test && Commit || Revert (TCR) style introduced by Kent Beck, follow these steps:

1. Write a test for your code.
2. Write the code to make the test pass.
3. Save your file.

When you save your file, the tests will run automatically. If they pass, the changes will be committed. If they fail, the changes will be reverted.

Here's how to set this up:

1. Ensure the `tcr.sh` script is executable. If it isn't, run `chmod +x tcr.sh`.
2. Open VS Code settings (File -> Preferences -> Settings or `Ctrl+,`).
3. Search for "Run on Save".
4. Check the box that says "Run On Save: Commands".
5. Add the following configuration to your `.vscode/settings.json` file (create this file in your `.vscode` directory if it doesn't exist):

```json
{
    "runOnSave.commands": [
        {
            "match": ".go$",
            "command": "./tcr.sh",
            "runIn": "terminal"
        }
    ]
}
```

This tells VS Code to run the `tcr.sh` script whenever a `.go` file is saved.

## Challenges

You can find the challenges in the `challenges/` directory. Each day's challenge is a separate Markdown file.

## Solutions

Your solutions should go in the `solutions/` directory. Each day should have its own directory with a `main.go` file for your code and a `main_test.go` file for your tests.

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.