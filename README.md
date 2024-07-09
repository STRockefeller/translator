# Translation Tool

This tool uses multiple translation APIs (including DeepL and NiuTrans) to translate text between various languages. It supports progress saving to allow resuming translation after interruptions.

## Features

1. **Text Translation**: Supports translation between multiple languages via HTTP protocol.
2. **Multiple API Support**: Configurable with multiple API keys, including DeepL and NiuTrans, and automatically switches keys when quota is exceeded.
3. **Progress Saving**: Automatically saves translation progress to ensure continuity after interruptions.
4. **Error Retrying**: Automatically retries translatable errors after a specified delay.

## Configuration File (config.json)

The configuration file contains all necessary settings, including API keys, retry settings, language settings, etc.

Example `config.json` file:

```json
{
    "ApiKeys": {
        "NiuTrans": ["your_niutrans_api_key1", "your_niutrans_api_key2"],
        "Deepl": ["your_deepl_api_key1", "your_deepl_api_key2"]
    },
    "RetryDelay": 5,
    "MaxRetries": 3,
    "SourceLang": "en",
    "TargetLang": "zh",
    "InputFilePath": "input.txt",
    "OutputFilePath": "output.txt"
}
```

### Configuration Parameters

- `ApiKeys`: Lists of API keys for each service.
- `RetryDelay`: Delay between retries in seconds.
- `MaxRetries`: Maximum number of retries.
- `SourceLang`: Source language.
- `TargetLang`: Target language.
- `InputFilePath`: Path to the input file.
- `OutputFilePath`: Path to the output file.

## Usage

### Command Line Arguments

Command line arguments can override the settings in the configuration file:

- `-config`: Path to the configuration file (default is `config.json`).
- `-source`: Source language.
- `-target`: Target language.
- `-input`: Path to the input file.
- `-output`: Path to the output file.
- `-progress`: Path to the progress file (default is `progress.json`).
- `-h`: Show help information.

### Example

Assume we have a file named `input.txt` with the following content:

```
Hello, how are you?
This is a test.
```

We can translate it using the following command:

```sh
go run main.go -source en -target zh -input input.txt -output output.txt
```

### Progress Saving

The tool automatically saves translation progress to a `progress.json` file. If the translation process is interrupted, rerun the tool to resume translation from where it left off.

## Supported Languages

Currently supported languages include, but are not limited to:

- Chinese (`zh`)
- English (`en`)
- Albanian (`sq`)
- Arabic (`ar`)
- Amharic (`am`)
- Achuar (`acu`)

For a full list of supported languages, please refer to the API documentation.