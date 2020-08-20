
cirruHtml = require('cirru-html')
fs = require('fs')
gaze = require('gaze')

data = require('./data')

compile = () ->
  code = fs.readFileSync("template.cirru", 'utf8')

  htmlString = cirruHtml.render(code, data)

  fs.writeFileSync('index.html', htmlString)

  console.log("Wrote to index.html")

compile()

gaze "template.cirru", (err, watcher) ->

  console.log "watching template..."

  watcher.on "changed", () ->
    try
      compile()
    catch error
      console.error error
