package main

import "fmt"

/*
Stack Last in first out
so the push function should take in a data type
and the pop since its the LIFO data structure will return a
data type and not take in a parameter
 */
type BrowserHistory struct {
	webUrls []string
}

func (b *BrowserHistory) Push(newUrl string) {
	b.webUrls = append(b.webUrls,newUrl)
}
func (b *BrowserHistory) Pop () string {
	stackLen := len(b.webUrls)
	lastElement := b.webUrls[stackLen -1]
	b.webUrls = b.webUrls[:stackLen -1]
	return lastElement
}

func main()  {
	browserHistory := BrowserHistory{}

	// Push
	fmt.Println("The history before: ",browserHistory)
	browserHistory.Push("https://google.com")
	browserHistory.Push("https://facebook.com")
	browserHistory.Push("https://bing.com")
	fmt.Println("The history now: ",browserHistory)

	// Pop
	browserHistory.Pop()
	fmt.Println("The history after pop: ",browserHistory)
}