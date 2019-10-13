import os
import logging
import subprocess

PROJECT_DIRECTORY = os.path.realpath(os.path.curdir)


if '{{cookiecutter.type}}'.lower() == "http":
    os.remove(os.path.join(PROJECT_DIRECTORY,  "msg_handlers.go"))

if '{{cookiecutter.type}}'.lower() == "msg":
    os.remove(os.path.join(PROJECT_DIRECTORY,  "http_handlers.go"))

if '{{cookiecutter.use_gcs}}'.lower() == "no":
    os.remove(os.path.join(PROJECT_DIRECTORY,  "storage.go"))
