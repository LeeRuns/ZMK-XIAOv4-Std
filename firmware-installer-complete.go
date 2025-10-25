package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const (
	githubAPIBase = "https://api.github.com"
	repoOwner     = "LeeRuns"
	repoName      = "ZMK-XIAOv4-Std"
	branchName    = "NS-5x6+6-ENC-TB-(73)"
)

type WorkflowRun struct {
	ID         int64     `json:"id"`
	Status     string    `json:"status"`
	Conclusion string    `json:"conclusion"`
	CreatedAt  time.Time `json:"created_at"`
	Artifacts  []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		URL  string `json:"archive_download_url"`
	} `json:"artifacts"`
}

type WorkflowRunsResponse struct {
	WorkflowRuns []WorkflowRun `json:"workflow_runs"`
}

func main() {
	fmt.Println("🔧 ZMK Firmware Installer")
	fmt.Println("=========================")
	fmt.Printf("Repository: %s/%s\n", repoOwner, repoName)
	fmt.Printf("Branch: %s\n\n", branchName)

	// Create firmware directory structure
	firmwareDir := "firmware"
	if err := setupFirmwareDirectory(firmwareDir); err != nil {
		fmt.Printf("❌ Error setting up firmware directory: %v\n", err)
		return
	}

	// Get latest successful workflow run
	workflowRun, err := getLatestSuccessfulRun()
	if err != nil {
		fmt.Printf("❌ Error getting workflow run: %v\n", err)
		return
	}

	if workflowRun == nil {
		fmt.Println("❌ No successful workflow runs found")
		fmt.Println("💡 You can still use the manual installation process")
		showManualInstructions(firmwareDir)
		return
	}

	fmt.Printf("✅ Found successful workflow run: %d\n", workflowRun.ID)
	fmt.Printf("📅 Created: %s\n", workflowRun.CreatedAt.Format("2006-01-02 15:04:05"))

	// Process artifacts
	processArtifacts(workflowRun.Artifacts, firmwareDir)

	// Show installation options
	showInstallationOptions(firmwareDir)
}

func setupFirmwareDirectory(firmwareDir string) error {
	// Create main firmware directory
	if err := os.MkdirAll(firmwareDir, 0755); err != nil {
		return err
	}

	// Create organized subdirectories
	dirs := []string{
		filepath.Join(firmwareDir, "left"),
		filepath.Join(firmwareDir, "right"),
		filepath.Join(firmwareDir, "archives"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	fmt.Printf("✅ Created firmware directory: %s\n", firmwareDir)
	fmt.Println("📁 Created organized directories:")
	fmt.Printf("  📂 %s (for left half firmware)\n", filepath.Join(firmwareDir, "left"))
	fmt.Printf("  📂 %s (for right half firmware)\n", filepath.Join(firmwareDir, "right"))
	fmt.Printf("  📂 %s (for downloaded archives)\n", filepath.Join(firmwareDir, "archives"))

	// Create installation guide
	return createInstallationGuide(firmwareDir)
}

func processArtifacts(artifacts []struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	URL  string `json:"archive_download_url"`
}, firmwareDir string) {
	firmwareFound := false

	for _, artifact := range artifacts {
		if strings.Contains(strings.ToLower(artifact.Name), "firmware") {
			firmwareFound = true
			fmt.Printf("\n📦 Found firmware artifact: %s\n", artifact.Name)
			fmt.Printf("🔗 Download URL: %s\n", artifact.URL)

			// Create placeholder files
			createPlaceholderFiles(firmwareDir, artifact.Name)
		}
	}

	if !firmwareFound {
		fmt.Println("\n⚠️  No firmware artifacts found in the latest run")
		fmt.Println("💡 You may need to wait for the build to complete or check manually")
	}
}

func createPlaceholderFiles(firmwareDir, artifactName string) {
	artifactDir := filepath.Join(firmwareDir, "archives", artifactName)
	os.MkdirAll(artifactDir, 0755)

	// Create placeholder files
	leftFirmware := filepath.Join(firmwareDir, "left", "skreecustom_left-seeeduino_xiao_ble-zmk.uf2")
	rightFirmware := filepath.Join(firmwareDir, "right", "skreecustom_right-seeeduino_xiao_ble-zmk.uf2")

	createPlaceholderFile(leftFirmware, "LEFT HALF FIRMWARE")
	createPlaceholderFile(rightFirmware, "RIGHT HALF FIRMWARE")

	fmt.Printf("📄 Created placeholder files in organized structure\n")
}

func createPlaceholderFile(filename, content string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Warning: Could not create placeholder %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	fmt.Fprintf(file, "# %s\n# Replace this file with the actual .uf2 firmware file\n# Downloaded from GitHub Actions\n", content)
}

func showInstallationOptions(firmwareDir string) {
	fmt.Println("\n🚀 Installation Options:")
	fmt.Println("1. 📥 Download firmware from GitHub Actions")
	fmt.Println("2. 🔧 Start installation process")
	fmt.Println("3. 📖 View detailed installation guide")
	fmt.Println("4. 🗂️  Open firmware folders")
	fmt.Println("5. ⚡ Quick install (like batch file)")
	fmt.Println("6. ❌ Exit")

	var choice string
	fmt.Print("\nSelect an option (1-6): ")
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		openGitHubActions()
	case "2":
		startInstallationProcess(firmwareDir)
	case "3":
		openInstallationGuide(firmwareDir)
	case "4":
		openFirmwareFolders(firmwareDir)
	case "5":
		quickInstallProcess(firmwareDir)
	case "6":
		fmt.Println("👋 Goodbye!")
	default:
		fmt.Println("❌ Invalid option. Please try again.")
		showInstallationOptions(firmwareDir)
	}
}

