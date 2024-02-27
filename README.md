## Links

- Issues tracker https://todo.sr.ht/~eliasnaur/gio
- Mailing list https://lists.sr.ht/~eliasnaur/gio
- gio patches https://lists.sr.ht/~eliasnaur/gio-patches

## gioui.org@v0.5.0


### window events https://gioui.org/news/2023-11

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

