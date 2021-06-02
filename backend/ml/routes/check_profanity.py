import os
from flask import request, jsonify
from utils import status_code
from utils.resp_out import err_resp, new_resp

def init_list_profanity_routes(app):
    @app.route("/api/v1/user/check", methods=["GET"])
    def check_profanity_route():
        if request.method == "GET":
            listBad = []
            rootDir = os.path.abspath(os.getcwd())
            assetsDir = os.path.join(rootDir,"assets")
            with open(os.path.join(assetsDir,"textList.txt"), "r+") as f:
                for line in f:
                    listBad.append({"text":line,"date":"08-04-2021"})
            
            return (
                jsonify(new_resp("success send data", {"texts": listBad}, status_code.CREATED)),
                status_code.CREATED,
            )

