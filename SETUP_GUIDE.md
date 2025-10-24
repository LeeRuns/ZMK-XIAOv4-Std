# ZMK XIAOv4 Keyboard Setup Guide

## Quick Start - Using the Web-Based Keymap Editor

This guide will help you customize your ZMK keyboard firmware using the web-based keymap editor, which is the easiest method for beginners.

### Prerequisites
- Your keyboard firmware repository forked on GitHub
- GitHub Actions enabled on your fork
- Access to the keymap editor website

### Step 1: Access the Keymap Editor
1. Go to [https://nickcoutsos.github.io/keymap-editor/](https://nickcoutsos.github.io/keymap-editor/)
2. Click "Authorize with GitHub" to connect your GitHub account
3. Select your repository: `leeruns/ZMK-XIAOv4-Std`
4. Choose the branch: `NS-5x6+6-ENC-TB-(73)`

### Step 2: Customize Your Keymap
1. The editor will load your current keymap configuration
2. Click on any key to change its function
3. Use the layer tabs to modify different layers (default, layer 1, layer 2, etc.)
4. Configure encoder functions (volume up/down by default)
5. Adjust trackball settings if needed

### Step 3: Save and Build
1. Click "Save" in the keymap editor
2. This will automatically commit your changes to your GitHub repository
3. GitHub Actions will automatically start building the firmware
4. Go to your repository's "Actions" tab to monitor the build progress

### Step 4: Download Firmware
1. Once the build completes successfully, go to the Actions tab
2. Click on the latest workflow run
3. Download the firmware artifacts (`.uf2` files for left and right halves)
4. Flash the firmware to your keyboard

### Step 5: Flash Your Keyboard
1. Put your keyboard into bootloader mode (usually by holding a specific key combination)
2. Copy the `.uf2` file to the keyboard's storage device
3. Safely eject the storage device
4. Your keyboard will restart with the new firmware

## Keyboard Layout

This keyboard features:
- **Layout**: 5x6+6 keys (73 total keys)
- **Split Design**: Left and right halves
- **Encoder**: Volume control on left side
- **Trackball**: Integrated trackball for mouse control
- **Layers**: Multiple layers for different functions
- **Bluetooth**: Wireless connectivity support

## Key Features

### Default Layer
- Standard QWERTY layout
- Function keys (F1-F12) on top row
- Number row (1-0) on second row
- Standard letter layout
- Thumb cluster with space, backspace, delete, enter, and shift keys

### Layer 1 (Bluetooth Layer)
- Bluetooth management functions
- Device selection (0-3)
- Clear Bluetooth connections
- Bootloader access

### Layer 2 (Additional Functions)
- Additional function keys
- Bootloader access
- Customizable functions

### Encoder Functions
- **Default**: Volume up/down
- **Layer 1**: Bluetooth device selection
- **Layer 2**: Additional functions

### Trackball Features
- **Left Trackball**: Scroll wheel functionality
- **Right Trackball**: Mouse cursor movement
- Configurable sensitivity and orientation

## Troubleshooting

### Build Issues
- Check the Actions tab for error messages
- Ensure all required files are present
- Verify the branch name matches your keyboard configuration

### Flashing Issues
- Make sure your keyboard is in bootloader mode
- Try a different USB cable or port
- Check that the firmware file is not corrupted

### Keymap Issues
- Verify your changes in the keymap editor before saving
- Test with a simple keymap first
- Check layer assignments and key bindings

## Advanced Customization

For advanced users who want to edit keymap files directly:
- Edit `config/skreecustom.keymap` for key assignments
- Modify `config/skreecustom.config` for configuration options
- Update `build.yaml` for build matrix changes

## Support

- Check the [ZMK Documentation](https://zmk.dev/docs) for detailed information
- Visit the [ZMK Discord](https://zmk.dev/community/discord/invite) for community support
- Review the original repository for additional configurations and examples
