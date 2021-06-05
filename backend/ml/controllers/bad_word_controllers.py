import os
import csv
import time
from better_profanity import profanity
from utils.helper import get_user_id, open_dir

def add_bad_word(text):
    user_id = get_user_id()
    with open(os.path.join(open_dir("assets"), user_id), mode='a') as f:
        writer = csv.writer(f, delimiter=',', quotechar='"', quoting=csv.QUOTE_MINIMAL)
        writer.writerow([user_id,text,time.time()])

def check_new_bad_word(textList):
    for word in list(textList.split("  ")):
        if profanity.contains_profanity(word) == True:
            add_bad_word(word)
