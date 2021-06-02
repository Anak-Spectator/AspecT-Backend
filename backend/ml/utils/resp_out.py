from utils import status_code

from . import status_code


def new_resp(msg, data, status=status_code.OK):
    return {"message": msg, "data": data, "status_code": status}


def err_resp(err):
    return {"error": err}
