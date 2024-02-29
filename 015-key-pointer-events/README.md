Testing new event logic with gioui.org@v0.5.0.

This demo app listens for events in different areas. 
It should react globally to a 'Q' press and respond to '1' or '2' keys on chart axes. 

Goals to resolve in this example:

+ Ensure there is a reaction to pressing 'Q'.
- Trigger a pointer.Leave event when the area is on the edge and the window loses focus.
- Determine how to prevent or cancel the app.DestroyEvent event."


## x11 focus events

- https://stackoverflow.com/questions/31438020/x11-event-when-app-loses-focus
    - https://tronche.com/gui/x/xlib/events/input-focus/

