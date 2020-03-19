#!/bin/bash
workdir=Log-MeterAgent
appname=meteragent
servicename=LogMeterAgent
lowservicename=logmeteragent
sed -i  "s/MP_buy_API/$workdir/g" `grep MP_buy_API -rl ./`
sed -i  "s/mpbuyservice/$lowservicename/g" `grep mpbuyservice -rl ./`
sed -i  "s/mpbuy/$appname/g" `grep mpbuy -rl ./`
sed -i  "s/MpBuyService/$servicename/g" `grep MpBuyService -rl ./`
sed -i  "s/MpbuyService/$servicename/g" `grep MpbuyService -rl ./`

mv pkg/api/mpbuy pkg/api/$appname
mv pkg/apiserver/mpbuy pkg/apiserver/$appname
mv pkg/apis/mpbuy pkg/apis/$appname
mv cmd/mpbuyservice cmd/${lowservicename}
mv package/mpbuyservice package/${lowservicename}
mv package/mpbuyservice/deploy/mpbuyservice.yaml package/${lowservicename}/deploy/${lowservicename}.yaml 
