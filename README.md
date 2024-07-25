# OpenRed
``` Open Redirect Scanner```

The problem. Sometimes admins make a mistakes inside configuration when it ness to redirect from http to https. 
For example nginx:

if ($uri !~ /page){
        return 301 https://$host$request_uri;
    }

$host - is variable and $request_uri - is variable. So, it is possible to change from user request header Host: 

The attacker can change the Host header (for example though xss) for phihing goals. 

``` Usage: ```

go run OpenRed.go www.fsaexample.com

For Scan list of targets use OpenRed.sh:

strings=('target1.com' 'target2.ru' 'target3.gg' )  --> change here the list for your purpose

--<-@  With Love from ZL...
