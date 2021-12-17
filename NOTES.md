# Notes

Writing a test is just like writing a function, with a few rules
- It needs to be in a file with a name like xxx_test.go
- The test function must start with the word Test
- The test function takes one argument only t *testing.T

`go doc <symbol>` to get documentation about a package, function, or any other symbol in std lib or our own code

`t.Run` inside a test to execute subtests

`t.Helper()` is needed to tell the test suite that this method is a helper. By doing this when it fails the line number reported will be in our function call rather than inside our test helper.`

TDD approach:
- Write a test
- Make the compiler pass
- Run the test, see that it fails and check if the error message is meaningful
- Write enough code to make the test pass
- Refactor 

named return values:
```go
func greetingPrefix(language string) (prefix string) {
    prefix = "something"
    return
}
```
prefix is already declared and return (without a variable) can be called

publish documentation and examples to pkg.go.dev with a public available godoc instance
`godoc -http=:6060 and navigate to http://localhost:6060/pkg/`

`%.2f` fmt format for two decimal places in a floating point


`errcheck`(go get -u github.com/kisielk/errcheck) for unchecked errors

maps can only hold comparable types has keys. see https://go.dev/ref/spec#Comparison_operators

maps are always pointers, even when they are passed by value (https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it). this means we need to be careful when changing a map between functions

maps can be `nil`. adding to it causes runtime panic. reading it will return empty map

writing to an existing key will override its value

very interesting that since we have full control over a mock struct, we are able to assert how many times a method has been called (a struct count field for example). See [countdown.go](mocking/countdown.go)

**race condition**: output is dependent on the timing and sequence of events that we have no control over. https://en.wikipedia.org/wiki/Race_condition

go's builtin race detector: https://go.dev/blog/race-detector

structs can have anonymous fields (cool when we don't know what name to give to bold see [concurrency.go](concurrency/concurrency.go) `result` struct

at the time of writing, go supports any type (similar to the generic type) with the type `interface {}` (empty interface). cons: we lose type safety

`sync.Mutex` zero value is an unlocked mutex