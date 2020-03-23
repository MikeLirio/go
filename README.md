# README - GO - [Golang.org](GO_Official)

This is my personal workspace of GO. I will push everything related with GO that I made here.

*This README is written with [Typora][] software.*

## Repository Three

* first-try: introduction to the language.

## Notes

* All projects are located in the same folder called **workspace**.

## Useful Links

[Typora]: https://typora.io/ "Typora official Webpage"
[GO_Official]: https://golang.org/ "Official site of the language."

## Useful Videos

[YTB_01]: https://www.youtube.com/watch?v=C8LgvuEBraI "Learn Go in 12 Minutes"

## POWERSHELL useful code

To be able to have your own scripts on PowerShell you'll have to do 2 things:

* Create the next file and insert your scripts: `C:\Users\micha\Documents\WindowsPowerShell\Microsoft.PowerShell_profile.ps1`
* Write in the console the next command to allow PowerShell to execute scripts:
   `Set-ExecutionPolicy Unrestricted`

```powershell
##################################
# General Functions
##################################

function printenv () {
  Get-ChildItem Env:
}

function debugON($value) {
  $value = $($value + '*')
  $env:DEBUG = $value
  Write-Output $env:DEBUG
}

function debugOFF() {
  $env:DEBUG = ""
  Write-Output $env:DEBUG
}

function setNodeEnv($value) {
  if(!$value) {
    $env:DEBUG = "develop"
  } else {
    $env:NODE_ENV = $value
  }
  Write-Output $env:NODE_ENV
}

##################################
# GO functions
##################################

function goWorkspace() {
  Set-Location "D:\Desarrollo\go-workspace"
}
```

