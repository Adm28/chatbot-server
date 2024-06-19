# Description
A chatbot server written in GoLang where a user can create room, broadcast messages. 
I'm building this for fun just to explore concepts such as goroutines and channels in golang. 

# Commands

- `/nick <name>` - get a name, otherwise user will staay anonymous
- `join <name>` - join a room, if room doesn't exist, a new room will be created. User can be only in one room at a time
- `/rooms` - show list of available rooms to join.
- `/msg <msg>` - broadcast message to everyone in the room
- `/quit` - disconnects from the chat server
 