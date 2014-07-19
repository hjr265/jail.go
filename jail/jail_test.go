package jail

import (
	"os"
)

func ExampleHello() {
	jail, err := New("", 16)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := jail.Dispose()
		if err != nil {
			panic(err)
		}
	}()

	proc, err := jail.StartProcess("/bin/echo", []string{"echo", "Hajimemashite"}, &os.ProcAttr{
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	})
	if err != nil {
		panic(err)
	}
	_, err = proc.Wait()
	if err != nil {
		panic(err)
	}

	// Output: Hajimemashite
}
