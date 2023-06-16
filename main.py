import os,sys
from os import path
from flask import Flask
from flask_login import LoginManager, current_user
from flask_uploads import DOCUMENTS, IMAGES, TEXT, UploadSet, configure_uploads
from flask_cors import CORS
from flaskwebgui import FlaskUI
from werkzeug.utils import secure_filename
from werkzeug.datastructures import  FileStorage
from datetime import timedelta

from database import create_db, get_migrate


from controllers import (
    setup_jwt,
    setup_flask_login
)

from views import (
    user_views,
    transactionLog_views,
    rentType_views,
    rent_views,
    student_views,
    locker_views,
    area_views,
    index_views,
    masterkey_views,
    key_views,
    report_views,
)

# New views must be imported and added to this list
views = [
    user_views,
    transactionLog_views,
    rentType_views,
    rent_views,
    student_views,
    locker_views,
    area_views,
    index_views,
    masterkey_views,
    key_views,
    report_views,
]

def add_views(app, views):
    for view in views:
        app.register_blueprint(view)


def loadConfig(app, config):
    app.config['ENV'] = os.environ.get('ENV', 'DEVELOPMENT')
    if app.config['ENV'] == "DEVELOPMENT":
        app.config.from_object('config')
        app.config['GIT_ENV'] = ""
        app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///' + os.path.join(os.environ.get('LOCALAPPDATA'),'test_database.db')
    else:
        app.config['SQLALCHEMY_DATABASE_URI'] = os.environ.get('SQLALCHEMY_DATABASE_URI')
        app.config['SECRET_KEY'] = os.environ.get('SECRET_KEY')
        app.config['JWT_EXPIRATION_DELTA'] =  timedelta(days=int(os.environ.get('JWT_EXPIRATION_DELTA')))
        app.config['DEBUG'] = os.environ.get('ENV').upper() != 'PRODUCTION'
        app.config['ENV'] = os.environ.get('ENV')
        
    for key, value in config.items():
        app.config[key] = config[key]

def create_app(config={}):
    app = Flask(__name__, static_url_path='/static')
    CORS(app)
    loadConfig(app, config)
    app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False
    app.config['TEMPLATES_AUTO_RELOAD'] = True
    app.config['PREFERRED_URL_SCHEME'] = 'https'
    app.config['UPLOADED_PHOTOS_DEST'] = "App/uploads"
    photos = UploadSet('photos', TEXT + DOCUMENTS + IMAGES)
    configure_uploads(app, photos)
    add_views(app, views)
    create_db(app)
    setup_jwt(app)
    setup_flask_login(app)
    app.app_context().push()
    return app

def start_flask(**server_kwargs):

    app = server_kwargs.pop("app", None)
    server_kwargs.pop("debug", None)

    try:
        import waitress

        waitress.serve(app, **server_kwargs)
    except:
        app.run(**server_kwargs)

app = create_app()
migrate = get_migrate(app)

if __name__ == "__main__":
    if app.config['GIT_ENV'] == "GITPOD" or app.config['ENV'].upper() == "PRODUCTION":
        app.run()
    else: 
        FlaskUI(app=app,width=1366, height=768, server=start_flask,server_kwargs={
            "app": app,
            "port": 3000,
        }).run()

        