---

Silo City Labs website using Github Pages.

Learn more about Github Pages at https://pages.github.com and the official [documentation](https://help.github.com/en/categories/github-pages-basics). Anyone can create a pull request to publish an article or make a fix based on issues. We just ask that you copy one of the templates in /content/post/samples/. Only admins are allowed to save drafts in the repo.

---

## Building locally for Admins

To work locally with this project, you'll have to follow the steps below:

1. Fork, Clone this project
2. [Install](https://gohugo.io/overview/installing/) Hugo
3. Preview your project: `hugo server`
4. Add content
5. Generate the website: `./build.sh`

Read more at Hugo's [documentation](https://gohugo.io/overview/introduction/).

### Preview your changes

If you clone or download this project to your local computer and run `hugo server`,
your site can be accessed under `localhost:1313/hugo/`.


### Setup

Install to `/opt/blog/`

Enable service `systemctl link /opt/blog/daemon/blog.service`

Copy settings file and edit

`cp settings.default.yaml settings.yaml`

Start Service `systemctl start blog`

Make sure port forwarding is handled on your routers