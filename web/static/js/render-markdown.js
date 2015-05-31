/**
* @version: 0.0.1
* @author: Indraniel Das https://github.com/indraniel
* @copyright: Copyright (c) 2015 Indraniel Das
* @license: Licensed under the BSD license. See http://choosealicense.com/licenses/bsd-2-clause/
*/

$(function() {
  markdownElement = document.getElementsByTagName('xmp')[0];
  markdown = markdownElement.textContent || markdownElement.innerText;
  mdTxt = marked(markdown);
  document.getElementById('example-content').innerHTML = mdTxt;
});
