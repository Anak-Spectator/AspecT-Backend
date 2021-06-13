import os
import csv
import time
import datetime
from better_profanity import profanity
from utils.helper import get_user_id, open_dir

def add_bad_word(text):
    user_id = get_user_id()
    timestamp = datetime.datetime.fromtimestamp(time.time()).strftime('%Y-%m-%d %H:%M')
    with open(os.path.join(open_dir("assets"), user_id), mode='a') as f:
        writer = csv.writer(f, delimiter=',', quotechar='"', quoting=csv.QUOTE_MINIMAL)
        #writer = csv.DictReader(f, fieldnames=[user_id, text, timestamp]);
        writer.writerow([user_id,text,timestamp])

def check_new_bad_word(textList):
    if profanity.contains_profanity(textList):
        add_bad_word(textList)
