Folder PATH listing for volume Games
Volume serial number is 4EB8-6771
D:.
|   .gitignore
|   build_and_deploy.ps1
|   config.json
|   config.template.json
|   go.mod
|   go.sum
|   LICENSE
|   WinSenseConnect.exe
|   WinSenseConnect.log
|   notes.txt
|   project_structure.md
|   readme.md
|   
+---backend
|   |   config.go
|   |   http_server.go
|   |   logger.go
|   |   main.go
|   |   mqtt.go
|   |   service.go
|   |   windows.go
|   |   
|   \---tmp
|           main.exe
|           WinSenseConnect.log
|           
+---frontend
|   |   .gitignore
|   |   app.vue
|   |   nuxt.config.ts
|   |   package-lock.json
|   |   package.json
|   |   README.md
|   |   tsconfig.json                       
|   +---app
|   |       spa-loading-template.html
|   |       
|   +---assets
|   |   \---css
|   |           tailwind.css
|   |                                       
|   +---pages
|   |   |   index.vue
|   |   |   
|   |   \---config
|   |           mqtt.vue
|   |           scripts.vue
|   |           
|   +---public
|   |       favicon.ico
|   |       robots.txt
|   |       
|   \---server
|           tsconfig.json
|           
\---scripts
        monitors_to_mac.ps1
        monitors_to_pc.ps1
        test_notification.ps1
        
