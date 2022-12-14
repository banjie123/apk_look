配合apkleaks进行使用
  -a string
    	Specify the path where the calling file is located -- 指定apkleaks.py所在路径
  -o string
    	The path where result save (default "result.txt")  -- 指定结果存储路径
  -p string
    	The folder where apk file save -- 指定apk文件所在文件夹

# apk_look
调用apkleaks批量跑apk文件，获取敏感信息
Windows运行效果
apk_look.exe -p C:\Users\19244\Desktop\huawei_apk -a C:\Users\19244\Desktop\apk\apkleaks-master\apkleaks-master\apkleaks.py -o result.txt
![image](https://user-images.githubusercontent.com/89896919/207629271-4eda4fe8-5e4b-40f1-bf76-1168682134ba.png)
![image](https://user-images.githubusercontent.com/89896919/207629370-b3e81299-67c8-4121-be1b-0590ff85d0b2.png)
Linux运行效果
 ./apk -a /root/apk/apkleaks/apkleaks.py -p /root/apk/apkleaks/apks/huawei_apk -o result.txt
 ![image](https://user-images.githubusercontent.com/89896919/207629627-9ae07c5a-3c41-46fd-a902-a56512988d07.png)
![image](https://user-images.githubusercontent.com/89896919/207629629-fe3ddf6c-79e6-4f42-bf27-1556dda08da5.png)

apk下载链接
https://github.com/dwisiswant0/apkleaks
