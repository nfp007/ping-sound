package main

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	deviceName = "Xiaomi Sound Pocket"
)

//go:embed ping.wav
var pingWavData embed.FS

// checkBluetoothConnected returns true if the device is connected
func checkBluetoothConnected() bool {
	cmd := exec.Command("bluetoothctl", "info")
	out, err := cmd.Output()
	if err != nil {
		return false
	}
	info := string(out)
	if !strings.Contains(info, deviceName) {
		return false
	}
	// Check for 'Connected: yes' in the device info
	lines := strings.Split(info, "\n")
	for _, line := range lines {
		if strings.Contains(line, deviceName) {
			for _, l := range lines {
				if strings.Contains(l, "Connected: yes") {
					return true
				}
			}
		}
	}
	return false
}

// playPing writes the embedded ping.wav to a temp file and plays it
func playPing() {
	data, err := pingWavData.ReadFile("ping.wav")
	if err != nil {
		fmt.Println("Failed to read embedded ping.wav:", err)
		return
	}

	tmpFile, err := os.CreateTemp("", "ping-*.wav")
	if err != nil {
		fmt.Println("Failed to create temp file for ping sound:", err)
		return
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write(data)
	if err != nil {
		fmt.Println("Failed to write ping sound to temp file:", err)
		tmpFile.Close()
		return
	}
	tmpFile.Close()

	cmd := exec.Command("aplay", tmpFile.Name())
	_ = cmd.Run()
}

func main() {
	fmt.Println("Monitoring Bluetooth device:", deviceName)
	for {
		if checkBluetoothConnected() {
			fmt.Println("Device connected. Sending ping.")
			playPing()
		} else {
			fmt.Println("Device not connected.")
		}
		time.Sleep(30 * time.Second)
	}
}
