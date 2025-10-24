# Firmware Installation Guide

## Step 1: Download the Firmware

1. **Go to your GitHub repository**: https://github.com/leeruns/ZMK-XIAOv4-Std
2. **Click on the "Actions" tab**
3. **Find the latest workflow run** (should be triggered by our recent commits)
4. **Click on the workflow run** to see the build details
5. **Download the firmware artifacts**:
   - Look for "Artifacts" section on the right side
   - Download the firmware archive (usually named "firmware" or similar)
   - Extract the downloaded ZIP file

## Step 2: Prepare Your Keyboard

### For Split Keyboards (like yours):
- You'll need to flash **both halves** separately
- The firmware archive should contain two `.uf2` files:
  - `skreecustom_left-seeeduino_xiao_ble-zmk.uf2` (left half)
  - `skreecustom_right-seeeduino_xiao_ble-zmk.uf2` (right half)

## Step 3: Flash the Left Half

1. **Connect the left half** of your keyboard to your computer via USB
2. **Put the keyboard in bootloader mode**:
   - Double-click the reset button on the XIAO BLE board
   - OR hold the reset button while plugging in USB
   - The device should appear as a USB storage device (like "XIAO-SENSE" or similar)
3. **Copy the left firmware file**:
   - Copy `skreecustom_left-seeeduino_xiao_ble-zmk.uf2` to the root of the USB storage device
   - The device will automatically restart and load the new firmware

## Step 4: Flash the Right Half

1. **Connect the right half** of your keyboard to your computer via USB
2. **Put the keyboard in bootloader mode** (same process as left half)
3. **Copy the right firmware file**:
   - Copy `skreecustom_right-seeeduino_xiao_ble-zmk.uf2` to the root of the USB storage device
   - The device will automatically restart and load the new firmware

## Step 5: Pair the Halves

1. **Reset both halves simultaneously**:
   - Press the reset button on both halves at the same time
   - OR unplug and replug both halves
2. **The halves should automatically pair** via Bluetooth
3. **Test the keyboard** to ensure both halves are working

## Troubleshooting

### If GitHub Actions hasn't built yet:
- Wait a few minutes for the build to complete
- Check the Actions tab for build status
- If build failed, check the error logs

### If bootloader mode doesn't work:
- Try different USB cables
- Try different USB ports
- Hold reset button longer (3-5 seconds)
- Check if the XIAO BLE board has a different reset procedure

### If flashing fails:
- Ensure the `.uf2` file is copied to the root directory
- Don't rename the `.uf2` file
- Try a different USB port or cable
- Check that the file transfer completed successfully

### If halves don't pair:
- Reset both halves simultaneously
- Check Bluetooth settings on your computer
- Ensure both halves have power (battery or USB)

## Next Steps After Installation

1. **Test all keys** to ensure they work correctly
2. **Test the encoder** (volume up/down)
3. **Test the trackball** (mouse movement and scrolling)
4. **Test Bluetooth connectivity** (if using wireless)
5. **Customize your keymap** using the web editor if needed

## Current Keymap Features

- **Default Layer**: Standard QWERTY layout
- **Layer 1**: Bluetooth management (hold MO(1) key)
- **Layer 2**: Additional functions (hold MO(2) key)
- **Encoder**: Volume control (left side)
- **Trackball**: Mouse cursor and scroll wheel
- **Bluetooth**: Multi-device support

Your keyboard should now be ready to use with the ZMK firmware!
