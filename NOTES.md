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