func openGitHubActions() {
	fmt.Println("\n🌐 Opening GitHub Actions...")
	url := fmt.Sprintf("https://github.com/%s/%s/actions", repoOwner, repoName)
	openURL(url)
	fmt.Println("📋 Instructions:")
	fmt.Println("1. Find the latest successful workflow run")
	fmt.Println("2. Click on the workflow run")
	fmt.Println("3. Download the firmware artifacts (ZIP file)")
	fmt.Println("4. Extract the .uf2 files to the appropriate folders")
}

func startInstallationProcess(firmwareDir string) {
	fmt.Println("\n🔧 Starting Installation Process...")

	leftFirmware := filepath.Join(firmwareDir, "left", "skreecustom_left-seeeduino_xiao_ble-zmk.uf2")
	rightFirmware := filepath.Join(firmwareDir, "right", "skreecustom_right-seeeduino_xiao_ble-zmk.uf2")

	// Check if firmware files exist
	leftExists := checkFirmwareFile(leftFirmware)
	rightExists := checkFirmwareFile(rightFirmware)

	if !leftExists || !rightExists {
		fmt.Println("❌ Firmware files not found!")
		fmt.Println("💡 Please download the firmware from GitHub Actions first")
		fmt.Println("   Run option 1 to open GitHub Actions")
		
		// Ask if user wants to continue anyway
		fmt.Print("\nDo you want to continue and open folders anyway? (y/n): ")
		var choice string
		fmt.Scanln(&choice)
		if strings.ToLower(choice) != "y" && strings.ToLower(choice) != "yes" {
			return
		}
	} else {
		fmt.Println("✅ Firmware files found!")
	}

	fmt.Println("\n🚀 Ready to install firmware")
	fmt.Println("\n📋 Installation Steps:")
	fmt.Println("1. Connect LEFT half of keyboard via USB")
	fmt.Println("2. Double-click reset button on XIAO BLE board")
	fmt.Println("3. Copy left firmware to USB storage device")
	fmt.Println("4. Repeat for RIGHT half")
	fmt.Println("5. Reset both halves to pair")

	fmt.Print("\nPress Enter to open firmware folders and installation guide...")
	fmt.Scanln()

	// Open firmware folders and installation guide
	openFirmwareFolders(firmwareDir)
	openInstallationGuide(firmwareDir)
	
	fmt.Println("\n✅ Installation process started!")
	fmt.Println("Check the opened folders and guide for next steps.")
}

func checkFirmwareFile(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}

	// Check if it's a placeholder file (small size)
	if info.Size() < 1000 {
		fmt.Printf("⚠️  %s appears to be a placeholder file\n", filepath.Base(filename))
		return false
	}

	return true
}

func openInstallationGuide(firmwareDir string) {
	guidePath := filepath.Join(firmwareDir, "INSTALLATION_GUIDE.md")
	fmt.Printf("\n📖 Opening installation guide: %s\n", guidePath)
	openFile(guidePath)
}

func openFirmwareFolders(firmwareDir string) {
	fmt.Println("\n📂 Opening firmware folders...")

	leftDir := filepath.Join(firmwareDir, "left")
	rightDir := filepath.Join(firmwareDir, "right")

	openFolder(leftDir)
	openFolder(rightDir)

	fmt.Println("✅ Folders opened! Copy the appropriate .uf2 files to your keyboard")
}

func openURL(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		fmt.Printf("Please open this URL manually: %s\n", url)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Printf("Could not open URL: %v\n", err)
		fmt.Printf("Please open manually: %s\n", url)
	}
}

func openFile(filename string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "", filename)
	case "darwin":
		cmd = exec.Command("open", filename)
	case "linux":
		cmd = exec.Command("xdg-open", filename)
	default:
		fmt.Printf("Please open this file manually: %s\n", filename)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Printf("Could not open file: %v\n", err)
		fmt.Printf("Please open manually: %s\n", filename)
	}
}

