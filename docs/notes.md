### Notes

- Might need a state manager, something to handle what the current state is: init, active, break, quit
- it should definitely have a state and a stringified version of that state. Maybe call it repr
- `State` should also be a struct with the `type`, `repr`, and `symbol`
- Make the command sessions goroutines so the user can stop them by sending a signal.
