package backend

type WailsTest struct{}

func (w *WailsTest) GoHello() string {
	return "Hello from Go file!"
}

func (w *WailsTest) GoSay(text string) string {
	return ("Here is your text: " + text)
}

// go to app.go, add the desired function we want in there
