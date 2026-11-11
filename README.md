# ascii-art-web

A small web page that turns your text into ASCII art, for anyone who wants to try ascii-art without the terminal. This project focuses on making the site more appealing, responsive and interactive, and on giving the user clear feedback, using CSS.

## Authors
- Imad Elalouaji (ImadeddineElalouaji)
- Zineb Nadak (zinebnadak)

## How it works
ascii-art-web is a Go HTTP server that gives a web GUI to our previous project, ascii-art.
You type a text, pick a banner (`standard`, `shadow` or `thinkertoy`), press the button, and the server sends back the page with your text drawn in ASCII art.

![frontend](images/frontend.png)

## Responsive design (Zineb)
- [ ] Works on phone (375px), tablet (768px) and desktop (1440px)
- [ ] No sideways page scroll, only the art box scrolls if the art is wide
- [ ] Title fits on one line on small screens
- [ ] Less empty space around the page on phones
- [ ] Banner buttons wrap neatly on small screens
- [ ] Input text is big enough to read on any screen
- [ ] Textarea only resizes vertically, so it can't overflow the screen

## Interactive design (Imad)
- [ ]
- [ ]
- [ ]

## Usage: how to run
You need Go installed.

```bash
git clone https://github.com/zinebnadak/ascii-art-web.git
cd ascii-art-web
go run .
```

Then open [http://localhost:8080](http://localhost:8080) in your browser.

| Method | Route        | What it does                                  |
|--------|--------------|-----------------------------------------------|
| GET    | `/`          | Shows the main page                           |
| POST   | `/ascii-art` | Receives `text` and `banner`, shows the art   |

Status codes:
- `200 OK` when everything works
- `400 Bad Request` for empty text, an unknown banner, or characters that are not printable ASCII
- `404 Not Found` for unknown pages, or when a template or banner file is missing
- `405 Method Not Allowed` when the wrong method is used on a route
- `500 Internal Server Error` for anything else that goes wrong

Built at [grit:lab](https://gritlab.ax/) Åland (01-edu peer-learning program).
