# ping-sound
Ping Sound service. It keeps the bluetooth speaker connected even without playing audio.

## Installation steps for your ping-sound service:

**Note:** This version uses a Bash script (`ping-sound.sh`) and does not require compiling a Go binary.

### Step 1: Copy the script and sound files to system locations
```sh
sudo cp ping-sound.sh /usr/local/bin/
sudo chmod +x /usr/local/bin/ping-sound.sh
sudo cp ping.wav /usr/local/share/ping-sound/
```

### Step 2: Copy the service file to systemd
```sh
sudo cp ping-sound.service /etc/systemd/system/
```

### Step 3: Reload systemd and enable the service
```sh
sudo systemctl daemon-reload
sudo systemctl enable ping-sound.service
```

### Step 4: Start the service (optional - to test immediately)
```sh
sudo systemctl start ping-sound.service
```

### Step 5: Check service status
```sh
sudo systemctl status ping-sound.service
```

### Step 6: View service logs (if needed)
```sh
sudo journalctl -u ping-sound.service -f
```

---

## The service is configured to:

- Start automatically after boot
- Restart automatically if it crashes
- Run as your user (`nfp`) to access audio properly
- Use proper environment variables for PulseAudio
- After installation, the service will start automatically on every boot and monitor for your Xiaomi Sound Pocket device!

---

### Script details

- The main logic is in [`ping-sound.sh`](ping-sound.sh), which checks if your Bluetooth device ("Xiaomi Sound Pocket") is connected and plays a ping sound every 30 seconds to keep it alive.
- The systemd service [`ping-sound.service`](ping-sound.service) runs the script and passes the sound file as an argument.
