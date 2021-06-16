# Initial greeting
write-host "Good Day!"

# Capture the user's name.  I used this method as the typical read-host appends a colon to the end the query
Write-Host "What is your name? " -NoNewline
$name = $Host.UI.ReadLine()

# Print our hello message with the user's name
Write-Host "Hello, "$name"!"

# Sleep for five seconds
Start-Sleep -s 5