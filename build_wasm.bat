@echo off

set GOOS=js
set GOARCH=wasm

del /f C:\Users\nicklvsa\Desktop\Gio-Testing\web\wasm_exec.js
echo f | xcopy /s T:\Go\misc\wasm\wasm_exec.js C:\Users\nicklvsa\Desktop\Gio-Testing\web\wasm_exec.js

go build -o main.wasm