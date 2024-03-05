[toc]

## GioUI Links

- Issues tracker https://todo.sr.ht/~eliasnaur/gio
- Mailing list https://lists.sr.ht/~eliasnaur/gio
- gio patches https://lists.sr.ht/~eliasnaur/gio-patches


## GioUI Links 

### window events https://gioui.org/news/2023-11

gioui.org@v0.5.0

```bash
gofmt -w -r "<-w.Events() -> w.NextEvent()"

gofmt -w -r "pointer.Type -> pointer.Kind"
gofmt -w -r "gesture.ClickType -> gesture.ClickKind"
gofmt -w -r "gesture.TypePress -> gesture.KindPress"
gofmt -w -r "gesture.TypeClick -> gesture.KindClick"
gofmt -w -r "gesture.TypeCancel -> gesture.KindCancel"
```


### event filters https://gioui.org/news/2024-02

```bash
gofmt -w -r 'op.InvalidateOp{}.Add(gtx.Ops) -> gtx.Execute(op.InvalidateCmd{})' .

gofmt -w -r 'system.DestroyEvent -> app.DestroyEvent' .
gofmt -w -r 'system.FrameEvent -> app.FrameEvent' .
gofmt -w -r 'layout.NewContext -> app.NewContext' .
gofmt -w -r 'system.Insets -> app.Insets' .
```


## overhaul

### Links

- Overhaul of package app https://todo.sr.ht/~eliasnaur/gio/555
- Overhaul of event routing https://todo.sr.ht/~eliasnaur/gio/550


## Styling References 

### GitHub Primer

- colors https://primer.style/foundations/color/overview
- https://primer.style/components/action-list
    - Action list items support three different sizes: small, medium, and large
    - On touch devices, the large size is used at all times to ensure usability when tapping.
- https://primer.style/components/action-menu
    - An action menu comprises a set of action list items
- https://primer.style/components/selectpanel
    - scrollable list consists of action list items


### Foundation Zurb

- https://get.foundation/sites/docs/kitchen-sink.html#0
