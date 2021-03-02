# CPUUtil
Utility for monitoring CPU on different OS

## Instalation guide
 If you do not have go follow instructions on https://golang.org/doc/install
 Clone the repository.
 Go into Sensor folder
 run it with command "go run sensor" for default values
 
## Usage guide
 General usage is sensor --help --type=<CPU_Temp,CPU_Usage> --duration=<seconds> --format=<JSON,YAML> --unit=<C,F> \n
 For help on how to use it use "go run sensor --help" \n
 Sensor has two modes of operation denoted by --type<CPU_Temp,CPU_Usage> \n
 To run it in Temperature mode use "go run sensor --type=CPU_Temp --duratin=<seconds> --format=<JSON,YAML> --unit=<C,F>" \n
 and change the variables to ones you would like \n
 To run it in Usage mode use "go run sensor --type=CPU_Usage --duration=<seconds> --format=<JSON,YAML>" \n
 and change the variables to ones you would like. \n
 Duration is time between two measurements.Unit is the degrees you want the result to be in. \n
 Sensor outputs on terminal in format that you specified. You can stop it on most OS by pressing Ctrl + C or terminating the process \n
