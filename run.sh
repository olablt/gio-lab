tmux send-keys -t mySession.0 C-c;
tmux send-keys -t mySession.0 C-l;
tmux send-keys -t mySession.0 "tmux clear-history" ENTER

# tmux send-keys -t mySession.0 "go run ." ENTER
# tmux send-keys -t mySession.0 "go run ./000-hello/" ENTER
# tmux send-keys -t mySession.0 "go run ./010-layout-stacked/" ENTER
# tmux send-keys -t mySession.0 "go run ./020-scroll-widgets/" ENTER
# tmux send-keys -t mySession.0 "go run ./000-black-on-white/" ENTER
# tmux send-keys -t mySession.0 "go run ./015-key-pointer-events/" ENTER
tmux send-keys -t mySession.0 "go run ./030-fuzzy/" ENTER
