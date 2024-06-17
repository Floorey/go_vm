Natürlich! Hier ist ein Beispiel für eine README-Datei auf Englisch, die du für dein Projekt verwenden kannst:

---

# Blockchain VM

This project is a virtual machine (VM) designed to execute smart contracts. The VM supports basic opcodes for arithmetic operations and memory management, with the capability to enforce gas limits for contract execution.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Directory Structure](#directory-structure)
- [Testing](#testing)
- [Scripts](#scripts)
- [Contributing](#contributing)
- [License](#license)

## Installation

To get started with the Blockchain VM, clone the repository and build the project:

```bash
git clone https://github.com/yourusername/blockchain-vm.git
cd blockchain-vm
./scripts/build.sh
```

## Usage

You can run the main application using:

```bash
./bin/blockchain-vm
```

This will execute the main entry point defined in `cmd/main.go`.

## Directory Structure

```
blockchain-vm/
│
├── cmd/                      # Main application
│   └── main.go               # Entry point
│
├── core/                     # Core VM logic
│   ├── vm.go                 # VM implementation
│   ├── context.go            # Execution context
│   └── opcodes.go            # Opcodes for the VM
│
├── scripts/                  # Scripts for tests and build process
│   ├── build.sh              # Build script
│   ├── test.sh               # Test script
│   └── lint.sh               # Linter script
│
└── tests/                    # Tests
    ├── vm_test.go            # Tests for the VM
    └── integration_test.go   # Integration tests
```

## Testing

To run all tests, use the provided script:

```bash
./scripts/test.sh
```

You can also run the tests directly using the `go test` command:

```bash
go test ./tests
```

## Scripts

### build.sh

Builds the application:

```bash
./scripts/build.sh
```

### test.sh

Runs all tests:

```bash
./scripts/test.sh
```

### lint.sh

Runs the linter:

```bash
./scripts/lint.sh
```

## Contributing

We welcome contributions! Please follow these steps to contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a new Pull Request.

Please make sure to add tests for any new features and run all tests before submitting your pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

This README provides a clear overview of the project, including installation instructions, usage, directory structure, testing instructions, available scripts, contributing guidelines, and license information. You can customize it further as needed for your specific project details.
