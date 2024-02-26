param(
    [Parameter(Mandatory = $true)]
    [string]$action
)

$containerName = "urlshort"
$imageName = "urlshort"
$portMapping = "8000:3000"

function Start-DockerContainer {
    # Check if the image exists
    $imageExists = docker images -q $imageName 2>$null
    if (-not $imageExists) {
        Write-Host "Building Docker image $imageName..."
        docker build -t $imageName .
    }
    else {
        Write-Host "Docker image $imageName already exists."
    }

    # Check if the container is already running
    $containerExists = docker ps -a -q -f name=$containerName 2>$null
    if ($containerExists) {
        $containerRunning = docker ps -q -f name=$containerName
        if ($containerRunning) {
            Write-Host "Container $containerName is already running."
        }
        else {
            Write-Host "Starting existing container $containerName..."
            docker start $containerName
        }
    }
    else {
        Write-Host "Creating and starting new container $containerName..."
        docker run -d --name $containerName -p $portMapping $imageName
    }
}

function Stop-DockerContainer {
    # Check if the container exists (regardless of its state)
    $containerExists = docker ps -a -q -f name=$containerName 2>$null
    if ($containerExists) {
        # Stop the container if it is running
        $containerRunning = docker ps -q -f name=$containerName
        if ($containerRunning) {
            Write-Host "Stopping container $containerName..."
            docker stop $containerName
        }
        else {
            Write-Host "Container $containerName is not running."
        }

        # Remove the container
        Write-Host "Removing container $containerName..."
        docker rm $containerName
    }
    else {
        Write-Host "Container $containerName does not exist."
    }

    # Attempt to remove the image if specified
    $imageExists = docker images -q $imageName 2>$null
    if ($imageExists) {
        Write-Host "Removing image $imageName..."
        docker rmi $imageName -f
    }
    else {
        Write-Host "Image $imageName does not exist."
    }
}

switch ($action.ToLower()) {
    "start" {
        Start-DockerContainer
    }
    "stop" {
        Stop-DockerContainer
    }
    default {
        Write-Host "Invalid action: $action. Please use 'start' or 'stop'."
    }
}
