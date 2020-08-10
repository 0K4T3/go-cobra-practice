package main

import (
    "go-cobra-practice/cmd"
)

func main() {
    cmd.SetArgs([]string{"hoge", "fuga"})
    cmd.Execute()
}
