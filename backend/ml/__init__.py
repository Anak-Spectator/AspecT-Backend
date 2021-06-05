import os
import uuid
from utils.helper import open_dir
from controllers.bad_word import check_new_bad_word

if not os.listdir(open_dir('assets')):
    open(os.path.join('assets', str(uuid.uuid4().hex)), 'w').close()

check_new_bad_word("Bullshit")
