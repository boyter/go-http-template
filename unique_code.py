#!/usr/bin/env python3

# This program runs through the Go code ensuring that the UniqueCode used for logging is indeed
# unique, and if not printing out the duplicated lines and returning a non zero exit code

import os
import re
import sys

found_unique = []
totally_unique = True

for root, dirs, files in os.walk(".", topdown=True):

    dirs[:] = [d for d in dirs if d not in ["vendor", ".git", ".idea"]]
    files[:] = [x for x in files if ".go" in x]

    for name in files:
        with open(os.path.join(root, name), "r", encoding="utf-8") as infile:
            cnt = 1
            lines = []
            for line in infile:
                if "common.UniqueCode" in line:

                    unique_key = re.findall('"[a-z0-9]{8}"', line)
                    if len(unique_key) != 0:

                        if unique_key[0] in found_unique:
                            lines.append("%d %s" % (cnt, line))
                            totally_unique = False
                        else:
                            found_unique.append(unique_key[0])

                cnt += 1

            if len(lines) != 0:
                print(os.path.join(root, name))
                for l in lines:
                    print(l.strip())


if totally_unique:
    sys.exit(0)
else:
    sys.exit(1)
