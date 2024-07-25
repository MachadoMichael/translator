> ### Package `translator`
> 
> The `translator` package provides a function to translate text from one language to another using the Google Translate API.
> 
> #### Function `Translate`
> ```go
> func Translate(txt string, sl string, tl string) (string, error)
> ```
> Translates a given text from a source language to a target language.
> 
> **Parameters:**
> - `txt` (string): The text to be translated.
> - `sl` (string): The source language code (e.g., "en" for English).
> - `tl` (string): The target language code (e.g., "es" for Spanish).
> 
> **Returns:**
> - `string`: The translated text.
> - `error`: Any error encountered during the translation process.
> 
> **Example:**
> ```go
> translatedText, err := translator.Translate("Hello, world!", "en", "es")
> if err != nil {
>     fmt.Println("Translation error:", err)
> } else {
>     fmt.Println("Translated text:", translatedText)
> }
> ```
> 
> **Details:**
> - The function constructs the Google Translate API URL by replacing spaces in the input text with `%20`.
> - It sends a GET request to the Google Translate API.
> - Reads the response body and unmarshals the JSON response.
> - Extracts and returns the translated text from the JSON response.
> 
> **Note:**
> - Ensure to handle the potential errors when making HTTP requests and processing the response.
> - This implementation is simplified and assumes the translated text is always present in the response.
