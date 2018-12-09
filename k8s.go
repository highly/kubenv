package main

import (
    "flag"
    "fmt"
    "github.com/fatih/color"
    "io"
    "os"
    "os/user"
    "strings"
)

var tip = "USAGE:   kubenv -e <extension name of config file> \n" +
    "EXAMPLE: cheate kubernetes config file in format of $HOME/.kube/config.xxx \n" +
    "         kubenv -e xxx"

var environment = flag.String("e", "", tip)

func main() {

    flag.Parse()

    megenta := color.New(color.FgCyan)
    red := color.New(color.FgRed)
    greenBold := color.New(color.FgGreen).Add(color.Bold)

    *environment = strings.TrimSpace(*environment)

    if *environment == "" {
        megenta.Println(tip)
        os.Exit(1)
    }

    // 获取用户 $HOME 路径
    machineInfo, err := user.Current()
    if err != nil {
        fmt.Println("> ", err.Error())
        os.Exit(2)
    }

    // k8s config file
    configFile := machineInfo.HomeDir + "/.kube/config"

    // wanted environment file
    targetFile := configFile + "." + *environment

    _, err = os.Stat(targetFile)
    if err != nil && os.IsNotExist(err) {
        red.Print("> ERROR: ")
        megenta.Printf("%s does not exist\n", targetFile)
        os.Exit(3)
    }

    _, err = CopyFile(configFile, configFile+"."+*environment)
    if err != nil {
        fmt.Println(err)
        os.Exit(4)
    }

    megenta.Print("> current kubernetes environment: ")
    greenBold.Println(*environment)
}

func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return 0, err
    }
    defer src.Close()

    dst, err := os.Create(dstName)
    if err != nil {
        return 0, err
    }
    defer dst.Close()

    return io.Copy(dst, src)
}
