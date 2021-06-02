from .bad_word_routes import init_bad_word_routes
from .check_profanity import init_list_profanity_routes

# routing all service
def routes(app):
    init_bad_word_routes(app)
    init_list_profanity_routes(app)