goLama — Neural Network API Integration Library goLama is a lightweight Go library for integrating neural network APIs, focusing on Ollama. It simplifies the process of making requests and parsing multi-part responses, speeding up the integration of AI services into your projects.

Features Streamlined Integration: Quickly connect to APIs like Ollama without writing excessive boilerplate code. Timeout Management: Pre-configured HTTP client prevents indefinite hanging of requests. Error Logging: Built-in utility for tracking and debugging issues. Customizable Settings: Easily adjust default parameters to fit your application needs. Installation Before using this library, ensure Ollama is installed and running. Then, install goLama via Go modules:

	go get github.com/Role1776/goLama
Usage Example Here's an example of how to use goLlama to interact with Ollama's API:

	package main

	import (
    	"fmt"
    	"github.com/Role1776/goLama"
	)

	func main() {

    	// Define the API endpoint
    	url := "http://localhost:11434/api/generate"

		// Specify the model and the prompt
   			model := "llama3"
    	prompt := "Explain quantum physics in simple terms."

    	// Send a request to the API and get the response
    	response, err := goLlama.GenerateResponse(url, model, prompt)
   		if err != nil {
	   		// Log any error encountered
	    		fmt.Println("Error:", err)
	    		return
    	}

    	// Display the received response
   	 	fmt.Println("Response:", response)
	}
Key Points Request Automation: Automatically constructs payloads and manages headers. Timeout Control: Comes with a default 100-second timeout, adjustable for your API's responsiveness. Multi-Part Decoding: Ensures complete data retrieval for chunked responses. Configuration You can modify the timeout duration by changing the defaultTimeout constant in the library:

	const defaultTimeout = 100 * time.Second
GitHub Short Description A Go library for integrating neural network APIs like Ollama, offering tools for streamlined request handling, multi-part response parsing, and error logging.
