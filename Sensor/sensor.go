package main

import (
	"flag"
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//function to get the current temperature of the CPU based on the reading of the BIOS
func temperature(degrees, format string) string {
	//different execution based on OS "windows"/"linux"/"darwin"(MacOs)
	switch runtime.GOOS {
	case "windows":
		//On Windows the app should call external process wmic /namespace:\\root\wmi PATH MSAcpi_ThermalZoneTemperature get CurrentTemperature.
		output, _ := exec.Command("cmd", "/c", "wmic /namespace:\\\\root\\wmi PATH MSAcpi_ThermalZoneTemperature get CurrentTemperature").Output()
		//get the time of the measurement
		hour, minute, second := time.Now().Clock()
		//Manipulate the output to make it easily usable
		s := strings.Split(string(output), "\n")
		s[1] = strings.TrimSpace(s[1])
		int1, _ := strconv.Atoi(s[1])
		coreTemp := 0.00
		//switch to calculate degrees as raw data is in Kelvin
		switch degrees {
		case "C":
			coreTemp = float64(int1)/10 - 273.15
		case "F":
			coreTemp = float64(int1)*9/50 - 459.67
		}
		//check what format is specified in the command line arguments and return the output
		if format == "JSON" {
			return fmt.Sprintf("{\"measuredAt\":%02d:%02d:%02d, \"value\":%.2f, \"unit\":%s}.", hour, minute, second, coreTemp, degrees)
		}
		return fmt.Sprintf("measuredAt:%02d:%02d:%02d\nvalue:%.2f\nunit:%s", hour, minute, second, coreTemp, degrees)
	//TODO Implement Linux and MacOs functionality
	case "linux":
		return ""
	case "darwin":
		return ""
	}
	//return error if something failed
	return "error"
}

//function to show number of CPU cores, clock speed and CPU utilization
func usage(format string) string {
	//different execution based on OS "windows"/"linux"/"darwin"(MacOs)
	switch runtime.GOOS {
	case "windows":
		//On Windows the app should call external process WMIC CPU Get NumberOfCores,MaxClockSpeed,loadpercentage
		output, _ := exec.Command("cmd", "/c", "WMIC CPU Get NumberOfCores,MaxClockSpeed,loadpercentage").Output()
		//data manipulation for easier access
		outputString := string(output)
		s := strings.Fields(outputString)
		//returns different strings based on YAML/JSON command line argument
		if format == "JSON" {
			return fmt.Sprintf("{\"cores\":%s, \"frequency\":%sMhz, \"usedPercent\":%s\"}", s[5], s[4], s[3])
		}
		return fmt.Sprintf("cores:%s\nfrequency:%sMhz\nusedPercent:%s", s[5], s[4], s[3])
	//TODO Implement Linux and MacOS
	case "linux":
		return ""
	case "darwin":
		return ""
	}
	//returns error if something failed
	return "error"
}
func main() {
	//set command line arguments
	operationType := flag.String("type", "CPU_Temp", "What mode of operation you want the program to run in. Temperature measurement or CPU usage. Default is Temperature measurement")
	help := flag.Bool("help", false, "Displays a help message")
	degreeUnit := flag.String("unit", "C", "What degrees you want the program to measure. Options are C for Celsius, F for Fahrenheit. Default is Celsius")
	outputType := flag.String("format", "JSON", "What format do you want the output to be in. Options are JSON or YAML. Default is JSON")
	tts := flag.Int("duration", 60, "How often do you want measurement to be taken in seconds. Default is 60 - once a minute")
	//parse command line arguments inputed by the user.
	flag.Parse()
	//Checks if all flags have correct values
	FlagsCorrect := ((*operationType == "CPU_Temp") && (*degreeUnit == "C" || *degreeUnit == "F") && (*outputType == "JSON" || *outputType == "YAML") && (*tts > 0)) || (*operationType == "CPU_Usage")
	// checks if the user used the help option anywhere in the command line. If they did, ignore all other options and print a usage message
	if *help || !FlagsCorrect {
		fmt.Println("Usage: sensor [--help] [--type=<CPU_Temp, CPU_Usage>] [--unit=<C, F>] [--duration=<seconds>] [--format=<JSON, YAML>].")
		fmt.Println("Default is sensor --type=CPU_Temp --unit=C --duration=60 --format=JSON")
		fmt.Println("Duration should be a positive number.")
	} else {
		switch *operationType {
		case "CPU_Temp":
			/* endless loop. Prints the temperature, calculates time for temperature function to take measure.
			Then sleeps the program for the specified time in --duration minus the execution time for more accurate measuring times
			If taking the temperature takes more time than the specified --duration, it will do it as soon as possible */
			for {
				start := time.Now()
				fmt.Println(temperature(*degreeUnit, *outputType))
				elapsed := time.Since(start)
				durationTTS := time.Duration(*tts)*time.Second - elapsed
				time.Sleep(durationTTS)
			}
		case "CPU_Usage":
			/* endless loop. Prints the CPU Usage, calculates time for CPU usage function to take measure.
			Then sleeps the program for the specified time in --duration minus the execution time for more accurate measuring times
			If taking the temperature takes more time than the specified --duration, it will do it as soon as possible */
			for {
				start := time.Now()
				fmt.Println(usage(*outputType))
				elapsed := time.Since(start)
				durationTTS := time.Duration(*tts)*time.Second - elapsed
				time.Sleep((durationTTS))
			}
		}
	}
}
