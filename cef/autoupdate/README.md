### energy lib framework auto update

#### Divided into three modules
* liblcl
* enregy
* cef

#### Control whether to enable the module update based on enable
#### Field description
> first level node is the module name
> 
> download: module download source
> 
> url: module download source url template
> 
> source: module download source, addrs
> 
> sourceSelect: Download the address source selection, select the source based on the index, and replace it with the Url template
> 
> enable: module update based on enable
> 
> latest: module latest version
> 
> versions: version list
> 
>> forced: forced update, a significant update that applies to the current version
>> 
>> content: update logs

---

#### update json template
```json
{
  "liblcl": {
    "download": {
      "url": "https://{url}/energye/energy/releases/download/{version}/{OSARCH}.zip",
      "source": [
        "gitee.com",
        "github.com"
      ],
      "sourceSelect": 0
    },
    "enable": true,
    "latest": "version.1",
    "versions": {
      "version.1": {
        "energyVersion": "version.1",
        "content": [
          "update log 1",
          "update log 2"
        ],
        "forced": false
      },
      "version.0": {
        "energyVersion": "",
        "content": [
          ""
        ],
        "forcede": false
      }
    }
  },
  "energy": {
    "download": {
      "url": "",
      "source": [],
      "sourceSelect": 0
    },
    "enable": false,
    "latest": "",
    "versions": {
      "": {
        "content": [
          ""
        ],
        "forcede": false
      }
    }
  },
  "cef": {
    "download": {
      "url": "",
      "source": [
        ""
      ],
      "sourceSelect": 0
    },
    "enable": false,
    "latest": "",
    "versions": {
      "": {
        "content": [
          ""
        ],
        "forcede": false
      }
    }
  }
}
```

---