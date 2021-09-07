#!/usr/bin/env python

import sys
import re

toorep = []

with open(sys.argv[1], "r") as f:
    b = f.readlines()

    for i, line in enumerate(b):
        if "map[string]interface{}" in line and "raw" not in line and "type" in line:
            bla = re.match(".* (.*) .*", line)
            search = bla.groups()[0]
            print(search)
            for j, line2 in enumerate(b):
                if search in line2 and "type" not in line2:
                    print(line2)
                    replacer = (search, i, j)
                    toorep.append(replacer)

for replace in toorep:
    search, i, j = replace
    print("___", search, i ,j)
    print(b[i])
    print(b[j])
    b[i] = "//" + b[i]
    b[j] = str.replace(b[j], search, "*apiextensions.JSON")
    print(b[i])
    print(b[j])

with open(sys.argv[1], "w") as f:
    f.writelines(b)