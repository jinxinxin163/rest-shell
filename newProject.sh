#!/bin/bash
workdir=rest-shell
appname=restshell
servicename=ShellService
sed -i  "s/MP_buy_API/$workdir/g" `grep MP_buy_API -rl ./`
sed -i  "s/mpbuy/$appname/g" `grep mpbuy -rl ./`
sed -i  "s/MpBuyService/$servicename/g" `grep MpBuyService -rl ./`

mv pkg/api/mpbuy pkg/api/$appname
mv pkg/apiserver/mpbuy pkg/apiserver/$appname
mv pkg/apis/mpbuy pkg/apis/$appname
mv cmd/mpbuyservice cmd/${appname}service
