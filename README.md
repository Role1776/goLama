# 🦙 GoLama

<div align="center">

[![Go Reference](https://pkg.go.dev/badge/github.com/Role1776/goLama/lama.svg)](https://pkg.go.dev/github.com/Role1776/goLama/lama)
[![Go Report Card](https://goreportcard.com/badge/github.com/Role1776/goLama)](https://goreportcard.com/report/github.com/Role1776/goLama)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

</div>

> 🚀 A fast and simple Go library for interacting with APIs compatible with `llama.cpp` or `Ollama` for text generation, providing both streaming and synchronous modes of operation.

## ✨ Highlights

*   🎯 **Simple & Intuitive** - Clean API design that just works.
*   🛡️ **Robust Error Handling** - Comprehensive error management with detailed logging.
*   ⚡ **High Performance** - Optimized for speed and efficiency.
*  🔄  **Streaming Support** - Real-time response streaming capabilities.
*   🛠️ **Highly Configurable** - Flexible timeout settings.
* ⚙️ **Reliable:** Asynchronous processing through goroutines.


## 🚀 Quick Start

### Installation
```bash 
    go get github.com/Role1776/goLama
```

Markdown
Example Usage
Streaming Response

```go    
    package main

    import (
	    "fmt"
	    "time"

	    "github.com/Role1776/goLama"
    )

    func main() {
        // Set your LLM endpoint URL
        url := "http://localhost:11434/api/generate"
    
        // Configure your request
        model := "llama3"
        prompt := "Explain quantum physics in simple terms."
	    timeout := 0 * time.Second // 0 means using default timeout - 100 seconds

        fmt.Println("✨ Streaming Output:")
        // Get the response through streaming channels
        respChan, errChan := lama.GenerateResponse(url, model, prompt, timeout)
	    for {
		    select {
		    case resp := <-respChan:
                fmt.Print(resp.Response)
                if resp.Done{
                    fmt.Println("\n✨ Done Streaming.")
				    return
                }
		    case err := <-errChan:
			    if err != nil {
				    fmt.Println("❌ Error:", err)
				    return
			    }
		    }
	    }
    }
```
Synchronous Response
```go  
    package main

    import (
	    "fmt"
	    "time"

	    "github.com/Role1776/goLama/lama"
    )

    func main() {
        // Set your LLM endpoint URL
        url := "http://localhost:11434/api/generate"
    
        // Configure your request
        model := "llama3"
        prompt := "Explain quantum physics in simple terms."
	    timeout := 0 * time.Second // 0 means using default timeout - 100 seconds
        fmt.Println("✨ Synchronous Response:")
	    // Get all response as one string
        response := lama.SyncResponse(url, model, prompt, timeout)
	    if response != ""{
		    fmt.Println("✨ Response:", response)
		    fmt.Println("\n✨ Done.")
	    } else {
		    fmt.Println("❌ Error in  response (look at logs)")
	    }
    }
``` 
📚 API Reference
🔧 GenerateResponse
```go 
    func GenerateResponse(url string, model string, prompt string, timeout time.Duration) (<-chan models.Response, <-chan error)
``` 
Sends a request to the specified API endpoint to generate text and returns two channels, using streaming mode of response receiving.

Parameters
Name	Type	Description
url	string	Your API endpoint URL
model	string	LLM model identifier
prompt	string	Your input prompt text
timeout	time.Duration	Request timeout. (0 will use default timeout, 100 sec).
Returns
<-chan models.Response: A channel of type models.Response, used for streaming generated text

<-chan error: Channel for handling any encountered errors

🔧 SyncResponse
```go 
    func SyncResponse(url string, model string, prompt string, timeout time.Duration) string
``` 
Sends a request and returns the entire generated text as a string when completed, if failed will return the empty string.

Parameters
Name	Type	Description
url	string	Your API endpoint URL
model	string	LLM model identifier
prompt	string	Your input prompt text
timeout	time.Duration	Request timeout. (0 will use default timeout, 100 sec).
Returns
string: The generated text response; will return empty string if the request failed

🛡️ Error Handling
GoLama provides comprehensive error handling for:

🌐 Network request failures

📊 Invalid HTTP status codes

🔄 JSON processing issues

📝 Response handling

All errors are automatically logged and include detailed context for debugging.

⚙️ Configuration
```go 
    const defaultTimeout = 100 * time.Second
``` 
Customize the timeout setting to match your needs.

📦 Response Format
```go 
    {
        "response": "Generated text response",
        "done": true
    }
``` 

🤝 Contributing
Contributions make the open source community amazing! Feel free to:

🐛 Report bugs

💡 Suggest new features

🔧 Submit PRs

📄 License
MIT License - feel free to use this in your projects!

<div align="center">
Made with ❤️ by the Go community
</div>
