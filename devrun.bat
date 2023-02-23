:: Build Vue Files
cd questionclient
call npm run build
cd ..

:: Run Server
$Env:CGO_ENABLED=1
call go run .
