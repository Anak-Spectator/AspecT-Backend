import os
import re

def get_user_id():
    dirId = open_dir("assets")
    for f in os.listdir(dirId):
        userCompile = re.compile('[0-9a-f]{32}\Z', re.I)
        userMatch = re.match(userCompile, f)
        userId = userMatch.group(0)

    return userId

def open_dir(dirname):
    rootDir = os.path.abspath(os.getcwd())
    subDir = os.path.join(rootDir, dirname)

    return subDir

