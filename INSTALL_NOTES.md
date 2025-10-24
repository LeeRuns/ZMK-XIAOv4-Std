# ZMK XIAOv4 Keyboard Installation Notes

## Installation Process Summary

**Date**: October 24, 2025  
**Repository**: https://github.com/leeruns/ZMK-XIAOv4-Std  
**Branch**: NS-5x6+6-ENC-TB-(73)  
**Original Repository**: https://github.com/WainingForests/ZMK-XIAOv4-Std  

## What Was Done

### 1. Repository Setup
- **Cloned Repository**: Successfully cloned the forked repository to `C:\Users\LeeRunyon\source\repos\ZMK-XIAOv4-Std`
- **Branch Selection**: Switched to the `NS-5x6+6-ENC-TB-(73)` branch which contains the specific keyboard configuration
- **Remote Configuration**: Added upstream remote to track the original repository
- **Branch Push**: Pushed the working branch to the fork to enable GitHub Actions

### 2. Repository Structure Analysis

#### Key Directories and Files:
```
ZMK-XIAOv4-Std/
├── .github/
│   └── workflows/
│       └── blank.yml          # GitHub Actions workflow
├── boards/
│   └── shields/
│       └── skreecustom/       # Keyboard shield definition
├── config/
│   ├── skreecustom.config     # Configuration options
│   ├── skreecustom.keymap    # Keymap definitions
│   ├── skreecustom.json      # Layout definition for keymap editor
│   └── west.yml              # Zephyr workspace configuration
├── build.yaml                 # Build matrix configuration
└── zephyr/                    # Zephyr RTOS framework
```

#### GitHub Actions Workflow:
- **File**: `.github/workflows/blank.yml`
- **Trigger**: Push, pull request, or manual dispatch
- **Action**: Uses `zmkfirmware/zmk/.github/workflows/build-user-config.yml@main`
- **Purpose**: Automatically builds firmware when changes are pushed

#### Build Configuration:
- **File**: `build.yaml`
- **Boards**: `seeeduino_xiao_ble` (XIAO BLE microcontroller)
- **Shields**: `skreecustom_left` and `skreecustom_right`
- **Output**: Separate firmware files for left and right keyboard halves

### 3. Keyboard Configuration Details

#### Physical Layout:
- **Matrix**: 6 rows × 12 columns (72 keys + 1 thumb key = 73 total)
- **Layout**: 5x6+6 configuration with split design
- **Thumb Cluster**: 6 additional keys for thumbs
- **Split**: Left and right halves communicate wirelessly

#### Hardware Features:
- **Microcontroller**: Seeeduino XIAO BLE
- **Encoder**: Alps EC11 rotary encoder on left side (volume control)
- **Trackball**: PMW3610 optical sensor
  - Left trackball: Scroll wheel functionality
  - Right trackball: Mouse cursor movement
- **Bluetooth**: Built-in BLE connectivity
- **Battery**: Integrated battery support

#### Keymap Structure:
- **Default Layer**: Standard QWERTY layout with function keys
- **Layer 1**: Bluetooth management and device selection
- **Layer 2**: Additional functions and bootloader access
- **Layer 3**: Reserved for trackball activation

#### GPIO Configuration:
- **Row Pins**: GPIO0.3, GPIO0.28, GPIO1.1, GPIO0.9, GPIO0.10, GPIO1.11, GPIO1.15
- **Column Pins**: GPIO1.7, GPIO1.14, GPIO1.5, GPIO1.13, GPIO1.3, GPIO1.12
- **Encoder Pins**: GPIO0.16 (A), GPIO1.10 (B)
- **Trackball**: SPI1 interface with CS on GPIO0.2, IRQ on GPIO0.29

### 4. Keymap Editor Integration

#### Web-Based Editor:
- **URL**: https://nickcoutsos.github.io/keymap-editor/
- **Integration**: Direct GitHub repository connection
- **Workflow**: Edit → Save → Auto-commit → GitHub Actions build
- **Layout File**: `config/skreecustom.json` defines the visual layout

#### Editor Features:
- Visual keymap editing
- Layer management
- Encoder configuration
- Trackball settings
- Real-time preview
- GitHub integration

### 5. Build Process

#### Automatic Building:
1. **Trigger**: Push to repository or manual workflow dispatch
2. **Environment**: GitHub Actions runner with ZMK build tools
3. **Process**: 
   - Initialize Zephyr workspace
   - Build left half firmware
   - Build right half firmware
   - Generate artifacts
4. **Output**: `.uf2` files for flashing

#### Manual Building (if needed):
```bash
west init -l app
west update
west zephyr-export
west build -p -b seeeduino_xiao_ble -- -DSHIELD=skreecustom_left
west build -p -b seeeduino_xiao_ble -- -DSHIELD=skreecustom_right
```

### 6. Flashing Process

#### Firmware Files:
- **Left Half**: `skreecustom_left-seeeduino_xiao_ble-zmk.uf2`
- **Right Half**: `skreecustom_right-seeeduino_xiao_ble-zmk.uf2`

#### Flashing Steps:
1. Put keyboard in bootloader mode
2. Copy `.uf2` file to keyboard storage
3. Safely eject storage
4. Keyboard reboots with new firmware

## Key Findings and Notes

### Repository Structure:
- Well-organized ZMK configuration
- Proper separation of hardware definitions and keymaps
- Clean GitHub Actions integration
- Multiple branch support for different keyboard variants

### Configuration Files:
- **skreecustom.keymap**: Contains all key bindings and layer definitions
- **skreecustom.dtsi**: Hardware definition including matrix, GPIO, and peripherals
- **skreecustom.json**: Layout definition for visual keymap editor
- **build.yaml**: Build matrix for GitHub Actions

### Hardware Capabilities:
- Full split keyboard with wireless communication
- Integrated trackball for mouse functionality
- Rotary encoder for volume control
- Multiple Bluetooth device support
- Battery-powered operation

### Software Features:
- ZMK firmware with modern features
- Multiple layers for different functions
- Bluetooth device management
- Trackball with configurable behavior
- Encoder with layer-specific functions

## Troubleshooting Notes

### Common Issues:
1. **Branch Mismatch**: Ensure using correct branch for your keyboard variant
2. **Build Failures**: Check GitHub Actions logs for specific errors
3. **Flashing Issues**: Verify bootloader mode and USB connection
4. **Keymap Problems**: Test with simple changes first

### Debugging Steps:
1. Check repository Actions tab for build status
2. Verify all required files are present
3. Test keymap changes incrementally
4. Use ZMK documentation for advanced configuration

## Next Steps

### Immediate Actions:
1. Test the keymap editor integration
2. Make a simple keymap change to verify the build process
3. Flash firmware to keyboard and test functionality

### Future Enhancements:
1. Customize keymap for personal preferences
2. Configure trackball sensitivity and behavior
3. Set up Bluetooth device profiles
4. Optimize battery usage settings

## Resources

- **ZMK Documentation**: https://zmk.dev/docs
- **Keymap Editor**: https://nickcoutsos.github.io/keymap-editor/
- **Original Repository**: https://github.com/WainingForests/ZMK-XIAOv4-Std
- **ZMK Community**: https://zmk.dev/community/discord/invite

## Conclusion

The ZMK XIAOv4 keyboard setup is now complete with:
- Repository cloned and configured
- GitHub Actions workflow enabled
- Keymap editor integration documented
- Build process automated
- Comprehensive documentation created

The setup provides a solid foundation for keyboard customization and firmware development using modern ZMK features and web-based tools.
