import click

@click.command()
@click.option('--template', '-t', type=str, help='name of the gitignore template to use')
def run(template):
    """Generates a .gitignore template from githubs repo"""
    url_template = f'https://raw.githubusercontent.com/github/gitignore/main/{template}.gitignore'
