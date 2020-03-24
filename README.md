# README - GO - [Golang.org](GO_Official)

This is my personal workspace of GO. I will push everything related with GO that I made here.

*This README is written with [Typora][] software.*

## Repository Three

* 
* Repositories created based on video [seeing the this video from YouTube][YTB_01].
  * **hello_there**: classic hello world .
  * 

## Notes

* All projects are located in the same folder called **workspace**. Because of that, all project you do should go inside of the folder `src`.  The final three should something similar to this:

  ```paths
  gopath/user/go #This can be complete customized
  	|- bin
          ........
  		list_of_compiled_executables.exe
  		........
  	|- src
  		|-project1
  		|-project2
          |-........
  		|-projectN
  	|- pkg
  		|-...........
          |-folder_packages_downloaded
          |-...........
  ```

* All application that you want to be executed *must have a package call* ***main***.

* Is **important** to take a look on the packages from [golang webpage][GO_pkg].

* ***GO DOES NOT HAVE EXEPTIONS !!!!***

* ***GO USE POINTERS !!! :blue_heart:***

## Useful Links

[Typora]: https://typora.io/ "Typora official Webpage"
[GO_pkg]: https://golang.org/pkg/ "Useful links with a lot of libraries."
[GO_Official]: https://golang.org/ "Official site of the language."

## Useful Videos

[YTB_01]: https://www.youtube.com/watch?v=C8LgvuEBraI "Learn Go in 12 Minutes"

## Commands

```bash
# Execute the file indicated
go run {file.go}

# install the project on bin folden
go install 
```

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

