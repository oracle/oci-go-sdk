#!/bin/bash
s=$(cd ${1%%/};pwd); d=$(cd $2;pwd); b=; while [ "${d#$s/}" == "${d}" ]
do s=$(dirname $s);b="../${b}"; done; echo ${b}${d#$s/}
