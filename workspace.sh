#!/bin/bash

session=$1

[ "$session" = "" ] && "The session name (first argument) is required" && exit 1

[ ! -d "./.devbox" ] && devbox install

tmux new-session -d -s "$session"

tmux rename-window -t 1 'IDE'

tmux send-keys -t 'IDE' 'devbox run nvim' C-m

tmux new-window -t $session:2 -n 'Server'

tmux send-keys -t 'Server' 'devbox run start' C-m

tmux new-window -t $session:3 -n 'Shell'

tmux send-keys -t 'Shell' 'devbox shell' C-m
