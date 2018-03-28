for /r "bin/odpsLogs" %%i in (*.txt) do mysql -h localhost -u root -p123456 bbc_odps < %%i
pause