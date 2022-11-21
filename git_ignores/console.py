import click
import json
import requests
from pathlib import Path

@click.command()
@click.option('--template', '-t', type=str, help='name of the gitignore template to use')
def run(template):
    """Generates a .gitignore template from githubs repo"""
    url_template = f'https://raw.githubusercontent.com/github/gitignore/main/{template}.gitignore'
    req = requests.get(url_template)
    if req.status_code == 200:
        data = req.text
        if Path('./.gitignore').is_file():
            f = open('./.gitignore', 'w')
            f.write(data)
            f.close()
            print('done')
        else:
            print(".gitignore file already")
