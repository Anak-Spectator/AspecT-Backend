import os

DEBUG   = not (os.getenv("ENV") == "prod")
PORT    = 7000
HOST    = "0.0.0.0"
USER    = "marionette"
PASS    = "12345678"
DB      = "aspect"

