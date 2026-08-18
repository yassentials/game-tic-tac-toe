# Tic-Tac-Toe using Typescript 

With my limited understanding of Typescript, I tried to create a simple Tic-Tac-Toe game.

## Diagram

```mermaid
flowchart TD
  A[Player] --> B((Enter Game))
  B --> C[Offline]
  B --> D[Offline Bot]
  B ------> E[Online]
  C --> F[Create Room] --> AP1[Add PLAYER 1 and PLAYER 2]
  D --> FF[Create Room] --> AB[Add BOT and PLAYER]
  E --> JR[Join Room]
  JR --> JC[Join with Code]
  JR --> JD[Join Random]
  E --> CN[Create New]
  CN --> PR[Private Room with Code]
  CN --> PUR[Public Room]
```

## Features

- Bot mode (play against bot)
- 2 Player mode (play against yourself)
- Online mode (not yet implemented!)

## Prerequisites

Copy the gen/ directory from server/ to client/

## TODO

- Add online mode

![Tic Tac Toe screenshot](./docs/images/screenshot1.jpeg)
