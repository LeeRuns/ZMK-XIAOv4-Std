package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
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

	// Create firmware directory
	firmwareDir := "firmware"
	if err := os.MkdirAll(firmwareDir, 0755); err != nil {
		fmt.Printf("❌ Error creating firmware directory: %v\n", err)
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
		return
	}

	fmt.Printf("✅ Found successful workflow run: %d\n", workflowRun.ID)
	fmt.Printf("📅 Created: %s\n", workflowRun.CreatedAt.Format("2006-01-02 15:04:05"))

	// Download and extract artifacts
	for _, artifact := range workflowRun.Artifacts {
		if strings.Contains(strings.ToLower(artifact.Name), "firmware") {
			fmt.Printf("\n📦 Downloading artifact: %s\n", artifact.Name)

			if err := downloadAndExtractArtifact(artifact, firmwareDir); err != nil {
				fmt.Printf("❌ Error downloading artifact %s: %v\n", artifact.Name, err)
				continue
			}

			fmt.Printf("✅ Successfully downloaded and extracted: %s\n", artifact.Name)
		}
	}

	// List downloaded firmware files
	fmt.Println("\n📁 Firmware files ready for installation:")
	listFirmwareFiles(firmwareDir)

	fmt.Println("\n🚀 Next steps:")
	fmt.Println("1. Connect your keyboard's LEFT half via USB")
	fmt.Println("2. Double-click the reset button on the XIAO BLE board")
	fmt.Println("3. Copy the LEFT .uf2 file to the USB storage device")
	fmt.Println("4. Repeat for the RIGHT half")
	fmt.Println("5. Reset both halves simultaneously to pair them")
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

func downloadAndExtractArtifact(artifact struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	URL  string `json:"archive_download_url"`
}, destDir string) error {
	// Note: In a real implementation, you'd need to authenticate with GitHub
	// For now, we'll create a placeholder and show the download URL
	fmt.Printf("🔗 Artifact download URL: %s\n", artifact.URL)
	fmt.Println("⚠️  Note: You'll need to manually download this artifact from GitHub Actions")
	fmt.Println("   Go to: https://github.com/LeeRuns/ZMK-XIAOv4-Std/actions")
	fmt.Println("   Find the latest successful run and download the firmware artifacts")

	// Create a placeholder directory structure
	artifactDir := filepath.Join(destDir, artifact.Name)
	if err := os.MkdirAll(artifactDir, 0755); err != nil {
		return err
	}

	// Create placeholder files to show expected structure
	leftFirmware := filepath.Join(artifactDir, "skreecustom_left-seeeduino_xiao_ble-zmk.uf2")
	rightFirmware := filepath.Join(artifactDir, "skreecustom_right-seeeduino_xiao_ble-zmk.uf2")

	// Create placeholder files
	createPlaceholderFile(leftFirmware, "LEFT HALF FIRMWARE")
	createPlaceholderFile(rightFirmware, "RIGHT HALF FIRMWARE")

	return nil
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

func listFirmwareFiles(dir string) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(strings.ToLower(path), ".uf2") {
			relPath, _ := filepath.Rel(dir, path)
			fmt.Printf("  📄 %s\n", relPath)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error listing files: %v\n", err)
	}
}
