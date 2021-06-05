import os
import csv
from flask import request, jsonify
from utils import status_code
from utils.resp_out import err_resp, new_resp
from utils.helper import get_user_id, open_dir

def init_list_profanity_routes(app):
    @app.route("/api/v1/user/check", methods=["GET"])
    def check_profanity_route():
        if request.method == "GET":
            listBad = []
            userId = get_user_id()
            with open(os.path.join(open_dir('assets'), userId), "r+") as f:
                reader = csv.reader(f, delimiter=",")
                line = 0
                for row in reader:
                    listBad.append({"id":row[0], "test":row[1], "date":row[2]})

            return (
                jsonify(new_resp("success send data", {"texts": listBad}, status_code.CREATED)),
                status_code.CREATED,
            )

