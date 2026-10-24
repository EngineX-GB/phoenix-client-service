# To run this script:
# powershell.exe -ExecutionPolicy Bypass -File "release.ps1"
# for the .NET client (phoenix-client-ui)

# Note: For usage:
# 1. Ensure the github cli directory exists
# 2. Ensure git directory exists
# 3. Ensure dotnet directory exists
# 4. Update the script to put the correct GH Token.

$GOLANG_HOME = "C:\Program Files\Go"
$GIT_HOME = "C:\Program Files\Git"
$GITHUB_CLI_HOME = "C:\Program Files\GitHub CLI"
$GITHUB_CLI_BIN = "$GITHUB_CLI_HOME\gh.exe"

$GIT_BIN = "$GIT_HOME\bin\git.exe"
$PROJECT_NAME = "phoenix-client-service"
$PROJECT_HOME_DIRECTORY = "C:\Users\Dell\GolandProjects\$PROJECT_NAME"
$RELEASE_FOLDER="$PROJECT_HOME_DIRECTORY/release"
$RELEASE_BINARY_PATH="$RELEASE_FOLDER\$PROJECT_NAME.exe"
$TEMP_FOLDER="$PROJECT_HOME_DIRECTORY/temp"
$GH_TOKEN = ""
$env:GH_TOKEN=$GH_TOKEN

if ($GH_TOKEN -eq "") {
    Write-Host("[ERROR] Github token is required to run this release pipeline.")
    exit 1
}

if (-not(Test-Path $GOLANG_HOME -PathType Container)) {
    Write-Host "[ERROR] $DOTNET_HOME does not exist"
    return
}

if (-not(Test-Path $GIT_HOME -PathType Container)) {
    Write-Host "[ERROR] $GIT_HOME does not exist"
    return
}

if (-not(Test-Path $GITHUB_CLI_HOME -PathType Container)) {
    Write-Host "[ERROR] $GITHUB_CLI_HOME does not exist"
    return
}

cd $PROJECT_HOME_DIRECTORY
& $GOLANG_HOME\bin\go.exe build -o $RELEASE_BINARY_PATH

IF ($LASTEXITCODE -ne 0) {
    Write-Host "[ERROR] Build failed with exit code $LASTEXITCODE"
    exit $LASTEXITCODE
}


# Here, write read the version manifest file in the client directory.
# Then read the version number

$VERSIONS_FILE = Join-Path $PROJECT_HOME_DIRECTORY "version.json"
$VERSIONS = Get-Content $VERSIONS_FILE -Raw | ConvertFrom-Json
$SERVICE_VERSION = $VERSIONS."version"

Write-Host $SERVICE_VERSION


# The increment the minor version

# Split the version into parts
$PARTS = $SERVICE_VERSION.Split('.')

# Increment the last digit
$PARTS[2] = [int]$PARTS[2] + 1

# Reassemble the version
$NEW_VERSION = $PARTS -join '.'

# Update the JSON property
$VERSIONS."version" = $NEW_VERSION

# Save the JSON back to the file
$VERSIONS | ConvertTo-Json -Depth 10 | Set-Content $VERSIONS_FILE

Write-Host "[SUCCESS] Version updated from $VERSIONS to $NEW_VERSION"


# Take the new version and tag the git master branch with this new version

Write-Host "[INFO] Creating tag on branch"
& $GIT_BIN tag -a $NEW_VERSION -m "Version $NEW_VERSION"
& $GIT_BIN push origin $NEW_VERSION

# Here, we need to commit the file version.json that has the up to date change log

git add version.json
Write-Host "[INFO] Committing updated version.json to repository"
& $GIT_BIN commit -m "Updating version.json change log with the new version ($NEW_VERSION)."
& $GIT_BIN push origin master




# now do an optimised release binary (exe) of the service.
& $GOLANG_HOME\bin\go.exe build

IF ($LASTEXITCODE -ne 0) {
    Write-Host "[ERROR] Optimised Release Build failed with exit code $LASTEXITCODE"
    exit $LASTEXITCODE
}

# create the temp folder
New-Item -Path $TEMP_FOLDER -ItemType Directory -Force | Out-Null


# The zip the archive (below)

IF (Test-Path $RELEASE_FOLDER -PathType Container) {
    # if the release folder exists, then zip it up
    Compress-Archive -Path $RELEASE_FOLDER -DestinationPath $TEMP_FOLDER/$PROJECT_NAME -CompressionLevel Optimal -Force
} else {
    Write Host "[ERROR] $RELEASE_FOLDER does not exist. Exiting."
    exit 1
}

# Publish as a release against the version using the Github CLI token

$RELEASE_TITLE="phoenix-client-service Release Title (test) ($NEW_VERSION)"
$RELEASE_NOTES="phoenix-client-service Release Notes (test) ($NEW_VERSION)"

Write-Host "[INFO] Publishing artefact..."
& $GITHUB_CLI_BIN release create $NEW_VERSION --notes $RELEASE_NOTES -t $RELEASE_TITLE
Write-Host ""
Write-Host "[INFO] Removing Temp Files"
Remove-Item "$RELEASE_FOLDER" -Recurse -Force
Remove-Item "$TEMP_FOLDER" -Recurse -Force
Write-Host ""
Write-Host "[INFO] Release Completed."