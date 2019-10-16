var main = {

  bigImgEl : null,
  numImgs : null,

  init : function() {
    var navbar = document.querySelector('.navbar');
    var main_navbar = document.querySelector('#main-navbar');
    var navlinks = document.querySelectorAll('.navlinks-parent');
    var searchBtn = document.querySelector('#searchBtn');
    var gitFooter = document.querySelector('#gitBuild');
    
    // Search bar
    searchBtn.addEventListener('click', function (e) {
      //TODO: close navbar
      setTimeout( function(){ document.getElementById('gsc-i-id1').focus(); }, 400);
    });

    // Shorten the navbar after scrolling a little bit down
    var prevScrollpos = window.pageYOffset;
    window.onscroll = function() {
      var currentScrollPos = window.pageYOffset;
      if (prevScrollpos > currentScrollPos) {
        navbar.classList.remove('top-nav-short');
      } else {
        navbar.classList.add('top-nav-short');
      }
      prevScrollpos = currentScrollPos;
    }
    
    // On mobile, hide the avatar when expanding the navbar menu
    main_navbar.addEventListener('show.bs.collapse', function () {
      navbar.classList.add('top-nav-expanded');
    });
    main_navbar.addEventListener('hidden.bs.collapse', function () {
      navbar.classList.remove('top-nav-expanded');
    });
  
    // On mobile, when clicking on a multi-level navbar menu, show the child links
    navlinks.forEach(function(link,i){
      link.addEventListener('click', function (e) {
        var target = e.target
        navlinks.forEach(function(child,key){
          if (child == target) {
            child.parentNode.classList.toggle("show-children");
          } else {
            child.parentNode.classList.remove("show-children");
          }
        });
      });
    });

    var lazyLoadInstance = new LazyLoad({
        elements_selector: ".lazy"
    });

    // show the big header image  
    main.initImgs();

    fetch(`https://api.github.com/repos/SiloCityLabs/blog/commits`).then(function(response) {
      return response.json();
    }).then(function(commits) {
      gitFooter.innerHTML = "Last built on " + new Date(commits[0].commit.author.date);
    });
  },
  
  initImgs : function() {
    // If the page was large images to randomly select from, choose an image
    if (document.querySelector("#header-big-imgs") !== null) {
      main.bigImgEl = document.querySelector("#header-big-imgs");
      main.numImgs = main.bigImgEl.getAttribute("data-num-img");

      // set an initial image
      var imgInfo = main.getImgInfo();
      var src = imgInfo.src;
      var desc = imgInfo.desc;
      main.setImg(src, desc);
    }
  },
  
  // Select random image and get the info for it
  getImgInfo : function() {
    var randNum = Math.floor((Math.random() * main.numImgs) + 1);
    var src = main.bigImgEl.getAttribute("data-img-src-" + randNum);
    var desc = main.bigImgEl.getAttribute("data-img-desc-" + randNum);
    
    return {
      src : src,
      desc : desc
    }
  },
  
  setImg : function(src, desc) {
    var header = document.querySelector(".intro-header.big-img");
    var imgdesc = document.querySelector(".img-desc");

    header.style.backgroundImage = 'url(' + src + ')';


    if (typeof desc !== typeof undefined && desc !== false) {
      imgdesc.textContent = desc;
      imgdesc.style.display = '';
    } else {
      imgdesc.style.display = 'none';
    }
  }
};

document.addEventListener('DOMContentLoaded', main.init);
