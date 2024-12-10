#!/bin/bash

day=${1}
dayName=${2}

if [ -z "${day}" ]; then
    echo "You must provide a day number"
    exit 1
fi
if [ -z "${dayName}" ]; then
    echo "You must provide a day name"
    exit 1
fi

if [ -f days/${day}.go ]; then
    echo "Day ${day} already created"
    exit 1
fi

cp days/blank.go "days/${day}.go"
cp days/blank_test.go "days/${day}_test.go"

sed -i "s/Blank/${dayName}/g" days/${day}.go
sed -i "s/Blank/${dayName}/g" days/${day}_test.go
sed -i "s/\/\/DAY PLACEHOLDER\/\//case ${day}:\nswitch *part {\ncase 1:\nout = strconv.Itoa(days.Day${dayName}Part1(lines))\ncase 2:\nout = strconv.Itoa(days.Day${dayName}Part2(lines))\n}\n\/\/DAY PLACEHOLDER\/\//" main.go

go fmt ./...