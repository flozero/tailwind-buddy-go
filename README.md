# Tailwind component library with GO

![Screenshot 2024-09-28 at 12 48 32 AM](https://github.com/user-attachments/assets/9e405c8e-5730-4355-a4d8-808945e7af87)

## Context

Recently, I developed a Node.js library to challenge [CVA](https://beta.cva.style/getting-started/whats-new). My library is called [TailwindBuddy](https://github.com/busbud/tailwind-buddy).

I set a personal challenge to create something similar in Go within just 3 hours. Having completed _The Tour of Go_ before (The interactive GUI), I wanted to refresh my memory and understand the basics, as I’ve gone through it in the past.

## Goal

- To build a stack in ≤ 3 hours using Go, a basic templating system, HTMX, and TailwindCSS.
- To ensure a developer experience (devX) that's as friendly as possible.

## Progress

Since I’m still new to Go, I struggled a lot with typing, mostly because I’ve been thinking in TypeScript.

I aimed to replicate a devX experience close to what I’m used to in Node.js. I’m sure Go professionals will probably cringe at my code, but that’s okay! Instead of criticism, I’d love to hear how you would improve it. Feel free to leave some feedback or suggestions via GitHub issues.

Despite my deadline, the project is actually working well even If I didn't touch at all to HTMX but not necessary here.

However, I couldn’t find a better Go templating system that handles it quite the way I envisioned—more like `template strings` in JS. I’m aware that Go doesn’t have that feature natively.

If you think this project deserves more attention, please consider starring it or contributing to the discussion [here](https://github.com/flozero/tailwind-buddy-go/issues/1).

Working on this project has inspired me to dive deeper into Go not necessary on this project. My next reads are:

- Let's Go
- Let's Go Further

## run the project

`go run cmd/main.go`
