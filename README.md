# t1-gol
# https://hackmd.io/@terminal1/assessment-conway

Conway’s Game of Life is a famous simulation that demonstrates cellular automaton. It is modeled as a grid with 4 simple rules:

Any live cell with fewer than two live neighbors dies, as if caused by under-population.
Any live cell with two or three live neighbors lives on to the next generation.
Any live cell with more than three live neighbors dies, as if by overcrowding.
Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
Create a multiplayer Web app version of Game of Life, with the following functions. At minimum, you must support the latest Google Chrome desktop browser, although you may be penalized if your game doesn’t work in other modern browsers. (For example, we grade using Google Chrome, but clients might open your work in Firefox or Safari.)

Implement the Game of Life browser frontend. You can use any representation such as <canvas>, simple DOM manipulation or even <table> cells. The game should tick automatically at a predefined interval, at say, 1 step per second.

The browser connects to a server, which allows multiple browsers to share the same, synchronized world view. Unless otherwise specified, the server may be written in Ruby, Node.js, or any other technology supported by Heroku. You may use any framework, e.g. Ruby on Rails, Sinatra, EventMachine, Hapi or just plain listening on a socket.

Each client is assigned a random color on initialization. From the browser, clicking on any grid will create a live cell on that grid with the client’s color. This change should be synchronized across all connected clients. (You can use any mechanism to achieve this, such as polling, comet, or WebSocket.)

When a dead cell revives by rule #4 “Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.”, it will be given a color that is the average of its neighbors (that revive it).

To make the evolution more interesting, include a toolbar that places some predefined patterns at random places with the player’s color, such as those found at here https://en.wikipedia.org/wiki/Conway’s_Game_of_Life#Examples_of_patterns (not necessary to implement all, just 3 - 4 is fine).

Skills to be graded
Algorithmic Programming
API and Service Design
Testing, CI/CD, and Site Reliability
UX Design and Prototyping
Web Frontend
Communication
Submission
You can find our grading guidelines at https://t1.gl/review.
Submit your solution at https://t1.gl/submit-assessment.
Pay specific attention to the following:
Correctness / Robustness – e.g. does it behave correctly if a client’s connection is unstable for a few seconds? Can we easily build and deploy?
Performance / Scalability – both in client side and server side.
Code Style / Terseness / Proper use of latest technologies.
UI aesthetics.
Your code must be ready to be deployed to Heroku without any issues.
Please spend no more than a few hours – if you run out of time, document what features are missing, and how you would approach them if you have more time.
FAQ
The description says “Each client is assigned a random color on initialization”. If the client disconnects and reconnects some time in the future, do they need to keep the same color?
This is unspecified. Use your best judgment. When we grade your solution, we will also look at your engineering choices and trade-offs.
What do you mean by the average of colors?
The mathematical average is sufficient, although feel free to use other definitions of color averaging. Remember to document your choices.


# https://hackmd.io/@terminal1/SkTMHmUpr#Review-of-Fullstack-Conway%E2%80%99s-Game-of-Life-Assessment---20191202




#test1
