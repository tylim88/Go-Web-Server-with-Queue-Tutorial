to setup

```bash
npm run setup
```

to run

```bash
npm run dev
```

original question: https://github.com/tylim88/se-take-home-assignment

my original frontend only solution: https://github.com/feedmepos/se-take-home-assignment/pull/69

This repository is part of my journey in learning Go

The original question ask for frontend OR backend solution, this repository provides frontend AND backend solution

I also simplify the question, ignoring all the side requirements(scripts and tests), focusing only the main requirements

frontend is written in ReactJS and backend is written in Go

Real time update is enabled by using server side event, so no polling is needed on front end

There is one mistake in this repository, this stack can only serve one user simultaneously because I didn't broadcast the channels.

Will fix it when I have time.