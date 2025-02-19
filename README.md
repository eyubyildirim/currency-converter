# Currency Converter

## Table of Contents

1. [Introduction](#introduction)
2. [Features](#features)
3. [Requirements](#requirements)
4. [Installation](#installation)
5. [Usage](#usage)

## Introduction

This project uses Go and Bubbletea to create a command-line tool to convert between 4 currencies (USD, TRY, GBP, EUR).
Exchange rates are obtained from Open Exchange Rates.

## Features

- Easy to use interface.
- Realtime currency conversion.
- -d tag for default currency conversion (1 USD to TRY).

## Requirements

- Golang 1.24 installed on the system.

## Installation

Step-by-step guide to getting the project up and running locally:

1. Clone the repository:
    ```bash
    git clone https://github.com/eyubyildirim/currency-converter
    ```
2. Navigate to the project directory:
    ```bash
    cd <project-directory>
    ```
3. Install dependencies:
    ```bash
    go mod download
    ```

## Usage

1. Run the project:
    ```bash
    go run main.go
    ```
2. Access the application or service as instructed.

You also need to provide an _appId_ via flags or environment variables. It can be obtained
from Open Exchange Rates website.
