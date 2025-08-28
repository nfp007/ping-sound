#!/bin/bash

# Configuration
DEVICE_NAME="Xiaomi Sound Pocket"

# Check if sound file parameter is provided
if [ $# -eq 0 ]; then
    echo "Usage: $0 <sound-file>"
    echo "Example: $0 debug-ping.wav"
    exit 1
fi

SOUND_FILE="$1"

# Check if sound file exists
if [ ! -f "$SOUND_FILE" ]; then
    echo "Error: Sound file '$SOUND_FILE' not found"
    exit 1
fi

echo "Monitoring Bluetooth device: $DEVICE_NAME"
echo "Using sound file: $SOUND_FILE"

# Function to check if device is connected
check_bluetooth_connected() {
    bluetoothctl info | grep -q "$DEVICE_NAME" && bluetoothctl info | grep -q "Connected: yes"
}

# Function to play ping sound
play_ping() {
    echo "Playing sound: $SOUND_FILE on Xiaomi Sound Pocket only"
    # Use the sink name for Xiaomi Sound Pocket
    SINK_NAME="bluez_output.AC_EF_92_C6_8D_78.1"
    paplay --device="$SINK_NAME" "$SOUND_FILE"
    if [ $? -eq 0 ]; then
        echo "Sound played successfully on Xiaomi Sound Pocket"
    else
        echo "Error playing sound on Xiaomi Sound Pocket"
    fi
}

# Main loop
while true; do
    if check_bluetooth_connected; then
        echo "Device connected. Sending ping."
        play_ping
    else
        echo "Device not connected."
    fi
    sleep 30
done