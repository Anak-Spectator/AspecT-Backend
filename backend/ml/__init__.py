import os
import csv
import uuid
from utils.helper import open_dir, get_user_id
from controllers.bad_word_controllers import check_new_bad_word
from models.db import send_db

if not os.listdir(open_dir('assets')):
    f=open(os.path.join('assets', str(uuid.uuid4().hex)), 'w')
    f.close()

check_new_bad_word("FUCK BITCH DAMN")

with open(os.path.join(open_dir('assets'), get_user_id()), "r") as f:
    reader = csv.reader(f)
    data = list(reader)

send_db(data)

