gominion
========

D&amp;D Administration Helper written in Go

## Dependencies
This app uses front end components managed by bower and grunt. Therefore, the
first thing that needs to happen is to install nodejs. The ubuntu package is
pretty woefully out of date, so I installed nodejs using the bash script
installer

````bash
curl -sL https://deb.nodesource.com/setup | sudo bash -
sudo apt-get install -y nodejs
````

Then I installed bower for dependency management

````bash
sudo npm install -g bower
````

In bower, I'm using
* bootstrap
* jquery (included by bootstrap)
* angularjs
* angularjs-ui-router

I get these bower dependencies by running `bower install <package> --save`

After following the prompts, I have a bower.json file. Next up: grunt.

````bash
npm install -g grunt-cli
````

Then to wire it all up, I made a package.json using npm:

````bash
npm init
npm install grunt--save-dev
npm install grunt-contrib-coffee --save-dev
npm install grunt-contrib-copy --save-dev
npm install grunt-contrib-uglify --save-dev
npm install grunt-contrib-watch --save-dev
````

And finally edited my `Gruntfile.js` to have it build my dependencies and drop
them in public/
