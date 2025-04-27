# 🐍 Terminal Snake Game in Go

This is a classic Snake game implemented in Go, running inside a terminal using
tview and tcell libraries.

It was created as a first project while learning Go and contains some
imperfections that are openly acknowledged and explained.

To run the project `cd` into the folder and run (go lang installation assumed):

```sh
go run main.go
```

The game starts with single snake element in the middle of the screen marked as
`*`. The snake starts moving once input via arrow keys is detected. The snake
grows when its head lands on tile with food `@` and the game is over when snake
collides with borders or with itself.

It is not the most polished snake version in the universe as the goal was to
learn Go concepts - not to make the best software.

## 🚀 Features

- Playable Snake game in the terminal.

- Real-time movement using arrow keys.

- Random food spawning across the arena.

- Snake grows after eating food.

- Wall and self-collision detection.

- Simple game over screen.

(the whole process has to be restarted once the game over screen is hit in order
to start a new game)

## ⚙️ Technologies Used

- Go (Golang)
- tview (Terminal UI library)
- tcell (Terminal event/input handling)

## 🛠️ Known Imperfections

- Design:
  - Heavy use of pointers instead of struct methods.
  - Some tight coupling between logic and rendering.
  - Simplified collision checking and random food generation without deeper
    validation (e.g., preventing food spawning inside the snake).
- Gameplay Bugs:
  - Known issue: Sometimes the snake seems to miss eating a food even when
    stepping on it. (likely small race between move and render cycle)
  - The status panel at the top does not update, it holds static placeholders
    right now only
- Other:
  - Certain parts of the code could be better modularized (everything lives in
    main package).
  - Some technical debt exists because polishing every small detail was not the
    goal — learning and completing the project was.
