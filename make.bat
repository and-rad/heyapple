@echo off
setlocal EnableDelayedExpansion

set DEV_OS=windows
set DEV_ARCH=amd64
set DEV_EXT=.exe

call .make\env.bat

if "%1"=="test" (
    call :TestServer
)

if "%1"=="build" (
    call :TestServer
    call :BuildWeb
    call :BuildServer
    call :BuildCLI
)

if "%1"=="run" (
    call :RunServer
)

exit

:TestServer
    set NUM_TESTS=0
    for /r %%i in (*_test.go) do (
        for /f %%j in ('type "%%i"^|find /c "{ //"') do (
            set /a NUM_TESTS += %%j
        )
    )

    set TEST_PACKAGES=
    for /f %%i in ('go list ./internal/...^|findstr /V /C:"/defaults/" /C:"/mock" /C:"/web"') do (
        set TEST_PACKAGES=!TEST_PACKAGES! %%i
    )

    echo testing %NUM_TESTS% cases:
    go test -short -cover -p 1 -timeout 30m %TEST_PACKAGES%

    exit /B 0

:BuildWeb
    call npm run build --prefix ./internal/web/src/login
	call npm run build --prefix ./internal/web/src/app
    exit /B 0

:BuildServer
    rd /s /q ".\out\server" >NUL 2>&1
    for /f %%i in ('where.exe scour.exe 2^>NUL') do set SCOUR_LOC=%%i
    if NOT "%SCOUR_LOC%"=="" (
        scour.exe .\assets\icons.svg .\internal\web\static\img\icons.svg --enable-id-stripping --protect-ids-noninkscape --remove-descriptive-elements --enable-comment-stripping --strip-xml-prolog --strip-xml-space
    )

    set CGO_ENABLED=0
    set GOOS=%DEV_OS%
    set GOARCH=%DEV_ARCH%
    go build -o ./out/server/heyapple-%DEV_OS%-%DEV_ARCH%%DEV_EXT% github.com/and-rad/heyapple/cmd/server
    exit /B 0

:BuildCLI
    rd /s /q ".\out\cli" >NUL 2>&1

    set CGO_ENABLED=0
    set GOOS=%DEV_OS%
    set GOARCH=%DEV_ARCH%
    go build -o ./out/cli/heyapple-cli-%DEV_OS%-%DEV_ARCH%%DEV_EXT% github.com/and-rad/heyapple/cmd/cli
    exit /B 0

:RunServer
    .\out\server\heyapple-%DEV_OS%-%DEV_ARCH%%DEV_EXT%
    exit /B 0
