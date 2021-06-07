import os
import re
from flask import request, jsonify
from utils import status_code
from utils.resp_out import err_resp, new_resp

def get_user_id(dirr):
    for f in os.listdir(dirr):
        user_compile = re.compile('[0-9a-f]{32}\Z', re.I)
        user_match = re.match(user_compile, f)
        user_id = user_match.group(0)
        
    return user_id

def init_list_profanity_routes(app):
    @app.route("/api/v1/user/check", methods=["GET"])
    def check_profanity_route():
        if request.method == "GET":
            listBad = []
            rootDir = os.path.abspath(os.getcwd())
            assetsDir = os.path.join(rootDir,"assets")
            userId = get_user_id(assetsDir) 
            with open(os.path.join(assetsDir, userId), "r+") as f:
                for line in f:
                    sline = line.rstrip()
                    listBad.append({"text":sline,"id":userId,"timestamp":"08-04-2021"})
            
            return (
                jsonify(new_resp("success send data", {"texts": listBad}, status_code.CREATED)),
                status_code.CREATED,
            )

