# Check if Go is already installed
if (Test-Path "$env:GOPATH\bin\go.exe") {
    Write-Host "Go is already installed."
    exit
}

# Download and install Go
Write-Host "Downloading and installing Go..."
$url = "https://golang.org/dl/"
$latestVersion = (Invoke-WebRequest -Uri $url).Content | Select-String -Pattern 'go[0-9]+\.[0-9]+(\.[0-9]+)?' | ForEach-Object { $_.Matches[0].Value } | Sort-Object -Unique | Select-Object -Last 1
$downloadUrl = "$url$latestVersion.windows-amd64.msi"
$installerPath = "$env:TEMP\go_installer.msi"

Invoke-WebRequest -Uri $downloadUrl -OutFile $installerPath
Start-Process -Wait -FilePath msiexec -ArgumentList "/i $installerPath /quiet"
Remove-Item $installerPath

# Add Go binary directory to PATH
$goPathBin = "$env:GOPATH\bin"
[Environment]::SetEnvironmentVariable('Path', "$env:Path;$goPathBin", [System.EnvironmentVariableTarget]::User)

# Print Go version
go version
