# ping-sound
Ping Sound service. It keeps the bluetooth speaker connected even without playing audio.

## Here are the complete installation steps for your ping-sound service:

Step 1: Copy the binary to system location
```
sudo cp /home/nfp/sound/ping-sound /usr/local/bin/
sudo chmod +x /usr/local/bin/ping-sound
```

Step 2: Copy the service file to systemd
```
sudo cp /home/nfp/sound/ping-sound.service /etc/systemd/system/
```

Step 3: Reload systemd and enable the service
```
sudo systemctl daemon-reload
sudo systemctl enable ping-sound.service
```

Step 4: Start the service (optional - to test immediately)
```
sudo systemctl start ping-sound.service
```

Step 5: Check service status
```
sudo systemctl status ping-sound.service
```

Step 6: View service logs (if needed)
```
sudo journalctl -u ping-sound.service -f
```


## The service is configured to:

- Start automatically after boot
- Restart automatically if it crashes
- Run as your user (nfp) to access audio properly
- Have proper environment variables for PulseAudio
- After installation, the service will start automatically on every boot and monitor for your Xiaomi Sound Pocket device!
