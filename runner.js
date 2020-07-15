const child = require('child_process')
const process = child.exec("./backend --DEBUG --PORT 2096");
process.stdout.on("data", (content) => console.log(content));
process.stderr.on("data", (content) => console.log(content));