func openFolder(dir string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "", dir)
	case "darwin":
		cmd = exec.Command("open", dir)
	case "linux":
		cmd = exec.Command("xdg-open", dir)
	default:
		fmt.Printf("Please open this folder manually: %s\n", dir)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Printf("Could not open folder: %v\n", err)
		fmt.Printf("Please open manually: %s\n", dir)
	}
}

func quickInstallProcess(firmwareDir string) {
	fmt.Println("\n⚡ Quick Install Process (Batch File Mode)")
	fmt.Println("==========================================")

	leftFirmware := filepath.Join(firmwareDir, "left", "skreecustom_left-seeeduino_xiao_ble-zmk.uf2")
	rightFirmware := filepath.Join(firmwareDir, "right", "skreecustom_right-seeeduino_xiao_ble-zmk.uf2")

	fmt.Println("📁 Checking firmware files...")
	
	leftExists := checkFirmwareFile(leftFirmware)
	rightExists := checkFirmwareFile(rightFirmware)

	if !leftExists {
		fmt.Printf("❌ Left firmware not found: %s\n", leftFirmware)
		fmt.Println("   Please download and extract firmware to the correct folders")
		fmt.Print("\nPress Enter to continue...")
		fmt.Scanln()
		return
	}

	if !rightExists {
		fmt.Printf("❌ Right firmware not found: %s\n", rightFirmware)
		fmt.Println("   Please download and extract firmware to the correct folders")
		fmt.Print("\nPress Enter to continue...")
		fmt.Scanln()
		return
	}

	fmt.Println("✅ Firmware files found!")
	fmt.Println()
	fmt.Println("🚀 Ready to install firmware")
	fmt.Println()
	fmt.Println("Instructions:")
	fmt.Println("1. Connect LEFT half of keyboard via USB")
	fmt.Println("2. Double-click reset button on XIAO BLE board")
	fmt.Println("3. Copy left firmware to USB storage device")
	fmt.Println("4. Repeat for RIGHT half")
	fmt.Println("5. Reset both halves to pair")
	fmt.Println()
	fmt.Print("Press Enter to open firmware folders...")
	fmt.Scanln()

	fmt.Println("📂 Opening firmware folders...")
	openFirmwareFolders(firmwareDir)

	fmt.Println()
	fmt.Println("📖 Opening installation guide...")
	openInstallationGuide(firmwareDir)

	fmt.Println()
	fmt.Println("✅ Installation script completed!")
	fmt.Println("Check the opened folders and guide for next steps.")
	fmt.Print("Press Enter to exit...")
	fmt.Scanln()
}

func showManualInstructions(firmwareDir string) {
	fmt.Println("\n📋 Manual Installation Instructions:")
	fmt.Println("1. Go to: https://github.com/LeeRuns/ZMK-XIAOv4-Std/actions")
	fmt.Println("2. Find the latest successful workflow run")
	fmt.Println("3. Download the firmware artifacts")
	fmt.Println("4. Extract the .uf2 files to the appropriate folders")
	fmt.Printf("5. Run this installer again or use the folders in: %s\n", firmwareDir)
}

func createInstallationGuide(firmwareDir string) error {
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

## Current Keymap Features
- **Default Layer**: Standard QWERTY layout
- **Layer 1**: Bluetooth management (hold MO(1) key)
- **Layer 2**: Additional functions (hold MO(2) key)
- **Encoder**: Volume control (left side)
- **Trackball**: Mouse cursor and scroll wheel
- **Bluetooth**: Multi-device support

Generated on: ` + time.Now().Format("2006-01-02 15:04:05") + `
`

	return os.WriteFile(guidePath, []byte(content), 0644)
}

func getLatestSuccessfulRun() (*WorkflowRun, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/actions/runs?branch=%s&status=completed&per_page=10",
		githubAPIBase, repoOwner, repoName, branchName)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned status %d", resp.StatusCode)
	}

	var response WorkflowRunsResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	// Find the first successful run
	for _, run := range response.WorkflowRuns {
		if run.Conclusion == "success" {
			// Get detailed run info with artifacts
			detailedRun, err := getWorkflowRunDetails(run.ID)
			if err != nil {
				continue
			}
			return detailedRun, nil
		}
	}

	return nil, nil
}

func getWorkflowRunDetails(runID int64) (*WorkflowRun, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/actions/runs/%d",
		githubAPIBase, repoOwner, repoName, runID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned status %d", resp.StatusCode)
	}

	var run WorkflowRun
	if err := json.NewDecoder(resp.Body).Decode(&run); err != nil {
		return nil, err
	}

	return &run, nil
}
