

## 2022年04月11日10:55:53 周一

### 安装apiato报错

执行安装命令: `composer create-project apiato/apiato my-api`


```
PHP Fatal error:  Declaration of Composer\DependencyResolver\Rule2Literals::getLiterals() must be compatible with Composer\DependencyResolver\Rule::getLiterals(): array in /Users/jw/workspace/demo/apiato/vendor/composer/composer/src/Composer/DependencyResolver/	.php on line 48

Fatal error: Declaration of Composer\DependencyResolver\Rule2Literals::getLiterals() must be compatible with Composer\DependencyResolver\Rule::getLiterals(): array in /Users/jw/workspace/demo/apiato/vendor/composer/composer/src/Composer/DependencyResolver/Rule2Literals.php on line 48
```



## 2022年04月08日10:42:17 周五



## 2022年04月07日11:16:51 周四


[JSON转URL参数](https://blog.csdn.net/lihefei_coder/article/details/81417311)

[Sort JavaScript object by key](https://stackoverflow.com/questions/5467129/sort-javascript-object-by-key)


### 接口签名 apifox postman 前置操作

```
let key = "blog";
let secret = "i1ydX9RtHyuJTrw7frcu";

let date = new Date();
let datetime = date.getFullYear() + "-" // "年"
    + ((date.getMonth() + 1) > 10 ? (date.getMonth() + 1) : "0" + (date.getMonth() + 1)) + "-" // "月"
    + (date.getDate() < 10 ? "0" + date.getDate() : date.getDate()) + " " // "日"
    + (date.getHours() < 10 ? "0" + date.getHours() : date.getHours()) + ":" // "小时"
    + (date.getMinutes() < 10 ? "0" + date.getMinutes() : date.getMinutes()) + ":" // "分钟"
    + (date.getSeconds() < 10 ? "0" + date.getSeconds() : date.getSeconds()); // "秒"

let path = "/echo";
let method = "POST";
let params = {d:'d1', a:'a1', c: 'c1 c2*'};
let sortParams = Object.keys(params).sort().reduce(
  (obj, key) => { 
    obj[key] = params[key]; 
    return obj;
  }, 
  {}
);
console.log(sortParams);
// let sortParamsEncode = decodeURIComponent(jQuery.param(ksort(params)));
let sortParamsEncode = decodeURIComponent(parseParams(sortParams));
let encryptStr = path + "|" + method.toUpperCase() + "|" + sortParamsEncode + "|" + datetime;
// let digest = CryptoJS.enc.Base64.stringify(CryptoJS.HmacSHA256(encryptStr, secret));
// console.log({authorization: key + " " + digest, date: datetime});
function parseParams(data) {
    try {
        var tempArr = [];
        for (var i in data) {
            var key = encodeURIComponent(i);
            var value = encodeURIComponent(data[i]);
            tempArr.push(key + '=' + value);
        }
        var urlParamsStr = tempArr.join('&');
        return urlParamsStr;
    } catch (err) {
        return '';
    }
}
```

### 涂鸦sign生成

```
(function () { 
    var timestamp = getTime();
    pm.environment.set("timestamp",timestamp);
    var clientId = pm.environment.get("client_id");
    var secret = pm.environment.get("secret");
    var sign = calcSign(clientId,secret,timestamp);
    pm.environment.set('easy_sign', sign);
})();

function getTime(){
    var timestamp = new Date().getTime();
    return timestamp;
}

function calcSign(clientId,secret,timestamp){
    var str = clientId + timestamp;
    var hash = CryptoJS.HmacSHA256(str, secret);
    var hashInBase64 = hash.toString();
    var signUp = hashInBase64.toUpperCase();
    return signUp;
}
```


### go run main.go -env fat 运行报错

```
go run main.go -env fat
# github.com/shirou/gopsutil/disk
iostat_darwin.c:28:2: warning: 'IOMasterPort' is deprecated: first deprecated in macOS 12.0 [-Wdeprecated-declarations]
/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX.sdk/System/Library/Frameworks/IOKit.framework/Headers/IOKitLib.h:132:1: note: 'IOMasterPort' has been explicitly marked deprecated here
# github.com/shirou/gopsutil/host
smc_darwin.c:75:41: warning: 'kIOMasterPortDefault' is deprecated: first deprecated in macOS 12.0 [-Wdeprecated-declarations]
/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX.sdk/System/Library/Frameworks/IOKit.framework/Headers/IOKitLib.h:123:19: note: 'kIOMasterPortDefault' has been explicitly marked deprecated here
exit status 1
```

处理方法: `go env -w CGO_ENABLED="0"`


### 覆盖变量

export PATH="/usr/local/opt/mysql@5.7/bin:$PATH"
export PATH="/usr/local/opt/php@7.4/bin:$PATH"
export PATH="/usr/local/opt/php@7.4/sbin:$PATH"

mv composer.phar /usr/local/opt/composer