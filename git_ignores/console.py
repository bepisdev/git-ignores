import click
import json
import requests
from pathlib import Path

def _write_to_file(filename, data):
    """Helper method that writes data to the given file"""
    filet = open(filename, 'w')
    filet.write(data)
    filet.close()

@click.command()
@click.option('--template', '-t', type=str, help='name of the gitignore template to use')
@click.option('--force', '-f', default=False, help='replace existing gitignore file if one exists')
def run(template, force):
    """Generates a .gitignore template from githubs repo"""
    url_template = f'https://raw.githubusercontent.com/github/gitignore/main/{template}.gitignore'
    req = requests.get(url_template)
    if req.status_code == 200:
        data = req.text
        fileexists = False
        
        if Path('./.gitignore').is_file():
            fileexists = True
        
        if fileexists is True:
            if force is False:
                print(".gitignore already exists, goodbye!")
                return
        
        _write_to_file('./.gitignore', data)