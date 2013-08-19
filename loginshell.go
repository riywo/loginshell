package loginshell

import (
    "runtime"
    "fmt"
    "os"
    "os/exec"
    "errors"
    "regexp"
)

func Shell() (string, error) {
    switch runtime.GOOS {
        case "linux":
          return LinuxShell()
        case "darwin":
          return DarwinShell()
    }
    return "", errors.New("Undefined GOOS: " + runtime.GOOS)
}

func LinuxShell() (string, error) {
    return os.Getenv("SHELL"), nil
}

func DarwinShell() (string, error) {
    dir := "Local/Default/Users/" + os.Getenv("USER")
    out, err := exec.Command("dscl", "localhost", "-read", dir, "UserShell").Output()
    if err != nil { return "", err }

    re := regexp.MustCompile("UserShell: (/[^ ]+)\n")
    matched := re.FindStringSubmatch(string(out))
    shell := matched[1]
    if shell == "" { return "", errors.New(fmt.Sprintf("Invalid output: %s", string(out))) }
    return shell, nil
}
