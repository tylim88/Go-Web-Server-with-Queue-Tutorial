to setup

```bash
npm run setup
```

to run 

```bash
npm run back
```

then open a new terminal

```bash
npm run front
```

The original question: https://github.com/tylim88/se-take-home-assignment

My original frontend only solution: https://github.com/feedmepos/se-take-home-assignment/pull/69

This repository is part of my journey in learning Go.

The original question asked for a frontend OR backend solution; this repository provides both a frontend AND backend solution.

I also simplified the question, ignoring all the side requirements (scripts and tests) and focusing only on the main requirements.

The frontend is written in ReactJS and the backend is written in Go.

Real-time updates are enabled by using Server-Sent Events, so no polling is needed on the frontend.

There is one mistake in this repository: this stack can only serve one user simultaneously because I didn't broadcast the channels.

I will fix it when I have time.

This repository was coded under AI mentorship. The AI explained the language features and best practices, but it is 95% hand-coded.
