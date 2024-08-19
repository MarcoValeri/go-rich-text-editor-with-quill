# go-rich-text-editor-with-quill
A simple web application built with Go and the Quill rich text editor. Follow along with the tutorial on [devwithgo.dev](https://www.devwithgo.dev/tutorial/how-to-create-rich-text-editor-field-with-go-and-quill) to learn how to create your own.

## Features
* Rich text editing with Quill's powerful toolbar
* Clean and minimal interface for distraction-free writing
* Easily save and load your documents 

## Prerequisites
* **Go:** Ensure Go is installed on your system. If not, download and install it from [go.dev/dl/](https://go.dev/dl/).

## Getting Started

1. **Clone the repository:**
   ```bash
   git clone https://github.com/MarcoValeri/go-rich-text-editor-with-quill.git
2. **Run the app in the local env**
    cd go-rich-text-editor-with-quill
    go mod init gorichtexteditorwithquill
    go run . 

## Project Structure
public/js/quill-editor.js - Build and manipulate Quill rich text editor
views/home.html - Frontend HTML, contains also Quill library
go.mod - Go module
main.go - main application logic