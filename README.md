# winclipboard
Windows剪贴板相关API / Windows clipboard related APIs

## Import
```
import "github.com/l3vi404/winclipboard"
```

# Example

## Read text content from clipboard
```
text, err = GetClipboardText()
if err != nil {
    fmt.Errorf("Faild GetClipboardText: %v", err)
}
fmt.Println(text)
```

## Write text content to the clipboard
```
err := SetClipboardText("Wishing for world peace！")
if err != nil {
    fmt.Errorf("Faild SetClipboardText: %v", err)
}
```

## Listen to the clipboard
```
AddClipboardUpdateListener(func(ctx any) {
    text, _ := GetClipboardText()
    fmt.Printf("The clipboard has been changed, content is: %q\n", text)
}, context.Background())

AddClipboardDestroyListener(func(ctx any) {
    fmt.Println("The clipboard destroyed")
}, context.Background())

go func() {
    if err := StartClipboardListener(); err != nil {
        fmt..Errorf("Faild StartClipboardListener: %v", err)
    }
}()

time.Sleep(10 * time.Second)
StopClipboardListener()
```