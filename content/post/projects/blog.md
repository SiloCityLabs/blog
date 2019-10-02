---
title: Silocitylabs Blog
type: post
date: 2015-09-15T22:36:25+00:00
comments: false
categories:
  - Projects
tags:
  - luis-portfolio
  - robert-portfolio
  - wordpress
  - hugo
  - seo

---

{{< image src="/uploads/luis-portfolio/blog.png" alt="Blog website" caption="The blog when this article was written.">}}

The Issue: 

Previously on Wordpress we started noticing performance issues early on into the websites time on the web. Wordpress unfortunately is not very good with performance on high traffic websites. 

The Solution:

We decided to try out a static site generator called hugo. After migrating we noticed extreme performance improvements. We made some more adjustments while the project was hot and decided to host on github public repo.

<!--more-->

Below you can see all the steps I took to move from Wordpress to Hugo + github pages and I went from a 45/100 seo score to currently now a 98/100 seo score on both mobile and desktop.

### Setting up hugo with a base theme

This was probably the simplest part of the entire project. Anyone can grab a theme right off of the Hugo website and get started. I decided not to make a custom theme but to just grab a bootstrap theme, Afterall bootstrap is most websites on the internet.

Download [the theme](https://themes.gohugo.io/beautifulhugo/)

After I grabbed the theme me and a friend started hacking away at all the bloat it had. Removing useless css and js like a js math symbols we knew we would never use.

### Export wordpress articles

Now that the theme is useable we decided to import some articles. We use a wordpress plugin to convert and export all the github pages to Markdown. It did a really good job at it but we still had to do some further tweaking.

Some of the custom html on the pages did not carry over as markdown. We began hacking away that html and converting it to markdown format. This only took less than a few days with the use of regex and pattern search.

[Wordpress to Hugo Export](https://downloads.techreanimate.com/toxlbb)

### Tweaking the theme

The theme had gotten us to basically where we were visually with wordpress. A basic blog site where users can read our posts and we can share the great knowledge we have learned. 

We wanted more now. We saw how great it was performing and decided we could improve it even more. Just by looking at our closed github issues you can see some of the [changes we made](https://github.com/SiloCityLabs/blog/issues?q=is%3Aissue+is%3Aclosed).

We added custom tags for better generated pages and made some quick seo changes. Little stuff here and there mostly.

### Comments system

A static site generated came with one major problem. No comments on my pages. After a bit of researching I found a service called [Disqus](https://disqus.com/) which is used by many major news orginizations and blogs. Its a universal login comment system that works with a simple js snippet.

It has some flaws that prevent us from getting 1 extra seo point on Google Pageinsights but its entirely worth it. We can enable and disable disqus by adding `comments: false` to any markdown page allowing us to turn it off for pages like About and Newsletter.

### Adding AMP support

The biggest seo change we made that actually does not reflect the seo score for page insights was adding AMP support to our website.

If you know AMP its a quick way for viewers to get articles by disabling much of the bottlenecks in todays websites.

Adding amp was quite a task, in doing so we ended up applying some of those modifications to our main website as well. We reduced the user of bootstrap submodules and shrunk the size of our css/js. We also ended up ditching jquery on the main website with the help of a library called bootstrap native which you can [read more about at sitepoint](https://www.sitepoint.com/use-bootstrap-components-without-jquery/).

You can see our amp support by prefixing /amp/ to any post url. For example [AMP: Installing nano text editor on Android](https://silocitylabs.com/amp/post/2019/09/26/installing-nano-android/)

### Lazy load

To further increase seo we took our custom markdown tag for hugo and added lazy load to images. This took our seo up by 4-5 points on Google pageinsights.

We took that and added it right to our main.js file to keep it consolidated.

### Minify and Combine CSS/JS

Hugo has a great built in feature that I beleive every hugo powered site should have enabled. I am honestly not sure why the theme did not have this done already. The setup was very easy. We took our files and combined the output withouth touching the files themselves. Just edit your templete like so:

```
{{ $jsmain := resources.Get "js/main.js" }}
{{ $jslazy := resources.Get "js/lazyload.js" }}
{{ $jshighlight := resources.Get "js/highlight.min.js" }}
{{ $jsnativebs := resources.Get "js/bootstrap-native-v4.min.js" }}

{{ $jssall := slice $jsnativebs $jshighlight $jslazy $jsmain  | resources.Concat "js/bundle.js" | resources.Minify }}
<script src="{{ $jssall.Permalink | absURL }}"></script>
```

### Setup github actions auto deploy

Github has granted our team Beta access to github actions wich allowed us to take our webhooks listener daemon and move it to a simpler build program. You can checkout the [build instructions here](https://github.com/SiloCityLabs/blog/blob/master/.github/workflows/build.yml).

Outline as follows:
 1. Sets up ubuntu
 2. Installs Go
 3. Installs Hugo
 4. Checks out Repo
 5. Runs Build Program
 6. Commit Changes back to repo
 7. Auto deployed using github pages.

### Re-enable Ads

At this point we felt comfortable with the site, its seo score and decided it was time to turn ads back on. We left this for the end to keep our seo score free of troublesome ads while we improved what we could. We kept it simple and used Google Autoads script and allowed it to put ads where it would like.

### Future stuff

We plan on taking the golang [build](https://github.com/SiloCityLabs/blog/tree/master/build) program and integrating it with newsletter and other more intense tasks that were normall handled by a few wordpress plugins. They will run when github actions auto deploys. You can checkout all our progress on [github issues page](https://github.com/SiloCityLabs/blog/issues)