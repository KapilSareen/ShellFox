<p align="center">
  <h1 align="center">ShellFox 🚀</h1>
<h3 align="center">the CLI browser of the future...or maybe not 🌐</h3>
</p>


**ShellFox** is an exciting project still in its early stages. While it’s not yet the next big web browser, it is a **fun exploration** into the internals of web technologies. The goal is to understand how browsers work, but in a minimalist, terminal-based environment.

This project is not focused on being the fastest or feature-rich browser; instead, it's about **exploring the core mechanics** of web fetching, parsing, and rendering, all while keeping things **lightweight and text-based**. Along the way, there’s fun with animations, interactive terminal features, and a touch of terminal magic. ✨



## Roadmap

### Completed ✅
- [x] **Basic TUI layout** with textarea, viewport, and spinner.
- [x] **Loading animations** with dynamic messages.
- [x] **Initial URL fetching functionality**.

### To-do 🚧
- [ ] **HTML Parsing**
  - [x] Parse the page and create the basic DOM tree with Text and Element nodes.
  - [ ] Extend Support for other nodes.
  - [ ] Extend Support for self closing tags
- [ ] **CSS Parsing**
  - [x] Implement a simple parser for Internal CSS. (Specificity pending)
  - [ ] Add support for external stylesheets and inline styles.
- [ ] **Render Tree Creation**
  - [ ] Combine DOM and CSSOM to create the Render Tree.
  - [ ] Ensure only visible elements are included in the Render Tree.
- [ ] **Layout (Reflow)**
  - [ ] Calculate the position and size of each element in the Render Tree.
  - [ ] Implement rules for layout (e.g., width, height, margins, padding, positioning).
- [ ] **Painting**
  - [ ] Paint each element with proper styles (e.g., text, images, borders).
  - [ ] Handle complex layers and z-index.

### Ambitious Goals 📝
- [ ] **JS Support**
  - [ ] Implement a basic JavaScript execution engine.
- [ ] **Accessibility Features**
  - [ ] Add keyboard navigation.
- [ ] **Offline Mode**
  - [ ] Cache previously fetched pages for quick reloads.
- [ ] **Improved Error Handling**
  - [ ] Provide detailed error messages.
  - [ ] Offer recovery options for failed requests.

---

## Why the hell I'm Building This Thing?

I'll be honest: I have no idea what I'm doing. But that's kind of the point.

I’ve always been curious about how web browsers work under the hood, and what better way to learn than by building my own (in a terminal, no less)? ShellFox is a little experiment, and I hope it grows into something that I'll have a lot of fun with. 

This isn't about building something perfect or fast; it's about learning, tinkering, and maybe finding a bit of joy in the weird and wonderful world of text-based interfaces. So, if you’re into that, or just want to see what happens when a web browser meets the terminal, come join me on this quirky ride!

Feel free to contribute, give feedback, or just laugh at the crazy stuff I try to implement (JavaScript in a terminal? Sure, why not?). 😎

---

## How to Use

1. Clone the repo:  
   `git clone https://github.com/KapilSareen/ShellFox.git`

2. Navigate to the project directory:  
   `cd ShellFox`

3. Install dependencies:  
   `go mod tidy`  
   `go mod vendor`
4. Build the project:  
   `go build -o shellfox ./cmd/main.go`

5. Run ShellFox:  
   `./shellfox`

---

Thanks for checking out **ShellFox**! 🌐🐾
