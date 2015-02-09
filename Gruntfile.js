var pkgjson = require('./package.json');
 
var config = {
  pkg: pkgjson,
  lib: 'bower_components',
  dist: 'public',
  app: 'assets/src'
}
 
module.exports = function (grunt) {
 
  // Configuration
  grunt.initConfig({
    config: config,
    pkg: config.pkg,
    watch: {
      tasks: ['coffee', 'copy'],
      files: ['<%= config.app %>/coffee/*.coffee']
    },
    copy: {
      dist: {
       files: [{
         expand: true,
         cwd: '<%= config.lib %>/font-awesome/css',
         src: 'font-awesome.min.css',
         dest: '<%= config.dist %>'
       },
       {
         expand: true,
         cwd: '<%= config.lib %>/bootstrap/dist/css',
         src: 'bootstrap.min.css',
         dest: '<%= config.dist %>'
       },
       {
         expand: true,
         cwd: '<%= config.lib %>/bootstrap/dist/fonts',
         src: './*',
         dest: '<%= config.dist %>'
       },
       {
         expand: true,
         cwd: '<%= config.lib %>/font-awesome/fonts',
         src: './*',
         dest: '<%= config.dist %>'
       },
       {
         expand: true,
         cwd: '<%= config.app %>/js/',
         src: 'app.js',
         dest: '<%= config.dist %>'
       }]
      }
    },
    coffee: {
      compile: {
        files: {
          '<%= config.app %>/js/app.js': ['<%= config.app %>/coffee/*.coffee'],
        }
      }
    },
    uglify: {
      options: {
        banner: '/*! <%= pkg.name %> lib - v<%= pkg.version %> -' +
          '<%= grunt.template.today("yyyy-mm-dd") %> */'
      },
      dist: {
        files: {
          '<%= config.dist %>/lib.min.js': [
            '<%= config.lib %>/jquery/dist/jquery.js',
            '<%= config.lib %>/bootstrap/dist/js/bootstrap.js',
            '<%= config.lib %>/angularjs/angular.js',
            '<%= config.lib %>/angular-ui-router/release/angular-ui-router.js',
            '<%= config.lib %>/angular-resource/angular-resource.min.js',
          ]
        }
      }
    }
  });

  grunt.loadNpmTasks('grunt-contrib-copy');
  grunt.loadNpmTasks('grunt-contrib-uglify');
  grunt.loadNpmTasks('grunt-contrib-coffee');
  grunt.loadNpmTasks('grunt-contrib-watch');

  grunt.registerTask('default', [
    'coffee',
    'copy',
    'uglify',
    'watch',
  ]);

};
