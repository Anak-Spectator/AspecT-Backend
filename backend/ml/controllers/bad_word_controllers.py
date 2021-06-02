from posixpath import dirname
from better_profanity import profanity
import os


def post_new_bad_word(text_list=[{"text": "", "time": ""}]):
    if len(text_list) == 1 and text_list[0]["time"] == "":
        return None, "there is no text"

    data = []
    text = []

    for dict in text_list:
        for key, val in dict.items():
            is_bad = False
            if key == "text":
                is_bad = profanity.contains_profanity(val)

            if is_bad:
                data.append({"text": dict["text"], "date": dict["date"]})
                text.append(dict["text"])
    rootDir = os.path.abspath(os.getcwd())
    assetsDir = os.path.join(rootDir,"assets")
    with open(os.path.join(assetsDir,"textList.txt"), "r+") as f:
        f.truncate(0)
        f.writelines("\n".join(text))
        f.close()
        # pass

    return {"texts": data}, None
