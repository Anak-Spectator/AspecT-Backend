from flask import request, jsonify
from controllers.bad_word_controllers import post_new_bad_word
from utils.resp_out import err_resp, new_resp
from utils import status_code
import json

# initializing app service
def init_bad_word_routes(app):
    @app.route("/api/v1/user/profanity/send", methods=["GET", "POST"])
    def route():
        if request.method == "GET":
            return jsonify(output="this is json response")

        if request.method == "POST":
            body = request.json

            text_list = body["texts"]
            resp, err = post_new_bad_word(text_list)

            if err is not None:
                return jsonify(err_resp(err)), status_code.SERVER_ERROR

            return (
                jsonify(new_resp("success send data", resp, status_code.CREATED)),
                status_code.CREATED,
            )
