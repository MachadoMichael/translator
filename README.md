# Branch Service Documentation

## Overview

The Branch Service is a simple HTTP server that provides a translation service. This service accepts text, source language, and target language as input, translates the text, and returns the translated result.

## API Endpoint

### `POST /translate`

**Description**: Translates a given text from the source language to the target language.

**Request**:

- **Method**: `POST`
- **Content-Type**: `application/json`

**Request Body**:

The request body should be a JSON object with the following fields:

```json
{
  "text": "string",
  "sl": "string",
  "tl": "string"
}

