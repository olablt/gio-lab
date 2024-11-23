# Modal Window Example

This example demonstrates how to implement a basic modal window in Gio UI with keyboard controls.

## Features

- Shows a centered modal window overlay
- Keyboard shortcuts:
  - Press `Ctrl+O` to open the modal
  - Press `ESC` to close the modal
- Modal stays centered on screen
- Main window content remains visible but inactive behind modal
- Demonstrates:
  - State management for UI elements
  - Keyboard event handling
  - Layout composition with background and overlays
  - Centering and positioning of UI elements

## Implementation Details

The modal is implemented using:
- Background layout to layer the modal over main content
- Keyboard event filtering for shortcuts
- Boolean state tracking for modal visibility
- Center layout for positioning
