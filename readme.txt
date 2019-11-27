
This project is for the assessment at:
https://hackmd.io/@terminal1/assessment-conway#Skills-to-be-graded

It has been pre configured for docker and can be easily pushed to heroku

working instance is here:
https://frozen-harbor-01265.herokuapp.com/


Infrastructure,
Client : ReactJS + d3 
- just wanted to try out the newer Hooks in 16.8

Backend : Go 
- Game of Life seems to have Go written all over it, its abbr is GOL.
- picked it as my last node project's response time was horrible, was hoping Go would give better performance

Skeleton.

The Project is basicly added to an existing websocket Chat.
In the backend, additional elements was addeed to the pool struct
including:
- ticker
- gameHandle - which points to a game struct

The game struct would include the *gameBoard and 
on a side note, a *[]PlayerStruct was originally designed in to be implemented later, but seeing that its not part of the requirements, and may open up a pandora's box, decided to not implemented this time.

basically, when a message is received, it will be parse in the GOLMessage function, it will parse out the message type and perform the relevant function needed.
the messag types may inlclude : goMove, goChat, goPresets, goReset etc.

the backend will send out messages with a tag of the type of message too so the client will know what to do with it. chat/gameBoard
Also, on initial connection, the backend will send out a random color to the client

The client will branch out the messages to either color/chat/goboard and update the corresponding state.
The gameboard was done with d3, as I thought i might be more better to add functionalities on top in the future.

After Thoughts
the game starts up the timer once its loaded. maybe should consider starting the timer after getting the first connection, but this will have to keep track of player connections which, as stated above may open the pandora's box.

Any player can bump off (change the color) of an existing live cell, do not see if it is allowed or not, but I consier this feature as well as the aesthetics, a low priority


performance
It is said that golang websocket would be able to concurrently connect up to 2000 connections, but this isn't tested.















