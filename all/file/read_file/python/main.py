#!/bin/env python3

file_name: str = "demo.txt"

file = open(file_name, 'r')
data: str = file.read()
file.close()

print(data)
