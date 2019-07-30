---
title: PHP JSON HTML Minify
author: Luis Rodriguez
type: post
date: 2015-09-15T22:36:25+00:00
comments: false
categories:
  - Projects
tags:
  - luis-portfolio
  - andrew-portfolio
  - maave-portfolio
  - php
  - json
  - html
  - output-buffer

---

Both of these functions are designed to be used with ob_start() and process the output before its sent to the user. You can also include caching into this to reduce overhead for static pages. It takes the output of the page and makes the content up to 50% smaller in size by removing unused comments and whitespace in code. [Download the file here](https://downloads.techreanimate.com/auowdk) or look at the code below.

<!--more-->

HTML minify function

```php
<?php

	/*
	 * Warning: do not call this in main includes, only in base of html pages
	 * Does not work with css, js, json. only html
	 * 
	 * Must be called like this ob_start("sanitize_page");
	 */

	function sanitize_page($buffer) {
		if(IS_DEV){
			return $buffer;
		}else{
			$search = array(
				'/\>[^\S ]+/s', /*strip whitespaces after tags, except space */
				'/[^\S ]+\</s', /*strip whitespaces before tags, except space */
				'/(\s)+/s' /* shorten multiple whitespace sequences */
			);
			
			$replace = array('>','<','\\1');
			
			return preg_replace($search, $replace, $buffer);
		}
		
	}
```

Json Minify Function

```php
<?php

	/*
	 * Warning: do not call this in main includes, only in base of html pages
	 * Does not work with css, js, html. only json
	 * 
	 * Must be called like this ob_start("json_minify");
	 */
	
	function json_minify($json) {
		if(IS_DEV){
			$tokenizer = "/\"|(\/\*)|(\*\/)|(\/\/)|\n|\r/";
			$in_string = false;
			$in_multiline_comment = false;
			$in_singleline_comment = false;
			$tmp; $tmp2; $new_str = array(); $ns = 0; $from = 0; $lc; $rc; $lastIndex = 0;
		
			while (preg_match($tokenizer,$json,$tmp,PREG_OFFSET_CAPTURE,$lastIndex)) {
				$tmp = $tmp[0];
				$lastIndex = $tmp[1] + strlen($tmp[0]);
				$lc = substr($json,0,$lastIndex - strlen($tmp[0]));
				$rc = substr($json,$lastIndex);
				if (!$in_multiline_comment && !$in_singleline_comment) {
					$tmp2 = substr($lc,$from);
					if (!$in_string) {
						$tmp2 = preg_replace("/(\n|\r|\s)*/","",$tmp2);
					}
					$new_str[] = $tmp2;
				}
				$from = $lastIndex;
		
				if ($tmp[0] == "\"" && !$in_multiline_comment && !$in_singleline_comment) {
					preg_match("/(\\\\)*$/",$lc,$tmp2);
					if (!$in_string || !$tmp2 || (strlen($tmp2[0]) % 2) == 0) {	/* start of string with ", or unescaped " character found to end string */
						$in_string = !$in_string;
					}
					$from--; /* include " character in next catch */
					$rc = substr($json,$from);
				}
				else if ($tmp[0] == "/*" && !$in_string && !$in_multiline_comment && !$in_singleline_comment) {
					$in_multiline_comment = true;
				}
				else if ($tmp[0] == "*/" && !$in_string && $in_multiline_comment && !$in_singleline_comment) {
					$in_multiline_comment = false;
				}
				else if ($tmp[0] == "//" && !$in_string && !$in_multiline_comment && !$in_singleline_comment) {
					$in_singleline_comment = true;
				}
				else if (($tmp[0] == "\n" || $tmp[0] == "\r") && !$in_string && !$in_multiline_comment && $in_singleline_comment) {
					$in_singleline_comment = false;
				}
				else if (!$in_multiline_comment && !$in_singleline_comment && !(preg_match("/\n|\r|\s/",$tmp[0]))) {
					$new_str[] = $tmp[0];
				}
			}
			$new_str[] = $rc;
			return implode("",$new_str);
		}else{
			return $json;
		}
	}
?>
```