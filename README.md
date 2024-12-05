# 🦙 GoLama

<div align="center">

[![Go Reference](https://pkg.go.dev/badge/github.com/Role1776/goLama.svg)](https://pkg.go.dev/github.com/Role1776/goLama)
[![Go Report Card](https://goreportcard.com/badge/github.com/Role1776/goLama)](https://goreportcard.com/report/github.com/Role1776/goLama)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

</div>

> 🚀 A lightning-fast Go library for seamless interaction with LLaMA language models through a clean and intuitive HTTP API interface.

## ✨ Highlights

- 🎯 **Simple & Intuitive** - Clean API design that just works
- 🛡️ **Robust Error Handling** - Comprehensive error management with detailed logging
- ⚡ **High Performance** - Optimized for speed and efficiency
- 🔄 **Streaming Support** - Real-time response streaming capabilities
- 🛠️ **Highly Configurable** - Flexible timeout and model settings

## 🚀 Quick Start

### Installation

```bash
go get github.com/Role1776/goLama
```

### Example Usage

```go
package main

import (
    "fmt"
    "github.com/Role1776/goLama/lama"
)

func main() {
    // Initialize your LLaMA endpoint
    url := "http://localhost:11434/api/generate"
    
    // Configure your request
    model := "llama3"
    prompt := "Explain quantum physics in simple terms."

    // Generate response
    response, err := lama.GenerateResponse(url, model, prompt)
    if err != nil {
        fmt.Println("❌ Error:", err)
        return
    }

    fmt.Println("✨ Response:", response)
}
```

## 📚 API Reference

### 🔧 GenerateResponse

```go
func GenerateResponse(url, model, prompt string) (string, error)
```

Generates an AI response using your specified LLaMA model.

#### Parameters
| Name | Type | Description |
|------|------|-------------|
| url | string | Your API endpoint URL |
| model | string | LLaMA model identifier |
| prompt | string | Your input prompt |

#### Returns
- `string`: The generated response
- `error`: Error information if the request fails

## 🛡️ Error Handling

GoLama provides comprehensive error handling for:
- 🌐 Network request failures
- 📊 Status code validation
- 🔄 JSON processing
- 📝 Response handling

All errors are automatically logged and include detailed context for debugging.

## ⚙️ Configuration

```go
const defaultTimeout = 100 * time.Second
```

Customize the timeout setting to match your needs.

## 📦 Response Format

```json
{
    "model": "llama3",
    "created_at": "2024-01-20T12:00:00Z",
    "response": "Generated text response",
    "done": true
}
```

## 🤝 Contributing

Contributions make the open source community amazing! Feel free to:

- 🐛 Report bugs
- 💡 Suggest features
- 🔧 Submit PRs

## 📄 License

MIT License - feel free to use this in your projects!

---

<div align="center">
Made with ❤️ by the Go community
</div>
