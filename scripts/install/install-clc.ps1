$repo = "estifanos-neway/CLC"
$arch = $env:PROCESSOR_ARCHITECTURE
$version = (Invoke-RestMethod "https://api.github.com/repos/$repo/releases/latest").tag_name
$os = "Windows"
$installDir = "$env:USERPROFILE\AppData\Local\Programs\CLC"
$url = ""

switch ($arch) {
    "AMD64" { $url = "https://github.com/$repo/releases/download/$version/CLC_Windows_x86_64.zip" }
    "ARM64" { $url = "https://github.com/$repo/releases/download/$version/CLC_Windows_arm64.zip" }
    "x86" { $url = "https://github.com/$repo/releases/download/$version/CLC_Windows_i386.zip" }
    default { Write-Host "Unsupported architecture: $arch"; exit 1 }
}

Write-Host "Downloading CLC for $arch..."

$tempZip = "$env:TEMP\CLC.zip"
Invoke-WebRequest -Uri $url -OutFile $tempZip

Expand-Archive -Path $tempZip -DestinationPath $installDir -Force
$exePath = Join-Path $installDir "CLC.exe"

if (-not (Test-Path $exePath)) {
    Write-Host "Error: CLC.exe not found in extracted archive."
    exit 1
}

# Add to PATH for current user
$envPath = [Environment]::GetEnvironmentVariable("PATH", "User")
if ($envPath -notlike "*$installDir*") {
    [Environment]::SetEnvironmentVariable("PATH", "$envPath;$installDir", "User")
    Write-Host "Added $installDir to PATH (you may need to restart your terminal)"
}

Write-Host "CLC installed successfully at $exePath"
Write-Host "You can now run: CLC"
