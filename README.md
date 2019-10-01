# vugu-starter-project

This repository can be used as a template to start a new
[vugu](https://vugu.org) project with development mode and to generate a
production-ready binary server using
[packr](https://github.com/gobuffalo/packr/tree/master/v2) to bundle static
assets.


## Using

### development

Run `make dev` and open http://localhost:8877.

If you make any changes at `root.vugu` or `root-data.go` reload the page to
reflect these changes.

### Distribution

Run `make dist` and you will have a binary file ready to deploy in production
at `bin/server`. Run `bin/server` and open http://localhost:8877.
