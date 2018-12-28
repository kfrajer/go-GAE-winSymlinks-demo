REFERENCE:  https://stackoverflow.com/questions/35230998/how-to-upload-a-google-app-engine-go-project-in-a-different-folder-than-the-ap/35261641#35261641

https://unify.my.salesforce.com/500f200001KVJTBAA5


TECHNICAL:

https://cloud.google.com/appengine/docs/standard/go/building-app/
https://cloud.google.com/appengine/docs/standard/go/configuration-files
app.yaml  https://cloud.google.com/appengine/docs/standard/go/config/appref
Symbolic links in Windows: https://www.techrepublic.com/article/how-to-take-advantage-of-symbolic-links-in-window-10/

mklink /D linkname folder-source
mklink /D D:\workspace\go\Projects\hw-1\appengine\build D:\workspace\go\Projects\hw-1\build


SETTING UP GO:

Path:   D:\Program Files\go1.11.2\Bbin
GOROOT: D:\Program Files\go1.11.2
GOPATH: D:\workspace\go\Projects
go get -u google.golang.org/appengine/...

dev_appengine.py app.yaml

Tested:
[LOCAL]
http://localhost:8080
http://localhost:8080/favicon.ico
http://localhost:8080/images/favicon.ico

[Deployed version go-extdir]
It does not work if deployed without corections

[Deployed version go-extdir - with corrections]
- Corrections using MKLink
- Update app.yaml references following static handlers described here https://cloud.google.com/appengine/docs/standard/go/config/appref
- Update myScript.go to fetch html from proper location
- Update html to get static resources from its corrent relative location
