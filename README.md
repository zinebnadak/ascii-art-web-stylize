# ascii-art-web-stylize

A small web page that turns your text into ASCII art, for anyone who wants to try ascii-art without the terminal. This project focuses on making the site more appealing, responsive and interactive, and on giving the user clear feedback, using CSS.

## Authors
- Imad Elalouaji (ImadeddineElalouaji)
- Zineb Nadak (zinebnadak)

## How it works
ascii-art-web-stylize is a Go HTTP server that gives a web GUI to our previous project, ascii-art.
You type a text, pick a banner (`standard`, `shadow` or `thinkertoy`), press the button, and the server sends back the page with your text drawn in ASCII art.

## Responsive design (Zineb)
- [x] Works on phone (375px), tablet (768px) and desktop (1440px) - open Cmd + Option + I
- [x] No sideways page scroll
- [x] Input text is big enough to read on any screen 
- [x] Textarea only resizes vertically
- [x] Error page with title and a hint to fix each status code

## Interactive design (Imad)
- [x] Real-time live ASCII art generation as you type or switch fonts using the Fetch API (no page reload required)
- [x] Live character counter tracking input length in real time
- [x] One-click "Copy Art" button using the Clipboard API with visual confirmation feedback
- [x] Interactive Dark Mode / Light Mode theme switcher, remembered across page loads
- [x] Live error feedback: invalid characters or an empty text show a message instead of stale art

## Usage: how to run and test
You need Go installed.

```bash
git clone https://github.com/zinebnadak/ascii-art-web-stylize.git
cd ascii-art-web-stylize
go run .
```

Then open [http://localhost:8080](http://localhost:8080) in Chrome to be able to use DevTools.

## Testing the error pages

| Code | How to trigger |
|------|----------------|
| 404 | Open http://localhost:8080/nope |
| 405 | Open http://localhost:8080/ascii-art |
| 400 | Type `é` in the textarea and click Generate |
| 404 (banner) | `mv banners/shadow.txt /tmp/`, generate with shadow, then `mv /tmp/shadow.txt banners/` |
| 500 | `chmod 000 banners/shadow.txt`, generate with shadow, then `chmod 644 banners/shadow.txt` |

## Routes

| Method | Route        | What it does                                  |
|--------|--------------|-----------------------------------------------|
| GET    | `/`          | Shows the main page                           |
| POST   | `/ascii-art` | Receives `text` and `banner`, shows the art   |

Built at [grit:lab](https://gritlab.ax/) Åland (01-edu peer-learning program).
