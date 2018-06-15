rd /s/q release
md release
go build -o restgo-admin.exe
COPY restgo-admin.exe release\
COPY favicon.ico release\favicon.ico
XCOPY config\*.* release\config\  /s /e 
XCOPY mnt\*.* release\mnt\  /s /e 
XCOPY asset\*.* release\asset\  /s /e 
XCOPY view\*.* release\view\  /s /e 
echo "打包成功,相关应用在release目录下"
