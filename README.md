---

Silo City Labs website using Github Pages.

Learn more about Github Pages at https://pages.github.com and the official [documentation](https://help.github.com/en/categories/github-pages-basics). Anyone can create a pull request to publish an article or make a fix based on issues. We just ask that you copy one of the templates in /content/post/samples/. Only admins are allowed to save drafts in the repo.

---

### Clone this site

1. Fork, Clone this project
2. Add secret "EMAIL" to your repository github actions

## Building locally for Admins

To work locally with this project, you'll have to follow the steps below:

1. Fork, Clone this project
2. Download hugo v0.123.6 to root of this project https://github.com/gohugoio/hugo/releases/tag/v0.123.6
3. Make executable: `chmod +x hugo`
4. Preview your project: `./hugo server`, your site can be accessed under `localhost:1313/`.
5. Add content

## Github Actions

We use github actions to make it easy for anyone to fork this repo withouth making a standalone daemon. Our action is set to run on push and once a month.
See our build script inside workflows for github actions.


Read more at Hugo's [documentation](https://gohugo.io/overview/introduction/).
