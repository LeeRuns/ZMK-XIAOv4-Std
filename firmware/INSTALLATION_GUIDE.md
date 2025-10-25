# ZMK Firmware Installation Guide

## Quick Installation Steps

### Step 1: Download Firmware
1. Go to: https://github.com/LeeRuns/ZMK-XIAOv4-Std/actions
2. Find the latest successful workflow run for branch "NS-5x6+6-ENC-TB-(73)"
3. Click on the workflow run
4. Download the firmware artifacts (ZIP file)
5. Extract the ZIP file to the "archives" folder

### Step 2: Organize Firmware Files
The firmware archive should contain these files:
- skreecustom_left-seeeduino_xiao_ble-zmk.uf2 → Copy to "left" folder
- skreecustom_right-seeeduino_xiao_ble-zmk.uf2 → Copy to "right" folder

### Step 3: Flash Left Half
1. Connect LEFT half of keyboard via USB
2. Double-click reset button on XIAO BLE board
3. Copy left/skreecustom_left-seeeduino_xiao_ble-zmk.uf2 to USB storage device
4. Device will automatically restart

### Step 4: Flash Right Half
1. Connect RIGHT half of keyboard via USB
2. Double-click reset button on XIAO BLE board
3. Copy right/skreecustom_right-seeeduino_xiao_ble-zmk.uf2 to USB storage device
4. Device will automatically restart

### Step 5: Pair Halves
1. Reset both halves simultaneously (press reset buttons together)
2. Halves should automatically pair via Bluetooth
3. Test keyboard functionality

## Troubleshooting

### Bootloader Mode Issues
- Try different USB cables
- Try different USB ports
- Hold reset button longer (3-5 seconds)
- Check if XIAO BLE board has different reset procedure

### Flashing Issues
- Ensure .uf2 file is copied to root of USB storage device
- Don't rename the .uf2 file
- Check that file transfer completed successfully

### Pairing Issues
- Reset both halves simultaneously
- Check Bluetooth settings on computer
- Ensure both halves have power (battery or USB)

## Current Keymap Features
- **Default Layer**: Standard QWERTY layout
- **Layer 1**: Bluetooth management (hold MO(1) key)
- **Layer 2**: Additional functions (hold MO(2) key)
- **Encoder**: Volume control (left side)
- **Trackball**: Mouse cursor and scroll wheel
- **Bluetooth**: Multi-device support

Generated on: 2025-10-25 06:51:10
