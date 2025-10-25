package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	fmt.Println("🔧 ZMK Firmware Installer")
	fmt.Println("=========================")
	fmt.Println("Repository: LeeRuns/ZMK-XIAOv4-Std")
	fmt.Println("Branch: NS-5x6+6-ENC-TB-(73)")
	fmt.Println()

	// Create firmware directory
	firmwareDir := "firmware"
	if err := os.MkdirAll(firmwareDir, 0755); err != nil {
		fmt.Printf("❌ Error creating firmware directory: %v\n", err)
		return
	}

	fmt.Printf("✅ Created firmware directory: %s\n", firmwareDir)

	// Create organized subdirectories
	leftDir := filepath.Join(firmwareDir, "left")
	rightDir := filepath.Join(firmwareDir, "right")
	archiveDir := filepath.Join(firmwareDir, "archives")

	for _, dir := range []string{leftDir, rightDir, archiveDir} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Printf("❌ Error creating directory %s: %v\n", dir, err)
			return
		}
	}

	fmt.Println("📁 Created organized directories:")
	fmt.Printf("  📂 %s (for left half firmware)\n", leftDir)
	fmt.Printf("  📂 %s (for right half firmware)\n", rightDir)
	fmt.Printf("  📂 %s (for downloaded archives)\n", archiveDir)

	// Create installation instructions
	createInstallationGuide(firmwareDir)

	// Create batch script for Windows
	createBatchScript(firmwareDir)

	fmt.Println("\n🚀 Installation Guide Created!")
	fmt.Println("📖 Check 'firmware/INSTALLATION_GUIDE.md' for detailed instructions")
	fmt.Println("🔧 Run 'firmware/install-firmware.bat' for automated installation")
}

func createInstallationGuide(firmwareDir string) {
	guidePath := filepath.Join(firmwareDir, "INSTALLATION_GUIDE.md")

	content := `# ZMK Firmware Installation Guide

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

## File Structure
firmware/
- left/                           # Left half firmware
  - skreecustom_left-seeeduino_xiao_ble-zmk.uf2
- right/                          # Right half firmware
  - skreecustom_right-seeeduino_xiao_ble-zmk.uf2
- archives/                       # Downloaded ZIP files
- INSTALLATION_GUIDE.md          # This guide
- install-firmware.bat           # Windows batch script

## Current Keymap Features
- **Default Layer**: Standard QWERTY layout
- **Layer 1**: Bluetooth management (hold MO(1) key)
- **Layer 2**: Additional functions (hold MO(2) key)
- **Encoder**: Volume control (left side)
- **Trackball**: Mouse cursor and scroll wheel
- **Bluetooth**: Multi-device support

Generated on: ` + time.Now().Format("2006-01-02 15:04:05") + `
`

	if err := os.WriteFile(guidePath, []byte(content), 0644); err != nil {
		fmt.Printf("❌ Error creating installation guide: %v\n", err)
		return
	}

	fmt.Printf("📖 Created installation guide: %s\n", guidePath)
}

func createBatchScript(firmwareDir string) {
	scriptPath := filepath.Join(firmwareDir, "install-firmware.bat")

	content := `@echo off
echo 🔧 ZMK Firmware Installation Script
echo ===================================
echo.

set "FIRMWARE_DIR=%~dp0"
set "LEFT_FIRMWARE=%FIRMWARE_DIR%left\skreecustom_left-seeeduino_xiao_ble-zmk.uf2"
set "RIGHT_FIRMWARE=%FIRMWARE_DIR%right\skreecustom_right-seeeduino_xiao_ble-zmk.uf2"

echo 📁 Checking firmware files...
if not exist "%LEFT_FIRMWARE%" (
    echo ❌ Left firmware not found: %LEFT_FIRMWARE%
    echo    Please download and extract firmware to the correct folders
    pause
    exit /b 1
)

if not exist "%RIGHT_FIRMWARE%" (
    echo ❌ Right firmware not found: %RIGHT_FIRMWARE%
    echo    Please download and extract firmware to the correct folders
    pause
    exit /b 1
)

echo ✅ Firmware files found!
echo.
echo 🚀 Ready to install firmware
echo.
echo Instructions:
echo 1. Connect LEFT half of keyboard via USB
echo 2. Double-click reset button on XIAO BLE board
echo 3. Copy left firmware to USB storage device
echo 4. Repeat for RIGHT half
echo 5. Reset both halves to pair
echo.
echo Press any key to open firmware folders...
pause >nul

echo 📂 Opening firmware folders...
start "" "%FIRMWARE_DIR%left"
start "" "%FIRMWARE_DIR%right"

echo.
echo 📖 Opening installation guide...
start "" "%FIRMWARE_DIR%INSTALLATION_GUIDE.md"

echo.
echo ✅ Installation script completed!
echo Check the opened folders and guide for next steps.
pause
`

	if err := os.WriteFile(scriptPath, []byte(content), 0644); err != nil {
		fmt.Printf("❌ Error creating batch script: %v\n", err)
		return
	}

	fmt.Printf("🔧 Created batch script: %s\n", scriptPath)
}
