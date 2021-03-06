package config

import (
	"os"
	"testing"

	"github.com/matryer/is"
)

func setTestEnv(key, val string) func() {
	os.Setenv(key, val)

	return func() {
		os.Unsetenv(key)
	}
}

func TestParsesConfigPipeHCL(t *testing.T) {
	is := is.New(t)

	c, err := ParseHCLFile("../test_fixtures/pipe/standard.hcl", nil)

	is.NoErr(err)             // error should have been nil
	is.Equal(1, len(c.Pipes)) // should have returned one pipe
}

func TestParsesConfigPipeHCLNoFail(t *testing.T) {
	is := is.New(t)

	c, err := ParseHCLFile("../test_fixtures/pipe/no_fail.hcl", nil)

	is.NoErr(err)                                          // error should have been nil
	is.Equal(1, len(c.Pipes))                              // should have returned one pipe
	is.Equal(0, len(c.Pipes["process_image_fail"].OnFail)) // should have returned 0 fail blocks
}

func TestParsesConfigPipeAndRetunsErrorWithInvalidExpiration(t *testing.T) {
	is := is.New(t)

	_, err := ParseHCLFile("../test_fixtures/pipe/invalid_expiration.bad", nil)

	is.True(err != nil) // should have returned an error
}

func TestParsesConfigPipeHCLNoSuccess(t *testing.T) {
	is := is.New(t)

	c, err := ParseHCLFile("../test_fixtures/pipe/no_success.hcl", nil)

	is.NoErr(err)                                                // error should have been nil
	is.Equal(1, len(c.Pipes))                                    // should have returned one pipe
	is.Equal(0, len(c.Pipes["process_image_success"].OnSuccess)) // should have returned 0 success blocks
}

func TestParsesFolder(t *testing.T) {
	is := is.New(t)

	c, err := ParseFolder("../test_fixtures", nil)

	is.NoErr(err)               // error should have been nil
	is.Equal(3, len(c.Pipes))   // should have returned three pipes
	is.Equal(2, len(c.Outputs)) // should have returned two output
	is.Equal(1, len(c.Inputs))  // should have returned one input
}

func TestParsesConfigContainingEnvironmentVariables(t *testing.T) {
	is := is.New(t)
	defer setTestEnv("abc", "123")()

	c, err := ParseHCLFile("../test_fixtures/pipe/interpolation.hcl.env", nil)

	is.NoErr(err)                                 // error should have been nil
	is.Equal("123", c.Pipes["environment"].Input) // should have set input to 123
}
