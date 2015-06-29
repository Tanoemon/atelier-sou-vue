'use strict';

var
  fs = require("fs"),
  gulp = require('gulp'),
  browserify = require("browserify"),
  babelify = require("babelify"),
  notifier = require('node-notifier'),
  source = require('vinyl-source-stream'),
  path = require('path'),
  del = require('del'),
  vueify = require('vueify'),
  uglifyify = require('uglifyify'),
  through2 = require('through2');


gulp.task('makeJs', () => {
  browserify('src/entry.js', {
      debug: false
    })
    .transform(babelify)
    .transform(vueify)
    .transform({
      global: true
    }, uglifyify)
    .bundle()
    .on('error', handleErrors)
    .pipe(source('bundle.js'))
    .pipe(gulp.dest('../dist/js'))
    .on('end', () => {
      console.log('end')
    });
});

gulp.task('makeView', () => {
  del(['../dist/index.html', '../dist/favicon.ico'], {
    force: true
  }, () => {
    gulp.src(['src/index.html', 'src/favicon.ico']).pipe(gulp.dest('../dist'));
  });
});

gulp.task('makeImages', () => {
  del(['../dist/images'], {
    force: true
  }, () => {
    gulp.src('src/app/images/*').pipe(gulp.dest('../dist/images'));
  });
});

gulp.task('watch', () => {
  gulp.watch(["src/entry.js", "src/app/**/*"], ['makeJs']);
  gulp.watch(["src/app/images/*"], ['makeImages']);
  gulp.watch(["src/index.html", "src/favicon.co"], ['makeView']);
});

gulp.task('default', ['watch', 'makeJs', 'makeImages', 'makeView']);

function handleErrors() {
  var em = arguments[0];

  console.log(em.message, em.stack);

  new notifier.Growl().notify({
    title: "Compile Error",
    message: em.message,
    icon: path.join(__dirname, 'error.png')
  });

  this.emit('end');
}
