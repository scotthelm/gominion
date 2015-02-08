var pkgjson = require('./package.json');
 
var config = {
  pkg: pkgjson,
  app: 'bower_components',
  dist: 'assets'
}
 
module.exports = function (grunt) {
 
  // Configuration
  grunt.initConfig({
    config: config,
    pkg: config.pkg,
    copy: {
      dist: {
       files: [{
         expand: true,
         cwd: '<%= config.app %>/font-awesome/css',
         src: 'font-awesome.min.css',
         dest: '<%= config.dist %>'
       },
       {
         expand: true,
         cwd: '<%= config.app %>/bootstrap/dist/css',
         src: 'bootstrap.min.css',
         dest: '<%= config.dist %>'
       },
       {
         expand: true,
         cwd: '<%= config.app %>/bootstrap/dist/fonts',
         src: './*',
         dest: '<%= config.dist %>'
       },
       {
         expand: true,
         cwd: '<%= config.app %>/font-awesome/fonts',
         src: './*',
         dest: '<%= config.dist %>'
       }]
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
            '<%= config.app %>/jquery/dist/jquery.js',
            '<%= config.app %>/bootstrap/dist/js/bootstrap.js',
          ]
        }
      }
    }
  });
 
  grunt.loadNpmTasks('grunt-contrib-copy');
  grunt.loadNpmTasks('grunt-contrib-uglify');
 
  grunt.registerTask('default', [
    'copy',
    'uglify',
  ]);
};
