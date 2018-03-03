rd /s/q release
md release
go build
COPY restgo.exe release\restgo.exe
COPY favicon.ico release\favicon.ico
XCOPY config\*.* release\config\  /s /e 
XCOPY mnt\*.* release\mnt\  /s /e 
XCOPY asset\*.* release\asset\  /s /e 
XCOPY view\*.* release\view\  /s /e
